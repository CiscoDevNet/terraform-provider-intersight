# ApplianceClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ClusterInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ClusterInfo"]
**DeploymentSize** | Pointer to **string** | The deployment size of the node requiring to join cluster. | [optional] 
**Gateway** | Pointer to **string** | Default gateway configured on the peer node. | [optional] 
**Hostip** | Pointer to **string** | Publicly accessible IP of the peer node. | [optional] 
**Hostname** | Pointer to **string** | Publicly accessible FQDN of the peer node. | [optional] 
**Peerkey** | Pointer to **string** | The public key of peer host. | [optional] 
**Responsekey** | Pointer to **string** | Public key returned to the client. | [optional] 
**Status** | Pointer to **string** | The status of the cluster join process. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. | [optional] [readonly] [default to "Unknown"]
**Subnetmask** | Pointer to **string** | Subnet Mask of the peer node. | [optional] 
**Uuid** | Pointer to **string** | The UUID of the peer appliance. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


