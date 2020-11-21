# NiaapiNewReleaseDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.NewReleaseDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.NewReleaseDetail"]
**Description** | Pointer to **string** | Description of this new verison release post. | [optional] 
**Link** | Pointer to **string** | Link of downloading the release file. | [optional] 
**ReleaseNoteLink** | Pointer to **string** | The link used to provide a gateway for customer to review the release note. | [optional] 
**ReleaseNoteLinkTitle** | Pointer to **string** | The link title used to provide a gateway for customer to review the release note. | [optional] 
**SoftwareDownloadLink** | Pointer to **string** | The link used to provide a gateway for customer to download the release. | [optional] 
**SoftwareDownloadLinkTitle** | Pointer to **string** | The link title used to provide a gateway for customer to download the release. | [optional] 
**Title** | Pointer to **string** | Title of the new verison release post. | [optional] 
**Version** | Pointer to **string** | Version number Associate with this Post. | [optional] 

## Methods

### NewNiaapiNewReleaseDetail

`func NewNiaapiNewReleaseDetail(classId string, objectType string, ) *NiaapiNewReleaseDetail`

NewNiaapiNewReleaseDetail instantiates a new NiaapiNewReleaseDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiNewReleaseDetailWithDefaults

`func NewNiaapiNewReleaseDetailWithDefaults() *NiaapiNewReleaseDetail`

NewNiaapiNewReleaseDetailWithDefaults instantiates a new NiaapiNewReleaseDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiNewReleaseDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiNewReleaseDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiNewReleaseDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiNewReleaseDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiNewReleaseDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiNewReleaseDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NiaapiNewReleaseDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiaapiNewReleaseDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiaapiNewReleaseDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiaapiNewReleaseDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLink

`func (o *NiaapiNewReleaseDetail) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *NiaapiNewReleaseDetail) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *NiaapiNewReleaseDetail) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *NiaapiNewReleaseDetail) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetReleaseNoteLink

`func (o *NiaapiNewReleaseDetail) GetReleaseNoteLink() string`

GetReleaseNoteLink returns the ReleaseNoteLink field if non-nil, zero value otherwise.

### GetReleaseNoteLinkOk

`func (o *NiaapiNewReleaseDetail) GetReleaseNoteLinkOk() (*string, bool)`

GetReleaseNoteLinkOk returns a tuple with the ReleaseNoteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNoteLink

`func (o *NiaapiNewReleaseDetail) SetReleaseNoteLink(v string)`

SetReleaseNoteLink sets ReleaseNoteLink field to given value.

### HasReleaseNoteLink

`func (o *NiaapiNewReleaseDetail) HasReleaseNoteLink() bool`

HasReleaseNoteLink returns a boolean if a field has been set.

### GetReleaseNoteLinkTitle

`func (o *NiaapiNewReleaseDetail) GetReleaseNoteLinkTitle() string`

GetReleaseNoteLinkTitle returns the ReleaseNoteLinkTitle field if non-nil, zero value otherwise.

### GetReleaseNoteLinkTitleOk

`func (o *NiaapiNewReleaseDetail) GetReleaseNoteLinkTitleOk() (*string, bool)`

GetReleaseNoteLinkTitleOk returns a tuple with the ReleaseNoteLinkTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNoteLinkTitle

`func (o *NiaapiNewReleaseDetail) SetReleaseNoteLinkTitle(v string)`

SetReleaseNoteLinkTitle sets ReleaseNoteLinkTitle field to given value.

### HasReleaseNoteLinkTitle

`func (o *NiaapiNewReleaseDetail) HasReleaseNoteLinkTitle() bool`

HasReleaseNoteLinkTitle returns a boolean if a field has been set.

### GetSoftwareDownloadLink

`func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLink() string`

GetSoftwareDownloadLink returns the SoftwareDownloadLink field if non-nil, zero value otherwise.

### GetSoftwareDownloadLinkOk

`func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLinkOk() (*string, bool)`

GetSoftwareDownloadLinkOk returns a tuple with the SoftwareDownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDownloadLink

`func (o *NiaapiNewReleaseDetail) SetSoftwareDownloadLink(v string)`

SetSoftwareDownloadLink sets SoftwareDownloadLink field to given value.

### HasSoftwareDownloadLink

`func (o *NiaapiNewReleaseDetail) HasSoftwareDownloadLink() bool`

HasSoftwareDownloadLink returns a boolean if a field has been set.

### GetSoftwareDownloadLinkTitle

`func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLinkTitle() string`

GetSoftwareDownloadLinkTitle returns the SoftwareDownloadLinkTitle field if non-nil, zero value otherwise.

### GetSoftwareDownloadLinkTitleOk

`func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLinkTitleOk() (*string, bool)`

GetSoftwareDownloadLinkTitleOk returns a tuple with the SoftwareDownloadLinkTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDownloadLinkTitle

`func (o *NiaapiNewReleaseDetail) SetSoftwareDownloadLinkTitle(v string)`

SetSoftwareDownloadLinkTitle sets SoftwareDownloadLinkTitle field to given value.

### HasSoftwareDownloadLinkTitle

`func (o *NiaapiNewReleaseDetail) HasSoftwareDownloadLinkTitle() bool`

HasSoftwareDownloadLinkTitle returns a boolean if a field has been set.

### GetTitle

`func (o *NiaapiNewReleaseDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NiaapiNewReleaseDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NiaapiNewReleaseDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NiaapiNewReleaseDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *NiaapiNewReleaseDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiaapiNewReleaseDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiaapiNewReleaseDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiaapiNewReleaseDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


