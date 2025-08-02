# ModelNetwork

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
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewModelNetwork

`func NewModelNetwork() *ModelNetwork`

NewModelNetwork instantiates a new ModelNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNetworkWithDefaults

`func NewModelNetworkWithDefaults() *ModelNetwork`

NewModelNetworkWithDefaults instantiates a new ModelNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *ModelNetwork) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ModelNetwork) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ModelNetwork) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ModelNetwork) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetGateways

`func (o *ModelNetwork) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ModelNetwork) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ModelNetwork) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *ModelNetwork) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetBroadcast

`func (o *ModelNetwork) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *ModelNetwork) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *ModelNetwork) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *ModelNetwork) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetNetmask

`func (o *ModelNetwork) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ModelNetwork) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ModelNetwork) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ModelNetwork) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetStatus

`func (o *ModelNetwork) GetStatus() ModelNetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelNetwork) GetStatusOk() (*ModelNetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelNetwork) SetStatus(v ModelNetworkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ModelNetwork) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelNetwork) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelNetwork) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelNetwork) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelNetwork) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *ModelNetwork) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *ModelNetwork) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *ModelNetwork) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *ModelNetwork) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetNamespace

`func (o *ModelNetwork) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ModelNetwork) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ModelNetwork) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ModelNetwork) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


