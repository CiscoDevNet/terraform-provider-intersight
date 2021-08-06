# WorkflowSelectorPropertyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SelectorProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SelectorProperty"]
**Body** | Pointer to **interface{}** | Content of the request body to send for POST request. | [optional] 
**Method** | Pointer to **string** | The HTTP method to be used. * &#x60;GET&#x60; - The HTTP GET method requests a representation of the specified resource. * &#x60;POST&#x60; - The HTTP POST method sends data to the server. | [optional] [default to "GET"]

## Methods

### NewWorkflowSelectorPropertyAllOf

`func NewWorkflowSelectorPropertyAllOf(classId string, objectType string, ) *WorkflowSelectorPropertyAllOf`

NewWorkflowSelectorPropertyAllOf instantiates a new WorkflowSelectorPropertyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSelectorPropertyAllOfWithDefaults

`func NewWorkflowSelectorPropertyAllOfWithDefaults() *WorkflowSelectorPropertyAllOf`

NewWorkflowSelectorPropertyAllOfWithDefaults instantiates a new WorkflowSelectorPropertyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSelectorPropertyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSelectorPropertyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSelectorPropertyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSelectorPropertyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSelectorPropertyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSelectorPropertyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *WorkflowSelectorPropertyAllOf) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WorkflowSelectorPropertyAllOf) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WorkflowSelectorPropertyAllOf) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *WorkflowSelectorPropertyAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *WorkflowSelectorPropertyAllOf) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *WorkflowSelectorPropertyAllOf) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetMethod

`func (o *WorkflowSelectorPropertyAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WorkflowSelectorPropertyAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WorkflowSelectorPropertyAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WorkflowSelectorPropertyAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


