package ipam

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"golang.org/x/net/context"

	"github.com/hatena/ipdrawer/gen/go/model"
	"github.com/hatena/ipdrawer/pkg/storage"
	nu "github.com/hatena/ipdrawer/pkg/utils/netutil"
)

const stressTestNS = "stress-test-ns"

// newRealRedis creates a connection to a real Redis instance.
// Set REDIS_ADDR environment variable (default: localhost:6379)
func newRealRedis(t *testing.T) (*storage.Redis, func()) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		DB:         15, // Use DB 15 for stress tests to avoid conflicts
		MaxRetries: 4,
	})

	_, err := client.Ping().Result()
	if err != nil {
		t.Skipf("Skipping stress test: Redis not available at %s: %v", addr, err)
	}

	// Clean up any existing test data
	client.FlushDB()

	return &storage.Redis{Client: client}, func() {
		client.FlushDB()
		client.Close()
	}
}

// getMemStats returns current memory statistics
func getMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m
}

// formatBytes formats bytes into human readable string
func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// prePopulatePool simulates a pool with a given percentage of IPs already allocated.
// It uses Redis pipeline for efficient bulk insertion.
func prePopulatePool(t *testing.T, r *storage.Redis, pool *model.Pool, allocPercent float64) int {
	s := net.ParseIP(pool.Start)
	e := net.ParseIP(pool.End)

	startInt := nu.IP2Uint(s)
	endInt := nu.IP2Uint(e)
	totalIPs := int(endInt - startInt + 1)
	allocCount := int(float64(totalIPs) * allocPercent / 100.0)

	t.Logf("Pool size: %d IPs, pre-allocating %d IPs (%.1f%%)", totalIPs, allocCount, allocPercent)

	// Generate random IPs to allocate (simulating real-world scattered allocation)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	allocatedSet := make(map[uint32]bool, allocCount)

	for len(allocatedSet) < allocCount {
		offset := uint32(rnd.Intn(totalIPs))
		ipInt := startInt + offset
		allocatedSet[ipInt] = true
	}

	// Bulk insert using pipeline
	zkey := makePoolUsedIPZSet(stressTestNS, s, e)
	batchSize := 10000
	pipe := r.Client.Pipeline()
	count := 0

	for ipInt := range allocatedSet {
		ip := nu.Int2IP(ipInt)
		z := redis.Z{
			Score:  float64(ipInt),
			Member: ip.String(),
		}
		pipe.ZAdd(zkey, z)
		count++

		if count%batchSize == 0 {
			_, err := pipe.Exec()
			if err != nil {
				t.Fatalf("Failed to execute pipeline: %v", err)
			}
			pipe = r.Client.Pipeline()
			if count%(batchSize*10) == 0 {
				t.Logf("Pre-populated %d/%d IPs...", count, allocCount)
			}
		}
	}

	// Execute remaining
	if count%batchSize != 0 {
		_, err := pipe.Exec()
		if err != nil {
			t.Fatalf("Failed to execute final pipeline: %v", err)
		}
	}

	t.Logf("Pre-population complete: %d IPs allocated", allocCount)
	return totalIPs - allocCount // Return number of free IPs
}

// TestStressLargePoolSequential tests sequential allocation on a large, mostly-full pool
func TestStressLargePoolSequential(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	r, cleanup := newRealRedis(t)
	defer cleanup()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// Test with /16 pool (65536 IPs) at 99% allocation
	// A /8 pool (16M IPs) would take too long to pre-populate and test
	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.0.255.254",
	}

	memBefore := getMemStats()
	t.Logf("Memory before pre-population: Alloc=%s, Sys=%s",
		formatBytes(memBefore.Alloc), formatBytes(memBefore.Sys))

	freeIPs := prePopulatePool(t, r, pool, 99.0)
	t.Logf("Free IPs remaining: %d", freeIPs)

	memAfterPopulate := getMemStats()
	t.Logf("Memory after pre-population: Alloc=%s, Sys=%s",
		formatBytes(memAfterPopulate.Alloc), formatBytes(memAfterPopulate.Sys))

	// Get Redis memory info
	info, _ := r.Client.Info("memory").Result()
	t.Logf("Redis memory info (excerpt): %s", info[:min(500, len(info))])

	// Test sequential allocation - measure time for multiple allocations
	numAllocations := min(100, freeIPs)
	times := make([]time.Duration, numAllocations)

	t.Logf("Starting %d sequential allocations...", numAllocations)

	for i := 0; i < numAllocations; i++ {
		uuid := fmt.Sprintf("stress-uuid-%d", i)
		start := time.Now()
		ip, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", false /* sequential */, true, false, false, nil)
		elapsed := time.Since(start)
		times[i] = elapsed

		if err != nil {
			t.Fatalf("DrawIP failed at iteration %d: %v", i, err)
		}

		if i < 5 || i%20 == 0 {
			t.Logf("Allocation %d: IP=%s, Time=%v", i, ip.String(), elapsed)
		}
	}

	// Calculate statistics
	var totalTime time.Duration
	var maxTime time.Duration
	minTime := times[0]
	for _, d := range times {
		totalTime += d
		if d > maxTime {
			maxTime = d
		}
		if d < minTime {
			minTime = d
		}
	}
	avgTime := totalTime / time.Duration(numAllocations)

	memAfter := getMemStats()

	t.Logf("\n=== Sequential Allocation Results ===")
	t.Logf("Pool: %s - %s (99%% allocated)", pool.Start, pool.End)
	t.Logf("Allocations: %d", numAllocations)
	t.Logf("Total time: %v", totalTime)
	t.Logf("Avg time: %v", avgTime)
	t.Logf("Min time: %v", minTime)
	t.Logf("Max time: %v", maxTime)
	t.Logf("Memory after test: Alloc=%s, Sys=%s",
		formatBytes(memAfter.Alloc), formatBytes(memAfter.Sys))

	// Performance threshold: avg should be under 100ms for /16 pool
	if avgTime > 100*time.Millisecond {
		t.Errorf("Sequential allocation too slow: avg=%v (threshold: 100ms)", avgTime)
	}
}

// TestStressLargePoolRandom tests random allocation on a large, mostly-full pool
func TestStressLargePoolRandom(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	r, cleanup := newRealRedis(t)
	defer cleanup()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// Test with /16 pool (65536 IPs) at 99% allocation
	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.0.255.254",
	}

	memBefore := getMemStats()
	t.Logf("Memory before pre-population: Alloc=%s, Sys=%s",
		formatBytes(memBefore.Alloc), formatBytes(memBefore.Sys))

	freeIPs := prePopulatePool(t, r, pool, 99.0)
	t.Logf("Free IPs remaining: %d", freeIPs)

	memAfterPopulate := getMemStats()
	t.Logf("Memory after pre-population: Alloc=%s, Sys=%s",
		formatBytes(memAfterPopulate.Alloc), formatBytes(memAfterPopulate.Sys))

	// Test random allocation
	numAllocations := min(100, freeIPs)
	times := make([]time.Duration, numAllocations)

	t.Logf("Starting %d random allocations...", numAllocations)

	for i := 0; i < numAllocations; i++ {
		uuid := fmt.Sprintf("stress-uuid-random-%d", i)
		start := time.Now()
		ip, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", true /* random */, true, false, false, nil)
		elapsed := time.Since(start)
		times[i] = elapsed

		if err != nil {
			t.Fatalf("DrawIP failed at iteration %d: %v", i, err)
		}

		if i < 5 || i%20 == 0 {
			t.Logf("Allocation %d: IP=%s, Time=%v", i, ip.String(), elapsed)
		}
	}

	// Calculate statistics
	var totalTime time.Duration
	var maxTime time.Duration
	minTime := times[0]
	for _, d := range times {
		totalTime += d
		if d > maxTime {
			maxTime = d
		}
		if d < minTime {
			minTime = d
		}
	}
	avgTime := totalTime / time.Duration(numAllocations)

	memAfter := getMemStats()

	t.Logf("\n=== Random Allocation Results ===")
	t.Logf("Pool: %s - %s (99%% allocated)", pool.Start, pool.End)
	t.Logf("Allocations: %d", numAllocations)
	t.Logf("Total time: %v", totalTime)
	t.Logf("Avg time: %v", avgTime)
	t.Logf("Min time: %v", minTime)
	t.Logf("Max time: %v", maxTime)
	t.Logf("Memory after test: Alloc=%s, Sys=%s",
		formatBytes(memAfter.Alloc), formatBytes(memAfter.Sys))

	// Performance threshold
	if avgTime > 100*time.Millisecond {
		t.Errorf("Random allocation too slow: avg=%v (threshold: 100ms)", avgTime)
	}
}

// TestStressProductionPool tests a /10 pool (4M IPs) - production scale
func TestStressProductionPool(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	// This test requires explicit opt-in due to time and resources needed
	if os.Getenv("RUN_PRODUCTION_POOL_TEST") == "" {
		t.Skip("Skipping production pool test. Set RUN_PRODUCTION_POOL_TEST=1 to run")
	}

	r, cleanup := newRealRedis(t)
	defer cleanup()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// /10 pool = 4,194,304 IPs (10.0.0.0 - 10.63.255.255)
	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.63.255.254",
	}

	s := net.ParseIP(pool.Start)
	e := net.ParseIP(pool.End)
	startInt := nu.IP2Uint(s)
	endInt := nu.IP2Uint(e)
	totalIPs := int(endInt - startInt + 1)

	t.Logf("=== /10 Pool Production Scale Test ===")
	t.Logf("Pool: %s - %s", pool.Start, pool.End)
	t.Logf("Total IPs: %d (~4M)", totalIPs)

	memBefore := getMemStats()
	t.Logf("Memory before pre-population: Alloc=%s, Sys=%s",
		formatBytes(memBefore.Alloc), formatBytes(memBefore.Sys))

	// Test at 99% allocation - this is the worst case scenario
	fillPercent := 99.0
	freeIPs := prePopulatePool(t, r, pool, fillPercent)
	t.Logf("Free IPs remaining: %d (%.1f%% free)", freeIPs, 100-fillPercent)

	// Get Redis memory usage
	memInfo, _ := r.Client.Info("memory").Result()
	// Extract just the used_memory line
	for _, line := range strings.Split(memInfo, "\n") {
		if strings.HasPrefix(line, "used_memory_human:") {
			t.Logf("Redis %s", strings.TrimSpace(line))
		}
	}

	memAfterPopulate := getMemStats()
	t.Logf("Go Memory after pre-population: Alloc=%s, Sys=%s",
		formatBytes(memAfterPopulate.Alloc), formatBytes(memAfterPopulate.Sys))

	// Test allocations
	numAllocations := 20

	t.Logf("\n--- Testing Sequential Allocation (99%% full) ---")
	seqTimes := make([]time.Duration, numAllocations)
	for i := 0; i < numAllocations; i++ {
		uuid := fmt.Sprintf("prod-seq-%d", i)
		start := time.Now()
		ip, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", false, true, false, false, nil)
		seqTimes[i] = time.Since(start)
		if err != nil {
			t.Fatalf("Sequential DrawIP failed at %d: %v", i, err)
		}
		t.Logf("Seq %d: IP=%s, Time=%v", i, ip.String(), seqTimes[i])
	}

	// Calculate sequential stats
	var seqTotal time.Duration
	seqMax := seqTimes[0]
	seqMin := seqTimes[0]
	for _, d := range seqTimes {
		seqTotal += d
		if d > seqMax {
			seqMax = d
		}
		if d < seqMin {
			seqMin = d
		}
	}

	t.Logf("\n--- Testing Random Allocation (99%% full) ---")
	// Reset for random test
	r.Client.FlushDB()
	prePopulatePool(t, r, pool, fillPercent)

	randTimes := make([]time.Duration, numAllocations)
	for i := 0; i < numAllocations; i++ {
		uuid := fmt.Sprintf("prod-rand-%d", i)
		start := time.Now()
		ip, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", true, true, false, false, nil)
		randTimes[i] = time.Since(start)
		if err != nil {
			t.Fatalf("Random DrawIP failed at %d: %v", i, err)
		}
		t.Logf("Rand %d: IP=%s, Time=%v", i, ip.String(), randTimes[i])
	}

	// Calculate random stats
	var randTotal time.Duration
	randMax := randTimes[0]
	randMin := randTimes[0]
	for _, d := range randTimes {
		randTotal += d
		if d > randMax {
			randMax = d
		}
		if d < randMin {
			randMin = d
		}
	}

	memAfter := getMemStats()

	t.Logf("\n=== /10 Pool (4M IPs, 99%% allocated) Results ===")
	t.Logf("Sequential: Total=%v, Avg=%v, Min=%v, Max=%v",
		seqTotal, seqTotal/time.Duration(numAllocations), seqMin, seqMax)
	t.Logf("Random: Total=%v, Avg=%v, Min=%v, Max=%v",
		randTotal, randTotal/time.Duration(numAllocations), randMin, randMax)
	t.Logf("Memory after test: Alloc=%s, Sys=%s",
		formatBytes(memAfter.Alloc), formatBytes(memAfter.Sys))

	// Warn if too slow
	seqAvg := seqTotal / time.Duration(numAllocations)
	randAvg := randTotal / time.Duration(numAllocations)
	if seqAvg > 1*time.Second {
		t.Logf("WARNING: Sequential allocation is very slow: avg=%v", seqAvg)
	}
	if randAvg > 1*time.Second {
		t.Logf("WARNING: Random allocation is very slow: avg=%v", randAvg)
	}
}

// TestStressVeryLargePool tests a /12 pool (1M IPs) - use sparingly
func TestStressVeryLargePool(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	// This test requires explicit opt-in
	if os.Getenv("RUN_VERY_LARGE_POOL_TEST") == "" {
		t.Skip("Skipping very large pool test. Set RUN_VERY_LARGE_POOL_TEST=1 to run")
	}

	r, cleanup := newRealRedis(t)
	defer cleanup()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// /12 pool = 1,048,576 IPs
	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.15.255.254",
	}

	memBefore := getMemStats()
	t.Logf("Memory before pre-population: Alloc=%s, Sys=%s",
		formatBytes(memBefore.Alloc), formatBytes(memBefore.Sys))

	// Only 95% allocation for very large pool to keep test time reasonable
	freeIPs := prePopulatePool(t, r, pool, 95.0)
	t.Logf("Free IPs remaining: %d", freeIPs)

	// Get Redis memory usage
	memInfo, _ := r.Client.Info("memory").Result()
	t.Logf("Redis memory info:\n%s", memInfo)

	memAfterPopulate := getMemStats()
	t.Logf("Go Memory after pre-population: Alloc=%s, Sys=%s",
		formatBytes(memAfterPopulate.Alloc), formatBytes(memAfterPopulate.Sys))

	// Test allocations
	numAllocations := 50
	
	t.Logf("\n--- Testing Sequential Allocation ---")
	seqTimes := make([]time.Duration, numAllocations)
	for i := 0; i < numAllocations; i++ {
		uuid := fmt.Sprintf("stress-large-seq-%d", i)
		start := time.Now()
		ip, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", false, true, false, false, nil)
		seqTimes[i] = time.Since(start)
		if err != nil {
			t.Fatalf("Sequential DrawIP failed at %d: %v", i, err)
		}
		if i < 3 || i%10 == 0 {
			t.Logf("Seq %d: IP=%s, Time=%v", i, ip.String(), seqTimes[i])
		}
	}

	// Clean up for random test
	r.Client.FlushDB()
	freeIPs = prePopulatePool(t, r, pool, 95.0)

	t.Logf("\n--- Testing Random Allocation ---")
	randTimes := make([]time.Duration, numAllocations)
	for i := 0; i < numAllocations; i++ {
		uuid := fmt.Sprintf("stress-large-rand-%d", i)
		start := time.Now()
		ip, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", true, true, false, false, nil)
		randTimes[i] = time.Since(start)
		if err != nil {
			t.Fatalf("Random DrawIP failed at %d: %v", i, err)
		}
		if i < 3 || i%10 == 0 {
			t.Logf("Rand %d: IP=%s, Time=%v", i, ip.String(), randTimes[i])
		}
	}

	// Report
	var seqTotal, randTotal time.Duration
	for i := 0; i < numAllocations; i++ {
		seqTotal += seqTimes[i]
		randTotal += randTimes[i]
	}

	t.Logf("\n=== Very Large Pool (/12, 1M IPs, 95%% allocated) Results ===")
	t.Logf("Sequential: Total=%v, Avg=%v", seqTotal, seqTotal/time.Duration(numAllocations))
	t.Logf("Random: Total=%v, Avg=%v", randTotal, randTotal/time.Duration(numAllocations))
}

// TestStressCompareAlgorithms compares sequential vs random at different fill levels
func TestStressCompareAlgorithms(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	r, cleanup := newRealRedis(t)
	defer cleanup()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// /18 pool = 16,384 IPs - good balance of size and test speed
	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.0.63.254",
	}

	fillLevels := []float64{50.0, 75.0, 90.0, 95.0, 99.0}
	numAllocations := 20

	type result struct {
		fillLevel float64
		seqAvg    time.Duration
		randAvg   time.Duration
	}
	results := make([]result, len(fillLevels))

	for idx, fillLevel := range fillLevels {
		r.Client.FlushDB()

		freeIPs := prePopulatePool(t, r, pool, fillLevel)
		allocations := min(numAllocations, freeIPs/2) // Leave room for both tests

		// Sequential test
		seqTimes := make([]time.Duration, allocations)
		for i := 0; i < allocations; i++ {
			uuid := fmt.Sprintf("cmp-seq-%d-%d", int(fillLevel), i)
			start := time.Now()
			_, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", false, true, false, false, nil)
			seqTimes[i] = time.Since(start)
			if err != nil {
				t.Logf("Sequential allocation failed at fill=%.0f%%, iter=%d: %v", fillLevel, i, err)
				break
			}
		}

		// Reset for random test
		r.Client.FlushDB()
		prePopulatePool(t, r, pool, fillLevel)

		// Random test
		randTimes := make([]time.Duration, allocations)
		for i := 0; i < allocations; i++ {
			uuid := fmt.Sprintf("cmp-rand-%d-%d", int(fillLevel), i)
			start := time.Now()
			_, err := m.DrawIP(ctx, stressTestNS, pool, uuid, "", true, true, false, false, nil)
			randTimes[i] = time.Since(start)
			if err != nil {
				t.Logf("Random allocation failed at fill=%.0f%%, iter=%d: %v", fillLevel, i, err)
				break
			}
		}

		var seqTotal, randTotal time.Duration
		for i := 0; i < allocations; i++ {
			seqTotal += seqTimes[i]
			randTotal += randTimes[i]
		}

		results[idx] = result{
			fillLevel: fillLevel,
			seqAvg:    seqTotal / time.Duration(allocations),
			randAvg:   randTotal / time.Duration(allocations),
		}

		t.Logf("Fill %.0f%%: Sequential avg=%v, Random avg=%v",
			fillLevel, results[idx].seqAvg, results[idx].randAvg)
	}

	t.Logf("\n=== Algorithm Comparison Summary ===")
	t.Logf("Pool: /18 (16,384 IPs)")
	t.Logf("%-10s %-15s %-15s %-10s", "Fill %", "Sequential", "Random", "Winner")
	for _, r := range results {
		winner := "Sequential"
		if r.randAvg < r.seqAvg {
			winner = "Random"
		}
		t.Logf("%-10.0f %-15v %-15v %-10s", r.fillLevel, r.seqAvg, r.randAvg, winner)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
