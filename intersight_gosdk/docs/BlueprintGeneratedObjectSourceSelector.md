# BlueprintGeneratedObjectSourceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectSourceSelector"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectSourceSelector"]
**Selector** | Pointer to **string** | The selector to identify the source object. | [optional] 

## Methods

### NewBlueprintGeneratedObjectSourceSelector

`func NewBlueprintGeneratedObjectSourceSelector(classId string, objectType string, ) *BlueprintGeneratedObjectSourceSelector`

NewBlueprintGeneratedObjectSourceSelector instantiates a new BlueprintGeneratedObjectSourceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectSourceSelectorWithDefaults

`func NewBlueprintGeneratedObjectSourceSelectorWithDefaults() *BlueprintGeneratedObjectSourceSelector`

NewBlueprintGeneratedObjectSourceSelectorWithDefaults instantiates a new BlueprintGeneratedObjectSourceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectSourceSelector) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectSourceSelector) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectSourceSelector) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectSourceSelector) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectSourceSelector) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectSourceSelector) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSelector

`func (o *BlueprintGeneratedObjectSourceSelector) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *BlueprintGeneratedObjectSourceSelector) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *BlueprintGeneratedObjectSourceSelector) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *BlueprintGeneratedObjectSourceSelector) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


