# ModelPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ModelPoolStatus**](ModelPoolStatus.md) |  | [optional] [default to ModelPoolStatus_0]
**Tags** | Pointer to [**[]ModelTag**](ModelTag.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewModelPool

`func NewModelPool() *ModelPool`

NewModelPool instantiates a new ModelPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPoolWithDefaults

`func NewModelPoolWithDefaults() *ModelPool`

NewModelPoolWithDefaults instantiates a new ModelPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ModelPool) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ModelPool) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ModelPool) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *ModelPool) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ModelPool) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ModelPool) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ModelPool) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ModelPool) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStatus

`func (o *ModelPool) GetStatus() ModelPoolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelPool) GetStatusOk() (*ModelPoolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelPool) SetStatus(v ModelPoolStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ModelPool) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelPool) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelPool) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelPool) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelPool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelPool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelPool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelPool) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *ModelPool) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *ModelPool) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *ModelPool) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *ModelPool) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetNamespace

`func (o *ModelPool) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ModelPool) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ModelPool) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ModelPool) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


