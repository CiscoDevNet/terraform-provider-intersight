# EquipmentInterconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Interconnect"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Interconnect"]
**InterconnectId** | Pointer to **string** | The identifier of the interconnect. | [optional] [readonly] 
**InterconnectType** | Pointer to **string** | Type of interconnectivity provided by this interconnect. * &#x60;Unknown&#x60; - Interconnect type is unknown. * &#x60;NVLink&#x60; - Interconnect type is NVLink. | [optional] [readonly] [default to "Unknown"]
**ComputeBoard** | Pointer to [**NullableComputeBoardRelationship**](ComputeBoardRelationship.md) |  | [optional] 
**InterconnectedGraphicsCards** | Pointer to [**[]GraphicsCardRelationship**](GraphicsCardRelationship.md) | An array of relationships to graphicsCard resources. | [optional] [readonly] 
**InterconnectedSharedGraphicsCards** | Pointer to [**[]EquipmentSharedGraphicsCardRelationship**](EquipmentSharedGraphicsCardRelationship.md) | An array of relationships to equipmentSharedGraphicsCard resources. | [optional] [readonly] 
**PciNode** | Pointer to [**NullablePciNodeRelationship**](PciNodeRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentInterconnect

`func NewEquipmentInterconnect(classId string, objectType string, ) *EquipmentInterconnect`

NewEquipmentInterconnect instantiates a new EquipmentInterconnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentInterconnectWithDefaults

`func NewEquipmentInterconnectWithDefaults() *EquipmentInterconnect`

NewEquipmentInterconnectWithDefaults instantiates a new EquipmentInterconnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentInterconnect) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentInterconnect) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentInterconnect) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentInterconnect) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentInterconnect) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentInterconnect) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterconnectId

`func (o *EquipmentInterconnect) GetInterconnectId() string`

GetInterconnectId returns the InterconnectId field if non-nil, zero value otherwise.

### GetInterconnectIdOk

`func (o *EquipmentInterconnect) GetInterconnectIdOk() (*string, bool)`

GetInterconnectIdOk returns a tuple with the InterconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectId

`func (o *EquipmentInterconnect) SetInterconnectId(v string)`

SetInterconnectId sets InterconnectId field to given value.

### HasInterconnectId

`func (o *EquipmentInterconnect) HasInterconnectId() bool`

HasInterconnectId returns a boolean if a field has been set.

### GetInterconnectType

`func (o *EquipmentInterconnect) GetInterconnectType() string`

GetInterconnectType returns the InterconnectType field if non-nil, zero value otherwise.

### GetInterconnectTypeOk

`func (o *EquipmentInterconnect) GetInterconnectTypeOk() (*string, bool)`

GetInterconnectTypeOk returns a tuple with the InterconnectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectType

`func (o *EquipmentInterconnect) SetInterconnectType(v string)`

SetInterconnectType sets InterconnectType field to given value.

### HasInterconnectType

`func (o *EquipmentInterconnect) HasInterconnectType() bool`

HasInterconnectType returns a boolean if a field has been set.

### GetComputeBoard

`func (o *EquipmentInterconnect) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *EquipmentInterconnect) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *EquipmentInterconnect) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *EquipmentInterconnect) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### SetComputeBoardNil

`func (o *EquipmentInterconnect) SetComputeBoardNil(b bool)`

 SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil

### UnsetComputeBoard
`func (o *EquipmentInterconnect) UnsetComputeBoard()`

UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
### GetInterconnectedGraphicsCards

`func (o *EquipmentInterconnect) GetInterconnectedGraphicsCards() []GraphicsCardRelationship`

GetInterconnectedGraphicsCards returns the InterconnectedGraphicsCards field if non-nil, zero value otherwise.

### GetInterconnectedGraphicsCardsOk

`func (o *EquipmentInterconnect) GetInterconnectedGraphicsCardsOk() (*[]GraphicsCardRelationship, bool)`

GetInterconnectedGraphicsCardsOk returns a tuple with the InterconnectedGraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectedGraphicsCards

`func (o *EquipmentInterconnect) SetInterconnectedGraphicsCards(v []GraphicsCardRelationship)`

SetInterconnectedGraphicsCards sets InterconnectedGraphicsCards field to given value.

### HasInterconnectedGraphicsCards

`func (o *EquipmentInterconnect) HasInterconnectedGraphicsCards() bool`

HasInterconnectedGraphicsCards returns a boolean if a field has been set.

### SetInterconnectedGraphicsCardsNil

`func (o *EquipmentInterconnect) SetInterconnectedGraphicsCardsNil(b bool)`

 SetInterconnectedGraphicsCardsNil sets the value for InterconnectedGraphicsCards to be an explicit nil

### UnsetInterconnectedGraphicsCards
`func (o *EquipmentInterconnect) UnsetInterconnectedGraphicsCards()`

UnsetInterconnectedGraphicsCards ensures that no value is present for InterconnectedGraphicsCards, not even an explicit nil
### GetInterconnectedSharedGraphicsCards

`func (o *EquipmentInterconnect) GetInterconnectedSharedGraphicsCards() []EquipmentSharedGraphicsCardRelationship`

GetInterconnectedSharedGraphicsCards returns the InterconnectedSharedGraphicsCards field if non-nil, zero value otherwise.

### GetInterconnectedSharedGraphicsCardsOk

`func (o *EquipmentInterconnect) GetInterconnectedSharedGraphicsCardsOk() (*[]EquipmentSharedGraphicsCardRelationship, bool)`

GetInterconnectedSharedGraphicsCardsOk returns a tuple with the InterconnectedSharedGraphicsCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectedSharedGraphicsCards

`func (o *EquipmentInterconnect) SetInterconnectedSharedGraphicsCards(v []EquipmentSharedGraphicsCardRelationship)`

SetInterconnectedSharedGraphicsCards sets InterconnectedSharedGraphicsCards field to given value.

### HasInterconnectedSharedGraphicsCards

`func (o *EquipmentInterconnect) HasInterconnectedSharedGraphicsCards() bool`

HasInterconnectedSharedGraphicsCards returns a boolean if a field has been set.

### SetInterconnectedSharedGraphicsCardsNil

`func (o *EquipmentInterconnect) SetInterconnectedSharedGraphicsCardsNil(b bool)`

 SetInterconnectedSharedGraphicsCardsNil sets the value for InterconnectedSharedGraphicsCards to be an explicit nil

### UnsetInterconnectedSharedGraphicsCards
`func (o *EquipmentInterconnect) UnsetInterconnectedSharedGraphicsCards()`

UnsetInterconnectedSharedGraphicsCards ensures that no value is present for InterconnectedSharedGraphicsCards, not even an explicit nil
### GetPciNode

`func (o *EquipmentInterconnect) GetPciNode() PciNodeRelationship`

GetPciNode returns the PciNode field if non-nil, zero value otherwise.

### GetPciNodeOk

`func (o *EquipmentInterconnect) GetPciNodeOk() (*PciNodeRelationship, bool)`

GetPciNodeOk returns a tuple with the PciNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciNode

`func (o *EquipmentInterconnect) SetPciNode(v PciNodeRelationship)`

SetPciNode sets PciNode field to given value.

### HasPciNode

`func (o *EquipmentInterconnect) HasPciNode() bool`

HasPciNode returns a boolean if a field has been set.

### SetPciNodeNil

`func (o *EquipmentInterconnect) SetPciNodeNil(b bool)`

 SetPciNodeNil sets the value for PciNode to be an explicit nil

### UnsetPciNode
`func (o *EquipmentInterconnect) UnsetPciNode()`

UnsetPciNode ensures that no value is present for PciNode, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentInterconnect) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentInterconnect) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentInterconnect) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentInterconnect) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EquipmentInterconnect) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EquipmentInterconnect) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


