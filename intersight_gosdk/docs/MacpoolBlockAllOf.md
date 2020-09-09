# MacpoolBlockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Starting address of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx. To ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use the following MAC prefix 00:25:B5:xx:xx:xx. | [optional] 
**To** | Pointer to **string** | Ending address of the block must be in hexadecimal format xx:xx:xx:xx:xx:xx. | [optional] 

## Methods

### NewMacpoolBlockAllOf

`func NewMacpoolBlockAllOf() *MacpoolBlockAllOf`

NewMacpoolBlockAllOf instantiates a new MacpoolBlockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolBlockAllOfWithDefaults

`func NewMacpoolBlockAllOfWithDefaults() *MacpoolBlockAllOf`

NewMacpoolBlockAllOfWithDefaults instantiates a new MacpoolBlockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MacpoolBlockAllOf) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MacpoolBlockAllOf) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MacpoolBlockAllOf) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MacpoolBlockAllOf) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MacpoolBlockAllOf) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MacpoolBlockAllOf) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MacpoolBlockAllOf) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MacpoolBlockAllOf) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


