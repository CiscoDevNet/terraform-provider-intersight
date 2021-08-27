# HyperflexSoftwareDistributionComponentAllOf

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

### NewHyperflexSoftwareDistributionComponentAllOf

`func NewHyperflexSoftwareDistributionComponentAllOf(classId string, objectType string, ) *HyperflexSoftwareDistributionComponentAllOf`

NewHyperflexSoftwareDistributionComponentAllOf instantiates a new HyperflexSoftwareDistributionComponentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSoftwareDistributionComponentAllOfWithDefaults

`func NewHyperflexSoftwareDistributionComponentAllOfWithDefaults() *HyperflexSoftwareDistributionComponentAllOf`

NewHyperflexSoftwareDistributionComponentAllOfWithDefaults instantiates a new HyperflexSoftwareDistributionComponentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBucketName

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *HyperflexSoftwareDistributionComponentAllOf) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetComponentId

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *HyperflexSoftwareDistributionComponentAllOf) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetComponentName

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *HyperflexSoftwareDistributionComponentAllOf) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetFilePath

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *HyperflexSoftwareDistributionComponentAllOf) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFilesToDownload

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilesToDownload() []string`

GetFilesToDownload returns the FilesToDownload field if non-nil, zero value otherwise.

### GetFilesToDownloadOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilesToDownloadOk() (*[]string, bool)`

GetFilesToDownloadOk returns a tuple with the FilesToDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesToDownload

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetFilesToDownload(v []string)`

SetFilesToDownload sets FilesToDownload field to given value.

### HasFilesToDownload

`func (o *HyperflexSoftwareDistributionComponentAllOf) HasFilesToDownload() bool`

HasFilesToDownload returns a boolean if a field has been set.

### SetFilesToDownloadNil

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetFilesToDownloadNil(b bool)`

 SetFilesToDownloadNil sets the value for FilesToDownload to be an explicit nil

### UnsetFilesToDownload
`func (o *HyperflexSoftwareDistributionComponentAllOf) UnsetFilesToDownload()`

UnsetFilesToDownload ensures that no value is present for FilesToDownload, not even an explicit nil
### GetVersion

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexSoftwareDistributionComponentAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSoftwareDistributionVersion

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetSoftwareDistributionVersion() HyperflexSoftwareDistributionVersionRelationship`

GetSoftwareDistributionVersion returns the SoftwareDistributionVersion field if non-nil, zero value otherwise.

### GetSoftwareDistributionVersionOk

`func (o *HyperflexSoftwareDistributionComponentAllOf) GetSoftwareDistributionVersionOk() (*HyperflexSoftwareDistributionVersionRelationship, bool)`

GetSoftwareDistributionVersionOk returns a tuple with the SoftwareDistributionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDistributionVersion

`func (o *HyperflexSoftwareDistributionComponentAllOf) SetSoftwareDistributionVersion(v HyperflexSoftwareDistributionVersionRelationship)`

SetSoftwareDistributionVersion sets SoftwareDistributionVersion field to given value.

### HasSoftwareDistributionVersion

`func (o *HyperflexSoftwareDistributionComponentAllOf) HasSoftwareDistributionVersion() bool`

HasSoftwareDistributionVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


