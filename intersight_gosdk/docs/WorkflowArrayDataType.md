# WorkflowArrayDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ArrayDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ArrayDataType"]
**ArrayItemType** | Pointer to [**NullableWorkflowArrayItem**](WorkflowArrayItem.md) |  | [optional] 
**Max** | Pointer to **int64** | Specify the maximum value of the array. | [optional] 
**Min** | Pointer to **int64** | Specify the minimum value of the array. | [optional] 

## Methods

### NewWorkflowArrayDataType

`func NewWorkflowArrayDataType(classId string, objectType string, ) *WorkflowArrayDataType`

NewWorkflowArrayDataType instantiates a new WorkflowArrayDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowArrayDataTypeWithDefaults

`func NewWorkflowArrayDataTypeWithDefaults() *WorkflowArrayDataType`

NewWorkflowArrayDataTypeWithDefaults instantiates a new WorkflowArrayDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowArrayDataType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowArrayDataType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowArrayDataType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowArrayDataType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowArrayDataType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowArrayDataType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayItemType

`func (o *WorkflowArrayDataType) GetArrayItemType() WorkflowArrayItem`

GetArrayItemType returns the ArrayItemType field if non-nil, zero value otherwise.

### GetArrayItemTypeOk

`func (o *WorkflowArrayDataType) GetArrayItemTypeOk() (*WorkflowArrayItem, bool)`

GetArrayItemTypeOk returns a tuple with the ArrayItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayItemType

`func (o *WorkflowArrayDataType) SetArrayItemType(v WorkflowArrayItem)`

SetArrayItemType sets ArrayItemType field to given value.

### HasArrayItemType

`func (o *WorkflowArrayDataType) HasArrayItemType() bool`

HasArrayItemType returns a boolean if a field has been set.

### SetArrayItemTypeNil

`func (o *WorkflowArrayDataType) SetArrayItemTypeNil(b bool)`

 SetArrayItemTypeNil sets the value for ArrayItemType to be an explicit nil

### UnsetArrayItemType
`func (o *WorkflowArrayDataType) UnsetArrayItemType()`

UnsetArrayItemType ensures that no value is present for ArrayItemType, not even an explicit nil
### GetMax

`func (o *WorkflowArrayDataType) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WorkflowArrayDataType) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *WorkflowArrayDataType) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *WorkflowArrayDataType) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *WorkflowArrayDataType) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WorkflowArrayDataType) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *WorkflowArrayDataType) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *WorkflowArrayDataType) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


