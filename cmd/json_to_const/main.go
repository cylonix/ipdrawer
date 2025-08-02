package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Reads all .json files in the input folder and encodes them as string
// literals in the output file.
func main() {
	input, outputFile, packageName := os.Args[1], os.Args[2], os.Args[3]
	fmt.Printf("input : %v\n", input)
	fmt.Printf("output: %v\n", outputFile)

	fs, err := os.ReadDir(input)
	if err != nil {
		log.Fatalf("Failed to open input directory %v: %v", input, err)
	}

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to open output file %v: %v", outputFile, err)
	}

	out.Write([]byte("package " + packageName + "\n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			name := strings.TrimPrefix(f.Name(), "server.")
			name = strings.TrimSuffix(name, ".json")
			name = strings.ToUpper(name)
			out.Write([]byte("\t" + name + " = `"))
			f, _ := os.Open(os.Args[1] + "/" + f.Name())
			fmt.Printf("Writing %v content to output file\n", f.Name())
			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}
