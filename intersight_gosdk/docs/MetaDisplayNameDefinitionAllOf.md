# MetaDisplayNameDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | A specification for constructing the displayname from the MO&#39;s properties. | [optional] [readonly] 
**IncludeAncestor** | Pointer to **bool** | An indication of whether the displayname should be contructed &#39;recursively&#39; including the displayname of the first ancestor with a similarly named displayname. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the displayname used as a key in the DisplayName map which is returned as part of an MO for a Rest request. | [optional] [readonly] 

## Methods

### NewMetaDisplayNameDefinitionAllOf

`func NewMetaDisplayNameDefinitionAllOf() *MetaDisplayNameDefinitionAllOf`

NewMetaDisplayNameDefinitionAllOf instantiates a new MetaDisplayNameDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDisplayNameDefinitionAllOfWithDefaults

`func NewMetaDisplayNameDefinitionAllOfWithDefaults() *MetaDisplayNameDefinitionAllOf`

NewMetaDisplayNameDefinitionAllOfWithDefaults instantiates a new MetaDisplayNameDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *MetaDisplayNameDefinitionAllOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MetaDisplayNameDefinitionAllOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MetaDisplayNameDefinitionAllOf) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MetaDisplayNameDefinitionAllOf) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIncludeAncestor

`func (o *MetaDisplayNameDefinitionAllOf) GetIncludeAncestor() bool`

GetIncludeAncestor returns the IncludeAncestor field if non-nil, zero value otherwise.

### GetIncludeAncestorOk

`func (o *MetaDisplayNameDefinitionAllOf) GetIncludeAncestorOk() (*bool, bool)`

GetIncludeAncestorOk returns a tuple with the IncludeAncestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAncestor

`func (o *MetaDisplayNameDefinitionAllOf) SetIncludeAncestor(v bool)`

SetIncludeAncestor sets IncludeAncestor field to given value.

### HasIncludeAncestor

`func (o *MetaDisplayNameDefinitionAllOf) HasIncludeAncestor() bool`

HasIncludeAncestor returns a boolean if a field has been set.

### GetName

`func (o *MetaDisplayNameDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaDisplayNameDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaDisplayNameDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaDisplayNameDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


