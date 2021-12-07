# VirtualizationIweVirtualDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweVirtualDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweVirtualDisk"]
**AccessMode** | Pointer to **string** | Access mode of the virtual disk. * &#x60;ReadWriteOnce&#x60; - Read write permisisons to a Virtual disk by a single virtual machine. * &#x60;ReadWriteMany&#x60; - Read write permisisons to a Virtual disk by multiple virtual machines. * &#x60;ReadOnlyMany&#x60; - Read only permisisons to a Virtual disk by multiple virtual machines. * &#x60;&#x60; - Unknown disk access mode. | [optional] [readonly] [default to "ReadWriteOnce"]
**Mode** | Pointer to **string** | File mode of the disk  example - Filesystem, Block. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [readonly] [default to "Block"]
**SourceFilePath** | Pointer to **string** | Source file path associated with virtual machine disk. | [optional] [readonly] 
**SourceVirtualDisk** | Pointer to **string** | Virtual disk used for cloning new disk. | [optional] 
**Status** | Pointer to [**NullableVirtualizationDiskStatus**](VirtualizationDiskStatus.md) |  | [optional] 
**Uuid** | Pointer to **string** | UUID of the virtual disk. | [optional] [readonly] 
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**VirtualizationIweVirtualMachineRelationship**](VirtualizationIweVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweVirtualDiskAllOf

`func NewVirtualizationIweVirtualDiskAllOf(classId string, objectType string, ) *VirtualizationIweVirtualDiskAllOf`

NewVirtualizationIweVirtualDiskAllOf instantiates a new VirtualizationIweVirtualDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweVirtualDiskAllOfWithDefaults

`func NewVirtualizationIweVirtualDiskAllOfWithDefaults() *VirtualizationIweVirtualDiskAllOf`

NewVirtualizationIweVirtualDiskAllOfWithDefaults instantiates a new VirtualizationIweVirtualDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweVirtualDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweVirtualDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweVirtualDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweVirtualDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessMode

`func (o *VirtualizationIweVirtualDiskAllOf) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *VirtualizationIweVirtualDiskAllOf) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *VirtualizationIweVirtualDiskAllOf) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationIweVirtualDiskAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationIweVirtualDiskAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationIweVirtualDiskAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *VirtualizationIweVirtualDiskAllOf) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *VirtualizationIweVirtualDiskAllOf) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *VirtualizationIweVirtualDiskAllOf) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetSourceVirtualDisk

`func (o *VirtualizationIweVirtualDiskAllOf) GetSourceVirtualDisk() string`

GetSourceVirtualDisk returns the SourceVirtualDisk field if non-nil, zero value otherwise.

### GetSourceVirtualDiskOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetSourceVirtualDiskOk() (*string, bool)`

GetSourceVirtualDiskOk returns a tuple with the SourceVirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVirtualDisk

`func (o *VirtualizationIweVirtualDiskAllOf) SetSourceVirtualDisk(v string)`

SetSourceVirtualDisk sets SourceVirtualDisk field to given value.

### HasSourceVirtualDisk

`func (o *VirtualizationIweVirtualDiskAllOf) HasSourceVirtualDisk() bool`

HasSourceVirtualDisk returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationIweVirtualDiskAllOf) GetStatus() VirtualizationDiskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetStatusOk() (*VirtualizationDiskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationIweVirtualDiskAllOf) SetStatus(v VirtualizationDiskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationIweVirtualDiskAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VirtualizationIweVirtualDiskAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VirtualizationIweVirtualDiskAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUuid

`func (o *VirtualizationIweVirtualDiskAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationIweVirtualDiskAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationIweVirtualDiskAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweVirtualDiskAllOf) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweVirtualDiskAllOf) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweVirtualDiskAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *VirtualizationIweVirtualDiskAllOf) GetVirtualMachine() VirtualizationIweVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualizationIweVirtualDiskAllOf) GetVirtualMachineOk() (*VirtualizationIweVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualizationIweVirtualDiskAllOf) SetVirtualMachine(v VirtualizationIweVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VirtualizationIweVirtualDiskAllOf) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


