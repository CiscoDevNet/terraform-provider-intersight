# StorageNetAppQtreeAllOf

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
**VolumeUuid** | Pointer to **string** | NetApp Volume uuid, unique identifier for the NetApp volume. | [optional] [readonly] 
**StorageContainer** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppQtreeAllOf

`func NewStorageNetAppQtreeAllOf(classId string, objectType string, ) *StorageNetAppQtreeAllOf`

NewStorageNetAppQtreeAllOf instantiates a new StorageNetAppQtreeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppQtreeAllOfWithDefaults

`func NewStorageNetAppQtreeAllOfWithDefaults() *StorageNetAppQtreeAllOf`

NewStorageNetAppQtreeAllOfWithDefaults instantiates a new StorageNetAppQtreeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppQtreeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppQtreeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppQtreeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppQtreeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppQtreeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppQtreeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExportPolicyId

`func (o *StorageNetAppQtreeAllOf) GetExportPolicyId() string`

GetExportPolicyId returns the ExportPolicyId field if non-nil, zero value otherwise.

### GetExportPolicyIdOk

`func (o *StorageNetAppQtreeAllOf) GetExportPolicyIdOk() (*string, bool)`

GetExportPolicyIdOk returns a tuple with the ExportPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicyId

`func (o *StorageNetAppQtreeAllOf) SetExportPolicyId(v string)`

SetExportPolicyId sets ExportPolicyId field to given value.

### HasExportPolicyId

`func (o *StorageNetAppQtreeAllOf) HasExportPolicyId() bool`

HasExportPolicyId returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppQtreeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppQtreeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppQtreeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppQtreeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *StorageNetAppQtreeAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageNetAppQtreeAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageNetAppQtreeAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageNetAppQtreeAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermission

`func (o *StorageNetAppQtreeAllOf) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *StorageNetAppQtreeAllOf) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *StorageNetAppQtreeAllOf) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *StorageNetAppQtreeAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetQtreeId

`func (o *StorageNetAppQtreeAllOf) GetQtreeId() int64`

GetQtreeId returns the QtreeId field if non-nil, zero value otherwise.

### GetQtreeIdOk

`func (o *StorageNetAppQtreeAllOf) GetQtreeIdOk() (*int64, bool)`

GetQtreeIdOk returns a tuple with the QtreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtreeId

`func (o *StorageNetAppQtreeAllOf) SetQtreeId(v int64)`

SetQtreeId sets QtreeId field to given value.

### HasQtreeId

`func (o *StorageNetAppQtreeAllOf) HasQtreeId() bool`

HasQtreeId returns a boolean if a field has been set.

### GetSecurityStyle

`func (o *StorageNetAppQtreeAllOf) GetSecurityStyle() string`

GetSecurityStyle returns the SecurityStyle field if non-nil, zero value otherwise.

### GetSecurityStyleOk

`func (o *StorageNetAppQtreeAllOf) GetSecurityStyleOk() (*string, bool)`

GetSecurityStyleOk returns a tuple with the SecurityStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStyle

`func (o *StorageNetAppQtreeAllOf) SetSecurityStyle(v string)`

SetSecurityStyle sets SecurityStyle field to given value.

### HasSecurityStyle

`func (o *StorageNetAppQtreeAllOf) HasSecurityStyle() bool`

HasSecurityStyle returns a boolean if a field has been set.

### GetVolumeUuid

`func (o *StorageNetAppQtreeAllOf) GetVolumeUuid() string`

GetVolumeUuid returns the VolumeUuid field if non-nil, zero value otherwise.

### GetVolumeUuidOk

`func (o *StorageNetAppQtreeAllOf) GetVolumeUuidOk() (*string, bool)`

GetVolumeUuidOk returns a tuple with the VolumeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUuid

`func (o *StorageNetAppQtreeAllOf) SetVolumeUuid(v string)`

SetVolumeUuid sets VolumeUuid field to given value.

### HasVolumeUuid

`func (o *StorageNetAppQtreeAllOf) HasVolumeUuid() bool`

HasVolumeUuid returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageNetAppQtreeAllOf) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppQtreeAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppQtreeAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppQtreeAllOf) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppQtreeAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppQtreeAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppQtreeAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppQtreeAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


