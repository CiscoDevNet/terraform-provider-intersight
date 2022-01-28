# StorageNetAppLunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppLun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppLun"]
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**Key** | Pointer to **string** | Unique identifier of Lun across data center. | [optional] [readonly] 
**Mapped** | Pointer to **bool** | Reports if the LUN is mapped to one or more initiator groups. | [optional] [readonly] 
**OsType** | Pointer to **string** | The operating system (OS) type for this LUN. * &#x60;Linux&#x60; - Family of open source Unix-like operating systems based on the Linux kernel. * &#x60;AIX&#x60; - Advanced Interactive Executive (AIX). * &#x60;HP-UX&#x60; - HP-UX is implementation of the Unix operating system, based on Unix System V. * &#x60;Hyper-V&#x60; - Windows Server 2008 or Windows Server 2012 Hyper-V. * &#x60;OpenVMS&#x60; - OpenVMS is multi-user, multiprocessing virtual memory-based operating system. * &#x60;Solaris&#x60; - Solaris is a Unix operating system. * &#x60;NetWare&#x60; - NetWare is a computer network operating system. * &#x60;VMware&#x60; - An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers. * &#x60;Windows&#x60; - Single-partition Windows disk using the Master Boot Record (MBR) partitioning style. * &#x60;Xen&#x60; - Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently. | [optional] [readonly] [default to "Linux"]
**Path** | Pointer to **string** | Path where the LUN is mounted. | [optional] [readonly] 
**Serial** | Pointer to **string** | Serial number for the provisioned LUN. | [optional] [readonly] 
**State** | Pointer to **string** | The administrative state of a LUN. * &#x60;offline&#x60; - The LUN is administratively offline, or a more detailed offline reason is not available. * &#x60;online&#x60; - The state of the LUN is online. | [optional] [readonly] [default to "offline"]
**Uuid** | Pointer to **string** | Universally unique identifier of the LUN. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppLunEventRelationship**](StorageNetAppLunEventRelationship.md) | An array of relationships to storageNetAppLunEvent resources. | [optional] [readonly] 
**Host** | Pointer to [**[]StorageNetAppInitiatorGroupRelationship**](StorageNetAppInitiatorGroupRelationship.md) | An array of relationships to storageNetAppInitiatorGroup resources. | [optional] [readonly] 
**StorageContainer** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppLunAllOf

`func NewStorageNetAppLunAllOf(classId string, objectType string, ) *StorageNetAppLunAllOf`

NewStorageNetAppLunAllOf instantiates a new StorageNetAppLunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppLunAllOfWithDefaults

`func NewStorageNetAppLunAllOfWithDefaults() *StorageNetAppLunAllOf`

NewStorageNetAppLunAllOfWithDefaults instantiates a new StorageNetAppLunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppLunAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppLunAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppLunAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppLunAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppLunAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppLunAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvgPerformanceMetrics

`func (o *StorageNetAppLunAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppLunAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppLunAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppLunAllOf) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetKey

`func (o *StorageNetAppLunAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppLunAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppLunAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppLunAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMapped

`func (o *StorageNetAppLunAllOf) GetMapped() bool`

GetMapped returns the Mapped field if non-nil, zero value otherwise.

### GetMappedOk

`func (o *StorageNetAppLunAllOf) GetMappedOk() (*bool, bool)`

GetMappedOk returns a tuple with the Mapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapped

`func (o *StorageNetAppLunAllOf) SetMapped(v bool)`

SetMapped sets Mapped field to given value.

### HasMapped

`func (o *StorageNetAppLunAllOf) HasMapped() bool`

HasMapped returns a boolean if a field has been set.

### GetOsType

`func (o *StorageNetAppLunAllOf) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *StorageNetAppLunAllOf) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *StorageNetAppLunAllOf) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *StorageNetAppLunAllOf) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetPath

`func (o *StorageNetAppLunAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageNetAppLunAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageNetAppLunAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageNetAppLunAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSerial

`func (o *StorageNetAppLunAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageNetAppLunAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageNetAppLunAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageNetAppLunAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppLunAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppLunAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppLunAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppLunAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppLunAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppLunAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppLunAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppLunAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppLunAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppLunAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppLunAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppLunAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppLunAllOf) GetEvents() []StorageNetAppLunEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppLunAllOf) GetEventsOk() (*[]StorageNetAppLunEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppLunAllOf) SetEvents(v []StorageNetAppLunEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppLunAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppLunAllOf) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppLunAllOf) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetHost

`func (o *StorageNetAppLunAllOf) GetHost() []StorageNetAppInitiatorGroupRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageNetAppLunAllOf) GetHostOk() (*[]StorageNetAppInitiatorGroupRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageNetAppLunAllOf) SetHost(v []StorageNetAppInitiatorGroupRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageNetAppLunAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *StorageNetAppLunAllOf) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *StorageNetAppLunAllOf) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStorageContainer

`func (o *StorageNetAppLunAllOf) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppLunAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppLunAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppLunAllOf) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


