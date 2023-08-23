# ApplianceNodeIpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NodeIpInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NodeIpInfo"]
**Gateway** | Pointer to **string** | Gateway ip address of the cluster node. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Hostname of the cluster node. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | Ip address of the cluster node. | [optional] [readonly] 
**Netmask** | Pointer to **string** | Netmask of the cluster node. | [optional] [readonly] 
**NodeId** | Pointer to **int64** | Id number of the cluster node. | [optional] [readonly] 
**NodeMoid** | Pointer to **string** | Moid of the corresponding appliance.ClusterInfo or appliance.NodeInfo mo. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the cluster node. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. | [optional] [readonly] [default to "Unknown"]

## Methods

### NewApplianceNodeIpInfo

`func NewApplianceNodeIpInfo(classId string, objectType string, ) *ApplianceNodeIpInfo`

NewApplianceNodeIpInfo instantiates a new ApplianceNodeIpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeIpInfoWithDefaults

`func NewApplianceNodeIpInfoWithDefaults() *ApplianceNodeIpInfo`

NewApplianceNodeIpInfoWithDefaults instantiates a new ApplianceNodeIpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeIpInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeIpInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeIpInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeIpInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeIpInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeIpInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGateway

`func (o *ApplianceNodeIpInfo) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ApplianceNodeIpInfo) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ApplianceNodeIpInfo) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ApplianceNodeIpInfo) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostname

`func (o *ApplianceNodeIpInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceNodeIpInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceNodeIpInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceNodeIpInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddress

`func (o *ApplianceNodeIpInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApplianceNodeIpInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApplianceNodeIpInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ApplianceNodeIpInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetmask

`func (o *ApplianceNodeIpInfo) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ApplianceNodeIpInfo) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ApplianceNodeIpInfo) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ApplianceNodeIpInfo) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceNodeIpInfo) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceNodeIpInfo) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceNodeIpInfo) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceNodeIpInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeMoid

`func (o *ApplianceNodeIpInfo) GetNodeMoid() string`

GetNodeMoid returns the NodeMoid field if non-nil, zero value otherwise.

### GetNodeMoidOk

`func (o *ApplianceNodeIpInfo) GetNodeMoidOk() (*string, bool)`

GetNodeMoidOk returns a tuple with the NodeMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMoid

`func (o *ApplianceNodeIpInfo) SetNodeMoid(v string)`

SetNodeMoid sets NodeMoid field to given value.

### HasNodeMoid

`func (o *ApplianceNodeIpInfo) HasNodeMoid() bool`

HasNodeMoid returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceNodeIpInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceNodeIpInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceNodeIpInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceNodeIpInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


