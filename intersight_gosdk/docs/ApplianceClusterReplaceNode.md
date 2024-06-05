# ApplianceClusterReplaceNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ClusterReplaceNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ClusterReplaceNode"]
**Hostname** | Pointer to **string** | Hostname of the node being replaced. | [optional] 
**NodeId** | Pointer to **int64** | Node id of the node being replaced. | [optional] 
**NodeIpChanged** | Pointer to **bool** | If the node being replaced has a different IP. | [optional] [readonly] 

## Methods

### NewApplianceClusterReplaceNode

`func NewApplianceClusterReplaceNode(classId string, objectType string, ) *ApplianceClusterReplaceNode`

NewApplianceClusterReplaceNode instantiates a new ApplianceClusterReplaceNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceClusterReplaceNodeWithDefaults

`func NewApplianceClusterReplaceNodeWithDefaults() *ApplianceClusterReplaceNode`

NewApplianceClusterReplaceNodeWithDefaults instantiates a new ApplianceClusterReplaceNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceClusterReplaceNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceClusterReplaceNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceClusterReplaceNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceClusterReplaceNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceClusterReplaceNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceClusterReplaceNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *ApplianceClusterReplaceNode) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceClusterReplaceNode) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceClusterReplaceNode) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceClusterReplaceNode) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceClusterReplaceNode) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceClusterReplaceNode) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceClusterReplaceNode) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceClusterReplaceNode) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeIpChanged

`func (o *ApplianceClusterReplaceNode) GetNodeIpChanged() bool`

GetNodeIpChanged returns the NodeIpChanged field if non-nil, zero value otherwise.

### GetNodeIpChangedOk

`func (o *ApplianceClusterReplaceNode) GetNodeIpChangedOk() (*bool, bool)`

GetNodeIpChangedOk returns a tuple with the NodeIpChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpChanged

`func (o *ApplianceClusterReplaceNode) SetNodeIpChanged(v bool)`

SetNodeIpChanged sets NodeIpChanged field to given value.

### HasNodeIpChanged

`func (o *ApplianceClusterReplaceNode) HasNodeIpChanged() bool`

HasNodeIpChanged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


