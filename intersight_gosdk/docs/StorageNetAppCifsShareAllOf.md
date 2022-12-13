# StorageNetAppCifsShareAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppCifsShare"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppCifsShare"]
**Comment** | Pointer to **string** | Description of the CIFS share. | [optional] [readonly] 
**Encryption** | Pointer to **string** | Indicates that SMB encryption must be used when accessing the share. | [optional] [readonly] 
**HomeDirectory** | Pointer to **string** | Indicates that the share is a home directory share, where the share and path names are dynamic. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the NetApp CIFS share. | [optional] [readonly] 
**NetAppCifsAcl** | Pointer to [**[]StorageNetAppCifsAcl**](StorageNetAppCifsAcl.md) |  | [optional] 
**Path** | Pointer to **string** | The fully-qualified pathname in the owning SVM namespace that is shared through the share. | [optional] [readonly] 
**SvmUuid** | Pointer to **string** | Unique identifier for the NetApp Storage Virtual Machine. | [optional] [readonly] 
**StorageContainer** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppCifsShareAllOf

`func NewStorageNetAppCifsShareAllOf(classId string, objectType string, ) *StorageNetAppCifsShareAllOf`

NewStorageNetAppCifsShareAllOf instantiates a new StorageNetAppCifsShareAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppCifsShareAllOfWithDefaults

`func NewStorageNetAppCifsShareAllOfWithDefaults() *StorageNetAppCifsShareAllOf`

NewStorageNetAppCifsShareAllOfWithDefaults instantiates a new StorageNetAppCifsShareAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppCifsShareAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppCifsShareAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppCifsShareAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppCifsShareAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppCifsShareAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppCifsShareAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComment

`func (o *StorageNetAppCifsShareAllOf) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StorageNetAppCifsShareAllOf) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StorageNetAppCifsShareAllOf) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StorageNetAppCifsShareAllOf) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEncryption

`func (o *StorageNetAppCifsShareAllOf) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *StorageNetAppCifsShareAllOf) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *StorageNetAppCifsShareAllOf) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *StorageNetAppCifsShareAllOf) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetHomeDirectory

`func (o *StorageNetAppCifsShareAllOf) GetHomeDirectory() string`

GetHomeDirectory returns the HomeDirectory field if non-nil, zero value otherwise.

### GetHomeDirectoryOk

`func (o *StorageNetAppCifsShareAllOf) GetHomeDirectoryOk() (*string, bool)`

GetHomeDirectoryOk returns a tuple with the HomeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDirectory

`func (o *StorageNetAppCifsShareAllOf) SetHomeDirectory(v string)`

SetHomeDirectory sets HomeDirectory field to given value.

### HasHomeDirectory

`func (o *StorageNetAppCifsShareAllOf) HasHomeDirectory() bool`

HasHomeDirectory returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppCifsShareAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppCifsShareAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppCifsShareAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppCifsShareAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetAppCifsAcl

`func (o *StorageNetAppCifsShareAllOf) GetNetAppCifsAcl() []StorageNetAppCifsAcl`

GetNetAppCifsAcl returns the NetAppCifsAcl field if non-nil, zero value otherwise.

### GetNetAppCifsAclOk

`func (o *StorageNetAppCifsShareAllOf) GetNetAppCifsAclOk() (*[]StorageNetAppCifsAcl, bool)`

GetNetAppCifsAclOk returns a tuple with the NetAppCifsAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppCifsAcl

`func (o *StorageNetAppCifsShareAllOf) SetNetAppCifsAcl(v []StorageNetAppCifsAcl)`

SetNetAppCifsAcl sets NetAppCifsAcl field to given value.

### HasNetAppCifsAcl

`func (o *StorageNetAppCifsShareAllOf) HasNetAppCifsAcl() bool`

HasNetAppCifsAcl returns a boolean if a field has been set.

### SetNetAppCifsAclNil

`func (o *StorageNetAppCifsShareAllOf) SetNetAppCifsAclNil(b bool)`

 SetNetAppCifsAclNil sets the value for NetAppCifsAcl to be an explicit nil

### UnsetNetAppCifsAcl
`func (o *StorageNetAppCifsShareAllOf) UnsetNetAppCifsAcl()`

UnsetNetAppCifsAcl ensures that no value is present for NetAppCifsAcl, not even an explicit nil
### GetPath

`func (o *StorageNetAppCifsShareAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageNetAppCifsShareAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageNetAppCifsShareAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageNetAppCifsShareAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSvmUuid

`func (o *StorageNetAppCifsShareAllOf) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppCifsShareAllOf) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppCifsShareAllOf) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppCifsShareAllOf) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageNetAppCifsShareAllOf) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppCifsShareAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppCifsShareAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppCifsShareAllOf) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppCifsShareAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppCifsShareAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppCifsShareAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppCifsShareAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


