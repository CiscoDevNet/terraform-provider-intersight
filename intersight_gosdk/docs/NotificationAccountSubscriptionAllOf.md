# NotificationAccountSubscriptionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.AccountSubscription"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.AccountSubscription"]
**Name** | Pointer to **string** | The name of the subscription. | [optional] 
**Type** | Pointer to **string** | The chosen subscription type imposes it is own validation rules. When &#39;email&#39; type is chosen, actions array can contain only one entry and it is entry should be of can be only notification.SendEmail; conditions can contain only notification.AlarmMoCondition and condition types should be unique. When the &#39;webhook&#39; type is chosen, the actions array can contain only one entry and it is entry should be of can be only notification.TriggerWebhook; conditions can contain up to a limited amount of entries and all of them should be of type notification.MoCondition. * &#x60;email&#x60; - Email type requires usage of notification.SendEmail complex types for actionsand notification.AlarmMoCondition complex types for conditions. * &#x60;webhook&#x60; - Webhook type requires usage of notification.TriggerWebhook complex types for actionsand notification.MoCondition complex types for conditions. | [optional] [default to "email"]
**Verify** | Pointer to **string** | Used to verify the actions of the Subscription MO. For a &#39;webhook&#39; type Ping event is sent to verify that the webhook server is accessible. For an &#39;email&#39; type there will be a verification email sent. * &#x60;none&#x60; - No actions will be verified. Default value. * &#x60;all&#x60; - All actions will be re-verified. The previous state of the verification will be preserved. | [optional] [default to "none"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewNotificationAccountSubscriptionAllOf

`func NewNotificationAccountSubscriptionAllOf(classId string, objectType string, ) *NotificationAccountSubscriptionAllOf`

NewNotificationAccountSubscriptionAllOf instantiates a new NotificationAccountSubscriptionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAccountSubscriptionAllOfWithDefaults

`func NewNotificationAccountSubscriptionAllOfWithDefaults() *NotificationAccountSubscriptionAllOf`

NewNotificationAccountSubscriptionAllOfWithDefaults instantiates a new NotificationAccountSubscriptionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationAccountSubscriptionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationAccountSubscriptionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationAccountSubscriptionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationAccountSubscriptionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationAccountSubscriptionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationAccountSubscriptionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *NotificationAccountSubscriptionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationAccountSubscriptionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationAccountSubscriptionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationAccountSubscriptionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NotificationAccountSubscriptionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationAccountSubscriptionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationAccountSubscriptionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationAccountSubscriptionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerify

`func (o *NotificationAccountSubscriptionAllOf) GetVerify() string`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *NotificationAccountSubscriptionAllOf) GetVerifyOk() (*string, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *NotificationAccountSubscriptionAllOf) SetVerify(v string)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *NotificationAccountSubscriptionAllOf) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetAccount

`func (o *NotificationAccountSubscriptionAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NotificationAccountSubscriptionAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NotificationAccountSubscriptionAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NotificationAccountSubscriptionAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


