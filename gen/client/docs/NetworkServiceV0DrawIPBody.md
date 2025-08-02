# NetworkServiceV0DrawIPBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IP** | Pointer to **string** | The ip address wanted or the network to draw ip from. Optional. | [optional] 
**Mask** | Pointer to **int32** |  | [optional] 
**PoolTag** | Pointer to [**ModelTag**](ModelTag.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TemporaryReserved** | Pointer to **bool** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**MustHaveWantIP** | Pointer to **bool** |  | [optional] 
**Sequential** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkServiceV0DrawIPBody

`func NewNetworkServiceV0DrawIPBody() *NetworkServiceV0DrawIPBody`

NewNetworkServiceV0DrawIPBody instantiates a new NetworkServiceV0DrawIPBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServiceV0DrawIPBodyWithDefaults

`func NewNetworkServiceV0DrawIPBodyWithDefaults() *NetworkServiceV0DrawIPBody`

NewNetworkServiceV0DrawIPBodyWithDefaults instantiates a new NetworkServiceV0DrawIPBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIP

`func (o *NetworkServiceV0DrawIPBody) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *NetworkServiceV0DrawIPBody) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *NetworkServiceV0DrawIPBody) SetIP(v string)`

SetIP sets IP field to given value.

### HasIP

`func (o *NetworkServiceV0DrawIPBody) HasIP() bool`

HasIP returns a boolean if a field has been set.

### GetMask

`func (o *NetworkServiceV0DrawIPBody) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *NetworkServiceV0DrawIPBody) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *NetworkServiceV0DrawIPBody) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *NetworkServiceV0DrawIPBody) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetPoolTag

`func (o *NetworkServiceV0DrawIPBody) GetPoolTag() ModelTag`

GetPoolTag returns the PoolTag field if non-nil, zero value otherwise.

### GetPoolTagOk

`func (o *NetworkServiceV0DrawIPBody) GetPoolTagOk() (*ModelTag, bool)`

GetPoolTagOk returns a tuple with the PoolTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTag

`func (o *NetworkServiceV0DrawIPBody) SetPoolTag(v ModelTag)`

SetPoolTag sets PoolTag field to given value.

### HasPoolTag

`func (o *NetworkServiceV0DrawIPBody) HasPoolTag() bool`

HasPoolTag returns a boolean if a field has been set.

### GetName

`func (o *NetworkServiceV0DrawIPBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkServiceV0DrawIPBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkServiceV0DrawIPBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkServiceV0DrawIPBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemporaryReserved

`func (o *NetworkServiceV0DrawIPBody) GetTemporaryReserved() bool`

GetTemporaryReserved returns the TemporaryReserved field if non-nil, zero value otherwise.

### GetTemporaryReservedOk

`func (o *NetworkServiceV0DrawIPBody) GetTemporaryReservedOk() (*bool, bool)`

GetTemporaryReservedOk returns a tuple with the TemporaryReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryReserved

`func (o *NetworkServiceV0DrawIPBody) SetTemporaryReserved(v bool)`

SetTemporaryReserved sets TemporaryReserved field to given value.

### HasTemporaryReserved

`func (o *NetworkServiceV0DrawIPBody) HasTemporaryReserved() bool`

HasTemporaryReserved returns a boolean if a field has been set.

### GetUUID

`func (o *NetworkServiceV0DrawIPBody) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *NetworkServiceV0DrawIPBody) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *NetworkServiceV0DrawIPBody) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *NetworkServiceV0DrawIPBody) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetMustHaveWantIP

`func (o *NetworkServiceV0DrawIPBody) GetMustHaveWantIP() bool`

GetMustHaveWantIP returns the MustHaveWantIP field if non-nil, zero value otherwise.

### GetMustHaveWantIPOk

`func (o *NetworkServiceV0DrawIPBody) GetMustHaveWantIPOk() (*bool, bool)`

GetMustHaveWantIPOk returns a tuple with the MustHaveWantIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustHaveWantIP

`func (o *NetworkServiceV0DrawIPBody) SetMustHaveWantIP(v bool)`

SetMustHaveWantIP sets MustHaveWantIP field to given value.

### HasMustHaveWantIP

`func (o *NetworkServiceV0DrawIPBody) HasMustHaveWantIP() bool`

HasMustHaveWantIP returns a boolean if a field has been set.

### GetSequential

`func (o *NetworkServiceV0DrawIPBody) GetSequential() bool`

GetSequential returns the Sequential field if non-nil, zero value otherwise.

### GetSequentialOk

`func (o *NetworkServiceV0DrawIPBody) GetSequentialOk() (*bool, bool)`

GetSequentialOk returns a tuple with the Sequential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequential

`func (o *NetworkServiceV0DrawIPBody) SetSequential(v bool)`

SetSequential sets Sequential field to given value.

### HasSequential

`func (o *NetworkServiceV0DrawIPBody) HasSequential() bool`

HasSequential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


