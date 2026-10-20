# TopologyAlarmCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.AlarmCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.AlarmCount"]
**Critical** | Pointer to **int32** | Count of Critical alarms  in the selected topology interval. | [optional] [readonly] 
**Health** | Pointer to **string** | Derived health state for a node, or Unknown for a topology-wide timeline. | [optional] [readonly] 
**Total** | Pointer to **int32** | Total countable alarms, calculated as Critical plus Warning. | [optional] [readonly] 
**Warning** | Pointer to **int32** | Count of Warning alarms in the selected topology interval. | [optional] [readonly] 

## Methods

### NewTopologyAlarmCount

`func NewTopologyAlarmCount(classId string, objectType string, ) *TopologyAlarmCount`

NewTopologyAlarmCount instantiates a new TopologyAlarmCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyAlarmCountWithDefaults

`func NewTopologyAlarmCountWithDefaults() *TopologyAlarmCount`

NewTopologyAlarmCountWithDefaults instantiates a new TopologyAlarmCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyAlarmCount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyAlarmCount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyAlarmCount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyAlarmCount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyAlarmCount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyAlarmCount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCritical

`func (o *TopologyAlarmCount) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *TopologyAlarmCount) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *TopologyAlarmCount) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *TopologyAlarmCount) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHealth

`func (o *TopologyAlarmCount) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *TopologyAlarmCount) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *TopologyAlarmCount) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *TopologyAlarmCount) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetTotal

`func (o *TopologyAlarmCount) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TopologyAlarmCount) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TopologyAlarmCount) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TopologyAlarmCount) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWarning

`func (o *TopologyAlarmCount) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *TopologyAlarmCount) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *TopologyAlarmCount) SetWarning(v int32)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *TopologyAlarmCount) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


