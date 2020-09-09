# WorkflowTargetDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDataTypeProperties** | Pointer to [**WorkflowCustomDataProperty**](workflow.CustomDataProperty.md) |  | [optional] 
**IsArray** | Pointer to **bool** | When this property is true then an array of targets can be passed as input. | [optional] 
**Max** | Pointer to **int64** | Specify the maximum value of the array. | [optional] 
**Min** | Pointer to **int64** | Specify the minimum value of the array. | [optional] 
**Properties** | Pointer to [**[]WorkflowTargetProperty**](workflow.TargetProperty.md) |  | [optional] 

## Methods

### NewWorkflowTargetDataType

`func NewWorkflowTargetDataType() *WorkflowTargetDataType`

NewWorkflowTargetDataType instantiates a new WorkflowTargetDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetDataTypeWithDefaults

`func NewWorkflowTargetDataTypeWithDefaults() *WorkflowTargetDataType`

NewWorkflowTargetDataTypeWithDefaults instantiates a new WorkflowTargetDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDataTypeProperties

`func (o *WorkflowTargetDataType) GetCustomDataTypeProperties() WorkflowCustomDataProperty`

GetCustomDataTypeProperties returns the CustomDataTypeProperties field if non-nil, zero value otherwise.

### GetCustomDataTypePropertiesOk

`func (o *WorkflowTargetDataType) GetCustomDataTypePropertiesOk() (*WorkflowCustomDataProperty, bool)`

GetCustomDataTypePropertiesOk returns a tuple with the CustomDataTypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDataTypeProperties

`func (o *WorkflowTargetDataType) SetCustomDataTypeProperties(v WorkflowCustomDataProperty)`

SetCustomDataTypeProperties sets CustomDataTypeProperties field to given value.

### HasCustomDataTypeProperties

`func (o *WorkflowTargetDataType) HasCustomDataTypeProperties() bool`

HasCustomDataTypeProperties returns a boolean if a field has been set.

### GetIsArray

`func (o *WorkflowTargetDataType) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *WorkflowTargetDataType) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *WorkflowTargetDataType) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *WorkflowTargetDataType) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetMax

`func (o *WorkflowTargetDataType) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WorkflowTargetDataType) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *WorkflowTargetDataType) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *WorkflowTargetDataType) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *WorkflowTargetDataType) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WorkflowTargetDataType) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *WorkflowTargetDataType) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *WorkflowTargetDataType) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTargetDataType) GetProperties() []WorkflowTargetProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTargetDataType) GetPropertiesOk() (*[]WorkflowTargetProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTargetDataType) SetProperties(v []WorkflowTargetProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowTargetDataType) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


