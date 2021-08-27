# StorageSasExpanderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.SasExpander"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.SasExpander"]
**ExpanderId** | Pointer to **int64** | Unique Identifier of the storage expander. | [optional] [readonly] 
**Name** | Pointer to **string** | The name  of the installed storage expander. | [optional] 
**OperState** | Pointer to **string** | The operational state of the storage expander. | [optional] [readonly] 
**Operability** | Pointer to **string** | The operability status of the storage expander. | [optional] [readonly] 
**SasAddress** | Pointer to **string** | The SAS address of the SAS expander. | [optional] [readonly] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) |  | [optional] 
**Controller** | Pointer to [**ManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageSasExpanderAllOf

`func NewStorageSasExpanderAllOf(classId string, objectType string, ) *StorageSasExpanderAllOf`

NewStorageSasExpanderAllOf instantiates a new StorageSasExpanderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSasExpanderAllOfWithDefaults

`func NewStorageSasExpanderAllOfWithDefaults() *StorageSasExpanderAllOf`

NewStorageSasExpanderAllOfWithDefaults instantiates a new StorageSasExpanderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageSasExpanderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageSasExpanderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageSasExpanderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageSasExpanderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageSasExpanderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageSasExpanderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpanderId

`func (o *StorageSasExpanderAllOf) GetExpanderId() int64`

GetExpanderId returns the ExpanderId field if non-nil, zero value otherwise.

### GetExpanderIdOk

`func (o *StorageSasExpanderAllOf) GetExpanderIdOk() (*int64, bool)`

GetExpanderIdOk returns a tuple with the ExpanderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanderId

`func (o *StorageSasExpanderAllOf) SetExpanderId(v int64)`

SetExpanderId sets ExpanderId field to given value.

### HasExpanderId

`func (o *StorageSasExpanderAllOf) HasExpanderId() bool`

HasExpanderId returns a boolean if a field has been set.

### GetName

`func (o *StorageSasExpanderAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSasExpanderAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSasExpanderAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageSasExpanderAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *StorageSasExpanderAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *StorageSasExpanderAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *StorageSasExpanderAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *StorageSasExpanderAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *StorageSasExpanderAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *StorageSasExpanderAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *StorageSasExpanderAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *StorageSasExpanderAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetSasAddress

`func (o *StorageSasExpanderAllOf) GetSasAddress() string`

GetSasAddress returns the SasAddress field if non-nil, zero value otherwise.

### GetSasAddressOk

`func (o *StorageSasExpanderAllOf) GetSasAddressOk() (*string, bool)`

GetSasAddressOk returns a tuple with the SasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasAddress

`func (o *StorageSasExpanderAllOf) SetSasAddress(v string)`

SetSasAddress sets SasAddress field to given value.

### HasSasAddress

`func (o *StorageSasExpanderAllOf) HasSasAddress() bool`

HasSasAddress returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *StorageSasExpanderAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *StorageSasExpanderAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *StorageSasExpanderAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *StorageSasExpanderAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetController

`func (o *StorageSasExpanderAllOf) GetController() ManagementControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StorageSasExpanderAllOf) GetControllerOk() (*ManagementControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StorageSasExpanderAllOf) SetController(v ManagementControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *StorageSasExpanderAllOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *StorageSasExpanderAllOf) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *StorageSasExpanderAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *StorageSasExpanderAllOf) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *StorageSasExpanderAllOf) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageSasExpanderAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageSasExpanderAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageSasExpanderAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageSasExpanderAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageSasExpanderAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageSasExpanderAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageSasExpanderAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageSasExpanderAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


