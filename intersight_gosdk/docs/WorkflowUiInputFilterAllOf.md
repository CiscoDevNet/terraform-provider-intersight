# WorkflowUiInputFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.UiInputFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.UiInputFilter"]
**Filters** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name for the input definition to which this filter applies. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. When defining the cascade filter for a sub property, use a period (.) to seperate each section of the name like \&quot;StorageConfig.Volume\&quot; where &#39;StorageConfig&#39; is an input name and &#39;Volume&#39; is a sub property defined through custom data type definition. | [optional] 
**UserHelpMessage** | Pointer to **string** | Help message shown to the user about which prior input needs to be selected to enable the input mapped to this filter. | [optional] 

## Methods

### NewWorkflowUiInputFilterAllOf

`func NewWorkflowUiInputFilterAllOf(classId string, objectType string, ) *WorkflowUiInputFilterAllOf`

NewWorkflowUiInputFilterAllOf instantiates a new WorkflowUiInputFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUiInputFilterAllOfWithDefaults

`func NewWorkflowUiInputFilterAllOfWithDefaults() *WorkflowUiInputFilterAllOf`

NewWorkflowUiInputFilterAllOfWithDefaults instantiates a new WorkflowUiInputFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowUiInputFilterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowUiInputFilterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowUiInputFilterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowUiInputFilterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowUiInputFilterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowUiInputFilterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilters

`func (o *WorkflowUiInputFilterAllOf) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WorkflowUiInputFilterAllOf) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WorkflowUiInputFilterAllOf) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *WorkflowUiInputFilterAllOf) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *WorkflowUiInputFilterAllOf) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *WorkflowUiInputFilterAllOf) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetName

`func (o *WorkflowUiInputFilterAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowUiInputFilterAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowUiInputFilterAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowUiInputFilterAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserHelpMessage

`func (o *WorkflowUiInputFilterAllOf) GetUserHelpMessage() string`

GetUserHelpMessage returns the UserHelpMessage field if non-nil, zero value otherwise.

### GetUserHelpMessageOk

`func (o *WorkflowUiInputFilterAllOf) GetUserHelpMessageOk() (*string, bool)`

GetUserHelpMessageOk returns a tuple with the UserHelpMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHelpMessage

`func (o *WorkflowUiInputFilterAllOf) SetUserHelpMessage(v string)`

SetUserHelpMessage sets UserHelpMessage field to given value.

### HasUserHelpMessage

`func (o *WorkflowUiInputFilterAllOf) HasUserHelpMessage() bool`

HasUserHelpMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


