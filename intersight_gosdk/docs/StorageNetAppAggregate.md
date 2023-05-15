# StorageNetAppAggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppAggregate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppAggregate"]
**AggregateType** | Pointer to **string** | Storage disk type for NetApp aggregate. * &#x60;HDD&#x60; - Hard Disk Drive disk type. * &#x60;Hybrid&#x60; - Solid State Hard Disk Drive. * &#x60;Hybrid (Flash Pool)&#x60; - SSHD in a flash pool disk type. * &#x60;SSD&#x60; - Solid State Disk disk type. * &#x60;SSD (FabricPool)&#x60; - SSD in a flash pool disk type. * &#x60;VMDisk (SDS)&#x60; - Storage disk with Hard disk drive. * &#x60;VMDisk (FabricPool)&#x60; - Storage disk with Non-volatile random-access memory drives. * &#x60;LUN (FlexArray)&#x60; - LUN (FlexArray) disk type. * &#x60;Not Mapped&#x60; - Storage disk is not mapped. | [optional] [readonly] [default to "HDD"]
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**CloudStorage** | Pointer to **[]string** |  | [optional] 
**EfficiencyRatio** | Pointer to **float32** | Data reduction ratio (logical_used / used). | [optional] [readonly] 
**Key** | Pointer to **string** | Unique identifier of NetApp Aggregate across data center. | [optional] [readonly] 
**NodeName** | Pointer to **string** | The node name for the aggregate. | [optional] [readonly] 
**RaidSize** | Pointer to **int64** | Size of the RAID group represented by number of disks in the group. | [optional] [readonly] 
**RaidType** | Pointer to **string** | The RAID configuration type. * &#x60;Unknown&#x60; - Default unknown RAID type. * &#x60;RAID0&#x60; - RAID0 splits (\&quot;stripes\&quot;) data evenly across two or more disks, without parity information. * &#x60;RAID1&#x60; - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * &#x60;RAID4&#x60; - RAID4 stripes block level data and dedicates a disk to parity. * &#x60;RAID5&#x60; - RAID5  distributes striping and parity at a block level. * &#x60;RAID6&#x60; - RAID6 level operates like RAID5 with distributed parity and striping. * &#x60;RAID10&#x60; - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * &#x60;RAIDDP&#x60; - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * &#x60;RAIDTEC&#x60; - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group. | [optional] [readonly] [default to "Unknown"]
**State** | Pointer to **string** | Current state of the NetApp aggregate. * &#x60;Unknown&#x60; - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server. * &#x60;Online&#x60; - Aggregate is ready and available. * &#x60;Onlining&#x60; - The state is transitioning online. * &#x60;Offline&#x60; - Aggregate is unavailable. * &#x60;Offlining&#x60; - The state is transitioning offline. * &#x60;Relocating&#x60; - The aggregate is being relocated. * &#x60;Restricted&#x60; - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed. * &#x60;Failed&#x60; - The aggregate cannot be brought online. * &#x60;Inconsistent&#x60; - The aggregate has been marked corrupted; contact technical support. * &#x60;Unmounted&#x60; - The aggregate is not mounted. | [optional] [readonly] [default to "Unknown"]
**Uuid** | Pointer to **string** | Uuid of NetApp Aggregate. | [optional] [readonly] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppAggregateEventRelationship**](StorageNetAppAggregateEventRelationship.md) | An array of relationships to storageNetAppAggregateEvent resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppAggregate

`func NewStorageNetAppAggregate(classId string, objectType string, ) *StorageNetAppAggregate`

NewStorageNetAppAggregate instantiates a new StorageNetAppAggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppAggregateWithDefaults

`func NewStorageNetAppAggregateWithDefaults() *StorageNetAppAggregate`

NewStorageNetAppAggregateWithDefaults instantiates a new StorageNetAppAggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppAggregate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppAggregate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppAggregate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppAggregate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppAggregate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppAggregate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregateType

`func (o *StorageNetAppAggregate) GetAggregateType() string`

GetAggregateType returns the AggregateType field if non-nil, zero value otherwise.

### GetAggregateTypeOk

`func (o *StorageNetAppAggregate) GetAggregateTypeOk() (*string, bool)`

GetAggregateTypeOk returns a tuple with the AggregateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateType

`func (o *StorageNetAppAggregate) SetAggregateType(v string)`

SetAggregateType sets AggregateType field to given value.

### HasAggregateType

`func (o *StorageNetAppAggregate) HasAggregateType() bool`

HasAggregateType returns a boolean if a field has been set.

### GetAvgPerformanceMetrics

`func (o *StorageNetAppAggregate) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppAggregate) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppAggregate) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppAggregate) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetCloudStorage

`func (o *StorageNetAppAggregate) GetCloudStorage() []string`

GetCloudStorage returns the CloudStorage field if non-nil, zero value otherwise.

### GetCloudStorageOk

`func (o *StorageNetAppAggregate) GetCloudStorageOk() (*[]string, bool)`

GetCloudStorageOk returns a tuple with the CloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorage

`func (o *StorageNetAppAggregate) SetCloudStorage(v []string)`

SetCloudStorage sets CloudStorage field to given value.

### HasCloudStorage

`func (o *StorageNetAppAggregate) HasCloudStorage() bool`

HasCloudStorage returns a boolean if a field has been set.

### SetCloudStorageNil

`func (o *StorageNetAppAggregate) SetCloudStorageNil(b bool)`

 SetCloudStorageNil sets the value for CloudStorage to be an explicit nil

### UnsetCloudStorage
`func (o *StorageNetAppAggregate) UnsetCloudStorage()`

UnsetCloudStorage ensures that no value is present for CloudStorage, not even an explicit nil
### GetEfficiencyRatio

`func (o *StorageNetAppAggregate) GetEfficiencyRatio() float32`

GetEfficiencyRatio returns the EfficiencyRatio field if non-nil, zero value otherwise.

### GetEfficiencyRatioOk

`func (o *StorageNetAppAggregate) GetEfficiencyRatioOk() (*float32, bool)`

GetEfficiencyRatioOk returns a tuple with the EfficiencyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfficiencyRatio

`func (o *StorageNetAppAggregate) SetEfficiencyRatio(v float32)`

SetEfficiencyRatio sets EfficiencyRatio field to given value.

### HasEfficiencyRatio

`func (o *StorageNetAppAggregate) HasEfficiencyRatio() bool`

HasEfficiencyRatio returns a boolean if a field has been set.

### GetKey

`func (o *StorageNetAppAggregate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppAggregate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppAggregate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppAggregate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNodeName

`func (o *StorageNetAppAggregate) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *StorageNetAppAggregate) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *StorageNetAppAggregate) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *StorageNetAppAggregate) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetRaidSize

`func (o *StorageNetAppAggregate) GetRaidSize() int64`

GetRaidSize returns the RaidSize field if non-nil, zero value otherwise.

### GetRaidSizeOk

`func (o *StorageNetAppAggregate) GetRaidSizeOk() (*int64, bool)`

GetRaidSizeOk returns a tuple with the RaidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidSize

`func (o *StorageNetAppAggregate) SetRaidSize(v int64)`

SetRaidSize sets RaidSize field to given value.

### HasRaidSize

`func (o *StorageNetAppAggregate) HasRaidSize() bool`

HasRaidSize returns a boolean if a field has been set.

### GetRaidType

`func (o *StorageNetAppAggregate) GetRaidType() string`

GetRaidType returns the RaidType field if non-nil, zero value otherwise.

### GetRaidTypeOk

`func (o *StorageNetAppAggregate) GetRaidTypeOk() (*string, bool)`

GetRaidTypeOk returns a tuple with the RaidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidType

`func (o *StorageNetAppAggregate) SetRaidType(v string)`

SetRaidType sets RaidType field to given value.

### HasRaidType

`func (o *StorageNetAppAggregate) HasRaidType() bool`

HasRaidType returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppAggregate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppAggregate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppAggregate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppAggregate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppAggregate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppAggregate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppAggregate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppAggregate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppAggregate) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppAggregate) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppAggregate) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppAggregate) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppAggregate) GetEvents() []StorageNetAppAggregateEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppAggregate) GetEventsOk() (*[]StorageNetAppAggregateEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppAggregate) SetEvents(v []StorageNetAppAggregateEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppAggregate) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppAggregate) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppAggregate) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


