# MemoryPersistentMemoryNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.PersistentMemoryNamespace"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.PersistentMemoryNamespace"]
**Capacity** | Pointer to **string** | Capacity in GiB of the Persistent Memory Namespace. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Namespace. | [optional] [readonly] 
**LabelVersion** | Pointer to **string** | Label version of the Persistent Memory Namespace. | [optional] [readonly] 
**Mode** | Pointer to **string** | Mode of the Persistent Memory Namespace. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Persistent Memory Namespace. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the Persistent Memory Namespace. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**MemoryPersistentMemoryRegion** | Pointer to [**MemoryPersistentMemoryRegionRelationship**](MemoryPersistentMemoryRegionRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryNamespace

`func NewMemoryPersistentMemoryNamespace(classId string, objectType string, ) *MemoryPersistentMemoryNamespace`

NewMemoryPersistentMemoryNamespace instantiates a new MemoryPersistentMemoryNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryNamespaceWithDefaults

`func NewMemoryPersistentMemoryNamespaceWithDefaults() *MemoryPersistentMemoryNamespace`

NewMemoryPersistentMemoryNamespaceWithDefaults instantiates a new MemoryPersistentMemoryNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPersistentMemoryNamespace) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryNamespace) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryNamespace) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPersistentMemoryNamespace) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryNamespace) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryNamespace) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *MemoryPersistentMemoryNamespace) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MemoryPersistentMemoryNamespace) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MemoryPersistentMemoryNamespace) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MemoryPersistentMemoryNamespace) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryNamespace) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryNamespace) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryNamespace) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryNamespace) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetLabelVersion

`func (o *MemoryPersistentMemoryNamespace) GetLabelVersion() string`

GetLabelVersion returns the LabelVersion field if non-nil, zero value otherwise.

### GetLabelVersionOk

`func (o *MemoryPersistentMemoryNamespace) GetLabelVersionOk() (*string, bool)`

GetLabelVersionOk returns a tuple with the LabelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelVersion

`func (o *MemoryPersistentMemoryNamespace) SetLabelVersion(v string)`

SetLabelVersion sets LabelVersion field to given value.

### HasLabelVersion

`func (o *MemoryPersistentMemoryNamespace) HasLabelVersion() bool`

HasLabelVersion returns a boolean if a field has been set.

### GetMode

`func (o *MemoryPersistentMemoryNamespace) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MemoryPersistentMemoryNamespace) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MemoryPersistentMemoryNamespace) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MemoryPersistentMemoryNamespace) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *MemoryPersistentMemoryNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemoryPersistentMemoryNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemoryPersistentMemoryNamespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemoryPersistentMemoryNamespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *MemoryPersistentMemoryNamespace) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MemoryPersistentMemoryNamespace) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MemoryPersistentMemoryNamespace) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MemoryPersistentMemoryNamespace) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespace) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryNamespace) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespace) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespace) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryRegion

`func (o *MemoryPersistentMemoryNamespace) GetMemoryPersistentMemoryRegion() MemoryPersistentMemoryRegionRelationship`

GetMemoryPersistentMemoryRegion returns the MemoryPersistentMemoryRegion field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryRegionOk

`func (o *MemoryPersistentMemoryNamespace) GetMemoryPersistentMemoryRegionOk() (*MemoryPersistentMemoryRegionRelationship, bool)`

GetMemoryPersistentMemoryRegionOk returns a tuple with the MemoryPersistentMemoryRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryRegion

`func (o *MemoryPersistentMemoryNamespace) SetMemoryPersistentMemoryRegion(v MemoryPersistentMemoryRegionRelationship)`

SetMemoryPersistentMemoryRegion sets MemoryPersistentMemoryRegion field to given value.

### HasMemoryPersistentMemoryRegion

`func (o *MemoryPersistentMemoryNamespace) HasMemoryPersistentMemoryRegion() bool`

HasMemoryPersistentMemoryRegion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *MemoryPersistentMemoryNamespace) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryNamespace) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryNamespace) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryNamespace) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


