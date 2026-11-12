# WorkflowDeprecationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.DeprecationInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.DeprecationInfo"]
**DeprecatedBy** | Pointer to **string** | The user or system email address that marked the versioned object as deprecated. The system populates this property only when the support status is Deprecated. | [optional] [readonly] 
**DeprecatedDependencies** | Pointer to [**[]WorkflowSupportObjectReference**](WorkflowSupportObjectReference.md) |  | [optional] 
**Description** | Pointer to **string** | The detailed description explaining why the versioned object is deprecated. Deprecated objects can use this field to provide additional support or migration information. | [optional] 
**ReplacementDefinition** | Pointer to [**NullableWorkflowSupportObjectReference**](WorkflowSupportObjectReference.md) |  | [optional] 
**SupportEndDate** | Pointer to **time.Time** | The date after which the deprecated object is no longer supported. | [optional] 

## Methods

### NewWorkflowDeprecationInfo

`func NewWorkflowDeprecationInfo(classId string, objectType string, ) *WorkflowDeprecationInfo`

NewWorkflowDeprecationInfo instantiates a new WorkflowDeprecationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDeprecationInfoWithDefaults

`func NewWorkflowDeprecationInfoWithDefaults() *WorkflowDeprecationInfo`

NewWorkflowDeprecationInfoWithDefaults instantiates a new WorkflowDeprecationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowDeprecationInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowDeprecationInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowDeprecationInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowDeprecationInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowDeprecationInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowDeprecationInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeprecatedBy

`func (o *WorkflowDeprecationInfo) GetDeprecatedBy() string`

GetDeprecatedBy returns the DeprecatedBy field if non-nil, zero value otherwise.

### GetDeprecatedByOk

`func (o *WorkflowDeprecationInfo) GetDeprecatedByOk() (*string, bool)`

GetDeprecatedByOk returns a tuple with the DeprecatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedBy

`func (o *WorkflowDeprecationInfo) SetDeprecatedBy(v string)`

SetDeprecatedBy sets DeprecatedBy field to given value.

### HasDeprecatedBy

`func (o *WorkflowDeprecationInfo) HasDeprecatedBy() bool`

HasDeprecatedBy returns a boolean if a field has been set.

### GetDeprecatedDependencies

`func (o *WorkflowDeprecationInfo) GetDeprecatedDependencies() []WorkflowSupportObjectReference`

GetDeprecatedDependencies returns the DeprecatedDependencies field if non-nil, zero value otherwise.

### GetDeprecatedDependenciesOk

`func (o *WorkflowDeprecationInfo) GetDeprecatedDependenciesOk() (*[]WorkflowSupportObjectReference, bool)`

GetDeprecatedDependenciesOk returns a tuple with the DeprecatedDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedDependencies

`func (o *WorkflowDeprecationInfo) SetDeprecatedDependencies(v []WorkflowSupportObjectReference)`

SetDeprecatedDependencies sets DeprecatedDependencies field to given value.

### HasDeprecatedDependencies

`func (o *WorkflowDeprecationInfo) HasDeprecatedDependencies() bool`

HasDeprecatedDependencies returns a boolean if a field has been set.

### SetDeprecatedDependenciesNil

`func (o *WorkflowDeprecationInfo) SetDeprecatedDependenciesNil(b bool)`

 SetDeprecatedDependenciesNil sets the value for DeprecatedDependencies to be an explicit nil

### UnsetDeprecatedDependencies
`func (o *WorkflowDeprecationInfo) UnsetDeprecatedDependencies()`

UnsetDeprecatedDependencies ensures that no value is present for DeprecatedDependencies, not even an explicit nil
### GetDescription

`func (o *WorkflowDeprecationInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowDeprecationInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowDeprecationInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowDeprecationInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReplacementDefinition

`func (o *WorkflowDeprecationInfo) GetReplacementDefinition() WorkflowSupportObjectReference`

GetReplacementDefinition returns the ReplacementDefinition field if non-nil, zero value otherwise.

### GetReplacementDefinitionOk

`func (o *WorkflowDeprecationInfo) GetReplacementDefinitionOk() (*WorkflowSupportObjectReference, bool)`

GetReplacementDefinitionOk returns a tuple with the ReplacementDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementDefinition

`func (o *WorkflowDeprecationInfo) SetReplacementDefinition(v WorkflowSupportObjectReference)`

SetReplacementDefinition sets ReplacementDefinition field to given value.

### HasReplacementDefinition

`func (o *WorkflowDeprecationInfo) HasReplacementDefinition() bool`

HasReplacementDefinition returns a boolean if a field has been set.

### SetReplacementDefinitionNil

`func (o *WorkflowDeprecationInfo) SetReplacementDefinitionNil(b bool)`

 SetReplacementDefinitionNil sets the value for ReplacementDefinition to be an explicit nil

### UnsetReplacementDefinition
`func (o *WorkflowDeprecationInfo) UnsetReplacementDefinition()`

UnsetReplacementDefinition ensures that no value is present for ReplacementDefinition, not even an explicit nil
### GetSupportEndDate

`func (o *WorkflowDeprecationInfo) GetSupportEndDate() time.Time`

GetSupportEndDate returns the SupportEndDate field if non-nil, zero value otherwise.

### GetSupportEndDateOk

`func (o *WorkflowDeprecationInfo) GetSupportEndDateOk() (*time.Time, bool)`

GetSupportEndDateOk returns a tuple with the SupportEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEndDate

`func (o *WorkflowDeprecationInfo) SetSupportEndDate(v time.Time)`

SetSupportEndDate sets SupportEndDate field to given value.

### HasSupportEndDate

`func (o *WorkflowDeprecationInfo) HasSupportEndDate() bool`

HasSupportEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


