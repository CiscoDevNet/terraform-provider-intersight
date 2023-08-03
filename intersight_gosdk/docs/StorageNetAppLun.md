# StorageNetAppLun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppLun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppLun"]
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**ContainerState** | Pointer to **string** | The state of the volume and aggregate that contain the LUN. LUNs are only available when their containers are available. | [optional] [readonly] 
**IsMapped** | Pointer to **string** | Reports if the LUN is mapped to one or more initiator groups. | [optional] [readonly] 
**Key** | Pointer to **string** | Unique identifier of LUN across data center. | [optional] [readonly] 
**Mapped** | Pointer to **bool** | Reports if the LUN is mapped to one or more initiator groups. | [optional] [readonly] 
**OsType** | Pointer to **string** | The operating system (OS) type for this LUN. * &#x60;Linux&#x60; - Family of open source Unix-like operating systems based on the Linux kernel. * &#x60;AIX&#x60; - Advanced Interactive Executive (AIX). * &#x60;HP-UX&#x60; - HP-UX is implementation of the Unix operating system, based on Unix System V. * &#x60;Hyper-V&#x60; - Windows Server 2008 or Windows Server 2012 Hyper-V. * &#x60;OpenVMS&#x60; - OpenVMS is multi-user, multiprocessing virtual memory-based operating system. * &#x60;Solaris&#x60; - Solaris is a Unix operating system. * &#x60;NetWare&#x60; - NetWare is a computer network operating system. * &#x60;VMware&#x60; - An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers. * &#x60;Windows&#x60; - Single-partition Windows disk using the Master Boot Record (MBR) partitioning style. * &#x60;Xen&#x60; - Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently. | [optional] [readonly] [default to "Linux"]
**Path** | Pointer to **string** | Path where the LUN is mounted. | [optional] [readonly] 
**Serial** | Pointer to **string** | Serial number for the provisioned LUN. | [optional] [readonly] 
**State** | Pointer to **string** | The administrative state of a LUN. * &#x60;offline&#x60; - The LUN is administratively offline, or a more detailed offline reason is not available. * &#x60;online&#x60; - The state of the LUN is online. | [optional] [readonly] [default to "offline"]
**SvmName** | Pointer to **string** | The storage virtual machine name for the lun. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally unique identifier of the LUN. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | The parent volume name for the lun. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppLunEventRelationship**](StorageNetAppLunEventRelationship.md) | An array of relationships to storageNetAppLunEvent resources. | [optional] [readonly] 
**Host** | Pointer to [**[]StorageNetAppInitiatorGroupRelationship**](StorageNetAppInitiatorGroupRelationship.md) | An array of relationships to storageNetAppInitiatorGroup resources. | [optional] [readonly] 
**StorageContainer** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppLun

`func NewStorageNetAppLun(classId string, objectType string, ) *StorageNetAppLun`

NewStorageNetAppLun instantiates a new StorageNetAppLun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppLunWithDefaults

`func NewStorageNetAppLunWithDefaults() *StorageNetAppLun`

NewStorageNetAppLunWithDefaults instantiates a new StorageNetAppLun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppLun) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppLun) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppLun) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppLun) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppLun) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppLun) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvgPerformanceMetrics

`func (o *StorageNetAppLun) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppLun) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppLun) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppLun) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetContainerState

`func (o *StorageNetAppLun) GetContainerState() string`

GetContainerState returns the ContainerState field if non-nil, zero value otherwise.

### GetContainerStateOk

`func (o *StorageNetAppLun) GetContainerStateOk() (*string, bool)`

GetContainerStateOk returns a tuple with the ContainerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerState

`func (o *StorageNetAppLun) SetContainerState(v string)`

SetContainerState sets ContainerState field to given value.

### HasContainerState

`func (o *StorageNetAppLun) HasContainerState() bool`

HasContainerState returns a boolean if a field has been set.

### GetIsMapped

`func (o *StorageNetAppLun) GetIsMapped() string`

GetIsMapped returns the IsMapped field if non-nil, zero value otherwise.

### GetIsMappedOk

`func (o *StorageNetAppLun) GetIsMappedOk() (*string, bool)`

GetIsMappedOk returns a tuple with the IsMapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMapped

`func (o *StorageNetAppLun) SetIsMapped(v string)`

SetIsMapped sets IsMapped field to given value.

### HasIsMapped

`func (o *StorageNetAppLun) HasIsMapped() bool`

HasIsMapped returns a boolean if a field has been set.

### GetKey

`func (o *StorageNetAppLun) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppLun) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppLun) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppLun) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMapped

`func (o *StorageNetAppLun) GetMapped() bool`

GetMapped returns the Mapped field if non-nil, zero value otherwise.

### GetMappedOk

`func (o *StorageNetAppLun) GetMappedOk() (*bool, bool)`

GetMappedOk returns a tuple with the Mapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapped

`func (o *StorageNetAppLun) SetMapped(v bool)`

SetMapped sets Mapped field to given value.

### HasMapped

`func (o *StorageNetAppLun) HasMapped() bool`

HasMapped returns a boolean if a field has been set.

### GetOsType

`func (o *StorageNetAppLun) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *StorageNetAppLun) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *StorageNetAppLun) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *StorageNetAppLun) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetPath

`func (o *StorageNetAppLun) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageNetAppLun) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageNetAppLun) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageNetAppLun) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSerial

`func (o *StorageNetAppLun) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StorageNetAppLun) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StorageNetAppLun) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *StorageNetAppLun) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppLun) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppLun) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppLun) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppLun) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSvmName

`func (o *StorageNetAppLun) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppLun) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppLun) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppLun) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppLun) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppLun) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppLun) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppLun) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageNetAppLun) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageNetAppLun) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageNetAppLun) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageNetAppLun) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppLun) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppLun) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppLun) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppLun) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppLun) GetEvents() []StorageNetAppLunEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppLun) GetEventsOk() (*[]StorageNetAppLunEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppLun) SetEvents(v []StorageNetAppLunEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppLun) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppLun) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppLun) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetHost

`func (o *StorageNetAppLun) GetHost() []StorageNetAppInitiatorGroupRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageNetAppLun) GetHostOk() (*[]StorageNetAppInitiatorGroupRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageNetAppLun) SetHost(v []StorageNetAppInitiatorGroupRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageNetAppLun) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *StorageNetAppLun) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *StorageNetAppLun) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetStorageContainer

`func (o *StorageNetAppLun) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppLun) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppLun) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppLun) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


