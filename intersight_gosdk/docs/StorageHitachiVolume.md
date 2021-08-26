# StorageHitachiVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiVolume"]
**Attributes** | Pointer to **[]string** |  | [optional] 
**ClprId** | Pointer to **int64** | CLPR (Cache Logical Partition) number of this volume. | [optional] [readonly] 
**DataReductionMode** | Pointer to **string** | Setting of the capacity saving function (dedupe and compression). * &#x60;N/A&#x60; - The capacity saving function is not available. * &#x60;Compression&#x60; - The capacity saving function (compression) is enabled. * &#x60;Compression Deduplication&#x60; - The capacity saving function (compression and deduplication) is enabled. * &#x60;Disabled&#x60; - The capacity saving function (compression and deduplication) is disabled. | [optional] [readonly] [default to "N/A"]
**DataReductionStatus** | Pointer to **string** | Status of the capacity saving function. * &#x60;N/A&#x60; - The capacity saving function is not available. * &#x60;Enabled&#x60; - The capacity saving function is enabled. * &#x60;Disabled&#x60; - The capacity saving function is disabled. * &#x60;Enabling&#x60; - The capacity saving function is being enabled. * &#x60;Rehydrating&#x60; - The capacity saving function is being disabled. * &#x60;Deleting&#x60; - The volumes for which the capacity saving function is enabled are being deleted. * &#x60;Failed&#x60; - An attempt to enable the capacity saving function failed. | [optional] [readonly] [default to "N/A"]
**DriveType** | Pointer to **string** | Code indicating the drive type of the drive belonging to the volume. | [optional] [readonly] 
**EmulationType** | Pointer to **string** | The volume emulation type or the volume status information. * &#x60;N/A&#x60; - Not available. * &#x60;NOT DEFINED&#x60; - The volume is not implemented. * &#x60;DEFINING&#x60; - The volume is being created. * &#x60;REMOVING&#x60; - The volume is being removed. * &#x60;OPEN-V&#x60; - To be provided by Hitachi. | [optional] [readonly] [default to "N/A"]
**IsFullAllocationEnabled** | Pointer to **bool** | Whether pages are reserved by the FullAllocation functionality. | [optional] [readonly] 
**Label** | Pointer to **string** | Label of the volume, as configured in the storage array. | [optional] [readonly] 
**NumOfPaths** | Pointer to **int64** | Number of paths set for the volume. | [optional] [readonly] 
**ParityGroupIds** | Pointer to **[]string** |  | [optional] 
**PoolId** | Pointer to **string** | ID of the pool with which the volume is associated. | [optional] [readonly] 
**RaidLevel** | Pointer to **string** | RAID level for the volume. * &#x60;N/A&#x60; - RAID level is unknown or multiple RAID levels are being used. * &#x60;RAID1&#x60; - RAID1. * &#x60;RAID5&#x60; - RAID5. * &#x60;RAID6&#x60; - RAID6. | [optional] [readonly] [default to "N/A"]
**RaidType** | Pointer to **string** | RAID type drive configuration. | [optional] [readonly] 
**Status** | Pointer to **string** | Status information of the volume. * &#x60;N/A&#x60; - The volume status is not available. * &#x60;NML&#x60; - The volume is in normal status. * &#x60;BLK&#x60; - The volume is in blocked state. * &#x60;BSY&#x60; - The volume status is being changed. * &#x60;Unknown&#x60; - The volume status is unknown (not supported). | [optional] [readonly] [default to "N/A"]
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**ParityGroups** | Pointer to [**[]StorageHitachiParityGroupRelationship**](StorageHitachiParityGroupRelationship.md) | An array of relationships to storageHitachiParityGroup resources. | [optional] [readonly] 
**Pool** | Pointer to [**StorageHitachiPoolRelationship**](StorageHitachiPoolRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiVolume

`func NewStorageHitachiVolume(classId string, objectType string, ) *StorageHitachiVolume`

NewStorageHitachiVolume instantiates a new StorageHitachiVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiVolumeWithDefaults

`func NewStorageHitachiVolumeWithDefaults() *StorageHitachiVolume`

NewStorageHitachiVolumeWithDefaults instantiates a new StorageHitachiVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiVolume) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiVolume) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiVolume) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiVolume) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiVolume) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiVolume) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributes

`func (o *StorageHitachiVolume) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StorageHitachiVolume) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StorageHitachiVolume) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *StorageHitachiVolume) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *StorageHitachiVolume) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *StorageHitachiVolume) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetClprId

`func (o *StorageHitachiVolume) GetClprId() int64`

GetClprId returns the ClprId field if non-nil, zero value otherwise.

### GetClprIdOk

`func (o *StorageHitachiVolume) GetClprIdOk() (*int64, bool)`

GetClprIdOk returns a tuple with the ClprId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClprId

`func (o *StorageHitachiVolume) SetClprId(v int64)`

SetClprId sets ClprId field to given value.

### HasClprId

`func (o *StorageHitachiVolume) HasClprId() bool`

HasClprId returns a boolean if a field has been set.

### GetDataReductionMode

`func (o *StorageHitachiVolume) GetDataReductionMode() string`

GetDataReductionMode returns the DataReductionMode field if non-nil, zero value otherwise.

### GetDataReductionModeOk

`func (o *StorageHitachiVolume) GetDataReductionModeOk() (*string, bool)`

GetDataReductionModeOk returns a tuple with the DataReductionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReductionMode

`func (o *StorageHitachiVolume) SetDataReductionMode(v string)`

SetDataReductionMode sets DataReductionMode field to given value.

### HasDataReductionMode

`func (o *StorageHitachiVolume) HasDataReductionMode() bool`

HasDataReductionMode returns a boolean if a field has been set.

### GetDataReductionStatus

`func (o *StorageHitachiVolume) GetDataReductionStatus() string`

GetDataReductionStatus returns the DataReductionStatus field if non-nil, zero value otherwise.

### GetDataReductionStatusOk

`func (o *StorageHitachiVolume) GetDataReductionStatusOk() (*string, bool)`

GetDataReductionStatusOk returns a tuple with the DataReductionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReductionStatus

`func (o *StorageHitachiVolume) SetDataReductionStatus(v string)`

SetDataReductionStatus sets DataReductionStatus field to given value.

### HasDataReductionStatus

`func (o *StorageHitachiVolume) HasDataReductionStatus() bool`

HasDataReductionStatus returns a boolean if a field has been set.

### GetDriveType

`func (o *StorageHitachiVolume) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *StorageHitachiVolume) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *StorageHitachiVolume) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *StorageHitachiVolume) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetEmulationType

`func (o *StorageHitachiVolume) GetEmulationType() string`

GetEmulationType returns the EmulationType field if non-nil, zero value otherwise.

### GetEmulationTypeOk

`func (o *StorageHitachiVolume) GetEmulationTypeOk() (*string, bool)`

GetEmulationTypeOk returns a tuple with the EmulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulationType

`func (o *StorageHitachiVolume) SetEmulationType(v string)`

SetEmulationType sets EmulationType field to given value.

### HasEmulationType

`func (o *StorageHitachiVolume) HasEmulationType() bool`

HasEmulationType returns a boolean if a field has been set.

### GetIsFullAllocationEnabled

`func (o *StorageHitachiVolume) GetIsFullAllocationEnabled() bool`

GetIsFullAllocationEnabled returns the IsFullAllocationEnabled field if non-nil, zero value otherwise.

### GetIsFullAllocationEnabledOk

`func (o *StorageHitachiVolume) GetIsFullAllocationEnabledOk() (*bool, bool)`

GetIsFullAllocationEnabledOk returns a tuple with the IsFullAllocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullAllocationEnabled

`func (o *StorageHitachiVolume) SetIsFullAllocationEnabled(v bool)`

SetIsFullAllocationEnabled sets IsFullAllocationEnabled field to given value.

### HasIsFullAllocationEnabled

`func (o *StorageHitachiVolume) HasIsFullAllocationEnabled() bool`

HasIsFullAllocationEnabled returns a boolean if a field has been set.

### GetLabel

`func (o *StorageHitachiVolume) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StorageHitachiVolume) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StorageHitachiVolume) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StorageHitachiVolume) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNumOfPaths

`func (o *StorageHitachiVolume) GetNumOfPaths() int64`

GetNumOfPaths returns the NumOfPaths field if non-nil, zero value otherwise.

### GetNumOfPathsOk

`func (o *StorageHitachiVolume) GetNumOfPathsOk() (*int64, bool)`

GetNumOfPathsOk returns a tuple with the NumOfPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfPaths

`func (o *StorageHitachiVolume) SetNumOfPaths(v int64)`

SetNumOfPaths sets NumOfPaths field to given value.

### HasNumOfPaths

`func (o *StorageHitachiVolume) HasNumOfPaths() bool`

HasNumOfPaths returns a boolean if a field has been set.

### GetParityGroupIds

`func (o *StorageHitachiVolume) GetParityGroupIds() []string`

GetParityGroupIds returns the ParityGroupIds field if non-nil, zero value otherwise.

### GetParityGroupIdsOk

`func (o *StorageHitachiVolume) GetParityGroupIdsOk() (*[]string, bool)`

GetParityGroupIdsOk returns a tuple with the ParityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParityGroupIds

`func (o *StorageHitachiVolume) SetParityGroupIds(v []string)`

SetParityGroupIds sets ParityGroupIds field to given value.

### HasParityGroupIds

`func (o *StorageHitachiVolume) HasParityGroupIds() bool`

HasParityGroupIds returns a boolean if a field has been set.

### SetParityGroupIdsNil

`func (o *StorageHitachiVolume) SetParityGroupIdsNil(b bool)`

 SetParityGroupIdsNil sets the value for ParityGroupIds to be an explicit nil

### UnsetParityGroupIds
`func (o *StorageHitachiVolume) UnsetParityGroupIds()`

UnsetParityGroupIds ensures that no value is present for ParityGroupIds, not even an explicit nil
### GetPoolId

`func (o *StorageHitachiVolume) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *StorageHitachiVolume) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *StorageHitachiVolume) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *StorageHitachiVolume) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetRaidLevel

`func (o *StorageHitachiVolume) GetRaidLevel() string`

GetRaidLevel returns the RaidLevel field if non-nil, zero value otherwise.

### GetRaidLevelOk

`func (o *StorageHitachiVolume) GetRaidLevelOk() (*string, bool)`

GetRaidLevelOk returns a tuple with the RaidLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidLevel

`func (o *StorageHitachiVolume) SetRaidLevel(v string)`

SetRaidLevel sets RaidLevel field to given value.

### HasRaidLevel

`func (o *StorageHitachiVolume) HasRaidLevel() bool`

HasRaidLevel returns a boolean if a field has been set.

### GetRaidType

`func (o *StorageHitachiVolume) GetRaidType() string`

GetRaidType returns the RaidType field if non-nil, zero value otherwise.

### GetRaidTypeOk

`func (o *StorageHitachiVolume) GetRaidTypeOk() (*string, bool)`

GetRaidTypeOk returns a tuple with the RaidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidType

`func (o *StorageHitachiVolume) SetRaidType(v string)`

SetRaidType sets RaidType field to given value.

### HasRaidType

`func (o *StorageHitachiVolume) HasRaidType() bool`

HasRaidType returns a boolean if a field has been set.

### GetStatus

`func (o *StorageHitachiVolume) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageHitachiVolume) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageHitachiVolume) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageHitachiVolume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiVolume) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiVolume) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiVolume) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiVolume) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetParityGroups

`func (o *StorageHitachiVolume) GetParityGroups() []StorageHitachiParityGroupRelationship`

GetParityGroups returns the ParityGroups field if non-nil, zero value otherwise.

### GetParityGroupsOk

`func (o *StorageHitachiVolume) GetParityGroupsOk() (*[]StorageHitachiParityGroupRelationship, bool)`

GetParityGroupsOk returns a tuple with the ParityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParityGroups

`func (o *StorageHitachiVolume) SetParityGroups(v []StorageHitachiParityGroupRelationship)`

SetParityGroups sets ParityGroups field to given value.

### HasParityGroups

`func (o *StorageHitachiVolume) HasParityGroups() bool`

HasParityGroups returns a boolean if a field has been set.

### SetParityGroupsNil

`func (o *StorageHitachiVolume) SetParityGroupsNil(b bool)`

 SetParityGroupsNil sets the value for ParityGroups to be an explicit nil

### UnsetParityGroups
`func (o *StorageHitachiVolume) UnsetParityGroups()`

UnsetParityGroups ensures that no value is present for ParityGroups, not even an explicit nil
### GetPool

`func (o *StorageHitachiVolume) GetPool() StorageHitachiPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *StorageHitachiVolume) GetPoolOk() (*StorageHitachiPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *StorageHitachiVolume) SetPool(v StorageHitachiPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *StorageHitachiVolume) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiVolume) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiVolume) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiVolume) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiVolume) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


