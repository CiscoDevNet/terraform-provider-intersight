# OpenapiOpenApiSpecificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.OpenApiSpecification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.OpenApiSpecification"]
**FilePath** | Pointer to **string** | The path of the file in S3/minio bucket. | [optional] [readonly] 
**FileUploadId** | Pointer to **string** | A unique tracking ID for the file being uploaded. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewOpenapiOpenApiSpecificationAllOf

`func NewOpenapiOpenApiSpecificationAllOf(classId string, objectType string, ) *OpenapiOpenApiSpecificationAllOf`

NewOpenapiOpenApiSpecificationAllOf instantiates a new OpenapiOpenApiSpecificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiOpenApiSpecificationAllOfWithDefaults

`func NewOpenapiOpenApiSpecificationAllOfWithDefaults() *OpenapiOpenApiSpecificationAllOf`

NewOpenapiOpenApiSpecificationAllOfWithDefaults instantiates a new OpenapiOpenApiSpecificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiOpenApiSpecificationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiOpenApiSpecificationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiOpenApiSpecificationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiOpenApiSpecificationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiOpenApiSpecificationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiOpenApiSpecificationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilePath

`func (o *OpenapiOpenApiSpecificationAllOf) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *OpenapiOpenApiSpecificationAllOf) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *OpenapiOpenApiSpecificationAllOf) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *OpenapiOpenApiSpecificationAllOf) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileUploadId

`func (o *OpenapiOpenApiSpecificationAllOf) GetFileUploadId() string`

GetFileUploadId returns the FileUploadId field if non-nil, zero value otherwise.

### GetFileUploadIdOk

`func (o *OpenapiOpenApiSpecificationAllOf) GetFileUploadIdOk() (*string, bool)`

GetFileUploadIdOk returns a tuple with the FileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadId

`func (o *OpenapiOpenApiSpecificationAllOf) SetFileUploadId(v string)`

SetFileUploadId sets FileUploadId field to given value.

### HasFileUploadId

`func (o *OpenapiOpenApiSpecificationAllOf) HasFileUploadId() bool`

HasFileUploadId returns a boolean if a field has been set.

### GetOrganization

`func (o *OpenapiOpenApiSpecificationAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OpenapiOpenApiSpecificationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OpenapiOpenApiSpecificationAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OpenapiOpenApiSpecificationAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


