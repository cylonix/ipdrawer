# ServerpbListIPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to [**[]ModelIPAddr**](ModelIPAddr.md) |  | [optional] 

## Methods

### NewServerpbListIPResponse

`func NewServerpbListIPResponse() *ServerpbListIPResponse`

NewServerpbListIPResponse instantiates a new ServerpbListIPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerpbListIPResponseWithDefaults

`func NewServerpbListIPResponseWithDefaults() *ServerpbListIPResponse`

NewServerpbListIPResponseWithDefaults instantiates a new ServerpbListIPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *ServerpbListIPResponse) GetIps() []ModelIPAddr`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ServerpbListIPResponse) GetIpsOk() (*[]ModelIPAddr, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ServerpbListIPResponse) SetIps(v []ModelIPAddr)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ServerpbListIPResponse) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


