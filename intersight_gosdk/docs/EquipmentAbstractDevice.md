# EquipmentAbstractDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Administrator defined name for the device. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Unique identity of the device. | [optional] [readonly] 
**Version** | Pointer to **string** | Current running software version of the device. | [optional] [readonly] 

## Methods

### NewEquipmentAbstractDevice

`func NewEquipmentAbstractDevice() *EquipmentAbstractDevice`

NewEquipmentAbstractDevice instantiates a new EquipmentAbstractDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentAbstractDeviceWithDefaults

`func NewEquipmentAbstractDeviceWithDefaults() *EquipmentAbstractDevice`

NewEquipmentAbstractDeviceWithDefaults instantiates a new EquipmentAbstractDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EquipmentAbstractDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentAbstractDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentAbstractDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentAbstractDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *EquipmentAbstractDevice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EquipmentAbstractDevice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EquipmentAbstractDevice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EquipmentAbstractDevice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *EquipmentAbstractDevice) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EquipmentAbstractDevice) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EquipmentAbstractDevice) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EquipmentAbstractDevice) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


