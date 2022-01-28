# StorageNetAppStorageVmAllOf

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

### NewStorageNetAppStorageVmAllOf

`func NewStorageNetAppStorageVmAllOf(classId string, objectType string, ) *StorageNetAppStorageVmAllOf`

NewStorageNetAppStorageVmAllOf instantiates a new StorageNetAppStorageVmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppStorageVmAllOfWithDefaults

`func NewStorageNetAppStorageVmAllOfWithDefaults() *StorageNetAppStorageVmAllOf`

NewStorageNetAppStorageVmAllOfWithDefaults instantiates a new StorageNetAppStorageVmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppStorageVmAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppStorageVmAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppStorageVmAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppStorageVmAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppStorageVmAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppStorageVmAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregates

`func (o *StorageNetAppStorageVmAllOf) GetAggregates() []string`

GetAggregates returns the Aggregates field if non-nil, zero value otherwise.

### GetAggregatesOk

`func (o *StorageNetAppStorageVmAllOf) GetAggregatesOk() (*[]string, bool)`

GetAggregatesOk returns a tuple with the Aggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregates

`func (o *StorageNetAppStorageVmAllOf) SetAggregates(v []string)`

SetAggregates sets Aggregates field to given value.

### HasAggregates

`func (o *StorageNetAppStorageVmAllOf) HasAggregates() bool`

HasAggregates returns a boolean if a field has been set.

### SetAggregatesNil

`func (o *StorageNetAppStorageVmAllOf) SetAggregatesNil(b bool)`

 SetAggregatesNil sets the value for Aggregates to be an explicit nil

### UnsetAggregates
`func (o *StorageNetAppStorageVmAllOf) UnsetAggregates()`

UnsetAggregates ensures that no value is present for Aggregates, not even an explicit nil
### GetAvgPerformanceMetrics

`func (o *StorageNetAppStorageVmAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppStorageVmAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppStorageVmAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppStorageVmAllOf) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetCifsEnabled

`func (o *StorageNetAppStorageVmAllOf) GetCifsEnabled() bool`

GetCifsEnabled returns the CifsEnabled field if non-nil, zero value otherwise.

### GetCifsEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetCifsEnabledOk() (*bool, bool)`

GetCifsEnabledOk returns a tuple with the CifsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsEnabled

`func (o *StorageNetAppStorageVmAllOf) SetCifsEnabled(v bool)`

SetCifsEnabled sets CifsEnabled field to given value.

### HasCifsEnabled

`func (o *StorageNetAppStorageVmAllOf) HasCifsEnabled() bool`

HasCifsEnabled returns a boolean if a field has been set.

### GetDnsDomains

`func (o *StorageNetAppStorageVmAllOf) GetDnsDomains() []string`

GetDnsDomains returns the DnsDomains field if non-nil, zero value otherwise.

### GetDnsDomainsOk

`func (o *StorageNetAppStorageVmAllOf) GetDnsDomainsOk() (*[]string, bool)`

GetDnsDomainsOk returns a tuple with the DnsDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomains

`func (o *StorageNetAppStorageVmAllOf) SetDnsDomains(v []string)`

SetDnsDomains sets DnsDomains field to given value.

### HasDnsDomains

`func (o *StorageNetAppStorageVmAllOf) HasDnsDomains() bool`

HasDnsDomains returns a boolean if a field has been set.

### SetDnsDomainsNil

`func (o *StorageNetAppStorageVmAllOf) SetDnsDomainsNil(b bool)`

 SetDnsDomainsNil sets the value for DnsDomains to be an explicit nil

### UnsetDnsDomains
`func (o *StorageNetAppStorageVmAllOf) UnsetDnsDomains()`

UnsetDnsDomains ensures that no value is present for DnsDomains, not even an explicit nil
### GetFcpEnabled

`func (o *StorageNetAppStorageVmAllOf) GetFcpEnabled() bool`

GetFcpEnabled returns the FcpEnabled field if non-nil, zero value otherwise.

### GetFcpEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetFcpEnabledOk() (*bool, bool)`

GetFcpEnabledOk returns a tuple with the FcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcpEnabled

`func (o *StorageNetAppStorageVmAllOf) SetFcpEnabled(v bool)`

SetFcpEnabled sets FcpEnabled field to given value.

### HasFcpEnabled

`func (o *StorageNetAppStorageVmAllOf) HasFcpEnabled() bool`

HasFcpEnabled returns a boolean if a field has been set.

### GetIpspace

`func (o *StorageNetAppStorageVmAllOf) GetIpspace() string`

GetIpspace returns the Ipspace field if non-nil, zero value otherwise.

### GetIpspaceOk

`func (o *StorageNetAppStorageVmAllOf) GetIpspaceOk() (*string, bool)`

GetIpspaceOk returns a tuple with the Ipspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpspace

`func (o *StorageNetAppStorageVmAllOf) SetIpspace(v string)`

SetIpspace sets Ipspace field to given value.

### HasIpspace

`func (o *StorageNetAppStorageVmAllOf) HasIpspace() bool`

HasIpspace returns a boolean if a field has been set.

### GetIscsiEnabled

`func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabled() bool`

GetIscsiEnabled returns the IscsiEnabled field if non-nil, zero value otherwise.

### GetIscsiEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabledOk() (*bool, bool)`

GetIscsiEnabledOk returns a tuple with the IscsiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiEnabled

`func (o *StorageNetAppStorageVmAllOf) SetIscsiEnabled(v bool)`

SetIscsiEnabled sets IscsiEnabled field to given value.

### HasIscsiEnabled

`func (o *StorageNetAppStorageVmAllOf) HasIscsiEnabled() bool`

HasIscsiEnabled returns a boolean if a field has been set.

### GetKey

`func (o *StorageNetAppStorageVmAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppStorageVmAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppStorageVmAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppStorageVmAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNameServers

`func (o *StorageNetAppStorageVmAllOf) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *StorageNetAppStorageVmAllOf) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *StorageNetAppStorageVmAllOf) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *StorageNetAppStorageVmAllOf) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### SetNameServersNil

`func (o *StorageNetAppStorageVmAllOf) SetNameServersNil(b bool)`

 SetNameServersNil sets the value for NameServers to be an explicit nil

### UnsetNameServers
`func (o *StorageNetAppStorageVmAllOf) UnsetNameServers()`

UnsetNameServers ensures that no value is present for NameServers, not even an explicit nil
### GetNfsEnabled

`func (o *StorageNetAppStorageVmAllOf) GetNfsEnabled() bool`

GetNfsEnabled returns the NfsEnabled field if non-nil, zero value otherwise.

### GetNfsEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetNfsEnabledOk() (*bool, bool)`

GetNfsEnabledOk returns a tuple with the NfsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsEnabled

`func (o *StorageNetAppStorageVmAllOf) SetNfsEnabled(v bool)`

SetNfsEnabled sets NfsEnabled field to given value.

### HasNfsEnabled

`func (o *StorageNetAppStorageVmAllOf) HasNfsEnabled() bool`

HasNfsEnabled returns a boolean if a field has been set.

### GetNfsV3Enabled

`func (o *StorageNetAppStorageVmAllOf) GetNfsV3Enabled() bool`

GetNfsV3Enabled returns the NfsV3Enabled field if non-nil, zero value otherwise.

### GetNfsV3EnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetNfsV3EnabledOk() (*bool, bool)`

GetNfsV3EnabledOk returns a tuple with the NfsV3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV3Enabled

`func (o *StorageNetAppStorageVmAllOf) SetNfsV3Enabled(v bool)`

SetNfsV3Enabled sets NfsV3Enabled field to given value.

### HasNfsV3Enabled

`func (o *StorageNetAppStorageVmAllOf) HasNfsV3Enabled() bool`

HasNfsV3Enabled returns a boolean if a field has been set.

### GetNfsV41Enabled

`func (o *StorageNetAppStorageVmAllOf) GetNfsV41Enabled() bool`

GetNfsV41Enabled returns the NfsV41Enabled field if non-nil, zero value otherwise.

### GetNfsV41EnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetNfsV41EnabledOk() (*bool, bool)`

GetNfsV41EnabledOk returns a tuple with the NfsV41Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV41Enabled

`func (o *StorageNetAppStorageVmAllOf) SetNfsV41Enabled(v bool)`

SetNfsV41Enabled sets NfsV41Enabled field to given value.

### HasNfsV41Enabled

`func (o *StorageNetAppStorageVmAllOf) HasNfsV41Enabled() bool`

HasNfsV41Enabled returns a boolean if a field has been set.

### GetNfsV4Enabled

`func (o *StorageNetAppStorageVmAllOf) GetNfsV4Enabled() bool`

GetNfsV4Enabled returns the NfsV4Enabled field if non-nil, zero value otherwise.

### GetNfsV4EnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetNfsV4EnabledOk() (*bool, bool)`

GetNfsV4EnabledOk returns a tuple with the NfsV4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV4Enabled

`func (o *StorageNetAppStorageVmAllOf) SetNfsV4Enabled(v bool)`

SetNfsV4Enabled sets NfsV4Enabled field to given value.

### HasNfsV4Enabled

`func (o *StorageNetAppStorageVmAllOf) HasNfsV4Enabled() bool`

HasNfsV4Enabled returns a boolean if a field has been set.

### GetNvmeEnabled

`func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabled() bool`

GetNvmeEnabled returns the NvmeEnabled field if non-nil, zero value otherwise.

### GetNvmeEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabledOk() (*bool, bool)`

GetNvmeEnabledOk returns a tuple with the NvmeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmeEnabled

`func (o *StorageNetAppStorageVmAllOf) SetNvmeEnabled(v bool)`

SetNvmeEnabled sets NvmeEnabled field to given value.

### HasNvmeEnabled

`func (o *StorageNetAppStorageVmAllOf) HasNvmeEnabled() bool`

HasNvmeEnabled returns a boolean if a field has been set.

### GetSubtype

`func (o *StorageNetAppStorageVmAllOf) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *StorageNetAppStorageVmAllOf) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *StorageNetAppStorageVmAllOf) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *StorageNetAppStorageVmAllOf) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppStorageVmAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppStorageVmAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppStorageVmAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppStorageVmAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetDiskPool

`func (o *StorageNetAppStorageVmAllOf) GetDiskPool() []StorageNetAppAggregateRelationship`

GetDiskPool returns the DiskPool field if non-nil, zero value otherwise.

### GetDiskPoolOk

`func (o *StorageNetAppStorageVmAllOf) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool)`

GetDiskPoolOk returns a tuple with the DiskPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPool

`func (o *StorageNetAppStorageVmAllOf) SetDiskPool(v []StorageNetAppAggregateRelationship)`

SetDiskPool sets DiskPool field to given value.

### HasDiskPool

`func (o *StorageNetAppStorageVmAllOf) HasDiskPool() bool`

HasDiskPool returns a boolean if a field has been set.

### SetDiskPoolNil

`func (o *StorageNetAppStorageVmAllOf) SetDiskPoolNil(b bool)`

 SetDiskPoolNil sets the value for DiskPool to be an explicit nil

### UnsetDiskPool
`func (o *StorageNetAppStorageVmAllOf) UnsetDiskPool()`

UnsetDiskPool ensures that no value is present for DiskPool, not even an explicit nil
### GetEvents

`func (o *StorageNetAppStorageVmAllOf) GetEvents() []StorageNetAppSvmEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppStorageVmAllOf) GetEventsOk() (*[]StorageNetAppSvmEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppStorageVmAllOf) SetEvents(v []StorageNetAppSvmEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppStorageVmAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppStorageVmAllOf) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppStorageVmAllOf) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


