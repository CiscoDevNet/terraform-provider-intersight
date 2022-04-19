# WorkflowTemplateParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TemplateParser"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TemplateParser"]
**Placeholders** | Pointer to [**[]WorkflowPrimitiveDataType**](WorkflowPrimitiveDataType.md) |  | [optional] 
**TemplateContent** | Pointer to **string** | The content of the entire template file. The content can either be a static or dynamic file that can contain placeholders. The placeholders are expected to conform to the golang template syntax i.e. it must be provided inside {{ }}. | [optional] 

## Methods

### NewWorkflowTemplateParser

`func NewWorkflowTemplateParser(classId string, objectType string, ) *WorkflowTemplateParser`

NewWorkflowTemplateParser instantiates a new WorkflowTemplateParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateParserWithDefaults

`func NewWorkflowTemplateParserWithDefaults() *WorkflowTemplateParser`

NewWorkflowTemplateParserWithDefaults instantiates a new WorkflowTemplateParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTemplateParser) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTemplateParser) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTemplateParser) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTemplateParser) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTemplateParser) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTemplateParser) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPlaceholders

`func (o *WorkflowTemplateParser) GetPlaceholders() []WorkflowPrimitiveDataType`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *WorkflowTemplateParser) GetPlaceholdersOk() (*[]WorkflowPrimitiveDataType, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *WorkflowTemplateParser) SetPlaceholders(v []WorkflowPrimitiveDataType)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *WorkflowTemplateParser) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### SetPlaceholdersNil

`func (o *WorkflowTemplateParser) SetPlaceholdersNil(b bool)`

 SetPlaceholdersNil sets the value for Placeholders to be an explicit nil

### UnsetPlaceholders
`func (o *WorkflowTemplateParser) UnsetPlaceholders()`

UnsetPlaceholders ensures that no value is present for Placeholders, not even an explicit nil
### GetTemplateContent

`func (o *WorkflowTemplateParser) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *WorkflowTemplateParser) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *WorkflowTemplateParser) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *WorkflowTemplateParser) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


