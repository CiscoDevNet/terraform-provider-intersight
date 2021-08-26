# PciCoprocessorCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pci.CoprocessorCard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pci.CoprocessorCard"]
**CardId** | Pointer to **int64** | The id of the coprocessor card. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The PCI slot name for the coprocessor card. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPciCoprocessorCard

`func NewPciCoprocessorCard(classId string, objectType string, ) *PciCoprocessorCard`

NewPciCoprocessorCard instantiates a new PciCoprocessorCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciCoprocessorCardWithDefaults

`func NewPciCoprocessorCardWithDefaults() *PciCoprocessorCard`

NewPciCoprocessorCardWithDefaults instantiates a new PciCoprocessorCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PciCoprocessorCard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PciCoprocessorCard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PciCoprocessorCard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PciCoprocessorCard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PciCoprocessorCard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PciCoprocessorCard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCardId

`func (o *PciCoprocessorCard) GetCardId() int64`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *PciCoprocessorCard) GetCardIdOk() (*int64, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *PciCoprocessorCard) SetCardId(v int64)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *PciCoprocessorCard) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetPciSlot

`func (o *PciCoprocessorCard) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *PciCoprocessorCard) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *PciCoprocessorCard) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *PciCoprocessorCard) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetComputeBoard

`func (o *PciCoprocessorCard) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *PciCoprocessorCard) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *PciCoprocessorCard) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *PciCoprocessorCard) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *PciCoprocessorCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciCoprocessorCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciCoprocessorCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciCoprocessorCard) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PciCoprocessorCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciCoprocessorCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciCoprocessorCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciCoprocessorCard) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


