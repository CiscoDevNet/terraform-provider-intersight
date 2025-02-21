# ApplianceClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ClusterInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ClusterInfo"]
**BuildType** | Pointer to **string** | The build type of the Intersight Virtual Appliance. | [optional] 
**DeploymentSize** | Pointer to **string** | The deployment size of the node requiring to join cluster. | [optional] 
**Gateway** | Pointer to **string** | Default gateway configured on the peer node. | [optional] 
**Hostip** | Pointer to **string** | Publicly accessible IP of the peer node. | [optional] 
**Hostname** | Pointer to **string** | Publicly accessible FQDN of the peer node. | [optional] 
**InstallerVersion** | Pointer to **string** | Installer version used to install on peer node. | [optional] 
**NodeId** | Pointer to **int64** | System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance starting with a value of 0. | [optional] [readonly] 
**NodeType** | Pointer to **string** | The node type of Intersight Virtual Appliance. * &#x60;standalone&#x60; - Single Node Intersight Virtual Appliance. * &#x60;management&#x60; - Management node type when Intersight Virtual Appliance is running as management-worker deployment. * &#x60;hamanagement&#x60; - Management node type when Intersight Virtual Appliance is running as multi node HA deployment. * &#x60;metrics&#x60; - Metrics node when Intersight Virtual Appliance is running management-metrics node. | [optional] [default to "standalone"]
**PartitionDatabase** | Pointer to **int64** | The partition size for /opt/database of this node. | [optional] 
**PartitionFileCisco** | Pointer to **int64** | The partition size for /Cisco of this node. | [optional] 
**PartitionOptData** | Pointer to **int64** | The partition size for /opt/cisco/data of this node. | [optional] 
**PartitionOptKafka** | Pointer to **int64** | The partition size for /opt/kafka of this node. | [optional] 
**PartitionOptMongo** | Pointer to **int64** | The partition size for /opt/mongodb of this node. | [optional] 
**PartitionVarLibDocker** | Pointer to **int64** | The partition size for /var/lib/docker of this node. | [optional] 
**PartitionVarLog** | Pointer to **int64** | The partition size for /var of this node. | [optional] 
**Peerkey** | Pointer to **string** | The public key of peer host. | [optional] 
**Responsekey** | Pointer to **string** | Public key returned to the client. | [optional] 
**Status** | Pointer to **string** | The status of the cluster join process. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. * &#x60;WorkerNodeInstInProgress&#x60; - The worker node installation is in progress. * &#x60;WorkerNodeInstSuccess&#x60; - The worker node installation succeeded. * &#x60;WorkerNodeInstFailed&#x60; - The worker node installation failed. | [optional] [readonly] [default to "Unknown"]
**Subnetmask** | Pointer to **string** | Subnet Mask of the peer node. | [optional] 
**Uuid** | Pointer to **string** | The UUID of the peer appliance. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceClusterInfo

`func NewApplianceClusterInfo(classId string, objectType string, ) *ApplianceClusterInfo`

NewApplianceClusterInfo instantiates a new ApplianceClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceClusterInfoWithDefaults

`func NewApplianceClusterInfoWithDefaults() *ApplianceClusterInfo`

NewApplianceClusterInfoWithDefaults instantiates a new ApplianceClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceClusterInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceClusterInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceClusterInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceClusterInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceClusterInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceClusterInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBuildType

`func (o *ApplianceClusterInfo) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *ApplianceClusterInfo) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *ApplianceClusterInfo) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *ApplianceClusterInfo) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetDeploymentSize

`func (o *ApplianceClusterInfo) GetDeploymentSize() string`

GetDeploymentSize returns the DeploymentSize field if non-nil, zero value otherwise.

### GetDeploymentSizeOk

`func (o *ApplianceClusterInfo) GetDeploymentSizeOk() (*string, bool)`

GetDeploymentSizeOk returns a tuple with the DeploymentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSize

`func (o *ApplianceClusterInfo) SetDeploymentSize(v string)`

SetDeploymentSize sets DeploymentSize field to given value.

### HasDeploymentSize

`func (o *ApplianceClusterInfo) HasDeploymentSize() bool`

HasDeploymentSize returns a boolean if a field has been set.

### GetGateway

`func (o *ApplianceClusterInfo) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ApplianceClusterInfo) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ApplianceClusterInfo) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ApplianceClusterInfo) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostip

`func (o *ApplianceClusterInfo) GetHostip() string`

GetHostip returns the Hostip field if non-nil, zero value otherwise.

### GetHostipOk

`func (o *ApplianceClusterInfo) GetHostipOk() (*string, bool)`

GetHostipOk returns a tuple with the Hostip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostip

`func (o *ApplianceClusterInfo) SetHostip(v string)`

SetHostip sets Hostip field to given value.

### HasHostip

`func (o *ApplianceClusterInfo) HasHostip() bool`

HasHostip returns a boolean if a field has been set.

### GetHostname

`func (o *ApplianceClusterInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceClusterInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceClusterInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceClusterInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInstallerVersion

`func (o *ApplianceClusterInfo) GetInstallerVersion() string`

GetInstallerVersion returns the InstallerVersion field if non-nil, zero value otherwise.

### GetInstallerVersionOk

`func (o *ApplianceClusterInfo) GetInstallerVersionOk() (*string, bool)`

GetInstallerVersionOk returns a tuple with the InstallerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerVersion

`func (o *ApplianceClusterInfo) SetInstallerVersion(v string)`

SetInstallerVersion sets InstallerVersion field to given value.

### HasInstallerVersion

`func (o *ApplianceClusterInfo) HasInstallerVersion() bool`

HasInstallerVersion returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceClusterInfo) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceClusterInfo) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceClusterInfo) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceClusterInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeType

`func (o *ApplianceClusterInfo) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ApplianceClusterInfo) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ApplianceClusterInfo) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ApplianceClusterInfo) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetPartitionDatabase

`func (o *ApplianceClusterInfo) GetPartitionDatabase() int64`

GetPartitionDatabase returns the PartitionDatabase field if non-nil, zero value otherwise.

### GetPartitionDatabaseOk

`func (o *ApplianceClusterInfo) GetPartitionDatabaseOk() (*int64, bool)`

GetPartitionDatabaseOk returns a tuple with the PartitionDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionDatabase

`func (o *ApplianceClusterInfo) SetPartitionDatabase(v int64)`

SetPartitionDatabase sets PartitionDatabase field to given value.

### HasPartitionDatabase

`func (o *ApplianceClusterInfo) HasPartitionDatabase() bool`

HasPartitionDatabase returns a boolean if a field has been set.

### GetPartitionFileCisco

`func (o *ApplianceClusterInfo) GetPartitionFileCisco() int64`

GetPartitionFileCisco returns the PartitionFileCisco field if non-nil, zero value otherwise.

### GetPartitionFileCiscoOk

`func (o *ApplianceClusterInfo) GetPartitionFileCiscoOk() (*int64, bool)`

GetPartitionFileCiscoOk returns a tuple with the PartitionFileCisco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionFileCisco

`func (o *ApplianceClusterInfo) SetPartitionFileCisco(v int64)`

SetPartitionFileCisco sets PartitionFileCisco field to given value.

### HasPartitionFileCisco

`func (o *ApplianceClusterInfo) HasPartitionFileCisco() bool`

HasPartitionFileCisco returns a boolean if a field has been set.

### GetPartitionOptData

`func (o *ApplianceClusterInfo) GetPartitionOptData() int64`

GetPartitionOptData returns the PartitionOptData field if non-nil, zero value otherwise.

### GetPartitionOptDataOk

`func (o *ApplianceClusterInfo) GetPartitionOptDataOk() (*int64, bool)`

GetPartitionOptDataOk returns a tuple with the PartitionOptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionOptData

`func (o *ApplianceClusterInfo) SetPartitionOptData(v int64)`

SetPartitionOptData sets PartitionOptData field to given value.

### HasPartitionOptData

`func (o *ApplianceClusterInfo) HasPartitionOptData() bool`

HasPartitionOptData returns a boolean if a field has been set.

### GetPartitionOptKafka

`func (o *ApplianceClusterInfo) GetPartitionOptKafka() int64`

GetPartitionOptKafka returns the PartitionOptKafka field if non-nil, zero value otherwise.

### GetPartitionOptKafkaOk

`func (o *ApplianceClusterInfo) GetPartitionOptKafkaOk() (*int64, bool)`

GetPartitionOptKafkaOk returns a tuple with the PartitionOptKafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionOptKafka

`func (o *ApplianceClusterInfo) SetPartitionOptKafka(v int64)`

SetPartitionOptKafka sets PartitionOptKafka field to given value.

### HasPartitionOptKafka

`func (o *ApplianceClusterInfo) HasPartitionOptKafka() bool`

HasPartitionOptKafka returns a boolean if a field has been set.

### GetPartitionOptMongo

`func (o *ApplianceClusterInfo) GetPartitionOptMongo() int64`

GetPartitionOptMongo returns the PartitionOptMongo field if non-nil, zero value otherwise.

### GetPartitionOptMongoOk

`func (o *ApplianceClusterInfo) GetPartitionOptMongoOk() (*int64, bool)`

GetPartitionOptMongoOk returns a tuple with the PartitionOptMongo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionOptMongo

`func (o *ApplianceClusterInfo) SetPartitionOptMongo(v int64)`

SetPartitionOptMongo sets PartitionOptMongo field to given value.

### HasPartitionOptMongo

`func (o *ApplianceClusterInfo) HasPartitionOptMongo() bool`

HasPartitionOptMongo returns a boolean if a field has been set.

### GetPartitionVarLibDocker

`func (o *ApplianceClusterInfo) GetPartitionVarLibDocker() int64`

GetPartitionVarLibDocker returns the PartitionVarLibDocker field if non-nil, zero value otherwise.

### GetPartitionVarLibDockerOk

`func (o *ApplianceClusterInfo) GetPartitionVarLibDockerOk() (*int64, bool)`

GetPartitionVarLibDockerOk returns a tuple with the PartitionVarLibDocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionVarLibDocker

`func (o *ApplianceClusterInfo) SetPartitionVarLibDocker(v int64)`

SetPartitionVarLibDocker sets PartitionVarLibDocker field to given value.

### HasPartitionVarLibDocker

`func (o *ApplianceClusterInfo) HasPartitionVarLibDocker() bool`

HasPartitionVarLibDocker returns a boolean if a field has been set.

### GetPartitionVarLog

`func (o *ApplianceClusterInfo) GetPartitionVarLog() int64`

GetPartitionVarLog returns the PartitionVarLog field if non-nil, zero value otherwise.

### GetPartitionVarLogOk

`func (o *ApplianceClusterInfo) GetPartitionVarLogOk() (*int64, bool)`

GetPartitionVarLogOk returns a tuple with the PartitionVarLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionVarLog

`func (o *ApplianceClusterInfo) SetPartitionVarLog(v int64)`

SetPartitionVarLog sets PartitionVarLog field to given value.

### HasPartitionVarLog

`func (o *ApplianceClusterInfo) HasPartitionVarLog() bool`

HasPartitionVarLog returns a boolean if a field has been set.

### GetPeerkey

`func (o *ApplianceClusterInfo) GetPeerkey() string`

GetPeerkey returns the Peerkey field if non-nil, zero value otherwise.

### GetPeerkeyOk

`func (o *ApplianceClusterInfo) GetPeerkeyOk() (*string, bool)`

GetPeerkeyOk returns a tuple with the Peerkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerkey

`func (o *ApplianceClusterInfo) SetPeerkey(v string)`

SetPeerkey sets Peerkey field to given value.

### HasPeerkey

`func (o *ApplianceClusterInfo) HasPeerkey() bool`

HasPeerkey returns a boolean if a field has been set.

### GetResponsekey

`func (o *ApplianceClusterInfo) GetResponsekey() string`

GetResponsekey returns the Responsekey field if non-nil, zero value otherwise.

### GetResponsekeyOk

`func (o *ApplianceClusterInfo) GetResponsekeyOk() (*string, bool)`

GetResponsekeyOk returns a tuple with the Responsekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsekey

`func (o *ApplianceClusterInfo) SetResponsekey(v string)`

SetResponsekey sets Responsekey field to given value.

### HasResponsekey

`func (o *ApplianceClusterInfo) HasResponsekey() bool`

HasResponsekey returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceClusterInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceClusterInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceClusterInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceClusterInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubnetmask

`func (o *ApplianceClusterInfo) GetSubnetmask() string`

GetSubnetmask returns the Subnetmask field if non-nil, zero value otherwise.

### GetSubnetmaskOk

`func (o *ApplianceClusterInfo) GetSubnetmaskOk() (*string, bool)`

GetSubnetmaskOk returns a tuple with the Subnetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetmask

`func (o *ApplianceClusterInfo) SetSubnetmask(v string)`

SetSubnetmask sets Subnetmask field to given value.

### HasSubnetmask

`func (o *ApplianceClusterInfo) HasSubnetmask() bool`

HasSubnetmask returns a boolean if a field has been set.

### GetUuid

`func (o *ApplianceClusterInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApplianceClusterInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApplianceClusterInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApplianceClusterInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceClusterInfo) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceClusterInfo) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceClusterInfo) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceClusterInfo) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceClusterInfo) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceClusterInfo) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


