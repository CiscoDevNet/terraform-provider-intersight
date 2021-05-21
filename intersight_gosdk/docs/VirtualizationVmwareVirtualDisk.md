# VirtualizationVmwareVirtualDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVirtualDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVirtualDisk"]
**CompatibilityMode** | Pointer to **string** | Compatibility mode of the raw disk mapping (RDM). * &#x60;notApplicable&#x60; - Value specified for any disk which is not of raw device mapping type. * &#x60;physicalMode&#x60; - A disk device backed by a physical compatibility mode raw disk mapping cannot use disk modes, and commands are passed straight through to the LUN indicated by the raw disk mapping. * &#x60;virtualMode&#x60; - A disk device backed by a virtual compatibility mode raw disk mapping can use disk modes. | [optional] [default to "notApplicable"]
**ControllerKey** | Pointer to **int64** | Key of the controller on which the disk is created. | [optional] 
**DeviceName** | Pointer to **string** | Host-specific device the LUN is being accessed through. If the target LUN is not available on the host then it is empty. For example, this could happen if it has accidentally been masked out. | [optional] 
**DiskMode** | Pointer to **string** | Persistence mode of the virtual disk. For RDM disks, it is only supported when the raw disk mapping is using virtual compatibility mode. * &#x60;persistent&#x60; - Changes are immediately and permanently written to the virtual disk. * &#x60;independent_persistent&#x60; - Changes are immediately and permanently written to the virtual disk and not affected by snapshots. * &#x60;independent_nonpersistent&#x60; - Changes to virtual disk are made to a redo log and discarded at power off and not affected by snapshots. * &#x60;nonpersistent&#x60; - Changes to virtual disk are made to a redo log and discarded at power off. * &#x60;undoable&#x60; - Changes to virtual disk are made to a redo log and has the option to commit or undo. * &#x60;append&#x60; - Changes are appended to the redo log and can be revoked by removing the undo log. | [optional] [default to "persistent"]
**DiskType** | Pointer to **string** | Specifies whether the virtual disk is a RDM or a Hard disk. * &#x60;flatDisk&#x60; - Specifies that it is a flat disk. * &#x60;rdmDisk&#x60; - Specifies that it is a raw device mapping disk. | [optional] [default to "flatDisk"]
**Key** | Pointer to **int64** | The internally assigned key of this disk. This entity is not manipulated by users. | [optional] 
**Limit** | Pointer to **int64** | The utilization of a virtual machine will not exceed this limit, even if there are available resources. Used to ensure a consistent performance of virtual machines independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). The unit is number of I/O per second. | [optional] 
**LunUuid** | Pointer to **string** | Unique identifier of the LUN accessed by the raw disk mapping (RDM). | [optional] 
**Serial** | Pointer to **string** | Serial ID of the storage device. | [optional] 
**Shares** | Pointer to [**NullableVirtualizationVmwareSharesInfo**](virtualization.VmwareSharesInfo.md) |  | [optional] 
**Sharing** | Pointer to **string** | Sharing mode of the virtual disk. * &#x60;sharingNone&#x60; - The virtual disk is not shared. * &#x60;sharingMultiWriter&#x60; - The virtual disk is shared between multiple virtual machines. | [optional] [default to "sharingNone"]
**StorageAllocationType** | Pointer to **string** | Allocation type for the virtual disk. * &#x60;notApplicable&#x60; - Value specified for a disk for which storage allocation type is not applicable. * &#x60;thin&#x60; - A thin provisioned disk consumes only the space that it needs for its initial operrations, and grows with time according to demand. It is the fastest method to create a virtual disk because it creates a disk with just the header information. * &#x60;lazyZeroedThick&#x60; - A lazy zeroed thick disk has all space allocated at the time of its creation. Data remaining on the physical device is not erased during creation, but is zeroed out on demand later on first write from the virtual machine. * &#x60;eagerZeroedThick&#x60; - An eager zeroed thick disk has all space allocated and wiped clean of any previous contents on the physical media at creation time. Such disks may take longer time during creation compared to other disk formats. | [optional] [default to "notApplicable"]
**UnitNumber** | Pointer to **int64** | Unit number of the disk on its controller. | [optional] 
**Uuid** | Pointer to **string** | UUID assigned by vCenter to every disk. | [optional] 
**VdiskId** | Pointer to **string** | Identity of the virtual disk object as the first class entity. | [optional] 
**Vendor** | Pointer to **string** | Vendor of the storage device. | [optional] 
**VirtualDiskPath** | Pointer to **string** | Path of the virtual disk. | [optional] 
**VmIdentity** | Pointer to **string** | Identity of the virtual machine where the virtual disk is created. | [optional] 
**Datastore** | Pointer to [**VirtualizationVmwareDatastoreRelationship**](virtualization.VmwareDatastore.Relationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**VirtualizationVmwareVirtualMachineRelationship**](virtualization.VmwareVirtualMachine.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareVirtualDisk

`func NewVirtualizationVmwareVirtualDisk(classId string, objectType string, ) *VirtualizationVmwareVirtualDisk`

NewVirtualizationVmwareVirtualDisk instantiates a new VirtualizationVmwareVirtualDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualDiskWithDefaults

`func NewVirtualizationVmwareVirtualDiskWithDefaults() *VirtualizationVmwareVirtualDisk`

NewVirtualizationVmwareVirtualDiskWithDefaults instantiates a new VirtualizationVmwareVirtualDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompatibilityMode

`func (o *VirtualizationVmwareVirtualDisk) GetCompatibilityMode() string`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *VirtualizationVmwareVirtualDisk) GetCompatibilityModeOk() (*string, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *VirtualizationVmwareVirtualDisk) SetCompatibilityMode(v string)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *VirtualizationVmwareVirtualDisk) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetControllerKey

`func (o *VirtualizationVmwareVirtualDisk) GetControllerKey() int64`

GetControllerKey returns the ControllerKey field if non-nil, zero value otherwise.

### GetControllerKeyOk

`func (o *VirtualizationVmwareVirtualDisk) GetControllerKeyOk() (*int64, bool)`

GetControllerKeyOk returns a tuple with the ControllerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerKey

`func (o *VirtualizationVmwareVirtualDisk) SetControllerKey(v int64)`

SetControllerKey sets ControllerKey field to given value.

### HasControllerKey

`func (o *VirtualizationVmwareVirtualDisk) HasControllerKey() bool`

HasControllerKey returns a boolean if a field has been set.

### GetDeviceName

`func (o *VirtualizationVmwareVirtualDisk) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *VirtualizationVmwareVirtualDisk) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *VirtualizationVmwareVirtualDisk) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *VirtualizationVmwareVirtualDisk) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDiskMode

`func (o *VirtualizationVmwareVirtualDisk) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *VirtualizationVmwareVirtualDisk) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *VirtualizationVmwareVirtualDisk) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *VirtualizationVmwareVirtualDisk) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *VirtualizationVmwareVirtualDisk) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *VirtualizationVmwareVirtualDisk) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *VirtualizationVmwareVirtualDisk) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *VirtualizationVmwareVirtualDisk) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetKey

`func (o *VirtualizationVmwareVirtualDisk) GetKey() int64`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VirtualizationVmwareVirtualDisk) GetKeyOk() (*int64, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VirtualizationVmwareVirtualDisk) SetKey(v int64)`

SetKey sets Key field to given value.

### HasKey

`func (o *VirtualizationVmwareVirtualDisk) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *VirtualizationVmwareVirtualDisk) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VirtualizationVmwareVirtualDisk) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VirtualizationVmwareVirtualDisk) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VirtualizationVmwareVirtualDisk) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLunUuid

`func (o *VirtualizationVmwareVirtualDisk) GetLunUuid() string`

GetLunUuid returns the LunUuid field if non-nil, zero value otherwise.

### GetLunUuidOk

`func (o *VirtualizationVmwareVirtualDisk) GetLunUuidOk() (*string, bool)`

GetLunUuidOk returns a tuple with the LunUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunUuid

`func (o *VirtualizationVmwareVirtualDisk) SetLunUuid(v string)`

SetLunUuid sets LunUuid field to given value.

### HasLunUuid

`func (o *VirtualizationVmwareVirtualDisk) HasLunUuid() bool`

HasLunUuid returns a boolean if a field has been set.

### GetSerial

`func (o *VirtualizationVmwareVirtualDisk) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualizationVmwareVirtualDisk) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualizationVmwareVirtualDisk) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualizationVmwareVirtualDisk) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetShares

`func (o *VirtualizationVmwareVirtualDisk) GetShares() VirtualizationVmwareSharesInfo`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *VirtualizationVmwareVirtualDisk) GetSharesOk() (*VirtualizationVmwareSharesInfo, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *VirtualizationVmwareVirtualDisk) SetShares(v VirtualizationVmwareSharesInfo)`

SetShares sets Shares field to given value.

### HasShares

`func (o *VirtualizationVmwareVirtualDisk) HasShares() bool`

HasShares returns a boolean if a field has been set.

### SetSharesNil

`func (o *VirtualizationVmwareVirtualDisk) SetSharesNil(b bool)`

 SetSharesNil sets the value for Shares to be an explicit nil

### UnsetShares
`func (o *VirtualizationVmwareVirtualDisk) UnsetShares()`

UnsetShares ensures that no value is present for Shares, not even an explicit nil
### GetSharing

`func (o *VirtualizationVmwareVirtualDisk) GetSharing() string`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *VirtualizationVmwareVirtualDisk) GetSharingOk() (*string, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *VirtualizationVmwareVirtualDisk) SetSharing(v string)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *VirtualizationVmwareVirtualDisk) HasSharing() bool`

HasSharing returns a boolean if a field has been set.

### GetStorageAllocationType

`func (o *VirtualizationVmwareVirtualDisk) GetStorageAllocationType() string`

GetStorageAllocationType returns the StorageAllocationType field if non-nil, zero value otherwise.

### GetStorageAllocationTypeOk

`func (o *VirtualizationVmwareVirtualDisk) GetStorageAllocationTypeOk() (*string, bool)`

GetStorageAllocationTypeOk returns a tuple with the StorageAllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAllocationType

`func (o *VirtualizationVmwareVirtualDisk) SetStorageAllocationType(v string)`

SetStorageAllocationType sets StorageAllocationType field to given value.

### HasStorageAllocationType

`func (o *VirtualizationVmwareVirtualDisk) HasStorageAllocationType() bool`

HasStorageAllocationType returns a boolean if a field has been set.

### GetUnitNumber

`func (o *VirtualizationVmwareVirtualDisk) GetUnitNumber() int64`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *VirtualizationVmwareVirtualDisk) GetUnitNumberOk() (*int64, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *VirtualizationVmwareVirtualDisk) SetUnitNumber(v int64)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *VirtualizationVmwareVirtualDisk) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualizationVmwareVirtualDisk) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualizationVmwareVirtualDisk) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualizationVmwareVirtualDisk) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualizationVmwareVirtualDisk) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVdiskId

`func (o *VirtualizationVmwareVirtualDisk) GetVdiskId() string`

GetVdiskId returns the VdiskId field if non-nil, zero value otherwise.

### GetVdiskIdOk

`func (o *VirtualizationVmwareVirtualDisk) GetVdiskIdOk() (*string, bool)`

GetVdiskIdOk returns a tuple with the VdiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiskId

`func (o *VirtualizationVmwareVirtualDisk) SetVdiskId(v string)`

SetVdiskId sets VdiskId field to given value.

### HasVdiskId

`func (o *VirtualizationVmwareVirtualDisk) HasVdiskId() bool`

HasVdiskId returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualizationVmwareVirtualDisk) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualizationVmwareVirtualDisk) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualizationVmwareVirtualDisk) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualizationVmwareVirtualDisk) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVirtualDiskPath

`func (o *VirtualizationVmwareVirtualDisk) GetVirtualDiskPath() string`

GetVirtualDiskPath returns the VirtualDiskPath field if non-nil, zero value otherwise.

### GetVirtualDiskPathOk

`func (o *VirtualizationVmwareVirtualDisk) GetVirtualDiskPathOk() (*string, bool)`

GetVirtualDiskPathOk returns a tuple with the VirtualDiskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskPath

`func (o *VirtualizationVmwareVirtualDisk) SetVirtualDiskPath(v string)`

SetVirtualDiskPath sets VirtualDiskPath field to given value.

### HasVirtualDiskPath

`func (o *VirtualizationVmwareVirtualDisk) HasVirtualDiskPath() bool`

HasVirtualDiskPath returns a boolean if a field has been set.

### GetVmIdentity

`func (o *VirtualizationVmwareVirtualDisk) GetVmIdentity() string`

GetVmIdentity returns the VmIdentity field if non-nil, zero value otherwise.

### GetVmIdentityOk

`func (o *VirtualizationVmwareVirtualDisk) GetVmIdentityOk() (*string, bool)`

GetVmIdentityOk returns a tuple with the VmIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIdentity

`func (o *VirtualizationVmwareVirtualDisk) SetVmIdentity(v string)`

SetVmIdentity sets VmIdentity field to given value.

### HasVmIdentity

`func (o *VirtualizationVmwareVirtualDisk) HasVmIdentity() bool`

HasVmIdentity returns a boolean if a field has been set.

### GetDatastore

`func (o *VirtualizationVmwareVirtualDisk) GetDatastore() VirtualizationVmwareDatastoreRelationship`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VirtualizationVmwareVirtualDisk) GetDatastoreOk() (*VirtualizationVmwareDatastoreRelationship, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VirtualizationVmwareVirtualDisk) SetDatastore(v VirtualizationVmwareDatastoreRelationship)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VirtualizationVmwareVirtualDisk) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *VirtualizationVmwareVirtualDisk) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualizationVmwareVirtualDisk) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualizationVmwareVirtualDisk) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VirtualizationVmwareVirtualDisk) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


