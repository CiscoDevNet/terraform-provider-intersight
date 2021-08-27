# NiaapiNewReleasePostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**PostDate** | Pointer to **time.Time** | The date when this new release notice is posted. | [optional] 
**PostDetail** | Pointer to [**NullableNiaapiNewReleaseDetail**](NiaapiNewReleaseDetail.md) |  | [optional] 
**PostType** | Pointer to **string** | The document type of this post. | [optional] 
**Postid** | Pointer to **string** | Identificator of this inbox post. | [optional] 
**Revision** | Pointer to **string** | Revision number of this notice. | [optional] 

## Methods

### NewNiaapiNewReleasePostAllOf

`func NewNiaapiNewReleasePostAllOf(classId string, objectType string, ) *NiaapiNewReleasePostAllOf`

NewNiaapiNewReleasePostAllOf instantiates a new NiaapiNewReleasePostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiNewReleasePostAllOfWithDefaults

`func NewNiaapiNewReleasePostAllOfWithDefaults() *NiaapiNewReleasePostAllOf`

NewNiaapiNewReleasePostAllOfWithDefaults instantiates a new NiaapiNewReleasePostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiNewReleasePostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiNewReleasePostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiNewReleasePostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiNewReleasePostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiNewReleasePostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiNewReleasePostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPostDate

`func (o *NiaapiNewReleasePostAllOf) GetPostDate() time.Time`

GetPostDate returns the PostDate field if non-nil, zero value otherwise.

### GetPostDateOk

`func (o *NiaapiNewReleasePostAllOf) GetPostDateOk() (*time.Time, bool)`

GetPostDateOk returns a tuple with the PostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDate

`func (o *NiaapiNewReleasePostAllOf) SetPostDate(v time.Time)`

SetPostDate sets PostDate field to given value.

### HasPostDate

`func (o *NiaapiNewReleasePostAllOf) HasPostDate() bool`

HasPostDate returns a boolean if a field has been set.

### GetPostDetail

`func (o *NiaapiNewReleasePostAllOf) GetPostDetail() NiaapiNewReleaseDetail`

GetPostDetail returns the PostDetail field if non-nil, zero value otherwise.

### GetPostDetailOk

`func (o *NiaapiNewReleasePostAllOf) GetPostDetailOk() (*NiaapiNewReleaseDetail, bool)`

GetPostDetailOk returns a tuple with the PostDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDetail

`func (o *NiaapiNewReleasePostAllOf) SetPostDetail(v NiaapiNewReleaseDetail)`

SetPostDetail sets PostDetail field to given value.

### HasPostDetail

`func (o *NiaapiNewReleasePostAllOf) HasPostDetail() bool`

HasPostDetail returns a boolean if a field has been set.

### SetPostDetailNil

`func (o *NiaapiNewReleasePostAllOf) SetPostDetailNil(b bool)`

 SetPostDetailNil sets the value for PostDetail to be an explicit nil

### UnsetPostDetail
`func (o *NiaapiNewReleasePostAllOf) UnsetPostDetail()`

UnsetPostDetail ensures that no value is present for PostDetail, not even an explicit nil
### GetPostType

`func (o *NiaapiNewReleasePostAllOf) GetPostType() string`

GetPostType returns the PostType field if non-nil, zero value otherwise.

### GetPostTypeOk

`func (o *NiaapiNewReleasePostAllOf) GetPostTypeOk() (*string, bool)`

GetPostTypeOk returns a tuple with the PostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostType

`func (o *NiaapiNewReleasePostAllOf) SetPostType(v string)`

SetPostType sets PostType field to given value.

### HasPostType

`func (o *NiaapiNewReleasePostAllOf) HasPostType() bool`

HasPostType returns a boolean if a field has been set.

### GetPostid

`func (o *NiaapiNewReleasePostAllOf) GetPostid() string`

GetPostid returns the Postid field if non-nil, zero value otherwise.

### GetPostidOk

`func (o *NiaapiNewReleasePostAllOf) GetPostidOk() (*string, bool)`

GetPostidOk returns a tuple with the Postid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostid

`func (o *NiaapiNewReleasePostAllOf) SetPostid(v string)`

SetPostid sets Postid field to given value.

### HasPostid

`func (o *NiaapiNewReleasePostAllOf) HasPostid() bool`

HasPostid returns a boolean if a field has been set.

### GetRevision

`func (o *NiaapiNewReleasePostAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NiaapiNewReleasePostAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NiaapiNewReleasePostAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *NiaapiNewReleasePostAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


