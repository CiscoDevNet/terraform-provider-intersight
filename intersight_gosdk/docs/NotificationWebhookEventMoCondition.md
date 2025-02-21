# NotificationWebhookEventMoCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.WebhookEventMoCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.WebhookEventMoCondition"]
**EndpointMoid** | Pointer to **string** | The required webhook.Endpoint Moid of a webhook.Event for the condition to be true. | [optional] 
**SchemaMoid** | Pointer to **string** | The required webhook.Schema Moid of a webhook.Event for the condition to be true. | [optional] 

## Methods

### NewNotificationWebhookEventMoCondition

`func NewNotificationWebhookEventMoCondition(classId string, objectType string, ) *NotificationWebhookEventMoCondition`

NewNotificationWebhookEventMoCondition instantiates a new NotificationWebhookEventMoCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWebhookEventMoConditionWithDefaults

`func NewNotificationWebhookEventMoConditionWithDefaults() *NotificationWebhookEventMoCondition`

NewNotificationWebhookEventMoConditionWithDefaults instantiates a new NotificationWebhookEventMoCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationWebhookEventMoCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationWebhookEventMoCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationWebhookEventMoCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationWebhookEventMoCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationWebhookEventMoCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationWebhookEventMoCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndpointMoid

`func (o *NotificationWebhookEventMoCondition) GetEndpointMoid() string`

GetEndpointMoid returns the EndpointMoid field if non-nil, zero value otherwise.

### GetEndpointMoidOk

`func (o *NotificationWebhookEventMoCondition) GetEndpointMoidOk() (*string, bool)`

GetEndpointMoidOk returns a tuple with the EndpointMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointMoid

`func (o *NotificationWebhookEventMoCondition) SetEndpointMoid(v string)`

SetEndpointMoid sets EndpointMoid field to given value.

### HasEndpointMoid

`func (o *NotificationWebhookEventMoCondition) HasEndpointMoid() bool`

HasEndpointMoid returns a boolean if a field has been set.

### GetSchemaMoid

`func (o *NotificationWebhookEventMoCondition) GetSchemaMoid() string`

GetSchemaMoid returns the SchemaMoid field if non-nil, zero value otherwise.

### GetSchemaMoidOk

`func (o *NotificationWebhookEventMoCondition) GetSchemaMoidOk() (*string, bool)`

GetSchemaMoidOk returns a tuple with the SchemaMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaMoid

`func (o *NotificationWebhookEventMoCondition) SetSchemaMoid(v string)`

SetSchemaMoid sets SchemaMoid field to given value.

### HasSchemaMoid

`func (o *NotificationWebhookEventMoCondition) HasSchemaMoid() bool`

HasSchemaMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


