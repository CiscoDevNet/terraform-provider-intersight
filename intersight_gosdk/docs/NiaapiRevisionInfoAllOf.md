# NiaapiRevisionInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.RevisionInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.RevisionInfo"]
**DatePublished** | Pointer to **time.Time** | The date the revision is made. | [optional] 
**RevisionComment** | Pointer to **string** | The changes made in this revision. | [optional] 
**RevisionNo** | Pointer to **string** | The Revision No. of this revision. | [optional] 

## Methods

### NewNiaapiRevisionInfoAllOf

`func NewNiaapiRevisionInfoAllOf(classId string, objectType string, ) *NiaapiRevisionInfoAllOf`

NewNiaapiRevisionInfoAllOf instantiates a new NiaapiRevisionInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiRevisionInfoAllOfWithDefaults

`func NewNiaapiRevisionInfoAllOfWithDefaults() *NiaapiRevisionInfoAllOf`

NewNiaapiRevisionInfoAllOfWithDefaults instantiates a new NiaapiRevisionInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiRevisionInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiRevisionInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiRevisionInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiRevisionInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiRevisionInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiRevisionInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatePublished

`func (o *NiaapiRevisionInfoAllOf) GetDatePublished() time.Time`

GetDatePublished returns the DatePublished field if non-nil, zero value otherwise.

### GetDatePublishedOk

`func (o *NiaapiRevisionInfoAllOf) GetDatePublishedOk() (*time.Time, bool)`

GetDatePublishedOk returns a tuple with the DatePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublished

`func (o *NiaapiRevisionInfoAllOf) SetDatePublished(v time.Time)`

SetDatePublished sets DatePublished field to given value.

### HasDatePublished

`func (o *NiaapiRevisionInfoAllOf) HasDatePublished() bool`

HasDatePublished returns a boolean if a field has been set.

### GetRevisionComment

`func (o *NiaapiRevisionInfoAllOf) GetRevisionComment() string`

GetRevisionComment returns the RevisionComment field if non-nil, zero value otherwise.

### GetRevisionCommentOk

`func (o *NiaapiRevisionInfoAllOf) GetRevisionCommentOk() (*string, bool)`

GetRevisionCommentOk returns a tuple with the RevisionComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionComment

`func (o *NiaapiRevisionInfoAllOf) SetRevisionComment(v string)`

SetRevisionComment sets RevisionComment field to given value.

### HasRevisionComment

`func (o *NiaapiRevisionInfoAllOf) HasRevisionComment() bool`

HasRevisionComment returns a boolean if a field has been set.

### GetRevisionNo

`func (o *NiaapiRevisionInfoAllOf) GetRevisionNo() string`

GetRevisionNo returns the RevisionNo field if non-nil, zero value otherwise.

### GetRevisionNoOk

`func (o *NiaapiRevisionInfoAllOf) GetRevisionNoOk() (*string, bool)`

GetRevisionNoOk returns a tuple with the RevisionNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNo

`func (o *NiaapiRevisionInfoAllOf) SetRevisionNo(v string)`

SetRevisionNo sets RevisionNo field to given value.

### HasRevisionNo

`func (o *NiaapiRevisionInfoAllOf) HasRevisionNo() bool`

HasRevisionNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


