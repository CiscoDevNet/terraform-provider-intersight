# HciNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Node"]
**BlockModel** | Pointer to **string** | The rackable unit model of the node. | [optional] [readonly] 
**BlockSerial** | Pointer to **string** | The rackable unit serial number of the node. | [optional] [readonly] 
**BootTimeUsecs** | Pointer to **int64** | The boot time in microseconds of the node. | [optional] [readonly] 
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The name of the cluster this node belongs to. | [optional] [readonly] 
**ControllerVmBackplaneAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**ControllerVmExternalAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**ControllerVmId** | Pointer to **int32** | The identifier number of the controller VM. | [optional] [readonly] 
**ControllerVmMaintanenceMode** | Pointer to **bool** | The maintenance mode status of the controller VM. | [optional] [readonly] 
**ControllerVmNatIp** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**ControllerVmNatPort** | Pointer to **int32** | The NAT port of the controller VM. | [optional] [readonly] 
**ControllerVmRdmaBackplaneAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**ControllerVmServerUuid** | Pointer to **string** | The Rackable unit UUID of the server. | [optional] [readonly] 
**CpuCapacityHz** | Pointer to **int64** | The CPU capacity in Hz of the node. | [optional] [readonly] 
**CpuFrequencyHz** | Pointer to **int64** | The CPU frequency in Hz on the node. | [optional] [readonly] 
**CpuModel** | Pointer to **string** | The CPU model of the node. | [optional] [readonly] 
**CpuUsageHz** | Pointer to **int64** | The CPU usage in Hz of the node. | [optional] [readonly] 
**DefaultVhdContainerUuid** | Pointer to **string** | The default VHD container UUID of the node. | [optional] [readonly] 
**DefaultVhdLocation** | Pointer to **string** | The default VHD location of the node. | [optional] [readonly] 
**DefaultVmContainerUuid** | Pointer to **string** | The default VM container UUID of the node. | [optional] [readonly] 
**DefaultVmLocation** | Pointer to **string** | The default VM location of the node. | [optional] [readonly] 
**DiskCount** | Pointer to **int64** | The number of disks on the node. | [optional] [readonly] 
**FailoverClusterFqdn** | Pointer to **string** | The failover cluster FQDN of the node. | [optional] [readonly] 
**FailoverClusterNodeStatus** | Pointer to **string** | The failover cluster node status of the node. | [optional] [readonly] 
**GpuCount** | Pointer to **int64** | The number of GPUs on the node. | [optional] [readonly] 
**GpuDriverVersion** | Pointer to **string** | The GPU driver version of the node. | [optional] [readonly] 
**HasCsr** | Pointer to **bool** | Certificate signing request status of the node. | [optional] [readonly] 
**HostName** | Pointer to **string** | The name of the host the node is running on. | [optional] [readonly] 
**HostType** | Pointer to **string** | The type of the host, e.g. HYPER_CONVERGED, COMPUTE_ONLY, STORAGE_ONLY. | [optional] [readonly] 
**HypervisorAcropolisConnectionState** | Pointer to **string** | The connection state of the hypervisor, e.g. CONNECTED, DISCONNECTED, NOT_AVAILABLE. | [optional] [readonly] 
**HypervisorExternalAddress** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**HypervisorNumberOfVms** | Pointer to **int64** | The number of VMs managed on this node. | [optional] [readonly] 
**HypervisorState** | Pointer to **string** | The hypervisor state e.g. ACROPOLIS_NORMAL, ENTERING_MAINTENANCE_MODE, ENTERED_MAINTENANCE_MODE, RESERVED_FOR_HA_FAILOVER, ENTERING_MAINTENANCE_MODE_FROM_HA_FAILOVER, RESERVING_FOR_HA_FAILOVER, HA_FAILOVER_SOURCE, HA_FAILOVER_TARGET, HA_HEALING_SOURCE, HA_HEALING_TARGET. | [optional] [readonly] 
**HypervisorType** | Pointer to **string** | The hypervisor type, e.g. AHV, ESX, HYPERV, XEN, NATIVEHOST etc. | [optional] [readonly] 
**HypervisorUserName** | Pointer to **string** | The user name of the hypervisor on this node. | [optional] [readonly] 
**HypervisorVersion** | Pointer to **string** | The version of the hypervisor on this node. | [optional] [readonly] 
**IpmiIp** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**IpmiUsername** | Pointer to **string** | The IPMI user name of the controller. | [optional] [readonly] 
**IsDegraded** | Pointer to **bool** | The degraded status of the node. | [optional] [readonly] 
**IsHardwareVirtualized** | Pointer to **bool** | The hardware virtualization status of the node. | [optional] [readonly] 
**IsSecureBooted** | Pointer to **bool** | The secure boot status of the node. | [optional] [readonly] 
**KeyManagementDeviceToCertStatus** | Pointer to [**[]HciKeyManagementDeviceToCertStatusInfo**](HciKeyManagementDeviceToCertStatusInfo.md) |  | [optional] 
**MaintenanceState** | Pointer to **string** | The maintenance state of the node. | [optional] [readonly] 
**MemoryCapacityBytes** | Pointer to **int64** | The memory capacity in bytes of the node. | [optional] [readonly] 
**MemorySizeBytes** | Pointer to **int64** | The memory size in bytes of the node. | [optional] [readonly] 
**MemoryUsageBytes** | Pointer to **int64** | The memory usage in bytes of the node. | [optional] [readonly] 
**NodeExtId** | Pointer to **string** | The unique identifier of the node. | [optional] [readonly] 
**NodeSerial** | Pointer to **string** | The serial number of this node. | [optional] [readonly] 
**NodeStatus** | Pointer to **string** | The status of the node such as NORMAL, TO_BE_REMOVED, OK_TO_BE_REMOVED, NEW_NODE, TO_BE_PREPROTECTED, PREPROTECTED. | [optional] [readonly] 
**NumberOfCpuCores** | Pointer to **int64** | The number of CPU cores on the node. | [optional] [readonly] 
**NumberOfCpuSockets** | Pointer to **int64** | The number of sockets on the node. | [optional] [readonly] 
**NumberOfCpuThreads** | Pointer to **int64** | The number of threads on the node. | [optional] [readonly] 
**RebootPending** | Pointer to **bool** | The reboot pending status of the node. | [optional] [readonly] 
**StorageCapacityBytes** | Pointer to **int64** | The storage capacity in bytes of the node. | [optional] [readonly] 
**StorageUsageBytes** | Pointer to **int64** | The storage usage in bytes of the node. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**Disks** | Pointer to [**[]HciDiskRelationship**](HciDiskRelationship.md) | An array of relationships to hciDisk resources. | [optional] [readonly] 
**Gpus** | Pointer to [**[]HciGpuRelationship**](HciGpuRelationship.md) | An array of relationships to hciGpu resources. | [optional] [readonly] 
**PhysicalServer** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vms** | Pointer to [**[]HciBaseVmRelationship**](HciBaseVmRelationship.md) | An array of relationships to hciBaseVm resources. | [optional] [readonly] 

## Methods

### NewHciNode

`func NewHciNode(classId string, objectType string, ) *HciNode`

NewHciNode instantiates a new HciNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciNodeWithDefaults

`func NewHciNodeWithDefaults() *HciNode`

NewHciNodeWithDefaults instantiates a new HciNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlockModel

`func (o *HciNode) GetBlockModel() string`

GetBlockModel returns the BlockModel field if non-nil, zero value otherwise.

### GetBlockModelOk

`func (o *HciNode) GetBlockModelOk() (*string, bool)`

GetBlockModelOk returns a tuple with the BlockModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockModel

`func (o *HciNode) SetBlockModel(v string)`

SetBlockModel sets BlockModel field to given value.

### HasBlockModel

`func (o *HciNode) HasBlockModel() bool`

HasBlockModel returns a boolean if a field has been set.

### GetBlockSerial

`func (o *HciNode) GetBlockSerial() string`

GetBlockSerial returns the BlockSerial field if non-nil, zero value otherwise.

### GetBlockSerialOk

`func (o *HciNode) GetBlockSerialOk() (*string, bool)`

GetBlockSerialOk returns a tuple with the BlockSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSerial

`func (o *HciNode) SetBlockSerial(v string)`

SetBlockSerial sets BlockSerial field to given value.

### HasBlockSerial

`func (o *HciNode) HasBlockSerial() bool`

HasBlockSerial returns a boolean if a field has been set.

### GetBootTimeUsecs

`func (o *HciNode) GetBootTimeUsecs() int64`

GetBootTimeUsecs returns the BootTimeUsecs field if non-nil, zero value otherwise.

### GetBootTimeUsecsOk

`func (o *HciNode) GetBootTimeUsecsOk() (*int64, bool)`

GetBootTimeUsecsOk returns a tuple with the BootTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTimeUsecs

`func (o *HciNode) SetBootTimeUsecs(v int64)`

SetBootTimeUsecs sets BootTimeUsecs field to given value.

### HasBootTimeUsecs

`func (o *HciNode) HasBootTimeUsecs() bool`

HasBootTimeUsecs returns a boolean if a field has been set.

### GetClusterExtId

`func (o *HciNode) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciNode) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciNode) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciNode) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetClusterName

`func (o *HciNode) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HciNode) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HciNode) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HciNode) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetControllerVmBackplaneAddress

`func (o *HciNode) GetControllerVmBackplaneAddress() HciIpAddress`

GetControllerVmBackplaneAddress returns the ControllerVmBackplaneAddress field if non-nil, zero value otherwise.

### GetControllerVmBackplaneAddressOk

`func (o *HciNode) GetControllerVmBackplaneAddressOk() (*HciIpAddress, bool)`

GetControllerVmBackplaneAddressOk returns a tuple with the ControllerVmBackplaneAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmBackplaneAddress

`func (o *HciNode) SetControllerVmBackplaneAddress(v HciIpAddress)`

SetControllerVmBackplaneAddress sets ControllerVmBackplaneAddress field to given value.

### HasControllerVmBackplaneAddress

`func (o *HciNode) HasControllerVmBackplaneAddress() bool`

HasControllerVmBackplaneAddress returns a boolean if a field has been set.

### SetControllerVmBackplaneAddressNil

`func (o *HciNode) SetControllerVmBackplaneAddressNil(b bool)`

 SetControllerVmBackplaneAddressNil sets the value for ControllerVmBackplaneAddress to be an explicit nil

### UnsetControllerVmBackplaneAddress
`func (o *HciNode) UnsetControllerVmBackplaneAddress()`

UnsetControllerVmBackplaneAddress ensures that no value is present for ControllerVmBackplaneAddress, not even an explicit nil
### GetControllerVmExternalAddress

`func (o *HciNode) GetControllerVmExternalAddress() HciIpAddress`

GetControllerVmExternalAddress returns the ControllerVmExternalAddress field if non-nil, zero value otherwise.

### GetControllerVmExternalAddressOk

`func (o *HciNode) GetControllerVmExternalAddressOk() (*HciIpAddress, bool)`

GetControllerVmExternalAddressOk returns a tuple with the ControllerVmExternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmExternalAddress

`func (o *HciNode) SetControllerVmExternalAddress(v HciIpAddress)`

SetControllerVmExternalAddress sets ControllerVmExternalAddress field to given value.

### HasControllerVmExternalAddress

`func (o *HciNode) HasControllerVmExternalAddress() bool`

HasControllerVmExternalAddress returns a boolean if a field has been set.

### SetControllerVmExternalAddressNil

`func (o *HciNode) SetControllerVmExternalAddressNil(b bool)`

 SetControllerVmExternalAddressNil sets the value for ControllerVmExternalAddress to be an explicit nil

### UnsetControllerVmExternalAddress
`func (o *HciNode) UnsetControllerVmExternalAddress()`

UnsetControllerVmExternalAddress ensures that no value is present for ControllerVmExternalAddress, not even an explicit nil
### GetControllerVmId

`func (o *HciNode) GetControllerVmId() int32`

GetControllerVmId returns the ControllerVmId field if non-nil, zero value otherwise.

### GetControllerVmIdOk

`func (o *HciNode) GetControllerVmIdOk() (*int32, bool)`

GetControllerVmIdOk returns a tuple with the ControllerVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmId

`func (o *HciNode) SetControllerVmId(v int32)`

SetControllerVmId sets ControllerVmId field to given value.

### HasControllerVmId

`func (o *HciNode) HasControllerVmId() bool`

HasControllerVmId returns a boolean if a field has been set.

### GetControllerVmMaintanenceMode

`func (o *HciNode) GetControllerVmMaintanenceMode() bool`

GetControllerVmMaintanenceMode returns the ControllerVmMaintanenceMode field if non-nil, zero value otherwise.

### GetControllerVmMaintanenceModeOk

`func (o *HciNode) GetControllerVmMaintanenceModeOk() (*bool, bool)`

GetControllerVmMaintanenceModeOk returns a tuple with the ControllerVmMaintanenceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmMaintanenceMode

`func (o *HciNode) SetControllerVmMaintanenceMode(v bool)`

SetControllerVmMaintanenceMode sets ControllerVmMaintanenceMode field to given value.

### HasControllerVmMaintanenceMode

`func (o *HciNode) HasControllerVmMaintanenceMode() bool`

HasControllerVmMaintanenceMode returns a boolean if a field has been set.

### GetControllerVmNatIp

`func (o *HciNode) GetControllerVmNatIp() HciIpAddress`

GetControllerVmNatIp returns the ControllerVmNatIp field if non-nil, zero value otherwise.

### GetControllerVmNatIpOk

`func (o *HciNode) GetControllerVmNatIpOk() (*HciIpAddress, bool)`

GetControllerVmNatIpOk returns a tuple with the ControllerVmNatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmNatIp

`func (o *HciNode) SetControllerVmNatIp(v HciIpAddress)`

SetControllerVmNatIp sets ControllerVmNatIp field to given value.

### HasControllerVmNatIp

`func (o *HciNode) HasControllerVmNatIp() bool`

HasControllerVmNatIp returns a boolean if a field has been set.

### SetControllerVmNatIpNil

`func (o *HciNode) SetControllerVmNatIpNil(b bool)`

 SetControllerVmNatIpNil sets the value for ControllerVmNatIp to be an explicit nil

### UnsetControllerVmNatIp
`func (o *HciNode) UnsetControllerVmNatIp()`

UnsetControllerVmNatIp ensures that no value is present for ControllerVmNatIp, not even an explicit nil
### GetControllerVmNatPort

`func (o *HciNode) GetControllerVmNatPort() int32`

GetControllerVmNatPort returns the ControllerVmNatPort field if non-nil, zero value otherwise.

### GetControllerVmNatPortOk

`func (o *HciNode) GetControllerVmNatPortOk() (*int32, bool)`

GetControllerVmNatPortOk returns a tuple with the ControllerVmNatPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmNatPort

`func (o *HciNode) SetControllerVmNatPort(v int32)`

SetControllerVmNatPort sets ControllerVmNatPort field to given value.

### HasControllerVmNatPort

`func (o *HciNode) HasControllerVmNatPort() bool`

HasControllerVmNatPort returns a boolean if a field has been set.

### GetControllerVmRdmaBackplaneAddress

`func (o *HciNode) GetControllerVmRdmaBackplaneAddress() HciIpAddress`

GetControllerVmRdmaBackplaneAddress returns the ControllerVmRdmaBackplaneAddress field if non-nil, zero value otherwise.

### GetControllerVmRdmaBackplaneAddressOk

`func (o *HciNode) GetControllerVmRdmaBackplaneAddressOk() (*HciIpAddress, bool)`

GetControllerVmRdmaBackplaneAddressOk returns a tuple with the ControllerVmRdmaBackplaneAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmRdmaBackplaneAddress

`func (o *HciNode) SetControllerVmRdmaBackplaneAddress(v HciIpAddress)`

SetControllerVmRdmaBackplaneAddress sets ControllerVmRdmaBackplaneAddress field to given value.

### HasControllerVmRdmaBackplaneAddress

`func (o *HciNode) HasControllerVmRdmaBackplaneAddress() bool`

HasControllerVmRdmaBackplaneAddress returns a boolean if a field has been set.

### SetControllerVmRdmaBackplaneAddressNil

`func (o *HciNode) SetControllerVmRdmaBackplaneAddressNil(b bool)`

 SetControllerVmRdmaBackplaneAddressNil sets the value for ControllerVmRdmaBackplaneAddress to be an explicit nil

### UnsetControllerVmRdmaBackplaneAddress
`func (o *HciNode) UnsetControllerVmRdmaBackplaneAddress()`

UnsetControllerVmRdmaBackplaneAddress ensures that no value is present for ControllerVmRdmaBackplaneAddress, not even an explicit nil
### GetControllerVmServerUuid

`func (o *HciNode) GetControllerVmServerUuid() string`

GetControllerVmServerUuid returns the ControllerVmServerUuid field if non-nil, zero value otherwise.

### GetControllerVmServerUuidOk

`func (o *HciNode) GetControllerVmServerUuidOk() (*string, bool)`

GetControllerVmServerUuidOk returns a tuple with the ControllerVmServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmServerUuid

`func (o *HciNode) SetControllerVmServerUuid(v string)`

SetControllerVmServerUuid sets ControllerVmServerUuid field to given value.

### HasControllerVmServerUuid

`func (o *HciNode) HasControllerVmServerUuid() bool`

HasControllerVmServerUuid returns a boolean if a field has been set.

### GetCpuCapacityHz

`func (o *HciNode) GetCpuCapacityHz() int64`

GetCpuCapacityHz returns the CpuCapacityHz field if non-nil, zero value otherwise.

### GetCpuCapacityHzOk

`func (o *HciNode) GetCpuCapacityHzOk() (*int64, bool)`

GetCpuCapacityHzOk returns a tuple with the CpuCapacityHz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCapacityHz

`func (o *HciNode) SetCpuCapacityHz(v int64)`

SetCpuCapacityHz sets CpuCapacityHz field to given value.

### HasCpuCapacityHz

`func (o *HciNode) HasCpuCapacityHz() bool`

HasCpuCapacityHz returns a boolean if a field has been set.

### GetCpuFrequencyHz

`func (o *HciNode) GetCpuFrequencyHz() int64`

GetCpuFrequencyHz returns the CpuFrequencyHz field if non-nil, zero value otherwise.

### GetCpuFrequencyHzOk

`func (o *HciNode) GetCpuFrequencyHzOk() (*int64, bool)`

GetCpuFrequencyHzOk returns a tuple with the CpuFrequencyHz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyHz

`func (o *HciNode) SetCpuFrequencyHz(v int64)`

SetCpuFrequencyHz sets CpuFrequencyHz field to given value.

### HasCpuFrequencyHz

`func (o *HciNode) HasCpuFrequencyHz() bool`

HasCpuFrequencyHz returns a boolean if a field has been set.

### GetCpuModel

`func (o *HciNode) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *HciNode) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *HciNode) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *HciNode) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCpuUsageHz

`func (o *HciNode) GetCpuUsageHz() int64`

GetCpuUsageHz returns the CpuUsageHz field if non-nil, zero value otherwise.

### GetCpuUsageHzOk

`func (o *HciNode) GetCpuUsageHzOk() (*int64, bool)`

GetCpuUsageHzOk returns a tuple with the CpuUsageHz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageHz

`func (o *HciNode) SetCpuUsageHz(v int64)`

SetCpuUsageHz sets CpuUsageHz field to given value.

### HasCpuUsageHz

`func (o *HciNode) HasCpuUsageHz() bool`

HasCpuUsageHz returns a boolean if a field has been set.

### GetDefaultVhdContainerUuid

`func (o *HciNode) GetDefaultVhdContainerUuid() string`

GetDefaultVhdContainerUuid returns the DefaultVhdContainerUuid field if non-nil, zero value otherwise.

### GetDefaultVhdContainerUuidOk

`func (o *HciNode) GetDefaultVhdContainerUuidOk() (*string, bool)`

GetDefaultVhdContainerUuidOk returns a tuple with the DefaultVhdContainerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVhdContainerUuid

`func (o *HciNode) SetDefaultVhdContainerUuid(v string)`

SetDefaultVhdContainerUuid sets DefaultVhdContainerUuid field to given value.

### HasDefaultVhdContainerUuid

`func (o *HciNode) HasDefaultVhdContainerUuid() bool`

HasDefaultVhdContainerUuid returns a boolean if a field has been set.

### GetDefaultVhdLocation

`func (o *HciNode) GetDefaultVhdLocation() string`

GetDefaultVhdLocation returns the DefaultVhdLocation field if non-nil, zero value otherwise.

### GetDefaultVhdLocationOk

`func (o *HciNode) GetDefaultVhdLocationOk() (*string, bool)`

GetDefaultVhdLocationOk returns a tuple with the DefaultVhdLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVhdLocation

`func (o *HciNode) SetDefaultVhdLocation(v string)`

SetDefaultVhdLocation sets DefaultVhdLocation field to given value.

### HasDefaultVhdLocation

`func (o *HciNode) HasDefaultVhdLocation() bool`

HasDefaultVhdLocation returns a boolean if a field has been set.

### GetDefaultVmContainerUuid

`func (o *HciNode) GetDefaultVmContainerUuid() string`

GetDefaultVmContainerUuid returns the DefaultVmContainerUuid field if non-nil, zero value otherwise.

### GetDefaultVmContainerUuidOk

`func (o *HciNode) GetDefaultVmContainerUuidOk() (*string, bool)`

GetDefaultVmContainerUuidOk returns a tuple with the DefaultVmContainerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVmContainerUuid

`func (o *HciNode) SetDefaultVmContainerUuid(v string)`

SetDefaultVmContainerUuid sets DefaultVmContainerUuid field to given value.

### HasDefaultVmContainerUuid

`func (o *HciNode) HasDefaultVmContainerUuid() bool`

HasDefaultVmContainerUuid returns a boolean if a field has been set.

### GetDefaultVmLocation

`func (o *HciNode) GetDefaultVmLocation() string`

GetDefaultVmLocation returns the DefaultVmLocation field if non-nil, zero value otherwise.

### GetDefaultVmLocationOk

`func (o *HciNode) GetDefaultVmLocationOk() (*string, bool)`

GetDefaultVmLocationOk returns a tuple with the DefaultVmLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVmLocation

`func (o *HciNode) SetDefaultVmLocation(v string)`

SetDefaultVmLocation sets DefaultVmLocation field to given value.

### HasDefaultVmLocation

`func (o *HciNode) HasDefaultVmLocation() bool`

HasDefaultVmLocation returns a boolean if a field has been set.

### GetDiskCount

`func (o *HciNode) GetDiskCount() int64`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *HciNode) GetDiskCountOk() (*int64, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *HciNode) SetDiskCount(v int64)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *HciNode) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### GetFailoverClusterFqdn

`func (o *HciNode) GetFailoverClusterFqdn() string`

GetFailoverClusterFqdn returns the FailoverClusterFqdn field if non-nil, zero value otherwise.

### GetFailoverClusterFqdnOk

`func (o *HciNode) GetFailoverClusterFqdnOk() (*string, bool)`

GetFailoverClusterFqdnOk returns a tuple with the FailoverClusterFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverClusterFqdn

`func (o *HciNode) SetFailoverClusterFqdn(v string)`

SetFailoverClusterFqdn sets FailoverClusterFqdn field to given value.

### HasFailoverClusterFqdn

`func (o *HciNode) HasFailoverClusterFqdn() bool`

HasFailoverClusterFqdn returns a boolean if a field has been set.

### GetFailoverClusterNodeStatus

`func (o *HciNode) GetFailoverClusterNodeStatus() string`

GetFailoverClusterNodeStatus returns the FailoverClusterNodeStatus field if non-nil, zero value otherwise.

### GetFailoverClusterNodeStatusOk

`func (o *HciNode) GetFailoverClusterNodeStatusOk() (*string, bool)`

GetFailoverClusterNodeStatusOk returns a tuple with the FailoverClusterNodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverClusterNodeStatus

`func (o *HciNode) SetFailoverClusterNodeStatus(v string)`

SetFailoverClusterNodeStatus sets FailoverClusterNodeStatus field to given value.

### HasFailoverClusterNodeStatus

`func (o *HciNode) HasFailoverClusterNodeStatus() bool`

HasFailoverClusterNodeStatus returns a boolean if a field has been set.

### GetGpuCount

`func (o *HciNode) GetGpuCount() int64`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *HciNode) GetGpuCountOk() (*int64, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *HciNode) SetGpuCount(v int64)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *HciNode) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuDriverVersion

`func (o *HciNode) GetGpuDriverVersion() string`

GetGpuDriverVersion returns the GpuDriverVersion field if non-nil, zero value otherwise.

### GetGpuDriverVersionOk

`func (o *HciNode) GetGpuDriverVersionOk() (*string, bool)`

GetGpuDriverVersionOk returns a tuple with the GpuDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDriverVersion

`func (o *HciNode) SetGpuDriverVersion(v string)`

SetGpuDriverVersion sets GpuDriverVersion field to given value.

### HasGpuDriverVersion

`func (o *HciNode) HasGpuDriverVersion() bool`

HasGpuDriverVersion returns a boolean if a field has been set.

### GetHasCsr

`func (o *HciNode) GetHasCsr() bool`

GetHasCsr returns the HasCsr field if non-nil, zero value otherwise.

### GetHasCsrOk

`func (o *HciNode) GetHasCsrOk() (*bool, bool)`

GetHasCsrOk returns a tuple with the HasCsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCsr

`func (o *HciNode) SetHasCsr(v bool)`

SetHasCsr sets HasCsr field to given value.

### HasHasCsr

`func (o *HciNode) HasHasCsr() bool`

HasHasCsr returns a boolean if a field has been set.

### GetHostName

`func (o *HciNode) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HciNode) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HciNode) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HciNode) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHostType

`func (o *HciNode) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *HciNode) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *HciNode) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *HciNode) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetHypervisorAcropolisConnectionState

`func (o *HciNode) GetHypervisorAcropolisConnectionState() string`

GetHypervisorAcropolisConnectionState returns the HypervisorAcropolisConnectionState field if non-nil, zero value otherwise.

### GetHypervisorAcropolisConnectionStateOk

`func (o *HciNode) GetHypervisorAcropolisConnectionStateOk() (*string, bool)`

GetHypervisorAcropolisConnectionStateOk returns a tuple with the HypervisorAcropolisConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorAcropolisConnectionState

`func (o *HciNode) SetHypervisorAcropolisConnectionState(v string)`

SetHypervisorAcropolisConnectionState sets HypervisorAcropolisConnectionState field to given value.

### HasHypervisorAcropolisConnectionState

`func (o *HciNode) HasHypervisorAcropolisConnectionState() bool`

HasHypervisorAcropolisConnectionState returns a boolean if a field has been set.

### GetHypervisorExternalAddress

`func (o *HciNode) GetHypervisorExternalAddress() HciIpAddress`

GetHypervisorExternalAddress returns the HypervisorExternalAddress field if non-nil, zero value otherwise.

### GetHypervisorExternalAddressOk

`func (o *HciNode) GetHypervisorExternalAddressOk() (*HciIpAddress, bool)`

GetHypervisorExternalAddressOk returns a tuple with the HypervisorExternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorExternalAddress

`func (o *HciNode) SetHypervisorExternalAddress(v HciIpAddress)`

SetHypervisorExternalAddress sets HypervisorExternalAddress field to given value.

### HasHypervisorExternalAddress

`func (o *HciNode) HasHypervisorExternalAddress() bool`

HasHypervisorExternalAddress returns a boolean if a field has been set.

### SetHypervisorExternalAddressNil

`func (o *HciNode) SetHypervisorExternalAddressNil(b bool)`

 SetHypervisorExternalAddressNil sets the value for HypervisorExternalAddress to be an explicit nil

### UnsetHypervisorExternalAddress
`func (o *HciNode) UnsetHypervisorExternalAddress()`

UnsetHypervisorExternalAddress ensures that no value is present for HypervisorExternalAddress, not even an explicit nil
### GetHypervisorNumberOfVms

`func (o *HciNode) GetHypervisorNumberOfVms() int64`

GetHypervisorNumberOfVms returns the HypervisorNumberOfVms field if non-nil, zero value otherwise.

### GetHypervisorNumberOfVmsOk

`func (o *HciNode) GetHypervisorNumberOfVmsOk() (*int64, bool)`

GetHypervisorNumberOfVmsOk returns a tuple with the HypervisorNumberOfVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorNumberOfVms

`func (o *HciNode) SetHypervisorNumberOfVms(v int64)`

SetHypervisorNumberOfVms sets HypervisorNumberOfVms field to given value.

### HasHypervisorNumberOfVms

`func (o *HciNode) HasHypervisorNumberOfVms() bool`

HasHypervisorNumberOfVms returns a boolean if a field has been set.

### GetHypervisorState

`func (o *HciNode) GetHypervisorState() string`

GetHypervisorState returns the HypervisorState field if non-nil, zero value otherwise.

### GetHypervisorStateOk

`func (o *HciNode) GetHypervisorStateOk() (*string, bool)`

GetHypervisorStateOk returns a tuple with the HypervisorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorState

`func (o *HciNode) SetHypervisorState(v string)`

SetHypervisorState sets HypervisorState field to given value.

### HasHypervisorState

`func (o *HciNode) HasHypervisorState() bool`

HasHypervisorState returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HciNode) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HciNode) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HciNode) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HciNode) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetHypervisorUserName

`func (o *HciNode) GetHypervisorUserName() string`

GetHypervisorUserName returns the HypervisorUserName field if non-nil, zero value otherwise.

### GetHypervisorUserNameOk

`func (o *HciNode) GetHypervisorUserNameOk() (*string, bool)`

GetHypervisorUserNameOk returns a tuple with the HypervisorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorUserName

`func (o *HciNode) SetHypervisorUserName(v string)`

SetHypervisorUserName sets HypervisorUserName field to given value.

### HasHypervisorUserName

`func (o *HciNode) HasHypervisorUserName() bool`

HasHypervisorUserName returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HciNode) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HciNode) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HciNode) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HciNode) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetIpmiIp

`func (o *HciNode) GetIpmiIp() HciIpAddress`

GetIpmiIp returns the IpmiIp field if non-nil, zero value otherwise.

### GetIpmiIpOk

`func (o *HciNode) GetIpmiIpOk() (*HciIpAddress, bool)`

GetIpmiIpOk returns a tuple with the IpmiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiIp

`func (o *HciNode) SetIpmiIp(v HciIpAddress)`

SetIpmiIp sets IpmiIp field to given value.

### HasIpmiIp

`func (o *HciNode) HasIpmiIp() bool`

HasIpmiIp returns a boolean if a field has been set.

### SetIpmiIpNil

`func (o *HciNode) SetIpmiIpNil(b bool)`

 SetIpmiIpNil sets the value for IpmiIp to be an explicit nil

### UnsetIpmiIp
`func (o *HciNode) UnsetIpmiIp()`

UnsetIpmiIp ensures that no value is present for IpmiIp, not even an explicit nil
### GetIpmiUsername

`func (o *HciNode) GetIpmiUsername() string`

GetIpmiUsername returns the IpmiUsername field if non-nil, zero value otherwise.

### GetIpmiUsernameOk

`func (o *HciNode) GetIpmiUsernameOk() (*string, bool)`

GetIpmiUsernameOk returns a tuple with the IpmiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiUsername

`func (o *HciNode) SetIpmiUsername(v string)`

SetIpmiUsername sets IpmiUsername field to given value.

### HasIpmiUsername

`func (o *HciNode) HasIpmiUsername() bool`

HasIpmiUsername returns a boolean if a field has been set.

### GetIsDegraded

`func (o *HciNode) GetIsDegraded() bool`

GetIsDegraded returns the IsDegraded field if non-nil, zero value otherwise.

### GetIsDegradedOk

`func (o *HciNode) GetIsDegradedOk() (*bool, bool)`

GetIsDegradedOk returns a tuple with the IsDegraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDegraded

`func (o *HciNode) SetIsDegraded(v bool)`

SetIsDegraded sets IsDegraded field to given value.

### HasIsDegraded

`func (o *HciNode) HasIsDegraded() bool`

HasIsDegraded returns a boolean if a field has been set.

### GetIsHardwareVirtualized

`func (o *HciNode) GetIsHardwareVirtualized() bool`

GetIsHardwareVirtualized returns the IsHardwareVirtualized field if non-nil, zero value otherwise.

### GetIsHardwareVirtualizedOk

`func (o *HciNode) GetIsHardwareVirtualizedOk() (*bool, bool)`

GetIsHardwareVirtualizedOk returns a tuple with the IsHardwareVirtualized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardwareVirtualized

`func (o *HciNode) SetIsHardwareVirtualized(v bool)`

SetIsHardwareVirtualized sets IsHardwareVirtualized field to given value.

### HasIsHardwareVirtualized

`func (o *HciNode) HasIsHardwareVirtualized() bool`

HasIsHardwareVirtualized returns a boolean if a field has been set.

### GetIsSecureBooted

`func (o *HciNode) GetIsSecureBooted() bool`

GetIsSecureBooted returns the IsSecureBooted field if non-nil, zero value otherwise.

### GetIsSecureBootedOk

`func (o *HciNode) GetIsSecureBootedOk() (*bool, bool)`

GetIsSecureBootedOk returns a tuple with the IsSecureBooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBooted

`func (o *HciNode) SetIsSecureBooted(v bool)`

SetIsSecureBooted sets IsSecureBooted field to given value.

### HasIsSecureBooted

`func (o *HciNode) HasIsSecureBooted() bool`

HasIsSecureBooted returns a boolean if a field has been set.

### GetKeyManagementDeviceToCertStatus

`func (o *HciNode) GetKeyManagementDeviceToCertStatus() []HciKeyManagementDeviceToCertStatusInfo`

GetKeyManagementDeviceToCertStatus returns the KeyManagementDeviceToCertStatus field if non-nil, zero value otherwise.

### GetKeyManagementDeviceToCertStatusOk

`func (o *HciNode) GetKeyManagementDeviceToCertStatusOk() (*[]HciKeyManagementDeviceToCertStatusInfo, bool)`

GetKeyManagementDeviceToCertStatusOk returns a tuple with the KeyManagementDeviceToCertStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagementDeviceToCertStatus

`func (o *HciNode) SetKeyManagementDeviceToCertStatus(v []HciKeyManagementDeviceToCertStatusInfo)`

SetKeyManagementDeviceToCertStatus sets KeyManagementDeviceToCertStatus field to given value.

### HasKeyManagementDeviceToCertStatus

`func (o *HciNode) HasKeyManagementDeviceToCertStatus() bool`

HasKeyManagementDeviceToCertStatus returns a boolean if a field has been set.

### SetKeyManagementDeviceToCertStatusNil

`func (o *HciNode) SetKeyManagementDeviceToCertStatusNil(b bool)`

 SetKeyManagementDeviceToCertStatusNil sets the value for KeyManagementDeviceToCertStatus to be an explicit nil

### UnsetKeyManagementDeviceToCertStatus
`func (o *HciNode) UnsetKeyManagementDeviceToCertStatus()`

UnsetKeyManagementDeviceToCertStatus ensures that no value is present for KeyManagementDeviceToCertStatus, not even an explicit nil
### GetMaintenanceState

`func (o *HciNode) GetMaintenanceState() string`

GetMaintenanceState returns the MaintenanceState field if non-nil, zero value otherwise.

### GetMaintenanceStateOk

`func (o *HciNode) GetMaintenanceStateOk() (*string, bool)`

GetMaintenanceStateOk returns a tuple with the MaintenanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceState

`func (o *HciNode) SetMaintenanceState(v string)`

SetMaintenanceState sets MaintenanceState field to given value.

### HasMaintenanceState

`func (o *HciNode) HasMaintenanceState() bool`

HasMaintenanceState returns a boolean if a field has been set.

### GetMemoryCapacityBytes

`func (o *HciNode) GetMemoryCapacityBytes() int64`

GetMemoryCapacityBytes returns the MemoryCapacityBytes field if non-nil, zero value otherwise.

### GetMemoryCapacityBytesOk

`func (o *HciNode) GetMemoryCapacityBytesOk() (*int64, bool)`

GetMemoryCapacityBytesOk returns a tuple with the MemoryCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacityBytes

`func (o *HciNode) SetMemoryCapacityBytes(v int64)`

SetMemoryCapacityBytes sets MemoryCapacityBytes field to given value.

### HasMemoryCapacityBytes

`func (o *HciNode) HasMemoryCapacityBytes() bool`

HasMemoryCapacityBytes returns a boolean if a field has been set.

### GetMemorySizeBytes

`func (o *HciNode) GetMemorySizeBytes() int64`

GetMemorySizeBytes returns the MemorySizeBytes field if non-nil, zero value otherwise.

### GetMemorySizeBytesOk

`func (o *HciNode) GetMemorySizeBytesOk() (*int64, bool)`

GetMemorySizeBytesOk returns a tuple with the MemorySizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeBytes

`func (o *HciNode) SetMemorySizeBytes(v int64)`

SetMemorySizeBytes sets MemorySizeBytes field to given value.

### HasMemorySizeBytes

`func (o *HciNode) HasMemorySizeBytes() bool`

HasMemorySizeBytes returns a boolean if a field has been set.

### GetMemoryUsageBytes

`func (o *HciNode) GetMemoryUsageBytes() int64`

GetMemoryUsageBytes returns the MemoryUsageBytes field if non-nil, zero value otherwise.

### GetMemoryUsageBytesOk

`func (o *HciNode) GetMemoryUsageBytesOk() (*int64, bool)`

GetMemoryUsageBytesOk returns a tuple with the MemoryUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageBytes

`func (o *HciNode) SetMemoryUsageBytes(v int64)`

SetMemoryUsageBytes sets MemoryUsageBytes field to given value.

### HasMemoryUsageBytes

`func (o *HciNode) HasMemoryUsageBytes() bool`

HasMemoryUsageBytes returns a boolean if a field has been set.

### GetNodeExtId

`func (o *HciNode) GetNodeExtId() string`

GetNodeExtId returns the NodeExtId field if non-nil, zero value otherwise.

### GetNodeExtIdOk

`func (o *HciNode) GetNodeExtIdOk() (*string, bool)`

GetNodeExtIdOk returns a tuple with the NodeExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExtId

`func (o *HciNode) SetNodeExtId(v string)`

SetNodeExtId sets NodeExtId field to given value.

### HasNodeExtId

`func (o *HciNode) HasNodeExtId() bool`

HasNodeExtId returns a boolean if a field has been set.

### GetNodeSerial

`func (o *HciNode) GetNodeSerial() string`

GetNodeSerial returns the NodeSerial field if non-nil, zero value otherwise.

### GetNodeSerialOk

`func (o *HciNode) GetNodeSerialOk() (*string, bool)`

GetNodeSerialOk returns a tuple with the NodeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSerial

`func (o *HciNode) SetNodeSerial(v string)`

SetNodeSerial sets NodeSerial field to given value.

### HasNodeSerial

`func (o *HciNode) HasNodeSerial() bool`

HasNodeSerial returns a boolean if a field has been set.

### GetNodeStatus

`func (o *HciNode) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *HciNode) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *HciNode) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *HciNode) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetNumberOfCpuCores

`func (o *HciNode) GetNumberOfCpuCores() int64`

GetNumberOfCpuCores returns the NumberOfCpuCores field if non-nil, zero value otherwise.

### GetNumberOfCpuCoresOk

`func (o *HciNode) GetNumberOfCpuCoresOk() (*int64, bool)`

GetNumberOfCpuCoresOk returns a tuple with the NumberOfCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCpuCores

`func (o *HciNode) SetNumberOfCpuCores(v int64)`

SetNumberOfCpuCores sets NumberOfCpuCores field to given value.

### HasNumberOfCpuCores

`func (o *HciNode) HasNumberOfCpuCores() bool`

HasNumberOfCpuCores returns a boolean if a field has been set.

### GetNumberOfCpuSockets

`func (o *HciNode) GetNumberOfCpuSockets() int64`

GetNumberOfCpuSockets returns the NumberOfCpuSockets field if non-nil, zero value otherwise.

### GetNumberOfCpuSocketsOk

`func (o *HciNode) GetNumberOfCpuSocketsOk() (*int64, bool)`

GetNumberOfCpuSocketsOk returns a tuple with the NumberOfCpuSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCpuSockets

`func (o *HciNode) SetNumberOfCpuSockets(v int64)`

SetNumberOfCpuSockets sets NumberOfCpuSockets field to given value.

### HasNumberOfCpuSockets

`func (o *HciNode) HasNumberOfCpuSockets() bool`

HasNumberOfCpuSockets returns a boolean if a field has been set.

### GetNumberOfCpuThreads

`func (o *HciNode) GetNumberOfCpuThreads() int64`

GetNumberOfCpuThreads returns the NumberOfCpuThreads field if non-nil, zero value otherwise.

### GetNumberOfCpuThreadsOk

`func (o *HciNode) GetNumberOfCpuThreadsOk() (*int64, bool)`

GetNumberOfCpuThreadsOk returns a tuple with the NumberOfCpuThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCpuThreads

`func (o *HciNode) SetNumberOfCpuThreads(v int64)`

SetNumberOfCpuThreads sets NumberOfCpuThreads field to given value.

### HasNumberOfCpuThreads

`func (o *HciNode) HasNumberOfCpuThreads() bool`

HasNumberOfCpuThreads returns a boolean if a field has been set.

### GetRebootPending

`func (o *HciNode) GetRebootPending() bool`

GetRebootPending returns the RebootPending field if non-nil, zero value otherwise.

### GetRebootPendingOk

`func (o *HciNode) GetRebootPendingOk() (*bool, bool)`

GetRebootPendingOk returns a tuple with the RebootPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootPending

`func (o *HciNode) SetRebootPending(v bool)`

SetRebootPending sets RebootPending field to given value.

### HasRebootPending

`func (o *HciNode) HasRebootPending() bool`

HasRebootPending returns a boolean if a field has been set.

### GetStorageCapacityBytes

`func (o *HciNode) GetStorageCapacityBytes() int64`

GetStorageCapacityBytes returns the StorageCapacityBytes field if non-nil, zero value otherwise.

### GetStorageCapacityBytesOk

`func (o *HciNode) GetStorageCapacityBytesOk() (*int64, bool)`

GetStorageCapacityBytesOk returns a tuple with the StorageCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacityBytes

`func (o *HciNode) SetStorageCapacityBytes(v int64)`

SetStorageCapacityBytes sets StorageCapacityBytes field to given value.

### HasStorageCapacityBytes

`func (o *HciNode) HasStorageCapacityBytes() bool`

HasStorageCapacityBytes returns a boolean if a field has been set.

### GetStorageUsageBytes

`func (o *HciNode) GetStorageUsageBytes() int64`

GetStorageUsageBytes returns the StorageUsageBytes field if non-nil, zero value otherwise.

### GetStorageUsageBytesOk

`func (o *HciNode) GetStorageUsageBytesOk() (*int64, bool)`

GetStorageUsageBytesOk returns a tuple with the StorageUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsageBytes

`func (o *HciNode) SetStorageUsageBytes(v int64)`

SetStorageUsageBytes sets StorageUsageBytes field to given value.

### HasStorageUsageBytes

`func (o *HciNode) HasStorageUsageBytes() bool`

HasStorageUsageBytes returns a boolean if a field has been set.

### GetCluster

`func (o *HciNode) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciNode) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciNode) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciNode) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciNode) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciNode) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDisks

`func (o *HciNode) GetDisks() []HciDiskRelationship`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *HciNode) GetDisksOk() (*[]HciDiskRelationship, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *HciNode) SetDisks(v []HciDiskRelationship)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *HciNode) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *HciNode) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *HciNode) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetGpus

`func (o *HciNode) GetGpus() []HciGpuRelationship`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *HciNode) GetGpusOk() (*[]HciGpuRelationship, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *HciNode) SetGpus(v []HciGpuRelationship)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *HciNode) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### SetGpusNil

`func (o *HciNode) SetGpusNil(b bool)`

 SetGpusNil sets the value for Gpus to be an explicit nil

### UnsetGpus
`func (o *HciNode) UnsetGpus()`

UnsetGpus ensures that no value is present for Gpus, not even an explicit nil
### GetPhysicalServer

`func (o *HciNode) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *HciNode) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *HciNode) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *HciNode) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.

### SetPhysicalServerNil

`func (o *HciNode) SetPhysicalServerNil(b bool)`

 SetPhysicalServerNil sets the value for PhysicalServer to be an explicit nil

### UnsetPhysicalServer
`func (o *HciNode) UnsetPhysicalServer()`

UnsetPhysicalServer ensures that no value is present for PhysicalServer, not even an explicit nil
### GetRegisteredDevice

`func (o *HciNode) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciNode) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciNode) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciNode) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciNode) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciNode) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVms

`func (o *HciNode) GetVms() []HciBaseVmRelationship`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *HciNode) GetVmsOk() (*[]HciBaseVmRelationship, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *HciNode) SetVms(v []HciBaseVmRelationship)`

SetVms sets Vms field to given value.

### HasVms

`func (o *HciNode) HasVms() bool`

HasVms returns a boolean if a field has been set.

### SetVmsNil

`func (o *HciNode) SetVmsNil(b bool)`

 SetVmsNil sets the value for Vms to be an explicit nil

### UnsetVms
`func (o *HciNode) UnsetVms()`

UnsetVms ensures that no value is present for Vms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


