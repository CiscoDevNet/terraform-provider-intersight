# BlueprintGeneratedObjectSourceLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectSourceLookup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectSourceLookup"]
**LookupType** | Pointer to **string** | The type of the object being looked up. * &#x60;Cloned&#x60; - The generated object that is being looked up is a result of a deep clone from a previous step. * &#x60;PreCreated&#x60; - The generated object that is being looked up is an existing or pre-created object from the system catalog or from the user account. | [optional] [default to "Cloned"]

## Methods

### NewBlueprintGeneratedObjectSourceLookup

`func NewBlueprintGeneratedObjectSourceLookup(classId string, objectType string, ) *BlueprintGeneratedObjectSourceLookup`

NewBlueprintGeneratedObjectSourceLookup instantiates a new BlueprintGeneratedObjectSourceLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectSourceLookupWithDefaults

`func NewBlueprintGeneratedObjectSourceLookupWithDefaults() *BlueprintGeneratedObjectSourceLookup`

NewBlueprintGeneratedObjectSourceLookupWithDefaults instantiates a new BlueprintGeneratedObjectSourceLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectSourceLookup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectSourceLookup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectSourceLookup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectSourceLookup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectSourceLookup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectSourceLookup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLookupType

`func (o *BlueprintGeneratedObjectSourceLookup) GetLookupType() string`

GetLookupType returns the LookupType field if non-nil, zero value otherwise.

### GetLookupTypeOk

`func (o *BlueprintGeneratedObjectSourceLookup) GetLookupTypeOk() (*string, bool)`

GetLookupTypeOk returns a tuple with the LookupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupType

`func (o *BlueprintGeneratedObjectSourceLookup) SetLookupType(v string)`

SetLookupType sets LookupType field to given value.

### HasLookupType

`func (o *BlueprintGeneratedObjectSourceLookup) HasLookupType() bool`

HasLookupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


