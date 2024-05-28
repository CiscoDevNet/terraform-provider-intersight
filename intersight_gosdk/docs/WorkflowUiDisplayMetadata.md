# WorkflowUiDisplayMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.UiDisplayMetadata"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.UiDisplayMetadata"]
**UiFormConfigs** | Pointer to [**[]WorkflowUiFormConfig**](WorkflowUiFormConfig.md) |  | [optional] 
**UiViewConfigs** | Pointer to [**[]WorkflowUiViewConfig**](WorkflowUiViewConfig.md) |  | [optional] 
**MetaDefinition** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewWorkflowUiDisplayMetadata

`func NewWorkflowUiDisplayMetadata(classId string, objectType string, ) *WorkflowUiDisplayMetadata`

NewWorkflowUiDisplayMetadata instantiates a new WorkflowUiDisplayMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUiDisplayMetadataWithDefaults

`func NewWorkflowUiDisplayMetadataWithDefaults() *WorkflowUiDisplayMetadata`

NewWorkflowUiDisplayMetadataWithDefaults instantiates a new WorkflowUiDisplayMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowUiDisplayMetadata) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowUiDisplayMetadata) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowUiDisplayMetadata) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowUiDisplayMetadata) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowUiDisplayMetadata) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowUiDisplayMetadata) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUiFormConfigs

`func (o *WorkflowUiDisplayMetadata) GetUiFormConfigs() []WorkflowUiFormConfig`

GetUiFormConfigs returns the UiFormConfigs field if non-nil, zero value otherwise.

### GetUiFormConfigsOk

`func (o *WorkflowUiDisplayMetadata) GetUiFormConfigsOk() (*[]WorkflowUiFormConfig, bool)`

GetUiFormConfigsOk returns a tuple with the UiFormConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiFormConfigs

`func (o *WorkflowUiDisplayMetadata) SetUiFormConfigs(v []WorkflowUiFormConfig)`

SetUiFormConfigs sets UiFormConfigs field to given value.

### HasUiFormConfigs

`func (o *WorkflowUiDisplayMetadata) HasUiFormConfigs() bool`

HasUiFormConfigs returns a boolean if a field has been set.

### SetUiFormConfigsNil

`func (o *WorkflowUiDisplayMetadata) SetUiFormConfigsNil(b bool)`

 SetUiFormConfigsNil sets the value for UiFormConfigs to be an explicit nil

### UnsetUiFormConfigs
`func (o *WorkflowUiDisplayMetadata) UnsetUiFormConfigs()`

UnsetUiFormConfigs ensures that no value is present for UiFormConfigs, not even an explicit nil
### GetUiViewConfigs

`func (o *WorkflowUiDisplayMetadata) GetUiViewConfigs() []WorkflowUiViewConfig`

GetUiViewConfigs returns the UiViewConfigs field if non-nil, zero value otherwise.

### GetUiViewConfigsOk

`func (o *WorkflowUiDisplayMetadata) GetUiViewConfigsOk() (*[]WorkflowUiViewConfig, bool)`

GetUiViewConfigsOk returns a tuple with the UiViewConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiViewConfigs

`func (o *WorkflowUiDisplayMetadata) SetUiViewConfigs(v []WorkflowUiViewConfig)`

SetUiViewConfigs sets UiViewConfigs field to given value.

### HasUiViewConfigs

`func (o *WorkflowUiDisplayMetadata) HasUiViewConfigs() bool`

HasUiViewConfigs returns a boolean if a field has been set.

### SetUiViewConfigsNil

`func (o *WorkflowUiDisplayMetadata) SetUiViewConfigsNil(b bool)`

 SetUiViewConfigsNil sets the value for UiViewConfigs to be an explicit nil

### UnsetUiViewConfigs
`func (o *WorkflowUiDisplayMetadata) UnsetUiViewConfigs()`

UnsetUiViewConfigs ensures that no value is present for UiViewConfigs, not even an explicit nil
### GetMetaDefinition

`func (o *WorkflowUiDisplayMetadata) GetMetaDefinition() MoBaseMoRelationship`

GetMetaDefinition returns the MetaDefinition field if non-nil, zero value otherwise.

### GetMetaDefinitionOk

`func (o *WorkflowUiDisplayMetadata) GetMetaDefinitionOk() (*MoBaseMoRelationship, bool)`

GetMetaDefinitionOk returns a tuple with the MetaDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDefinition

`func (o *WorkflowUiDisplayMetadata) SetMetaDefinition(v MoBaseMoRelationship)`

SetMetaDefinition sets MetaDefinition field to given value.

### HasMetaDefinition

`func (o *WorkflowUiDisplayMetadata) HasMetaDefinition() bool`

HasMetaDefinition returns a boolean if a field has been set.

### SetMetaDefinitionNil

`func (o *WorkflowUiDisplayMetadata) SetMetaDefinitionNil(b bool)`

 SetMetaDefinitionNil sets the value for MetaDefinition to be an explicit nil

### UnsetMetaDefinition
`func (o *WorkflowUiDisplayMetadata) UnsetMetaDefinition()`

UnsetMetaDefinition ensures that no value is present for MetaDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


