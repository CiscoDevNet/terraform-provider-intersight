# WorkflowTargetDataTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TargetDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TargetDataType"]
**CustomDataTypeProperties** | Pointer to [**NullableWorkflowCustomDataProperty**](WorkflowCustomDataProperty.md) |  | [optional] 
**IsArray** | Pointer to **bool** | When this property is true then an array of targets can be passed as input. | [optional] [default to false]
**Max** | Pointer to **int64** | Specify the maximum value of the array. | [optional] 
**Min** | Pointer to **int64** | Specify the minimum value of the array. | [optional] 
**Properties** | Pointer to [**[]WorkflowTargetProperty**](WorkflowTargetProperty.md) |  | [optional] 

## Methods

### NewWorkflowTargetDataTypeAllOf

`func NewWorkflowTargetDataTypeAllOf(classId string, objectType string, ) *WorkflowTargetDataTypeAllOf`

NewWorkflowTargetDataTypeAllOf instantiates a new WorkflowTargetDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetDataTypeAllOfWithDefaults

`func NewWorkflowTargetDataTypeAllOfWithDefaults() *WorkflowTargetDataTypeAllOf`

NewWorkflowTargetDataTypeAllOfWithDefaults instantiates a new WorkflowTargetDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTargetDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTargetDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTargetDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTargetDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTargetDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTargetDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCustomDataTypeProperties

`func (o *WorkflowTargetDataTypeAllOf) GetCustomDataTypeProperties() WorkflowCustomDataProperty`

GetCustomDataTypeProperties returns the CustomDataTypeProperties field if non-nil, zero value otherwise.

### GetCustomDataTypePropertiesOk

`func (o *WorkflowTargetDataTypeAllOf) GetCustomDataTypePropertiesOk() (*WorkflowCustomDataProperty, bool)`

GetCustomDataTypePropertiesOk returns a tuple with the CustomDataTypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDataTypeProperties

`func (o *WorkflowTargetDataTypeAllOf) SetCustomDataTypeProperties(v WorkflowCustomDataProperty)`

SetCustomDataTypeProperties sets CustomDataTypeProperties field to given value.

### HasCustomDataTypeProperties

`func (o *WorkflowTargetDataTypeAllOf) HasCustomDataTypeProperties() bool`

HasCustomDataTypeProperties returns a boolean if a field has been set.

### SetCustomDataTypePropertiesNil

`func (o *WorkflowTargetDataTypeAllOf) SetCustomDataTypePropertiesNil(b bool)`

 SetCustomDataTypePropertiesNil sets the value for CustomDataTypeProperties to be an explicit nil

### UnsetCustomDataTypeProperties
`func (o *WorkflowTargetDataTypeAllOf) UnsetCustomDataTypeProperties()`

UnsetCustomDataTypeProperties ensures that no value is present for CustomDataTypeProperties, not even an explicit nil
### GetIsArray

`func (o *WorkflowTargetDataTypeAllOf) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *WorkflowTargetDataTypeAllOf) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *WorkflowTargetDataTypeAllOf) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *WorkflowTargetDataTypeAllOf) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetMax

`func (o *WorkflowTargetDataTypeAllOf) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *WorkflowTargetDataTypeAllOf) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *WorkflowTargetDataTypeAllOf) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *WorkflowTargetDataTypeAllOf) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *WorkflowTargetDataTypeAllOf) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *WorkflowTargetDataTypeAllOf) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *WorkflowTargetDataTypeAllOf) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *WorkflowTargetDataTypeAllOf) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTargetDataTypeAllOf) GetProperties() []WorkflowTargetProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTargetDataTypeAllOf) GetPropertiesOk() (*[]WorkflowTargetProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTargetDataTypeAllOf) SetProperties(v []WorkflowTargetProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowTargetDataTypeAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowTargetDataTypeAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowTargetDataTypeAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


