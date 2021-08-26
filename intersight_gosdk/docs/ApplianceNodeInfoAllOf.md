# ApplianceNodeInfoAllOf

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

### NewApplianceNodeInfoAllOf

`func NewApplianceNodeInfoAllOf(classId string, objectType string, ) *ApplianceNodeInfoAllOf`

NewApplianceNodeInfoAllOf instantiates a new ApplianceNodeInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeInfoAllOfWithDefaults

`func NewApplianceNodeInfoAllOfWithDefaults() *ApplianceNodeInfoAllOf`

NewApplianceNodeInfoAllOfWithDefaults instantiates a new ApplianceNodeInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *ApplianceNodeInfoAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceNodeInfoAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceNodeInfoAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceNodeInfoAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceNodeInfoAllOf) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceNodeInfoAllOf) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceNodeInfoAllOf) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceNodeInfoAllOf) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeIpV4Config

`func (o *ApplianceNodeInfoAllOf) GetNodeIpV4Config() CommIpV4Interface`

GetNodeIpV4Config returns the NodeIpV4Config field if non-nil, zero value otherwise.

### GetNodeIpV4ConfigOk

`func (o *ApplianceNodeInfoAllOf) GetNodeIpV4ConfigOk() (*CommIpV4Interface, bool)`

GetNodeIpV4ConfigOk returns a tuple with the NodeIpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpV4Config

`func (o *ApplianceNodeInfoAllOf) SetNodeIpV4Config(v CommIpV4Interface)`

SetNodeIpV4Config sets NodeIpV4Config field to given value.

### HasNodeIpV4Config

`func (o *ApplianceNodeInfoAllOf) HasNodeIpV4Config() bool`

HasNodeIpV4Config returns a boolean if a field has been set.

### SetNodeIpV4ConfigNil

`func (o *ApplianceNodeInfoAllOf) SetNodeIpV4ConfigNil(b bool)`

 SetNodeIpV4ConfigNil sets the value for NodeIpV4Config to be an explicit nil

### UnsetNodeIpV4Config
`func (o *ApplianceNodeInfoAllOf) UnsetNodeIpV4Config()`

UnsetNodeIpV4Config ensures that no value is present for NodeIpV4Config, not even an explicit nil
### GetNodeIpV6Config

`func (o *ApplianceNodeInfoAllOf) GetNodeIpV6Config() CommIpV6Interface`

GetNodeIpV6Config returns the NodeIpV6Config field if non-nil, zero value otherwise.

### GetNodeIpV6ConfigOk

`func (o *ApplianceNodeInfoAllOf) GetNodeIpV6ConfigOk() (*CommIpV6Interface, bool)`

GetNodeIpV6ConfigOk returns a tuple with the NodeIpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpV6Config

`func (o *ApplianceNodeInfoAllOf) SetNodeIpV6Config(v CommIpV6Interface)`

SetNodeIpV6Config sets NodeIpV6Config field to given value.

### HasNodeIpV6Config

`func (o *ApplianceNodeInfoAllOf) HasNodeIpV6Config() bool`

HasNodeIpV6Config returns a boolean if a field has been set.

### SetNodeIpV6ConfigNil

`func (o *ApplianceNodeInfoAllOf) SetNodeIpV6ConfigNil(b bool)`

 SetNodeIpV6ConfigNil sets the value for NodeIpV6Config to be an explicit nil

### UnsetNodeIpV6Config
`func (o *ApplianceNodeInfoAllOf) UnsetNodeIpV6Config()`

UnsetNodeIpV6Config ensures that no value is present for NodeIpV6Config, not even an explicit nil
### GetOperationalStatus

`func (o *ApplianceNodeInfoAllOf) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceNodeInfoAllOf) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceNodeInfoAllOf) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceNodeInfoAllOf) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


