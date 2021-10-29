# VirtualizationVirtualDiskAllOf

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

### NewVirtualizationVirtualDiskAllOf

`func NewVirtualizationVirtualDiskAllOf(classId string, objectType string, ) *VirtualizationVirtualDiskAllOf`

NewVirtualizationVirtualDiskAllOf instantiates a new VirtualizationVirtualDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualDiskAllOfWithDefaults

`func NewVirtualizationVirtualDiskAllOfWithDefaults() *VirtualizationVirtualDiskAllOf`

NewVirtualizationVirtualDiskAllOfWithDefaults instantiates a new VirtualizationVirtualDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnitId

`func (o *VirtualizationVirtualDiskAllOf) GetBillingUnitId() string`

GetBillingUnitId returns the BillingUnitId field if non-nil, zero value otherwise.

### GetBillingUnitIdOk

`func (o *VirtualizationVirtualDiskAllOf) GetBillingUnitIdOk() (*string, bool)`

GetBillingUnitIdOk returns a tuple with the BillingUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnitId

`func (o *VirtualizationVirtualDiskAllOf) SetBillingUnitId(v string)`

SetBillingUnitId sets BillingUnitId field to given value.

### HasBillingUnitId

`func (o *VirtualizationVirtualDiskAllOf) HasBillingUnitId() bool`

HasBillingUnitId returns a boolean if a field has been set.

### GetCapacity

`func (o *VirtualizationVirtualDiskAllOf) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVirtualDiskAllOf) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVirtualDiskAllOf) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVirtualDiskAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDiscovered

`func (o *VirtualizationVirtualDiskAllOf) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *VirtualizationVirtualDiskAllOf) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *VirtualizationVirtualDiskAllOf) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *VirtualizationVirtualDiskAllOf) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDiskAction

`func (o *VirtualizationVirtualDiskAllOf) GetDiskAction() string`

GetDiskAction returns the DiskAction field if non-nil, zero value otherwise.

### GetDiskActionOk

`func (o *VirtualizationVirtualDiskAllOf) GetDiskActionOk() (*string, bool)`

GetDiskActionOk returns a tuple with the DiskAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAction

`func (o *VirtualizationVirtualDiskAllOf) SetDiskAction(v string)`

SetDiskAction sets DiskAction field to given value.

### HasDiskAction

`func (o *VirtualizationVirtualDiskAllOf) HasDiskAction() bool`

HasDiskAction returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *VirtualizationVirtualDiskAllOf) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *VirtualizationVirtualDiskAllOf) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *VirtualizationVirtualDiskAllOf) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *VirtualizationVirtualDiskAllOf) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEncryptionType

`func (o *VirtualizationVirtualDiskAllOf) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *VirtualizationVirtualDiskAllOf) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *VirtualizationVirtualDiskAllOf) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *VirtualizationVirtualDiskAllOf) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationVirtualDiskAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationVirtualDiskAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationVirtualDiskAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationVirtualDiskAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceCerts

`func (o *VirtualizationVirtualDiskAllOf) GetSourceCerts() string`

GetSourceCerts returns the SourceCerts field if non-nil, zero value otherwise.

### GetSourceCertsOk

`func (o *VirtualizationVirtualDiskAllOf) GetSourceCertsOk() (*string, bool)`

GetSourceCertsOk returns a tuple with the SourceCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCerts

`func (o *VirtualizationVirtualDiskAllOf) SetSourceCerts(v string)`

SetSourceCerts sets SourceCerts field to given value.

### HasSourceCerts

`func (o *VirtualizationVirtualDiskAllOf) HasSourceCerts() bool`

HasSourceCerts returns a boolean if a field has been set.

### GetSourceDiskToClone

`func (o *VirtualizationVirtualDiskAllOf) GetSourceDiskToClone() string`

GetSourceDiskToClone returns the SourceDiskToClone field if non-nil, zero value otherwise.

### GetSourceDiskToCloneOk

`func (o *VirtualizationVirtualDiskAllOf) GetSourceDiskToCloneOk() (*string, bool)`

GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDiskToClone

`func (o *VirtualizationVirtualDiskAllOf) SetSourceDiskToClone(v string)`

SetSourceDiskToClone sets SourceDiskToClone field to given value.

### HasSourceDiskToClone

`func (o *VirtualizationVirtualDiskAllOf) HasSourceDiskToClone() bool`

HasSourceDiskToClone returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *VirtualizationVirtualDiskAllOf) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *VirtualizationVirtualDiskAllOf) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *VirtualizationVirtualDiskAllOf) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *VirtualizationVirtualDiskAllOf) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetVolumeIopsInfo

`func (o *VirtualizationVirtualDiskAllOf) GetVolumeIopsInfo() CloudVolumeIopsInfo`

GetVolumeIopsInfo returns the VolumeIopsInfo field if non-nil, zero value otherwise.

### GetVolumeIopsInfoOk

`func (o *VirtualizationVirtualDiskAllOf) GetVolumeIopsInfoOk() (*CloudVolumeIopsInfo, bool)`

GetVolumeIopsInfoOk returns a tuple with the VolumeIopsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIopsInfo

`func (o *VirtualizationVirtualDiskAllOf) SetVolumeIopsInfo(v CloudVolumeIopsInfo)`

SetVolumeIopsInfo sets VolumeIopsInfo field to given value.

### HasVolumeIopsInfo

`func (o *VirtualizationVirtualDiskAllOf) HasVolumeIopsInfo() bool`

HasVolumeIopsInfo returns a boolean if a field has been set.

### SetVolumeIopsInfoNil

`func (o *VirtualizationVirtualDiskAllOf) SetVolumeIopsInfoNil(b bool)`

 SetVolumeIopsInfoNil sets the value for VolumeIopsInfo to be an explicit nil

### UnsetVolumeIopsInfo
`func (o *VirtualizationVirtualDiskAllOf) UnsetVolumeIopsInfo()`

UnsetVolumeIopsInfo ensures that no value is present for VolumeIopsInfo, not even an explicit nil
### GetZone

`func (o *VirtualizationVirtualDiskAllOf) GetZone() CloudAvailabilityZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *VirtualizationVirtualDiskAllOf) GetZoneOk() (*CloudAvailabilityZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *VirtualizationVirtualDiskAllOf) SetZone(v CloudAvailabilityZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *VirtualizationVirtualDiskAllOf) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *VirtualizationVirtualDiskAllOf) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *VirtualizationVirtualDiskAllOf) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetCluster

`func (o *VirtualizationVirtualDiskAllOf) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualDiskAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualDiskAllOf) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualDiskAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationVirtualDiskAllOf) GetInventory() VirtualizationBaseVirtualDiskRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationVirtualDiskAllOf) GetInventoryOk() (*VirtualizationBaseVirtualDiskRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationVirtualDiskAllOf) SetInventory(v VirtualizationBaseVirtualDiskRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationVirtualDiskAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVirtualDiskAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVirtualDiskAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVirtualDiskAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVirtualDiskAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationVirtualDiskAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationVirtualDiskAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationVirtualDiskAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationVirtualDiskAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


