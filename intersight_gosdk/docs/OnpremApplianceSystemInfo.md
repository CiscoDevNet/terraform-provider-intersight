# OnpremApplianceSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.ApplianceSystemInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.ApplianceSystemInfo"]
**DeploymentSize** | Pointer to **string** | Current running deployment size for the Intersight Appliance node. Eg. tiny, small, medium, large etc. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Cluster node&#39;s FQDN or IP address. | [optional] [readonly] 
**NodeId** | Pointer to **int64** | System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. | [optional] [readonly] 
**NodeIpV4Config** | Pointer to [**NullableCommIpV4Interface**](CommIpV4Interface.md) |  | [optional] 
**NodeIpV6Config** | Pointer to [**NullableCommIpV6Interface**](CommIpV6Interface.md) |  | [optional] 
**NodeType** | Pointer to **string** | The node type of Intersight Virtual Appliance. * &#x60;standalone&#x60; - Single Node Intersight Virtual Appliance. * &#x60;management&#x60; - Management node type when Intersight Virtual Appliance is running as management-worker deployment. * &#x60;hamanagement&#x60; - Management node type when Intersight Virtual Appliance is running as multi node HA deployment. * &#x60;metrics&#x60; - Metrics node when Intersight Virtual Appliance is running management-metrics node. | [optional] [readonly] [default to "standalone"]
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance node. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. * &#x60;WorkerNodeInstInProgress&#x60; - The worker node installation is in progress. * &#x60;WorkerNodeInstSuccess&#x60; - The worker node installation succeeded. * &#x60;WorkerNodeInstFailed&#x60; - The worker node installation failed. | [optional] [readonly] [default to "Unknown"]
**SerialId** | Pointer to **string** | SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance. | [optional] [readonly] 
**Version** | Pointer to **string** | Current software running version of the Intersight Appliance. | [optional] [readonly] 

## Methods

### NewOnpremApplianceSystemInfo

`func NewOnpremApplianceSystemInfo(classId string, objectType string, ) *OnpremApplianceSystemInfo`

NewOnpremApplianceSystemInfo instantiates a new OnpremApplianceSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremApplianceSystemInfoWithDefaults

`func NewOnpremApplianceSystemInfoWithDefaults() *OnpremApplianceSystemInfo`

NewOnpremApplianceSystemInfoWithDefaults instantiates a new OnpremApplianceSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremApplianceSystemInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremApplianceSystemInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremApplianceSystemInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremApplianceSystemInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremApplianceSystemInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremApplianceSystemInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeploymentSize

`func (o *OnpremApplianceSystemInfo) GetDeploymentSize() string`

GetDeploymentSize returns the DeploymentSize field if non-nil, zero value otherwise.

### GetDeploymentSizeOk

`func (o *OnpremApplianceSystemInfo) GetDeploymentSizeOk() (*string, bool)`

GetDeploymentSizeOk returns a tuple with the DeploymentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSize

`func (o *OnpremApplianceSystemInfo) SetDeploymentSize(v string)`

SetDeploymentSize sets DeploymentSize field to given value.

### HasDeploymentSize

`func (o *OnpremApplianceSystemInfo) HasDeploymentSize() bool`

HasDeploymentSize returns a boolean if a field has been set.

### GetHostname

`func (o *OnpremApplianceSystemInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *OnpremApplianceSystemInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *OnpremApplianceSystemInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *OnpremApplianceSystemInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *OnpremApplianceSystemInfo) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *OnpremApplianceSystemInfo) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *OnpremApplianceSystemInfo) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *OnpremApplianceSystemInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeIpV4Config

`func (o *OnpremApplianceSystemInfo) GetNodeIpV4Config() CommIpV4Interface`

GetNodeIpV4Config returns the NodeIpV4Config field if non-nil, zero value otherwise.

### GetNodeIpV4ConfigOk

`func (o *OnpremApplianceSystemInfo) GetNodeIpV4ConfigOk() (*CommIpV4Interface, bool)`

GetNodeIpV4ConfigOk returns a tuple with the NodeIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpV4Config

`func (o *OnpremApplianceSystemInfo) SetNodeIpV4Config(v CommIpV4Interface)`

SetNodeIpV4Config sets NodeIpV4Config field to given value.

### HasNodeIpV4Config

`func (o *OnpremApplianceSystemInfo) HasNodeIpV4Config() bool`

HasNodeIpV4Config returns a boolean if a field has been set.

### SetNodeIpV4ConfigNil

`func (o *OnpremApplianceSystemInfo) SetNodeIpV4ConfigNil(b bool)`

 SetNodeIpV4ConfigNil sets the value for NodeIpV4Config to be an explicit nil

### UnsetNodeIpV4Config
`func (o *OnpremApplianceSystemInfo) UnsetNodeIpV4Config()`

UnsetNodeIpV4Config ensures that no value is present for NodeIpV4Config, not even an explicit nil
### GetNodeIpV6Config

`func (o *OnpremApplianceSystemInfo) GetNodeIpV6Config() CommIpV6Interface`

GetNodeIpV6Config returns the NodeIpV6Config field if non-nil, zero value otherwise.

### GetNodeIpV6ConfigOk

`func (o *OnpremApplianceSystemInfo) GetNodeIpV6ConfigOk() (*CommIpV6Interface, bool)`

GetNodeIpV6ConfigOk returns a tuple with the NodeIpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpV6Config

`func (o *OnpremApplianceSystemInfo) SetNodeIpV6Config(v CommIpV6Interface)`

SetNodeIpV6Config sets NodeIpV6Config field to given value.

### HasNodeIpV6Config

`func (o *OnpremApplianceSystemInfo) HasNodeIpV6Config() bool`

HasNodeIpV6Config returns a boolean if a field has been set.

### SetNodeIpV6ConfigNil

`func (o *OnpremApplianceSystemInfo) SetNodeIpV6ConfigNil(b bool)`

 SetNodeIpV6ConfigNil sets the value for NodeIpV6Config to be an explicit nil

### UnsetNodeIpV6Config
`func (o *OnpremApplianceSystemInfo) UnsetNodeIpV6Config()`

UnsetNodeIpV6Config ensures that no value is present for NodeIpV6Config, not even an explicit nil
### GetNodeType

`func (o *OnpremApplianceSystemInfo) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *OnpremApplianceSystemInfo) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *OnpremApplianceSystemInfo) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *OnpremApplianceSystemInfo) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *OnpremApplianceSystemInfo) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *OnpremApplianceSystemInfo) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *OnpremApplianceSystemInfo) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *OnpremApplianceSystemInfo) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetSerialId

`func (o *OnpremApplianceSystemInfo) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *OnpremApplianceSystemInfo) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *OnpremApplianceSystemInfo) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *OnpremApplianceSystemInfo) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetVersion

`func (o *OnpremApplianceSystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OnpremApplianceSystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OnpremApplianceSystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OnpremApplianceSystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


