# NotificationTriggerWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.TriggerWorkflow"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.TriggerWorkflow"]
**WorkflowDefinitionMoid** | Pointer to **string** | Target workflow definition moid to trigger. | [optional] 

## Methods

### NewNotificationTriggerWorkflow

`func NewNotificationTriggerWorkflow(classId string, objectType string, ) *NotificationTriggerWorkflow`

NewNotificationTriggerWorkflow instantiates a new NotificationTriggerWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTriggerWorkflowWithDefaults

`func NewNotificationTriggerWorkflowWithDefaults() *NotificationTriggerWorkflow`

NewNotificationTriggerWorkflowWithDefaults instantiates a new NotificationTriggerWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationTriggerWorkflow) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationTriggerWorkflow) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationTriggerWorkflow) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationTriggerWorkflow) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationTriggerWorkflow) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationTriggerWorkflow) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetWorkflowDefinitionMoid

`func (o *NotificationTriggerWorkflow) GetWorkflowDefinitionMoid() string`

GetWorkflowDefinitionMoid returns the WorkflowDefinitionMoid field if non-nil, zero value otherwise.

### GetWorkflowDefinitionMoidOk

`func (o *NotificationTriggerWorkflow) GetWorkflowDefinitionMoidOk() (*string, bool)`

GetWorkflowDefinitionMoidOk returns a tuple with the WorkflowDefinitionMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinitionMoid

`func (o *NotificationTriggerWorkflow) SetWorkflowDefinitionMoid(v string)`

SetWorkflowDefinitionMoid sets WorkflowDefinitionMoid field to given value.

### HasWorkflowDefinitionMoid

`func (o *NotificationTriggerWorkflow) HasWorkflowDefinitionMoid() bool`

HasWorkflowDefinitionMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


