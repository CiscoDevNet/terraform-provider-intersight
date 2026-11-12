# BlueprintVersionCompatibilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.VersionCompatibilityInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.VersionCompatibilityInfo"]
**Blueprints** | Pointer to [**[]BlueprintVersionReference**](BlueprintVersionReference.md) |  | [optional] 
**CompatibilityType** | Pointer to **string** | The update compatibility type for the listed blueprint versions relative to the current blueprint version. * &#x60;NotCompatible&#x60; - The previously published blueprint version cannot be updated to the current blueprint version. * &#x60;AutoUpdate&#x60; - The previously published blueprint version can be updated to the current blueprint version automatically. * &#x60;NeedsNewVersion&#x60; - The previously published blueprint version can move to the current blueprint version only by creating a new version of the dependent definition. | [optional] [default to "NotCompatible"]

## Methods

### NewBlueprintVersionCompatibilityInfo

`func NewBlueprintVersionCompatibilityInfo(classId string, objectType string, ) *BlueprintVersionCompatibilityInfo`

NewBlueprintVersionCompatibilityInfo instantiates a new BlueprintVersionCompatibilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintVersionCompatibilityInfoWithDefaults

`func NewBlueprintVersionCompatibilityInfoWithDefaults() *BlueprintVersionCompatibilityInfo`

NewBlueprintVersionCompatibilityInfoWithDefaults instantiates a new BlueprintVersionCompatibilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintVersionCompatibilityInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintVersionCompatibilityInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintVersionCompatibilityInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintVersionCompatibilityInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintVersionCompatibilityInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintVersionCompatibilityInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprints

`func (o *BlueprintVersionCompatibilityInfo) GetBlueprints() []BlueprintVersionReference`

GetBlueprints returns the Blueprints field if non-nil, zero value otherwise.

### GetBlueprintsOk

`func (o *BlueprintVersionCompatibilityInfo) GetBlueprintsOk() (*[]BlueprintVersionReference, bool)`

GetBlueprintsOk returns a tuple with the Blueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprints

`func (o *BlueprintVersionCompatibilityInfo) SetBlueprints(v []BlueprintVersionReference)`

SetBlueprints sets Blueprints field to given value.

### HasBlueprints

`func (o *BlueprintVersionCompatibilityInfo) HasBlueprints() bool`

HasBlueprints returns a boolean if a field has been set.

### SetBlueprintsNil

`func (o *BlueprintVersionCompatibilityInfo) SetBlueprintsNil(b bool)`

 SetBlueprintsNil sets the value for Blueprints to be an explicit nil

### UnsetBlueprints
`func (o *BlueprintVersionCompatibilityInfo) UnsetBlueprints()`

UnsetBlueprints ensures that no value is present for Blueprints, not even an explicit nil
### GetCompatibilityType

`func (o *BlueprintVersionCompatibilityInfo) GetCompatibilityType() string`

GetCompatibilityType returns the CompatibilityType field if non-nil, zero value otherwise.

### GetCompatibilityTypeOk

`func (o *BlueprintVersionCompatibilityInfo) GetCompatibilityTypeOk() (*string, bool)`

GetCompatibilityTypeOk returns a tuple with the CompatibilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityType

`func (o *BlueprintVersionCompatibilityInfo) SetCompatibilityType(v string)`

SetCompatibilityType sets CompatibilityType field to given value.

### HasCompatibilityType

`func (o *BlueprintVersionCompatibilityInfo) HasCompatibilityType() bool`

HasCompatibilityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


