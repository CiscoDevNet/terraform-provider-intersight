# BlueprintGeneratedObjectSourceMerge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "blueprint.GeneratedObjectSourceMerge"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "blueprint.GeneratedObjectSourceMerge"]
**MergeAction** | Pointer to **string** | The action to be performed when merging the source object with the target object. When the action is set to \&quot;Merge\&quot;, all empty source properties are ignored. \&quot;Replace\&quot; action replaces all properties in the target object with the source object properties. * &#x60;Merge&#x60; - The null properties/relationships of the source MO will be ignored for the target MO. The non-null properties/relationships of the source will override the target MO properties/relationships. * &#x60;Replace&#x60; - Merge action as described in RFC 7386. The null properties/relationships of the source MO will be deleted on the target MO.The non-null properties/relationships of the source will override the target MO properties/relationships.When source object type is different from target, only the properties common to both source and target  will be affected.Other properties on the target will be ignored. | [optional] [default to "Merge"]

## Methods

### NewBlueprintGeneratedObjectSourceMerge

`func NewBlueprintGeneratedObjectSourceMerge(classId string, objectType string, ) *BlueprintGeneratedObjectSourceMerge`

NewBlueprintGeneratedObjectSourceMerge instantiates a new BlueprintGeneratedObjectSourceMerge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintGeneratedObjectSourceMergeWithDefaults

`func NewBlueprintGeneratedObjectSourceMergeWithDefaults() *BlueprintGeneratedObjectSourceMerge`

NewBlueprintGeneratedObjectSourceMergeWithDefaults instantiates a new BlueprintGeneratedObjectSourceMerge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintGeneratedObjectSourceMerge) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintGeneratedObjectSourceMerge) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintGeneratedObjectSourceMerge) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintGeneratedObjectSourceMerge) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintGeneratedObjectSourceMerge) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintGeneratedObjectSourceMerge) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMergeAction

`func (o *BlueprintGeneratedObjectSourceMerge) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *BlueprintGeneratedObjectSourceMerge) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *BlueprintGeneratedObjectSourceMerge) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *BlueprintGeneratedObjectSourceMerge) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


