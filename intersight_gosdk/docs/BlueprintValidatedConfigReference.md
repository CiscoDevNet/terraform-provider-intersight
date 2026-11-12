# BlueprintValidatedConfigReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.ValidatedConfigReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.ValidatedConfigReference"]
**Name** | Pointer to **string** | The name for the object in the blueprint definition. | [optional] 
**Selector** | Pointer to **string** | The selector to identify the object in the shared catalog. | [optional] 
**ValidatedObject** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewBlueprintValidatedConfigReference

`func NewBlueprintValidatedConfigReference(classId string, objectType string, ) *BlueprintValidatedConfigReference`

NewBlueprintValidatedConfigReference instantiates a new BlueprintValidatedConfigReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintValidatedConfigReferenceWithDefaults

`func NewBlueprintValidatedConfigReferenceWithDefaults() *BlueprintValidatedConfigReference`

NewBlueprintValidatedConfigReferenceWithDefaults instantiates a new BlueprintValidatedConfigReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintValidatedConfigReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintValidatedConfigReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintValidatedConfigReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintValidatedConfigReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintValidatedConfigReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintValidatedConfigReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *BlueprintValidatedConfigReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintValidatedConfigReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintValidatedConfigReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintValidatedConfigReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelector

`func (o *BlueprintValidatedConfigReference) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *BlueprintValidatedConfigReference) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *BlueprintValidatedConfigReference) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *BlueprintValidatedConfigReference) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetValidatedObject

`func (o *BlueprintValidatedConfigReference) GetValidatedObject() MoMoRef`

GetValidatedObject returns the ValidatedObject field if non-nil, zero value otherwise.

### GetValidatedObjectOk

`func (o *BlueprintValidatedConfigReference) GetValidatedObjectOk() (*MoMoRef, bool)`

GetValidatedObjectOk returns a tuple with the ValidatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedObject

`func (o *BlueprintValidatedConfigReference) SetValidatedObject(v MoMoRef)`

SetValidatedObject sets ValidatedObject field to given value.

### HasValidatedObject

`func (o *BlueprintValidatedConfigReference) HasValidatedObject() bool`

HasValidatedObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


