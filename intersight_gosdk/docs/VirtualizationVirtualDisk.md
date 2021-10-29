# VirtualizationVirtualDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VirtualDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VirtualDisk"]
**BillingUnitId** | Pointer to **string** | Billing rate for this resource. | [optional] 
**Capacity** | Pointer to **string** | Disk capacity to be provided with units example - 10Gi. | [optional] 
**Discovered** | Pointer to **bool** | Flag to indicate whether the configuration is created from inventory object. | [optional] [readonly] 
**DiskAction** | Pointer to **string** | Action to perform on the disk example resize, shrink, defragment etc. | [optional] [readonly] 
**EncryptionKey** | Pointer to **string** | Encryption key used if volume is encrypted. | [optional] 
**EncryptionType** | Pointer to **string** | Encryption method used to encrypt the volume. | [optional] 
**Mode** | Pointer to **string** | File mode of the disk  example - Filesystem, Block. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [default to "Block"]
**Name** | Pointer to **string** | Name of the storage disk. Name must be unique within a Datastore. | [optional] 
**SourceCerts** | Pointer to **string** | Base64 encoded CA certificates of the https source to check against. | [optional] 
**SourceDiskToClone** | Pointer to **string** | Source disk from which the content is copied. | [optional] 
**SourceFilePath** | Pointer to **string** | Image path used to import on the created disk. | [optional] 
**VolumeIopsInfo** | Pointer to [**NullableCloudVolumeIopsInfo**](CloudVolumeIopsInfo.md) |  | [optional] 
**Zone** | Pointer to [**NullableCloudAvailabilityZone**](CloudAvailabilityZone.md) |  | [optional] 
**Cluster** | Pointer to [**VirtualizationBaseClusterRelationship**](VirtualizationBaseClusterRelationship.md) |  | [optional] 
**Inventory** | Pointer to [**VirtualizationBaseVirtualDiskRelationship**](VirtualizationBaseVirtualDiskRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVirtualDisk

`func NewVirtualizationVirtualDisk(classId string, objectType string, ) *VirtualizationVirtualDisk`

NewVirtualizationVirtualDisk instantiates a new VirtualizationVirtualDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualDiskWithDefaults

`func NewVirtualizationVirtualDiskWithDefaults() *VirtualizationVirtualDisk`

NewVirtualizationVirtualDiskWithDefaults instantiates a new VirtualizationVirtualDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnitId

`func (o *VirtualizationVirtualDisk) GetBillingUnitId() string`

GetBillingUnitId returns the BillingUnitId field if non-nil, zero value otherwise.

### GetBillingUnitIdOk

`func (o *VirtualizationVirtualDisk) GetBillingUnitIdOk() (*string, bool)`

GetBillingUnitIdOk returns a tuple with the BillingUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnitId

`func (o *VirtualizationVirtualDisk) SetBillingUnitId(v string)`

SetBillingUnitId sets BillingUnitId field to given value.

### HasBillingUnitId

`func (o *VirtualizationVirtualDisk) HasBillingUnitId() bool`

HasBillingUnitId returns a boolean if a field has been set.

### GetCapacity

`func (o *VirtualizationVirtualDisk) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVirtualDisk) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVirtualDisk) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVirtualDisk) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDiscovered

`func (o *VirtualizationVirtualDisk) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *VirtualizationVirtualDisk) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *VirtualizationVirtualDisk) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *VirtualizationVirtualDisk) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDiskAction

`func (o *VirtualizationVirtualDisk) GetDiskAction() string`

GetDiskAction returns the DiskAction field if non-nil, zero value otherwise.

### GetDiskActionOk

`func (o *VirtualizationVirtualDisk) GetDiskActionOk() (*string, bool)`

GetDiskActionOk returns a tuple with the DiskAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAction

`func (o *VirtualizationVirtualDisk) SetDiskAction(v string)`

SetDiskAction sets DiskAction field to given value.

### HasDiskAction

`func (o *VirtualizationVirtualDisk) HasDiskAction() bool`

HasDiskAction returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *VirtualizationVirtualDisk) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *VirtualizationVirtualDisk) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *VirtualizationVirtualDisk) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *VirtualizationVirtualDisk) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEncryptionType

`func (o *VirtualizationVirtualDisk) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *VirtualizationVirtualDisk) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *VirtualizationVirtualDisk) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *VirtualizationVirtualDisk) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationVirtualDisk) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationVirtualDisk) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationVirtualDisk) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationVirtualDisk) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceCerts

`func (o *VirtualizationVirtualDisk) GetSourceCerts() string`

GetSourceCerts returns the SourceCerts field if non-nil, zero value otherwise.

### GetSourceCertsOk

`func (o *VirtualizationVirtualDisk) GetSourceCertsOk() (*string, bool)`

GetSourceCertsOk returns a tuple with the SourceCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCerts

`func (o *VirtualizationVirtualDisk) SetSourceCerts(v string)`

SetSourceCerts sets SourceCerts field to given value.

### HasSourceCerts

`func (o *VirtualizationVirtualDisk) HasSourceCerts() bool`

HasSourceCerts returns a boolean if a field has been set.

### GetSourceDiskToClone

`func (o *VirtualizationVirtualDisk) GetSourceDiskToClone() string`

GetSourceDiskToClone returns the SourceDiskToClone field if non-nil, zero value otherwise.

### GetSourceDiskToCloneOk

`func (o *VirtualizationVirtualDisk) GetSourceDiskToCloneOk() (*string, bool)`

GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDiskToClone

`func (o *VirtualizationVirtualDisk) SetSourceDiskToClone(v string)`

SetSourceDiskToClone sets SourceDiskToClone field to given value.

### HasSourceDiskToClone

`func (o *VirtualizationVirtualDisk) HasSourceDiskToClone() bool`

HasSourceDiskToClone returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *VirtualizationVirtualDisk) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *VirtualizationVirtualDisk) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *VirtualizationVirtualDisk) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *VirtualizationVirtualDisk) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetVolumeIopsInfo

`func (o *VirtualizationVirtualDisk) GetVolumeIopsInfo() CloudVolumeIopsInfo`

GetVolumeIopsInfo returns the VolumeIopsInfo field if non-nil, zero value otherwise.

### GetVolumeIopsInfoOk

`func (o *VirtualizationVirtualDisk) GetVolumeIopsInfoOk() (*CloudVolumeIopsInfo, bool)`

GetVolumeIopsInfoOk returns a tuple with the VolumeIopsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIopsInfo

`func (o *VirtualizationVirtualDisk) SetVolumeIopsInfo(v CloudVolumeIopsInfo)`

SetVolumeIopsInfo sets VolumeIopsInfo field to given value.

### HasVolumeIopsInfo

`func (o *VirtualizationVirtualDisk) HasVolumeIopsInfo() bool`

HasVolumeIopsInfo returns a boolean if a field has been set.

### SetVolumeIopsInfoNil

`func (o *VirtualizationVirtualDisk) SetVolumeIopsInfoNil(b bool)`

 SetVolumeIopsInfoNil sets the value for VolumeIopsInfo to be an explicit nil

### UnsetVolumeIopsInfo
`func (o *VirtualizationVirtualDisk) UnsetVolumeIopsInfo()`

UnsetVolumeIopsInfo ensures that no value is present for VolumeIopsInfo, not even an explicit nil
### GetZone

`func (o *VirtualizationVirtualDisk) GetZone() CloudAvailabilityZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *VirtualizationVirtualDisk) GetZoneOk() (*CloudAvailabilityZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *VirtualizationVirtualDisk) SetZone(v CloudAvailabilityZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *VirtualizationVirtualDisk) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *VirtualizationVirtualDisk) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *VirtualizationVirtualDisk) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetCluster

`func (o *VirtualizationVirtualDisk) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualDisk) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualDisk) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualDisk) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationVirtualDisk) GetInventory() VirtualizationBaseVirtualDiskRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationVirtualDisk) GetInventoryOk() (*VirtualizationBaseVirtualDiskRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationVirtualDisk) SetInventory(v VirtualizationBaseVirtualDiskRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationVirtualDisk) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVirtualDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVirtualDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVirtualDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVirtualDisk) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationVirtualDisk) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationVirtualDisk) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationVirtualDisk) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationVirtualDisk) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


