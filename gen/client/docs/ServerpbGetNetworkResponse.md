# ServerpbGetNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **string** |  | [optional] 
**DefaultGateways** | Pointer to **[]string** |  | [optional] 
**Broadcast** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]ModelTag**](ModelTag.md) |  | [optional] 

## Methods

### NewServerpbGetNetworkResponse

`func NewServerpbGetNetworkResponse() *ServerpbGetNetworkResponse`

NewServerpbGetNetworkResponse instantiates a new ServerpbGetNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerpbGetNetworkResponseWithDefaults

`func NewServerpbGetNetworkResponseWithDefaults() *ServerpbGetNetworkResponse`

NewServerpbGetNetworkResponseWithDefaults instantiates a new ServerpbGetNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ServerpbGetNetworkResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ServerpbGetNetworkResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ServerpbGetNetworkResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ServerpbGetNetworkResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDefaultGateways

`func (o *ServerpbGetNetworkResponse) GetDefaultGateways() []string`

GetDefaultGateways returns the DefaultGateways field if non-nil, zero value otherwise.

### GetDefaultGatewaysOk

`func (o *ServerpbGetNetworkResponse) GetDefaultGatewaysOk() (*[]string, bool)`

GetDefaultGatewaysOk returns a tuple with the DefaultGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateways

`func (o *ServerpbGetNetworkResponse) SetDefaultGateways(v []string)`

SetDefaultGateways sets DefaultGateways field to given value.

### HasDefaultGateways

`func (o *ServerpbGetNetworkResponse) HasDefaultGateways() bool`

HasDefaultGateways returns a boolean if a field has been set.

### GetBroadcast

`func (o *ServerpbGetNetworkResponse) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *ServerpbGetNetworkResponse) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *ServerpbGetNetworkResponse) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *ServerpbGetNetworkResponse) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetNetmask

`func (o *ServerpbGetNetworkResponse) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ServerpbGetNetworkResponse) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ServerpbGetNetworkResponse) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ServerpbGetNetworkResponse) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetTags

`func (o *ServerpbGetNetworkResponse) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerpbGetNetworkResponse) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerpbGetNetworkResponse) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerpbGetNetworkResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


