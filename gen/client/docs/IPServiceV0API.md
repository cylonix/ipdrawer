# \IPServiceV0API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IPServiceV0ActivateIP**](IPServiceV0API.md#IPServiceV0ActivateIP) | **Post** /api/v0/{namespace}/ip/{ip}/activate | 
[**IPServiceV0CreateIP**](IPServiceV0API.md#IPServiceV0CreateIP) | **Post** /api/v0/{namespace}/ip/{ip}/create | 
[**IPServiceV0DeactivateIP**](IPServiceV0API.md#IPServiceV0DeactivateIP) | **Post** /api/v0/{namespace}/ip/{ip}/deactivate | 
[**IPServiceV0GetNetworkIncludingIP**](IPServiceV0API.md#IPServiceV0GetNetworkIncludingIP) | **Get** /api/v0/{namespace}/ip/{ip}/network | 
[**IPServiceV0ListIP**](IPServiceV0API.md#IPServiceV0ListIP) | **Get** /api/v0/{namespace}/ip/list | 
[**IPServiceV0ListTemporaryReservedIP**](IPServiceV0API.md#IPServiceV0ListTemporaryReservedIP) | **Get** /api/v0/{namespace}/ip/temporary_reserved/list | 
[**IPServiceV0UpdateIP**](IPServiceV0API.md#IPServiceV0UpdateIP) | **Post** /api/v0/{namespace}/ip/{ip}/update | 



## IPServiceV0ActivateIP

> map[string]interface{} IPServiceV0ActivateIP(ctx, namespace, ip).Body(body).Execute()



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
	body := *openapiclient.NewIPServiceV0ActivateIPBody() // IPServiceV0ActivateIPBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPServiceV0API.IPServiceV0ActivateIP(context.Background(), namespace, ip).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPServiceV0API.IPServiceV0ActivateIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IPServiceV0ActivateIP`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IPServiceV0API.IPServiceV0ActivateIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIPServiceV0ActivateIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IPServiceV0ActivateIPBody**](IPServiceV0ActivateIPBody.md) |  | 

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


## IPServiceV0CreateIP

> map[string]interface{} IPServiceV0CreateIP(ctx, namespace, ip).Body(body).Execute()



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
	body := *openapiclient.NewIPServiceV0CreateIPBody() // IPServiceV0CreateIPBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPServiceV0API.IPServiceV0CreateIP(context.Background(), namespace, ip).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPServiceV0API.IPServiceV0CreateIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IPServiceV0CreateIP`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IPServiceV0API.IPServiceV0CreateIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIPServiceV0CreateIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IPServiceV0CreateIPBody**](IPServiceV0CreateIPBody.md) |  | 

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


## IPServiceV0DeactivateIP

> map[string]interface{} IPServiceV0DeactivateIP(ctx, namespace, ip).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPServiceV0API.IPServiceV0DeactivateIP(context.Background(), namespace, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPServiceV0API.IPServiceV0DeactivateIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IPServiceV0DeactivateIP`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IPServiceV0API.IPServiceV0DeactivateIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIPServiceV0DeactivateIPRequest struct via the builder pattern


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


## IPServiceV0GetNetworkIncludingIP

> ServerpbGetNetworkResponse IPServiceV0GetNetworkIncludingIP(ctx, namespace, ip).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPServiceV0API.IPServiceV0GetNetworkIncludingIP(context.Background(), namespace, ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPServiceV0API.IPServiceV0GetNetworkIncludingIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IPServiceV0GetNetworkIncludingIP`: ServerpbGetNetworkResponse
	fmt.Fprintf(os.Stdout, "Response from `IPServiceV0API.IPServiceV0GetNetworkIncludingIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIPServiceV0GetNetworkIncludingIPRequest struct via the builder pattern


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


## IPServiceV0ListIP

> ServerpbListIPResponse IPServiceV0ListIP(ctx, namespace).Execute()



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
	resp, r, err := apiClient.IPServiceV0API.IPServiceV0ListIP(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPServiceV0API.IPServiceV0ListIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IPServiceV0ListIP`: ServerpbListIPResponse
	fmt.Fprintf(os.Stdout, "Response from `IPServiceV0API.IPServiceV0ListIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIPServiceV0ListIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerpbListIPResponse**](ServerpbListIPResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IPServiceV0ListTemporaryReservedIP

> ServerpbListTemporaryReservedIPResponse IPServiceV0ListTemporaryReservedIP(ctx, namespace).Execute()



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
	resp, r, err := apiClient.IPServiceV0API.IPServiceV0ListTemporaryReservedIP(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPServiceV0API.IPServiceV0ListTemporaryReservedIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IPServiceV0ListTemporaryReservedIP`: ServerpbListTemporaryReservedIPResponse
	fmt.Fprintf(os.Stdout, "Response from `IPServiceV0API.IPServiceV0ListTemporaryReservedIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIPServiceV0ListTemporaryReservedIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerpbListTemporaryReservedIPResponse**](ServerpbListTemporaryReservedIPResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IPServiceV0UpdateIP

> map[string]interface{} IPServiceV0UpdateIP(ctx, namespace, ip).Body(body).Execute()



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
	body := *openapiclient.NewIPServiceV0UpdateIPBody() // IPServiceV0UpdateIPBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPServiceV0API.IPServiceV0UpdateIP(context.Background(), namespace, ip).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPServiceV0API.IPServiceV0UpdateIP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IPServiceV0UpdateIP`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IPServiceV0API.IPServiceV0UpdateIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**ip** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIPServiceV0UpdateIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IPServiceV0UpdateIPBody**](IPServiceV0UpdateIPBody.md) |  | 

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

