# HyperflexFilePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.FilePath"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.FilePath"]
**DsInfo** | Pointer to [**NullableHyperflexDatastoreInfo**](HyperflexDatastoreInfo.md) |  | [optional] 
**RelativeFilePath** | Pointer to **string** | Relative file path within the datastore. | [optional] [readonly] 

## Methods

### NewHyperflexFilePath

`func NewHyperflexFilePath(classId string, objectType string, ) *HyperflexFilePath`

NewHyperflexFilePath instantiates a new HyperflexFilePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexFilePathWithDefaults

`func NewHyperflexFilePathWithDefaults() *HyperflexFilePath`

NewHyperflexFilePathWithDefaults instantiates a new HyperflexFilePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexFilePath) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexFilePath) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexFilePath) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexFilePath) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexFilePath) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexFilePath) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDsInfo

`func (o *HyperflexFilePath) GetDsInfo() HyperflexDatastoreInfo`

GetDsInfo returns the DsInfo field if non-nil, zero value otherwise.

### GetDsInfoOk

`func (o *HyperflexFilePath) GetDsInfoOk() (*HyperflexDatastoreInfo, bool)`

GetDsInfoOk returns a tuple with the DsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsInfo

`func (o *HyperflexFilePath) SetDsInfo(v HyperflexDatastoreInfo)`

SetDsInfo sets DsInfo field to given value.

### HasDsInfo

`func (o *HyperflexFilePath) HasDsInfo() bool`

HasDsInfo returns a boolean if a field has been set.

### SetDsInfoNil

`func (o *HyperflexFilePath) SetDsInfoNil(b bool)`

 SetDsInfoNil sets the value for DsInfo to be an explicit nil

### UnsetDsInfo
`func (o *HyperflexFilePath) UnsetDsInfo()`

UnsetDsInfo ensures that no value is present for DsInfo, not even an explicit nil
### GetRelativeFilePath

`func (o *HyperflexFilePath) GetRelativeFilePath() string`

GetRelativeFilePath returns the RelativeFilePath field if non-nil, zero value otherwise.

### GetRelativeFilePathOk

`func (o *HyperflexFilePath) GetRelativeFilePathOk() (*string, bool)`

GetRelativeFilePathOk returns a tuple with the RelativeFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeFilePath

`func (o *HyperflexFilePath) SetRelativeFilePath(v string)`

SetRelativeFilePath sets RelativeFilePath field to given value.

### HasRelativeFilePath

`func (o *HyperflexFilePath) HasRelativeFilePath() bool`

HasRelativeFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


