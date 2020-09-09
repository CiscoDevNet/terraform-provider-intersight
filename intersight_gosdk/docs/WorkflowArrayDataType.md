# WorkflowArrayDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrayItemType** | Pointer to [**WorkflowArrayItem**](workflow.ArrayItem.md) |  | [optional] 
**Max** | Pointer to **int64** | Specify the maximum value of the array. | [optional] 
**Min** | Pointer to **int64** | Specify the minimum value of the array. | [optional] 

## Methods

### NewWorkflowArrayDataType

`func NewWorkflowArrayDataType() *WorkflowArrayDataType`

NewWorkflowArrayDataType instantiates a new WorkflowArrayDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowArrayDataTypeWithDefaults

`func NewWorkflowArrayDataTypeWithDefaults() *WorkflowArrayDataType`

NewWorkflowArrayDataTypeWithDefaults instantiates a new WorkflowArrayDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrayItemType

`func (o *WorkflowArrayDataType) GetArrayItemType() WorkflowArrayItem`

GetArrayItemType returns the ArrayItemType field if non-nil, zero value otherwise.

### GetArrayItemTypeOk

`func (o *WorkflowArrayDataType) GetArrayItemTypeOk() (*WorkflowArrayItem, bool)`

GetArrayItemTypeOk returns a tuple with the ArrayItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayItemType

`func (o *WorkflowArrayDataType) SetArrayItemType(v WorkflowArrayItem)`

SetArrayItemType sets ArrayItemType field to given value.

### HasArrayItemType

`func (o *WorkflowArrayDataType) HasArrayItemType() bool`

HasArrayItemType returns a boolean if a field has been set.

### GetMax

`func (o *WorkflowArrayDataType) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WorkflowArrayDataType) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *WorkflowArrayDataType) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *WorkflowArrayDataType) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *WorkflowArrayDataType) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WorkflowArrayDataType) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *WorkflowArrayDataType) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *WorkflowArrayDataType) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


