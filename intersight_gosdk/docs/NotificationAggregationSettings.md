# NotificationAggregationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.AggregationSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.AggregationSettings"]
**Enabled** | Pointer to **bool** | Show if aggregation is enabled or not. If aggregation is enabled, events will be grouped and delivered as one notification, if it is disabled - each event will be delivered as independent notification. | [optional] 
**EventsLimit** | Pointer to **int64** | Limit of events in one group. If this limit is reached, a window will be closed regardless of the time limit. | [optional] 
**Size** | Pointer to **string** | Time in seconds which means the max time after which the window will be closed. | [optional] 
**Step** | Pointer to **string** | Time in seconds which means how long after the last event Intersight should wait for the next event to come. If there&#39;s no new event that matches with the same Subscription within this time, then the window will be closed. | [optional] 

## Methods

### NewNotificationAggregationSettings

`func NewNotificationAggregationSettings(classId string, objectType string, ) *NotificationAggregationSettings`

NewNotificationAggregationSettings instantiates a new NotificationAggregationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAggregationSettingsWithDefaults

`func NewNotificationAggregationSettingsWithDefaults() *NotificationAggregationSettings`

NewNotificationAggregationSettingsWithDefaults instantiates a new NotificationAggregationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationAggregationSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationAggregationSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationAggregationSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationAggregationSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationAggregationSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationAggregationSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *NotificationAggregationSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationAggregationSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationAggregationSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationAggregationSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventsLimit

`func (o *NotificationAggregationSettings) GetEventsLimit() int64`

GetEventsLimit returns the EventsLimit field if non-nil, zero value otherwise.

### GetEventsLimitOk

`func (o *NotificationAggregationSettings) GetEventsLimitOk() (*int64, bool)`

GetEventsLimitOk returns a tuple with the EventsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsLimit

`func (o *NotificationAggregationSettings) SetEventsLimit(v int64)`

SetEventsLimit sets EventsLimit field to given value.

### HasEventsLimit

`func (o *NotificationAggregationSettings) HasEventsLimit() bool`

HasEventsLimit returns a boolean if a field has been set.

### GetSize

`func (o *NotificationAggregationSettings) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NotificationAggregationSettings) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NotificationAggregationSettings) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *NotificationAggregationSettings) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStep

`func (o *NotificationAggregationSettings) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *NotificationAggregationSettings) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *NotificationAggregationSettings) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *NotificationAggregationSettings) HasStep() bool`

HasStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


