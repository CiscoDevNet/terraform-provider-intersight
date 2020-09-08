# BootIscsi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** | Port ID of the ISCSI boot device. | [optional] 
**Slot** | Pointer to **string** | The slot id of the device. Supported values are (1 - 255, \&quot;MLOM\&quot;, \&quot;L\&quot;, \&quot;L1\&quot;, \&quot;L2\&quot;, \&quot;OCP\&quot;). | [optional] 

## Methods

### NewBootIscsi

`func NewBootIscsi() *BootIscsi`

NewBootIscsi instantiates a new BootIscsi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootIscsiWithDefaults

`func NewBootIscsiWithDefaults() *BootIscsi`

NewBootIscsiWithDefaults instantiates a new BootIscsi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *BootIscsi) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BootIscsi) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BootIscsi) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *BootIscsi) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSlot

`func (o *BootIscsi) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *BootIscsi) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *BootIscsi) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *BootIscsi) HasSlot() bool`

HasSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


