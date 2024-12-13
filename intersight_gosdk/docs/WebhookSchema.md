# WebhookSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "webhook.Schema"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "webhook.Schema"]
**Data** | Pointer to **string** | The actual JSON schema data. The data must be a valid JSON schema string. | [optional] 
**Description** | Pointer to **string** | The description of the schema. | [optional] 
**EventType** | Pointer to **string** | The event type of the schema. | [optional] 
**Source** | Pointer to **string** | The source of the schema. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewWebhookSchema

`func NewWebhookSchema(classId string, objectType string, ) *WebhookSchema`

NewWebhookSchema instantiates a new WebhookSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSchemaWithDefaults

`func NewWebhookSchemaWithDefaults() *WebhookSchema`

NewWebhookSchemaWithDefaults instantiates a new WebhookSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WebhookSchema) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WebhookSchema) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WebhookSchema) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WebhookSchema) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WebhookSchema) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WebhookSchema) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetData

`func (o *WebhookSchema) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookSchema) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookSchema) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *WebhookSchema) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventType

`func (o *WebhookSchema) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookSchema) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookSchema) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *WebhookSchema) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *WebhookSchema) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebhookSchema) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebhookSchema) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *WebhookSchema) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAccount

`func (o *WebhookSchema) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookSchema) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookSchema) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *WebhookSchema) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *WebhookSchema) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *WebhookSchema) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


