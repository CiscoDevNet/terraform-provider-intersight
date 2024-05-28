# WorkflowServiceItemAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemAttribute"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemAttribute"]
**Attribute** | Pointer to **interface{}** | Service item attribute for a service item instance and the format is specified by attribute definition of the service item definition. | [optional] 
**Datatype** | Pointer to **string** | Datatype, if any, backing the service item attribute definition. | [optional] [readonly] 
**Name** | Pointer to **string** | Attribute name which is used in the attribute definition of the service item. | [optional] 
**Type** | Pointer to **string** | Type of the service item attribute. * &#x60;None&#x60; - Default value if the service item attribute does not belong to any of the existing types. * &#x60;Configuration&#x60; - The service item attribute is a configuration from the designer or the end user. * &#x60;Inventory&#x60; - The service item attribute captures the inventory of the resource created by the service item deployment. * &#x60;Health&#x60; - The service item attribute describes the health of the resource created by the service item deployment. * &#x60;Output&#x60; - The service item attribute captures the artifact generated after performing an action on the service item. | [optional] [readonly] [default to "None"]
**ServiceItemDefinition** | Pointer to [**NullableWorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 
**ServiceItemInstance** | Pointer to [**NullableWorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemAttribute

`func NewWorkflowServiceItemAttribute(classId string, objectType string, ) *WorkflowServiceItemAttribute`

NewWorkflowServiceItemAttribute instantiates a new WorkflowServiceItemAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemAttributeWithDefaults

`func NewWorkflowServiceItemAttributeWithDefaults() *WorkflowServiceItemAttribute`

NewWorkflowServiceItemAttributeWithDefaults instantiates a new WorkflowServiceItemAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemAttribute) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemAttribute) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemAttribute) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemAttribute) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemAttribute) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemAttribute) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttribute

`func (o *WorkflowServiceItemAttribute) GetAttribute() interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *WorkflowServiceItemAttribute) GetAttributeOk() (*interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *WorkflowServiceItemAttribute) SetAttribute(v interface{})`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *WorkflowServiceItemAttribute) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### SetAttributeNil

`func (o *WorkflowServiceItemAttribute) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *WorkflowServiceItemAttribute) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetDatatype

`func (o *WorkflowServiceItemAttribute) GetDatatype() string`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *WorkflowServiceItemAttribute) GetDatatypeOk() (*string, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *WorkflowServiceItemAttribute) SetDatatype(v string)`

SetDatatype sets Datatype field to given value.

### HasDatatype

`func (o *WorkflowServiceItemAttribute) HasDatatype() bool`

HasDatatype returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WorkflowServiceItemAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowServiceItemAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowServiceItemAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowServiceItemAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServiceItemDefinition

`func (o *WorkflowServiceItemAttribute) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemAttribute) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemAttribute) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemAttribute) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.

### SetServiceItemDefinitionNil

`func (o *WorkflowServiceItemAttribute) SetServiceItemDefinitionNil(b bool)`

 SetServiceItemDefinitionNil sets the value for ServiceItemDefinition to be an explicit nil

### UnsetServiceItemDefinition
`func (o *WorkflowServiceItemAttribute) UnsetServiceItemDefinition()`

UnsetServiceItemDefinition ensures that no value is present for ServiceItemDefinition, not even an explicit nil
### GetServiceItemInstance

`func (o *WorkflowServiceItemAttribute) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkflowServiceItemAttribute) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkflowServiceItemAttribute) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkflowServiceItemAttribute) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.

### SetServiceItemInstanceNil

`func (o *WorkflowServiceItemAttribute) SetServiceItemInstanceNil(b bool)`

 SetServiceItemInstanceNil sets the value for ServiceItemInstance to be an explicit nil

### UnsetServiceItemInstance
`func (o *WorkflowServiceItemAttribute) UnsetServiceItemInstance()`

UnsetServiceItemInstance ensures that no value is present for ServiceItemInstance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


