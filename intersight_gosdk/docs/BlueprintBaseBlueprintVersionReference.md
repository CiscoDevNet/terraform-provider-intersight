# BlueprintBaseBlueprintVersionReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**BlueprintDefinition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**CatalogMoid** | Pointer to **string** | The catalog MOID that this blueprint belongs to. When the catalog MOID is not specified, the catalog of the parent blueprint is used first. If no blueprint can be found in that catalog, the system catalog is used. | [optional] 
**Version** | Pointer to **int64** | The version number of the blueprint. | [optional] 

## Methods

### NewBlueprintBaseBlueprintVersionReference

`func NewBlueprintBaseBlueprintVersionReference(classId string, objectType string, ) *BlueprintBaseBlueprintVersionReference`

NewBlueprintBaseBlueprintVersionReference instantiates a new BlueprintBaseBlueprintVersionReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintBaseBlueprintVersionReferenceWithDefaults

`func NewBlueprintBaseBlueprintVersionReferenceWithDefaults() *BlueprintBaseBlueprintVersionReference`

NewBlueprintBaseBlueprintVersionReferenceWithDefaults instantiates a new BlueprintBaseBlueprintVersionReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintBaseBlueprintVersionReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintBaseBlueprintVersionReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintBaseBlueprintVersionReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintBaseBlueprintVersionReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintBaseBlueprintVersionReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintBaseBlueprintVersionReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprintDefinition

`func (o *BlueprintBaseBlueprintVersionReference) GetBlueprintDefinition() MoMoRef`

GetBlueprintDefinition returns the BlueprintDefinition field if non-nil, zero value otherwise.

### GetBlueprintDefinitionOk

`func (o *BlueprintBaseBlueprintVersionReference) GetBlueprintDefinitionOk() (*MoMoRef, bool)`

GetBlueprintDefinitionOk returns a tuple with the BlueprintDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintDefinition

`func (o *BlueprintBaseBlueprintVersionReference) SetBlueprintDefinition(v MoMoRef)`

SetBlueprintDefinition sets BlueprintDefinition field to given value.

### HasBlueprintDefinition

`func (o *BlueprintBaseBlueprintVersionReference) HasBlueprintDefinition() bool`

HasBlueprintDefinition returns a boolean if a field has been set.

### GetCatalogMoid

`func (o *BlueprintBaseBlueprintVersionReference) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *BlueprintBaseBlueprintVersionReference) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *BlueprintBaseBlueprintVersionReference) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *BlueprintBaseBlueprintVersionReference) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetVersion

`func (o *BlueprintBaseBlueprintVersionReference) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintBaseBlueprintVersionReference) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintBaseBlueprintVersionReference) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlueprintBaseBlueprintVersionReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


