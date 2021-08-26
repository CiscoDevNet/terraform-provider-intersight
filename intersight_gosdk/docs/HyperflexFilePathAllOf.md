# HyperflexFilePathAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.FilePath"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.FilePath"]
**DsInfo** | Pointer to [**NullableHyperflexDatastoreInfo**](HyperflexDatastoreInfo.md) |  | [optional] 
**RelativeFilePath** | Pointer to **string** | Relative file path within the datastore. | [optional] [readonly] 

## Methods

### NewHyperflexFilePathAllOf

`func NewHyperflexFilePathAllOf(classId string, objectType string, ) *HyperflexFilePathAllOf`

NewHyperflexFilePathAllOf instantiates a new HyperflexFilePathAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexFilePathAllOfWithDefaults

`func NewHyperflexFilePathAllOfWithDefaults() *HyperflexFilePathAllOf`

NewHyperflexFilePathAllOfWithDefaults instantiates a new HyperflexFilePathAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexFilePathAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexFilePathAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexFilePathAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexFilePathAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexFilePathAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexFilePathAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDsInfo

`func (o *HyperflexFilePathAllOf) GetDsInfo() HyperflexDatastoreInfo`

GetDsInfo returns the DsInfo field if non-nil, zero value otherwise.

### GetDsInfoOk

`func (o *HyperflexFilePathAllOf) GetDsInfoOk() (*HyperflexDatastoreInfo, bool)`

GetDsInfoOk returns a tuple with the DsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsInfo

`func (o *HyperflexFilePathAllOf) SetDsInfo(v HyperflexDatastoreInfo)`

SetDsInfo sets DsInfo field to given value.

### HasDsInfo

`func (o *HyperflexFilePathAllOf) HasDsInfo() bool`

HasDsInfo returns a boolean if a field has been set.

### SetDsInfoNil

`func (o *HyperflexFilePathAllOf) SetDsInfoNil(b bool)`

 SetDsInfoNil sets the value for DsInfo to be an explicit nil

### UnsetDsInfo
`func (o *HyperflexFilePathAllOf) UnsetDsInfo()`

UnsetDsInfo ensures that no value is present for DsInfo, not even an explicit nil
### GetRelativeFilePath

`func (o *HyperflexFilePathAllOf) GetRelativeFilePath() string`

GetRelativeFilePath returns the RelativeFilePath field if non-nil, zero value otherwise.

### GetRelativeFilePathOk

`func (o *HyperflexFilePathAllOf) GetRelativeFilePathOk() (*string, bool)`

GetRelativeFilePathOk returns a tuple with the RelativeFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeFilePath

`func (o *HyperflexFilePathAllOf) SetRelativeFilePath(v string)`

SetRelativeFilePath sets RelativeFilePath field to given value.

### HasRelativeFilePath

`func (o *HyperflexFilePathAllOf) HasRelativeFilePath() bool`

HasRelativeFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


