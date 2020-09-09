# WorkflowConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnumList** | Pointer to [**[]WorkflowEnumEntry**](workflow.EnumEntry.md) |  | [optional] 
**Max** | Pointer to **float64** | Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. | [optional] 
**Min** | Pointer to **float64** | Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. | [optional] 
**Regex** | Pointer to **string** | When the parameter is a string this regular expression is used to ensure the value is valid. | [optional] 

## Methods

### NewWorkflowConstraints

`func NewWorkflowConstraints() *WorkflowConstraints`

NewWorkflowConstraints instantiates a new WorkflowConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowConstraintsWithDefaults

`func NewWorkflowConstraintsWithDefaults() *WorkflowConstraints`

NewWorkflowConstraintsWithDefaults instantiates a new WorkflowConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnumList

`func (o *WorkflowConstraints) GetEnumList() []WorkflowEnumEntry`

GetEnumList returns the EnumList field if non-nil, zero value otherwise.

### GetEnumListOk

`func (o *WorkflowConstraints) GetEnumListOk() (*[]WorkflowEnumEntry, bool)`

GetEnumListOk returns a tuple with the EnumList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumList

`func (o *WorkflowConstraints) SetEnumList(v []WorkflowEnumEntry)`

SetEnumList sets EnumList field to given value.

### HasEnumList

`func (o *WorkflowConstraints) HasEnumList() bool`

HasEnumList returns a boolean if a field has been set.

### GetMax

`func (o *WorkflowConstraints) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WorkflowConstraints) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *WorkflowConstraints) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *WorkflowConstraints) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *WorkflowConstraints) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WorkflowConstraints) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *WorkflowConstraints) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *WorkflowConstraints) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetRegex

`func (o *WorkflowConstraints) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *WorkflowConstraints) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *WorkflowConstraints) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *WorkflowConstraints) HasRegex() bool`

HasRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


