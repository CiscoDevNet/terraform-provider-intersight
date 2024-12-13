# ResourceServerQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AssetTags** | Pointer to **[]string** |  | [optional] 
**Pids** | Pointer to **[]string** |  | [optional] 
**UserLabels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResourceServerQualifier

`func NewResourceServerQualifier(classId string, objectType string, ) *ResourceServerQualifier`

NewResourceServerQualifier instantiates a new ResourceServerQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceServerQualifierWithDefaults

`func NewResourceServerQualifierWithDefaults() *ResourceServerQualifier`

NewResourceServerQualifierWithDefaults instantiates a new ResourceServerQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceServerQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceServerQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceServerQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceServerQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceServerQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceServerQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetTags

`func (o *ResourceServerQualifier) GetAssetTags() []string`

GetAssetTags returns the AssetTags field if non-nil, zero value otherwise.

### GetAssetTagsOk

`func (o *ResourceServerQualifier) GetAssetTagsOk() (*[]string, bool)`

GetAssetTagsOk returns a tuple with the AssetTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTags

`func (o *ResourceServerQualifier) SetAssetTags(v []string)`

SetAssetTags sets AssetTags field to given value.

### HasAssetTags

`func (o *ResourceServerQualifier) HasAssetTags() bool`

HasAssetTags returns a boolean if a field has been set.

### SetAssetTagsNil

`func (o *ResourceServerQualifier) SetAssetTagsNil(b bool)`

 SetAssetTagsNil sets the value for AssetTags to be an explicit nil

### UnsetAssetTags
`func (o *ResourceServerQualifier) UnsetAssetTags()`

UnsetAssetTags ensures that no value is present for AssetTags, not even an explicit nil
### GetPids

`func (o *ResourceServerQualifier) GetPids() []string`

GetPids returns the Pids field if non-nil, zero value otherwise.

### GetPidsOk

`func (o *ResourceServerQualifier) GetPidsOk() (*[]string, bool)`

GetPidsOk returns a tuple with the Pids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPids

`func (o *ResourceServerQualifier) SetPids(v []string)`

SetPids sets Pids field to given value.

### HasPids

`func (o *ResourceServerQualifier) HasPids() bool`

HasPids returns a boolean if a field has been set.

### SetPidsNil

`func (o *ResourceServerQualifier) SetPidsNil(b bool)`

 SetPidsNil sets the value for Pids to be an explicit nil

### UnsetPids
`func (o *ResourceServerQualifier) UnsetPids()`

UnsetPids ensures that no value is present for Pids, not even an explicit nil
### GetUserLabels

`func (o *ResourceServerQualifier) GetUserLabels() []string`

GetUserLabels returns the UserLabels field if non-nil, zero value otherwise.

### GetUserLabelsOk

`func (o *ResourceServerQualifier) GetUserLabelsOk() (*[]string, bool)`

GetUserLabelsOk returns a tuple with the UserLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabels

`func (o *ResourceServerQualifier) SetUserLabels(v []string)`

SetUserLabels sets UserLabels field to given value.

### HasUserLabels

`func (o *ResourceServerQualifier) HasUserLabels() bool`

HasUserLabels returns a boolean if a field has been set.

### SetUserLabelsNil

`func (o *ResourceServerQualifier) SetUserLabelsNil(b bool)`

 SetUserLabelsNil sets the value for UserLabels to be an explicit nil

### UnsetUserLabels
`func (o *ResourceServerQualifier) UnsetUserLabels()`

UnsetUserLabels ensures that no value is present for UserLabels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


