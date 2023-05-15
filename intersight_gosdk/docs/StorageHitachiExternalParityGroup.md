# StorageHitachiExternalParityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiExternalParityGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiExternalParityGroup"]
**AllocatableOpenVolumeCapacity** | Pointer to **int64** | From among the open volumes in the external parity group, the total capacity of volumes to which paths can be allocated (KB). | [optional] [readonly] 
**AllocatedOpenVolumeCapacity** | Pointer to **int64** | From among the open volumes in the external parity group, the total capacity of volumes to which paths are allocated (KB). | [optional] [readonly] 
**AvailableVolumeCapacity** | Pointer to **int64** | Available capacity of the external parity group, represented in bytes. | [optional] [readonly] 
**ClprId** | Pointer to **int64** | Number of CLPR to which the external parity group belongs. | [optional] [readonly] 
**EmulationType** | Pointer to **string** | Emulation type of the external parity group. | [optional] [readonly] 
**ExternalProductId** | Pointer to **string** | Storage system that is connected using the external storage connection functionality of Universal Volume Manager. | [optional] [readonly] 
**LargestAvailableCapacity** | Pointer to **int64** | Maximum capacity of the non-volume areas in the external parity group (KB). | [optional] [readonly] 
**Name** | Pointer to **string** | External parity group number. | [optional] [readonly] 
**ReservedOpenVolumeCapacity** | Pointer to **int64** | From among the open volumes in the external parity group, the total capacity of volumes which are reserved (KB). | [optional] [readonly] 
**Spaces** | Pointer to [**[]StorageSpace**](StorageSpace.md) |  | [optional] 
**StorageUtilization** | Pointer to [**StorageHitachiCapacity**](StorageHitachiCapacity.md) |  | [optional] 
**TotalOpenVolumeCapacity** | Pointer to **int64** | Total volume capacity of the open volumes in the external parity group (KB). | [optional] [readonly] 
**UnallocatedOpenVolumeCapacity** | Pointer to **int64** | From among the open volumes in the external parity group, the total capacity of volumes to which paths are not allocated (KB). | [optional] [readonly] 
**UsedCapacityRate** | Pointer to **int64** | Usage rate of the external parity group. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiExternalParityGroup

`func NewStorageHitachiExternalParityGroup(classId string, objectType string, ) *StorageHitachiExternalParityGroup`

NewStorageHitachiExternalParityGroup instantiates a new StorageHitachiExternalParityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiExternalParityGroupWithDefaults

`func NewStorageHitachiExternalParityGroupWithDefaults() *StorageHitachiExternalParityGroup`

NewStorageHitachiExternalParityGroupWithDefaults instantiates a new StorageHitachiExternalParityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiExternalParityGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiExternalParityGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiExternalParityGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiExternalParityGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiExternalParityGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiExternalParityGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatableOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) GetAllocatableOpenVolumeCapacity() int64`

GetAllocatableOpenVolumeCapacity returns the AllocatableOpenVolumeCapacity field if non-nil, zero value otherwise.

### GetAllocatableOpenVolumeCapacityOk

`func (o *StorageHitachiExternalParityGroup) GetAllocatableOpenVolumeCapacityOk() (*int64, bool)`

GetAllocatableOpenVolumeCapacityOk returns a tuple with the AllocatableOpenVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatableOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) SetAllocatableOpenVolumeCapacity(v int64)`

SetAllocatableOpenVolumeCapacity sets AllocatableOpenVolumeCapacity field to given value.

### HasAllocatableOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) HasAllocatableOpenVolumeCapacity() bool`

HasAllocatableOpenVolumeCapacity returns a boolean if a field has been set.

### GetAllocatedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) GetAllocatedOpenVolumeCapacity() int64`

GetAllocatedOpenVolumeCapacity returns the AllocatedOpenVolumeCapacity field if non-nil, zero value otherwise.

### GetAllocatedOpenVolumeCapacityOk

`func (o *StorageHitachiExternalParityGroup) GetAllocatedOpenVolumeCapacityOk() (*int64, bool)`

GetAllocatedOpenVolumeCapacityOk returns a tuple with the AllocatedOpenVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) SetAllocatedOpenVolumeCapacity(v int64)`

SetAllocatedOpenVolumeCapacity sets AllocatedOpenVolumeCapacity field to given value.

### HasAllocatedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) HasAllocatedOpenVolumeCapacity() bool`

HasAllocatedOpenVolumeCapacity returns a boolean if a field has been set.

### GetAvailableVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) GetAvailableVolumeCapacity() int64`

GetAvailableVolumeCapacity returns the AvailableVolumeCapacity field if non-nil, zero value otherwise.

### GetAvailableVolumeCapacityOk

`func (o *StorageHitachiExternalParityGroup) GetAvailableVolumeCapacityOk() (*int64, bool)`

GetAvailableVolumeCapacityOk returns a tuple with the AvailableVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) SetAvailableVolumeCapacity(v int64)`

SetAvailableVolumeCapacity sets AvailableVolumeCapacity field to given value.

### HasAvailableVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) HasAvailableVolumeCapacity() bool`

HasAvailableVolumeCapacity returns a boolean if a field has been set.

### GetClprId

`func (o *StorageHitachiExternalParityGroup) GetClprId() int64`

GetClprId returns the ClprId field if non-nil, zero value otherwise.

### GetClprIdOk

`func (o *StorageHitachiExternalParityGroup) GetClprIdOk() (*int64, bool)`

GetClprIdOk returns a tuple with the ClprId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClprId

`func (o *StorageHitachiExternalParityGroup) SetClprId(v int64)`

SetClprId sets ClprId field to given value.

### HasClprId

`func (o *StorageHitachiExternalParityGroup) HasClprId() bool`

HasClprId returns a boolean if a field has been set.

### GetEmulationType

`func (o *StorageHitachiExternalParityGroup) GetEmulationType() string`

GetEmulationType returns the EmulationType field if non-nil, zero value otherwise.

### GetEmulationTypeOk

`func (o *StorageHitachiExternalParityGroup) GetEmulationTypeOk() (*string, bool)`

GetEmulationTypeOk returns a tuple with the EmulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulationType

`func (o *StorageHitachiExternalParityGroup) SetEmulationType(v string)`

SetEmulationType sets EmulationType field to given value.

### HasEmulationType

`func (o *StorageHitachiExternalParityGroup) HasEmulationType() bool`

HasEmulationType returns a boolean if a field has been set.

### GetExternalProductId

`func (o *StorageHitachiExternalParityGroup) GetExternalProductId() string`

GetExternalProductId returns the ExternalProductId field if non-nil, zero value otherwise.

### GetExternalProductIdOk

`func (o *StorageHitachiExternalParityGroup) GetExternalProductIdOk() (*string, bool)`

GetExternalProductIdOk returns a tuple with the ExternalProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProductId

`func (o *StorageHitachiExternalParityGroup) SetExternalProductId(v string)`

SetExternalProductId sets ExternalProductId field to given value.

### HasExternalProductId

`func (o *StorageHitachiExternalParityGroup) HasExternalProductId() bool`

HasExternalProductId returns a boolean if a field has been set.

### GetLargestAvailableCapacity

`func (o *StorageHitachiExternalParityGroup) GetLargestAvailableCapacity() int64`

GetLargestAvailableCapacity returns the LargestAvailableCapacity field if non-nil, zero value otherwise.

### GetLargestAvailableCapacityOk

`func (o *StorageHitachiExternalParityGroup) GetLargestAvailableCapacityOk() (*int64, bool)`

GetLargestAvailableCapacityOk returns a tuple with the LargestAvailableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestAvailableCapacity

`func (o *StorageHitachiExternalParityGroup) SetLargestAvailableCapacity(v int64)`

SetLargestAvailableCapacity sets LargestAvailableCapacity field to given value.

### HasLargestAvailableCapacity

`func (o *StorageHitachiExternalParityGroup) HasLargestAvailableCapacity() bool`

HasLargestAvailableCapacity returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiExternalParityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiExternalParityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiExternalParityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiExternalParityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReservedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) GetReservedOpenVolumeCapacity() int64`

GetReservedOpenVolumeCapacity returns the ReservedOpenVolumeCapacity field if non-nil, zero value otherwise.

### GetReservedOpenVolumeCapacityOk

`func (o *StorageHitachiExternalParityGroup) GetReservedOpenVolumeCapacityOk() (*int64, bool)`

GetReservedOpenVolumeCapacityOk returns a tuple with the ReservedOpenVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) SetReservedOpenVolumeCapacity(v int64)`

SetReservedOpenVolumeCapacity sets ReservedOpenVolumeCapacity field to given value.

### HasReservedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) HasReservedOpenVolumeCapacity() bool`

HasReservedOpenVolumeCapacity returns a boolean if a field has been set.

### GetSpaces

`func (o *StorageHitachiExternalParityGroup) GetSpaces() []StorageSpace`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *StorageHitachiExternalParityGroup) GetSpacesOk() (*[]StorageSpace, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *StorageHitachiExternalParityGroup) SetSpaces(v []StorageSpace)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *StorageHitachiExternalParityGroup) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.

### SetSpacesNil

`func (o *StorageHitachiExternalParityGroup) SetSpacesNil(b bool)`

 SetSpacesNil sets the value for Spaces to be an explicit nil

### UnsetSpaces
`func (o *StorageHitachiExternalParityGroup) UnsetSpaces()`

UnsetSpaces ensures that no value is present for Spaces, not even an explicit nil
### GetStorageUtilization

`func (o *StorageHitachiExternalParityGroup) GetStorageUtilization() StorageHitachiCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageHitachiExternalParityGroup) GetStorageUtilizationOk() (*StorageHitachiCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageHitachiExternalParityGroup) SetStorageUtilization(v StorageHitachiCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageHitachiExternalParityGroup) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetTotalOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) GetTotalOpenVolumeCapacity() int64`

GetTotalOpenVolumeCapacity returns the TotalOpenVolumeCapacity field if non-nil, zero value otherwise.

### GetTotalOpenVolumeCapacityOk

`func (o *StorageHitachiExternalParityGroup) GetTotalOpenVolumeCapacityOk() (*int64, bool)`

GetTotalOpenVolumeCapacityOk returns a tuple with the TotalOpenVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) SetTotalOpenVolumeCapacity(v int64)`

SetTotalOpenVolumeCapacity sets TotalOpenVolumeCapacity field to given value.

### HasTotalOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) HasTotalOpenVolumeCapacity() bool`

HasTotalOpenVolumeCapacity returns a boolean if a field has been set.

### GetUnallocatedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) GetUnallocatedOpenVolumeCapacity() int64`

GetUnallocatedOpenVolumeCapacity returns the UnallocatedOpenVolumeCapacity field if non-nil, zero value otherwise.

### GetUnallocatedOpenVolumeCapacityOk

`func (o *StorageHitachiExternalParityGroup) GetUnallocatedOpenVolumeCapacityOk() (*int64, bool)`

GetUnallocatedOpenVolumeCapacityOk returns a tuple with the UnallocatedOpenVolumeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnallocatedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) SetUnallocatedOpenVolumeCapacity(v int64)`

SetUnallocatedOpenVolumeCapacity sets UnallocatedOpenVolumeCapacity field to given value.

### HasUnallocatedOpenVolumeCapacity

`func (o *StorageHitachiExternalParityGroup) HasUnallocatedOpenVolumeCapacity() bool`

HasUnallocatedOpenVolumeCapacity returns a boolean if a field has been set.

### GetUsedCapacityRate

`func (o *StorageHitachiExternalParityGroup) GetUsedCapacityRate() int64`

GetUsedCapacityRate returns the UsedCapacityRate field if non-nil, zero value otherwise.

### GetUsedCapacityRateOk

`func (o *StorageHitachiExternalParityGroup) GetUsedCapacityRateOk() (*int64, bool)`

GetUsedCapacityRateOk returns a tuple with the UsedCapacityRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacityRate

`func (o *StorageHitachiExternalParityGroup) SetUsedCapacityRate(v int64)`

SetUsedCapacityRate sets UsedCapacityRate field to given value.

### HasUsedCapacityRate

`func (o *StorageHitachiExternalParityGroup) HasUsedCapacityRate() bool`

HasUsedCapacityRate returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiExternalParityGroup) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiExternalParityGroup) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiExternalParityGroup) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiExternalParityGroup) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiExternalParityGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiExternalParityGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiExternalParityGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiExternalParityGroup) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


