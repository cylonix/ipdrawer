# ModelIPAddr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IP** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ModelIPAddrStatus**](ModelIPAddrStatus.md) |  | [optional] [default to ModelIPAddrStatus_0]
**Tags** | Pointer to [**[]ModelTag**](ModelTag.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewModelIPAddr

`func NewModelIPAddr() *ModelIPAddr`

NewModelIPAddr instantiates a new ModelIPAddr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelIPAddrWithDefaults

`func NewModelIPAddrWithDefaults() *ModelIPAddr`

NewModelIPAddrWithDefaults instantiates a new ModelIPAddr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIP

`func (o *ModelIPAddr) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *ModelIPAddr) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *ModelIPAddr) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *ModelIPAddr) HasIP() bool`

HasIP returns a boolean if a field has been set.

### GetStatus

`func (o *ModelIPAddr) GetStatus() ModelIPAddrStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelIPAddr) GetStatusOk() (*ModelIPAddrStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelIPAddr) SetStatus(v ModelIPAddrStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelIPAddr) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ModelIPAddr) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelIPAddr) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelIPAddr) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelIPAddr) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelIPAddr) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelIPAddr) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelIPAddr) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelIPAddr) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *ModelIPAddr) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *ModelIPAddr) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *ModelIPAddr) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *ModelIPAddr) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetUUID

`func (o *ModelIPAddr) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *ModelIPAddr) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *ModelIPAddr) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *ModelIPAddr) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetNamespace

`func (o *ModelIPAddr) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ModelIPAddr) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ModelIPAddr) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ModelIPAddr) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


