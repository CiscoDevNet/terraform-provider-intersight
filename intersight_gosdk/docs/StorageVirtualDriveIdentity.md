# StorageVirtualDriveIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.VirtualDriveIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.VirtualDriveIdentity"]
**Name** | Pointer to **string** | The VirtualDrive Name which belongs to the Storage VirtualDrive. | [optional] [readonly] 
**ServerProfile** | Pointer to [**ServerProfileRelationship**](ServerProfileRelationship.md) |  | [optional] 
**StoragePolicy** | Pointer to [**StorageStoragePolicyRelationship**](StorageStoragePolicyRelationship.md) |  | [optional] 
**VirtualDrive** | Pointer to [**StorageVirtualDriveRelationship**](StorageVirtualDriveRelationship.md) |  | [optional] 

## Methods

### NewStorageVirtualDriveIdentity

`func NewStorageVirtualDriveIdentity(classId string, objectType string, ) *StorageVirtualDriveIdentity`

NewStorageVirtualDriveIdentity instantiates a new StorageVirtualDriveIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveIdentityWithDefaults

`func NewStorageVirtualDriveIdentityWithDefaults() *StorageVirtualDriveIdentity`

NewStorageVirtualDriveIdentityWithDefaults instantiates a new StorageVirtualDriveIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageVirtualDriveIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageVirtualDriveIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageVirtualDriveIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageVirtualDriveIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageVirtualDriveIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageVirtualDriveIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *StorageVirtualDriveIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerProfile

`func (o *StorageVirtualDriveIdentity) GetServerProfile() ServerProfileRelationship`

GetServerProfile returns the ServerProfile field if non-nil, zero value otherwise.

### GetServerProfileOk

`func (o *StorageVirtualDriveIdentity) GetServerProfileOk() (*ServerProfileRelationship, bool)`

GetServerProfileOk returns a tuple with the ServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProfile

`func (o *StorageVirtualDriveIdentity) SetServerProfile(v ServerProfileRelationship)`

SetServerProfile sets ServerProfile field to given value.

### HasServerProfile

`func (o *StorageVirtualDriveIdentity) HasServerProfile() bool`

HasServerProfile returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *StorageVirtualDriveIdentity) GetStoragePolicy() StorageStoragePolicyRelationship`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *StorageVirtualDriveIdentity) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *StorageVirtualDriveIdentity) SetStoragePolicy(v StorageStoragePolicyRelationship)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *StorageVirtualDriveIdentity) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageVirtualDriveIdentity) GetVirtualDrive() StorageVirtualDriveRelationship`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageVirtualDriveIdentity) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageVirtualDriveIdentity) SetVirtualDrive(v StorageVirtualDriveRelationship)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageVirtualDriveIdentity) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


