# BlueprintBlueprintDefinitionDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.BlueprintDefinitionDependency"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.BlueprintDefinitionDependency"]
**BlueprintDefinition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**BlueprintName** | Pointer to **string** | The qualified name of blueprint. | [optional] 
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this blueprint belongs too. When catalog moid is not specified then the catalog of the parent blueprint is used first and if no blueprint can be found in that catalog, then the system catalog is used. | [optional] 
**Description** | Pointer to **string** | The description of that describes the purpose of including this dependent blueprint. | [optional] 
**InputParameters** | Pointer to **interface{}** | If any input of the nested blueprint needs to be modified in some form to include data from the parent blueprint, then the input parameters can be specified here. | [optional] 
**Label** | Pointer to **string** | A user defined label identifier for this blueprint. | [optional] 
**Version** | Pointer to **int64** | The blueprint version. If not specified, the default version is used. | [optional] 

## Methods

### NewBlueprintBlueprintDefinitionDependency

`func NewBlueprintBlueprintDefinitionDependency(classId string, objectType string, ) *BlueprintBlueprintDefinitionDependency`

NewBlueprintBlueprintDefinitionDependency instantiates a new BlueprintBlueprintDefinitionDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintBlueprintDefinitionDependencyWithDefaults

`func NewBlueprintBlueprintDefinitionDependencyWithDefaults() *BlueprintBlueprintDefinitionDependency`

NewBlueprintBlueprintDefinitionDependencyWithDefaults instantiates a new BlueprintBlueprintDefinitionDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintBlueprintDefinitionDependency) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintBlueprintDefinitionDependency) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintBlueprintDefinitionDependency) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintBlueprintDefinitionDependency) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintBlueprintDefinitionDependency) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintBlueprintDefinitionDependency) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprintDefinition

`func (o *BlueprintBlueprintDefinitionDependency) GetBlueprintDefinition() MoMoRef`

GetBlueprintDefinition returns the BlueprintDefinition field if non-nil, zero value otherwise.

### GetBlueprintDefinitionOk

`func (o *BlueprintBlueprintDefinitionDependency) GetBlueprintDefinitionOk() (*MoMoRef, bool)`

GetBlueprintDefinitionOk returns a tuple with the BlueprintDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintDefinition

`func (o *BlueprintBlueprintDefinitionDependency) SetBlueprintDefinition(v MoMoRef)`

SetBlueprintDefinition sets BlueprintDefinition field to given value.

### HasBlueprintDefinition

`func (o *BlueprintBlueprintDefinitionDependency) HasBlueprintDefinition() bool`

HasBlueprintDefinition returns a boolean if a field has been set.

### GetBlueprintName

`func (o *BlueprintBlueprintDefinitionDependency) GetBlueprintName() string`

GetBlueprintName returns the BlueprintName field if non-nil, zero value otherwise.

### GetBlueprintNameOk

`func (o *BlueprintBlueprintDefinitionDependency) GetBlueprintNameOk() (*string, bool)`

GetBlueprintNameOk returns a tuple with the BlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintName

`func (o *BlueprintBlueprintDefinitionDependency) SetBlueprintName(v string)`

SetBlueprintName sets BlueprintName field to given value.

### HasBlueprintName

`func (o *BlueprintBlueprintDefinitionDependency) HasBlueprintName() bool`

HasBlueprintName returns a boolean if a field has been set.

### GetCatalogMoid

`func (o *BlueprintBlueprintDefinitionDependency) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *BlueprintBlueprintDefinitionDependency) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *BlueprintBlueprintDefinitionDependency) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *BlueprintBlueprintDefinitionDependency) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetDescription

`func (o *BlueprintBlueprintDefinitionDependency) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintBlueprintDefinitionDependency) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintBlueprintDefinitionDependency) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintBlueprintDefinitionDependency) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *BlueprintBlueprintDefinitionDependency) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *BlueprintBlueprintDefinitionDependency) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *BlueprintBlueprintDefinitionDependency) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *BlueprintBlueprintDefinitionDependency) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *BlueprintBlueprintDefinitionDependency) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *BlueprintBlueprintDefinitionDependency) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetLabel

`func (o *BlueprintBlueprintDefinitionDependency) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BlueprintBlueprintDefinitionDependency) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BlueprintBlueprintDefinitionDependency) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BlueprintBlueprintDefinitionDependency) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetVersion

`func (o *BlueprintBlueprintDefinitionDependency) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintBlueprintDefinitionDependency) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintBlueprintDefinitionDependency) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlueprintBlueprintDefinitionDependency) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


