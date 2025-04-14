# WorkflowMigrateServiceItemActionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MigrateServiceItemActionProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MigrateServiceItemActionProperty"]
**SourceVersions** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewWorkflowMigrateServiceItemActionProperty

`func NewWorkflowMigrateServiceItemActionProperty(classId string, objectType string, ) *WorkflowMigrateServiceItemActionProperty`

NewWorkflowMigrateServiceItemActionProperty instantiates a new WorkflowMigrateServiceItemActionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMigrateServiceItemActionPropertyWithDefaults

`func NewWorkflowMigrateServiceItemActionPropertyWithDefaults() *WorkflowMigrateServiceItemActionProperty`

NewWorkflowMigrateServiceItemActionPropertyWithDefaults instantiates a new WorkflowMigrateServiceItemActionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMigrateServiceItemActionProperty) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMigrateServiceItemActionProperty) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMigrateServiceItemActionProperty) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMigrateServiceItemActionProperty) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMigrateServiceItemActionProperty) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMigrateServiceItemActionProperty) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSourceVersions

`func (o *WorkflowMigrateServiceItemActionProperty) GetSourceVersions() []int64`

GetSourceVersions returns the SourceVersions field if non-nil, zero value otherwise.

### GetSourceVersionsOk

`func (o *WorkflowMigrateServiceItemActionProperty) GetSourceVersionsOk() (*[]int64, bool)`

GetSourceVersionsOk returns a tuple with the SourceVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersions

`func (o *WorkflowMigrateServiceItemActionProperty) SetSourceVersions(v []int64)`

SetSourceVersions sets SourceVersions field to given value.

### HasSourceVersions

`func (o *WorkflowMigrateServiceItemActionProperty) HasSourceVersions() bool`

HasSourceVersions returns a boolean if a field has been set.

### SetSourceVersionsNil

`func (o *WorkflowMigrateServiceItemActionProperty) SetSourceVersionsNil(b bool)`

 SetSourceVersionsNil sets the value for SourceVersions to be an explicit nil

### UnsetSourceVersions
`func (o *WorkflowMigrateServiceItemActionProperty) UnsetSourceVersions()`

UnsetSourceVersions ensures that no value is present for SourceVersions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


