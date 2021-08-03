# WorkflowCustomDataTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CustomDataTypeDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CustomDataTypeDefinition"]
**CompositeType** | Pointer to **bool** | When true this data type definition is a collection of type definitions to represent composite data like JSON. | [optional] [default to false]
**Description** | Pointer to **string** | A human-friendly description of this custom data type indicating it&#39;s domain and usage. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote (&#39;), or an underscore (_). | [optional] 
**Name** | Pointer to **string** | The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters &#39;-&#39; and &#39;_&#39;. | [optional] 
**ParameterSet** | Pointer to [**[]WorkflowParameterSet**](WorkflowParameterSet.md) |  | [optional] 
**Properties** | Pointer to [**NullableWorkflowCustomDataTypeProperties**](workflow.CustomDataTypeProperties.md) |  | [optional] 
**TypeDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](workflow.Catalog.Relationship.md) |  | [optional] 
**ClonedFrom** | Pointer to [**WorkflowCustomDataTypeDefinitionRelationship**](workflow.CustomDataTypeDefinition.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowCustomDataTypeDefinition

`func NewWorkflowCustomDataTypeDefinition(classId string, objectType string, ) *WorkflowCustomDataTypeDefinition`

NewWorkflowCustomDataTypeDefinition instantiates a new WorkflowCustomDataTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCustomDataTypeDefinitionWithDefaults

`func NewWorkflowCustomDataTypeDefinitionWithDefaults() *WorkflowCustomDataTypeDefinition`

NewWorkflowCustomDataTypeDefinitionWithDefaults instantiates a new WorkflowCustomDataTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCustomDataTypeDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCustomDataTypeDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCustomDataTypeDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCustomDataTypeDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCustomDataTypeDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCustomDataTypeDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompositeType

`func (o *WorkflowCustomDataTypeDefinition) GetCompositeType() bool`

GetCompositeType returns the CompositeType field if non-nil, zero value otherwise.

### GetCompositeTypeOk

`func (o *WorkflowCustomDataTypeDefinition) GetCompositeTypeOk() (*bool, bool)`

GetCompositeTypeOk returns a tuple with the CompositeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeType

`func (o *WorkflowCustomDataTypeDefinition) SetCompositeType(v bool)`

SetCompositeType sets CompositeType field to given value.

### HasCompositeType

`func (o *WorkflowCustomDataTypeDefinition) HasCompositeType() bool`

HasCompositeType returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowCustomDataTypeDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCustomDataTypeDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCustomDataTypeDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCustomDataTypeDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowCustomDataTypeDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowCustomDataTypeDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowCustomDataTypeDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowCustomDataTypeDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowCustomDataTypeDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowCustomDataTypeDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowCustomDataTypeDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowCustomDataTypeDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameterSet

`func (o *WorkflowCustomDataTypeDefinition) GetParameterSet() []WorkflowParameterSet`

GetParameterSet returns the ParameterSet field if non-nil, zero value otherwise.

### GetParameterSetOk

`func (o *WorkflowCustomDataTypeDefinition) GetParameterSetOk() (*[]WorkflowParameterSet, bool)`

GetParameterSetOk returns a tuple with the ParameterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterSet

`func (o *WorkflowCustomDataTypeDefinition) SetParameterSet(v []WorkflowParameterSet)`

SetParameterSet sets ParameterSet field to given value.

### HasParameterSet

`func (o *WorkflowCustomDataTypeDefinition) HasParameterSet() bool`

HasParameterSet returns a boolean if a field has been set.

### SetParameterSetNil

`func (o *WorkflowCustomDataTypeDefinition) SetParameterSetNil(b bool)`

 SetParameterSetNil sets the value for ParameterSet to be an explicit nil

### UnsetParameterSet
`func (o *WorkflowCustomDataTypeDefinition) UnsetParameterSet()`

UnsetParameterSet ensures that no value is present for ParameterSet, not even an explicit nil
### GetProperties

`func (o *WorkflowCustomDataTypeDefinition) GetProperties() WorkflowCustomDataTypeProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowCustomDataTypeDefinition) GetPropertiesOk() (*WorkflowCustomDataTypeProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowCustomDataTypeDefinition) SetProperties(v WorkflowCustomDataTypeProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowCustomDataTypeDefinition) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowCustomDataTypeDefinition) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowCustomDataTypeDefinition) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetTypeDefinition

`func (o *WorkflowCustomDataTypeDefinition) GetTypeDefinition() []WorkflowBaseDataType`

GetTypeDefinition returns the TypeDefinition field if non-nil, zero value otherwise.

### GetTypeDefinitionOk

`func (o *WorkflowCustomDataTypeDefinition) GetTypeDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetTypeDefinitionOk returns a tuple with the TypeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDefinition

`func (o *WorkflowCustomDataTypeDefinition) SetTypeDefinition(v []WorkflowBaseDataType)`

SetTypeDefinition sets TypeDefinition field to given value.

### HasTypeDefinition

`func (o *WorkflowCustomDataTypeDefinition) HasTypeDefinition() bool`

HasTypeDefinition returns a boolean if a field has been set.

### SetTypeDefinitionNil

`func (o *WorkflowCustomDataTypeDefinition) SetTypeDefinitionNil(b bool)`

 SetTypeDefinitionNil sets the value for TypeDefinition to be an explicit nil

### UnsetTypeDefinition
`func (o *WorkflowCustomDataTypeDefinition) UnsetTypeDefinition()`

UnsetTypeDefinition ensures that no value is present for TypeDefinition, not even an explicit nil
### GetCatalog

`func (o *WorkflowCustomDataTypeDefinition) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowCustomDataTypeDefinition) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowCustomDataTypeDefinition) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowCustomDataTypeDefinition) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetClonedFrom

`func (o *WorkflowCustomDataTypeDefinition) GetClonedFrom() WorkflowCustomDataTypeDefinitionRelationship`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *WorkflowCustomDataTypeDefinition) GetClonedFromOk() (*WorkflowCustomDataTypeDefinitionRelationship, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *WorkflowCustomDataTypeDefinition) SetClonedFrom(v WorkflowCustomDataTypeDefinitionRelationship)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *WorkflowCustomDataTypeDefinition) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


