# HyperflexCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Cluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Cluster"]
**Capability** | Pointer to [**NullableHyperflexCapability**](HyperflexCapability.md) |  | [optional] 
**ClusterType** | Pointer to **int64** | The storage type of this cluster (All Flash or Hybrid). | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | The unique identifier for this HyperFlex cluster. | [optional] [readonly] 
**DeviceId** | Pointer to **string** | The unique identifier of the device registration that represents this HyperFlex cluster&#39;s connection to Intersight. | [optional] [readonly] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**EncryptionStatus** | Pointer to **string** | This captures the encryption status for a HyperFlex cluster. Currently it will have the status if HXA-CLU-0020 alarm is raised. In the future it can capture other details. | [optional] 
**FltAggr** | Pointer to **int64** | The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms. | [optional] [readonly] 
**HxdpBuildVersion** | Pointer to **string** | The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version. | [optional] [readonly] 
**NetworkConfiguration** | Pointer to [**NullableHyperflexNetworkConfiguration**](HyperflexNetworkConfiguration.md) |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to [**NullableHyperflexSummary**](HyperflexSummary.md) |  | [optional] 
**UpgradeStatus** | Pointer to **string** | The upgrade status of the HyperFlex cluster. * &#x60;Unknown&#x60; - The upgrade status of the HyperFlex cluster could not be determined. * &#x60;Ok&#x60; - The upgrade of the HyperFlex cluster is complete. * &#x60;InProgress&#x60; - The upgrade of the HyperFlex cluster is in-progress. * &#x60;Failed&#x60; - The upgrade of the HyperFlex cluster has failed. * &#x60;Waiting&#x60; - The upgrade of the HyperFlex cluster is waiting to continue execution. | [optional] [readonly] [default to "Unknown"]
**UplinkSpeed** | Pointer to **string** | The uplink speed information of the HyperFlex cluster. * &#x60;Unknown&#x60; - The uplink speed could not be determined. The physical servers are potentially not claimed. * &#x60;10G&#x60; - The uplink speed is 10G. * &#x60;1G&#x60; - The uplink speed is 1G. | [optional] [readonly] [default to "Unknown"]
**VcenterConfiguration** | Pointer to [**NullableHyperflexVcenterConfiguration**](HyperflexVcenterConfiguration.md) |  | [optional] 
**VmCount** | Pointer to **int64** | The number of virtual machines present on this cluster. | [optional] [readonly] 
**ZoneType** | Pointer to **string** | The type of availability zone used by the cluster. Physical zones are always used in HyperFlex Stretched Clusters. Logical zones may be used if a cluster has Logical Availability Zones (LAZ) enabled. * &#x60;UNKNOWN&#x60; - The type of zone configured on the HyperFlex cluster is not known. * &#x60;NOT_CONFIGURED&#x60; - The zone type is not configured. * &#x60;LOGICAL&#x60; - The zone is a logical zone created when the logical availability zones (LAZ) feature is enabled on the HyperFlex cluster. * &#x60;PHYSICAL&#x60; - The zone is a physical zone configured on a stretched HyperFlex cluster. | [optional] [readonly] [default to "UNKNOWN"]
**Alarm** | Pointer to [**[]HyperflexAlarmRelationship**](HyperflexAlarmRelationship.md) | An array of relationships to hyperflexAlarm resources. | [optional] [readonly] 
**Encryption** | Pointer to [**HyperflexEncryptionRelationship**](HyperflexEncryptionRelationship.md) |  | [optional] 
**Health** | Pointer to [**HyperflexHealthRelationship**](HyperflexHealthRelationship.md) |  | [optional] 
**License** | Pointer to [**HyperflexLicenseRelationship**](HyperflexLicenseRelationship.md) |  | [optional] 
**Nodes** | Pointer to [**[]HyperflexNodeRelationship**](HyperflexNodeRelationship.md) | An array of relationships to hyperflexNode resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageClientIpPools** | Pointer to [**[]IppoolPoolRelationship**](IppoolPoolRelationship.md) | An array of relationships to ippoolPool resources. | [optional] [readonly] 
**StorageClientVrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 
**StorageContainers** | Pointer to [**[]StorageHyperFlexStorageContainerRelationship**](StorageHyperFlexStorageContainerRelationship.md) | An array of relationships to storageHyperFlexStorageContainer resources. | [optional] [readonly] 
**Volumes** | Pointer to [**[]StorageHyperFlexVolumeRelationship**](StorageHyperFlexVolumeRelationship.md) | An array of relationships to storageHyperFlexVolume resources. | [optional] [readonly] 

## Methods

### NewHyperflexCluster

`func NewHyperflexCluster(classId string, objectType string, ) *HyperflexCluster`

NewHyperflexCluster instantiates a new HyperflexCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterWithDefaults

`func NewHyperflexClusterWithDefaults() *HyperflexCluster`

NewHyperflexClusterWithDefaults instantiates a new HyperflexCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapability

`func (o *HyperflexCluster) GetCapability() HyperflexCapability`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *HyperflexCluster) GetCapabilityOk() (*HyperflexCapability, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *HyperflexCluster) SetCapability(v HyperflexCapability)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *HyperflexCluster) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### SetCapabilityNil

`func (o *HyperflexCluster) SetCapabilityNil(b bool)`

 SetCapabilityNil sets the value for Capability to be an explicit nil

### UnsetCapability
`func (o *HyperflexCluster) UnsetCapability()`

UnsetCapability ensures that no value is present for Capability, not even an explicit nil
### GetClusterType

`func (o *HyperflexCluster) GetClusterType() int64`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *HyperflexCluster) GetClusterTypeOk() (*int64, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *HyperflexCluster) SetClusterType(v int64)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *HyperflexCluster) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetClusterUuid

`func (o *HyperflexCluster) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexCluster) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexCluster) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexCluster) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetDeviceId

`func (o *HyperflexCluster) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HyperflexCluster) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HyperflexCluster) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HyperflexCluster) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDnsServers

`func (o *HyperflexCluster) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *HyperflexCluster) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *HyperflexCluster) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *HyperflexCluster) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *HyperflexCluster) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *HyperflexCluster) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetEncryptionStatus

`func (o *HyperflexCluster) GetEncryptionStatus() string`

GetEncryptionStatus returns the EncryptionStatus field if non-nil, zero value otherwise.

### GetEncryptionStatusOk

`func (o *HyperflexCluster) GetEncryptionStatusOk() (*string, bool)`

GetEncryptionStatusOk returns a tuple with the EncryptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionStatus

`func (o *HyperflexCluster) SetEncryptionStatus(v string)`

SetEncryptionStatus sets EncryptionStatus field to given value.

### HasEncryptionStatus

`func (o *HyperflexCluster) HasEncryptionStatus() bool`

HasEncryptionStatus returns a boolean if a field has been set.

### GetFltAggr

`func (o *HyperflexCluster) GetFltAggr() int64`

GetFltAggr returns the FltAggr field if non-nil, zero value otherwise.

### GetFltAggrOk

`func (o *HyperflexCluster) GetFltAggrOk() (*int64, bool)`

GetFltAggrOk returns a tuple with the FltAggr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFltAggr

`func (o *HyperflexCluster) SetFltAggr(v int64)`

SetFltAggr sets FltAggr field to given value.

### HasFltAggr

`func (o *HyperflexCluster) HasFltAggr() bool`

HasFltAggr returns a boolean if a field has been set.

### GetHxdpBuildVersion

`func (o *HyperflexCluster) GetHxdpBuildVersion() string`

GetHxdpBuildVersion returns the HxdpBuildVersion field if non-nil, zero value otherwise.

### GetHxdpBuildVersionOk

`func (o *HyperflexCluster) GetHxdpBuildVersionOk() (*string, bool)`

GetHxdpBuildVersionOk returns a tuple with the HxdpBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpBuildVersion

`func (o *HyperflexCluster) SetHxdpBuildVersion(v string)`

SetHxdpBuildVersion sets HxdpBuildVersion field to given value.

### HasHxdpBuildVersion

`func (o *HyperflexCluster) HasHxdpBuildVersion() bool`

HasHxdpBuildVersion returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *HyperflexCluster) GetNetworkConfiguration() HyperflexNetworkConfiguration`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *HyperflexCluster) GetNetworkConfigurationOk() (*HyperflexNetworkConfiguration, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *HyperflexCluster) SetNetworkConfiguration(v HyperflexNetworkConfiguration)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *HyperflexCluster) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.

### SetNetworkConfigurationNil

`func (o *HyperflexCluster) SetNetworkConfigurationNil(b bool)`

 SetNetworkConfigurationNil sets the value for NetworkConfiguration to be an explicit nil

### UnsetNetworkConfiguration
`func (o *HyperflexCluster) UnsetNetworkConfiguration()`

UnsetNetworkConfiguration ensures that no value is present for NetworkConfiguration, not even an explicit nil
### GetNtpServers

`func (o *HyperflexCluster) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *HyperflexCluster) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *HyperflexCluster) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *HyperflexCluster) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *HyperflexCluster) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *HyperflexCluster) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetSummary

`func (o *HyperflexCluster) GetSummary() HyperflexSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *HyperflexCluster) GetSummaryOk() (*HyperflexSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *HyperflexCluster) SetSummary(v HyperflexSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *HyperflexCluster) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *HyperflexCluster) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *HyperflexCluster) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetUpgradeStatus

`func (o *HyperflexCluster) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *HyperflexCluster) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *HyperflexCluster) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *HyperflexCluster) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### GetUplinkSpeed

`func (o *HyperflexCluster) GetUplinkSpeed() string`

GetUplinkSpeed returns the UplinkSpeed field if non-nil, zero value otherwise.

### GetUplinkSpeedOk

`func (o *HyperflexCluster) GetUplinkSpeedOk() (*string, bool)`

GetUplinkSpeedOk returns a tuple with the UplinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkSpeed

`func (o *HyperflexCluster) SetUplinkSpeed(v string)`

SetUplinkSpeed sets UplinkSpeed field to given value.

### HasUplinkSpeed

`func (o *HyperflexCluster) HasUplinkSpeed() bool`

HasUplinkSpeed returns a boolean if a field has been set.

### GetVcenterConfiguration

`func (o *HyperflexCluster) GetVcenterConfiguration() HyperflexVcenterConfiguration`

GetVcenterConfiguration returns the VcenterConfiguration field if non-nil, zero value otherwise.

### GetVcenterConfigurationOk

`func (o *HyperflexCluster) GetVcenterConfigurationOk() (*HyperflexVcenterConfiguration, bool)`

GetVcenterConfigurationOk returns a tuple with the VcenterConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterConfiguration

`func (o *HyperflexCluster) SetVcenterConfiguration(v HyperflexVcenterConfiguration)`

SetVcenterConfiguration sets VcenterConfiguration field to given value.

### HasVcenterConfiguration

`func (o *HyperflexCluster) HasVcenterConfiguration() bool`

HasVcenterConfiguration returns a boolean if a field has been set.

### SetVcenterConfigurationNil

`func (o *HyperflexCluster) SetVcenterConfigurationNil(b bool)`

 SetVcenterConfigurationNil sets the value for VcenterConfiguration to be an explicit nil

### UnsetVcenterConfiguration
`func (o *HyperflexCluster) UnsetVcenterConfiguration()`

UnsetVcenterConfiguration ensures that no value is present for VcenterConfiguration, not even an explicit nil
### GetVmCount

`func (o *HyperflexCluster) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *HyperflexCluster) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *HyperflexCluster) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *HyperflexCluster) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetZoneType

`func (o *HyperflexCluster) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *HyperflexCluster) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *HyperflexCluster) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *HyperflexCluster) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetAlarm

`func (o *HyperflexCluster) GetAlarm() []HyperflexAlarmRelationship`

GetAlarm returns the Alarm field if non-nil, zero value otherwise.

### GetAlarmOk

`func (o *HyperflexCluster) GetAlarmOk() (*[]HyperflexAlarmRelationship, bool)`

GetAlarmOk returns a tuple with the Alarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarm

`func (o *HyperflexCluster) SetAlarm(v []HyperflexAlarmRelationship)`

SetAlarm sets Alarm field to given value.

### HasAlarm

`func (o *HyperflexCluster) HasAlarm() bool`

HasAlarm returns a boolean if a field has been set.

### SetAlarmNil

`func (o *HyperflexCluster) SetAlarmNil(b bool)`

 SetAlarmNil sets the value for Alarm to be an explicit nil

### UnsetAlarm
`func (o *HyperflexCluster) UnsetAlarm()`

UnsetAlarm ensures that no value is present for Alarm, not even an explicit nil
### GetEncryption

`func (o *HyperflexCluster) GetEncryption() HyperflexEncryptionRelationship`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *HyperflexCluster) GetEncryptionOk() (*HyperflexEncryptionRelationship, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *HyperflexCluster) SetEncryption(v HyperflexEncryptionRelationship)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *HyperflexCluster) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetHealth

`func (o *HyperflexCluster) GetHealth() HyperflexHealthRelationship`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HyperflexCluster) GetHealthOk() (*HyperflexHealthRelationship, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HyperflexCluster) SetHealth(v HyperflexHealthRelationship)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *HyperflexCluster) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLicense

`func (o *HyperflexCluster) GetLicense() HyperflexLicenseRelationship`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *HyperflexCluster) GetLicenseOk() (*HyperflexLicenseRelationship, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *HyperflexCluster) SetLicense(v HyperflexLicenseRelationship)`

SetLicense sets License field to given value.

### HasLicense

`func (o *HyperflexCluster) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetNodes

`func (o *HyperflexCluster) GetNodes() []HyperflexNodeRelationship`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *HyperflexCluster) GetNodesOk() (*[]HyperflexNodeRelationship, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *HyperflexCluster) SetNodes(v []HyperflexNodeRelationship)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *HyperflexCluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *HyperflexCluster) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *HyperflexCluster) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetRegisteredDevice

`func (o *HyperflexCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageClientIpPools

`func (o *HyperflexCluster) GetStorageClientIpPools() []IppoolPoolRelationship`

GetStorageClientIpPools returns the StorageClientIpPools field if non-nil, zero value otherwise.

### GetStorageClientIpPoolsOk

`func (o *HyperflexCluster) GetStorageClientIpPoolsOk() (*[]IppoolPoolRelationship, bool)`

GetStorageClientIpPoolsOk returns a tuple with the StorageClientIpPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClientIpPools

`func (o *HyperflexCluster) SetStorageClientIpPools(v []IppoolPoolRelationship)`

SetStorageClientIpPools sets StorageClientIpPools field to given value.

### HasStorageClientIpPools

`func (o *HyperflexCluster) HasStorageClientIpPools() bool`

HasStorageClientIpPools returns a boolean if a field has been set.

### SetStorageClientIpPoolsNil

`func (o *HyperflexCluster) SetStorageClientIpPoolsNil(b bool)`

 SetStorageClientIpPoolsNil sets the value for StorageClientIpPools to be an explicit nil

### UnsetStorageClientIpPools
`func (o *HyperflexCluster) UnsetStorageClientIpPools()`

UnsetStorageClientIpPools ensures that no value is present for StorageClientIpPools, not even an explicit nil
### GetStorageClientVrf

`func (o *HyperflexCluster) GetStorageClientVrf() VrfVrfRelationship`

GetStorageClientVrf returns the StorageClientVrf field if non-nil, zero value otherwise.

### GetStorageClientVrfOk

`func (o *HyperflexCluster) GetStorageClientVrfOk() (*VrfVrfRelationship, bool)`

GetStorageClientVrfOk returns a tuple with the StorageClientVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClientVrf

`func (o *HyperflexCluster) SetStorageClientVrf(v VrfVrfRelationship)`

SetStorageClientVrf sets StorageClientVrf field to given value.

### HasStorageClientVrf

`func (o *HyperflexCluster) HasStorageClientVrf() bool`

HasStorageClientVrf returns a boolean if a field has been set.

### GetStorageContainers

`func (o *HyperflexCluster) GetStorageContainers() []StorageHyperFlexStorageContainerRelationship`

GetStorageContainers returns the StorageContainers field if non-nil, zero value otherwise.

### GetStorageContainersOk

`func (o *HyperflexCluster) GetStorageContainersOk() (*[]StorageHyperFlexStorageContainerRelationship, bool)`

GetStorageContainersOk returns a tuple with the StorageContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainers

`func (o *HyperflexCluster) SetStorageContainers(v []StorageHyperFlexStorageContainerRelationship)`

SetStorageContainers sets StorageContainers field to given value.

### HasStorageContainers

`func (o *HyperflexCluster) HasStorageContainers() bool`

HasStorageContainers returns a boolean if a field has been set.

### SetStorageContainersNil

`func (o *HyperflexCluster) SetStorageContainersNil(b bool)`

 SetStorageContainersNil sets the value for StorageContainers to be an explicit nil

### UnsetStorageContainers
`func (o *HyperflexCluster) UnsetStorageContainers()`

UnsetStorageContainers ensures that no value is present for StorageContainers, not even an explicit nil
### GetVolumes

`func (o *HyperflexCluster) GetVolumes() []StorageHyperFlexVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *HyperflexCluster) GetVolumesOk() (*[]StorageHyperFlexVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *HyperflexCluster) SetVolumes(v []StorageHyperFlexVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *HyperflexCluster) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *HyperflexCluster) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *HyperflexCluster) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


