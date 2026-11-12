# BlueprintBlueprintDefinitionDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.BlueprintDefinitionDependency"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.BlueprintDefinitionDependency"]
**BlueprintName** | Pointer to **string** | The qualified name of the required blueprint. | [optional] 
**DependencyQualifier** | Pointer to [**BlueprintOdataDependencyQualifier**](BlueprintOdataDependencyQualifier.md) |  | [optional] 
**Description** | Pointer to **string** | Specifies why the required blueprint is included. | [optional] 
**InputParameters** | Pointer to **interface{}** | Specifies input mappings for the required blueprint, including values derived from the referencing blueprint. | [optional] 
**Label** | Pointer to **string** | A user defined label identifier for this blueprint. | [optional] 

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

### GetDependencyQualifier

`func (o *BlueprintBlueprintDefinitionDependency) GetDependencyQualifier() BlueprintOdataDependencyQualifier`

GetDependencyQualifier returns the DependencyQualifier field if non-nil, zero value otherwise.

### GetDependencyQualifierOk

`func (o *BlueprintBlueprintDefinitionDependency) GetDependencyQualifierOk() (*BlueprintOdataDependencyQualifier, bool)`

GetDependencyQualifierOk returns a tuple with the DependencyQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyQualifier

`func (o *BlueprintBlueprintDefinitionDependency) SetDependencyQualifier(v BlueprintOdataDependencyQualifier)`

SetDependencyQualifier sets DependencyQualifier field to given value.

### HasDependencyQualifier

`func (o *BlueprintBlueprintDefinitionDependency) HasDependencyQualifier() bool`

HasDependencyQualifier returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


