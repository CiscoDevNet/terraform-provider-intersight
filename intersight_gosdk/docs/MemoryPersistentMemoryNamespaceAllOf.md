# MemoryPersistentMemoryNamespaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** | Capacity in GiB of the Persistent Memory Namespace. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Namespace. | [optional] [readonly] 
**LabelVersion** | Pointer to **string** | Label version of the Persistent Memory Namespace. | [optional] [readonly] 
**Mode** | Pointer to **string** | Mode of the Persistent Memory Namespace. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Persistent Memory Namespace. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the Persistent Memory Namespace. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryPersistentMemoryRegion** | Pointer to [**MemoryPersistentMemoryRegionRelationship**](memory.PersistentMemoryRegion.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryNamespaceAllOf

`func NewMemoryPersistentMemoryNamespaceAllOf() *MemoryPersistentMemoryNamespaceAllOf`

NewMemoryPersistentMemoryNamespaceAllOf instantiates a new MemoryPersistentMemoryNamespaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryNamespaceAllOfWithDefaults

`func NewMemoryPersistentMemoryNamespaceAllOfWithDefaults() *MemoryPersistentMemoryNamespaceAllOf`

NewMemoryPersistentMemoryNamespaceAllOfWithDefaults instantiates a new MemoryPersistentMemoryNamespaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetLabelVersion

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetLabelVersion() string`

GetLabelVersion returns the LabelVersion field if non-nil, zero value otherwise.

### GetLabelVersionOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetLabelVersionOk() (*string, bool)`

GetLabelVersionOk returns a tuple with the LabelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelVersion

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetLabelVersion(v string)`

SetLabelVersion sets LabelVersion field to given value.

### HasLabelVersion

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasLabelVersion() bool`

HasLabelVersion returns a boolean if a field has been set.

### GetMode

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryRegion

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetMemoryPersistentMemoryRegion() MemoryPersistentMemoryRegionRelationship`

GetMemoryPersistentMemoryRegion returns the MemoryPersistentMemoryRegion field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryRegionOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetMemoryPersistentMemoryRegionOk() (*MemoryPersistentMemoryRegionRelationship, bool)`

GetMemoryPersistentMemoryRegionOk returns a tuple with the MemoryPersistentMemoryRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryRegion

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetMemoryPersistentMemoryRegion(v MemoryPersistentMemoryRegionRelationship)`

SetMemoryPersistentMemoryRegion sets MemoryPersistentMemoryRegion field to given value.

### HasMemoryPersistentMemoryRegion

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasMemoryPersistentMemoryRegion() bool`

HasMemoryPersistentMemoryRegion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryNamespaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryNamespaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryNamespaceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


