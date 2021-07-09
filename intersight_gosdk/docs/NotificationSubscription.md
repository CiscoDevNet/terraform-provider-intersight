# NotificationSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "notification.AccountSubscription"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "notification.AccountSubscription"]
**Actions** | Pointer to [**[]NotificationAction**](NotificationAction.md) |  | [optional] 
**Conditions** | Pointer to [**[]NotificationAbstractCondition**](NotificationAbstractCondition.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Subscription can be switched on/off with out necessity to change the subscription settings: notification methods, conditions etc. Ex.: Subscription MO can be configured, but switched off. | [optional] 

## Methods

### NewNotificationSubscription

`func NewNotificationSubscription(classId string, objectType string, ) *NotificationSubscription`

NewNotificationSubscription instantiates a new NotificationSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSubscriptionWithDefaults

`func NewNotificationSubscriptionWithDefaults() *NotificationSubscription`

NewNotificationSubscriptionWithDefaults instantiates a new NotificationSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationSubscription) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationSubscription) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationSubscription) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationSubscription) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationSubscription) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationSubscription) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActions

`func (o *NotificationSubscription) GetActions() []NotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NotificationSubscription) GetActionsOk() (*[]NotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NotificationSubscription) SetActions(v []NotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *NotificationSubscription) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *NotificationSubscription) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *NotificationSubscription) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetConditions

`func (o *NotificationSubscription) GetConditions() []NotificationAbstractCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *NotificationSubscription) GetConditionsOk() (*[]NotificationAbstractCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *NotificationSubscription) SetConditions(v []NotificationAbstractCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *NotificationSubscription) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *NotificationSubscription) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *NotificationSubscription) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetEnabled

`func (o *NotificationSubscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationSubscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationSubscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationSubscription) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


