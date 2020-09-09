# WorkflowBaseDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to [**WorkflowDefaultValue**](workflow.DefaultValue.md) |  | [optional] 
**Description** | Pointer to **string** | Provide a detailed description of the data type. | [optional] 
**Label** | Pointer to **string** | Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. | [optional] 
**Name** | Pointer to **string** | Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. | [optional] 
**Required** | Pointer to **bool** | Specifies whether this parameter is required. The field is applicable for task and workflow. | [optional] 

## Methods

### NewWorkflowBaseDataType

`func NewWorkflowBaseDataType() *WorkflowBaseDataType`

NewWorkflowBaseDataType instantiates a new WorkflowBaseDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBaseDataTypeWithDefaults

`func NewWorkflowBaseDataTypeWithDefaults() *WorkflowBaseDataType`

NewWorkflowBaseDataTypeWithDefaults instantiates a new WorkflowBaseDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *WorkflowBaseDataType) GetDefault() WorkflowDefaultValue`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WorkflowBaseDataType) GetDefaultOk() (*WorkflowDefaultValue, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WorkflowBaseDataType) SetDefault(v WorkflowDefaultValue)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WorkflowBaseDataType) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowBaseDataType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowBaseDataType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowBaseDataType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowBaseDataType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowBaseDataType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowBaseDataType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowBaseDataType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowBaseDataType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowBaseDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBaseDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBaseDataType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowBaseDataType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *WorkflowBaseDataType) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *WorkflowBaseDataType) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *WorkflowBaseDataType) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *WorkflowBaseDataType) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


