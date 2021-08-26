# WorkflowArrayDataTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ArrayDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ArrayDataType"]
**ArrayItemType** | Pointer to [**NullableWorkflowArrayItem**](WorkflowArrayItem.md) |  | [optional] 
**Max** | Pointer to **int64** | Specify the maximum value of the array. | [optional] 
**Min** | Pointer to **int64** | Specify the minimum value of the array. | [optional] 

## Methods

### NewWorkflowArrayDataTypeAllOf

`func NewWorkflowArrayDataTypeAllOf(classId string, objectType string, ) *WorkflowArrayDataTypeAllOf`

NewWorkflowArrayDataTypeAllOf instantiates a new WorkflowArrayDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowArrayDataTypeAllOfWithDefaults

`func NewWorkflowArrayDataTypeAllOfWithDefaults() *WorkflowArrayDataTypeAllOf`

NewWorkflowArrayDataTypeAllOfWithDefaults instantiates a new WorkflowArrayDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowArrayDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowArrayDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowArrayDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowArrayDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowArrayDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowArrayDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayItemType

`func (o *WorkflowArrayDataTypeAllOf) GetArrayItemType() WorkflowArrayItem`

GetArrayItemType returns the ArrayItemType field if non-nil, zero value otherwise.

### GetArrayItemTypeOk

`func (o *WorkflowArrayDataTypeAllOf) GetArrayItemTypeOk() (*WorkflowArrayItem, bool)`

GetArrayItemTypeOk returns a tuple with the ArrayItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayItemType

`func (o *WorkflowArrayDataTypeAllOf) SetArrayItemType(v WorkflowArrayItem)`

SetArrayItemType sets ArrayItemType field to given value.

### HasArrayItemType

`func (o *WorkflowArrayDataTypeAllOf) HasArrayItemType() bool`

HasArrayItemType returns a boolean if a field has been set.

### SetArrayItemTypeNil

`func (o *WorkflowArrayDataTypeAllOf) SetArrayItemTypeNil(b bool)`

 SetArrayItemTypeNil sets the value for ArrayItemType to be an explicit nil

### UnsetArrayItemType
`func (o *WorkflowArrayDataTypeAllOf) UnsetArrayItemType()`

UnsetArrayItemType ensures that no value is present for ArrayItemType, not even an explicit nil
### GetMax

`func (o *WorkflowArrayDataTypeAllOf) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WorkflowArrayDataTypeAllOf) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *WorkflowArrayDataTypeAllOf) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *WorkflowArrayDataTypeAllOf) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *WorkflowArrayDataTypeAllOf) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WorkflowArrayDataTypeAllOf) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *WorkflowArrayDataTypeAllOf) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *WorkflowArrayDataTypeAllOf) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


