# ApplianceFileGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.FileGateway"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.FileGateway"]
**BucketName** | Pointer to **string** | Bucket name in the cloud storage service where the file is stored. | [optional] [readonly] 
**ExternalHost** | Pointer to **bool** | Flag to specify if the requested file is served from an external host. An external host is a host other than the Intersight endpoint URL that resides outside of an endpoint device&#39;s local network. A download client (e.g. the Intersight Appliance device connector) should use this property to determine if a proxy is required to reach the host. | [optional] [readonly] 
**FileSize** | Pointer to **int64** | Size of the file in bytes. FileSize maybe zero if the storage service does not report file size. | [optional] [readonly] 
**FileTime** | Pointer to **time.Time** | File timestamp as reported by the cloud storage service. | [optional] [readonly] 
**FileType** | Pointer to **string** | User defined file type supplied by the caller. | [optional] [readonly] 
**Filename** | Pointer to **string** | Full name of the file as it exists in the cloud storage service. If the file is in a subfolder, then the filename must contain the full path. | [optional] [readonly] 
**PresignedUrl** | Pointer to **string** | Pre-signed URL obtained from the cloud storage service. | [optional] [readonly] 
**ServerCert** | Pointer to **string** | SSL certificate of the cloud storage service. | [optional] [readonly] 
**ValidityPeriod** | Pointer to **int64** | Signed URL&#39;s validity period in minutes. | [optional] [readonly] 
**Version** | Pointer to **string** | File version as reported by the cloud storage service. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceFileGateway

`func NewApplianceFileGateway(classId string, objectType string, ) *ApplianceFileGateway`

NewApplianceFileGateway instantiates a new ApplianceFileGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceFileGatewayWithDefaults

`func NewApplianceFileGatewayWithDefaults() *ApplianceFileGateway`

NewApplianceFileGatewayWithDefaults instantiates a new ApplianceFileGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceFileGateway) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceFileGateway) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceFileGateway) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceFileGateway) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceFileGateway) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceFileGateway) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBucketName

`func (o *ApplianceFileGateway) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ApplianceFileGateway) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ApplianceFileGateway) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ApplianceFileGateway) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetExternalHost

`func (o *ApplianceFileGateway) GetExternalHost() bool`

GetExternalHost returns the ExternalHost field if non-nil, zero value otherwise.

### GetExternalHostOk

`func (o *ApplianceFileGateway) GetExternalHostOk() (*bool, bool)`

GetExternalHostOk returns a tuple with the ExternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHost

`func (o *ApplianceFileGateway) SetExternalHost(v bool)`

SetExternalHost sets ExternalHost field to given value.

### HasExternalHost

`func (o *ApplianceFileGateway) HasExternalHost() bool`

HasExternalHost returns a boolean if a field has been set.

### GetFileSize

`func (o *ApplianceFileGateway) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ApplianceFileGateway) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ApplianceFileGateway) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *ApplianceFileGateway) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileTime

`func (o *ApplianceFileGateway) GetFileTime() time.Time`

GetFileTime returns the FileTime field if non-nil, zero value otherwise.

### GetFileTimeOk

`func (o *ApplianceFileGateway) GetFileTimeOk() (*time.Time, bool)`

GetFileTimeOk returns a tuple with the FileTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTime

`func (o *ApplianceFileGateway) SetFileTime(v time.Time)`

SetFileTime sets FileTime field to given value.

### HasFileTime

`func (o *ApplianceFileGateway) HasFileTime() bool`

HasFileTime returns a boolean if a field has been set.

### GetFileType

`func (o *ApplianceFileGateway) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ApplianceFileGateway) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ApplianceFileGateway) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *ApplianceFileGateway) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFilename

`func (o *ApplianceFileGateway) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ApplianceFileGateway) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ApplianceFileGateway) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ApplianceFileGateway) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPresignedUrl

`func (o *ApplianceFileGateway) GetPresignedUrl() string`

GetPresignedUrl returns the PresignedUrl field if non-nil, zero value otherwise.

### GetPresignedUrlOk

`func (o *ApplianceFileGateway) GetPresignedUrlOk() (*string, bool)`

GetPresignedUrlOk returns a tuple with the PresignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresignedUrl

`func (o *ApplianceFileGateway) SetPresignedUrl(v string)`

SetPresignedUrl sets PresignedUrl field to given value.

### HasPresignedUrl

`func (o *ApplianceFileGateway) HasPresignedUrl() bool`

HasPresignedUrl returns a boolean if a field has been set.

### GetServerCert

`func (o *ApplianceFileGateway) GetServerCert() string`

GetServerCert returns the ServerCert field if non-nil, zero value otherwise.

### GetServerCertOk

`func (o *ApplianceFileGateway) GetServerCertOk() (*string, bool)`

GetServerCertOk returns a tuple with the ServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCert

`func (o *ApplianceFileGateway) SetServerCert(v string)`

SetServerCert sets ServerCert field to given value.

### HasServerCert

`func (o *ApplianceFileGateway) HasServerCert() bool`

HasServerCert returns a boolean if a field has been set.

### GetValidityPeriod

`func (o *ApplianceFileGateway) GetValidityPeriod() int64`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *ApplianceFileGateway) GetValidityPeriodOk() (*int64, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *ApplianceFileGateway) SetValidityPeriod(v int64)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *ApplianceFileGateway) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceFileGateway) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceFileGateway) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceFileGateway) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceFileGateway) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceFileGateway) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceFileGateway) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceFileGateway) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceFileGateway) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceFileGateway) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceFileGateway) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


