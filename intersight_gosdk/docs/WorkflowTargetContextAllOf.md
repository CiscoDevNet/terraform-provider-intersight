# WorkflowTargetContextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TargetContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TargetContext"]
**TargetMoid** | Pointer to **string** | Moid of the target Intersight managed object. | [optional] [readonly] 
**TargetName** | Pointer to **string** | Name of the target instance. | [optional] [readonly] 
**TargetType** | Pointer to **string** | Object type of the target Intersight managed object. | [optional] [readonly] 

## Methods

### NewWorkflowTargetContextAllOf

`func NewWorkflowTargetContextAllOf(classId string, objectType string, ) *WorkflowTargetContextAllOf`

NewWorkflowTargetContextAllOf instantiates a new WorkflowTargetContextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetContextAllOfWithDefaults

`func NewWorkflowTargetContextAllOfWithDefaults() *WorkflowTargetContextAllOf`

NewWorkflowTargetContextAllOfWithDefaults instantiates a new WorkflowTargetContextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTargetContextAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTargetContextAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTargetContextAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTargetContextAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTargetContextAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTargetContextAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetMoid

`func (o *WorkflowTargetContextAllOf) GetTargetMoid() string`

GetTargetMoid returns the TargetMoid field if non-nil, zero value otherwise.

### GetTargetMoidOk

`func (o *WorkflowTargetContextAllOf) GetTargetMoidOk() (*string, bool)`

GetTargetMoidOk returns a tuple with the TargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoid

`func (o *WorkflowTargetContextAllOf) SetTargetMoid(v string)`

SetTargetMoid sets TargetMoid field to given value.

### HasTargetMoid

`func (o *WorkflowTargetContextAllOf) HasTargetMoid() bool`

HasTargetMoid returns a boolean if a field has been set.

### GetTargetName

`func (o *WorkflowTargetContextAllOf) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *WorkflowTargetContextAllOf) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *WorkflowTargetContextAllOf) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *WorkflowTargetContextAllOf) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *WorkflowTargetContextAllOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *WorkflowTargetContextAllOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *WorkflowTargetContextAllOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *WorkflowTargetContextAllOf) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


