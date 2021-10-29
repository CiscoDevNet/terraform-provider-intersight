# StorageNetAppCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppCluster"]
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**ClusterEfficiency** | Pointer to [**NullableStorageNetAppStorageClusterEfficiency**](StorageNetAppStorageClusterEfficiency.md) |  | [optional] 
**ClusterHealthStatus** | Pointer to **string** | The health status of the cluster. Possible states are ok, ok-with-suppressed, degraded, and unreachable. | [optional] [readonly] 
**DnsDomains** | Pointer to **[]string** |  | [optional] 
**Key** | Pointer to **string** | Unique identifier of NetApp Cluster across data center. | [optional] [readonly] 
**Location** | Pointer to **string** | Location of the storage controller. | [optional] [readonly] 
**ManagementAddress** | Pointer to **string** | FQDN or IP Address of Storage Cluster. | [optional] [readonly] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppCluster

`func NewStorageNetAppCluster(classId string, objectType string, ) *StorageNetAppCluster`

NewStorageNetAppCluster instantiates a new StorageNetAppCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppClusterWithDefaults

`func NewStorageNetAppClusterWithDefaults() *StorageNetAppCluster`

NewStorageNetAppClusterWithDefaults instantiates a new StorageNetAppCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvgPerformanceMetrics

`func (o *StorageNetAppCluster) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppCluster) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppCluster) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppCluster) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetClusterEfficiency

`func (o *StorageNetAppCluster) GetClusterEfficiency() StorageNetAppStorageClusterEfficiency`

GetClusterEfficiency returns the ClusterEfficiency field if non-nil, zero value otherwise.

### GetClusterEfficiencyOk

`func (o *StorageNetAppCluster) GetClusterEfficiencyOk() (*StorageNetAppStorageClusterEfficiency, bool)`

GetClusterEfficiencyOk returns a tuple with the ClusterEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEfficiency

`func (o *StorageNetAppCluster) SetClusterEfficiency(v StorageNetAppStorageClusterEfficiency)`

SetClusterEfficiency sets ClusterEfficiency field to given value.

### HasClusterEfficiency

`func (o *StorageNetAppCluster) HasClusterEfficiency() bool`

HasClusterEfficiency returns a boolean if a field has been set.

### SetClusterEfficiencyNil

`func (o *StorageNetAppCluster) SetClusterEfficiencyNil(b bool)`

 SetClusterEfficiencyNil sets the value for ClusterEfficiency to be an explicit nil

### UnsetClusterEfficiency
`func (o *StorageNetAppCluster) UnsetClusterEfficiency()`

UnsetClusterEfficiency ensures that no value is present for ClusterEfficiency, not even an explicit nil
### GetClusterHealthStatus

`func (o *StorageNetAppCluster) GetClusterHealthStatus() string`

GetClusterHealthStatus returns the ClusterHealthStatus field if non-nil, zero value otherwise.

### GetClusterHealthStatusOk

`func (o *StorageNetAppCluster) GetClusterHealthStatusOk() (*string, bool)`

GetClusterHealthStatusOk returns a tuple with the ClusterHealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHealthStatus

`func (o *StorageNetAppCluster) SetClusterHealthStatus(v string)`

SetClusterHealthStatus sets ClusterHealthStatus field to given value.

### HasClusterHealthStatus

`func (o *StorageNetAppCluster) HasClusterHealthStatus() bool`

HasClusterHealthStatus returns a boolean if a field has been set.

### GetDnsDomains

`func (o *StorageNetAppCluster) GetDnsDomains() []string`

GetDnsDomains returns the DnsDomains field if non-nil, zero value otherwise.

### GetDnsDomainsOk

`func (o *StorageNetAppCluster) GetDnsDomainsOk() (*[]string, bool)`

GetDnsDomainsOk returns a tuple with the DnsDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomains

`func (o *StorageNetAppCluster) SetDnsDomains(v []string)`

SetDnsDomains sets DnsDomains field to given value.

### HasDnsDomains

`func (o *StorageNetAppCluster) HasDnsDomains() bool`

HasDnsDomains returns a boolean if a field has been set.

### SetDnsDomainsNil

`func (o *StorageNetAppCluster) SetDnsDomainsNil(b bool)`

 SetDnsDomainsNil sets the value for DnsDomains to be an explicit nil

### UnsetDnsDomains
`func (o *StorageNetAppCluster) UnsetDnsDomains()`

UnsetDnsDomains ensures that no value is present for DnsDomains, not even an explicit nil
### GetKey

`func (o *StorageNetAppCluster) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppCluster) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppCluster) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppCluster) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *StorageNetAppCluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageNetAppCluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageNetAppCluster) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *StorageNetAppCluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManagementAddress

`func (o *StorageNetAppCluster) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *StorageNetAppCluster) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *StorageNetAppCluster) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *StorageNetAppCluster) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetNameServers

`func (o *StorageNetAppCluster) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *StorageNetAppCluster) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *StorageNetAppCluster) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *StorageNetAppCluster) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### SetNameServersNil

`func (o *StorageNetAppCluster) SetNameServersNil(b bool)`

 SetNameServersNil sets the value for NameServers to be an explicit nil

### UnsetNameServers
`func (o *StorageNetAppCluster) UnsetNameServers()`

UnsetNameServers ensures that no value is present for NameServers, not even an explicit nil
### GetNtpServers

`func (o *StorageNetAppCluster) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *StorageNetAppCluster) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *StorageNetAppCluster) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *StorageNetAppCluster) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *StorageNetAppCluster) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *StorageNetAppCluster) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageNetAppCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageNetAppCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageNetAppCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageNetAppCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


