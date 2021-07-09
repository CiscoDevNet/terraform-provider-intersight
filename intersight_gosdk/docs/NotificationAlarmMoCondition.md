# NotificationAlarmMoCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.AlarmMoCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.AlarmMoCondition"]
**Severity** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotificationAlarmMoCondition

`func NewNotificationAlarmMoCondition(classId string, objectType string, ) *NotificationAlarmMoCondition`

NewNotificationAlarmMoCondition instantiates a new NotificationAlarmMoCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAlarmMoConditionWithDefaults

`func NewNotificationAlarmMoConditionWithDefaults() *NotificationAlarmMoCondition`

NewNotificationAlarmMoConditionWithDefaults instantiates a new NotificationAlarmMoCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationAlarmMoCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationAlarmMoCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationAlarmMoCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationAlarmMoCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationAlarmMoCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationAlarmMoCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSeverity

`func (o *NotificationAlarmMoCondition) GetSeverity() []string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationAlarmMoCondition) GetSeverityOk() (*[]string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationAlarmMoCondition) SetSeverity(v []string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationAlarmMoCondition) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *NotificationAlarmMoCondition) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *NotificationAlarmMoCondition) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


