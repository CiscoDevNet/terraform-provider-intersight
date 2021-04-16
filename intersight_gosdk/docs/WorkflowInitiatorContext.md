# WorkflowInitiatorContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.InitiatorContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.InitiatorContext"]
**InitiatorMoid** | Pointer to **string** | The moid of the Intersight managed object that initiated the workflow. | [optional] 
**InitiatorName** | Pointer to **string** | Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow. | [optional] 
**InitiatorType** | Pointer to **string** | Type of Intersight managed object that initiated the workflow. | [optional] 

## Methods

### NewWorkflowInitiatorContext

`func NewWorkflowInitiatorContext(classId string, objectType string, ) *WorkflowInitiatorContext`

NewWorkflowInitiatorContext instantiates a new WorkflowInitiatorContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInitiatorContextWithDefaults

`func NewWorkflowInitiatorContextWithDefaults() *WorkflowInitiatorContext`

NewWorkflowInitiatorContextWithDefaults instantiates a new WorkflowInitiatorContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowInitiatorContext) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowInitiatorContext) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowInitiatorContext) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowInitiatorContext) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowInitiatorContext) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowInitiatorContext) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInitiatorMoid

`func (o *WorkflowInitiatorContext) GetInitiatorMoid() string`

GetInitiatorMoid returns the InitiatorMoid field if non-nil, zero value otherwise.

### GetInitiatorMoidOk

`func (o *WorkflowInitiatorContext) GetInitiatorMoidOk() (*string, bool)`

GetInitiatorMoidOk returns a tuple with the InitiatorMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorMoid

`func (o *WorkflowInitiatorContext) SetInitiatorMoid(v string)`

SetInitiatorMoid sets InitiatorMoid field to given value.

### HasInitiatorMoid

`func (o *WorkflowInitiatorContext) HasInitiatorMoid() bool`

HasInitiatorMoid returns a boolean if a field has been set.

### GetInitiatorName

`func (o *WorkflowInitiatorContext) GetInitiatorName() string`

GetInitiatorName returns the InitiatorName field if non-nil, zero value otherwise.

### GetInitiatorNameOk

`func (o *WorkflowInitiatorContext) GetInitiatorNameOk() (*string, bool)`

GetInitiatorNameOk returns a tuple with the InitiatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorName

`func (o *WorkflowInitiatorContext) SetInitiatorName(v string)`

SetInitiatorName sets InitiatorName field to given value.

### HasInitiatorName

`func (o *WorkflowInitiatorContext) HasInitiatorName() bool`

HasInitiatorName returns a boolean if a field has been set.

### GetInitiatorType

`func (o *WorkflowInitiatorContext) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *WorkflowInitiatorContext) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *WorkflowInitiatorContext) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *WorkflowInitiatorContext) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


