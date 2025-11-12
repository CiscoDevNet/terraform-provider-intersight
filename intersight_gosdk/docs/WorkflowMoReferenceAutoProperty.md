# WorkflowMoReferenceAutoProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MoReferenceAutoProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MoReferenceAutoProperty"]
**DisplayAttributes** | Pointer to **[]string** |  | [optional] 
**Filters** | Pointer to **[]string** |  | [optional] 
**OrderBy** | Pointer to **string** | Determines  properties that are used to sort the collection of resources. | [optional] 
**Rule** | Pointer to [**NullableWorkflowAbstractResourceSelector**](WorkflowAbstractResourceSelector.md) |  | [optional] 

## Methods

### NewWorkflowMoReferenceAutoProperty

`func NewWorkflowMoReferenceAutoProperty(classId string, objectType string, ) *WorkflowMoReferenceAutoProperty`

NewWorkflowMoReferenceAutoProperty instantiates a new WorkflowMoReferenceAutoProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoReferenceAutoPropertyWithDefaults

`func NewWorkflowMoReferenceAutoPropertyWithDefaults() *WorkflowMoReferenceAutoProperty`

NewWorkflowMoReferenceAutoPropertyWithDefaults instantiates a new WorkflowMoReferenceAutoProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoReferenceAutoProperty) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoReferenceAutoProperty) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoReferenceAutoProperty) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoReferenceAutoProperty) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoReferenceAutoProperty) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoReferenceAutoProperty) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplayAttributes

`func (o *WorkflowMoReferenceAutoProperty) GetDisplayAttributes() []string`

GetDisplayAttributes returns the DisplayAttributes field if non-nil, zero value otherwise.

### GetDisplayAttributesOk

`func (o *WorkflowMoReferenceAutoProperty) GetDisplayAttributesOk() (*[]string, bool)`

GetDisplayAttributesOk returns a tuple with the DisplayAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAttributes

`func (o *WorkflowMoReferenceAutoProperty) SetDisplayAttributes(v []string)`

SetDisplayAttributes sets DisplayAttributes field to given value.

### HasDisplayAttributes

`func (o *WorkflowMoReferenceAutoProperty) HasDisplayAttributes() bool`

HasDisplayAttributes returns a boolean if a field has been set.

### SetDisplayAttributesNil

`func (o *WorkflowMoReferenceAutoProperty) SetDisplayAttributesNil(b bool)`

 SetDisplayAttributesNil sets the value for DisplayAttributes to be an explicit nil

### UnsetDisplayAttributes
`func (o *WorkflowMoReferenceAutoProperty) UnsetDisplayAttributes()`

UnsetDisplayAttributes ensures that no value is present for DisplayAttributes, not even an explicit nil
### GetFilters

`func (o *WorkflowMoReferenceAutoProperty) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WorkflowMoReferenceAutoProperty) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WorkflowMoReferenceAutoProperty) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *WorkflowMoReferenceAutoProperty) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *WorkflowMoReferenceAutoProperty) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *WorkflowMoReferenceAutoProperty) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetOrderBy

`func (o *WorkflowMoReferenceAutoProperty) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *WorkflowMoReferenceAutoProperty) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *WorkflowMoReferenceAutoProperty) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *WorkflowMoReferenceAutoProperty) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetRule

`func (o *WorkflowMoReferenceAutoProperty) GetRule() WorkflowAbstractResourceSelector`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *WorkflowMoReferenceAutoProperty) GetRuleOk() (*WorkflowAbstractResourceSelector, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *WorkflowMoReferenceAutoProperty) SetRule(v WorkflowAbstractResourceSelector)`

SetRule sets Rule field to given value.

### HasRule

`func (o *WorkflowMoReferenceAutoProperty) HasRule() bool`

HasRule returns a boolean if a field has been set.

### SetRuleNil

`func (o *WorkflowMoReferenceAutoProperty) SetRuleNil(b bool)`

 SetRuleNil sets the value for Rule to be an explicit nil

### UnsetRule
`func (o *WorkflowMoReferenceAutoProperty) UnsetRule()`

UnsetRule ensures that no value is present for Rule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


