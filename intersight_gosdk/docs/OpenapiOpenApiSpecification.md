# OpenapiOpenApiSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.OpenApiSpecification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.OpenApiSpecification"]
**FilePath** | Pointer to **string** | The path of the file in S3/minio bucket. | [optional] [readonly] 
**FileUploadId** | Pointer to **string** | A unique tracking ID for the file being uploaded. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewOpenapiOpenApiSpecification

`func NewOpenapiOpenApiSpecification(classId string, objectType string, ) *OpenapiOpenApiSpecification`

NewOpenapiOpenApiSpecification instantiates a new OpenapiOpenApiSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiOpenApiSpecificationWithDefaults

`func NewOpenapiOpenApiSpecificationWithDefaults() *OpenapiOpenApiSpecification`

NewOpenapiOpenApiSpecificationWithDefaults instantiates a new OpenapiOpenApiSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiOpenApiSpecification) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiOpenApiSpecification) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiOpenApiSpecification) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiOpenApiSpecification) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiOpenApiSpecification) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiOpenApiSpecification) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilePath

`func (o *OpenapiOpenApiSpecification) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *OpenapiOpenApiSpecification) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *OpenapiOpenApiSpecification) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *OpenapiOpenApiSpecification) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileUploadId

`func (o *OpenapiOpenApiSpecification) GetFileUploadId() string`

GetFileUploadId returns the FileUploadId field if non-nil, zero value otherwise.

### GetFileUploadIdOk

`func (o *OpenapiOpenApiSpecification) GetFileUploadIdOk() (*string, bool)`

GetFileUploadIdOk returns a tuple with the FileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadId

`func (o *OpenapiOpenApiSpecification) SetFileUploadId(v string)`

SetFileUploadId sets FileUploadId field to given value.

### HasFileUploadId

`func (o *OpenapiOpenApiSpecification) HasFileUploadId() bool`

HasFileUploadId returns a boolean if a field has been set.

### GetOrganization

`func (o *OpenapiOpenApiSpecification) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OpenapiOpenApiSpecification) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OpenapiOpenApiSpecification) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OpenapiOpenApiSpecification) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


