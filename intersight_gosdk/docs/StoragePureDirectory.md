# StoragePureDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureDirectory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureDirectory"]
**CreatedTime** | Pointer to **time.Time** | The managed directory creation time, measured in milliseconds since the UNIX epoch. | [optional] [readonly] 
**DataReduction** | Pointer to **int64** | The ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5:1 means that for every 5 MB the host writes to the array, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Destroyed** | Pointer to **bool** | A boolean value, if set to true, lists only destroyed objects that are in the eradication pending state. If set to false, lists only objects that are not destroyed. | [optional] [readonly] 
**DirectoryName** | Pointer to **string** | The managed directory name without the file system name prefix. | [optional] [readonly] 
**ExportPolicyList** | Pointer to **[]string** |  | [optional] 
**FileSystemName** | Pointer to **string** | Name of file syetem associated with the directory. | [optional] [readonly] 
**MemberName** | Pointer to **string** | Absolute path of the managed directory in the file system. | [optional] [readonly] 
**MemberResourceType** | Pointer to **string** | Absolute path of the managed directory in the file system. | [optional] [readonly] 
**Name** | Pointer to **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [readonly] 
**Path** | Pointer to **string** | Absolute path of the managed directory in the file system. | [optional] [readonly] 
**PolicyResourceType** | Pointer to **string** | Absolute path of the managed directory in the file system. | [optional] [readonly] 
**QuotaPolicyNames** | Pointer to **[]string** |  | [optional] 
**Snapshots** | Pointer to **int64** | The physical space occupied by data unique to one or more snapshots. Measured in bytes. | [optional] [readonly] 
**ThinProvisioning** | Pointer to **int64** | The percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed. | [optional] [readonly] 
**TotalPhysical** | Pointer to **int64** | The total physical space occupied by system, shared space, volume, and snapshot data. Measured in bytes. | [optional] [readonly] 
**TotalReduction** | Pointer to **int64** | The ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10:1 means that for every 10 MB of provisioned space, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**TotalUsed** | Pointer to **int64** | The total space contributed by customer data, measured in bytes. | [optional] [readonly] 
**Unique** | Pointer to **int64** | The unique physical space occupied by customer data. Unique physical space does not include shared space, snapshots, and internal array metadata. Measured in bytes. On Evergreen//One arrays, this is the effective space contributed by unique customer data, measured in bytes. Unique data does not include shared space, snapshots, and internal array metadata. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**ExportPolicies** | Pointer to [**[]StoragePureDirectoryExportRelationship**](StoragePureDirectoryExportRelationship.md) | An array of relationships to storagePureDirectoryExport resources. | [optional] [readonly] 
**FileSystems** | Pointer to [**NullableStoragePureFileSystemsRelationship**](StoragePureFileSystemsRelationship.md) |  | [optional] 
**QuotaPolicies** | Pointer to [**[]StoragePureDirectoryQuotaRelationship**](StoragePureDirectoryQuotaRelationship.md) | An array of relationships to storagePureDirectoryQuota resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureDirectory

`func NewStoragePureDirectory(classId string, objectType string, ) *StoragePureDirectory`

NewStoragePureDirectory instantiates a new StoragePureDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureDirectoryWithDefaults

`func NewStoragePureDirectoryWithDefaults() *StoragePureDirectory`

NewStoragePureDirectoryWithDefaults instantiates a new StoragePureDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureDirectory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureDirectory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureDirectory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureDirectory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureDirectory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureDirectory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedTime

`func (o *StoragePureDirectory) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StoragePureDirectory) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StoragePureDirectory) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StoragePureDirectory) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDataReduction

`func (o *StoragePureDirectory) GetDataReduction() int64`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StoragePureDirectory) GetDataReductionOk() (*int64, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StoragePureDirectory) SetDataReduction(v int64)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StoragePureDirectory) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetDestroyed

`func (o *StoragePureDirectory) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureDirectory) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureDirectory) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureDirectory) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetDirectoryName

`func (o *StoragePureDirectory) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *StoragePureDirectory) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *StoragePureDirectory) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *StoragePureDirectory) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### GetExportPolicyList

`func (o *StoragePureDirectory) GetExportPolicyList() []string`

GetExportPolicyList returns the ExportPolicyList field if non-nil, zero value otherwise.

### GetExportPolicyListOk

`func (o *StoragePureDirectory) GetExportPolicyListOk() (*[]string, bool)`

GetExportPolicyListOk returns a tuple with the ExportPolicyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicyList

`func (o *StoragePureDirectory) SetExportPolicyList(v []string)`

SetExportPolicyList sets ExportPolicyList field to given value.

### HasExportPolicyList

`func (o *StoragePureDirectory) HasExportPolicyList() bool`

HasExportPolicyList returns a boolean if a field has been set.

### SetExportPolicyListNil

`func (o *StoragePureDirectory) SetExportPolicyListNil(b bool)`

 SetExportPolicyListNil sets the value for ExportPolicyList to be an explicit nil

### UnsetExportPolicyList
`func (o *StoragePureDirectory) UnsetExportPolicyList()`

UnsetExportPolicyList ensures that no value is present for ExportPolicyList, not even an explicit nil
### GetFileSystemName

`func (o *StoragePureDirectory) GetFileSystemName() string`

GetFileSystemName returns the FileSystemName field if non-nil, zero value otherwise.

### GetFileSystemNameOk

`func (o *StoragePureDirectory) GetFileSystemNameOk() (*string, bool)`

GetFileSystemNameOk returns a tuple with the FileSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemName

`func (o *StoragePureDirectory) SetFileSystemName(v string)`

SetFileSystemName sets FileSystemName field to given value.

### HasFileSystemName

`func (o *StoragePureDirectory) HasFileSystemName() bool`

HasFileSystemName returns a boolean if a field has been set.

### GetMemberName

`func (o *StoragePureDirectory) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *StoragePureDirectory) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *StoragePureDirectory) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *StoragePureDirectory) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### GetMemberResourceType

`func (o *StoragePureDirectory) GetMemberResourceType() string`

GetMemberResourceType returns the MemberResourceType field if non-nil, zero value otherwise.

### GetMemberResourceTypeOk

`func (o *StoragePureDirectory) GetMemberResourceTypeOk() (*string, bool)`

GetMemberResourceTypeOk returns a tuple with the MemberResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberResourceType

`func (o *StoragePureDirectory) SetMemberResourceType(v string)`

SetMemberResourceType sets MemberResourceType field to given value.

### HasMemberResourceType

`func (o *StoragePureDirectory) HasMemberResourceType() bool`

HasMemberResourceType returns a boolean if a field has been set.

### GetName

`func (o *StoragePureDirectory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureDirectory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureDirectory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureDirectory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *StoragePureDirectory) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StoragePureDirectory) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StoragePureDirectory) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StoragePureDirectory) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPolicyResourceType

`func (o *StoragePureDirectory) GetPolicyResourceType() string`

GetPolicyResourceType returns the PolicyResourceType field if non-nil, zero value otherwise.

### GetPolicyResourceTypeOk

`func (o *StoragePureDirectory) GetPolicyResourceTypeOk() (*string, bool)`

GetPolicyResourceTypeOk returns a tuple with the PolicyResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyResourceType

`func (o *StoragePureDirectory) SetPolicyResourceType(v string)`

SetPolicyResourceType sets PolicyResourceType field to given value.

### HasPolicyResourceType

`func (o *StoragePureDirectory) HasPolicyResourceType() bool`

HasPolicyResourceType returns a boolean if a field has been set.

### GetQuotaPolicyNames

`func (o *StoragePureDirectory) GetQuotaPolicyNames() []string`

GetQuotaPolicyNames returns the QuotaPolicyNames field if non-nil, zero value otherwise.

### GetQuotaPolicyNamesOk

`func (o *StoragePureDirectory) GetQuotaPolicyNamesOk() (*[]string, bool)`

GetQuotaPolicyNamesOk returns a tuple with the QuotaPolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaPolicyNames

`func (o *StoragePureDirectory) SetQuotaPolicyNames(v []string)`

SetQuotaPolicyNames sets QuotaPolicyNames field to given value.

### HasQuotaPolicyNames

`func (o *StoragePureDirectory) HasQuotaPolicyNames() bool`

HasQuotaPolicyNames returns a boolean if a field has been set.

### SetQuotaPolicyNamesNil

`func (o *StoragePureDirectory) SetQuotaPolicyNamesNil(b bool)`

 SetQuotaPolicyNamesNil sets the value for QuotaPolicyNames to be an explicit nil

### UnsetQuotaPolicyNames
`func (o *StoragePureDirectory) UnsetQuotaPolicyNames()`

UnsetQuotaPolicyNames ensures that no value is present for QuotaPolicyNames, not even an explicit nil
### GetSnapshots

`func (o *StoragePureDirectory) GetSnapshots() int64`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *StoragePureDirectory) GetSnapshotsOk() (*int64, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *StoragePureDirectory) SetSnapshots(v int64)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *StoragePureDirectory) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetThinProvisioning

`func (o *StoragePureDirectory) GetThinProvisioning() int64`

GetThinProvisioning returns the ThinProvisioning field if non-nil, zero value otherwise.

### GetThinProvisioningOk

`func (o *StoragePureDirectory) GetThinProvisioningOk() (*int64, bool)`

GetThinProvisioningOk returns a tuple with the ThinProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioning

`func (o *StoragePureDirectory) SetThinProvisioning(v int64)`

SetThinProvisioning sets ThinProvisioning field to given value.

### HasThinProvisioning

`func (o *StoragePureDirectory) HasThinProvisioning() bool`

HasThinProvisioning returns a boolean if a field has been set.

### GetTotalPhysical

`func (o *StoragePureDirectory) GetTotalPhysical() int64`

GetTotalPhysical returns the TotalPhysical field if non-nil, zero value otherwise.

### GetTotalPhysicalOk

`func (o *StoragePureDirectory) GetTotalPhysicalOk() (*int64, bool)`

GetTotalPhysicalOk returns a tuple with the TotalPhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysical

`func (o *StoragePureDirectory) SetTotalPhysical(v int64)`

SetTotalPhysical sets TotalPhysical field to given value.

### HasTotalPhysical

`func (o *StoragePureDirectory) HasTotalPhysical() bool`

HasTotalPhysical returns a boolean if a field has been set.

### GetTotalReduction

`func (o *StoragePureDirectory) GetTotalReduction() int64`

GetTotalReduction returns the TotalReduction field if non-nil, zero value otherwise.

### GetTotalReductionOk

`func (o *StoragePureDirectory) GetTotalReductionOk() (*int64, bool)`

GetTotalReductionOk returns a tuple with the TotalReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReduction

`func (o *StoragePureDirectory) SetTotalReduction(v int64)`

SetTotalReduction sets TotalReduction field to given value.

### HasTotalReduction

`func (o *StoragePureDirectory) HasTotalReduction() bool`

HasTotalReduction returns a boolean if a field has been set.

### GetTotalUsed

`func (o *StoragePureDirectory) GetTotalUsed() int64`

GetTotalUsed returns the TotalUsed field if non-nil, zero value otherwise.

### GetTotalUsedOk

`func (o *StoragePureDirectory) GetTotalUsedOk() (*int64, bool)`

GetTotalUsedOk returns a tuple with the TotalUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsed

`func (o *StoragePureDirectory) SetTotalUsed(v int64)`

SetTotalUsed sets TotalUsed field to given value.

### HasTotalUsed

`func (o *StoragePureDirectory) HasTotalUsed() bool`

HasTotalUsed returns a boolean if a field has been set.

### GetUnique

`func (o *StoragePureDirectory) GetUnique() int64`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *StoragePureDirectory) GetUniqueOk() (*int64, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *StoragePureDirectory) SetUnique(v int64)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *StoragePureDirectory) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureDirectory) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureDirectory) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureDirectory) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureDirectory) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureDirectory) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureDirectory) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetExportPolicies

`func (o *StoragePureDirectory) GetExportPolicies() []StoragePureDirectoryExportRelationship`

GetExportPolicies returns the ExportPolicies field if non-nil, zero value otherwise.

### GetExportPoliciesOk

`func (o *StoragePureDirectory) GetExportPoliciesOk() (*[]StoragePureDirectoryExportRelationship, bool)`

GetExportPoliciesOk returns a tuple with the ExportPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicies

`func (o *StoragePureDirectory) SetExportPolicies(v []StoragePureDirectoryExportRelationship)`

SetExportPolicies sets ExportPolicies field to given value.

### HasExportPolicies

`func (o *StoragePureDirectory) HasExportPolicies() bool`

HasExportPolicies returns a boolean if a field has been set.

### SetExportPoliciesNil

`func (o *StoragePureDirectory) SetExportPoliciesNil(b bool)`

 SetExportPoliciesNil sets the value for ExportPolicies to be an explicit nil

### UnsetExportPolicies
`func (o *StoragePureDirectory) UnsetExportPolicies()`

UnsetExportPolicies ensures that no value is present for ExportPolicies, not even an explicit nil
### GetFileSystems

`func (o *StoragePureDirectory) GetFileSystems() StoragePureFileSystemsRelationship`

GetFileSystems returns the FileSystems field if non-nil, zero value otherwise.

### GetFileSystemsOk

`func (o *StoragePureDirectory) GetFileSystemsOk() (*StoragePureFileSystemsRelationship, bool)`

GetFileSystemsOk returns a tuple with the FileSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystems

`func (o *StoragePureDirectory) SetFileSystems(v StoragePureFileSystemsRelationship)`

SetFileSystems sets FileSystems field to given value.

### HasFileSystems

`func (o *StoragePureDirectory) HasFileSystems() bool`

HasFileSystems returns a boolean if a field has been set.

### SetFileSystemsNil

`func (o *StoragePureDirectory) SetFileSystemsNil(b bool)`

 SetFileSystemsNil sets the value for FileSystems to be an explicit nil

### UnsetFileSystems
`func (o *StoragePureDirectory) UnsetFileSystems()`

UnsetFileSystems ensures that no value is present for FileSystems, not even an explicit nil
### GetQuotaPolicies

`func (o *StoragePureDirectory) GetQuotaPolicies() []StoragePureDirectoryQuotaRelationship`

GetQuotaPolicies returns the QuotaPolicies field if non-nil, zero value otherwise.

### GetQuotaPoliciesOk

`func (o *StoragePureDirectory) GetQuotaPoliciesOk() (*[]StoragePureDirectoryQuotaRelationship, bool)`

GetQuotaPoliciesOk returns a tuple with the QuotaPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaPolicies

`func (o *StoragePureDirectory) SetQuotaPolicies(v []StoragePureDirectoryQuotaRelationship)`

SetQuotaPolicies sets QuotaPolicies field to given value.

### HasQuotaPolicies

`func (o *StoragePureDirectory) HasQuotaPolicies() bool`

HasQuotaPolicies returns a boolean if a field has been set.

### SetQuotaPoliciesNil

`func (o *StoragePureDirectory) SetQuotaPoliciesNil(b bool)`

 SetQuotaPoliciesNil sets the value for QuotaPolicies to be an explicit nil

### UnsetQuotaPolicies
`func (o *StoragePureDirectory) UnsetQuotaPolicies()`

UnsetQuotaPolicies ensures that no value is present for QuotaPolicies, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureDirectory) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureDirectory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureDirectory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureDirectory) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureDirectory) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureDirectory) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


