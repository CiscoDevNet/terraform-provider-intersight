# WorkflowSelectorProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SelectorProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SelectorProperty"]
**Body** | Pointer to **interface{}** | Content of the request body to send for POST request. | [optional] 
**Method** | Pointer to **string** | The HTTP method to be used. * &#x60;GET&#x60; - The HTTP GET method requests a representation of the specified resource. * &#x60;POST&#x60; - The HTTP POST method sends data to the server. | [optional] [default to "GET"]

## Methods

### NewWorkflowSelectorProperty

`func NewWorkflowSelectorProperty(classId string, objectType string, ) *WorkflowSelectorProperty`

NewWorkflowSelectorProperty instantiates a new WorkflowSelectorProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSelectorPropertyWithDefaults

`func NewWorkflowSelectorPropertyWithDefaults() *WorkflowSelectorProperty`

NewWorkflowSelectorPropertyWithDefaults instantiates a new WorkflowSelectorProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSelectorProperty) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSelectorProperty) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSelectorProperty) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSelectorProperty) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSelectorProperty) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSelectorProperty) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *WorkflowSelectorProperty) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WorkflowSelectorProperty) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WorkflowSelectorProperty) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *WorkflowSelectorProperty) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *WorkflowSelectorProperty) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *WorkflowSelectorProperty) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetMethod

`func (o *WorkflowSelectorProperty) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WorkflowSelectorProperty) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WorkflowSelectorProperty) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WorkflowSelectorProperty) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


