# StorageNetAppClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppCluster"]
**AutoSupport** | Pointer to [**NullableStorageNetAppAutoSupport**](StorageNetAppAutoSupport.md) |  | [optional] 
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**ClusterEfficiency** | Pointer to [**NullableStorageNetAppStorageClusterEfficiency**](StorageNetAppStorageClusterEfficiency.md) |  | [optional] 
**ClusterHealthStatus** | Pointer to **string** | The health status of the cluster. Possible states are ok, ok-with-suppressed, degraded, and unreachable. * &#x60;Unreachable&#x60; - Cluster status is unreachable. * &#x60;OK&#x60; - Cluster status is either ok or ok-with-suppressed. * &#x60;Degraded&#x60; - Cluster status is degraded. | [optional] [readonly] [default to "Unreachable"]
**DnsDomains** | Pointer to **[]string** |  | [optional] 
**Key** | Pointer to **string** | Unique identifier of NetApp Cluster across data center. | [optional] [readonly] 
**Location** | Pointer to **string** | Location of the storage controller. | [optional] [readonly] 
**ManagementAddress** | Pointer to **string** | FQDN or IP Address of Storage Cluster. | [optional] [readonly] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**RshEnabled** | Pointer to **bool** | Indicates whether or not rsh is enabled on the cluster. | [optional] [readonly] 
**TelnetEnabled** | Pointer to **bool** | Indicates whether or not telnet is enabled on the cluster. | [optional] [readonly] 
**VersionGeneration** | Pointer to **int64** | The generation portion of the version. | [optional] [readonly] 
**VersionMajor** | Pointer to **int64** | The major portion of the version. | [optional] [readonly] 
**VersionMinor** | Pointer to **int64** | The minor portion of the version. | [optional] [readonly] 
**Events** | Pointer to [**[]StorageNetAppClusterEventRelationship**](StorageNetAppClusterEventRelationship.md) | An array of relationships to storageNetAppClusterEvent resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppClusterAllOf

`func NewStorageNetAppClusterAllOf(classId string, objectType string, ) *StorageNetAppClusterAllOf`

NewStorageNetAppClusterAllOf instantiates a new StorageNetAppClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppClusterAllOfWithDefaults

`func NewStorageNetAppClusterAllOfWithDefaults() *StorageNetAppClusterAllOf`

NewStorageNetAppClusterAllOfWithDefaults instantiates a new StorageNetAppClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoSupport

`func (o *StorageNetAppClusterAllOf) GetAutoSupport() StorageNetAppAutoSupport`

GetAutoSupport returns the AutoSupport field if non-nil, zero value otherwise.

### GetAutoSupportOk

`func (o *StorageNetAppClusterAllOf) GetAutoSupportOk() (*StorageNetAppAutoSupport, bool)`

GetAutoSupportOk returns a tuple with the AutoSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSupport

`func (o *StorageNetAppClusterAllOf) SetAutoSupport(v StorageNetAppAutoSupport)`

SetAutoSupport sets AutoSupport field to given value.

### HasAutoSupport

`func (o *StorageNetAppClusterAllOf) HasAutoSupport() bool`

HasAutoSupport returns a boolean if a field has been set.

### SetAutoSupportNil

`func (o *StorageNetAppClusterAllOf) SetAutoSupportNil(b bool)`

 SetAutoSupportNil sets the value for AutoSupport to be an explicit nil

### UnsetAutoSupport
`func (o *StorageNetAppClusterAllOf) UnsetAutoSupport()`

UnsetAutoSupport ensures that no value is present for AutoSupport, not even an explicit nil
### GetAvgPerformanceMetrics

`func (o *StorageNetAppClusterAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppClusterAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppClusterAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppClusterAllOf) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetClusterEfficiency

`func (o *StorageNetAppClusterAllOf) GetClusterEfficiency() StorageNetAppStorageClusterEfficiency`

GetClusterEfficiency returns the ClusterEfficiency field if non-nil, zero value otherwise.

### GetClusterEfficiencyOk

`func (o *StorageNetAppClusterAllOf) GetClusterEfficiencyOk() (*StorageNetAppStorageClusterEfficiency, bool)`

GetClusterEfficiencyOk returns a tuple with the ClusterEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEfficiency

`func (o *StorageNetAppClusterAllOf) SetClusterEfficiency(v StorageNetAppStorageClusterEfficiency)`

SetClusterEfficiency sets ClusterEfficiency field to given value.

### HasClusterEfficiency

`func (o *StorageNetAppClusterAllOf) HasClusterEfficiency() bool`

HasClusterEfficiency returns a boolean if a field has been set.

### SetClusterEfficiencyNil

`func (o *StorageNetAppClusterAllOf) SetClusterEfficiencyNil(b bool)`

 SetClusterEfficiencyNil sets the value for ClusterEfficiency to be an explicit nil

### UnsetClusterEfficiency
`func (o *StorageNetAppClusterAllOf) UnsetClusterEfficiency()`

UnsetClusterEfficiency ensures that no value is present for ClusterEfficiency, not even an explicit nil
### GetClusterHealthStatus

`func (o *StorageNetAppClusterAllOf) GetClusterHealthStatus() string`

GetClusterHealthStatus returns the ClusterHealthStatus field if non-nil, zero value otherwise.

### GetClusterHealthStatusOk

`func (o *StorageNetAppClusterAllOf) GetClusterHealthStatusOk() (*string, bool)`

GetClusterHealthStatusOk returns a tuple with the ClusterHealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHealthStatus

`func (o *StorageNetAppClusterAllOf) SetClusterHealthStatus(v string)`

SetClusterHealthStatus sets ClusterHealthStatus field to given value.

### HasClusterHealthStatus

`func (o *StorageNetAppClusterAllOf) HasClusterHealthStatus() bool`

HasClusterHealthStatus returns a boolean if a field has been set.

### GetDnsDomains

`func (o *StorageNetAppClusterAllOf) GetDnsDomains() []string`

GetDnsDomains returns the DnsDomains field if non-nil, zero value otherwise.

### GetDnsDomainsOk

`func (o *StorageNetAppClusterAllOf) GetDnsDomainsOk() (*[]string, bool)`

GetDnsDomainsOk returns a tuple with the DnsDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomains

`func (o *StorageNetAppClusterAllOf) SetDnsDomains(v []string)`

SetDnsDomains sets DnsDomains field to given value.

### HasDnsDomains

`func (o *StorageNetAppClusterAllOf) HasDnsDomains() bool`

HasDnsDomains returns a boolean if a field has been set.

### SetDnsDomainsNil

`func (o *StorageNetAppClusterAllOf) SetDnsDomainsNil(b bool)`

 SetDnsDomainsNil sets the value for DnsDomains to be an explicit nil

### UnsetDnsDomains
`func (o *StorageNetAppClusterAllOf) UnsetDnsDomains()`

UnsetDnsDomains ensures that no value is present for DnsDomains, not even an explicit nil
### GetKey

`func (o *StorageNetAppClusterAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppClusterAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppClusterAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppClusterAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *StorageNetAppClusterAllOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageNetAppClusterAllOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageNetAppClusterAllOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *StorageNetAppClusterAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManagementAddress

`func (o *StorageNetAppClusterAllOf) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *StorageNetAppClusterAllOf) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *StorageNetAppClusterAllOf) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *StorageNetAppClusterAllOf) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetNameServers

`func (o *StorageNetAppClusterAllOf) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *StorageNetAppClusterAllOf) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *StorageNetAppClusterAllOf) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *StorageNetAppClusterAllOf) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### SetNameServersNil

`func (o *StorageNetAppClusterAllOf) SetNameServersNil(b bool)`

 SetNameServersNil sets the value for NameServers to be an explicit nil

### UnsetNameServers
`func (o *StorageNetAppClusterAllOf) UnsetNameServers()`

UnsetNameServers ensures that no value is present for NameServers, not even an explicit nil
### GetNtpServers

`func (o *StorageNetAppClusterAllOf) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *StorageNetAppClusterAllOf) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *StorageNetAppClusterAllOf) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *StorageNetAppClusterAllOf) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *StorageNetAppClusterAllOf) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *StorageNetAppClusterAllOf) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetRshEnabled

`func (o *StorageNetAppClusterAllOf) GetRshEnabled() bool`

GetRshEnabled returns the RshEnabled field if non-nil, zero value otherwise.

### GetRshEnabledOk

`func (o *StorageNetAppClusterAllOf) GetRshEnabledOk() (*bool, bool)`

GetRshEnabledOk returns a tuple with the RshEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRshEnabled

`func (o *StorageNetAppClusterAllOf) SetRshEnabled(v bool)`

SetRshEnabled sets RshEnabled field to given value.

### HasRshEnabled

`func (o *StorageNetAppClusterAllOf) HasRshEnabled() bool`

HasRshEnabled returns a boolean if a field has been set.

### GetTelnetEnabled

`func (o *StorageNetAppClusterAllOf) GetTelnetEnabled() bool`

GetTelnetEnabled returns the TelnetEnabled field if non-nil, zero value otherwise.

### GetTelnetEnabledOk

`func (o *StorageNetAppClusterAllOf) GetTelnetEnabledOk() (*bool, bool)`

GetTelnetEnabledOk returns a tuple with the TelnetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnetEnabled

`func (o *StorageNetAppClusterAllOf) SetTelnetEnabled(v bool)`

SetTelnetEnabled sets TelnetEnabled field to given value.

### HasTelnetEnabled

`func (o *StorageNetAppClusterAllOf) HasTelnetEnabled() bool`

HasTelnetEnabled returns a boolean if a field has been set.

### GetVersionGeneration

`func (o *StorageNetAppClusterAllOf) GetVersionGeneration() int64`

GetVersionGeneration returns the VersionGeneration field if non-nil, zero value otherwise.

### GetVersionGenerationOk

`func (o *StorageNetAppClusterAllOf) GetVersionGenerationOk() (*int64, bool)`

GetVersionGenerationOk returns a tuple with the VersionGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGeneration

`func (o *StorageNetAppClusterAllOf) SetVersionGeneration(v int64)`

SetVersionGeneration sets VersionGeneration field to given value.

### HasVersionGeneration

`func (o *StorageNetAppClusterAllOf) HasVersionGeneration() bool`

HasVersionGeneration returns a boolean if a field has been set.

### GetVersionMajor

`func (o *StorageNetAppClusterAllOf) GetVersionMajor() int64`

GetVersionMajor returns the VersionMajor field if non-nil, zero value otherwise.

### GetVersionMajorOk

`func (o *StorageNetAppClusterAllOf) GetVersionMajorOk() (*int64, bool)`

GetVersionMajorOk returns a tuple with the VersionMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMajor

`func (o *StorageNetAppClusterAllOf) SetVersionMajor(v int64)`

SetVersionMajor sets VersionMajor field to given value.

### HasVersionMajor

`func (o *StorageNetAppClusterAllOf) HasVersionMajor() bool`

HasVersionMajor returns a boolean if a field has been set.

### GetVersionMinor

`func (o *StorageNetAppClusterAllOf) GetVersionMinor() int64`

GetVersionMinor returns the VersionMinor field if non-nil, zero value otherwise.

### GetVersionMinorOk

`func (o *StorageNetAppClusterAllOf) GetVersionMinorOk() (*int64, bool)`

GetVersionMinorOk returns a tuple with the VersionMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMinor

`func (o *StorageNetAppClusterAllOf) SetVersionMinor(v int64)`

SetVersionMinor sets VersionMinor field to given value.

### HasVersionMinor

`func (o *StorageNetAppClusterAllOf) HasVersionMinor() bool`

HasVersionMinor returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppClusterAllOf) GetEvents() []StorageNetAppClusterEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppClusterAllOf) GetEventsOk() (*[]StorageNetAppClusterEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppClusterAllOf) SetEvents(v []StorageNetAppClusterEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppClusterAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppClusterAllOf) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppClusterAllOf) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageNetAppClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageNetAppClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageNetAppClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageNetAppClusterAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


