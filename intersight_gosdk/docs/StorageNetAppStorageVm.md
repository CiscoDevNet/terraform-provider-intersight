# StorageNetAppStorageVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppStorageVm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppStorageVm"]
**Aggregates** | Pointer to **[]string** |  | [optional] 
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**CifsEnabled** | Pointer to **bool** | Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers. | [optional] [readonly] 
**DnsDomains** | Pointer to **[]string** |  | [optional] 
**FcpEnabled** | Pointer to **bool** | Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers. | [optional] [readonly] 
**Ipspace** | Pointer to **string** | IPspace name. IPspaces are distinct IP address spaces in which storage virtual machines (SVMs) reside. | [optional] [readonly] 
**IscsiEnabled** | Pointer to **bool** | Status for iSCSI protocol allowed to run on Vservers. | [optional] [readonly] 
**Key** | Pointer to **string** | Unique identifier of VServer across data center. | [optional] [readonly] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**NfsEnabled** | Pointer to **bool** | Status for Network File System Protocol ( NFS ) allowed to run on  Vservers. | [optional] [readonly] 
**NfsV3Enabled** | Pointer to **bool** | Status for Network File System Protocol ( NFSv3 ) allowed to run on  Vservers. | [optional] [readonly] 
**NfsV41Enabled** | Pointer to **bool** | Status for Network File System Protocol ( NFSv4.1 ) allowed to run on  Vservers. | [optional] [readonly] 
**NfsV4Enabled** | Pointer to **bool** | Status for Network File System Protocol ( NFSv4 ) allowed to run on  Vservers. | [optional] [readonly] 
**NvmeEnabled** | Pointer to **bool** | Status for NVME protocol allowed to run on Vservers. | [optional] [readonly] 
**Subtype** | Pointer to **string** | SVM subtype (default, dp_destination, sync_source, or sync_destination). The SVM subtype sync_destination is created automatically when an SVM of subtype sync_source is created on the source MetroCluster cluster. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**DiskPool** | Pointer to [**[]StorageNetAppAggregateRelationship**](StorageNetAppAggregateRelationship.md) | An array of relationships to storageNetAppAggregate resources. | [optional] [readonly] 
**Events** | Pointer to [**[]StorageNetAppSvmEventRelationship**](StorageNetAppSvmEventRelationship.md) | An array of relationships to storageNetAppSvmEvent resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppStorageVm

`func NewStorageNetAppStorageVm(classId string, objectType string, ) *StorageNetAppStorageVm`

NewStorageNetAppStorageVm instantiates a new StorageNetAppStorageVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppStorageVmWithDefaults

`func NewStorageNetAppStorageVmWithDefaults() *StorageNetAppStorageVm`

NewStorageNetAppStorageVmWithDefaults instantiates a new StorageNetAppStorageVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppStorageVm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppStorageVm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppStorageVm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppStorageVm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppStorageVm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppStorageVm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregates

`func (o *StorageNetAppStorageVm) GetAggregates() []string`

GetAggregates returns the Aggregates field if non-nil, zero value otherwise.

### GetAggregatesOk

`func (o *StorageNetAppStorageVm) GetAggregatesOk() (*[]string, bool)`

GetAggregatesOk returns a tuple with the Aggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregates

`func (o *StorageNetAppStorageVm) SetAggregates(v []string)`

SetAggregates sets Aggregates field to given value.

### HasAggregates

`func (o *StorageNetAppStorageVm) HasAggregates() bool`

HasAggregates returns a boolean if a field has been set.

### SetAggregatesNil

`func (o *StorageNetAppStorageVm) SetAggregatesNil(b bool)`

 SetAggregatesNil sets the value for Aggregates to be an explicit nil

### UnsetAggregates
`func (o *StorageNetAppStorageVm) UnsetAggregates()`

UnsetAggregates ensures that no value is present for Aggregates, not even an explicit nil
### GetAvgPerformanceMetrics

`func (o *StorageNetAppStorageVm) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppStorageVm) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppStorageVm) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppStorageVm) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetCifsEnabled

`func (o *StorageNetAppStorageVm) GetCifsEnabled() bool`

GetCifsEnabled returns the CifsEnabled field if non-nil, zero value otherwise.

### GetCifsEnabledOk

`func (o *StorageNetAppStorageVm) GetCifsEnabledOk() (*bool, bool)`

GetCifsEnabledOk returns a tuple with the CifsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsEnabled

`func (o *StorageNetAppStorageVm) SetCifsEnabled(v bool)`

SetCifsEnabled sets CifsEnabled field to given value.

### HasCifsEnabled

`func (o *StorageNetAppStorageVm) HasCifsEnabled() bool`

HasCifsEnabled returns a boolean if a field has been set.

### GetDnsDomains

`func (o *StorageNetAppStorageVm) GetDnsDomains() []string`

GetDnsDomains returns the DnsDomains field if non-nil, zero value otherwise.

### GetDnsDomainsOk

`func (o *StorageNetAppStorageVm) GetDnsDomainsOk() (*[]string, bool)`

GetDnsDomainsOk returns a tuple with the DnsDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomains

`func (o *StorageNetAppStorageVm) SetDnsDomains(v []string)`

SetDnsDomains sets DnsDomains field to given value.

### HasDnsDomains

`func (o *StorageNetAppStorageVm) HasDnsDomains() bool`

HasDnsDomains returns a boolean if a field has been set.

### SetDnsDomainsNil

`func (o *StorageNetAppStorageVm) SetDnsDomainsNil(b bool)`

 SetDnsDomainsNil sets the value for DnsDomains to be an explicit nil

### UnsetDnsDomains
`func (o *StorageNetAppStorageVm) UnsetDnsDomains()`

UnsetDnsDomains ensures that no value is present for DnsDomains, not even an explicit nil
### GetFcpEnabled

`func (o *StorageNetAppStorageVm) GetFcpEnabled() bool`

GetFcpEnabled returns the FcpEnabled field if non-nil, zero value otherwise.

### GetFcpEnabledOk

`func (o *StorageNetAppStorageVm) GetFcpEnabledOk() (*bool, bool)`

GetFcpEnabledOk returns a tuple with the FcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcpEnabled

`func (o *StorageNetAppStorageVm) SetFcpEnabled(v bool)`

SetFcpEnabled sets FcpEnabled field to given value.

### HasFcpEnabled

`func (o *StorageNetAppStorageVm) HasFcpEnabled() bool`

HasFcpEnabled returns a boolean if a field has been set.

### GetIpspace

`func (o *StorageNetAppStorageVm) GetIpspace() string`

GetIpspace returns the Ipspace field if non-nil, zero value otherwise.

### GetIpspaceOk

`func (o *StorageNetAppStorageVm) GetIpspaceOk() (*string, bool)`

GetIpspaceOk returns a tuple with the Ipspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpspace

`func (o *StorageNetAppStorageVm) SetIpspace(v string)`

SetIpspace sets Ipspace field to given value.

### HasIpspace

`func (o *StorageNetAppStorageVm) HasIpspace() bool`

HasIpspace returns a boolean if a field has been set.

### GetIscsiEnabled

`func (o *StorageNetAppStorageVm) GetIscsiEnabled() bool`

GetIscsiEnabled returns the IscsiEnabled field if non-nil, zero value otherwise.

### GetIscsiEnabledOk

`func (o *StorageNetAppStorageVm) GetIscsiEnabledOk() (*bool, bool)`

GetIscsiEnabledOk returns a tuple with the IscsiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiEnabled

`func (o *StorageNetAppStorageVm) SetIscsiEnabled(v bool)`

SetIscsiEnabled sets IscsiEnabled field to given value.

### HasIscsiEnabled

`func (o *StorageNetAppStorageVm) HasIscsiEnabled() bool`

HasIscsiEnabled returns a boolean if a field has been set.

### GetKey

`func (o *StorageNetAppStorageVm) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppStorageVm) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppStorageVm) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppStorageVm) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNameServers

`func (o *StorageNetAppStorageVm) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *StorageNetAppStorageVm) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *StorageNetAppStorageVm) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *StorageNetAppStorageVm) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### SetNameServersNil

`func (o *StorageNetAppStorageVm) SetNameServersNil(b bool)`

 SetNameServersNil sets the value for NameServers to be an explicit nil

### UnsetNameServers
`func (o *StorageNetAppStorageVm) UnsetNameServers()`

UnsetNameServers ensures that no value is present for NameServers, not even an explicit nil
### GetNfsEnabled

`func (o *StorageNetAppStorageVm) GetNfsEnabled() bool`

GetNfsEnabled returns the NfsEnabled field if non-nil, zero value otherwise.

### GetNfsEnabledOk

`func (o *StorageNetAppStorageVm) GetNfsEnabledOk() (*bool, bool)`

GetNfsEnabledOk returns a tuple with the NfsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsEnabled

`func (o *StorageNetAppStorageVm) SetNfsEnabled(v bool)`

SetNfsEnabled sets NfsEnabled field to given value.

### HasNfsEnabled

`func (o *StorageNetAppStorageVm) HasNfsEnabled() bool`

HasNfsEnabled returns a boolean if a field has been set.

### GetNfsV3Enabled

`func (o *StorageNetAppStorageVm) GetNfsV3Enabled() bool`

GetNfsV3Enabled returns the NfsV3Enabled field if non-nil, zero value otherwise.

### GetNfsV3EnabledOk

`func (o *StorageNetAppStorageVm) GetNfsV3EnabledOk() (*bool, bool)`

GetNfsV3EnabledOk returns a tuple with the NfsV3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV3Enabled

`func (o *StorageNetAppStorageVm) SetNfsV3Enabled(v bool)`

SetNfsV3Enabled sets NfsV3Enabled field to given value.

### HasNfsV3Enabled

`func (o *StorageNetAppStorageVm) HasNfsV3Enabled() bool`

HasNfsV3Enabled returns a boolean if a field has been set.

### GetNfsV41Enabled

`func (o *StorageNetAppStorageVm) GetNfsV41Enabled() bool`

GetNfsV41Enabled returns the NfsV41Enabled field if non-nil, zero value otherwise.

### GetNfsV41EnabledOk

`func (o *StorageNetAppStorageVm) GetNfsV41EnabledOk() (*bool, bool)`

GetNfsV41EnabledOk returns a tuple with the NfsV41Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV41Enabled

`func (o *StorageNetAppStorageVm) SetNfsV41Enabled(v bool)`

SetNfsV41Enabled sets NfsV41Enabled field to given value.

### HasNfsV41Enabled

`func (o *StorageNetAppStorageVm) HasNfsV41Enabled() bool`

HasNfsV41Enabled returns a boolean if a field has been set.

### GetNfsV4Enabled

`func (o *StorageNetAppStorageVm) GetNfsV4Enabled() bool`

GetNfsV4Enabled returns the NfsV4Enabled field if non-nil, zero value otherwise.

### GetNfsV4EnabledOk

`func (o *StorageNetAppStorageVm) GetNfsV4EnabledOk() (*bool, bool)`

GetNfsV4EnabledOk returns a tuple with the NfsV4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV4Enabled

`func (o *StorageNetAppStorageVm) SetNfsV4Enabled(v bool)`

SetNfsV4Enabled sets NfsV4Enabled field to given value.

### HasNfsV4Enabled

`func (o *StorageNetAppStorageVm) HasNfsV4Enabled() bool`

HasNfsV4Enabled returns a boolean if a field has been set.

### GetNvmeEnabled

`func (o *StorageNetAppStorageVm) GetNvmeEnabled() bool`

GetNvmeEnabled returns the NvmeEnabled field if non-nil, zero value otherwise.

### GetNvmeEnabledOk

`func (o *StorageNetAppStorageVm) GetNvmeEnabledOk() (*bool, bool)`

GetNvmeEnabledOk returns a tuple with the NvmeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmeEnabled

`func (o *StorageNetAppStorageVm) SetNvmeEnabled(v bool)`

SetNvmeEnabled sets NvmeEnabled field to given value.

### HasNvmeEnabled

`func (o *StorageNetAppStorageVm) HasNvmeEnabled() bool`

HasNvmeEnabled returns a boolean if a field has been set.

### GetSubtype

`func (o *StorageNetAppStorageVm) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *StorageNetAppStorageVm) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *StorageNetAppStorageVm) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *StorageNetAppStorageVm) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppStorageVm) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppStorageVm) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppStorageVm) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppStorageVm) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetDiskPool

`func (o *StorageNetAppStorageVm) GetDiskPool() []StorageNetAppAggregateRelationship`

GetDiskPool returns the DiskPool field if non-nil, zero value otherwise.

### GetDiskPoolOk

`func (o *StorageNetAppStorageVm) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool)`

GetDiskPoolOk returns a tuple with the DiskPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPool

`func (o *StorageNetAppStorageVm) SetDiskPool(v []StorageNetAppAggregateRelationship)`

SetDiskPool sets DiskPool field to given value.

### HasDiskPool

`func (o *StorageNetAppStorageVm) HasDiskPool() bool`

HasDiskPool returns a boolean if a field has been set.

### SetDiskPoolNil

`func (o *StorageNetAppStorageVm) SetDiskPoolNil(b bool)`

 SetDiskPoolNil sets the value for DiskPool to be an explicit nil

### UnsetDiskPool
`func (o *StorageNetAppStorageVm) UnsetDiskPool()`

UnsetDiskPool ensures that no value is present for DiskPool, not even an explicit nil
### GetEvents

`func (o *StorageNetAppStorageVm) GetEvents() []StorageNetAppSvmEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppStorageVm) GetEventsOk() (*[]StorageNetAppSvmEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppStorageVm) SetEvents(v []StorageNetAppSvmEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppStorageVm) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppStorageVm) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppStorageVm) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


