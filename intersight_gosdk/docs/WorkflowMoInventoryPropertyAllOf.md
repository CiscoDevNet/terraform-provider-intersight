# WorkflowMoInventoryPropertyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MoInventoryProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MoInventoryProperty"]
**Attributes** | Pointer to **[]string** |  | [optional] 
**ReferenceObjectType** | Pointer to **string** | ObjectType for which the attributes need to be collected. | [optional] 
**ReferenceType** | Pointer to **string** | Defines how the reference to the shadow resource is done. Base case is via an Moid, but reference via a selector is also possible. * &#x60;Moid&#x60; - The reference to the original resource is via an Moid. * &#x60;Selector&#x60; - The reference to the original resource is via a selector query. This can potentially lead to tracking data for multiple resources. | [optional] [default to "Moid"]

## Methods

### NewWorkflowMoInventoryPropertyAllOf

`func NewWorkflowMoInventoryPropertyAllOf(classId string, objectType string, ) *WorkflowMoInventoryPropertyAllOf`

NewWorkflowMoInventoryPropertyAllOf instantiates a new WorkflowMoInventoryPropertyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoInventoryPropertyAllOfWithDefaults

`func NewWorkflowMoInventoryPropertyAllOfWithDefaults() *WorkflowMoInventoryPropertyAllOf`

NewWorkflowMoInventoryPropertyAllOfWithDefaults instantiates a new WorkflowMoInventoryPropertyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoInventoryPropertyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoInventoryPropertyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoInventoryPropertyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoInventoryPropertyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoInventoryPropertyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoInventoryPropertyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttributes

`func (o *WorkflowMoInventoryPropertyAllOf) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowMoInventoryPropertyAllOf) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowMoInventoryPropertyAllOf) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkflowMoInventoryPropertyAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WorkflowMoInventoryPropertyAllOf) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WorkflowMoInventoryPropertyAllOf) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetReferenceObjectType

`func (o *WorkflowMoInventoryPropertyAllOf) GetReferenceObjectType() string`

GetReferenceObjectType returns the ReferenceObjectType field if non-nil, zero value otherwise.

### GetReferenceObjectTypeOk

`func (o *WorkflowMoInventoryPropertyAllOf) GetReferenceObjectTypeOk() (*string, bool)`

GetReferenceObjectTypeOk returns a tuple with the ReferenceObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceObjectType

`func (o *WorkflowMoInventoryPropertyAllOf) SetReferenceObjectType(v string)`

SetReferenceObjectType sets ReferenceObjectType field to given value.

### HasReferenceObjectType

`func (o *WorkflowMoInventoryPropertyAllOf) HasReferenceObjectType() bool`

HasReferenceObjectType returns a boolean if a field has been set.

### GetReferenceType

`func (o *WorkflowMoInventoryPropertyAllOf) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *WorkflowMoInventoryPropertyAllOf) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *WorkflowMoInventoryPropertyAllOf) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *WorkflowMoInventoryPropertyAllOf) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


