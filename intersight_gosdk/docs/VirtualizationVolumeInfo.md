# VirtualizationVolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VolumeInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VolumeInfo"]
**Bootable** | Pointer to **bool** | Set to true, if the volume should be a root disk. | [optional] 
**DeleteOnTermination** | Pointer to **bool** | Set to true, to delete the volume on termination of the VM the volume is attached to. | [optional] 
**Encryption** | Pointer to **bool** | Set to true, if the volume should be encrypted. | [optional] 
**Iops** | Pointer to **int64** | IOPS for the volume for applicable volume types. | [optional] 
**Order** | Pointer to **int64** | Order of the disk attachment to the VM. | [optional] 
**VolumeId** | Pointer to **string** | Unique volume id assigned by the cloud provider. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | Name assigned to the volume created. | [optional] 
**VolumeSize** | Pointer to **int64** | Size of the volume created in GB. | [optional] 
**VolumeType** | Pointer to **string** | Id of the volume or storage type of this volume. | [optional] 

## Methods

### NewVirtualizationVolumeInfo

`func NewVirtualizationVolumeInfo(classId string, objectType string, ) *VirtualizationVolumeInfo`

NewVirtualizationVolumeInfo instantiates a new VirtualizationVolumeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVolumeInfoWithDefaults

`func NewVirtualizationVolumeInfoWithDefaults() *VirtualizationVolumeInfo`

NewVirtualizationVolumeInfoWithDefaults instantiates a new VirtualizationVolumeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVolumeInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVolumeInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVolumeInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVolumeInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVolumeInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVolumeInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootable

`func (o *VirtualizationVolumeInfo) GetBootable() bool`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *VirtualizationVolumeInfo) GetBootableOk() (*bool, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *VirtualizationVolumeInfo) SetBootable(v bool)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *VirtualizationVolumeInfo) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetDeleteOnTermination

`func (o *VirtualizationVolumeInfo) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *VirtualizationVolumeInfo) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *VirtualizationVolumeInfo) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.

### HasDeleteOnTermination

`func (o *VirtualizationVolumeInfo) HasDeleteOnTermination() bool`

HasDeleteOnTermination returns a boolean if a field has been set.

### GetEncryption

`func (o *VirtualizationVolumeInfo) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VirtualizationVolumeInfo) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VirtualizationVolumeInfo) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VirtualizationVolumeInfo) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetIops

`func (o *VirtualizationVolumeInfo) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *VirtualizationVolumeInfo) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *VirtualizationVolumeInfo) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *VirtualizationVolumeInfo) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetOrder

`func (o *VirtualizationVolumeInfo) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VirtualizationVolumeInfo) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VirtualizationVolumeInfo) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VirtualizationVolumeInfo) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetVolumeId

`func (o *VirtualizationVolumeInfo) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VirtualizationVolumeInfo) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VirtualizationVolumeInfo) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VirtualizationVolumeInfo) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetVolumeName

`func (o *VirtualizationVolumeInfo) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VirtualizationVolumeInfo) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VirtualizationVolumeInfo) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *VirtualizationVolumeInfo) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeSize

`func (o *VirtualizationVolumeInfo) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *VirtualizationVolumeInfo) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *VirtualizationVolumeInfo) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *VirtualizationVolumeInfo) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### GetVolumeType

`func (o *VirtualizationVolumeInfo) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VirtualizationVolumeInfo) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VirtualizationVolumeInfo) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *VirtualizationVolumeInfo) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


