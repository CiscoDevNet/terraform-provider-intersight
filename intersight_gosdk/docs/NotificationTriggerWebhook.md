# NotificationTriggerWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.TriggerWebhook"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.TriggerWebhook"]
**FirstFailedRequest** | Pointer to **time.Time** | Holds the timestamp of the first failed request. The first time the notification is not delivered to the webhook server, the user will have the Warning alarm in the system. Next 48 hours the system still will try to notify the webhook server. If after 48 hours the server is not recovered, the system will mark this webhook as Inactive, and the user will have a critical alarm in the system. | [optional] [readonly] 
**IsSecretSet** | Pointer to **bool** | Indicates whether the value of the &#39;secret&#39; property has been set. | [optional] [readonly] [default to false]
**LastNetworkError** | Pointer to **string** | Holds the error message for the user of the last response. | [optional] [readonly] 
**LastResponseCode** | Pointer to **int64** | Holds the code of the last response, which helps to debug the issue in case if webhook server is not reachable. | [optional] [readonly] 
**Secret** | Pointer to **string** | The secret is used to build the Authorization header, which will be attached to each webhook notification. By this header developers of the webhooks servers can make sure that events are received from the trusted source - Intersight. | [optional] 
**State** | Pointer to **string** | State of the action shows whether this action passes the verification or not. If this property holds &#39;Inactive&#39; value, this action will not be executed. To verify action again, use the Verify property from the MO. * &#x60;Inactive&#x60; - Inactive state means action didn&#39;t pass the verification and it won&#39;t be executed. * &#x60;Active&#x60; - Active state means that action successfully passed the verification and it is ready to be performed. | [optional] [readonly] [default to "Inactive"]
**Url** | Pointer to **string** | Payload URL of the recipient app, which is intended to serve the events that happens in Intersight. | [optional] 

## Methods

### NewNotificationTriggerWebhook

`func NewNotificationTriggerWebhook(classId string, objectType string, ) *NotificationTriggerWebhook`

NewNotificationTriggerWebhook instantiates a new NotificationTriggerWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTriggerWebhookWithDefaults

`func NewNotificationTriggerWebhookWithDefaults() *NotificationTriggerWebhook`

NewNotificationTriggerWebhookWithDefaults instantiates a new NotificationTriggerWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationTriggerWebhook) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationTriggerWebhook) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationTriggerWebhook) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationTriggerWebhook) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationTriggerWebhook) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationTriggerWebhook) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFirstFailedRequest

`func (o *NotificationTriggerWebhook) GetFirstFailedRequest() time.Time`

GetFirstFailedRequest returns the FirstFailedRequest field if non-nil, zero value otherwise.

### GetFirstFailedRequestOk

`func (o *NotificationTriggerWebhook) GetFirstFailedRequestOk() (*time.Time, bool)`

GetFirstFailedRequestOk returns a tuple with the FirstFailedRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFailedRequest

`func (o *NotificationTriggerWebhook) SetFirstFailedRequest(v time.Time)`

SetFirstFailedRequest sets FirstFailedRequest field to given value.

### HasFirstFailedRequest

`func (o *NotificationTriggerWebhook) HasFirstFailedRequest() bool`

HasFirstFailedRequest returns a boolean if a field has been set.

### GetIsSecretSet

`func (o *NotificationTriggerWebhook) GetIsSecretSet() bool`

GetIsSecretSet returns the IsSecretSet field if non-nil, zero value otherwise.

### GetIsSecretSetOk

`func (o *NotificationTriggerWebhook) GetIsSecretSetOk() (*bool, bool)`

GetIsSecretSetOk returns a tuple with the IsSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecretSet

`func (o *NotificationTriggerWebhook) SetIsSecretSet(v bool)`

SetIsSecretSet sets IsSecretSet field to given value.

### HasIsSecretSet

`func (o *NotificationTriggerWebhook) HasIsSecretSet() bool`

HasIsSecretSet returns a boolean if a field has been set.

### GetLastNetworkError

`func (o *NotificationTriggerWebhook) GetLastNetworkError() string`

GetLastNetworkError returns the LastNetworkError field if non-nil, zero value otherwise.

### GetLastNetworkErrorOk

`func (o *NotificationTriggerWebhook) GetLastNetworkErrorOk() (*string, bool)`

GetLastNetworkErrorOk returns a tuple with the LastNetworkError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNetworkError

`func (o *NotificationTriggerWebhook) SetLastNetworkError(v string)`

SetLastNetworkError sets LastNetworkError field to given value.

### HasLastNetworkError

`func (o *NotificationTriggerWebhook) HasLastNetworkError() bool`

HasLastNetworkError returns a boolean if a field has been set.

### GetLastResponseCode

`func (o *NotificationTriggerWebhook) GetLastResponseCode() int64`

GetLastResponseCode returns the LastResponseCode field if non-nil, zero value otherwise.

### GetLastResponseCodeOk

`func (o *NotificationTriggerWebhook) GetLastResponseCodeOk() (*int64, bool)`

GetLastResponseCodeOk returns a tuple with the LastResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseCode

`func (o *NotificationTriggerWebhook) SetLastResponseCode(v int64)`

SetLastResponseCode sets LastResponseCode field to given value.

### HasLastResponseCode

`func (o *NotificationTriggerWebhook) HasLastResponseCode() bool`

HasLastResponseCode returns a boolean if a field has been set.

### GetSecret

`func (o *NotificationTriggerWebhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *NotificationTriggerWebhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *NotificationTriggerWebhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *NotificationTriggerWebhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetState

`func (o *NotificationTriggerWebhook) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NotificationTriggerWebhook) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NotificationTriggerWebhook) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NotificationTriggerWebhook) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationTriggerWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationTriggerWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationTriggerWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationTriggerWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


