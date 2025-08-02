# ServerpbGetIPInPoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool** | Pointer to [**ModelPool**](ModelPool.md) |  | [optional] 
**Ips** | Pointer to [**[]ModelIPAddr**](ModelIPAddr.md) |  | [optional] 

## Methods

### NewServerpbGetIPInPoolResponse

`func NewServerpbGetIPInPoolResponse() *ServerpbGetIPInPoolResponse`

NewServerpbGetIPInPoolResponse instantiates a new ServerpbGetIPInPoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerpbGetIPInPoolResponseWithDefaults

`func NewServerpbGetIPInPoolResponseWithDefaults() *ServerpbGetIPInPoolResponse`

NewServerpbGetIPInPoolResponseWithDefaults instantiates a new ServerpbGetIPInPoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool

`func (o *ServerpbGetIPInPoolResponse) GetPool() ModelPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ServerpbGetIPInPoolResponse) GetPoolOk() (*ModelPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ServerpbGetIPInPoolResponse) SetPool(v ModelPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ServerpbGetIPInPoolResponse) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetIps

`func (o *ServerpbGetIPInPoolResponse) GetIps() []ModelIPAddr`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ServerpbGetIPInPoolResponse) GetIpsOk() (*[]ModelIPAddr, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ServerpbGetIPInPoolResponse) SetIps(v []ModelIPAddr)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ServerpbGetIPInPoolResponse) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


