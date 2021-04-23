# TamS3DataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.S3DataSource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.S3DataSource"]
**Queries** | Pointer to [**[]TamQueryEntry**](TamQueryEntry.md) |  | [optional] 
**S3Path** | Pointer to **string** | Path used to access file in s3 containing data. | [optional] 

## Methods

### NewTamS3DataSource

`func NewTamS3DataSource(classId string, objectType string, ) *TamS3DataSource`

NewTamS3DataSource instantiates a new TamS3DataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamS3DataSourceWithDefaults

`func NewTamS3DataSourceWithDefaults() *TamS3DataSource`

NewTamS3DataSourceWithDefaults instantiates a new TamS3DataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamS3DataSource) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamS3DataSource) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamS3DataSource) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamS3DataSource) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamS3DataSource) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamS3DataSource) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetQueries

`func (o *TamS3DataSource) GetQueries() []TamQueryEntry`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TamS3DataSource) GetQueriesOk() (*[]TamQueryEntry, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TamS3DataSource) SetQueries(v []TamQueryEntry)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TamS3DataSource) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### SetQueriesNil

`func (o *TamS3DataSource) SetQueriesNil(b bool)`

 SetQueriesNil sets the value for Queries to be an explicit nil

### UnsetQueries
`func (o *TamS3DataSource) UnsetQueries()`

UnsetQueries ensures that no value is present for Queries, not even an explicit nil
### GetS3Path

`func (o *TamS3DataSource) GetS3Path() string`

GetS3Path returns the S3Path field if non-nil, zero value otherwise.

### GetS3PathOk

`func (o *TamS3DataSource) GetS3PathOk() (*string, bool)`

GetS3PathOk returns a tuple with the S3Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Path

`func (o *TamS3DataSource) SetS3Path(v string)`

SetS3Path sets S3Path field to given value.

### HasS3Path

`func (o *TamS3DataSource) HasS3Path() bool`

HasS3Path returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


