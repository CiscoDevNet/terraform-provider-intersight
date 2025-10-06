# BlueprintServiceItemDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.ServiceItemDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.ServiceItemDefinition"]
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this service item belongs. When catalog moid is not specified then the catalog of the service item is used first and if no service item can be found in that catalog, then the system catalog is used. | [optional] 
**Description** | Pointer to **string** | The description of that describes the purpose of including this service item within the blueprint. | [optional] 
**InputParameters** | Pointer to **interface{}** | Capture the mapping of blueprint inputs to service item inputs. The mapping for inputs can refer to blueprint inputs or to the generated objects which are part of this blueprint and all nested blueprints. The mapping will be referred using the convention of ${&lt;BlueprintName&gt;.input.&lt;InputName&gt;} or ${&lt;BlueprintName&gt;.generatedObject.&lt;GeneratedObjectName&gt;}. | [optional] 
**Label** | Pointer to **string** | A user defined label identifier for this Service item. | [optional] 
**ServiceItemDefinitionName** | Pointer to **string** | The qualified name of service item that should be executed. | [optional] 
**Version** | Pointer to **int64** | The service item definition version. If not specified, the default version of the service item is used. | [optional] 

## Methods

### NewBlueprintServiceItemDefinition

`func NewBlueprintServiceItemDefinition(classId string, objectType string, ) *BlueprintServiceItemDefinition`

NewBlueprintServiceItemDefinition instantiates a new BlueprintServiceItemDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintServiceItemDefinitionWithDefaults

`func NewBlueprintServiceItemDefinitionWithDefaults() *BlueprintServiceItemDefinition`

NewBlueprintServiceItemDefinitionWithDefaults instantiates a new BlueprintServiceItemDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintServiceItemDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintServiceItemDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintServiceItemDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintServiceItemDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintServiceItemDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintServiceItemDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *BlueprintServiceItemDefinition) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *BlueprintServiceItemDefinition) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *BlueprintServiceItemDefinition) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *BlueprintServiceItemDefinition) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetDescription

`func (o *BlueprintServiceItemDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintServiceItemDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintServiceItemDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintServiceItemDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *BlueprintServiceItemDefinition) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *BlueprintServiceItemDefinition) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *BlueprintServiceItemDefinition) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *BlueprintServiceItemDefinition) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *BlueprintServiceItemDefinition) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *BlueprintServiceItemDefinition) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetLabel

`func (o *BlueprintServiceItemDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BlueprintServiceItemDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BlueprintServiceItemDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BlueprintServiceItemDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetServiceItemDefinitionName

`func (o *BlueprintServiceItemDefinition) GetServiceItemDefinitionName() string`

GetServiceItemDefinitionName returns the ServiceItemDefinitionName field if non-nil, zero value otherwise.

### GetServiceItemDefinitionNameOk

`func (o *BlueprintServiceItemDefinition) GetServiceItemDefinitionNameOk() (*string, bool)`

GetServiceItemDefinitionNameOk returns a tuple with the ServiceItemDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinitionName

`func (o *BlueprintServiceItemDefinition) SetServiceItemDefinitionName(v string)`

SetServiceItemDefinitionName sets ServiceItemDefinitionName field to given value.

### HasServiceItemDefinitionName

`func (o *BlueprintServiceItemDefinition) HasServiceItemDefinitionName() bool`

HasServiceItemDefinitionName returns a boolean if a field has been set.

### GetVersion

`func (o *BlueprintServiceItemDefinition) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintServiceItemDefinition) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintServiceItemDefinition) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlueprintServiceItemDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


