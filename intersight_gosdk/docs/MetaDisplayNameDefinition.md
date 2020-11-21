# MetaDisplayNameDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meta.DisplayNameDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meta.DisplayNameDefinition"]
**Format** | Pointer to **string** | A specification for constructing the displayname from the MO&#39;s properties. | [optional] [readonly] 
**IncludeAncestor** | Pointer to **bool** | An indication of whether the displayname should be contructed &#39;recursively&#39; including the displayname of the first ancestor with a similarly named displayname. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the displayname used as a key in the DisplayName map which is returned as part of an MO for a Rest request. | [optional] [readonly] 

## Methods

### NewMetaDisplayNameDefinition

`func NewMetaDisplayNameDefinition(classId string, objectType string, ) *MetaDisplayNameDefinition`

NewMetaDisplayNameDefinition instantiates a new MetaDisplayNameDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDisplayNameDefinitionWithDefaults

`func NewMetaDisplayNameDefinitionWithDefaults() *MetaDisplayNameDefinition`

NewMetaDisplayNameDefinitionWithDefaults instantiates a new MetaDisplayNameDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetaDisplayNameDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetaDisplayNameDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetaDisplayNameDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetaDisplayNameDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetaDisplayNameDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetaDisplayNameDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFormat

`func (o *MetaDisplayNameDefinition) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MetaDisplayNameDefinition) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MetaDisplayNameDefinition) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MetaDisplayNameDefinition) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIncludeAncestor

`func (o *MetaDisplayNameDefinition) GetIncludeAncestor() bool`

GetIncludeAncestor returns the IncludeAncestor field if non-nil, zero value otherwise.

### GetIncludeAncestorOk

`func (o *MetaDisplayNameDefinition) GetIncludeAncestorOk() (*bool, bool)`

GetIncludeAncestorOk returns a tuple with the IncludeAncestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAncestor

`func (o *MetaDisplayNameDefinition) SetIncludeAncestor(v bool)`

SetIncludeAncestor sets IncludeAncestor field to given value.

### HasIncludeAncestor

`func (o *MetaDisplayNameDefinition) HasIncludeAncestor() bool`

HasIncludeAncestor returns a boolean if a field has been set.

### GetName

`func (o *MetaDisplayNameDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaDisplayNameDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaDisplayNameDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaDisplayNameDefinition) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


