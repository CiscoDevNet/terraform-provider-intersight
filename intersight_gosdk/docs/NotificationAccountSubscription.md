# NotificationAccountSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.AccountSubscription"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.AccountSubscription"]
**Description** | Pointer to **string** | The description for the subscription. | [optional] 
**Name** | Pointer to **string** | The name of the subscription. | [optional] 
**Type** | Pointer to **string** | The chosen subscription type imposes it is own validation rules. When &#39;email&#39; type is chosen, actions array can contain only one entry and it is entry should be of can be only notification.SendEmail; conditions can contain only notification.AlarmMoCondition and condition types should be unique. When the &#39;webhook&#39; type is chosen, the actions array can contain only one entry and it is entry should be of can be only notification.TriggerWebhook; conditions can contain up to a limited amount of entries and all of them should be of type notification.MoCondition. When the &#39;workflow&#39; type is chosen, the actions array can contain only one entry and it is entry should  be of type notification.TriggerWorkflow notification.TriggerWorkflow; conditions can contain up to a  limited amount of entries and all of them should be of type notification.WebhookEventMoCondition. * &#x60;email&#x60; - Email type requires usage of notification.SendEmail complex types for actionsand notification.AlarmMoCondition complex types for conditions. * &#x60;webhook&#x60; - Webhook type requires usage of notification.TriggerWebhook complex types for actionsand notification.MoCondition complex types for conditions. * &#x60;workflow&#x60; - Workflow type requires usage of notification.TriggerWorkflow complex types for actionsand notification.WebhookEventMoCondition complex types for conditions. | [optional] [default to "email"]
**Verify** | Pointer to **string** | Used to verify the actions of the Subscription MO. For a &#39;webhook&#39; type Ping event is sent to verify that the webhook server is accessible. For an &#39;email&#39; type there will be a verification email sent. * &#x60;none&#x60; - No actions will be verified. Default value. * &#x60;all&#x60; - All actions will be re-verified. The previous state of the verification will be preserved. | [optional] [default to "none"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewNotificationAccountSubscription

`func NewNotificationAccountSubscription(classId string, objectType string, ) *NotificationAccountSubscription`

NewNotificationAccountSubscription instantiates a new NotificationAccountSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAccountSubscriptionWithDefaults

`func NewNotificationAccountSubscriptionWithDefaults() *NotificationAccountSubscription`

NewNotificationAccountSubscriptionWithDefaults instantiates a new NotificationAccountSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationAccountSubscription) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationAccountSubscription) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationAccountSubscription) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationAccountSubscription) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationAccountSubscription) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationAccountSubscription) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NotificationAccountSubscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationAccountSubscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationAccountSubscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationAccountSubscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *NotificationAccountSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationAccountSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationAccountSubscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationAccountSubscription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NotificationAccountSubscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationAccountSubscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationAccountSubscription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationAccountSubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerify

`func (o *NotificationAccountSubscription) GetVerify() string`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *NotificationAccountSubscription) GetVerifyOk() (*string, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *NotificationAccountSubscription) SetVerify(v string)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *NotificationAccountSubscription) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetAccount

`func (o *NotificationAccountSubscription) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NotificationAccountSubscription) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NotificationAccountSubscription) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NotificationAccountSubscription) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *NotificationAccountSubscription) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *NotificationAccountSubscription) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


