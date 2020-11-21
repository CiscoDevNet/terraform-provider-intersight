# WorkflowTargetContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TargetContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TargetContext"]
**TargetMoid** | Pointer to **string** | Moid of the target Intersight managed object. | [optional] 
**TargetName** | Pointer to **string** | Name of the target instance. | [optional] 
**TargetType** | Pointer to **string** | Object type of the target Intersight managed object. | [optional] 

## Methods

### NewWorkflowTargetContext

`func NewWorkflowTargetContext(classId string, objectType string, ) *WorkflowTargetContext`

NewWorkflowTargetContext instantiates a new WorkflowTargetContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetContextWithDefaults

`func NewWorkflowTargetContextWithDefaults() *WorkflowTargetContext`

NewWorkflowTargetContextWithDefaults instantiates a new WorkflowTargetContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTargetContext) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTargetContext) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTargetContext) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTargetContext) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTargetContext) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTargetContext) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetMoid

`func (o *WorkflowTargetContext) GetTargetMoid() string`

GetTargetMoid returns the TargetMoid field if non-nil, zero value otherwise.

### GetTargetMoidOk

`func (o *WorkflowTargetContext) GetTargetMoidOk() (*string, bool)`

GetTargetMoidOk returns a tuple with the TargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoid

`func (o *WorkflowTargetContext) SetTargetMoid(v string)`

SetTargetMoid sets TargetMoid field to given value.

### HasTargetMoid

`func (o *WorkflowTargetContext) HasTargetMoid() bool`

HasTargetMoid returns a boolean if a field has been set.

### GetTargetName

`func (o *WorkflowTargetContext) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *WorkflowTargetContext) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *WorkflowTargetContext) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *WorkflowTargetContext) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *WorkflowTargetContext) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *WorkflowTargetContext) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *WorkflowTargetContext) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *WorkflowTargetContext) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


