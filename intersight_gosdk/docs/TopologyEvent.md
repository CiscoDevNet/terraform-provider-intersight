# TopologyEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.Event"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.Event"]
**Alarms** | Pointer to **interface{}** | Alarm detail entries returned when the requested category is Alarms. | [optional] [readonly] 
**Requests** | Pointer to **interface{}** | Workflow request detail entries returned when the requested category is Requests. | [optional] [readonly] 

## Methods

### NewTopologyEvent

`func NewTopologyEvent(classId string, objectType string, ) *TopologyEvent`

NewTopologyEvent instantiates a new TopologyEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyEventWithDefaults

`func NewTopologyEventWithDefaults() *TopologyEvent`

NewTopologyEventWithDefaults instantiates a new TopologyEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyEvent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyEvent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyEvent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarms

`func (o *TopologyEvent) GetAlarms() interface{}`

GetAlarms returns the Alarms field if non-nil, zero value otherwise.

### GetAlarmsOk

`func (o *TopologyEvent) GetAlarmsOk() (*interface{}, bool)`

GetAlarmsOk returns a tuple with the Alarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarms

`func (o *TopologyEvent) SetAlarms(v interface{})`

SetAlarms sets Alarms field to given value.

### HasAlarms

`func (o *TopologyEvent) HasAlarms() bool`

HasAlarms returns a boolean if a field has been set.

### SetAlarmsNil

`func (o *TopologyEvent) SetAlarmsNil(b bool)`

 SetAlarmsNil sets the value for Alarms to be an explicit nil

### UnsetAlarms
`func (o *TopologyEvent) UnsetAlarms()`

UnsetAlarms ensures that no value is present for Alarms, not even an explicit nil
### GetRequests

`func (o *TopologyEvent) GetRequests() interface{}`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TopologyEvent) GetRequestsOk() (*interface{}, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TopologyEvent) SetRequests(v interface{})`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *TopologyEvent) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *TopologyEvent) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *TopologyEvent) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


