# NotificationAbstractMoConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Enabled** | Pointer to **bool** | The condition can be switched on/off with out necessity to change the subscription settings: actions, conditions, etc. Ex.: Subscription MO can be configured, but switched off. | [optional] 
**MoType** | Pointer to **string** | MoType for which the condition is created. | [optional] 
**Operations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotificationAbstractMoConditionAllOf

`func NewNotificationAbstractMoConditionAllOf(classId string, objectType string, ) *NotificationAbstractMoConditionAllOf`

NewNotificationAbstractMoConditionAllOf instantiates a new NotificationAbstractMoConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAbstractMoConditionAllOfWithDefaults

`func NewNotificationAbstractMoConditionAllOfWithDefaults() *NotificationAbstractMoConditionAllOf`

NewNotificationAbstractMoConditionAllOfWithDefaults instantiates a new NotificationAbstractMoConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationAbstractMoConditionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationAbstractMoConditionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationAbstractMoConditionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationAbstractMoConditionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationAbstractMoConditionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationAbstractMoConditionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *NotificationAbstractMoConditionAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationAbstractMoConditionAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationAbstractMoConditionAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationAbstractMoConditionAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMoType

`func (o *NotificationAbstractMoConditionAllOf) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *NotificationAbstractMoConditionAllOf) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *NotificationAbstractMoConditionAllOf) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *NotificationAbstractMoConditionAllOf) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetOperations

`func (o *NotificationAbstractMoConditionAllOf) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *NotificationAbstractMoConditionAllOf) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *NotificationAbstractMoConditionAllOf) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *NotificationAbstractMoConditionAllOf) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperationsNil

`func (o *NotificationAbstractMoConditionAllOf) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *NotificationAbstractMoConditionAllOf) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


