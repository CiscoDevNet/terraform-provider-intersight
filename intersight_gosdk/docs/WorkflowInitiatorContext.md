# WorkflowInitiatorContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorMoid** | Pointer to **string** | The moid of the Intersigt managed object that initiated the workflow. | [optional] 
**InitiatorName** | Pointer to **string** | Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow. | [optional] 
**InitiatorType** | Pointer to **string** | Type of Intersight managed object that initiated the workflow. | [optional] 

## Methods

### NewWorkflowInitiatorContext

`func NewWorkflowInitiatorContext() *WorkflowInitiatorContext`

NewWorkflowInitiatorContext instantiates a new WorkflowInitiatorContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInitiatorContextWithDefaults

`func NewWorkflowInitiatorContextWithDefaults() *WorkflowInitiatorContext`

NewWorkflowInitiatorContextWithDefaults instantiates a new WorkflowInitiatorContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


