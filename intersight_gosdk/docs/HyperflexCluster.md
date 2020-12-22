# HyperflexCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Cluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Cluster"]
**AlarmSummary** | Pointer to [**NullableHyperflexAlarmSummary**](hyperflex.AlarmSummary.md) |  | [optional] 
**CapacityRunway** | Pointer to **int64** | The number of days remaining before the cluster&#39;s storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \&quot;Unknown\&quot; for a cluster that is not connected or with not sufficient data. | [optional] [readonly] [default to 2147483647]
**ClusterName** | Pointer to **string** | The name of this HyperFlex cluster. | [optional] [readonly] 
**ClusterType** | Pointer to **int64** | The storage type of this cluster (All Flash or Hybrid). | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | The unique identifier for this HyperFlex cluster. | [optional] [readonly] 
**ComputeNodeCount** | Pointer to **int64** | The number of compute nodes that belong to this cluster. | [optional] [readonly] 
**ConvergedNodeCount** | Pointer to **int64** | The number of converged nodes that belong to this cluster. | [optional] [readonly] 
**DeploymentType** | Pointer to **string** | The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as &#39;NA&#39; (not available). * &#x60;NA&#x60; - The deployment type of the HyperFlex cluster is not available. * &#x60;Datacenter&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * &#x60;Stretched Cluster&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * &#x60;Edge&#x60; - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes. | [optional] [readonly] [default to "NA"]
**DeviceId** | Pointer to **string** | The unique identifier of the device registration that represents this HyperFlex cluster&#39;s connection to Intersight. | [optional] [readonly] 
**FltAggr** | Pointer to **int64** | The number of yellow (warning) and red (critical) alarms stored as an aggregate. The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms. | [optional] [readonly] 
**HxVersion** | Pointer to **string** | The HyperFlex Data Platform version of this cluster. | [optional] [readonly] 
**HxdpBuildVersion** | Pointer to **string** | The version and build number of the HyperFlex Data Platform for this cluster. After a cluster upgrade, this version string will be updated on the next inventory cycle to reflect the newly installed version. | [optional] [readonly] 
**HypervisorType** | Pointer to **string** | The type of hypervisor running on this cluster. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [readonly] [default to "ESXi"]
**HypervisorVersion** | Pointer to **string** | The version of hypervisor running on this cluster. | [optional] [readonly] 
**Summary** | Pointer to [**NullableHyperflexSummary**](hyperflex.Summary.md) |  | [optional] 
**UtilizationPercentage** | Pointer to **float32** | The storage utilization percentage is computed based on total capacity and current capacity utilization. | [optional] [readonly] 
**UtilizationTrendPercentage** | Pointer to **float32** | The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | The number of virtual machines present on this cluster. | [optional] [readonly] 
**Alarm** | Pointer to [**[]HyperflexAlarmRelationship**](HyperflexAlarmRelationship.md) | An array of relationships to hyperflexAlarm resources. | [optional] [readonly] 
**Health** | Pointer to [**HyperflexHealthRelationship**](hyperflex.Health.Relationship.md) |  | [optional] 
**Nodes** | Pointer to [**[]HyperflexNodeRelationship**](HyperflexNodeRelationship.md) | An array of relationships to hyperflexNode resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
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


### GetAlarmSummary

`func (o *HyperflexCluster) GetAlarmSummary() HyperflexAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *HyperflexCluster) GetAlarmSummaryOk() (*HyperflexAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *HyperflexCluster) SetAlarmSummary(v HyperflexAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *HyperflexCluster) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *HyperflexCluster) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *HyperflexCluster) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetCapacityRunway

`func (o *HyperflexCluster) GetCapacityRunway() int64`

GetCapacityRunway returns the CapacityRunway field if non-nil, zero value otherwise.

### GetCapacityRunwayOk

`func (o *HyperflexCluster) GetCapacityRunwayOk() (*int64, bool)`

GetCapacityRunwayOk returns a tuple with the CapacityRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRunway

`func (o *HyperflexCluster) SetCapacityRunway(v int64)`

SetCapacityRunway sets CapacityRunway field to given value.

### HasCapacityRunway

`func (o *HyperflexCluster) HasCapacityRunway() bool`

HasCapacityRunway returns a boolean if a field has been set.

### GetClusterName

`func (o *HyperflexCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HyperflexCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HyperflexCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HyperflexCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

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

### GetComputeNodeCount

`func (o *HyperflexCluster) GetComputeNodeCount() int64`

GetComputeNodeCount returns the ComputeNodeCount field if non-nil, zero value otherwise.

### GetComputeNodeCountOk

`func (o *HyperflexCluster) GetComputeNodeCountOk() (*int64, bool)`

GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodeCount

`func (o *HyperflexCluster) SetComputeNodeCount(v int64)`

SetComputeNodeCount sets ComputeNodeCount field to given value.

### HasComputeNodeCount

`func (o *HyperflexCluster) HasComputeNodeCount() bool`

HasComputeNodeCount returns a boolean if a field has been set.

### GetConvergedNodeCount

`func (o *HyperflexCluster) GetConvergedNodeCount() int64`

GetConvergedNodeCount returns the ConvergedNodeCount field if non-nil, zero value otherwise.

### GetConvergedNodeCountOk

`func (o *HyperflexCluster) GetConvergedNodeCountOk() (*int64, bool)`

GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergedNodeCount

`func (o *HyperflexCluster) SetConvergedNodeCount(v int64)`

SetConvergedNodeCount sets ConvergedNodeCount field to given value.

### HasConvergedNodeCount

`func (o *HyperflexCluster) HasConvergedNodeCount() bool`

HasConvergedNodeCount returns a boolean if a field has been set.

### GetDeploymentType

`func (o *HyperflexCluster) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *HyperflexCluster) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *HyperflexCluster) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *HyperflexCluster) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

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

### GetHxVersion

`func (o *HyperflexCluster) GetHxVersion() string`

GetHxVersion returns the HxVersion field if non-nil, zero value otherwise.

### GetHxVersionOk

`func (o *HyperflexCluster) GetHxVersionOk() (*string, bool)`

GetHxVersionOk returns a tuple with the HxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxVersion

`func (o *HyperflexCluster) SetHxVersion(v string)`

SetHxVersion sets HxVersion field to given value.

### HasHxVersion

`func (o *HyperflexCluster) HasHxVersion() bool`

HasHxVersion returns a boolean if a field has been set.

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

### GetHypervisorType

`func (o *HyperflexCluster) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HyperflexCluster) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HyperflexCluster) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HyperflexCluster) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HyperflexCluster) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HyperflexCluster) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HyperflexCluster) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HyperflexCluster) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

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
### GetUtilizationPercentage

`func (o *HyperflexCluster) GetUtilizationPercentage() float32`

GetUtilizationPercentage returns the UtilizationPercentage field if non-nil, zero value otherwise.

### GetUtilizationPercentageOk

`func (o *HyperflexCluster) GetUtilizationPercentageOk() (*float32, bool)`

GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercentage

`func (o *HyperflexCluster) SetUtilizationPercentage(v float32)`

SetUtilizationPercentage sets UtilizationPercentage field to given value.

### HasUtilizationPercentage

`func (o *HyperflexCluster) HasUtilizationPercentage() bool`

HasUtilizationPercentage returns a boolean if a field has been set.

### GetUtilizationTrendPercentage

`func (o *HyperflexCluster) GetUtilizationTrendPercentage() float32`

GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field if non-nil, zero value otherwise.

### GetUtilizationTrendPercentageOk

`func (o *HyperflexCluster) GetUtilizationTrendPercentageOk() (*float32, bool)`

GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationTrendPercentage

`func (o *HyperflexCluster) SetUtilizationTrendPercentage(v float32)`

SetUtilizationTrendPercentage sets UtilizationTrendPercentage field to given value.

### HasUtilizationTrendPercentage

`func (o *HyperflexCluster) HasUtilizationTrendPercentage() bool`

HasUtilizationTrendPercentage returns a boolean if a field has been set.

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


