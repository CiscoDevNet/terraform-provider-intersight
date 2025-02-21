# ApplianceClusterWorkerNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ClusterWorkerNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ClusterWorkerNode"]
**Hostname** | Pointer to **string** | Hostname of the worker node being added. | [optional] 
**ReuseNode** | Pointer to **bool** | Indicates if the worker node being added is being reused. | [optional] 
**SessionId** | Pointer to **string** | Session Moid for the user session. | [optional] [readonly] 

## Methods

### NewApplianceClusterWorkerNode

`func NewApplianceClusterWorkerNode(classId string, objectType string, ) *ApplianceClusterWorkerNode`

NewApplianceClusterWorkerNode instantiates a new ApplianceClusterWorkerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceClusterWorkerNodeWithDefaults

`func NewApplianceClusterWorkerNodeWithDefaults() *ApplianceClusterWorkerNode`

NewApplianceClusterWorkerNodeWithDefaults instantiates a new ApplianceClusterWorkerNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceClusterWorkerNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceClusterWorkerNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceClusterWorkerNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceClusterWorkerNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceClusterWorkerNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceClusterWorkerNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *ApplianceClusterWorkerNode) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceClusterWorkerNode) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceClusterWorkerNode) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceClusterWorkerNode) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetReuseNode

`func (o *ApplianceClusterWorkerNode) GetReuseNode() bool`

GetReuseNode returns the ReuseNode field if non-nil, zero value otherwise.

### GetReuseNodeOk

`func (o *ApplianceClusterWorkerNode) GetReuseNodeOk() (*bool, bool)`

GetReuseNodeOk returns a tuple with the ReuseNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseNode

`func (o *ApplianceClusterWorkerNode) SetReuseNode(v bool)`

SetReuseNode sets ReuseNode field to given value.

### HasReuseNode

`func (o *ApplianceClusterWorkerNode) HasReuseNode() bool`

HasReuseNode returns a boolean if a field has been set.

### GetSessionId

`func (o *ApplianceClusterWorkerNode) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplianceClusterWorkerNode) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApplianceClusterWorkerNode) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ApplianceClusterWorkerNode) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


