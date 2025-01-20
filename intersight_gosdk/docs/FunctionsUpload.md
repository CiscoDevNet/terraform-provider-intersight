# FunctionsUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "functions.Upload"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "functions.Upload"]
**Action** | Pointer to **string** | Action against the Upload. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;CompleteUploading&#x60; - Mark the instance of a Upload as uploaded. | [optional] [default to "None"]
**CreateUser** | Pointer to **string** | The user identifier who created the Upload. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the Upload. | [optional] 
**FileName** | Pointer to **string** | The file name of the Upload. File name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**FileSize** | Pointer to **int64** | The size (in bytes) of the file. | [optional] 
**ModUser** | Pointer to **string** | The user identifier who last updated the Upload. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Upload. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**PartSize** | Pointer to **int64** | The chunk size (in bytes) for each part of the file to be uploaded. | [optional] [default to 8388608]
**State** | Pointer to **string** | Current representation of the state of Upload. * &#x60;Uploading&#x60; - File uploading is in progress. * &#x60;Uploaded&#x60; - File uploading is completed. * &#x60;Failed&#x60; - File uploading is failed. | [optional] [readonly] [default to "Uploading"]
**UploadUrls** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFunctionsUpload

`func NewFunctionsUpload(classId string, objectType string, ) *FunctionsUpload`

NewFunctionsUpload instantiates a new FunctionsUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsUploadWithDefaults

`func NewFunctionsUploadWithDefaults() *FunctionsUpload`

NewFunctionsUploadWithDefaults instantiates a new FunctionsUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FunctionsUpload) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FunctionsUpload) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FunctionsUpload) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FunctionsUpload) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FunctionsUpload) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FunctionsUpload) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *FunctionsUpload) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FunctionsUpload) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FunctionsUpload) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FunctionsUpload) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreateUser

`func (o *FunctionsUpload) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *FunctionsUpload) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *FunctionsUpload) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *FunctionsUpload) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDescription

`func (o *FunctionsUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionsUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionsUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionsUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileName

`func (o *FunctionsUpload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FunctionsUpload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FunctionsUpload) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FunctionsUpload) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *FunctionsUpload) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FunctionsUpload) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FunctionsUpload) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FunctionsUpload) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetModUser

`func (o *FunctionsUpload) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *FunctionsUpload) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *FunctionsUpload) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *FunctionsUpload) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *FunctionsUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsUpload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionsUpload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartSize

`func (o *FunctionsUpload) GetPartSize() int64`

GetPartSize returns the PartSize field if non-nil, zero value otherwise.

### GetPartSizeOk

`func (o *FunctionsUpload) GetPartSizeOk() (*int64, bool)`

GetPartSizeOk returns a tuple with the PartSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartSize

`func (o *FunctionsUpload) SetPartSize(v int64)`

SetPartSize sets PartSize field to given value.

### HasPartSize

`func (o *FunctionsUpload) HasPartSize() bool`

HasPartSize returns a boolean if a field has been set.

### GetState

`func (o *FunctionsUpload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FunctionsUpload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FunctionsUpload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FunctionsUpload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUploadUrls

`func (o *FunctionsUpload) GetUploadUrls() []string`

GetUploadUrls returns the UploadUrls field if non-nil, zero value otherwise.

### GetUploadUrlsOk

`func (o *FunctionsUpload) GetUploadUrlsOk() (*[]string, bool)`

GetUploadUrlsOk returns a tuple with the UploadUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrls

`func (o *FunctionsUpload) SetUploadUrls(v []string)`

SetUploadUrls sets UploadUrls field to given value.

### HasUploadUrls

`func (o *FunctionsUpload) HasUploadUrls() bool`

HasUploadUrls returns a boolean if a field has been set.

### SetUploadUrlsNil

`func (o *FunctionsUpload) SetUploadUrlsNil(b bool)`

 SetUploadUrlsNil sets the value for UploadUrls to be an explicit nil

### UnsetUploadUrls
`func (o *FunctionsUpload) UnsetUploadUrls()`

UnsetUploadUrls ensures that no value is present for UploadUrls, not even an explicit nil
### GetOrganization

`func (o *FunctionsUpload) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FunctionsUpload) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FunctionsUpload) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FunctionsUpload) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FunctionsUpload) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FunctionsUpload) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


