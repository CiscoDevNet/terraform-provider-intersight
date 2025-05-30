# CoremanagementDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "coremanagement.Download"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "coremanagement.Download"]
**CoreFile** | Pointer to [**NullableCoremanagementCoreFileRelationship**](CoremanagementCoreFileRelationship.md) |  | [optional] 

## Methods

### NewCoremanagementDownload

`func NewCoremanagementDownload(classId string, objectType string, ) *CoremanagementDownload`

NewCoremanagementDownload instantiates a new CoremanagementDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoremanagementDownloadWithDefaults

`func NewCoremanagementDownloadWithDefaults() *CoremanagementDownload`

NewCoremanagementDownloadWithDefaults instantiates a new CoremanagementDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CoremanagementDownload) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CoremanagementDownload) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CoremanagementDownload) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CoremanagementDownload) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CoremanagementDownload) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CoremanagementDownload) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCoreFile

`func (o *CoremanagementDownload) GetCoreFile() CoremanagementCoreFileRelationship`

GetCoreFile returns the CoreFile field if non-nil, zero value otherwise.

### GetCoreFileOk

`func (o *CoremanagementDownload) GetCoreFileOk() (*CoremanagementCoreFileRelationship, bool)`

GetCoreFileOk returns a tuple with the CoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreFile

`func (o *CoremanagementDownload) SetCoreFile(v CoremanagementCoreFileRelationship)`

SetCoreFile sets CoreFile field to given value.

### HasCoreFile

`func (o *CoremanagementDownload) HasCoreFile() bool`

HasCoreFile returns a boolean if a field has been set.

### SetCoreFileNil

`func (o *CoremanagementDownload) SetCoreFileNil(b bool)`

 SetCoreFileNil sets the value for CoreFile to be an explicit nil

### UnsetCoreFile
`func (o *CoremanagementDownload) UnsetCoreFile()`

UnsetCoreFile ensures that no value is present for CoreFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


