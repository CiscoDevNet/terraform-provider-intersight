# StorageNvmeRaidDriveGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NvmeRaidDriveGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NvmeRaidDriveGroup"]
**Configurations** | Pointer to [**[]StorageNvmeVirtualDriveConfiguration**](StorageNvmeVirtualDriveConfiguration.md) |  | [optional] 
**DedicatedHotSparesForDriveGroup** | Pointer to [**[]StorageNvmeDedicatedHotSpareConfiguration**](StorageNvmeDedicatedHotSpareConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | The DriveGroup Name which is used to create virtual Drives at endpoint. | [optional] [readonly] 

## Methods

### NewStorageNvmeRaidDriveGroup

`func NewStorageNvmeRaidDriveGroup(classId string, objectType string, ) *StorageNvmeRaidDriveGroup`

NewStorageNvmeRaidDriveGroup instantiates a new StorageNvmeRaidDriveGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNvmeRaidDriveGroupWithDefaults

`func NewStorageNvmeRaidDriveGroupWithDefaults() *StorageNvmeRaidDriveGroup`

NewStorageNvmeRaidDriveGroupWithDefaults instantiates a new StorageNvmeRaidDriveGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNvmeRaidDriveGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNvmeRaidDriveGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNvmeRaidDriveGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNvmeRaidDriveGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNvmeRaidDriveGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNvmeRaidDriveGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigurations

`func (o *StorageNvmeRaidDriveGroup) GetConfigurations() []StorageNvmeVirtualDriveConfiguration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *StorageNvmeRaidDriveGroup) GetConfigurationsOk() (*[]StorageNvmeVirtualDriveConfiguration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *StorageNvmeRaidDriveGroup) SetConfigurations(v []StorageNvmeVirtualDriveConfiguration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *StorageNvmeRaidDriveGroup) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### SetConfigurationsNil

`func (o *StorageNvmeRaidDriveGroup) SetConfigurationsNil(b bool)`

 SetConfigurationsNil sets the value for Configurations to be an explicit nil

### UnsetConfigurations
`func (o *StorageNvmeRaidDriveGroup) UnsetConfigurations()`

UnsetConfigurations ensures that no value is present for Configurations, not even an explicit nil
### GetDedicatedHotSparesForDriveGroup

`func (o *StorageNvmeRaidDriveGroup) GetDedicatedHotSparesForDriveGroup() []StorageNvmeDedicatedHotSpareConfiguration`

GetDedicatedHotSparesForDriveGroup returns the DedicatedHotSparesForDriveGroup field if non-nil, zero value otherwise.

### GetDedicatedHotSparesForDriveGroupOk

`func (o *StorageNvmeRaidDriveGroup) GetDedicatedHotSparesForDriveGroupOk() (*[]StorageNvmeDedicatedHotSpareConfiguration, bool)`

GetDedicatedHotSparesForDriveGroupOk returns a tuple with the DedicatedHotSparesForDriveGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedHotSparesForDriveGroup

`func (o *StorageNvmeRaidDriveGroup) SetDedicatedHotSparesForDriveGroup(v []StorageNvmeDedicatedHotSpareConfiguration)`

SetDedicatedHotSparesForDriveGroup sets DedicatedHotSparesForDriveGroup field to given value.

### HasDedicatedHotSparesForDriveGroup

`func (o *StorageNvmeRaidDriveGroup) HasDedicatedHotSparesForDriveGroup() bool`

HasDedicatedHotSparesForDriveGroup returns a boolean if a field has been set.

### SetDedicatedHotSparesForDriveGroupNil

`func (o *StorageNvmeRaidDriveGroup) SetDedicatedHotSparesForDriveGroupNil(b bool)`

 SetDedicatedHotSparesForDriveGroupNil sets the value for DedicatedHotSparesForDriveGroup to be an explicit nil

### UnsetDedicatedHotSparesForDriveGroup
`func (o *StorageNvmeRaidDriveGroup) UnsetDedicatedHotSparesForDriveGroup()`

UnsetDedicatedHotSparesForDriveGroup ensures that no value is present for DedicatedHotSparesForDriveGroup, not even an explicit nil
### GetName

`func (o *StorageNvmeRaidDriveGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNvmeRaidDriveGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNvmeRaidDriveGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNvmeRaidDriveGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


