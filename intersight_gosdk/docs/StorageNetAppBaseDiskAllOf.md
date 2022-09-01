# StorageNetAppBaseDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppBaseDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppBaseDisk"]
**BaseDiskModel** | Pointer to **string** | The NetApp base disk model. | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | Unique identity of the device. | [optional] [readonly] 
**ContainerType** | Pointer to **string** | Supported container type for NetApp disk. * &#x60;Unknown&#x60; - Default container type is currently unknown. * &#x60;Aggregate&#x60; - Disk is used as a physical disk in an aggregate. * &#x60;Broken&#x60; - Disk is in a broken pool. * &#x60;Label Maintenance&#x60; - Disk is in online label maintenance list. * &#x60;Foreign&#x60; - Array LUN has been marked foreign. * &#x60;Maintenance&#x60; - Disk is in maintenance center. * &#x60;Mediator&#x60; - A mediator disk is a disk used on non-shared HA systems hosted by an external node which is used to communicate the viability of the storage failover between non-shared HA nodes. * &#x60;Shared&#x60; - Disk is partitioned or in a storage pool. * &#x60;Remote&#x60; - Disk belongs to a remote cluster. * &#x60;Spare&#x60; - The disk is a spare disk. * &#x60;Unassigned&#x60; - Disk ownership has not been assigned. * &#x60;Unsupported&#x60; - The disk is not supported. | [optional] [readonly] [default to "Unknown"]
**DiskBay** | Pointer to **int64** | NetApp base disk shelf bay. | [optional] [readonly] 
**DiskSerialNumber** | Pointer to **string** | NetApp base disk serial number. | [optional] [readonly] 
**DiskShelfId** | Pointer to **string** | NetApp base disk shelf id. | [optional] [readonly] 
**DiskShelfModel** | Pointer to **string** | NetApp base disk shelf model. | [optional] [readonly] 
**DiskShelfName** | Pointer to **string** | NetApp base disk shelf name. | [optional] [readonly] 
**DiskType** | Pointer to **string** | The type of the NetApp disk. * &#x60;Unknown&#x60; - Default unknown disk type. * &#x60;SSDNVM&#x60; - Solid state disk with Non-Volatile Memory Express protocol enabled. * &#x60;ATA&#x60; - Advanced Technology Attachment is a type of disk drive that integrates the drive controller directly on the drive itself. * &#x60;FCAL&#x60; - For the FC-AL disk connection type, disk shelves are connected to the controller in a loop. * &#x60;BSAS&#x60; - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * &#x60;FSAS&#x60; - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, and rotational speed of traditional enterprise-class SATA drives with the fully capable SAS interface typical for classic SAS drives. * &#x60;LUN&#x60; - Logical Unit Number refers to a logical disk. * &#x60;SAS&#x60; - Storage disk with serial attached SCSI. * &#x60;MSATA&#x60; - SATA disk in multi-disk carrier storage shelf. * &#x60;SSD&#x60; - Storage disk with Solid state disk. * &#x60;VMDISK&#x60; - Virtual machine Data Disk. | [optional] [readonly] [default to "Unknown"]
**State** | Pointer to **string** | Current state of the NetApp disk. * &#x60;Present&#x60; - Storage disk state type is present. * &#x60;Copy&#x60; - Storage disk state type is copy. * &#x60;Broken&#x60; - Storage disk state type is broken. * &#x60;Maintenance&#x60; - Storage disk state type is maintenance. * &#x60;Partner&#x60; - Storage disk state type is partner. * &#x60;Pending&#x60; - Storage disk state type is pending. * &#x60;Reconstructing&#x60; - Storage disk state type is reconstructing. * &#x60;Removed&#x60; - Storage disk state type is removed. * &#x60;Spare&#x60; - Storage disk state type is spare. * &#x60;Unfail&#x60; - Storage disk state type is unfail. * &#x60;Zeroing&#x60; - Storage disk state type is zeroing. | [optional] [readonly] [default to "Present"]
**Uuid** | Pointer to **string** | Universally unique identifier of the NetApp Disk. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**DiskPool** | Pointer to [**[]StorageNetAppAggregateRelationship**](StorageNetAppAggregateRelationship.md) | An array of relationships to storageNetAppAggregate resources. | [optional] [readonly] 
**Events** | Pointer to [**[]StorageNetAppDiskEventRelationship**](StorageNetAppDiskEventRelationship.md) | An array of relationships to storageNetAppDiskEvent resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppBaseDiskAllOf

`func NewStorageNetAppBaseDiskAllOf(classId string, objectType string, ) *StorageNetAppBaseDiskAllOf`

NewStorageNetAppBaseDiskAllOf instantiates a new StorageNetAppBaseDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppBaseDiskAllOfWithDefaults

`func NewStorageNetAppBaseDiskAllOfWithDefaults() *StorageNetAppBaseDiskAllOf`

NewStorageNetAppBaseDiskAllOfWithDefaults instantiates a new StorageNetAppBaseDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppBaseDiskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppBaseDiskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppBaseDiskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppBaseDiskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppBaseDiskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppBaseDiskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBaseDiskModel

`func (o *StorageNetAppBaseDiskAllOf) GetBaseDiskModel() string`

GetBaseDiskModel returns the BaseDiskModel field if non-nil, zero value otherwise.

### GetBaseDiskModelOk

`func (o *StorageNetAppBaseDiskAllOf) GetBaseDiskModelOk() (*string, bool)`

GetBaseDiskModelOk returns a tuple with the BaseDiskModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDiskModel

`func (o *StorageNetAppBaseDiskAllOf) SetBaseDiskModel(v string)`

SetBaseDiskModel sets BaseDiskModel field to given value.

### HasBaseDiskModel

`func (o *StorageNetAppBaseDiskAllOf) HasBaseDiskModel() bool`

HasBaseDiskModel returns a boolean if a field has been set.

### GetClusterUuid

`func (o *StorageNetAppBaseDiskAllOf) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *StorageNetAppBaseDiskAllOf) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *StorageNetAppBaseDiskAllOf) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *StorageNetAppBaseDiskAllOf) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetContainerType

`func (o *StorageNetAppBaseDiskAllOf) GetContainerType() string`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *StorageNetAppBaseDiskAllOf) GetContainerTypeOk() (*string, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *StorageNetAppBaseDiskAllOf) SetContainerType(v string)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *StorageNetAppBaseDiskAllOf) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetDiskBay

`func (o *StorageNetAppBaseDiskAllOf) GetDiskBay() int64`

GetDiskBay returns the DiskBay field if non-nil, zero value otherwise.

### GetDiskBayOk

`func (o *StorageNetAppBaseDiskAllOf) GetDiskBayOk() (*int64, bool)`

GetDiskBayOk returns a tuple with the DiskBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskBay

`func (o *StorageNetAppBaseDiskAllOf) SetDiskBay(v int64)`

SetDiskBay sets DiskBay field to given value.

### HasDiskBay

`func (o *StorageNetAppBaseDiskAllOf) HasDiskBay() bool`

HasDiskBay returns a boolean if a field has been set.

### GetDiskSerialNumber

`func (o *StorageNetAppBaseDiskAllOf) GetDiskSerialNumber() string`

GetDiskSerialNumber returns the DiskSerialNumber field if non-nil, zero value otherwise.

### GetDiskSerialNumberOk

`func (o *StorageNetAppBaseDiskAllOf) GetDiskSerialNumberOk() (*string, bool)`

GetDiskSerialNumberOk returns a tuple with the DiskSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSerialNumber

`func (o *StorageNetAppBaseDiskAllOf) SetDiskSerialNumber(v string)`

SetDiskSerialNumber sets DiskSerialNumber field to given value.

### HasDiskSerialNumber

`func (o *StorageNetAppBaseDiskAllOf) HasDiskSerialNumber() bool`

HasDiskSerialNumber returns a boolean if a field has been set.

### GetDiskShelfId

`func (o *StorageNetAppBaseDiskAllOf) GetDiskShelfId() string`

GetDiskShelfId returns the DiskShelfId field if non-nil, zero value otherwise.

### GetDiskShelfIdOk

`func (o *StorageNetAppBaseDiskAllOf) GetDiskShelfIdOk() (*string, bool)`

GetDiskShelfIdOk returns a tuple with the DiskShelfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskShelfId

`func (o *StorageNetAppBaseDiskAllOf) SetDiskShelfId(v string)`

SetDiskShelfId sets DiskShelfId field to given value.

### HasDiskShelfId

`func (o *StorageNetAppBaseDiskAllOf) HasDiskShelfId() bool`

HasDiskShelfId returns a boolean if a field has been set.

### GetDiskShelfModel

`func (o *StorageNetAppBaseDiskAllOf) GetDiskShelfModel() string`

GetDiskShelfModel returns the DiskShelfModel field if non-nil, zero value otherwise.

### GetDiskShelfModelOk

`func (o *StorageNetAppBaseDiskAllOf) GetDiskShelfModelOk() (*string, bool)`

GetDiskShelfModelOk returns a tuple with the DiskShelfModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskShelfModel

`func (o *StorageNetAppBaseDiskAllOf) SetDiskShelfModel(v string)`

SetDiskShelfModel sets DiskShelfModel field to given value.

### HasDiskShelfModel

`func (o *StorageNetAppBaseDiskAllOf) HasDiskShelfModel() bool`

HasDiskShelfModel returns a boolean if a field has been set.

### GetDiskShelfName

`func (o *StorageNetAppBaseDiskAllOf) GetDiskShelfName() string`

GetDiskShelfName returns the DiskShelfName field if non-nil, zero value otherwise.

### GetDiskShelfNameOk

`func (o *StorageNetAppBaseDiskAllOf) GetDiskShelfNameOk() (*string, bool)`

GetDiskShelfNameOk returns a tuple with the DiskShelfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskShelfName

`func (o *StorageNetAppBaseDiskAllOf) SetDiskShelfName(v string)`

SetDiskShelfName sets DiskShelfName field to given value.

### HasDiskShelfName

`func (o *StorageNetAppBaseDiskAllOf) HasDiskShelfName() bool`

HasDiskShelfName returns a boolean if a field has been set.

### GetDiskType

`func (o *StorageNetAppBaseDiskAllOf) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *StorageNetAppBaseDiskAllOf) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *StorageNetAppBaseDiskAllOf) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *StorageNetAppBaseDiskAllOf) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppBaseDiskAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppBaseDiskAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppBaseDiskAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppBaseDiskAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppBaseDiskAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppBaseDiskAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppBaseDiskAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppBaseDiskAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppBaseDiskAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppBaseDiskAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppBaseDiskAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppBaseDiskAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppBaseDiskAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppBaseDiskAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppBaseDiskAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppBaseDiskAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetDiskPool

`func (o *StorageNetAppBaseDiskAllOf) GetDiskPool() []StorageNetAppAggregateRelationship`

GetDiskPool returns the DiskPool field if non-nil, zero value otherwise.

### GetDiskPoolOk

`func (o *StorageNetAppBaseDiskAllOf) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool)`

GetDiskPoolOk returns a tuple with the DiskPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPool

`func (o *StorageNetAppBaseDiskAllOf) SetDiskPool(v []StorageNetAppAggregateRelationship)`

SetDiskPool sets DiskPool field to given value.

### HasDiskPool

`func (o *StorageNetAppBaseDiskAllOf) HasDiskPool() bool`

HasDiskPool returns a boolean if a field has been set.

### SetDiskPoolNil

`func (o *StorageNetAppBaseDiskAllOf) SetDiskPoolNil(b bool)`

 SetDiskPoolNil sets the value for DiskPool to be an explicit nil

### UnsetDiskPool
`func (o *StorageNetAppBaseDiskAllOf) UnsetDiskPool()`

UnsetDiskPool ensures that no value is present for DiskPool, not even an explicit nil
### GetEvents

`func (o *StorageNetAppBaseDiskAllOf) GetEvents() []StorageNetAppDiskEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppBaseDiskAllOf) GetEventsOk() (*[]StorageNetAppDiskEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppBaseDiskAllOf) SetEvents(v []StorageNetAppDiskEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppBaseDiskAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppBaseDiskAllOf) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppBaseDiskAllOf) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


