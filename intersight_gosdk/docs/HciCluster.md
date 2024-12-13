# HciCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Cluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Cluster"]
**AlarmSummary** | Pointer to [**NullableHciAlarmSummary**](HciAlarmSummary.md) |  | [optional] 
**Backplane** | Pointer to [**NullableHciBackplaneNetworkParams**](HciBackplaneNetworkParams.md) |  | [optional] 
**BuildInfoBuildType** | Pointer to **string** | The software build type, such as \&quot;release\&quot; or \&quot;debug\&quot; build. | [optional] [readonly] 
**BuildInfoCommitId** | Pointer to **string** | The software commit id for this build image. | [optional] [readonly] 
**BuildInfoFullVersion** | Pointer to **string** | The longer form of software version. It usually includes the commit id. | [optional] [readonly] 
**BuildInfoShortCommitId** | Pointer to **string** | The short version of the software commit id for this build image. | [optional] [readonly] 
**BuildInfoVersion** | Pointer to **string** | The software version from the build. | [optional] [readonly] 
**ClusterArch** | Pointer to **string** | The CPU architecture of the cluster server such as x86_64 and PPC64LE. | [optional] [readonly] 
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster. | [optional] [readonly] 
**ClusterFunction** | Pointer to **[]string** |  | [optional] 
**ClusterSoftwareMap** | Pointer to [**[]HciSoftwareType**](HciSoftwareType.md) |  | [optional] 
**ContainerName** | Pointer to **string** | The name of the default container created as part of cluster creation. | [optional] [readonly] 
**CpuCapacityHz** | Pointer to **int64** | The CPU capacity in Hz of the cluster. | [optional] [readonly] 
**CpuUsageHz** | Pointer to **int64** | The CPU usage in Hz of the cluster. | [optional] [readonly] 
**EncryptionInTransitStatus** | Pointer to **bool** | Indicate if encryption-in-transit is enabled or not. | [optional] [readonly] 
**EncryptionScope** | Pointer to **[]string** |  | [optional] 
**ExternalAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**ExternalDataServiceIp** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**ExternalSubnet** | Pointer to **string** | The external subnet of the cluster. | [optional] [readonly] 
**FaultToleranceState** | Pointer to [**NullableHciFaultToleranceState**](HciFaultToleranceState.md) |  | [optional] 
**HypervisorTypes** | Pointer to **[]string** |  | [optional] 
**Incarnationid** | Pointer to **string** | Cluster incarnation Id, part of payload for cluster update operation only. | [optional] [readonly] 
**InefficientVmCount** | Pointer to **int64** | The number of inefficient VMs in this cluster. | [optional] [readonly] 
**InternalSubnet** | Pointer to **string** | The internal subnet of the cluster. | [optional] [readonly] 
**IsLts** | Pointer to **bool** | The LTS status indicates whether the release is categorized as Long-term or not. | [optional] [readonly] 
**KeyManagementServerType** | Pointer to **string** | The key management server type of the cluster. | [optional] [readonly] 
**ManagementServer** | Pointer to [**NullableHciManagementServer**](HciManagementServer.md) |  | [optional] 
**MasqueradingIp** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**MasqueradingPort** | Pointer to **int32** | The masquerading port of the cluster. | [optional] [readonly] 
**MemoryCapacityBytes** | Pointer to **int64** | The memory capacity in bytes of the cluster. | [optional] [readonly] 
**MemoryUsageBytes** | Pointer to **int64** | The memory usage in bytes of the cluster. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the reported cluster. | [optional] [readonly] 
**NameServerIpList** | Pointer to [**[]HciIpAddressOrFqdn**](HciIpAddressOrFqdn.md) |  | [optional] 
**NtpServerIpList** | Pointer to [**[]HciIpAddressOrFqdn**](HciIpAddressOrFqdn.md) |  | [optional] 
**NumberOfNodes** | Pointer to **int32** | The number of nodes in the cluster. | [optional] [readonly] 
**OperationMode** | Pointer to **string** | The operation mode of the cluster such as NORMAL, READ_ONLY, STAND_ALONE, SWITCH_TO_TWO_NODE, OVERRIDE. | [optional] [readonly] 
**PasswordRemoteLoginEnabled** | Pointer to **bool** | Indicates whether the password ssh into the cluster is enabled or not. | [optional] [readonly] 
**PcExtId** | Pointer to **string** | The unique identifier of the domain manager (Prism Central) instance which manages this cluster. | [optional] [readonly] 
**PulseStatus** | Pointer to [**NullableHciPulseStatus**](HciPulseStatus.md) |  | [optional] 
**RedundancyFactor** | Pointer to **int64** | The redundancy factor of the cluster. | [optional] [readonly] 
**RemoteSupport** | Pointer to **bool** | The remote support status of the cluster. | [optional] [readonly] 
**StorageCapacityBytes** | Pointer to **int64** | The storage capacity in bytes of the cluster. | [optional] [readonly] 
**StorageUsageBytes** | Pointer to **int64** | The storage usage in bytes of the cluster. | [optional] [readonly] 
**Timezone** | Pointer to **string** | The timezone of the cluster. | [optional] [readonly] 
**UpgradeStatus** | Pointer to **string** | The upgrade status of a cluster includes the following known values: PENDING, DOWNLOADING, QUEUED, PREUPGRADE, UPGRADING, SUCCEEDED, FAILED, CANCELLED, and SCHEDULED.The upgrade status of a cluster. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | The number of VMs running on this cluster. | [optional] [readonly] 
**Compliance** | Pointer to [**NullableHciComplianceRelationship**](HciComplianceRelationship.md) |  | [optional] 
**DomainManager** | Pointer to [**NullableHciDomainManagerRelationship**](HciDomainManagerRelationship.md) |  | [optional] 
**Entitlement** | Pointer to [**NullableHciEntitlementRelationship**](HciEntitlementRelationship.md) |  | [optional] 
**Nodes** | Pointer to [**[]HciNodeRelationship**](HciNodeRelationship.md) | An array of relationships to hciNode resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Violation** | Pointer to [**NullableHciViolationRelationship**](HciViolationRelationship.md) |  | [optional] 

## Methods

### NewHciCluster

`func NewHciCluster(classId string, objectType string, ) *HciCluster`

NewHciCluster instantiates a new HciCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciClusterWithDefaults

`func NewHciClusterWithDefaults() *HciCluster`

NewHciClusterWithDefaults instantiates a new HciCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmSummary

`func (o *HciCluster) GetAlarmSummary() HciAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *HciCluster) GetAlarmSummaryOk() (*HciAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *HciCluster) SetAlarmSummary(v HciAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *HciCluster) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *HciCluster) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *HciCluster) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetBackplane

`func (o *HciCluster) GetBackplane() HciBackplaneNetworkParams`

GetBackplane returns the Backplane field if non-nil, zero value otherwise.

### GetBackplaneOk

`func (o *HciCluster) GetBackplaneOk() (*HciBackplaneNetworkParams, bool)`

GetBackplaneOk returns a tuple with the Backplane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackplane

`func (o *HciCluster) SetBackplane(v HciBackplaneNetworkParams)`

SetBackplane sets Backplane field to given value.

### HasBackplane

`func (o *HciCluster) HasBackplane() bool`

HasBackplane returns a boolean if a field has been set.

### SetBackplaneNil

`func (o *HciCluster) SetBackplaneNil(b bool)`

 SetBackplaneNil sets the value for Backplane to be an explicit nil

### UnsetBackplane
`func (o *HciCluster) UnsetBackplane()`

UnsetBackplane ensures that no value is present for Backplane, not even an explicit nil
### GetBuildInfoBuildType

`func (o *HciCluster) GetBuildInfoBuildType() string`

GetBuildInfoBuildType returns the BuildInfoBuildType field if non-nil, zero value otherwise.

### GetBuildInfoBuildTypeOk

`func (o *HciCluster) GetBuildInfoBuildTypeOk() (*string, bool)`

GetBuildInfoBuildTypeOk returns a tuple with the BuildInfoBuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfoBuildType

`func (o *HciCluster) SetBuildInfoBuildType(v string)`

SetBuildInfoBuildType sets BuildInfoBuildType field to given value.

### HasBuildInfoBuildType

`func (o *HciCluster) HasBuildInfoBuildType() bool`

HasBuildInfoBuildType returns a boolean if a field has been set.

### GetBuildInfoCommitId

`func (o *HciCluster) GetBuildInfoCommitId() string`

GetBuildInfoCommitId returns the BuildInfoCommitId field if non-nil, zero value otherwise.

### GetBuildInfoCommitIdOk

`func (o *HciCluster) GetBuildInfoCommitIdOk() (*string, bool)`

GetBuildInfoCommitIdOk returns a tuple with the BuildInfoCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfoCommitId

`func (o *HciCluster) SetBuildInfoCommitId(v string)`

SetBuildInfoCommitId sets BuildInfoCommitId field to given value.

### HasBuildInfoCommitId

`func (o *HciCluster) HasBuildInfoCommitId() bool`

HasBuildInfoCommitId returns a boolean if a field has been set.

### GetBuildInfoFullVersion

`func (o *HciCluster) GetBuildInfoFullVersion() string`

GetBuildInfoFullVersion returns the BuildInfoFullVersion field if non-nil, zero value otherwise.

### GetBuildInfoFullVersionOk

`func (o *HciCluster) GetBuildInfoFullVersionOk() (*string, bool)`

GetBuildInfoFullVersionOk returns a tuple with the BuildInfoFullVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfoFullVersion

`func (o *HciCluster) SetBuildInfoFullVersion(v string)`

SetBuildInfoFullVersion sets BuildInfoFullVersion field to given value.

### HasBuildInfoFullVersion

`func (o *HciCluster) HasBuildInfoFullVersion() bool`

HasBuildInfoFullVersion returns a boolean if a field has been set.

### GetBuildInfoShortCommitId

`func (o *HciCluster) GetBuildInfoShortCommitId() string`

GetBuildInfoShortCommitId returns the BuildInfoShortCommitId field if non-nil, zero value otherwise.

### GetBuildInfoShortCommitIdOk

`func (o *HciCluster) GetBuildInfoShortCommitIdOk() (*string, bool)`

GetBuildInfoShortCommitIdOk returns a tuple with the BuildInfoShortCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfoShortCommitId

`func (o *HciCluster) SetBuildInfoShortCommitId(v string)`

SetBuildInfoShortCommitId sets BuildInfoShortCommitId field to given value.

### HasBuildInfoShortCommitId

`func (o *HciCluster) HasBuildInfoShortCommitId() bool`

HasBuildInfoShortCommitId returns a boolean if a field has been set.

### GetBuildInfoVersion

`func (o *HciCluster) GetBuildInfoVersion() string`

GetBuildInfoVersion returns the BuildInfoVersion field if non-nil, zero value otherwise.

### GetBuildInfoVersionOk

`func (o *HciCluster) GetBuildInfoVersionOk() (*string, bool)`

GetBuildInfoVersionOk returns a tuple with the BuildInfoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfoVersion

`func (o *HciCluster) SetBuildInfoVersion(v string)`

SetBuildInfoVersion sets BuildInfoVersion field to given value.

### HasBuildInfoVersion

`func (o *HciCluster) HasBuildInfoVersion() bool`

HasBuildInfoVersion returns a boolean if a field has been set.

### GetClusterArch

`func (o *HciCluster) GetClusterArch() string`

GetClusterArch returns the ClusterArch field if non-nil, zero value otherwise.

### GetClusterArchOk

`func (o *HciCluster) GetClusterArchOk() (*string, bool)`

GetClusterArchOk returns a tuple with the ClusterArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArch

`func (o *HciCluster) SetClusterArch(v string)`

SetClusterArch sets ClusterArch field to given value.

### HasClusterArch

`func (o *HciCluster) HasClusterArch() bool`

HasClusterArch returns a boolean if a field has been set.

### GetClusterExtId

`func (o *HciCluster) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciCluster) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciCluster) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciCluster) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetClusterFunction

`func (o *HciCluster) GetClusterFunction() []string`

GetClusterFunction returns the ClusterFunction field if non-nil, zero value otherwise.

### GetClusterFunctionOk

`func (o *HciCluster) GetClusterFunctionOk() (*[]string, bool)`

GetClusterFunctionOk returns a tuple with the ClusterFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFunction

`func (o *HciCluster) SetClusterFunction(v []string)`

SetClusterFunction sets ClusterFunction field to given value.

### HasClusterFunction

`func (o *HciCluster) HasClusterFunction() bool`

HasClusterFunction returns a boolean if a field has been set.

### SetClusterFunctionNil

`func (o *HciCluster) SetClusterFunctionNil(b bool)`

 SetClusterFunctionNil sets the value for ClusterFunction to be an explicit nil

### UnsetClusterFunction
`func (o *HciCluster) UnsetClusterFunction()`

UnsetClusterFunction ensures that no value is present for ClusterFunction, not even an explicit nil
### GetClusterSoftwareMap

`func (o *HciCluster) GetClusterSoftwareMap() []HciSoftwareType`

GetClusterSoftwareMap returns the ClusterSoftwareMap field if non-nil, zero value otherwise.

### GetClusterSoftwareMapOk

`func (o *HciCluster) GetClusterSoftwareMapOk() (*[]HciSoftwareType, bool)`

GetClusterSoftwareMapOk returns a tuple with the ClusterSoftwareMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSoftwareMap

`func (o *HciCluster) SetClusterSoftwareMap(v []HciSoftwareType)`

SetClusterSoftwareMap sets ClusterSoftwareMap field to given value.

### HasClusterSoftwareMap

`func (o *HciCluster) HasClusterSoftwareMap() bool`

HasClusterSoftwareMap returns a boolean if a field has been set.

### SetClusterSoftwareMapNil

`func (o *HciCluster) SetClusterSoftwareMapNil(b bool)`

 SetClusterSoftwareMapNil sets the value for ClusterSoftwareMap to be an explicit nil

### UnsetClusterSoftwareMap
`func (o *HciCluster) UnsetClusterSoftwareMap()`

UnsetClusterSoftwareMap ensures that no value is present for ClusterSoftwareMap, not even an explicit nil
### GetContainerName

`func (o *HciCluster) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *HciCluster) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *HciCluster) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *HciCluster) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetCpuCapacityHz

`func (o *HciCluster) GetCpuCapacityHz() int64`

GetCpuCapacityHz returns the CpuCapacityHz field if non-nil, zero value otherwise.

### GetCpuCapacityHzOk

`func (o *HciCluster) GetCpuCapacityHzOk() (*int64, bool)`

GetCpuCapacityHzOk returns a tuple with the CpuCapacityHz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCapacityHz

`func (o *HciCluster) SetCpuCapacityHz(v int64)`

SetCpuCapacityHz sets CpuCapacityHz field to given value.

### HasCpuCapacityHz

`func (o *HciCluster) HasCpuCapacityHz() bool`

HasCpuCapacityHz returns a boolean if a field has been set.

### GetCpuUsageHz

`func (o *HciCluster) GetCpuUsageHz() int64`

GetCpuUsageHz returns the CpuUsageHz field if non-nil, zero value otherwise.

### GetCpuUsageHzOk

`func (o *HciCluster) GetCpuUsageHzOk() (*int64, bool)`

GetCpuUsageHzOk returns a tuple with the CpuUsageHz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageHz

`func (o *HciCluster) SetCpuUsageHz(v int64)`

SetCpuUsageHz sets CpuUsageHz field to given value.

### HasCpuUsageHz

`func (o *HciCluster) HasCpuUsageHz() bool`

HasCpuUsageHz returns a boolean if a field has been set.

### GetEncryptionInTransitStatus

`func (o *HciCluster) GetEncryptionInTransitStatus() bool`

GetEncryptionInTransitStatus returns the EncryptionInTransitStatus field if non-nil, zero value otherwise.

### GetEncryptionInTransitStatusOk

`func (o *HciCluster) GetEncryptionInTransitStatusOk() (*bool, bool)`

GetEncryptionInTransitStatusOk returns a tuple with the EncryptionInTransitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionInTransitStatus

`func (o *HciCluster) SetEncryptionInTransitStatus(v bool)`

SetEncryptionInTransitStatus sets EncryptionInTransitStatus field to given value.

### HasEncryptionInTransitStatus

`func (o *HciCluster) HasEncryptionInTransitStatus() bool`

HasEncryptionInTransitStatus returns a boolean if a field has been set.

### GetEncryptionScope

`func (o *HciCluster) GetEncryptionScope() []string`

GetEncryptionScope returns the EncryptionScope field if non-nil, zero value otherwise.

### GetEncryptionScopeOk

`func (o *HciCluster) GetEncryptionScopeOk() (*[]string, bool)`

GetEncryptionScopeOk returns a tuple with the EncryptionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionScope

`func (o *HciCluster) SetEncryptionScope(v []string)`

SetEncryptionScope sets EncryptionScope field to given value.

### HasEncryptionScope

`func (o *HciCluster) HasEncryptionScope() bool`

HasEncryptionScope returns a boolean if a field has been set.

### SetEncryptionScopeNil

`func (o *HciCluster) SetEncryptionScopeNil(b bool)`

 SetEncryptionScopeNil sets the value for EncryptionScope to be an explicit nil

### UnsetEncryptionScope
`func (o *HciCluster) UnsetEncryptionScope()`

UnsetEncryptionScope ensures that no value is present for EncryptionScope, not even an explicit nil
### GetExternalAddress

`func (o *HciCluster) GetExternalAddress() HciIpAddress`

GetExternalAddress returns the ExternalAddress field if non-nil, zero value otherwise.

### GetExternalAddressOk

`func (o *HciCluster) GetExternalAddressOk() (*HciIpAddress, bool)`

GetExternalAddressOk returns a tuple with the ExternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddress

`func (o *HciCluster) SetExternalAddress(v HciIpAddress)`

SetExternalAddress sets ExternalAddress field to given value.

### HasExternalAddress

`func (o *HciCluster) HasExternalAddress() bool`

HasExternalAddress returns a boolean if a field has been set.

### SetExternalAddressNil

`func (o *HciCluster) SetExternalAddressNil(b bool)`

 SetExternalAddressNil sets the value for ExternalAddress to be an explicit nil

### UnsetExternalAddress
`func (o *HciCluster) UnsetExternalAddress()`

UnsetExternalAddress ensures that no value is present for ExternalAddress, not even an explicit nil
### GetExternalDataServiceIp

`func (o *HciCluster) GetExternalDataServiceIp() HciIpAddress`

GetExternalDataServiceIp returns the ExternalDataServiceIp field if non-nil, zero value otherwise.

### GetExternalDataServiceIpOk

`func (o *HciCluster) GetExternalDataServiceIpOk() (*HciIpAddress, bool)`

GetExternalDataServiceIpOk returns a tuple with the ExternalDataServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDataServiceIp

`func (o *HciCluster) SetExternalDataServiceIp(v HciIpAddress)`

SetExternalDataServiceIp sets ExternalDataServiceIp field to given value.

### HasExternalDataServiceIp

`func (o *HciCluster) HasExternalDataServiceIp() bool`

HasExternalDataServiceIp returns a boolean if a field has been set.

### SetExternalDataServiceIpNil

`func (o *HciCluster) SetExternalDataServiceIpNil(b bool)`

 SetExternalDataServiceIpNil sets the value for ExternalDataServiceIp to be an explicit nil

### UnsetExternalDataServiceIp
`func (o *HciCluster) UnsetExternalDataServiceIp()`

UnsetExternalDataServiceIp ensures that no value is present for ExternalDataServiceIp, not even an explicit nil
### GetExternalSubnet

`func (o *HciCluster) GetExternalSubnet() string`

GetExternalSubnet returns the ExternalSubnet field if non-nil, zero value otherwise.

### GetExternalSubnetOk

`func (o *HciCluster) GetExternalSubnetOk() (*string, bool)`

GetExternalSubnetOk returns a tuple with the ExternalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnet

`func (o *HciCluster) SetExternalSubnet(v string)`

SetExternalSubnet sets ExternalSubnet field to given value.

### HasExternalSubnet

`func (o *HciCluster) HasExternalSubnet() bool`

HasExternalSubnet returns a boolean if a field has been set.

### GetFaultToleranceState

`func (o *HciCluster) GetFaultToleranceState() HciFaultToleranceState`

GetFaultToleranceState returns the FaultToleranceState field if non-nil, zero value otherwise.

### GetFaultToleranceStateOk

`func (o *HciCluster) GetFaultToleranceStateOk() (*HciFaultToleranceState, bool)`

GetFaultToleranceStateOk returns a tuple with the FaultToleranceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultToleranceState

`func (o *HciCluster) SetFaultToleranceState(v HciFaultToleranceState)`

SetFaultToleranceState sets FaultToleranceState field to given value.

### HasFaultToleranceState

`func (o *HciCluster) HasFaultToleranceState() bool`

HasFaultToleranceState returns a boolean if a field has been set.

### SetFaultToleranceStateNil

`func (o *HciCluster) SetFaultToleranceStateNil(b bool)`

 SetFaultToleranceStateNil sets the value for FaultToleranceState to be an explicit nil

### UnsetFaultToleranceState
`func (o *HciCluster) UnsetFaultToleranceState()`

UnsetFaultToleranceState ensures that no value is present for FaultToleranceState, not even an explicit nil
### GetHypervisorTypes

`func (o *HciCluster) GetHypervisorTypes() []string`

GetHypervisorTypes returns the HypervisorTypes field if non-nil, zero value otherwise.

### GetHypervisorTypesOk

`func (o *HciCluster) GetHypervisorTypesOk() (*[]string, bool)`

GetHypervisorTypesOk returns a tuple with the HypervisorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorTypes

`func (o *HciCluster) SetHypervisorTypes(v []string)`

SetHypervisorTypes sets HypervisorTypes field to given value.

### HasHypervisorTypes

`func (o *HciCluster) HasHypervisorTypes() bool`

HasHypervisorTypes returns a boolean if a field has been set.

### SetHypervisorTypesNil

`func (o *HciCluster) SetHypervisorTypesNil(b bool)`

 SetHypervisorTypesNil sets the value for HypervisorTypes to be an explicit nil

### UnsetHypervisorTypes
`func (o *HciCluster) UnsetHypervisorTypes()`

UnsetHypervisorTypes ensures that no value is present for HypervisorTypes, not even an explicit nil
### GetIncarnationid

`func (o *HciCluster) GetIncarnationid() string`

GetIncarnationid returns the Incarnationid field if non-nil, zero value otherwise.

### GetIncarnationidOk

`func (o *HciCluster) GetIncarnationidOk() (*string, bool)`

GetIncarnationidOk returns a tuple with the Incarnationid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationid

`func (o *HciCluster) SetIncarnationid(v string)`

SetIncarnationid sets Incarnationid field to given value.

### HasIncarnationid

`func (o *HciCluster) HasIncarnationid() bool`

HasIncarnationid returns a boolean if a field has been set.

### GetInefficientVmCount

`func (o *HciCluster) GetInefficientVmCount() int64`

GetInefficientVmCount returns the InefficientVmCount field if non-nil, zero value otherwise.

### GetInefficientVmCountOk

`func (o *HciCluster) GetInefficientVmCountOk() (*int64, bool)`

GetInefficientVmCountOk returns a tuple with the InefficientVmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInefficientVmCount

`func (o *HciCluster) SetInefficientVmCount(v int64)`

SetInefficientVmCount sets InefficientVmCount field to given value.

### HasInefficientVmCount

`func (o *HciCluster) HasInefficientVmCount() bool`

HasInefficientVmCount returns a boolean if a field has been set.

### GetInternalSubnet

`func (o *HciCluster) GetInternalSubnet() string`

GetInternalSubnet returns the InternalSubnet field if non-nil, zero value otherwise.

### GetInternalSubnetOk

`func (o *HciCluster) GetInternalSubnetOk() (*string, bool)`

GetInternalSubnetOk returns a tuple with the InternalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSubnet

`func (o *HciCluster) SetInternalSubnet(v string)`

SetInternalSubnet sets InternalSubnet field to given value.

### HasInternalSubnet

`func (o *HciCluster) HasInternalSubnet() bool`

HasInternalSubnet returns a boolean if a field has been set.

### GetIsLts

`func (o *HciCluster) GetIsLts() bool`

GetIsLts returns the IsLts field if non-nil, zero value otherwise.

### GetIsLtsOk

`func (o *HciCluster) GetIsLtsOk() (*bool, bool)`

GetIsLtsOk returns a tuple with the IsLts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLts

`func (o *HciCluster) SetIsLts(v bool)`

SetIsLts sets IsLts field to given value.

### HasIsLts

`func (o *HciCluster) HasIsLts() bool`

HasIsLts returns a boolean if a field has been set.

### GetKeyManagementServerType

`func (o *HciCluster) GetKeyManagementServerType() string`

GetKeyManagementServerType returns the KeyManagementServerType field if non-nil, zero value otherwise.

### GetKeyManagementServerTypeOk

`func (o *HciCluster) GetKeyManagementServerTypeOk() (*string, bool)`

GetKeyManagementServerTypeOk returns a tuple with the KeyManagementServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagementServerType

`func (o *HciCluster) SetKeyManagementServerType(v string)`

SetKeyManagementServerType sets KeyManagementServerType field to given value.

### HasKeyManagementServerType

`func (o *HciCluster) HasKeyManagementServerType() bool`

HasKeyManagementServerType returns a boolean if a field has been set.

### GetManagementServer

`func (o *HciCluster) GetManagementServer() HciManagementServer`

GetManagementServer returns the ManagementServer field if non-nil, zero value otherwise.

### GetManagementServerOk

`func (o *HciCluster) GetManagementServerOk() (*HciManagementServer, bool)`

GetManagementServerOk returns a tuple with the ManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServer

`func (o *HciCluster) SetManagementServer(v HciManagementServer)`

SetManagementServer sets ManagementServer field to given value.

### HasManagementServer

`func (o *HciCluster) HasManagementServer() bool`

HasManagementServer returns a boolean if a field has been set.

### SetManagementServerNil

`func (o *HciCluster) SetManagementServerNil(b bool)`

 SetManagementServerNil sets the value for ManagementServer to be an explicit nil

### UnsetManagementServer
`func (o *HciCluster) UnsetManagementServer()`

UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
### GetMasqueradingIp

`func (o *HciCluster) GetMasqueradingIp() HciIpAddress`

GetMasqueradingIp returns the MasqueradingIp field if non-nil, zero value otherwise.

### GetMasqueradingIpOk

`func (o *HciCluster) GetMasqueradingIpOk() (*HciIpAddress, bool)`

GetMasqueradingIpOk returns a tuple with the MasqueradingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasqueradingIp

`func (o *HciCluster) SetMasqueradingIp(v HciIpAddress)`

SetMasqueradingIp sets MasqueradingIp field to given value.

### HasMasqueradingIp

`func (o *HciCluster) HasMasqueradingIp() bool`

HasMasqueradingIp returns a boolean if a field has been set.

### SetMasqueradingIpNil

`func (o *HciCluster) SetMasqueradingIpNil(b bool)`

 SetMasqueradingIpNil sets the value for MasqueradingIp to be an explicit nil

### UnsetMasqueradingIp
`func (o *HciCluster) UnsetMasqueradingIp()`

UnsetMasqueradingIp ensures that no value is present for MasqueradingIp, not even an explicit nil
### GetMasqueradingPort

`func (o *HciCluster) GetMasqueradingPort() int32`

GetMasqueradingPort returns the MasqueradingPort field if non-nil, zero value otherwise.

### GetMasqueradingPortOk

`func (o *HciCluster) GetMasqueradingPortOk() (*int32, bool)`

GetMasqueradingPortOk returns a tuple with the MasqueradingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasqueradingPort

`func (o *HciCluster) SetMasqueradingPort(v int32)`

SetMasqueradingPort sets MasqueradingPort field to given value.

### HasMasqueradingPort

`func (o *HciCluster) HasMasqueradingPort() bool`

HasMasqueradingPort returns a boolean if a field has been set.

### GetMemoryCapacityBytes

`func (o *HciCluster) GetMemoryCapacityBytes() int64`

GetMemoryCapacityBytes returns the MemoryCapacityBytes field if non-nil, zero value otherwise.

### GetMemoryCapacityBytesOk

`func (o *HciCluster) GetMemoryCapacityBytesOk() (*int64, bool)`

GetMemoryCapacityBytesOk returns a tuple with the MemoryCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacityBytes

`func (o *HciCluster) SetMemoryCapacityBytes(v int64)`

SetMemoryCapacityBytes sets MemoryCapacityBytes field to given value.

### HasMemoryCapacityBytes

`func (o *HciCluster) HasMemoryCapacityBytes() bool`

HasMemoryCapacityBytes returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *HciCluster) GetMemoryUsageBytes() int64`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *HciCluster) GetMemoryUsageBytesOk() (*int64, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *HciCluster) SetMemoryUsageBytes(v int64)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *HciCluster) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.

### GetName

`func (o *HciCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameServerIpList

`func (o *HciCluster) GetNameServerIpList() []HciIpAddressOrFqdn`

GetNameServerIpList returns the NameServerIpList field if non-nil, zero value otherwise.

### GetNameServerIpListOk

`func (o *HciCluster) GetNameServerIpListOk() (*[]HciIpAddressOrFqdn, bool)`

GetNameServerIpListOk returns a tuple with the NameServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServerIpList

`func (o *HciCluster) SetNameServerIpList(v []HciIpAddressOrFqdn)`

SetNameServerIpList sets NameServerIpList field to given value.

### HasNameServerIpList

`func (o *HciCluster) HasNameServerIpList() bool`

HasNameServerIpList returns a boolean if a field has been set.

### SetNameServerIpListNil

`func (o *HciCluster) SetNameServerIpListNil(b bool)`

 SetNameServerIpListNil sets the value for NameServerIpList to be an explicit nil

### UnsetNameServerIpList
`func (o *HciCluster) UnsetNameServerIpList()`

UnsetNameServerIpList ensures that no value is present for NameServerIpList, not even an explicit nil
### GetNtpServerIpList

`func (o *HciCluster) GetNtpServerIpList() []HciIpAddressOrFqdn`

GetNtpServerIpList returns the NtpServerIpList field if non-nil, zero value otherwise.

### GetNtpServerIpListOk

`func (o *HciCluster) GetNtpServerIpListOk() (*[]HciIpAddressOrFqdn, bool)`

GetNtpServerIpListOk returns a tuple with the NtpServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServerIpList

`func (o *HciCluster) SetNtpServerIpList(v []HciIpAddressOrFqdn)`

SetNtpServerIpList sets NtpServerIpList field to given value.

### HasNtpServerIpList

`func (o *HciCluster) HasNtpServerIpList() bool`

HasNtpServerIpList returns a boolean if a field has been set.

### SetNtpServerIpListNil

`func (o *HciCluster) SetNtpServerIpListNil(b bool)`

 SetNtpServerIpListNil sets the value for NtpServerIpList to be an explicit nil

### UnsetNtpServerIpList
`func (o *HciCluster) UnsetNtpServerIpList()`

UnsetNtpServerIpList ensures that no value is present for NtpServerIpList, not even an explicit nil
### GetNumberOfNodes

`func (o *HciCluster) GetNumberOfNodes() int32`

GetNumberOfNodes returns the NumberOfNodes field if non-nil, zero value otherwise.

### GetNumberOfNodesOk

`func (o *HciCluster) GetNumberOfNodesOk() (*int32, bool)`

GetNumberOfNodesOk returns a tuple with the NumberOfNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfNodes

`func (o *HciCluster) SetNumberOfNodes(v int32)`

SetNumberOfNodes sets NumberOfNodes field to given value.

### HasNumberOfNodes

`func (o *HciCluster) HasNumberOfNodes() bool`

HasNumberOfNodes returns a boolean if a field has been set.

### GetOperationMode

`func (o *HciCluster) GetOperationMode() string`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *HciCluster) GetOperationModeOk() (*string, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *HciCluster) SetOperationMode(v string)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *HciCluster) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.

### GetPasswordRemoteLoginEnabled

`func (o *HciCluster) GetPasswordRemoteLoginEnabled() bool`

GetPasswordRemoteLoginEnabled returns the PasswordRemoteLoginEnabled field if non-nil, zero value otherwise.

### GetPasswordRemoteLoginEnabledOk

`func (o *HciCluster) GetPasswordRemoteLoginEnabledOk() (*bool, bool)`

GetPasswordRemoteLoginEnabledOk returns a tuple with the PasswordRemoteLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRemoteLoginEnabled

`func (o *HciCluster) SetPasswordRemoteLoginEnabled(v bool)`

SetPasswordRemoteLoginEnabled sets PasswordRemoteLoginEnabled field to given value.

### HasPasswordRemoteLoginEnabled

`func (o *HciCluster) HasPasswordRemoteLoginEnabled() bool`

HasPasswordRemoteLoginEnabled returns a boolean if a field has been set.

### GetPcExtId

`func (o *HciCluster) GetPcExtId() string`

GetPcExtId returns the PcExtId field if non-nil, zero value otherwise.

### GetPcExtIdOk

`func (o *HciCluster) GetPcExtIdOk() (*string, bool)`

GetPcExtIdOk returns a tuple with the PcExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcExtId

`func (o *HciCluster) SetPcExtId(v string)`

SetPcExtId sets PcExtId field to given value.

### HasPcExtId

`func (o *HciCluster) HasPcExtId() bool`

HasPcExtId returns a boolean if a field has been set.

### GetPulseStatus

`func (o *HciCluster) GetPulseStatus() HciPulseStatus`

GetPulseStatus returns the PulseStatus field if non-nil, zero value otherwise.

### GetPulseStatusOk

`func (o *HciCluster) GetPulseStatusOk() (*HciPulseStatus, bool)`

GetPulseStatusOk returns a tuple with the PulseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulseStatus

`func (o *HciCluster) SetPulseStatus(v HciPulseStatus)`

SetPulseStatus sets PulseStatus field to given value.

### HasPulseStatus

`func (o *HciCluster) HasPulseStatus() bool`

HasPulseStatus returns a boolean if a field has been set.

### SetPulseStatusNil

`func (o *HciCluster) SetPulseStatusNil(b bool)`

 SetPulseStatusNil sets the value for PulseStatus to be an explicit nil

### UnsetPulseStatus
`func (o *HciCluster) UnsetPulseStatus()`

UnsetPulseStatus ensures that no value is present for PulseStatus, not even an explicit nil
### GetRedundancyFactor

`func (o *HciCluster) GetRedundancyFactor() int64`

GetRedundancyFactor returns the RedundancyFactor field if non-nil, zero value otherwise.

### GetRedundancyFactorOk

`func (o *HciCluster) GetRedundancyFactorOk() (*int64, bool)`

GetRedundancyFactorOk returns a tuple with the RedundancyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyFactor

`func (o *HciCluster) SetRedundancyFactor(v int64)`

SetRedundancyFactor sets RedundancyFactor field to given value.

### HasRedundancyFactor

`func (o *HciCluster) HasRedundancyFactor() bool`

HasRedundancyFactor returns a boolean if a field has been set.

### GetRemoteSupport

`func (o *HciCluster) GetRemoteSupport() bool`

GetRemoteSupport returns the RemoteSupport field if non-nil, zero value otherwise.

### GetRemoteSupportOk

`func (o *HciCluster) GetRemoteSupportOk() (*bool, bool)`

GetRemoteSupportOk returns a tuple with the RemoteSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSupport

`func (o *HciCluster) SetRemoteSupport(v bool)`

SetRemoteSupport sets RemoteSupport field to given value.

### HasRemoteSupport

`func (o *HciCluster) HasRemoteSupport() bool`

HasRemoteSupport returns a boolean if a field has been set.

### GetStorageCapacityBytes

`func (o *HciCluster) GetStorageCapacityBytes() int64`

GetStorageCapacityBytes returns the StorageCapacityBytes field if non-nil, zero value otherwise.

### GetStorageCapacityBytesOk

`func (o *HciCluster) GetStorageCapacityBytesOk() (*int64, bool)`

GetStorageCapacityBytesOk returns a tuple with the StorageCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacityBytes

`func (o *HciCluster) SetStorageCapacityBytes(v int64)`

SetStorageCapacityBytes sets StorageCapacityBytes field to given value.

### HasStorageCapacityBytes

`func (o *HciCluster) HasStorageCapacityBytes() bool`

HasStorageCapacityBytes returns a boolean if a field has been set.

### GetStorageUsageBytes

`func (o *HciCluster) GetStorageUsageBytes() int64`

GetStorageUsageBytes returns the StorageUsageBytes field if non-nil, zero value otherwise.

### GetStorageUsageBytesOk

`func (o *HciCluster) GetStorageUsageBytesOk() (*int64, bool)`

GetStorageUsageBytesOk returns a tuple with the StorageUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsageBytes

`func (o *HciCluster) SetStorageUsageBytes(v int64)`

SetStorageUsageBytes sets StorageUsageBytes field to given value.

### HasStorageUsageBytes

`func (o *HciCluster) HasStorageUsageBytes() bool`

HasStorageUsageBytes returns a boolean if a field has been set.

### GetTimezone

`func (o *HciCluster) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *HciCluster) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *HciCluster) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *HciCluster) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *HciCluster) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *HciCluster) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *HciCluster) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *HciCluster) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### GetVmCount

`func (o *HciCluster) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *HciCluster) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *HciCluster) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *HciCluster) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetCompliance

`func (o *HciCluster) GetCompliance() HciComplianceRelationship`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *HciCluster) GetComplianceOk() (*HciComplianceRelationship, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *HciCluster) SetCompliance(v HciComplianceRelationship)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *HciCluster) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.

### SetComplianceNil

`func (o *HciCluster) SetComplianceNil(b bool)`

 SetComplianceNil sets the value for Compliance to be an explicit nil

### UnsetCompliance
`func (o *HciCluster) UnsetCompliance()`

UnsetCompliance ensures that no value is present for Compliance, not even an explicit nil
### GetDomainManager

`func (o *HciCluster) GetDomainManager() HciDomainManagerRelationship`

GetDomainManager returns the DomainManager field if non-nil, zero value otherwise.

### GetDomainManagerOk

`func (o *HciCluster) GetDomainManagerOk() (*HciDomainManagerRelationship, bool)`

GetDomainManagerOk returns a tuple with the DomainManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainManager

`func (o *HciCluster) SetDomainManager(v HciDomainManagerRelationship)`

SetDomainManager sets DomainManager field to given value.

### HasDomainManager

`func (o *HciCluster) HasDomainManager() bool`

HasDomainManager returns a boolean if a field has been set.

### SetDomainManagerNil

`func (o *HciCluster) SetDomainManagerNil(b bool)`

 SetDomainManagerNil sets the value for DomainManager to be an explicit nil

### UnsetDomainManager
`func (o *HciCluster) UnsetDomainManager()`

UnsetDomainManager ensures that no value is present for DomainManager, not even an explicit nil
### GetEntitlement

`func (o *HciCluster) GetEntitlement() HciEntitlementRelationship`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *HciCluster) GetEntitlementOk() (*HciEntitlementRelationship, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *HciCluster) SetEntitlement(v HciEntitlementRelationship)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *HciCluster) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### SetEntitlementNil

`func (o *HciCluster) SetEntitlementNil(b bool)`

 SetEntitlementNil sets the value for Entitlement to be an explicit nil

### UnsetEntitlement
`func (o *HciCluster) UnsetEntitlement()`

UnsetEntitlement ensures that no value is present for Entitlement, not even an explicit nil
### GetNodes

`func (o *HciCluster) GetNodes() []HciNodeRelationship`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *HciCluster) GetNodesOk() (*[]HciNodeRelationship, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *HciCluster) SetNodes(v []HciNodeRelationship)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *HciCluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *HciCluster) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *HciCluster) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetRegisteredDevice

`func (o *HciCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciCluster) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciCluster) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetViolation

`func (o *HciCluster) GetViolation() HciViolationRelationship`

GetViolation returns the Violation field if non-nil, zero value otherwise.

### GetViolationOk

`func (o *HciCluster) GetViolationOk() (*HciViolationRelationship, bool)`

GetViolationOk returns a tuple with the Violation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolation

`func (o *HciCluster) SetViolation(v HciViolationRelationship)`

SetViolation sets Violation field to given value.

### HasViolation

`func (o *HciCluster) HasViolation() bool`

HasViolation returns a boolean if a field has been set.

### SetViolationNil

`func (o *HciCluster) SetViolationNil(b bool)`

 SetViolationNil sets the value for Violation to be an explicit nil

### UnsetViolation
`func (o *HciCluster) UnsetViolation()`

UnsetViolation ensures that no value is present for Violation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


