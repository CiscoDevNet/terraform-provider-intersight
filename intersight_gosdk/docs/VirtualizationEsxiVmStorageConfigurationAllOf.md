# VirtualizationEsxiVmStorageConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiVmStorageConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiVmStorageConfiguration"]
**Datastore** | Pointer to **string** | Datastore where virtual machine is deployed. | [optional] 
**Disks** | Pointer to [**[]VirtualizationVmEsxiDisk**](VirtualizationVmEsxiDisk.md) |  | [optional] 

## Methods

### NewVirtualizationEsxiVmStorageConfigurationAllOf

`func NewVirtualizationEsxiVmStorageConfigurationAllOf(classId string, objectType string, ) *VirtualizationEsxiVmStorageConfigurationAllOf`

NewVirtualizationEsxiVmStorageConfigurationAllOf instantiates a new VirtualizationEsxiVmStorageConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiVmStorageConfigurationAllOfWithDefaults

`func NewVirtualizationEsxiVmStorageConfigurationAllOfWithDefaults() *VirtualizationEsxiVmStorageConfigurationAllOf`

NewVirtualizationEsxiVmStorageConfigurationAllOfWithDefaults instantiates a new VirtualizationEsxiVmStorageConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatastore

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDisks

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDisks() []VirtualizationVmEsxiDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) GetDisksOk() (*[]VirtualizationVmEsxiDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetDisks(v []VirtualizationVmEsxiDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *VirtualizationEsxiVmStorageConfigurationAllOf) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


