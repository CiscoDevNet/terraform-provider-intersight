# ApplianceNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NodeInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NodeInfo"]
**Hostname** | Pointer to **string** | Cluster node&#39;s FQDN or IP address. | [optional] [readonly] 
**NodeId** | Pointer to **int64** | System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. | [optional] [readonly] 
**NodeIpV4Config** | Pointer to [**NullableCommIpV4Interface**](CommIpV4Interface.md) |  | [optional] 
**NodeIpV6Config** | Pointer to [**NullableCommIpV6Interface**](CommIpV6Interface.md) |  | [optional] 
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance node. * &#x60;Unknown&#x60; - Operational status of the Intersight Appliance entity is Unknown. * &#x60;Operational&#x60; - Operational status of the Intersight Appliance entity is Operational. * &#x60;Impaired&#x60; - Operational status of the Intersight Appliance entity is Impaired. * &#x60;AttentionNeeded&#x60; - Operational status of the Intersight Appliance entity is AttentionNeeded. | [optional] [readonly] [default to "Unknown"]

## Methods

### NewApplianceNodeInfo

`func NewApplianceNodeInfo(classId string, objectType string, ) *ApplianceNodeInfo`

NewApplianceNodeInfo instantiates a new ApplianceNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeInfoWithDefaults

`func NewApplianceNodeInfoWithDefaults() *ApplianceNodeInfo`

NewApplianceNodeInfoWithDefaults instantiates a new ApplianceNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *ApplianceNodeInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceNodeInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceNodeInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceNodeInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceNodeInfo) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceNodeInfo) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceNodeInfo) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceNodeInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeIpV4Config

`func (o *ApplianceNodeInfo) GetNodeIpV4Config() CommIpV4Interface`

GetNodeIpV4Config returns the NodeIpV4Config field if non-nil, zero value otherwise.

### GetNodeIpV4ConfigOk

`func (o *ApplianceNodeInfo) GetNodeIpV4ConfigOk() (*CommIpV4Interface, bool)`

GetNodeIpV4ConfigOk returns a tuple with the NodeIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpV4Config

`func (o *ApplianceNodeInfo) SetNodeIpV4Config(v CommIpV4Interface)`

SetNodeIpV4Config sets NodeIpV4Config field to given value.

### HasNodeIpV4Config

`func (o *ApplianceNodeInfo) HasNodeIpV4Config() bool`

HasNodeIpV4Config returns a boolean if a field has been set.

### SetNodeIpV4ConfigNil

`func (o *ApplianceNodeInfo) SetNodeIpV4ConfigNil(b bool)`

 SetNodeIpV4ConfigNil sets the value for NodeIpV4Config to be an explicit nil

### UnsetNodeIpV4Config
`func (o *ApplianceNodeInfo) UnsetNodeIpV4Config()`

UnsetNodeIpV4Config ensures that no value is present for NodeIpV4Config, not even an explicit nil
### GetNodeIpV6Config

`func (o *ApplianceNodeInfo) GetNodeIpV6Config() CommIpV6Interface`

GetNodeIpV6Config returns the NodeIpV6Config field if non-nil, zero value otherwise.

### GetNodeIpV6ConfigOk

`func (o *ApplianceNodeInfo) GetNodeIpV6ConfigOk() (*CommIpV6Interface, bool)`

GetNodeIpV6ConfigOk returns a tuple with the NodeIpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpV6Config

`func (o *ApplianceNodeInfo) SetNodeIpV6Config(v CommIpV6Interface)`

SetNodeIpV6Config sets NodeIpV6Config field to given value.

### HasNodeIpV6Config

`func (o *ApplianceNodeInfo) HasNodeIpV6Config() bool`

HasNodeIpV6Config returns a boolean if a field has been set.

### SetNodeIpV6ConfigNil

`func (o *ApplianceNodeInfo) SetNodeIpV6ConfigNil(b bool)`

 SetNodeIpV6ConfigNil sets the value for NodeIpV6Config to be an explicit nil

### UnsetNodeIpV6Config
`func (o *ApplianceNodeInfo) UnsetNodeIpV6Config()`

UnsetNodeIpV6Config ensures that no value is present for NodeIpV6Config, not even an explicit nil
### GetOperationalStatus

`func (o *ApplianceNodeInfo) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceNodeInfo) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceNodeInfo) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceNodeInfo) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


