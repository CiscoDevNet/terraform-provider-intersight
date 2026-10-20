# TopologyEventDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "topology.EventDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "topology.EventDetails"]
**Category** | Pointer to **string** | Category of details to return. * &#x60;Unknown&#x60; - Unknown or unspecified category. * &#x60;Alarms&#x60; - Alarm detail records for the selected topology interval. * &#x60;Requests&#x60; - Workflow request detail records for the selected topology interval. | [optional] [default to "Unknown"]
**Events** | Pointer to [**NullableTopologyEvent**](TopologyEvent.md) |  | [optional] 
**IncludeAll** | Pointer to **bool** | When true, include all scoped records contributing to the selected value. | [optional] 

## Methods

### NewTopologyEventDetails

`func NewTopologyEventDetails(classId string, objectType string, ) *TopologyEventDetails`

NewTopologyEventDetails instantiates a new TopologyEventDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyEventDetailsWithDefaults

`func NewTopologyEventDetailsWithDefaults() *TopologyEventDetails`

NewTopologyEventDetailsWithDefaults instantiates a new TopologyEventDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopologyEventDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopologyEventDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopologyEventDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopologyEventDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopologyEventDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopologyEventDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *TopologyEventDetails) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TopologyEventDetails) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TopologyEventDetails) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TopologyEventDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEvents

`func (o *TopologyEventDetails) GetEvents() TopologyEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TopologyEventDetails) GetEventsOk() (*TopologyEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TopologyEventDetails) SetEvents(v TopologyEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TopologyEventDetails) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *TopologyEventDetails) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *TopologyEventDetails) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetIncludeAll

`func (o *TopologyEventDetails) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *TopologyEventDetails) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *TopologyEventDetails) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.

### HasIncludeAll

`func (o *TopologyEventDetails) HasIncludeAll() bool`

HasIncludeAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


