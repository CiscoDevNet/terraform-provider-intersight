# WorkflowMoReferenceAutoPropertyAllOf

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

### NewWorkflowMoReferenceAutoPropertyAllOf

`func NewWorkflowMoReferenceAutoPropertyAllOf(classId string, objectType string, ) *WorkflowMoReferenceAutoPropertyAllOf`

NewWorkflowMoReferenceAutoPropertyAllOf instantiates a new WorkflowMoReferenceAutoPropertyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoReferenceAutoPropertyAllOfWithDefaults

`func NewWorkflowMoReferenceAutoPropertyAllOfWithDefaults() *WorkflowMoReferenceAutoPropertyAllOf`

NewWorkflowMoReferenceAutoPropertyAllOfWithDefaults instantiates a new WorkflowMoReferenceAutoPropertyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplayAttributes

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetDisplayAttributes() []string`

GetDisplayAttributes returns the DisplayAttributes field if non-nil, zero value otherwise.

### GetDisplayAttributesOk

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetDisplayAttributesOk() (*[]string, bool)`

GetDisplayAttributesOk returns a tuple with the DisplayAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAttributes

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetDisplayAttributes(v []string)`

SetDisplayAttributes sets DisplayAttributes field to given value.

### HasDisplayAttributes

`func (o *WorkflowMoReferenceAutoPropertyAllOf) HasDisplayAttributes() bool`

HasDisplayAttributes returns a boolean if a field has been set.

### SetDisplayAttributesNil

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetDisplayAttributesNil(b bool)`

 SetDisplayAttributesNil sets the value for DisplayAttributes to be an explicit nil

### UnsetDisplayAttributes
`func (o *WorkflowMoReferenceAutoPropertyAllOf) UnsetDisplayAttributes()`

UnsetDisplayAttributes ensures that no value is present for DisplayAttributes, not even an explicit nil
### GetFilters

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *WorkflowMoReferenceAutoPropertyAllOf) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *WorkflowMoReferenceAutoPropertyAllOf) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetOrderBy

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *WorkflowMoReferenceAutoPropertyAllOf) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetRule

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetRule() WorkflowAbstractResourceSelector`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *WorkflowMoReferenceAutoPropertyAllOf) GetRuleOk() (*WorkflowAbstractResourceSelector, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetRule(v WorkflowAbstractResourceSelector)`

SetRule sets Rule field to given value.

### HasRule

`func (o *WorkflowMoReferenceAutoPropertyAllOf) HasRule() bool`

HasRule returns a boolean if a field has been set.

### SetRuleNil

`func (o *WorkflowMoReferenceAutoPropertyAllOf) SetRuleNil(b bool)`

 SetRuleNil sets the value for Rule to be an explicit nil

### UnsetRule
`func (o *WorkflowMoReferenceAutoPropertyAllOf) UnsetRule()`

UnsetRule ensures that no value is present for Rule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


