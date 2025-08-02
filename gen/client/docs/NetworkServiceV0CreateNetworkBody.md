# NetworkServiceV0CreateNetworkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGateways** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]ModelTag**](ModelTag.md) |  | [optional] 
**Status** | Pointer to [**ModelNetworkStatus**](ModelNetworkStatus.md) |  | [optional] [default to ModelNetworkStatus_0]

## Methods

### NewNetworkServiceV0CreateNetworkBody

`func NewNetworkServiceV0CreateNetworkBody() *NetworkServiceV0CreateNetworkBody`

NewNetworkServiceV0CreateNetworkBody instantiates a new NetworkServiceV0CreateNetworkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServiceV0CreateNetworkBodyWithDefaults

`func NewNetworkServiceV0CreateNetworkBodyWithDefaults() *NetworkServiceV0CreateNetworkBody`

NewNetworkServiceV0CreateNetworkBodyWithDefaults instantiates a new NetworkServiceV0CreateNetworkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGateways

`func (o *NetworkServiceV0CreateNetworkBody) GetDefaultGateways() []string`

GetDefaultGateways returns the DefaultGateways field if non-nil, zero value otherwise.

### GetDefaultGatewaysOk

`func (o *NetworkServiceV0CreateNetworkBody) GetDefaultGatewaysOk() (*[]string, bool)`

GetDefaultGatewaysOk returns a tuple with the DefaultGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateways

`func (o *NetworkServiceV0CreateNetworkBody) SetDefaultGateways(v []string)`

SetDefaultGateways sets DefaultGateways field to given value.

### HasDefaultGateways

`func (o *NetworkServiceV0CreateNetworkBody) HasDefaultGateways() bool`

HasDefaultGateways returns a boolean if a field has been set.

### GetTags

`func (o *NetworkServiceV0CreateNetworkBody) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkServiceV0CreateNetworkBody) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkServiceV0CreateNetworkBody) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetworkServiceV0CreateNetworkBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkServiceV0CreateNetworkBody) GetStatus() ModelNetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkServiceV0CreateNetworkBody) GetStatusOk() (*ModelNetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkServiceV0CreateNetworkBody) SetStatus(v ModelNetworkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkServiceV0CreateNetworkBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


