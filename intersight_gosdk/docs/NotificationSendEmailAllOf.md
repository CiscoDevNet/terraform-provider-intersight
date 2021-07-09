# NotificationSendEmailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.SendEmail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.SendEmail"]
**Email** | Pointer to **string** | Email of the recipient, who expects to be notified about the event that happens in Intersight. | [optional] 

## Methods

### NewNotificationSendEmailAllOf

`func NewNotificationSendEmailAllOf(classId string, objectType string, ) *NotificationSendEmailAllOf`

NewNotificationSendEmailAllOf instantiates a new NotificationSendEmailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSendEmailAllOfWithDefaults

`func NewNotificationSendEmailAllOfWithDefaults() *NotificationSendEmailAllOf`

NewNotificationSendEmailAllOfWithDefaults instantiates a new NotificationSendEmailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationSendEmailAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationSendEmailAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationSendEmailAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationSendEmailAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationSendEmailAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationSendEmailAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEmail

`func (o *NotificationSendEmailAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NotificationSendEmailAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NotificationSendEmailAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NotificationSendEmailAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


