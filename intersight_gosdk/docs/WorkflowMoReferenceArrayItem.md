# WorkflowMoReferenceArrayItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MoReferenceArrayItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MoReferenceArrayItem"]
**Properties** | Pointer to [**[]WorkflowMoReferenceProperty**](WorkflowMoReferenceProperty.md) |  | [optional] 

## Methods

### NewWorkflowMoReferenceArrayItem

`func NewWorkflowMoReferenceArrayItem(classId string, objectType string, ) *WorkflowMoReferenceArrayItem`

NewWorkflowMoReferenceArrayItem instantiates a new WorkflowMoReferenceArrayItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoReferenceArrayItemWithDefaults

`func NewWorkflowMoReferenceArrayItemWithDefaults() *WorkflowMoReferenceArrayItem`

NewWorkflowMoReferenceArrayItemWithDefaults instantiates a new WorkflowMoReferenceArrayItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoReferenceArrayItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoReferenceArrayItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoReferenceArrayItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoReferenceArrayItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoReferenceArrayItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoReferenceArrayItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProperties

`func (o *WorkflowMoReferenceArrayItem) GetProperties() []WorkflowMoReferenceProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowMoReferenceArrayItem) GetPropertiesOk() (*[]WorkflowMoReferenceProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowMoReferenceArrayItem) SetProperties(v []WorkflowMoReferenceProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowMoReferenceArrayItem) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowMoReferenceArrayItem) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowMoReferenceArrayItem) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


