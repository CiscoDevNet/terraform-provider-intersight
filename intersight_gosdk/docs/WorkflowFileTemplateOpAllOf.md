# WorkflowFileTemplateOpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.FileTemplateOp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.FileTemplateOp"]
**TemplateFilePath** | Pointer to **string** | Path of the template file on the connected device. | [optional] 
**TemplateValues** | Pointer to **interface{}** | Input values to render text output file from template file. | [optional] 

## Methods

### NewWorkflowFileTemplateOpAllOf

`func NewWorkflowFileTemplateOpAllOf(classId string, objectType string, ) *WorkflowFileTemplateOpAllOf`

NewWorkflowFileTemplateOpAllOf instantiates a new WorkflowFileTemplateOpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowFileTemplateOpAllOfWithDefaults

`func NewWorkflowFileTemplateOpAllOfWithDefaults() *WorkflowFileTemplateOpAllOf`

NewWorkflowFileTemplateOpAllOfWithDefaults instantiates a new WorkflowFileTemplateOpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowFileTemplateOpAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowFileTemplateOpAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowFileTemplateOpAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowFileTemplateOpAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowFileTemplateOpAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowFileTemplateOpAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTemplateFilePath

`func (o *WorkflowFileTemplateOpAllOf) GetTemplateFilePath() string`

GetTemplateFilePath returns the TemplateFilePath field if non-nil, zero value otherwise.

### GetTemplateFilePathOk

`func (o *WorkflowFileTemplateOpAllOf) GetTemplateFilePathOk() (*string, bool)`

GetTemplateFilePathOk returns a tuple with the TemplateFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFilePath

`func (o *WorkflowFileTemplateOpAllOf) SetTemplateFilePath(v string)`

SetTemplateFilePath sets TemplateFilePath field to given value.

### HasTemplateFilePath

`func (o *WorkflowFileTemplateOpAllOf) HasTemplateFilePath() bool`

HasTemplateFilePath returns a boolean if a field has been set.

### GetTemplateValues

`func (o *WorkflowFileTemplateOpAllOf) GetTemplateValues() interface{}`

GetTemplateValues returns the TemplateValues field if non-nil, zero value otherwise.

### GetTemplateValuesOk

`func (o *WorkflowFileTemplateOpAllOf) GetTemplateValuesOk() (*interface{}, bool)`

GetTemplateValuesOk returns a tuple with the TemplateValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateValues

`func (o *WorkflowFileTemplateOpAllOf) SetTemplateValues(v interface{})`

SetTemplateValues sets TemplateValues field to given value.

### HasTemplateValues

`func (o *WorkflowFileTemplateOpAllOf) HasTemplateValues() bool`

HasTemplateValues returns a boolean if a field has been set.

### SetTemplateValuesNil

`func (o *WorkflowFileTemplateOpAllOf) SetTemplateValuesNil(b bool)`

 SetTemplateValuesNil sets the value for TemplateValues to be an explicit nil

### UnsetTemplateValues
`func (o *WorkflowFileTemplateOpAllOf) UnsetTemplateValues()`

UnsetTemplateValues ensures that no value is present for TemplateValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


