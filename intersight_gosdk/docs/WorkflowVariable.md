# WorkflowVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.Variable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.Variable"]
**CreateUser** | Pointer to **string** | The user identifier who created the environment variable. | [optional] [readonly] 
**Definition** | Pointer to [**NullableWorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**ModUser** | Pointer to **string** | The user identifier who last updated the environment variable. | [optional] [readonly] 
**Name** | Pointer to **string** | This defines the name of the variable and it is set by the system. The name used inside the datatype definition will be set as the name of the variable. | [optional] [readonly] 
**Value** | Pointer to **interface{}** | This defines the value of the variable and it will be validated against the datatype defined in the definition. | [optional] 
**Catalog** | Pointer to [**NullableWorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkflowVariable

`func NewWorkflowVariable(classId string, objectType string, ) *WorkflowVariable`

NewWorkflowVariable instantiates a new WorkflowVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowVariableWithDefaults

`func NewWorkflowVariableWithDefaults() *WorkflowVariable`

NewWorkflowVariableWithDefaults instantiates a new WorkflowVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowVariable) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowVariable) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowVariable) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowVariable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowVariable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowVariable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreateUser

`func (o *WorkflowVariable) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *WorkflowVariable) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *WorkflowVariable) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *WorkflowVariable) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefinition

`func (o *WorkflowVariable) GetDefinition() WorkflowBaseDataType`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *WorkflowVariable) GetDefinitionOk() (*WorkflowBaseDataType, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *WorkflowVariable) SetDefinition(v WorkflowBaseDataType)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *WorkflowVariable) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### SetDefinitionNil

`func (o *WorkflowVariable) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *WorkflowVariable) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetModUser

`func (o *WorkflowVariable) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *WorkflowVariable) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *WorkflowVariable) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *WorkflowVariable) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *WorkflowVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowVariable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowVariable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowVariable) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowVariable) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowVariable) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WorkflowVariable) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WorkflowVariable) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCatalog

`func (o *WorkflowVariable) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowVariable) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowVariable) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowVariable) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *WorkflowVariable) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *WorkflowVariable) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


