# BlueprintGeneratedObjectMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectMetadata"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectMetadata"]
**Conditions** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name for the generated object. The name given for reference in later parts of the blueprint definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. | [optional] 
**ObjectIteration** | Pointer to **string** | A template for controlling loop behavior when generating multiple instances of the generated object. The resolved value of the template is expected to be an integer which is used to determine the number of iterations in the loop. Each iteration creates a new instance of the generated object. | [optional] 
**ObjectSource** | Pointer to [**NullableBlueprintBaseGeneratedObjectSourceMetadata**](BlueprintBaseGeneratedObjectSourceMetadata.md) |  | [optional] 
**PreGenerateOperations** | Pointer to [**[]BlueprintPreGenerateOperation**](BlueprintPreGenerateOperation.md) |  | [optional] 
**PropertyIteration** | Pointer to [**[]BlueprintPropertyIteration**](BlueprintPropertyIteration.md) |  | [optional] 
**PropertyParameters** | Pointer to **interface{}** | A list of key value pairs where key is the property path and value is the template to derive the value. | [optional] 
**Required** | Pointer to **bool** | The flag to indicate if this object must be generated for this blueprint. | [optional] [default to false]

## Methods

### NewBlueprintGeneratedObjectMetadata

`func NewBlueprintGeneratedObjectMetadata(classId string, objectType string, ) *BlueprintGeneratedObjectMetadata`

NewBlueprintGeneratedObjectMetadata instantiates a new BlueprintGeneratedObjectMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectMetadataWithDefaults

`func NewBlueprintGeneratedObjectMetadataWithDefaults() *BlueprintGeneratedObjectMetadata`

NewBlueprintGeneratedObjectMetadataWithDefaults instantiates a new BlueprintGeneratedObjectMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectMetadata) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectMetadata) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectMetadata) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectMetadata) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectMetadata) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectMetadata) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConditions

`func (o *BlueprintGeneratedObjectMetadata) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *BlueprintGeneratedObjectMetadata) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *BlueprintGeneratedObjectMetadata) SetConditions(v []string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *BlueprintGeneratedObjectMetadata) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *BlueprintGeneratedObjectMetadata) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *BlueprintGeneratedObjectMetadata) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetName

`func (o *BlueprintGeneratedObjectMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintGeneratedObjectMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintGeneratedObjectMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintGeneratedObjectMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectIteration

`func (o *BlueprintGeneratedObjectMetadata) GetObjectIteration() string`

GetObjectIteration returns the ObjectIteration field if non-nil, zero value otherwise.

### GetObjectIterationOk

`func (o *BlueprintGeneratedObjectMetadata) GetObjectIterationOk() (*string, bool)`

GetObjectIterationOk returns a tuple with the ObjectIteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIteration

`func (o *BlueprintGeneratedObjectMetadata) SetObjectIteration(v string)`

SetObjectIteration sets ObjectIteration field to given value.

### HasObjectIteration

`func (o *BlueprintGeneratedObjectMetadata) HasObjectIteration() bool`

HasObjectIteration returns a boolean if a field has been set.

### GetObjectSource

`func (o *BlueprintGeneratedObjectMetadata) GetObjectSource() BlueprintBaseGeneratedObjectSourceMetadata`

GetObjectSource returns the ObjectSource field if non-nil, zero value otherwise.

### GetObjectSourceOk

`func (o *BlueprintGeneratedObjectMetadata) GetObjectSourceOk() (*BlueprintBaseGeneratedObjectSourceMetadata, bool)`

GetObjectSourceOk returns a tuple with the ObjectSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSource

`func (o *BlueprintGeneratedObjectMetadata) SetObjectSource(v BlueprintBaseGeneratedObjectSourceMetadata)`

SetObjectSource sets ObjectSource field to given value.

### HasObjectSource

`func (o *BlueprintGeneratedObjectMetadata) HasObjectSource() bool`

HasObjectSource returns a boolean if a field has been set.

### SetObjectSourceNil

`func (o *BlueprintGeneratedObjectMetadata) SetObjectSourceNil(b bool)`

 SetObjectSourceNil sets the value for ObjectSource to be an explicit nil

### UnsetObjectSource
`func (o *BlueprintGeneratedObjectMetadata) UnsetObjectSource()`

UnsetObjectSource ensures that no value is present for ObjectSource, not even an explicit nil
### GetPreGenerateOperations

`func (o *BlueprintGeneratedObjectMetadata) GetPreGenerateOperations() []BlueprintPreGenerateOperation`

GetPreGenerateOperations returns the PreGenerateOperations field if non-nil, zero value otherwise.

### GetPreGenerateOperationsOk

`func (o *BlueprintGeneratedObjectMetadata) GetPreGenerateOperationsOk() (*[]BlueprintPreGenerateOperation, bool)`

GetPreGenerateOperationsOk returns a tuple with the PreGenerateOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreGenerateOperations

`func (o *BlueprintGeneratedObjectMetadata) SetPreGenerateOperations(v []BlueprintPreGenerateOperation)`

SetPreGenerateOperations sets PreGenerateOperations field to given value.

### HasPreGenerateOperations

`func (o *BlueprintGeneratedObjectMetadata) HasPreGenerateOperations() bool`

HasPreGenerateOperations returns a boolean if a field has been set.

### SetPreGenerateOperationsNil

`func (o *BlueprintGeneratedObjectMetadata) SetPreGenerateOperationsNil(b bool)`

 SetPreGenerateOperationsNil sets the value for PreGenerateOperations to be an explicit nil

### UnsetPreGenerateOperations
`func (o *BlueprintGeneratedObjectMetadata) UnsetPreGenerateOperations()`

UnsetPreGenerateOperations ensures that no value is present for PreGenerateOperations, not even an explicit nil
### GetPropertyIteration

`func (o *BlueprintGeneratedObjectMetadata) GetPropertyIteration() []BlueprintPropertyIteration`

GetPropertyIteration returns the PropertyIteration field if non-nil, zero value otherwise.

### GetPropertyIterationOk

`func (o *BlueprintGeneratedObjectMetadata) GetPropertyIterationOk() (*[]BlueprintPropertyIteration, bool)`

GetPropertyIterationOk returns a tuple with the PropertyIteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIteration

`func (o *BlueprintGeneratedObjectMetadata) SetPropertyIteration(v []BlueprintPropertyIteration)`

SetPropertyIteration sets PropertyIteration field to given value.

### HasPropertyIteration

`func (o *BlueprintGeneratedObjectMetadata) HasPropertyIteration() bool`

HasPropertyIteration returns a boolean if a field has been set.

### SetPropertyIterationNil

`func (o *BlueprintGeneratedObjectMetadata) SetPropertyIterationNil(b bool)`

 SetPropertyIterationNil sets the value for PropertyIteration to be an explicit nil

### UnsetPropertyIteration
`func (o *BlueprintGeneratedObjectMetadata) UnsetPropertyIteration()`

UnsetPropertyIteration ensures that no value is present for PropertyIteration, not even an explicit nil
### GetPropertyParameters

`func (o *BlueprintGeneratedObjectMetadata) GetPropertyParameters() interface{}`

GetPropertyParameters returns the PropertyParameters field if non-nil, zero value otherwise.

### GetPropertyParametersOk

`func (o *BlueprintGeneratedObjectMetadata) GetPropertyParametersOk() (*interface{}, bool)`

GetPropertyParametersOk returns a tuple with the PropertyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParameters

`func (o *BlueprintGeneratedObjectMetadata) SetPropertyParameters(v interface{})`

SetPropertyParameters sets PropertyParameters field to given value.

### HasPropertyParameters

`func (o *BlueprintGeneratedObjectMetadata) HasPropertyParameters() bool`

HasPropertyParameters returns a boolean if a field has been set.

### SetPropertyParametersNil

`func (o *BlueprintGeneratedObjectMetadata) SetPropertyParametersNil(b bool)`

 SetPropertyParametersNil sets the value for PropertyParameters to be an explicit nil

### UnsetPropertyParameters
`func (o *BlueprintGeneratedObjectMetadata) UnsetPropertyParameters()`

UnsetPropertyParameters ensures that no value is present for PropertyParameters, not even an explicit nil
### GetRequired

`func (o *BlueprintGeneratedObjectMetadata) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BlueprintGeneratedObjectMetadata) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BlueprintGeneratedObjectMetadata) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BlueprintGeneratedObjectMetadata) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


