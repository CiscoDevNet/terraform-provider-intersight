# TechsupportmanagementPlatformParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionType** | Pointer to **int32** | CollectionType specifies if basic or detailed techsupport needs to be collected. * &#x60;1&#x60; - Collect basic techsupport. * &#x60;2&#x60; - Collect detailed techsupport. | [optional] [default to 1]
**IncludeCore** | Pointer to **bool** | Controls whether to include core file in the techsupport bundle. | [optional] 

## Methods

### NewTechsupportmanagementPlatformParam

`func NewTechsupportmanagementPlatformParam() *TechsupportmanagementPlatformParam`

NewTechsupportmanagementPlatformParam instantiates a new TechsupportmanagementPlatformParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementPlatformParamWithDefaults

`func NewTechsupportmanagementPlatformParamWithDefaults() *TechsupportmanagementPlatformParam`

NewTechsupportmanagementPlatformParamWithDefaults instantiates a new TechsupportmanagementPlatformParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionType

`func (o *TechsupportmanagementPlatformParam) GetCollectionType() int32`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *TechsupportmanagementPlatformParam) GetCollectionTypeOk() (*int32, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *TechsupportmanagementPlatformParam) SetCollectionType(v int32)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *TechsupportmanagementPlatformParam) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### GetIncludeCore

`func (o *TechsupportmanagementPlatformParam) GetIncludeCore() bool`

GetIncludeCore returns the IncludeCore field if non-nil, zero value otherwise.

### GetIncludeCoreOk

`func (o *TechsupportmanagementPlatformParam) GetIncludeCoreOk() (*bool, bool)`

GetIncludeCoreOk returns a tuple with the IncludeCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCore

`func (o *TechsupportmanagementPlatformParam) SetIncludeCore(v bool)`

SetIncludeCore sets IncludeCore field to given value.

### HasIncludeCore

`func (o *TechsupportmanagementPlatformParam) HasIncludeCore() bool`

HasIncludeCore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


