# ResourceTagQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.TagQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.TagQualifier"]
**ChassisTags** | Pointer to [**[]ResourceTag**](ResourceTag.md) |  | [optional] 
**DomainProfileTags** | Pointer to [**[]ResourceTag**](ResourceTag.md) |  | [optional] 
**ServerTags** | Pointer to [**[]ResourceTag**](ResourceTag.md) |  | [optional] 

## Methods

### NewResourceTagQualifier

`func NewResourceTagQualifier(classId string, objectType string, ) *ResourceTagQualifier`

NewResourceTagQualifier instantiates a new ResourceTagQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTagQualifierWithDefaults

`func NewResourceTagQualifierWithDefaults() *ResourceTagQualifier`

NewResourceTagQualifierWithDefaults instantiates a new ResourceTagQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceTagQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceTagQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceTagQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceTagQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceTagQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceTagQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisTags

`func (o *ResourceTagQualifier) GetChassisTags() []ResourceTag`

GetChassisTags returns the ChassisTags field if non-nil, zero value otherwise.

### GetChassisTagsOk

`func (o *ResourceTagQualifier) GetChassisTagsOk() (*[]ResourceTag, bool)`

GetChassisTagsOk returns a tuple with the ChassisTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisTags

`func (o *ResourceTagQualifier) SetChassisTags(v []ResourceTag)`

SetChassisTags sets ChassisTags field to given value.

### HasChassisTags

`func (o *ResourceTagQualifier) HasChassisTags() bool`

HasChassisTags returns a boolean if a field has been set.

### SetChassisTagsNil

`func (o *ResourceTagQualifier) SetChassisTagsNil(b bool)`

 SetChassisTagsNil sets the value for ChassisTags to be an explicit nil

### UnsetChassisTags
`func (o *ResourceTagQualifier) UnsetChassisTags()`

UnsetChassisTags ensures that no value is present for ChassisTags, not even an explicit nil
### GetDomainProfileTags

`func (o *ResourceTagQualifier) GetDomainProfileTags() []ResourceTag`

GetDomainProfileTags returns the DomainProfileTags field if non-nil, zero value otherwise.

### GetDomainProfileTagsOk

`func (o *ResourceTagQualifier) GetDomainProfileTagsOk() (*[]ResourceTag, bool)`

GetDomainProfileTagsOk returns a tuple with the DomainProfileTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainProfileTags

`func (o *ResourceTagQualifier) SetDomainProfileTags(v []ResourceTag)`

SetDomainProfileTags sets DomainProfileTags field to given value.

### HasDomainProfileTags

`func (o *ResourceTagQualifier) HasDomainProfileTags() bool`

HasDomainProfileTags returns a boolean if a field has been set.

### SetDomainProfileTagsNil

`func (o *ResourceTagQualifier) SetDomainProfileTagsNil(b bool)`

 SetDomainProfileTagsNil sets the value for DomainProfileTags to be an explicit nil

### UnsetDomainProfileTags
`func (o *ResourceTagQualifier) UnsetDomainProfileTags()`

UnsetDomainProfileTags ensures that no value is present for DomainProfileTags, not even an explicit nil
### GetServerTags

`func (o *ResourceTagQualifier) GetServerTags() []ResourceTag`

GetServerTags returns the ServerTags field if non-nil, zero value otherwise.

### GetServerTagsOk

`func (o *ResourceTagQualifier) GetServerTagsOk() (*[]ResourceTag, bool)`

GetServerTagsOk returns a tuple with the ServerTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTags

`func (o *ResourceTagQualifier) SetServerTags(v []ResourceTag)`

SetServerTags sets ServerTags field to given value.

### HasServerTags

`func (o *ResourceTagQualifier) HasServerTags() bool`

HasServerTags returns a boolean if a field has been set.

### SetServerTagsNil

`func (o *ResourceTagQualifier) SetServerTagsNil(b bool)`

 SetServerTagsNil sets the value for ServerTags to be an explicit nil

### UnsetServerTags
`func (o *ResourceTagQualifier) UnsetServerTags()`

UnsetServerTags ensures that no value is present for ServerTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


