# HyperflexSoftwareDistributionComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SoftwareDistributionComponent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SoftwareDistributionComponent"]
**BucketName** | Pointer to **string** | The bucket name where the files are present, if source is external cloud store. | [optional] 
**ComponentId** | Pointer to **string** | The HyperFlex Software Distribution Component Identifier. | [optional] 
**ComponentName** | Pointer to **string** | The HyperFlex Software Distribution Component Name. | [optional] 
**FilePath** | Pointer to **string** | File location on the cloud storage. | [optional] 
**FilesToDownload** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** | The HyperFlex Software Distribution Component Version. | [optional] 
**SoftwareDistributionVersion** | Pointer to [**HyperflexSoftwareDistributionVersionRelationship**](HyperflexSoftwareDistributionVersionRelationship.md) |  | [optional] 

## Methods

### NewHyperflexSoftwareDistributionComponent

`func NewHyperflexSoftwareDistributionComponent(classId string, objectType string, ) *HyperflexSoftwareDistributionComponent`

NewHyperflexSoftwareDistributionComponent instantiates a new HyperflexSoftwareDistributionComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSoftwareDistributionComponentWithDefaults

`func NewHyperflexSoftwareDistributionComponentWithDefaults() *HyperflexSoftwareDistributionComponent`

NewHyperflexSoftwareDistributionComponentWithDefaults instantiates a new HyperflexSoftwareDistributionComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSoftwareDistributionComponent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSoftwareDistributionComponent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSoftwareDistributionComponent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSoftwareDistributionComponent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSoftwareDistributionComponent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSoftwareDistributionComponent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBucketName

`func (o *HyperflexSoftwareDistributionComponent) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *HyperflexSoftwareDistributionComponent) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *HyperflexSoftwareDistributionComponent) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *HyperflexSoftwareDistributionComponent) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetComponentId

`func (o *HyperflexSoftwareDistributionComponent) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *HyperflexSoftwareDistributionComponent) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *HyperflexSoftwareDistributionComponent) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *HyperflexSoftwareDistributionComponent) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetComponentName

`func (o *HyperflexSoftwareDistributionComponent) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *HyperflexSoftwareDistributionComponent) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *HyperflexSoftwareDistributionComponent) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *HyperflexSoftwareDistributionComponent) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetFilePath

`func (o *HyperflexSoftwareDistributionComponent) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *HyperflexSoftwareDistributionComponent) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *HyperflexSoftwareDistributionComponent) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *HyperflexSoftwareDistributionComponent) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFilesToDownload

`func (o *HyperflexSoftwareDistributionComponent) GetFilesToDownload() []string`

GetFilesToDownload returns the FilesToDownload field if non-nil, zero value otherwise.

### GetFilesToDownloadOk

`func (o *HyperflexSoftwareDistributionComponent) GetFilesToDownloadOk() (*[]string, bool)`

GetFilesToDownloadOk returns a tuple with the FilesToDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesToDownload

`func (o *HyperflexSoftwareDistributionComponent) SetFilesToDownload(v []string)`

SetFilesToDownload sets FilesToDownload field to given value.

### HasFilesToDownload

`func (o *HyperflexSoftwareDistributionComponent) HasFilesToDownload() bool`

HasFilesToDownload returns a boolean if a field has been set.

### SetFilesToDownloadNil

`func (o *HyperflexSoftwareDistributionComponent) SetFilesToDownloadNil(b bool)`

 SetFilesToDownloadNil sets the value for FilesToDownload to be an explicit nil

### UnsetFilesToDownload
`func (o *HyperflexSoftwareDistributionComponent) UnsetFilesToDownload()`

UnsetFilesToDownload ensures that no value is present for FilesToDownload, not even an explicit nil
### GetVersion

`func (o *HyperflexSoftwareDistributionComponent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexSoftwareDistributionComponent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexSoftwareDistributionComponent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexSoftwareDistributionComponent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSoftwareDistributionVersion

`func (o *HyperflexSoftwareDistributionComponent) GetSoftwareDistributionVersion() HyperflexSoftwareDistributionVersionRelationship`

GetSoftwareDistributionVersion returns the SoftwareDistributionVersion field if non-nil, zero value otherwise.

### GetSoftwareDistributionVersionOk

`func (o *HyperflexSoftwareDistributionComponent) GetSoftwareDistributionVersionOk() (*HyperflexSoftwareDistributionVersionRelationship, bool)`

GetSoftwareDistributionVersionOk returns a tuple with the SoftwareDistributionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDistributionVersion

`func (o *HyperflexSoftwareDistributionComponent) SetSoftwareDistributionVersion(v HyperflexSoftwareDistributionVersionRelationship)`

SetSoftwareDistributionVersion sets SoftwareDistributionVersion field to given value.

### HasSoftwareDistributionVersion

`func (o *HyperflexSoftwareDistributionComponent) HasSoftwareDistributionVersion() bool`

HasSoftwareDistributionVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


