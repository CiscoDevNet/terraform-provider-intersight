# BlueprintInputOperationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.InputOperationType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.InputOperationType"]
**ImpactType** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | The name of the input as defined within the definition. | [optional] 
**OperationModes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBlueprintInputOperationType

`func NewBlueprintInputOperationType(classId string, objectType string, ) *BlueprintInputOperationType`

NewBlueprintInputOperationType instantiates a new BlueprintInputOperationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintInputOperationTypeWithDefaults

`func NewBlueprintInputOperationTypeWithDefaults() *BlueprintInputOperationType`

NewBlueprintInputOperationTypeWithDefaults instantiates a new BlueprintInputOperationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintInputOperationType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintInputOperationType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintInputOperationType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintInputOperationType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintInputOperationType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintInputOperationType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImpactType

`func (o *BlueprintInputOperationType) GetImpactType() []string`

GetImpactType returns the ImpactType field if non-nil, zero value otherwise.

### GetImpactTypeOk

`func (o *BlueprintInputOperationType) GetImpactTypeOk() (*[]string, bool)`

GetImpactTypeOk returns a tuple with the ImpactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactType

`func (o *BlueprintInputOperationType) SetImpactType(v []string)`

SetImpactType sets ImpactType field to given value.

### HasImpactType

`func (o *BlueprintInputOperationType) HasImpactType() bool`

HasImpactType returns a boolean if a field has been set.

### SetImpactTypeNil

`func (o *BlueprintInputOperationType) SetImpactTypeNil(b bool)`

 SetImpactTypeNil sets the value for ImpactType to be an explicit nil

### UnsetImpactType
`func (o *BlueprintInputOperationType) UnsetImpactType()`

UnsetImpactType ensures that no value is present for ImpactType, not even an explicit nil
### GetName

`func (o *BlueprintInputOperationType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintInputOperationType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintInputOperationType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintInputOperationType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationModes

`func (o *BlueprintInputOperationType) GetOperationModes() []string`

GetOperationModes returns the OperationModes field if non-nil, zero value otherwise.

### GetOperationModesOk

`func (o *BlueprintInputOperationType) GetOperationModesOk() (*[]string, bool)`

GetOperationModesOk returns a tuple with the OperationModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationModes

`func (o *BlueprintInputOperationType) SetOperationModes(v []string)`

SetOperationModes sets OperationModes field to given value.

### HasOperationModes

`func (o *BlueprintInputOperationType) HasOperationModes() bool`

HasOperationModes returns a boolean if a field has been set.

### SetOperationModesNil

`func (o *BlueprintInputOperationType) SetOperationModesNil(b bool)`

 SetOperationModesNil sets the value for OperationModes to be an explicit nil

### UnsetOperationModes
`func (o *BlueprintInputOperationType) UnsetOperationModes()`

UnsetOperationModes ensures that no value is present for OperationModes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


