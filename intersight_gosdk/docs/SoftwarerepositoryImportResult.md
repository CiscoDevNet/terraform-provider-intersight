# SoftwarerepositoryImportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.ImportResult"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.ImportResult"]
**ErrorMessage** | Pointer to **string** | The reason for the failure of an import operation, if applicable. | [optional] [readonly] 
**Progress** | Pointer to **int64** | The progress percentage for the import operation. | [optional] [readonly] 

## Methods

### NewSoftwarerepositoryImportResult

`func NewSoftwarerepositoryImportResult(classId string, objectType string, ) *SoftwarerepositoryImportResult`

NewSoftwarerepositoryImportResult instantiates a new SoftwarerepositoryImportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryImportResultWithDefaults

`func NewSoftwarerepositoryImportResultWithDefaults() *SoftwarerepositoryImportResult`

NewSoftwarerepositoryImportResultWithDefaults instantiates a new SoftwarerepositoryImportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryImportResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryImportResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryImportResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryImportResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryImportResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryImportResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorMessage

`func (o *SoftwarerepositoryImportResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SoftwarerepositoryImportResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SoftwarerepositoryImportResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SoftwarerepositoryImportResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetProgress

`func (o *SoftwarerepositoryImportResult) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SoftwarerepositoryImportResult) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SoftwarerepositoryImportResult) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SoftwarerepositoryImportResult) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


