# StorageNetAppQtree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppQtree"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppQtree"]
**ExportPolicyId** | Pointer to **string** | Unique identifier of NetApp export policy. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the NetApp Qtree. | [optional] [readonly] 
**Path** | Pointer to **string** | Client visible path to the qtree. | [optional] [readonly] 
**Permission** | Pointer to **string** | Identifies the UNIX permissions for the qtree. | [optional] [readonly] 
**QtreeId** | Pointer to **int64** | NetApp Qtree ID, unique within the qtree&#39;s volume. | [optional] [readonly] 
**SecurityStyle** | Pointer to **string** | Identifies the security style for the qtree, it determines how access to the qtree is controlled. * &#x60;UNIX&#x60; - Security style for UNIX uid, gid and mode bits. * &#x60;NTFS&#x60; - Security style for CIFS ACLs. * &#x60;Mixed&#x60; - Security style for NFS and CIFS access. | [optional] [readonly] [default to "UNIX"]
**SvmName** | Pointer to **string** | The storage virtual machine name for the qtree. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | The parent volume name for the qtree. | [optional] [readonly] 
**VolumeUuid** | Pointer to **string** | NetApp Volume uuid, unique identifier for the NetApp volume. | [optional] [readonly] 
**StorageContainer** | Pointer to [**NullableStorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppQtree

`func NewStorageNetAppQtree(classId string, objectType string, ) *StorageNetAppQtree`

NewStorageNetAppQtree instantiates a new StorageNetAppQtree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppQtreeWithDefaults

`func NewStorageNetAppQtreeWithDefaults() *StorageNetAppQtree`

NewStorageNetAppQtreeWithDefaults instantiates a new StorageNetAppQtree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppQtree) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppQtree) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppQtree) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppQtree) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppQtree) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppQtree) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExportPolicyId

`func (o *StorageNetAppQtree) GetExportPolicyId() string`

GetExportPolicyId returns the ExportPolicyId field if non-nil, zero value otherwise.

### GetExportPolicyIdOk

`func (o *StorageNetAppQtree) GetExportPolicyIdOk() (*string, bool)`

GetExportPolicyIdOk returns a tuple with the ExportPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicyId

`func (o *StorageNetAppQtree) SetExportPolicyId(v string)`

SetExportPolicyId sets ExportPolicyId field to given value.

### HasExportPolicyId

`func (o *StorageNetAppQtree) HasExportPolicyId() bool`

HasExportPolicyId returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppQtree) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppQtree) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppQtree) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppQtree) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *StorageNetAppQtree) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageNetAppQtree) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageNetAppQtree) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageNetAppQtree) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermission

`func (o *StorageNetAppQtree) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *StorageNetAppQtree) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *StorageNetAppQtree) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *StorageNetAppQtree) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetQtreeId

`func (o *StorageNetAppQtree) GetQtreeId() int64`

GetQtreeId returns the QtreeId field if non-nil, zero value otherwise.

### GetQtreeIdOk

`func (o *StorageNetAppQtree) GetQtreeIdOk() (*int64, bool)`

GetQtreeIdOk returns a tuple with the QtreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtreeId

`func (o *StorageNetAppQtree) SetQtreeId(v int64)`

SetQtreeId sets QtreeId field to given value.

### HasQtreeId

`func (o *StorageNetAppQtree) HasQtreeId() bool`

HasQtreeId returns a boolean if a field has been set.

### GetSecurityStyle

`func (o *StorageNetAppQtree) GetSecurityStyle() string`

GetSecurityStyle returns the SecurityStyle field if non-nil, zero value otherwise.

### GetSecurityStyleOk

`func (o *StorageNetAppQtree) GetSecurityStyleOk() (*string, bool)`

GetSecurityStyleOk returns a tuple with the SecurityStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStyle

`func (o *StorageNetAppQtree) SetSecurityStyle(v string)`

SetSecurityStyle sets SecurityStyle field to given value.

### HasSecurityStyle

`func (o *StorageNetAppQtree) HasSecurityStyle() bool`

HasSecurityStyle returns a boolean if a field has been set.

### GetSvmName

`func (o *StorageNetAppQtree) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppQtree) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppQtree) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppQtree) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageNetAppQtree) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageNetAppQtree) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageNetAppQtree) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageNetAppQtree) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeUuid

`func (o *StorageNetAppQtree) GetVolumeUuid() string`

GetVolumeUuid returns the VolumeUuid field if non-nil, zero value otherwise.

### GetVolumeUuidOk

`func (o *StorageNetAppQtree) GetVolumeUuidOk() (*string, bool)`

GetVolumeUuidOk returns a tuple with the VolumeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUuid

`func (o *StorageNetAppQtree) SetVolumeUuid(v string)`

SetVolumeUuid sets VolumeUuid field to given value.

### HasVolumeUuid

`func (o *StorageNetAppQtree) HasVolumeUuid() bool`

HasVolumeUuid returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageNetAppQtree) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppQtree) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppQtree) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppQtree) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### SetStorageContainerNil

`func (o *StorageNetAppQtree) SetStorageContainerNil(b bool)`

 SetStorageContainerNil sets the value for StorageContainer to be an explicit nil

### UnsetStorageContainer
`func (o *StorageNetAppQtree) UnsetStorageContainer()`

UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
### GetTenant

`func (o *StorageNetAppQtree) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppQtree) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppQtree) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppQtree) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *StorageNetAppQtree) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *StorageNetAppQtree) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


