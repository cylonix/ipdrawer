# \PoolServiceV0API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoolServiceV0DeletePool**](PoolServiceV0API.md#PoolServiceV0DeletePool) | **Post** /api/v0/{namespace}/pool/{rangeStart}/{rangeEnd}/delete | 
[**PoolServiceV0GetIPInPool**](PoolServiceV0API.md#PoolServiceV0GetIPInPool) | **Get** /api/v0/{namespace}/pool/{rangeStart}/{rangeEnd}/ip | 
[**PoolServiceV0ListPool**](PoolServiceV0API.md#PoolServiceV0ListPool) | **Get** /api/v0/{namespace}/pool/list | 
[**PoolServiceV0UpdatePool**](PoolServiceV0API.md#PoolServiceV0UpdatePool) | **Post** /api/v0/{namespace}/pool/{start}/{end}/update | 



## PoolServiceV0DeletePool

> map[string]interface{} PoolServiceV0DeletePool(ctx, namespace, rangeStart, rangeEnd).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolServiceV0API.PoolServiceV0DeletePool(context.Background(), namespace, rangeStart, rangeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolServiceV0API.PoolServiceV0DeletePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoolServiceV0DeletePool`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolServiceV0API.PoolServiceV0DeletePool`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPoolServiceV0DeletePoolRequest struct via the builder pattern


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


## PoolServiceV0GetIPInPool

> ServerpbGetIPInPoolResponse PoolServiceV0GetIPInPool(ctx, namespace, rangeStart, rangeEnd).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolServiceV0API.PoolServiceV0GetIPInPool(context.Background(), namespace, rangeStart, rangeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolServiceV0API.PoolServiceV0GetIPInPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoolServiceV0GetIPInPool`: ServerpbGetIPInPoolResponse
	fmt.Fprintf(os.Stdout, "Response from `PoolServiceV0API.PoolServiceV0GetIPInPool`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPoolServiceV0GetIPInPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServerpbGetIPInPoolResponse**](ServerpbGetIPInPoolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolServiceV0ListPool

> ServerpbListPoolResponse PoolServiceV0ListPool(ctx, namespace).Execute()



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
	resp, r, err := apiClient.PoolServiceV0API.PoolServiceV0ListPool(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolServiceV0API.PoolServiceV0ListPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoolServiceV0ListPool`: ServerpbListPoolResponse
	fmt.Fprintf(os.Stdout, "Response from `PoolServiceV0API.PoolServiceV0ListPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolServiceV0ListPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerpbListPoolResponse**](ServerpbListPoolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoolServiceV0UpdatePool

> map[string]interface{} PoolServiceV0UpdatePool(ctx, namespace, start, end).Body(body).Execute()



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
	start := "start_example" // string | 
	end := "end_example" // string | 
	body := *openapiclient.NewPoolServiceV0UpdatePoolBody() // PoolServiceV0UpdatePoolBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolServiceV0API.PoolServiceV0UpdatePool(context.Background(), namespace, start, end).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolServiceV0API.PoolServiceV0UpdatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoolServiceV0UpdatePool`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolServiceV0API.PoolServiceV0UpdatePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**start** | **string** |  | 
**end** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoolServiceV0UpdatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**PoolServiceV0UpdatePoolBody**](PoolServiceV0UpdatePoolBody.md) |  | 

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

