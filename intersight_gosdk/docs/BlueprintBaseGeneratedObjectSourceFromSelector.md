# BlueprintBaseGeneratedObjectSourceFromSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**SourceSelectors** | Pointer to [**[]BlueprintGeneratedObjectSourceSelector**](BlueprintGeneratedObjectSourceSelector.md) |  | [optional] 

## Methods

### NewBlueprintBaseGeneratedObjectSourceFromSelector

`func NewBlueprintBaseGeneratedObjectSourceFromSelector(classId string, objectType string, ) *BlueprintBaseGeneratedObjectSourceFromSelector`

NewBlueprintBaseGeneratedObjectSourceFromSelector instantiates a new BlueprintBaseGeneratedObjectSourceFromSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintBaseGeneratedObjectSourceFromSelectorWithDefaults

`func NewBlueprintBaseGeneratedObjectSourceFromSelectorWithDefaults() *BlueprintBaseGeneratedObjectSourceFromSelector`

NewBlueprintBaseGeneratedObjectSourceFromSelectorWithDefaults instantiates a new BlueprintBaseGeneratedObjectSourceFromSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSourceSelectors

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) GetSourceSelectors() []BlueprintGeneratedObjectSourceSelector`

GetSourceSelectors returns the SourceSelectors field if non-nil, zero value otherwise.

### GetSourceSelectorsOk

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) GetSourceSelectorsOk() (*[]BlueprintGeneratedObjectSourceSelector, bool)`

GetSourceSelectorsOk returns a tuple with the SourceSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSelectors

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) SetSourceSelectors(v []BlueprintGeneratedObjectSourceSelector)`

SetSourceSelectors sets SourceSelectors field to given value.

### HasSourceSelectors

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) HasSourceSelectors() bool`

HasSourceSelectors returns a boolean if a field has been set.

### SetSourceSelectorsNil

`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) SetSourceSelectorsNil(b bool)`

 SetSourceSelectorsNil sets the value for SourceSelectors to be an explicit nil

### UnsetSourceSelectors
`func (o *BlueprintBaseGeneratedObjectSourceFromSelector) UnsetSourceSelectors()`

UnsetSourceSelectors ensures that no value is present for SourceSelectors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


