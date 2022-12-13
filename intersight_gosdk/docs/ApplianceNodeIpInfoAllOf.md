# ApplianceNodeIpInfoAllOf

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
**Status** | Pointer to **string** | Status of the cluster node. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. | [optional] [readonly] [default to "Unknown"]

## Methods

### NewApplianceNodeIpInfoAllOf

`func NewApplianceNodeIpInfoAllOf(classId string, objectType string, ) *ApplianceNodeIpInfoAllOf`

NewApplianceNodeIpInfoAllOf instantiates a new ApplianceNodeIpInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeIpInfoAllOfWithDefaults

`func NewApplianceNodeIpInfoAllOfWithDefaults() *ApplianceNodeIpInfoAllOf`

NewApplianceNodeIpInfoAllOfWithDefaults instantiates a new ApplianceNodeIpInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeIpInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeIpInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeIpInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeIpInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeIpInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeIpInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGateway

`func (o *ApplianceNodeIpInfoAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ApplianceNodeIpInfoAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ApplianceNodeIpInfoAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ApplianceNodeIpInfoAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostname

`func (o *ApplianceNodeIpInfoAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceNodeIpInfoAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceNodeIpInfoAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceNodeIpInfoAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddress

`func (o *ApplianceNodeIpInfoAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApplianceNodeIpInfoAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApplianceNodeIpInfoAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ApplianceNodeIpInfoAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetmask

`func (o *ApplianceNodeIpInfoAllOf) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ApplianceNodeIpInfoAllOf) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ApplianceNodeIpInfoAllOf) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ApplianceNodeIpInfoAllOf) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceNodeIpInfoAllOf) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceNodeIpInfoAllOf) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceNodeIpInfoAllOf) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceNodeIpInfoAllOf) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeMoid

`func (o *ApplianceNodeIpInfoAllOf) GetNodeMoid() string`

GetNodeMoid returns the NodeMoid field if non-nil, zero value otherwise.

### GetNodeMoidOk

`func (o *ApplianceNodeIpInfoAllOf) GetNodeMoidOk() (*string, bool)`

GetNodeMoidOk returns a tuple with the NodeMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMoid

`func (o *ApplianceNodeIpInfoAllOf) SetNodeMoid(v string)`

SetNodeMoid sets NodeMoid field to given value.

### HasNodeMoid

`func (o *ApplianceNodeIpInfoAllOf) HasNodeMoid() bool`

HasNodeMoid returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceNodeIpInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceNodeIpInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceNodeIpInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceNodeIpInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


