# WorkflowEnumEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote (&#39;), forward slash (/), or an underscore (_). | [optional] 
**Value** | Pointer to **string** | Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_). | [optional] 

## Methods

### NewWorkflowEnumEntry

`func NewWorkflowEnumEntry() *WorkflowEnumEntry`

NewWorkflowEnumEntry instantiates a new WorkflowEnumEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowEnumEntryWithDefaults

`func NewWorkflowEnumEntryWithDefaults() *WorkflowEnumEntry`

NewWorkflowEnumEntryWithDefaults instantiates a new WorkflowEnumEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *WorkflowEnumEntry) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowEnumEntry) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowEnumEntry) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowEnumEntry) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowEnumEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowEnumEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowEnumEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowEnumEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


