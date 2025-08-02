# NetworkServiceV0UpdateNetworkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** |  | [optional] 
**Gateways** | Pointer to **[]string** |  | [optional] 
**Broadcast** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ModelNetworkStatus**](ModelNetworkStatus.md) |  | [optional] [default to ModelNetworkStatus_0]
**Tags** | Pointer to [**[]ModelTag**](ModelTag.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNetworkServiceV0UpdateNetworkBody

`func NewNetworkServiceV0UpdateNetworkBody() *NetworkServiceV0UpdateNetworkBody`

NewNetworkServiceV0UpdateNetworkBody instantiates a new NetworkServiceV0UpdateNetworkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServiceV0UpdateNetworkBodyWithDefaults

`func NewNetworkServiceV0UpdateNetworkBodyWithDefaults() *NetworkServiceV0UpdateNetworkBody`

NewNetworkServiceV0UpdateNetworkBodyWithDefaults instantiates a new NetworkServiceV0UpdateNetworkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *NetworkServiceV0UpdateNetworkBody) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *NetworkServiceV0UpdateNetworkBody) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *NetworkServiceV0UpdateNetworkBody) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetGateways

`func (o *NetworkServiceV0UpdateNetworkBody) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *NetworkServiceV0UpdateNetworkBody) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *NetworkServiceV0UpdateNetworkBody) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetBroadcast

`func (o *NetworkServiceV0UpdateNetworkBody) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *NetworkServiceV0UpdateNetworkBody) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *NetworkServiceV0UpdateNetworkBody) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetNetmask

`func (o *NetworkServiceV0UpdateNetworkBody) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *NetworkServiceV0UpdateNetworkBody) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *NetworkServiceV0UpdateNetworkBody) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkServiceV0UpdateNetworkBody) GetStatus() ModelNetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetStatusOk() (*ModelNetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkServiceV0UpdateNetworkBody) SetStatus(v ModelNetworkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkServiceV0UpdateNetworkBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *NetworkServiceV0UpdateNetworkBody) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkServiceV0UpdateNetworkBody) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetworkServiceV0UpdateNetworkBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkServiceV0UpdateNetworkBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkServiceV0UpdateNetworkBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkServiceV0UpdateNetworkBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *NetworkServiceV0UpdateNetworkBody) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *NetworkServiceV0UpdateNetworkBody) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *NetworkServiceV0UpdateNetworkBody) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *NetworkServiceV0UpdateNetworkBody) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


