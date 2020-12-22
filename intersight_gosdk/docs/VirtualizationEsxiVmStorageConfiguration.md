# VirtualizationEsxiVmStorageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiVmStorageConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiVmStorageConfiguration"]
**Datastore** | Pointer to **string** | Datastore where virtual machine is deployed. | [optional] 
**Disks** | Pointer to [**[]VirtualizationVmEsxiDisk**](VirtualizationVmEsxiDisk.md) |  | [optional] 

## Methods

### NewVirtualizationEsxiVmStorageConfiguration

`func NewVirtualizationEsxiVmStorageConfiguration(classId string, objectType string, ) *VirtualizationEsxiVmStorageConfiguration`

NewVirtualizationEsxiVmStorageConfiguration instantiates a new VirtualizationEsxiVmStorageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiVmStorageConfigurationWithDefaults

`func NewVirtualizationEsxiVmStorageConfigurationWithDefaults() *VirtualizationEsxiVmStorageConfiguration`

NewVirtualizationEsxiVmStorageConfigurationWithDefaults instantiates a new VirtualizationEsxiVmStorageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiVmStorageConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiVmStorageConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiVmStorageConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiVmStorageConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiVmStorageConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiVmStorageConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatastore

`func (o *VirtualizationEsxiVmStorageConfiguration) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VirtualizationEsxiVmStorageConfiguration) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VirtualizationEsxiVmStorageConfiguration) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VirtualizationEsxiVmStorageConfiguration) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDisks

`func (o *VirtualizationEsxiVmStorageConfiguration) GetDisks() []VirtualizationVmEsxiDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VirtualizationEsxiVmStorageConfiguration) GetDisksOk() (*[]VirtualizationVmEsxiDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VirtualizationEsxiVmStorageConfiguration) SetDisks(v []VirtualizationVmEsxiDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VirtualizationEsxiVmStorageConfiguration) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *VirtualizationEsxiVmStorageConfiguration) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *VirtualizationEsxiVmStorageConfiguration) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


