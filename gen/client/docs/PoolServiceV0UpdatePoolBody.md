# PoolServiceV0UpdatePoolBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ModelPoolStatus**](ModelPoolStatus.md) |  | [optional] [default to ModelPoolStatus_0]
**Tags** | Pointer to [**[]ModelTag**](ModelTag.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPoolServiceV0UpdatePoolBody

`func NewPoolServiceV0UpdatePoolBody() *PoolServiceV0UpdatePoolBody`

NewPoolServiceV0UpdatePoolBody instantiates a new PoolServiceV0UpdatePoolBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolServiceV0UpdatePoolBodyWithDefaults

`func NewPoolServiceV0UpdatePoolBodyWithDefaults() *PoolServiceV0UpdatePoolBody`

NewPoolServiceV0UpdatePoolBodyWithDefaults instantiates a new PoolServiceV0UpdatePoolBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PoolServiceV0UpdatePoolBody) GetStatus() ModelPoolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolServiceV0UpdatePoolBody) GetStatusOk() (*ModelPoolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolServiceV0UpdatePoolBody) SetStatus(v ModelPoolStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoolServiceV0UpdatePoolBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PoolServiceV0UpdatePoolBody) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PoolServiceV0UpdatePoolBody) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PoolServiceV0UpdatePoolBody) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PoolServiceV0UpdatePoolBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PoolServiceV0UpdatePoolBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoolServiceV0UpdatePoolBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoolServiceV0UpdatePoolBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PoolServiceV0UpdatePoolBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *PoolServiceV0UpdatePoolBody) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *PoolServiceV0UpdatePoolBody) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *PoolServiceV0UpdatePoolBody) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *PoolServiceV0UpdatePoolBody) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


