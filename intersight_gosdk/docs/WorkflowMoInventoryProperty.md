# WorkflowMoInventoryProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MoInventoryProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MoInventoryProperty"]
**Attributes** | Pointer to **[]string** |  | [optional] 
**ReferenceObjectType** | Pointer to **string** | ObjectType for which the attributes need to be collected. | [optional] 
**ReferenceType** | Pointer to **string** | Defines how the reference to the shadow resource is done. Base case is via an Moid, but reference via a selector is also possible. * &#x60;Moid&#x60; - The reference to the original resource is via an Moid. * &#x60;Selector&#x60; - The reference to the original resource is via a selector query. This can potentially lead to tracking data for multiple resources. | [optional] [default to "Moid"]

## Methods

### NewWorkflowMoInventoryProperty

`func NewWorkflowMoInventoryProperty(classId string, objectType string, ) *WorkflowMoInventoryProperty`

NewWorkflowMoInventoryProperty instantiates a new WorkflowMoInventoryProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoInventoryPropertyWithDefaults

`func NewWorkflowMoInventoryPropertyWithDefaults() *WorkflowMoInventoryProperty`

NewWorkflowMoInventoryPropertyWithDefaults instantiates a new WorkflowMoInventoryProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoInventoryProperty) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoInventoryProperty) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoInventoryProperty) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoInventoryProperty) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoInventoryProperty) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoInventoryProperty) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributes

`func (o *WorkflowMoInventoryProperty) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowMoInventoryProperty) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowMoInventoryProperty) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkflowMoInventoryProperty) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WorkflowMoInventoryProperty) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WorkflowMoInventoryProperty) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetReferenceObjectType

`func (o *WorkflowMoInventoryProperty) GetReferenceObjectType() string`

GetReferenceObjectType returns the ReferenceObjectType field if non-nil, zero value otherwise.

### GetReferenceObjectTypeOk

`func (o *WorkflowMoInventoryProperty) GetReferenceObjectTypeOk() (*string, bool)`

GetReferenceObjectTypeOk returns a tuple with the ReferenceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceObjectType

`func (o *WorkflowMoInventoryProperty) SetReferenceObjectType(v string)`

SetReferenceObjectType sets ReferenceObjectType field to given value.

### HasReferenceObjectType

`func (o *WorkflowMoInventoryProperty) HasReferenceObjectType() bool`

HasReferenceObjectType returns a boolean if a field has been set.

### GetReferenceType

`func (o *WorkflowMoInventoryProperty) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *WorkflowMoInventoryProperty) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *WorkflowMoInventoryProperty) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *WorkflowMoInventoryProperty) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


