# IPServiceV0UpdateIPBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ModelIPAddrStatus**](ModelIPAddrStatus.md) |  | [optional] [default to ModelIPAddrStatus_0]
**Tags** | Pointer to [**[]ModelTag**](ModelTag.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewIPServiceV0UpdateIPBody

`func NewIPServiceV0UpdateIPBody() *IPServiceV0UpdateIPBody`

NewIPServiceV0UpdateIPBody instantiates a new IPServiceV0UpdateIPBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPServiceV0UpdateIPBodyWithDefaults

`func NewIPServiceV0UpdateIPBodyWithDefaults() *IPServiceV0UpdateIPBody`

NewIPServiceV0UpdateIPBodyWithDefaults instantiates a new IPServiceV0UpdateIPBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IPServiceV0UpdateIPBody) GetStatus() ModelIPAddrStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPServiceV0UpdateIPBody) GetStatusOk() (*ModelIPAddrStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPServiceV0UpdateIPBody) SetStatus(v ModelIPAddrStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IPServiceV0UpdateIPBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *IPServiceV0UpdateIPBody) GetTags() []ModelTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPServiceV0UpdateIPBody) GetTagsOk() (*[]ModelTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPServiceV0UpdateIPBody) SetTags(v []ModelTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPServiceV0UpdateIPBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPServiceV0UpdateIPBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPServiceV0UpdateIPBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPServiceV0UpdateIPBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPServiceV0UpdateIPBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *IPServiceV0UpdateIPBody) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *IPServiceV0UpdateIPBody) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *IPServiceV0UpdateIPBody) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *IPServiceV0UpdateIPBody) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetUUID

`func (o *IPServiceV0UpdateIPBody) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *IPServiceV0UpdateIPBody) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *IPServiceV0UpdateIPBody) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *IPServiceV0UpdateIPBody) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


