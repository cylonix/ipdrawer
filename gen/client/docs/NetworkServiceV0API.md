# \NetworkServiceV0API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkServiceV0CreateNetwork**](NetworkServiceV0API.md#NetworkServiceV0CreateNetwork) | **Post** /api/v0/{namespace}/network/{ip}/{mask}/create | 
[**NetworkServiceV0CreatePool**](NetworkServiceV0API.md#NetworkServiceV0CreatePool) | **Post** /api/v0/{namespace}/network/{ip}/{mask}/pool/create | 
[**NetworkServiceV0DeleteNetwork**](NetworkServiceV0API.md#NetworkServiceV0DeleteNetwork) | **Post** /api/v0/{namespace}/network/{ip}/{mask}/delete | 
[**NetworkServiceV0DrawIP**](NetworkServiceV0API.md#NetworkServiceV0DrawIP) | **Post** /api/v0/{namespace}/network/{ip}/{mask}/drawip | 
[**NetworkServiceV0DrawIP2**](NetworkServiceV0API.md#NetworkServiceV0DrawIP2) | **Post** /api/v0/{namespace}/network/{name}/drawip | 
[**NetworkServiceV0DrawIP3**](NetworkServiceV0API.md#NetworkServiceV0DrawIP3) | **Post** /api/v0/{namespace}/pool/{rangeStart}/{rangeEnd}/drawip | 
[**NetworkServiceV0DrawIPEstimatingNetwork**](NetworkServiceV0API.md#NetworkServiceV0DrawIPEstimatingNetwork) | **Get** /api/v0/{namespace}/drawip | 
[**NetworkServiceV0GetEstimatedNetwork**](NetworkServiceV0API.md#NetworkServiceV0GetEstimatedNetwork) | **Get** /api/v0/{namespace}/network | 
[**NetworkServiceV0GetNetwork**](NetworkServiceV0API.md#NetworkServiceV0GetNetwork) | **Get** /api/v0/{namespace}/network/{ip}/{mask} | 
[**NetworkServiceV0GetNetwork2**](NetworkServiceV0API.md#NetworkServiceV0GetNetwork2) | **Get** /api/v0/{namespace}/network/{name} | 
[**NetworkServiceV0GetPoolsInNetwork**](NetworkServiceV0API.md#NetworkServiceV0GetPoolsInNetwork) | **Get** /api/v0/{namespace}/network/{ip}/{mask}/pools | 
[**NetworkServiceV0ListNetwork**](NetworkServiceV0API.md#NetworkServiceV0ListNetwork) | **Get** /api/v0/{namespace}/networks | 
[**NetworkServiceV0UpdateNetwork**](NetworkServiceV0API.md#NetworkServiceV0UpdateNetwork) | **Post** /api/v0/{namespace}/network/update | 



## NetworkServiceV0CreateNetwork

> map[string]interface{} NetworkServiceV0CreateNetwork(ctx, namespace, ip, mask).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	ip := "ip_example" // string | 
	mask := int32(56) // int32 | 
	body := *openapiclient.NewNetworkServiceV0CreateNetworkBody() // NetworkServiceV0CreateNetworkBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0CreateNetwork(context.Background(), namespace, ip, mask).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0CreateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0CreateNetwork`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0CreateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 
**mask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0CreateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**NetworkServiceV0CreateNetworkBody**](NetworkServiceV0CreateNetworkBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0CreatePool

> map[string]interface{} NetworkServiceV0CreatePool(ctx, namespace, ip, mask).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	ip := "ip_example" // string | 
	mask := int32(56) // int32 | 
	body := *openapiclient.NewNetworkServiceV0CreatePoolBody() // NetworkServiceV0CreatePoolBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0CreatePool(context.Background(), namespace, ip, mask).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0CreatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0CreatePool`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0CreatePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 
**mask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0CreatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**NetworkServiceV0CreatePoolBody**](NetworkServiceV0CreatePoolBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0DeleteNetwork

> map[string]interface{} NetworkServiceV0DeleteNetwork(ctx, namespace, ip, mask).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	ip := "ip_example" // string | 
	mask := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0DeleteNetwork(context.Background(), namespace, ip, mask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0DeleteNetwork`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 
**mask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0DeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0DrawIP

> ServerpbDrawIPResponse NetworkServiceV0DrawIP(ctx, namespace, ip, mask).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	ip := "ip_example" // string | The ip address wanted or the network to draw ip from. Optional.
	mask := int32(56) // int32 | 
	body := *openapiclient.NewNetworkServiceV0DrawIPBody() // NetworkServiceV0DrawIPBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0DrawIP(context.Background(), namespace, ip, mask).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0DrawIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0DrawIP`: ServerpbDrawIPResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0DrawIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** | The ip address wanted or the network to draw ip from. Optional. | 
**mask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0DrawIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**NetworkServiceV0DrawIPBody**](NetworkServiceV0DrawIPBody.md) |  | 

### Return type

[**ServerpbDrawIPResponse**](ServerpbDrawIPResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0DrawIP2

> ServerpbDrawIPResponse NetworkServiceV0DrawIP2(ctx, namespace, name).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	name := "name_example" // string | 
	body := *openapiclient.NewNetworkServiceV0DrawIPBody() // NetworkServiceV0DrawIPBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0DrawIP2(context.Background(), namespace, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0DrawIP2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0DrawIP2`: ServerpbDrawIPResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0DrawIP2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0DrawIP2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NetworkServiceV0DrawIPBody**](NetworkServiceV0DrawIPBody.md) |  | 

### Return type

[**ServerpbDrawIPResponse**](ServerpbDrawIPResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0DrawIP3

> ServerpbDrawIPResponse NetworkServiceV0DrawIP3(ctx, namespace, rangeStart, rangeEnd).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	rangeStart := "rangeStart_example" // string | 
	rangeEnd := "rangeEnd_example" // string | 
	body := *openapiclient.NewNetworkServiceV0DrawIPBody() // NetworkServiceV0DrawIPBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0DrawIP3(context.Background(), namespace, rangeStart, rangeEnd).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0DrawIP3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0DrawIP3`: ServerpbDrawIPResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0DrawIP3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**rangeStart** | **string** |  | 
**rangeEnd** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0DrawIP3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**NetworkServiceV0DrawIPBody**](NetworkServiceV0DrawIPBody.md) |  | 

### Return type

[**ServerpbDrawIPResponse**](ServerpbDrawIPResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0DrawIPEstimatingNetwork

> ServerpbDrawIPResponse NetworkServiceV0DrawIPEstimatingNetwork(ctx, namespace).PoolTagKey(poolTagKey).PoolTagValue(poolTagValue).TemporaryReserved(temporaryReserved).Sequential(sequential).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	poolTagKey := "poolTagKey_example" // string |  (optional)
	poolTagValue := "poolTagValue_example" // string |  (optional)
	temporaryReserved := true // bool |  (optional)
	sequential := true // bool | Default false i.e. randomized (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0DrawIPEstimatingNetwork(context.Background(), namespace).PoolTagKey(poolTagKey).PoolTagValue(poolTagValue).TemporaryReserved(temporaryReserved).Sequential(sequential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0DrawIPEstimatingNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0DrawIPEstimatingNetwork`: ServerpbDrawIPResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0DrawIPEstimatingNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0DrawIPEstimatingNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolTagKey** | **string** |  | 
 **poolTagValue** | **string** |  | 
 **temporaryReserved** | **bool** |  | 
 **sequential** | **bool** | Default false i.e. randomized | 

### Return type

[**ServerpbDrawIPResponse**](ServerpbDrawIPResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0GetEstimatedNetwork

> ServerpbGetNetworkResponse NetworkServiceV0GetEstimatedNetwork(ctx, namespace).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0GetEstimatedNetwork(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0GetEstimatedNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0GetEstimatedNetwork`: ServerpbGetNetworkResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0GetEstimatedNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0GetEstimatedNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerpbGetNetworkResponse**](ServerpbGetNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0GetNetwork

> ServerpbGetNetworkResponse NetworkServiceV0GetNetwork(ctx, namespace, ip, mask).Name(name).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	ip := "ip_example" // string | 
	mask := int32(56) // int32 | 
	name := "name_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0GetNetwork(context.Background(), namespace, ip, mask).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0GetNetwork`: ServerpbGetNetworkResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 
**mask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0GetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **string** |  | 

### Return type

[**ServerpbGetNetworkResponse**](ServerpbGetNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0GetNetwork2

> ServerpbGetNetworkResponse NetworkServiceV0GetNetwork2(ctx, namespace, name).Ip(ip).Mask(mask).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	name := "name_example" // string | 
	ip := "ip_example" // string |  (optional)
	mask := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0GetNetwork2(context.Background(), namespace, name).Ip(ip).Mask(mask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0GetNetwork2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0GetNetwork2`: ServerpbGetNetworkResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0GetNetwork2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0GetNetwork2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ip** | **string** |  | 
 **mask** | **int32** |  | 

### Return type

[**ServerpbGetNetworkResponse**](ServerpbGetNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0GetPoolsInNetwork

> ServerpbGetPoolsInNetworkResponse NetworkServiceV0GetPoolsInNetwork(ctx, namespace, ip, mask).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	ip := "ip_example" // string | 
	mask := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0GetPoolsInNetwork(context.Background(), namespace, ip, mask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0GetPoolsInNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0GetPoolsInNetwork`: ServerpbGetPoolsInNetworkResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0GetPoolsInNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 
**mask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0GetPoolsInNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServerpbGetPoolsInNetworkResponse**](ServerpbGetPoolsInNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0ListNetwork

> ServerpbListNetworkResponse NetworkServiceV0ListNetwork(ctx, namespace).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0ListNetwork(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0ListNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0ListNetwork`: ServerpbListNetworkResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0ListNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0ListNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerpbListNetworkResponse**](ServerpbListNetworkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceV0UpdateNetwork

> map[string]interface{} NetworkServiceV0UpdateNetwork(ctx, namespace).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hatena/ipdrawer/gen/client"
)

func main() {
	namespace := "namespace_example" // string | 
	body := *openapiclient.NewNetworkServiceV0UpdateNetworkBody() // NetworkServiceV0UpdateNetworkBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkServiceV0API.NetworkServiceV0UpdateNetwork(context.Background(), namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceV0API.NetworkServiceV0UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkServiceV0UpdateNetwork`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkServiceV0API.NetworkServiceV0UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceV0UpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NetworkServiceV0UpdateNetworkBody**](NetworkServiceV0UpdateNetworkBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

