# TechsupportmanagementPlatformParamAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.PlatformParam"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.PlatformParam"]
**CollectionType** | Pointer to **int32** | CollectionType specifies if basic or detailed techsupport needs to be collected. * &#x60;1&#x60; - Collect basic techsupport. * &#x60;2&#x60; - Collect detailed techsupport. | [optional] [default to 1]
**IncludeCore** | Pointer to **bool** | Controls whether to include core file in the techsupport bundle. | [optional] 

## Methods

### NewTechsupportmanagementPlatformParamAllOf

`func NewTechsupportmanagementPlatformParamAllOf(classId string, objectType string, ) *TechsupportmanagementPlatformParamAllOf`

NewTechsupportmanagementPlatformParamAllOf instantiates a new TechsupportmanagementPlatformParamAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementPlatformParamAllOfWithDefaults

`func NewTechsupportmanagementPlatformParamAllOfWithDefaults() *TechsupportmanagementPlatformParamAllOf`

NewTechsupportmanagementPlatformParamAllOfWithDefaults instantiates a new TechsupportmanagementPlatformParamAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementPlatformParamAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementPlatformParamAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementPlatformParamAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementPlatformParamAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementPlatformParamAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementPlatformParamAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectionType

`func (o *TechsupportmanagementPlatformParamAllOf) GetCollectionType() int32`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *TechsupportmanagementPlatformParamAllOf) GetCollectionTypeOk() (*int32, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *TechsupportmanagementPlatformParamAllOf) SetCollectionType(v int32)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *TechsupportmanagementPlatformParamAllOf) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### GetIncludeCore

`func (o *TechsupportmanagementPlatformParamAllOf) GetIncludeCore() bool`

GetIncludeCore returns the IncludeCore field if non-nil, zero value otherwise.

### GetIncludeCoreOk

`func (o *TechsupportmanagementPlatformParamAllOf) GetIncludeCoreOk() (*bool, bool)`

GetIncludeCoreOk returns a tuple with the IncludeCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCore

`func (o *TechsupportmanagementPlatformParamAllOf) SetIncludeCore(v bool)`

SetIncludeCore sets IncludeCore field to given value.

### HasIncludeCore

`func (o *TechsupportmanagementPlatformParamAllOf) HasIncludeCore() bool`

HasIncludeCore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


