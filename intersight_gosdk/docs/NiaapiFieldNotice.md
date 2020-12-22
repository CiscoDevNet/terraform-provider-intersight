# NiaapiFieldNotice

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

### NewNiaapiFieldNotice

`func NewNiaapiFieldNotice(classId string, objectType string, ) *NiaapiFieldNotice`

NewNiaapiFieldNotice instantiates a new NiaapiFieldNotice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiFieldNoticeWithDefaults

`func NewNiaapiFieldNoticeWithDefaults() *NiaapiFieldNotice`

NewNiaapiFieldNoticeWithDefaults instantiates a new NiaapiFieldNotice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiFieldNotice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiFieldNotice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiFieldNotice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiFieldNotice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiFieldNotice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiFieldNotice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBugid

`func (o *NiaapiFieldNotice) GetBugid() string`

GetBugid returns the Bugid field if non-nil, zero value otherwise.

### GetBugidOk

`func (o *NiaapiFieldNotice) GetBugidOk() (*string, bool)`

GetBugidOk returns a tuple with the Bugid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugid

`func (o *NiaapiFieldNotice) SetBugid(v string)`

SetBugid sets Bugid field to given value.

### HasBugid

`func (o *NiaapiFieldNotice) HasBugid() bool`

HasBugid returns a boolean if a field has been set.

### GetFieldNoticeDesc

`func (o *NiaapiFieldNotice) GetFieldNoticeDesc() string`

GetFieldNoticeDesc returns the FieldNoticeDesc field if non-nil, zero value otherwise.

### GetFieldNoticeDescOk

`func (o *NiaapiFieldNotice) GetFieldNoticeDescOk() (*string, bool)`

GetFieldNoticeDescOk returns a tuple with the FieldNoticeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNoticeDesc

`func (o *NiaapiFieldNotice) SetFieldNoticeDesc(v string)`

SetFieldNoticeDesc sets FieldNoticeDesc field to given value.

### HasFieldNoticeDesc

`func (o *NiaapiFieldNotice) HasFieldNoticeDesc() bool`

HasFieldNoticeDesc returns a boolean if a field has been set.

### GetFieldNoticeId

`func (o *NiaapiFieldNotice) GetFieldNoticeId() string`

GetFieldNoticeId returns the FieldNoticeId field if non-nil, zero value otherwise.

### GetFieldNoticeIdOk

`func (o *NiaapiFieldNotice) GetFieldNoticeIdOk() (*string, bool)`

GetFieldNoticeIdOk returns a tuple with the FieldNoticeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNoticeId

`func (o *NiaapiFieldNotice) SetFieldNoticeId(v string)`

SetFieldNoticeId sets FieldNoticeId field to given value.

### HasFieldNoticeId

`func (o *NiaapiFieldNotice) HasFieldNoticeId() bool`

HasFieldNoticeId returns a boolean if a field has been set.

### GetFieldNoticeUrl

`func (o *NiaapiFieldNotice) GetFieldNoticeUrl() string`

GetFieldNoticeUrl returns the FieldNoticeUrl field if non-nil, zero value otherwise.

### GetFieldNoticeUrlOk

`func (o *NiaapiFieldNotice) GetFieldNoticeUrlOk() (*string, bool)`

GetFieldNoticeUrlOk returns a tuple with the FieldNoticeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNoticeUrl

`func (o *NiaapiFieldNotice) SetFieldNoticeUrl(v string)`

SetFieldNoticeUrl sets FieldNoticeUrl field to given value.

### HasFieldNoticeUrl

`func (o *NiaapiFieldNotice) HasFieldNoticeUrl() bool`

HasFieldNoticeUrl returns a boolean if a field has been set.

### GetHeadline

`func (o *NiaapiFieldNotice) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *NiaapiFieldNotice) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *NiaapiFieldNotice) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *NiaapiFieldNotice) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetHwpid

`func (o *NiaapiFieldNotice) GetHwpid() string`

GetHwpid returns the Hwpid field if non-nil, zero value otherwise.

### GetHwpidOk

`func (o *NiaapiFieldNotice) GetHwpidOk() (*string, bool)`

GetHwpidOk returns a tuple with the Hwpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwpid

`func (o *NiaapiFieldNotice) SetHwpid(v string)`

SetHwpid sets Hwpid field to given value.

### HasHwpid

`func (o *NiaapiFieldNotice) HasHwpid() bool`

HasHwpid returns a boolean if a field has been set.

### GetRevisionInfo

`func (o *NiaapiFieldNotice) GetRevisionInfo() []NiaapiRevisionInfo`

GetRevisionInfo returns the RevisionInfo field if non-nil, zero value otherwise.

### GetRevisionInfoOk

`func (o *NiaapiFieldNotice) GetRevisionInfoOk() (*[]NiaapiRevisionInfo, bool)`

GetRevisionInfoOk returns a tuple with the RevisionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionInfo

`func (o *NiaapiFieldNotice) SetRevisionInfo(v []NiaapiRevisionInfo)`

SetRevisionInfo sets RevisionInfo field to given value.

### HasRevisionInfo

`func (o *NiaapiFieldNotice) HasRevisionInfo() bool`

HasRevisionInfo returns a boolean if a field has been set.

### SetRevisionInfoNil

`func (o *NiaapiFieldNotice) SetRevisionInfoNil(b bool)`

 SetRevisionInfoNil sets the value for RevisionInfo to be an explicit nil

### UnsetRevisionInfo
`func (o *NiaapiFieldNotice) UnsetRevisionInfo()`

UnsetRevisionInfo ensures that no value is present for RevisionInfo, not even an explicit nil
### GetSwRelease

`func (o *NiaapiFieldNotice) GetSwRelease() string`

GetSwRelease returns the SwRelease field if non-nil, zero value otherwise.

### GetSwReleaseOk

`func (o *NiaapiFieldNotice) GetSwReleaseOk() (*string, bool)`

GetSwReleaseOk returns a tuple with the SwRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwRelease

`func (o *NiaapiFieldNotice) SetSwRelease(v string)`

SetSwRelease sets SwRelease field to given value.

### HasSwRelease

`func (o *NiaapiFieldNotice) HasSwRelease() bool`

HasSwRelease returns a boolean if a field has been set.

### GetWorkaroundUrl

`func (o *NiaapiFieldNotice) GetWorkaroundUrl() string`

GetWorkaroundUrl returns the WorkaroundUrl field if non-nil, zero value otherwise.

### GetWorkaroundUrlOk

`func (o *NiaapiFieldNotice) GetWorkaroundUrlOk() (*string, bool)`

GetWorkaroundUrlOk returns a tuple with the WorkaroundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkaroundUrl

`func (o *NiaapiFieldNotice) SetWorkaroundUrl(v string)`

SetWorkaroundUrl sets WorkaroundUrl field to given value.

### HasWorkaroundUrl

`func (o *NiaapiFieldNotice) HasWorkaroundUrl() bool`

HasWorkaroundUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


