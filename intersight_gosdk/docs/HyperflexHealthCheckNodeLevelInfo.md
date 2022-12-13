# HyperflexHealthCheckNodeLevelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HealthCheckNodeLevelInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HealthCheckNodeLevelInfo"]
**NodeCause** | Pointer to **string** | Node-specific check failure cause. | [optional] [readonly] 
**NodeCheckResult** | Pointer to **string** | Node-specific check result. | [optional] [readonly] 
**NodeEsxIpAddress** | Pointer to **string** | The IP Address of the ESXi server. | [optional] [readonly] 
**NodeIpAddress** | Pointer to **string** | The IP Address of cluster node. | [optional] [readonly] 
**NodeName** | Pointer to **string** | Cluster node name on which the check was run. | [optional] [readonly] 
**NodeResolution** | Pointer to **string** | Node-specific check failure suggested resolution. | [optional] [readonly] 

## Methods

### NewHyperflexHealthCheckNodeLevelInfo

`func NewHyperflexHealthCheckNodeLevelInfo(classId string, objectType string, ) *HyperflexHealthCheckNodeLevelInfo`

NewHyperflexHealthCheckNodeLevelInfo instantiates a new HyperflexHealthCheckNodeLevelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckNodeLevelInfoWithDefaults

`func NewHyperflexHealthCheckNodeLevelInfoWithDefaults() *HyperflexHealthCheckNodeLevelInfo`

NewHyperflexHealthCheckNodeLevelInfoWithDefaults instantiates a new HyperflexHealthCheckNodeLevelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckNodeLevelInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckNodeLevelInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckNodeLevelInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckNodeLevelInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodeCause

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeCause() string`

GetNodeCause returns the NodeCause field if non-nil, zero value otherwise.

### GetNodeCauseOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeCauseOk() (*string, bool)`

GetNodeCauseOk returns a tuple with the NodeCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCause

`func (o *HyperflexHealthCheckNodeLevelInfo) SetNodeCause(v string)`

SetNodeCause sets NodeCause field to given value.

### HasNodeCause

`func (o *HyperflexHealthCheckNodeLevelInfo) HasNodeCause() bool`

HasNodeCause returns a boolean if a field has been set.

### GetNodeCheckResult

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeCheckResult() string`

GetNodeCheckResult returns the NodeCheckResult field if non-nil, zero value otherwise.

### GetNodeCheckResultOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeCheckResultOk() (*string, bool)`

GetNodeCheckResultOk returns a tuple with the NodeCheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCheckResult

`func (o *HyperflexHealthCheckNodeLevelInfo) SetNodeCheckResult(v string)`

SetNodeCheckResult sets NodeCheckResult field to given value.

### HasNodeCheckResult

`func (o *HyperflexHealthCheckNodeLevelInfo) HasNodeCheckResult() bool`

HasNodeCheckResult returns a boolean if a field has been set.

### GetNodeEsxIpAddress

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeEsxIpAddress() string`

GetNodeEsxIpAddress returns the NodeEsxIpAddress field if non-nil, zero value otherwise.

### GetNodeEsxIpAddressOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeEsxIpAddressOk() (*string, bool)`

GetNodeEsxIpAddressOk returns a tuple with the NodeEsxIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeEsxIpAddress

`func (o *HyperflexHealthCheckNodeLevelInfo) SetNodeEsxIpAddress(v string)`

SetNodeEsxIpAddress sets NodeEsxIpAddress field to given value.

### HasNodeEsxIpAddress

`func (o *HyperflexHealthCheckNodeLevelInfo) HasNodeEsxIpAddress() bool`

HasNodeEsxIpAddress returns a boolean if a field has been set.

### GetNodeIpAddress

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeIpAddress() string`

GetNodeIpAddress returns the NodeIpAddress field if non-nil, zero value otherwise.

### GetNodeIpAddressOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeIpAddressOk() (*string, bool)`

GetNodeIpAddressOk returns a tuple with the NodeIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpAddress

`func (o *HyperflexHealthCheckNodeLevelInfo) SetNodeIpAddress(v string)`

SetNodeIpAddress sets NodeIpAddress field to given value.

### HasNodeIpAddress

`func (o *HyperflexHealthCheckNodeLevelInfo) HasNodeIpAddress() bool`

HasNodeIpAddress returns a boolean if a field has been set.

### GetNodeName

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *HyperflexHealthCheckNodeLevelInfo) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *HyperflexHealthCheckNodeLevelInfo) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeResolution

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeResolution() string`

GetNodeResolution returns the NodeResolution field if non-nil, zero value otherwise.

### GetNodeResolutionOk

`func (o *HyperflexHealthCheckNodeLevelInfo) GetNodeResolutionOk() (*string, bool)`

GetNodeResolutionOk returns a tuple with the NodeResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeResolution

`func (o *HyperflexHealthCheckNodeLevelInfo) SetNodeResolution(v string)`

SetNodeResolution sets NodeResolution field to given value.

### HasNodeResolution

`func (o *HyperflexHealthCheckNodeLevelInfo) HasNodeResolution() bool`

HasNodeResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


