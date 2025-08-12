# NotificationSimpleFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.SimpleFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.SimpleFilter"]
**Operator** | Pointer to **string** | The operator to apply (eq, ne, contains, in). | [optional] 
**Property** | Pointer to **string** | The property name to filter on (e.g., HostName, Domain,  ServerProfile, Organization). | [optional] 
**Value** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotificationSimpleFilter

`func NewNotificationSimpleFilter(classId string, objectType string, ) *NotificationSimpleFilter`

NewNotificationSimpleFilter instantiates a new NotificationSimpleFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSimpleFilterWithDefaults

`func NewNotificationSimpleFilterWithDefaults() *NotificationSimpleFilter`

NewNotificationSimpleFilterWithDefaults instantiates a new NotificationSimpleFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationSimpleFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationSimpleFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationSimpleFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationSimpleFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationSimpleFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationSimpleFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperator

`func (o *NotificationSimpleFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NotificationSimpleFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NotificationSimpleFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *NotificationSimpleFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetProperty

`func (o *NotificationSimpleFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *NotificationSimpleFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *NotificationSimpleFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *NotificationSimpleFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *NotificationSimpleFilter) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NotificationSimpleFilter) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NotificationSimpleFilter) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NotificationSimpleFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *NotificationSimpleFilter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *NotificationSimpleFilter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


