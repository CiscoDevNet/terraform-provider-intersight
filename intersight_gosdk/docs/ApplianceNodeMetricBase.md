# ApplianceNodeMetricBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**NodeHostname** | Pointer to **string** | Hostname of the Intersight Appliance node. This human-readable identifier  corresponds to the Kubernetes node name used for metric queries and provides  an intuitive way for users to correlate resource utilization data  with specific physical or virtual machines in the cluster topology. | [optional] [readonly] 
**NodeId** | Pointer to **int64** | System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. | [optional] [readonly] 

## Methods

### NewApplianceNodeMetricBase

`func NewApplianceNodeMetricBase(classId string, objectType string, ) *ApplianceNodeMetricBase`

NewApplianceNodeMetricBase instantiates a new ApplianceNodeMetricBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeMetricBaseWithDefaults

`func NewApplianceNodeMetricBaseWithDefaults() *ApplianceNodeMetricBase`

NewApplianceNodeMetricBaseWithDefaults instantiates a new ApplianceNodeMetricBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeMetricBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeMetricBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeMetricBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeMetricBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeMetricBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeMetricBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodeHostname

`func (o *ApplianceNodeMetricBase) GetNodeHostname() string`

GetNodeHostname returns the NodeHostname field if non-nil, zero value otherwise.

### GetNodeHostnameOk

`func (o *ApplianceNodeMetricBase) GetNodeHostnameOk() (*string, bool)`

GetNodeHostnameOk returns a tuple with the NodeHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHostname

`func (o *ApplianceNodeMetricBase) SetNodeHostname(v string)`

SetNodeHostname sets NodeHostname field to given value.

### HasNodeHostname

`func (o *ApplianceNodeMetricBase) HasNodeHostname() bool`

HasNodeHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceNodeMetricBase) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceNodeMetricBase) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceNodeMetricBase) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceNodeMetricBase) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


