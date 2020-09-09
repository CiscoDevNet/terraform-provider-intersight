# StorageSasExpander

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpanderId** | Pointer to **int64** | Unique Identifier of the storage expander. | [optional] [readonly] 
**Name** | Pointer to **string** | The name  of the installed storage expander. | [optional] 
**OperState** | Pointer to **string** | The operational state of the storage expander. | [optional] [readonly] 
**Operability** | Pointer to **string** | The operability status of the storage expander. | [optional] [readonly] 
**Presence** | Pointer to **string** | The availability of the storage expander. | [optional] [readonly] 
**SasAddress** | Pointer to **string** | The SAS address of the SAS expander. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**Controller** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStorageSasExpander

`func NewStorageSasExpander() *StorageSasExpander`

NewStorageSasExpander instantiates a new StorageSasExpander object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSasExpanderWithDefaults

`func NewStorageSasExpanderWithDefaults() *StorageSasExpander`

NewStorageSasExpanderWithDefaults instantiates a new StorageSasExpander object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpanderId

`func (o *StorageSasExpander) GetExpanderId() int64`

GetExpanderId returns the ExpanderId field if non-nil, zero value otherwise.

### GetExpanderIdOk

`func (o *StorageSasExpander) GetExpanderIdOk() (*int64, bool)`

GetExpanderIdOk returns a tuple with the ExpanderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderId

`func (o *StorageSasExpander) SetExpanderId(v int64)`

SetExpanderId sets ExpanderId field to given value.

### HasExpanderId

`func (o *StorageSasExpander) HasExpanderId() bool`

HasExpanderId returns a boolean if a field has been set.

### GetName

`func (o *StorageSasExpander) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSasExpander) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSasExpander) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageSasExpander) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *StorageSasExpander) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageSasExpander) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageSasExpander) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageSasExpander) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageSasExpander) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageSasExpander) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageSasExpander) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageSasExpander) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPresence

`func (o *StorageSasExpander) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageSasExpander) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageSasExpander) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageSasExpander) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetSasAddress

`func (o *StorageSasExpander) GetSasAddress() string`

GetSasAddress returns the SasAddress field if non-nil, zero value otherwise.

### GetSasAddressOk

`func (o *StorageSasExpander) GetSasAddressOk() (*string, bool)`

GetSasAddressOk returns a tuple with the SasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasAddress

`func (o *StorageSasExpander) SetSasAddress(v string)`

SetSasAddress sets SasAddress field to given value.

### HasSasAddress

`func (o *StorageSasExpander) HasSasAddress() bool`

HasSasAddress returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *StorageSasExpander) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *StorageSasExpander) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *StorageSasExpander) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *StorageSasExpander) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetController

`func (o *StorageSasExpander) GetController() ManagementControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StorageSasExpander) GetControllerOk() (*ManagementControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StorageSasExpander) SetController(v ManagementControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *StorageSasExpander) HasController() bool`

HasController returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *StorageSasExpander) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *StorageSasExpander) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *StorageSasExpander) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *StorageSasExpander) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageSasExpander) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageSasExpander) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageSasExpander) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageSasExpander) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageSasExpander) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageSasExpander) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageSasExpander) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageSasExpander) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


