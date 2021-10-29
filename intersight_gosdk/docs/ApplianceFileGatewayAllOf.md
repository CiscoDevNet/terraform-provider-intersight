# ApplianceFileGatewayAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.FileGateway"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.FileGateway"]
**BucketName** | Pointer to **string** | Bucket name in the cloud storage service where the file is stored. | [optional] [readonly] 
**FileSize** | Pointer to **int64** | Size of the file in bytes. FileSize maybe zero if the storage service does not report file size. | [optional] [readonly] 
**FileTime** | Pointer to **time.Time** | File timestamp as reported by the cloud storage service. | [optional] [readonly] 
**FileType** | Pointer to **string** | User defined file type supplied by the caller. | [optional] [readonly] 
**Filename** | Pointer to **string** | Full name of the file as it exists in the cloud storage service. If the file is in a subfolder, then the filename must contain the full path. | [optional] [readonly] 
**PresignedUrl** | Pointer to **string** | Pre-signed URL obtained from the cloud storage service. | [optional] [readonly] 
**ServerCert** | Pointer to **string** | SSL certificate of the cloud storage service. | [optional] [readonly] 
**ValidityPeriod** | Pointer to **int64** | Signed URL&#39;s validity period in minutes. | [optional] [readonly] 
**Version** | Pointer to **string** | File version as reported by the cloud storage service. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceFileGatewayAllOf

`func NewApplianceFileGatewayAllOf(classId string, objectType string, ) *ApplianceFileGatewayAllOf`

NewApplianceFileGatewayAllOf instantiates a new ApplianceFileGatewayAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceFileGatewayAllOfWithDefaults

`func NewApplianceFileGatewayAllOfWithDefaults() *ApplianceFileGatewayAllOf`

NewApplianceFileGatewayAllOfWithDefaults instantiates a new ApplianceFileGatewayAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceFileGatewayAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceFileGatewayAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceFileGatewayAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceFileGatewayAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceFileGatewayAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceFileGatewayAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBucketName

`func (o *ApplianceFileGatewayAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ApplianceFileGatewayAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ApplianceFileGatewayAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ApplianceFileGatewayAllOf) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetFileSize

`func (o *ApplianceFileGatewayAllOf) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ApplianceFileGatewayAllOf) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ApplianceFileGatewayAllOf) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *ApplianceFileGatewayAllOf) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileTime

`func (o *ApplianceFileGatewayAllOf) GetFileTime() time.Time`

GetFileTime returns the FileTime field if non-nil, zero value otherwise.

### GetFileTimeOk

`func (o *ApplianceFileGatewayAllOf) GetFileTimeOk() (*time.Time, bool)`

GetFileTimeOk returns a tuple with the FileTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTime

`func (o *ApplianceFileGatewayAllOf) SetFileTime(v time.Time)`

SetFileTime sets FileTime field to given value.

### HasFileTime

`func (o *ApplianceFileGatewayAllOf) HasFileTime() bool`

HasFileTime returns a boolean if a field has been set.

### GetFileType

`func (o *ApplianceFileGatewayAllOf) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ApplianceFileGatewayAllOf) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ApplianceFileGatewayAllOf) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *ApplianceFileGatewayAllOf) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFilename

`func (o *ApplianceFileGatewayAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ApplianceFileGatewayAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ApplianceFileGatewayAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ApplianceFileGatewayAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPresignedUrl

`func (o *ApplianceFileGatewayAllOf) GetPresignedUrl() string`

GetPresignedUrl returns the PresignedUrl field if non-nil, zero value otherwise.

### GetPresignedUrlOk

`func (o *ApplianceFileGatewayAllOf) GetPresignedUrlOk() (*string, bool)`

GetPresignedUrlOk returns a tuple with the PresignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresignedUrl

`func (o *ApplianceFileGatewayAllOf) SetPresignedUrl(v string)`

SetPresignedUrl sets PresignedUrl field to given value.

### HasPresignedUrl

`func (o *ApplianceFileGatewayAllOf) HasPresignedUrl() bool`

HasPresignedUrl returns a boolean if a field has been set.

### GetServerCert

`func (o *ApplianceFileGatewayAllOf) GetServerCert() string`

GetServerCert returns the ServerCert field if non-nil, zero value otherwise.

### GetServerCertOk

`func (o *ApplianceFileGatewayAllOf) GetServerCertOk() (*string, bool)`

GetServerCertOk returns a tuple with the ServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCert

`func (o *ApplianceFileGatewayAllOf) SetServerCert(v string)`

SetServerCert sets ServerCert field to given value.

### HasServerCert

`func (o *ApplianceFileGatewayAllOf) HasServerCert() bool`

HasServerCert returns a boolean if a field has been set.

### GetValidityPeriod

`func (o *ApplianceFileGatewayAllOf) GetValidityPeriod() int64`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *ApplianceFileGatewayAllOf) GetValidityPeriodOk() (*int64, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *ApplianceFileGatewayAllOf) SetValidityPeriod(v int64)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *ApplianceFileGatewayAllOf) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceFileGatewayAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceFileGatewayAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceFileGatewayAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceFileGatewayAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceFileGatewayAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceFileGatewayAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceFileGatewayAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceFileGatewayAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


