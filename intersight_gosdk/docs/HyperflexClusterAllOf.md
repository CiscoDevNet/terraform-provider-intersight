# HyperflexClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Cluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Cluster"]
**ClusterType** | Pointer to **int64** | The storage type of this cluster (All Flash or Hybrid). | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | The unique identifier for this HyperFlex cluster. | [optional] [readonly] 
**DeviceId** | Pointer to **string** | The unique identifier of the device registration that represents this HyperFlex cluster&#39;s connection to Intersight. | [optional] [readonly] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**FltAggr** | Pointer to **int64** | The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms. | [optional] [readonly] 
**HxdpBuildVersion** | Pointer to **string** | The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version. | [optional] [readonly] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to [**NullableHyperflexSummary**](HyperflexSummary.md) |  | [optional] 
**UpgradeStatus** | Pointer to **string** | The upgrade status of the HyperFlex cluster. * &#x60;Unknown&#x60; - The upgrade status of the HyperFlex cluster could not be determined. * &#x60;Ok&#x60; - The upgrade of the HyperFlex cluster is complete. * &#x60;InProgress&#x60; - The upgrade of the HyperFlex cluster is in-progress. * &#x60;Failed&#x60; - The upgrade of the HyperFlex cluster has failed. * &#x60;Waiting&#x60; - The upgrade of the HyperFlex cluster is waiting to continue execution. | [optional] [readonly] [default to "Unknown"]
**VmCount** | Pointer to **int64** | The number of virtual machines present on this cluster. | [optional] [readonly] 
**Alarm** | Pointer to [**[]HyperflexAlarmRelationship**](HyperflexAlarmRelationship.md) | An array of relationships to hyperflexAlarm resources. | [optional] [readonly] 
**Health** | Pointer to [**HyperflexHealthRelationship**](HyperflexHealthRelationship.md) |  | [optional] 
**License** | Pointer to [**HyperflexLicenseRelationship**](HyperflexLicenseRelationship.md) |  | [optional] 
**Nodes** | Pointer to [**[]HyperflexNodeRelationship**](HyperflexNodeRelationship.md) | An array of relationships to hyperflexNode resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageContainers** | Pointer to [**[]StorageHyperFlexStorageContainerRelationship**](StorageHyperFlexStorageContainerRelationship.md) | An array of relationships to storageHyperFlexStorageContainer resources. | [optional] [readonly] 
**Volumes** | Pointer to [**[]StorageHyperFlexVolumeRelationship**](StorageHyperFlexVolumeRelationship.md) | An array of relationships to storageHyperFlexVolume resources. | [optional] [readonly] 

## Methods

### NewHyperflexClusterAllOf

`func NewHyperflexClusterAllOf(classId string, objectType string, ) *HyperflexClusterAllOf`

NewHyperflexClusterAllOf instantiates a new HyperflexClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterAllOfWithDefaults

`func NewHyperflexClusterAllOfWithDefaults() *HyperflexClusterAllOf`

NewHyperflexClusterAllOfWithDefaults instantiates a new HyperflexClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterType

`func (o *HyperflexClusterAllOf) GetClusterType() int64`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *HyperflexClusterAllOf) GetClusterTypeOk() (*int64, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *HyperflexClusterAllOf) SetClusterType(v int64)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *HyperflexClusterAllOf) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetClusterUuid

`func (o *HyperflexClusterAllOf) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexClusterAllOf) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexClusterAllOf) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexClusterAllOf) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetDeviceId

`func (o *HyperflexClusterAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HyperflexClusterAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HyperflexClusterAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HyperflexClusterAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDnsServers

`func (o *HyperflexClusterAllOf) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *HyperflexClusterAllOf) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *HyperflexClusterAllOf) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *HyperflexClusterAllOf) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *HyperflexClusterAllOf) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *HyperflexClusterAllOf) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetFltAggr

`func (o *HyperflexClusterAllOf) GetFltAggr() int64`

GetFltAggr returns the FltAggr field if non-nil, zero value otherwise.

### GetFltAggrOk

`func (o *HyperflexClusterAllOf) GetFltAggrOk() (*int64, bool)`

GetFltAggrOk returns a tuple with the FltAggr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFltAggr

`func (o *HyperflexClusterAllOf) SetFltAggr(v int64)`

SetFltAggr sets FltAggr field to given value.

### HasFltAggr

`func (o *HyperflexClusterAllOf) HasFltAggr() bool`

HasFltAggr returns a boolean if a field has been set.

### GetHxdpBuildVersion

`func (o *HyperflexClusterAllOf) GetHxdpBuildVersion() string`

GetHxdpBuildVersion returns the HxdpBuildVersion field if non-nil, zero value otherwise.

### GetHxdpBuildVersionOk

`func (o *HyperflexClusterAllOf) GetHxdpBuildVersionOk() (*string, bool)`

GetHxdpBuildVersionOk returns a tuple with the HxdpBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpBuildVersion

`func (o *HyperflexClusterAllOf) SetHxdpBuildVersion(v string)`

SetHxdpBuildVersion sets HxdpBuildVersion field to given value.

### HasHxdpBuildVersion

`func (o *HyperflexClusterAllOf) HasHxdpBuildVersion() bool`

HasHxdpBuildVersion returns a boolean if a field has been set.

### GetNtpServers

`func (o *HyperflexClusterAllOf) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *HyperflexClusterAllOf) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *HyperflexClusterAllOf) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *HyperflexClusterAllOf) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *HyperflexClusterAllOf) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *HyperflexClusterAllOf) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetSummary

`func (o *HyperflexClusterAllOf) GetSummary() HyperflexSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *HyperflexClusterAllOf) GetSummaryOk() (*HyperflexSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *HyperflexClusterAllOf) SetSummary(v HyperflexSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *HyperflexClusterAllOf) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *HyperflexClusterAllOf) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *HyperflexClusterAllOf) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetUpgradeStatus

`func (o *HyperflexClusterAllOf) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *HyperflexClusterAllOf) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *HyperflexClusterAllOf) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *HyperflexClusterAllOf) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### GetVmCount

`func (o *HyperflexClusterAllOf) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *HyperflexClusterAllOf) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *HyperflexClusterAllOf) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *HyperflexClusterAllOf) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetAlarm

`func (o *HyperflexClusterAllOf) GetAlarm() []HyperflexAlarmRelationship`

GetAlarm returns the Alarm field if non-nil, zero value otherwise.

### GetAlarmOk

`func (o *HyperflexClusterAllOf) GetAlarmOk() (*[]HyperflexAlarmRelationship, bool)`

GetAlarmOk returns a tuple with the Alarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarm

`func (o *HyperflexClusterAllOf) SetAlarm(v []HyperflexAlarmRelationship)`

SetAlarm sets Alarm field to given value.

### HasAlarm

`func (o *HyperflexClusterAllOf) HasAlarm() bool`

HasAlarm returns a boolean if a field has been set.

### SetAlarmNil

`func (o *HyperflexClusterAllOf) SetAlarmNil(b bool)`

 SetAlarmNil sets the value for Alarm to be an explicit nil

### UnsetAlarm
`func (o *HyperflexClusterAllOf) UnsetAlarm()`

UnsetAlarm ensures that no value is present for Alarm, not even an explicit nil
### GetHealth

`func (o *HyperflexClusterAllOf) GetHealth() HyperflexHealthRelationship`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HyperflexClusterAllOf) GetHealthOk() (*HyperflexHealthRelationship, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HyperflexClusterAllOf) SetHealth(v HyperflexHealthRelationship)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *HyperflexClusterAllOf) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLicense

`func (o *HyperflexClusterAllOf) GetLicense() HyperflexLicenseRelationship`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *HyperflexClusterAllOf) GetLicenseOk() (*HyperflexLicenseRelationship, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *HyperflexClusterAllOf) SetLicense(v HyperflexLicenseRelationship)`

SetLicense sets License field to given value.

### HasLicense

`func (o *HyperflexClusterAllOf) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetNodes

`func (o *HyperflexClusterAllOf) GetNodes() []HyperflexNodeRelationship`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *HyperflexClusterAllOf) GetNodesOk() (*[]HyperflexNodeRelationship, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *HyperflexClusterAllOf) SetNodes(v []HyperflexNodeRelationship)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *HyperflexClusterAllOf) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *HyperflexClusterAllOf) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *HyperflexClusterAllOf) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetRegisteredDevice

`func (o *HyperflexClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexClusterAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageContainers

`func (o *HyperflexClusterAllOf) GetStorageContainers() []StorageHyperFlexStorageContainerRelationship`

GetStorageContainers returns the StorageContainers field if non-nil, zero value otherwise.

### GetStorageContainersOk

`func (o *HyperflexClusterAllOf) GetStorageContainersOk() (*[]StorageHyperFlexStorageContainerRelationship, bool)`

GetStorageContainersOk returns a tuple with the StorageContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainers

`func (o *HyperflexClusterAllOf) SetStorageContainers(v []StorageHyperFlexStorageContainerRelationship)`

SetStorageContainers sets StorageContainers field to given value.

### HasStorageContainers

`func (o *HyperflexClusterAllOf) HasStorageContainers() bool`

HasStorageContainers returns a boolean if a field has been set.

### SetStorageContainersNil

`func (o *HyperflexClusterAllOf) SetStorageContainersNil(b bool)`

 SetStorageContainersNil sets the value for StorageContainers to be an explicit nil

### UnsetStorageContainers
`func (o *HyperflexClusterAllOf) UnsetStorageContainers()`

UnsetStorageContainers ensures that no value is present for StorageContainers, not even an explicit nil
### GetVolumes

`func (o *HyperflexClusterAllOf) GetVolumes() []StorageHyperFlexVolumeRelationship`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *HyperflexClusterAllOf) GetVolumesOk() (*[]StorageHyperFlexVolumeRelationship, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *HyperflexClusterAllOf) SetVolumes(v []StorageHyperFlexVolumeRelationship)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *HyperflexClusterAllOf) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *HyperflexClusterAllOf) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *HyperflexClusterAllOf) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


