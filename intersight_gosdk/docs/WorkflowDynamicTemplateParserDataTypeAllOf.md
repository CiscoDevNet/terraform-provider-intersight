# WorkflowDynamicTemplateParserDataTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DynamicTemplateParserDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DynamicTemplateParserDataType"]
**IsTemplateSecure** | Pointer to **bool** | When set to true, the template is marked as secure and the content is encrypted and stored. | [optional] 
**TemplateType** | Pointer to **string** | Template type decides on the API to be used to fetch the placeholders present inside the template. * &#x60;OsInstall&#x60; - This refers to the OS configuration template MO. | [optional] [default to "OsInstall"]

## Methods

### NewWorkflowDynamicTemplateParserDataTypeAllOf

`func NewWorkflowDynamicTemplateParserDataTypeAllOf(classId string, objectType string, ) *WorkflowDynamicTemplateParserDataTypeAllOf`

NewWorkflowDynamicTemplateParserDataTypeAllOf instantiates a new WorkflowDynamicTemplateParserDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDynamicTemplateParserDataTypeAllOfWithDefaults

`func NewWorkflowDynamicTemplateParserDataTypeAllOfWithDefaults() *WorkflowDynamicTemplateParserDataTypeAllOf`

NewWorkflowDynamicTemplateParserDataTypeAllOfWithDefaults instantiates a new WorkflowDynamicTemplateParserDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsTemplateSecure

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetIsTemplateSecure() bool`

GetIsTemplateSecure returns the IsTemplateSecure field if non-nil, zero value otherwise.

### GetIsTemplateSecureOk

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetIsTemplateSecureOk() (*bool, bool)`

GetIsTemplateSecureOk returns a tuple with the IsTemplateSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplateSecure

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetIsTemplateSecure(v bool)`

SetIsTemplateSecure sets IsTemplateSecure field to given value.

### HasIsTemplateSecure

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) HasIsTemplateSecure() bool`

HasIsTemplateSecure returns a boolean if a field has been set.

### GetTemplateType

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *WorkflowDynamicTemplateParserDataTypeAllOf) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


