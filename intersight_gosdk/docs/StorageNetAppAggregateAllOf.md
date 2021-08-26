# StorageNetAppAggregateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppAggregate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppAggregate"]
**AggregateType** | Pointer to **string** | Storage disk type for NetApp aggregate. * &#x60;HDD&#x60; - Hard Disk Drive. * &#x60;Hybrid&#x60; - Solid State Hard Disk Drive. * &#x60;Hybrid (Flash Pool)&#x60; - SSHD in a flash pool. * &#x60;SSD&#x60; - Solid State Disk. * &#x60;SSD (FabricPool)&#x60; - SSD in a flash pool. * &#x60;VMDisk (SDS)&#x60; - Storage disk with Hard disk drive. * &#x60;VMDisk (FabricPool)&#x60; - Storage disk with Non-volatile random-access memory drives. * &#x60;LUN (FlexArray)&#x60; - LUN as a disk. * &#x60;Not Mapped&#x60; - Storage disk is not mapped. | [optional] [readonly] [default to "HDD"]
**RaidSize** | Pointer to **int64** | Size of the RAID group represented by number of disks in the group. | [optional] [readonly] 
**RaidType** | Pointer to **string** | The RAID configuration type. * &#x60;Unknown&#x60; - Default unknown RAID type. * &#x60;RAID0&#x60; - RAID0 splits (\&quot;stripes\&quot;) data evenly across two or more disks, without parity information. * &#x60;RAID1&#x60; - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * &#x60;RAID4&#x60; - RAID4 stripes block level data and dedicates a disk to parity. * &#x60;RAID5&#x60; - RAID5  distributes striping and parity at a block level. * &#x60;RAID6&#x60; - RAID6 level operates like RAID5 with distributed parity and striping. * &#x60;RAID10&#x60; - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * &#x60;RAIDDP&#x60; - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * &#x60;RAIDTEC&#x60; - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group. | [optional] [readonly] [default to "Unknown"]
**State** | Pointer to **string** | Current state of the NetApp aggregate. * &#x60;Unknown&#x60; - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server. * &#x60;Online&#x60; - Aggregate is ready and available. * &#x60;Onlining&#x60; - Transitioning online. * &#x60;Offline&#x60; - Aggregate is unavailable. * &#x60;Offlining&#x60; - Transitioning offline. * &#x60;Relocating&#x60; - The aggregate is being relocated. * &#x60;Restricted&#x60; - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed. * &#x60;Failed&#x60; - The aggregate cannot be brought online. * &#x60;Inconsistent&#x60; - The aggregate has been marked corrupted; contact technical support. * &#x60;Unmounted&#x60; - The aggregate is not mounted. | [optional] [readonly] [default to "Unknown"]
**Uuid** | Pointer to **string** | Uuid of  NetApp Aggregate. | [optional] [readonly] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppAggregateAllOf

`func NewStorageNetAppAggregateAllOf(classId string, objectType string, ) *StorageNetAppAggregateAllOf`

NewStorageNetAppAggregateAllOf instantiates a new StorageNetAppAggregateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppAggregateAllOfWithDefaults

`func NewStorageNetAppAggregateAllOfWithDefaults() *StorageNetAppAggregateAllOf`

NewStorageNetAppAggregateAllOfWithDefaults instantiates a new StorageNetAppAggregateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppAggregateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppAggregateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppAggregateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppAggregateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppAggregateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppAggregateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregateType

`func (o *StorageNetAppAggregateAllOf) GetAggregateType() string`

GetAggregateType returns the AggregateType field if non-nil, zero value otherwise.

### GetAggregateTypeOk

`func (o *StorageNetAppAggregateAllOf) GetAggregateTypeOk() (*string, bool)`

GetAggregateTypeOk returns a tuple with the AggregateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateType

`func (o *StorageNetAppAggregateAllOf) SetAggregateType(v string)`

SetAggregateType sets AggregateType field to given value.

### HasAggregateType

`func (o *StorageNetAppAggregateAllOf) HasAggregateType() bool`

HasAggregateType returns a boolean if a field has been set.

### GetRaidSize

`func (o *StorageNetAppAggregateAllOf) GetRaidSize() int64`

GetRaidSize returns the RaidSize field if non-nil, zero value otherwise.

### GetRaidSizeOk

`func (o *StorageNetAppAggregateAllOf) GetRaidSizeOk() (*int64, bool)`

GetRaidSizeOk returns a tuple with the RaidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidSize

`func (o *StorageNetAppAggregateAllOf) SetRaidSize(v int64)`

SetRaidSize sets RaidSize field to given value.

### HasRaidSize

`func (o *StorageNetAppAggregateAllOf) HasRaidSize() bool`

HasRaidSize returns a boolean if a field has been set.

### GetRaidType

`func (o *StorageNetAppAggregateAllOf) GetRaidType() string`

GetRaidType returns the RaidType field if non-nil, zero value otherwise.

### GetRaidTypeOk

`func (o *StorageNetAppAggregateAllOf) GetRaidTypeOk() (*string, bool)`

GetRaidTypeOk returns a tuple with the RaidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidType

`func (o *StorageNetAppAggregateAllOf) SetRaidType(v string)`

SetRaidType sets RaidType field to given value.

### HasRaidType

`func (o *StorageNetAppAggregateAllOf) HasRaidType() bool`

HasRaidType returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppAggregateAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppAggregateAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppAggregateAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppAggregateAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppAggregateAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppAggregateAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppAggregateAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppAggregateAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppAggregateAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppAggregateAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppAggregateAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppAggregateAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


