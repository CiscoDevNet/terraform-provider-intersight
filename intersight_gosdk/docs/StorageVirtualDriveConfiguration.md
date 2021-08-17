# StorageVirtualDriveConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.VirtualDriveConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.VirtualDriveConfiguration"]
**BootDrive** | Pointer to **bool** | This flag enables this virtual drive to be used as a boot drive. | [optional] 
**ExpandToAvailable** | Pointer to **bool** | This flag enables the virtual drive to use all the space available in the disk group. When this flag is enabled, the size property is ignored. | [optional] 
**Name** | Pointer to **string** | The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. | [optional] 
**Size** | Pointer to **int64** | Virtual drive size in MebiBytes. Size is mandatory field except when the Expand to Available option is enabled. | [optional] 
**VirtualDrivePolicy** | Pointer to [**NullableStorageVirtualDrivePolicy**](storage.VirtualDrivePolicy.md) |  | [optional] 

## Methods

### NewStorageVirtualDriveConfiguration

`func NewStorageVirtualDriveConfiguration(classId string, objectType string, ) *StorageVirtualDriveConfiguration`

NewStorageVirtualDriveConfiguration instantiates a new StorageVirtualDriveConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveConfigurationWithDefaults

`func NewStorageVirtualDriveConfigurationWithDefaults() *StorageVirtualDriveConfiguration`

NewStorageVirtualDriveConfigurationWithDefaults instantiates a new StorageVirtualDriveConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageVirtualDriveConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageVirtualDriveConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageVirtualDriveConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageVirtualDriveConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageVirtualDriveConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageVirtualDriveConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootDrive

`func (o *StorageVirtualDriveConfiguration) GetBootDrive() bool`

GetBootDrive returns the BootDrive field if non-nil, zero value otherwise.

### GetBootDriveOk

`func (o *StorageVirtualDriveConfiguration) GetBootDriveOk() (*bool, bool)`

GetBootDriveOk returns a tuple with the BootDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDrive

`func (o *StorageVirtualDriveConfiguration) SetBootDrive(v bool)`

SetBootDrive sets BootDrive field to given value.

### HasBootDrive

`func (o *StorageVirtualDriveConfiguration) HasBootDrive() bool`

HasBootDrive returns a boolean if a field has been set.

### GetExpandToAvailable

`func (o *StorageVirtualDriveConfiguration) GetExpandToAvailable() bool`

GetExpandToAvailable returns the ExpandToAvailable field if non-nil, zero value otherwise.

### GetExpandToAvailableOk

`func (o *StorageVirtualDriveConfiguration) GetExpandToAvailableOk() (*bool, bool)`

GetExpandToAvailableOk returns a tuple with the ExpandToAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandToAvailable

`func (o *StorageVirtualDriveConfiguration) SetExpandToAvailable(v bool)`

SetExpandToAvailable sets ExpandToAvailable field to given value.

### HasExpandToAvailable

`func (o *StorageVirtualDriveConfiguration) HasExpandToAvailable() bool`

HasExpandToAvailable returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDriveConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageVirtualDriveConfiguration) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageVirtualDriveConfiguration) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageVirtualDriveConfiguration) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageVirtualDriveConfiguration) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVirtualDrivePolicy

`func (o *StorageVirtualDriveConfiguration) GetVirtualDrivePolicy() StorageVirtualDrivePolicy`

GetVirtualDrivePolicy returns the VirtualDrivePolicy field if non-nil, zero value otherwise.

### GetVirtualDrivePolicyOk

`func (o *StorageVirtualDriveConfiguration) GetVirtualDrivePolicyOk() (*StorageVirtualDrivePolicy, bool)`

GetVirtualDrivePolicyOk returns a tuple with the VirtualDrivePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrivePolicy

`func (o *StorageVirtualDriveConfiguration) SetVirtualDrivePolicy(v StorageVirtualDrivePolicy)`

SetVirtualDrivePolicy sets VirtualDrivePolicy field to given value.

### HasVirtualDrivePolicy

`func (o *StorageVirtualDriveConfiguration) HasVirtualDrivePolicy() bool`

HasVirtualDrivePolicy returns a boolean if a field has been set.

### SetVirtualDrivePolicyNil

`func (o *StorageVirtualDriveConfiguration) SetVirtualDrivePolicyNil(b bool)`

 SetVirtualDrivePolicyNil sets the value for VirtualDrivePolicy to be an explicit nil

### UnsetVirtualDrivePolicy
`func (o *StorageVirtualDriveConfiguration) UnsetVirtualDrivePolicy()`

UnsetVirtualDrivePolicy ensures that no value is present for VirtualDrivePolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


