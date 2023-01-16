# OpenapiProcessFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.ProcessFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.ProcessFile"]
**FailureReason** | Pointer to **string** | An error message for when download, validation or processing of OpenAPI Specification fails. | [optional] [readonly] 
**FileDownloadStatus** | Pointer to **string** | Status of the internal OpenAPI specification fetch operation * &#x60;none&#x60; - Indicates the default status * &#x60;InProgress&#x60; - Indicates that operation is in progress * &#x60;Completed&#x60; - Indicates that the operation is complete * &#x60;Failed&#x60; - Indicates that the operation has failed. Check the failureReason attribute for more details. | [optional] [readonly] [default to "none"]
**FileProcessingStatus** | Pointer to **string** | Status of the OpenAPI specification processing operation. The OpenAPI specification file is processed to create APIMethod objects. * &#x60;none&#x60; - Indicates the default status * &#x60;InProgress&#x60; - Indicates that operation is in progress * &#x60;Completed&#x60; - Indicates that the operation is complete * &#x60;Failed&#x60; - Indicates that the operation has failed. Check the failureReason attribute for more details. | [optional] [readonly] [default to "none"]
**FileValidationStatus** | Pointer to **string** | Status of the OpenAPI specification validation operation. * &#x60;none&#x60; - Indicates the default status * &#x60;InProgress&#x60; - Indicates that operation is in progress * &#x60;Completed&#x60; - Indicates that the operation is complete * &#x60;Failed&#x60; - Indicates that the operation has failed. Check the failureReason attribute for more details. | [optional] [readonly] [default to "none"]
**SpecFilePath** | Pointer to **string** | The location of the previously uploaded OpenAPI specification file. | [optional] 
**FileInfo** | Pointer to [**OpenapiOpenApiSpecificationRelationship**](OpenapiOpenApiSpecificationRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewOpenapiProcessFile

`func NewOpenapiProcessFile(classId string, objectType string, ) *OpenapiProcessFile`

NewOpenapiProcessFile instantiates a new OpenapiProcessFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiProcessFileWithDefaults

`func NewOpenapiProcessFileWithDefaults() *OpenapiProcessFile`

NewOpenapiProcessFileWithDefaults instantiates a new OpenapiProcessFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiProcessFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiProcessFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiProcessFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiProcessFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiProcessFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiProcessFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailureReason

`func (o *OpenapiProcessFile) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *OpenapiProcessFile) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *OpenapiProcessFile) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *OpenapiProcessFile) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFileDownloadStatus

`func (o *OpenapiProcessFile) GetFileDownloadStatus() string`

GetFileDownloadStatus returns the FileDownloadStatus field if non-nil, zero value otherwise.

### GetFileDownloadStatusOk

`func (o *OpenapiProcessFile) GetFileDownloadStatusOk() (*string, bool)`

GetFileDownloadStatusOk returns a tuple with the FileDownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadStatus

`func (o *OpenapiProcessFile) SetFileDownloadStatus(v string)`

SetFileDownloadStatus sets FileDownloadStatus field to given value.

### HasFileDownloadStatus

`func (o *OpenapiProcessFile) HasFileDownloadStatus() bool`

HasFileDownloadStatus returns a boolean if a field has been set.

### GetFileProcessingStatus

`func (o *OpenapiProcessFile) GetFileProcessingStatus() string`

GetFileProcessingStatus returns the FileProcessingStatus field if non-nil, zero value otherwise.

### GetFileProcessingStatusOk

`func (o *OpenapiProcessFile) GetFileProcessingStatusOk() (*string, bool)`

GetFileProcessingStatusOk returns a tuple with the FileProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileProcessingStatus

`func (o *OpenapiProcessFile) SetFileProcessingStatus(v string)`

SetFileProcessingStatus sets FileProcessingStatus field to given value.

### HasFileProcessingStatus

`func (o *OpenapiProcessFile) HasFileProcessingStatus() bool`

HasFileProcessingStatus returns a boolean if a field has been set.

### GetFileValidationStatus

`func (o *OpenapiProcessFile) GetFileValidationStatus() string`

GetFileValidationStatus returns the FileValidationStatus field if non-nil, zero value otherwise.

### GetFileValidationStatusOk

`func (o *OpenapiProcessFile) GetFileValidationStatusOk() (*string, bool)`

GetFileValidationStatusOk returns a tuple with the FileValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileValidationStatus

`func (o *OpenapiProcessFile) SetFileValidationStatus(v string)`

SetFileValidationStatus sets FileValidationStatus field to given value.

### HasFileValidationStatus

`func (o *OpenapiProcessFile) HasFileValidationStatus() bool`

HasFileValidationStatus returns a boolean if a field has been set.

### GetSpecFilePath

`func (o *OpenapiProcessFile) GetSpecFilePath() string`

GetSpecFilePath returns the SpecFilePath field if non-nil, zero value otherwise.

### GetSpecFilePathOk

`func (o *OpenapiProcessFile) GetSpecFilePathOk() (*string, bool)`

GetSpecFilePathOk returns a tuple with the SpecFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecFilePath

`func (o *OpenapiProcessFile) SetSpecFilePath(v string)`

SetSpecFilePath sets SpecFilePath field to given value.

### HasSpecFilePath

`func (o *OpenapiProcessFile) HasSpecFilePath() bool`

HasSpecFilePath returns a boolean if a field has been set.

### GetFileInfo

`func (o *OpenapiProcessFile) GetFileInfo() OpenapiOpenApiSpecificationRelationship`

GetFileInfo returns the FileInfo field if non-nil, zero value otherwise.

### GetFileInfoOk

`func (o *OpenapiProcessFile) GetFileInfoOk() (*OpenapiOpenApiSpecificationRelationship, bool)`

GetFileInfoOk returns a tuple with the FileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileInfo

`func (o *OpenapiProcessFile) SetFileInfo(v OpenapiOpenApiSpecificationRelationship)`

SetFileInfo sets FileInfo field to given value.

### HasFileInfo

`func (o *OpenapiProcessFile) HasFileInfo() bool`

HasFileInfo returns a boolean if a field has been set.

### GetOrganization

`func (o *OpenapiProcessFile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OpenapiProcessFile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OpenapiProcessFile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OpenapiProcessFile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


