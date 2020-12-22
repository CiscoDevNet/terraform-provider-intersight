# NiaapiFieldNoticeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Bugid** | Pointer to **string** | Bug Id associated with this notice. | [optional] 
**FieldNoticeDesc** | Pointer to **string** | Field notice Description. | [optional] 
**FieldNoticeId** | Pointer to **string** | Fieldnotice Id of this notice. | [optional] 
**FieldNoticeUrl** | Pointer to **string** | Field notice URL link to the notice webpage. | [optional] 
**Headline** | Pointer to **string** | The headline of this field notice. | [optional] 
**Hwpid** | Pointer to **string** | Hardware PID for affected models. | [optional] 
**RevisionInfo** | Pointer to [**[]NiaapiRevisionInfo**](NiaapiRevisionInfo.md) |  | [optional] 
**SwRelease** | Pointer to **string** | Software Release number for affected versions. | [optional] 
**WorkaroundUrl** | Pointer to **string** | URL of workaround of this notice. | [optional] 

## Methods

### NewNiaapiFieldNoticeAllOf

`func NewNiaapiFieldNoticeAllOf(classId string, objectType string, ) *NiaapiFieldNoticeAllOf`

NewNiaapiFieldNoticeAllOf instantiates a new NiaapiFieldNoticeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiFieldNoticeAllOfWithDefaults

`func NewNiaapiFieldNoticeAllOfWithDefaults() *NiaapiFieldNoticeAllOf`

NewNiaapiFieldNoticeAllOfWithDefaults instantiates a new NiaapiFieldNoticeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiFieldNoticeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiFieldNoticeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiFieldNoticeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiFieldNoticeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiFieldNoticeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiFieldNoticeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBugid

`func (o *NiaapiFieldNoticeAllOf) GetBugid() string`

GetBugid returns the Bugid field if non-nil, zero value otherwise.

### GetBugidOk

`func (o *NiaapiFieldNoticeAllOf) GetBugidOk() (*string, bool)`

GetBugidOk returns a tuple with the Bugid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugid

`func (o *NiaapiFieldNoticeAllOf) SetBugid(v string)`

SetBugid sets Bugid field to given value.

### HasBugid

`func (o *NiaapiFieldNoticeAllOf) HasBugid() bool`

HasBugid returns a boolean if a field has been set.

### GetFieldNoticeDesc

`func (o *NiaapiFieldNoticeAllOf) GetFieldNoticeDesc() string`

GetFieldNoticeDesc returns the FieldNoticeDesc field if non-nil, zero value otherwise.

### GetFieldNoticeDescOk

`func (o *NiaapiFieldNoticeAllOf) GetFieldNoticeDescOk() (*string, bool)`

GetFieldNoticeDescOk returns a tuple with the FieldNoticeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNoticeDesc

`func (o *NiaapiFieldNoticeAllOf) SetFieldNoticeDesc(v string)`

SetFieldNoticeDesc sets FieldNoticeDesc field to given value.

### HasFieldNoticeDesc

`func (o *NiaapiFieldNoticeAllOf) HasFieldNoticeDesc() bool`

HasFieldNoticeDesc returns a boolean if a field has been set.

### GetFieldNoticeId

`func (o *NiaapiFieldNoticeAllOf) GetFieldNoticeId() string`

GetFieldNoticeId returns the FieldNoticeId field if non-nil, zero value otherwise.

### GetFieldNoticeIdOk

`func (o *NiaapiFieldNoticeAllOf) GetFieldNoticeIdOk() (*string, bool)`

GetFieldNoticeIdOk returns a tuple with the FieldNoticeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNoticeId

`func (o *NiaapiFieldNoticeAllOf) SetFieldNoticeId(v string)`

SetFieldNoticeId sets FieldNoticeId field to given value.

### HasFieldNoticeId

`func (o *NiaapiFieldNoticeAllOf) HasFieldNoticeId() bool`

HasFieldNoticeId returns a boolean if a field has been set.

### GetFieldNoticeUrl

`func (o *NiaapiFieldNoticeAllOf) GetFieldNoticeUrl() string`

GetFieldNoticeUrl returns the FieldNoticeUrl field if non-nil, zero value otherwise.

### GetFieldNoticeUrlOk

`func (o *NiaapiFieldNoticeAllOf) GetFieldNoticeUrlOk() (*string, bool)`

GetFieldNoticeUrlOk returns a tuple with the FieldNoticeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNoticeUrl

`func (o *NiaapiFieldNoticeAllOf) SetFieldNoticeUrl(v string)`

SetFieldNoticeUrl sets FieldNoticeUrl field to given value.

### HasFieldNoticeUrl

`func (o *NiaapiFieldNoticeAllOf) HasFieldNoticeUrl() bool`

HasFieldNoticeUrl returns a boolean if a field has been set.

### GetHeadline

`func (o *NiaapiFieldNoticeAllOf) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *NiaapiFieldNoticeAllOf) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *NiaapiFieldNoticeAllOf) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *NiaapiFieldNoticeAllOf) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetHwpid

`func (o *NiaapiFieldNoticeAllOf) GetHwpid() string`

GetHwpid returns the Hwpid field if non-nil, zero value otherwise.

### GetHwpidOk

`func (o *NiaapiFieldNoticeAllOf) GetHwpidOk() (*string, bool)`

GetHwpidOk returns a tuple with the Hwpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwpid

`func (o *NiaapiFieldNoticeAllOf) SetHwpid(v string)`

SetHwpid sets Hwpid field to given value.

### HasHwpid

`func (o *NiaapiFieldNoticeAllOf) HasHwpid() bool`

HasHwpid returns a boolean if a field has been set.

### GetRevisionInfo

`func (o *NiaapiFieldNoticeAllOf) GetRevisionInfo() []NiaapiRevisionInfo`

GetRevisionInfo returns the RevisionInfo field if non-nil, zero value otherwise.

### GetRevisionInfoOk

`func (o *NiaapiFieldNoticeAllOf) GetRevisionInfoOk() (*[]NiaapiRevisionInfo, bool)`

GetRevisionInfoOk returns a tuple with the RevisionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionInfo

`func (o *NiaapiFieldNoticeAllOf) SetRevisionInfo(v []NiaapiRevisionInfo)`

SetRevisionInfo sets RevisionInfo field to given value.

### HasRevisionInfo

`func (o *NiaapiFieldNoticeAllOf) HasRevisionInfo() bool`

HasRevisionInfo returns a boolean if a field has been set.

### SetRevisionInfoNil

`func (o *NiaapiFieldNoticeAllOf) SetRevisionInfoNil(b bool)`

 SetRevisionInfoNil sets the value for RevisionInfo to be an explicit nil

### UnsetRevisionInfo
`func (o *NiaapiFieldNoticeAllOf) UnsetRevisionInfo()`

UnsetRevisionInfo ensures that no value is present for RevisionInfo, not even an explicit nil
### GetSwRelease

`func (o *NiaapiFieldNoticeAllOf) GetSwRelease() string`

GetSwRelease returns the SwRelease field if non-nil, zero value otherwise.

### GetSwReleaseOk

`func (o *NiaapiFieldNoticeAllOf) GetSwReleaseOk() (*string, bool)`

GetSwReleaseOk returns a tuple with the SwRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwRelease

`func (o *NiaapiFieldNoticeAllOf) SetSwRelease(v string)`

SetSwRelease sets SwRelease field to given value.

### HasSwRelease

`func (o *NiaapiFieldNoticeAllOf) HasSwRelease() bool`

HasSwRelease returns a boolean if a field has been set.

### GetWorkaroundUrl

`func (o *NiaapiFieldNoticeAllOf) GetWorkaroundUrl() string`

GetWorkaroundUrl returns the WorkaroundUrl field if non-nil, zero value otherwise.

### GetWorkaroundUrlOk

`func (o *NiaapiFieldNoticeAllOf) GetWorkaroundUrlOk() (*string, bool)`

GetWorkaroundUrlOk returns a tuple with the WorkaroundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkaroundUrl

`func (o *NiaapiFieldNoticeAllOf) SetWorkaroundUrl(v string)`

SetWorkaroundUrl sets WorkaroundUrl field to given value.

### HasWorkaroundUrl

`func (o *NiaapiFieldNoticeAllOf) HasWorkaroundUrl() bool`

HasWorkaroundUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


