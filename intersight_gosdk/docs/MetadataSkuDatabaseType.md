# MetadataSkuDatabaseType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metadata.SkuDatabaseType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metadata.SkuDatabaseType"]
**DatabaseEdition** | Pointer to **string** | The database edition. For e.g. standard or enterprise. | [optional] 
**DatabaseEngine** | Pointer to **string** | The database engine. For e.g. SQL Server, Oracle, PostgreSQL. | [optional] 
**LicenseModel** | Pointer to **string** | The licensing option for the database. For e.g. license required or not. | [optional] 
**NetworkPerformance** | Pointer to **string** | Network performance of this instance type. | [optional] 

## Methods

### NewMetadataSkuDatabaseType

`func NewMetadataSkuDatabaseType(classId string, objectType string, ) *MetadataSkuDatabaseType`

NewMetadataSkuDatabaseType instantiates a new MetadataSkuDatabaseType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSkuDatabaseTypeWithDefaults

`func NewMetadataSkuDatabaseTypeWithDefaults() *MetadataSkuDatabaseType`

NewMetadataSkuDatabaseTypeWithDefaults instantiates a new MetadataSkuDatabaseType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetadataSkuDatabaseType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetadataSkuDatabaseType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetadataSkuDatabaseType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetadataSkuDatabaseType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetadataSkuDatabaseType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetadataSkuDatabaseType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatabaseEdition

`func (o *MetadataSkuDatabaseType) GetDatabaseEdition() string`

GetDatabaseEdition returns the DatabaseEdition field if non-nil, zero value otherwise.

### GetDatabaseEditionOk

`func (o *MetadataSkuDatabaseType) GetDatabaseEditionOk() (*string, bool)`

GetDatabaseEditionOk returns a tuple with the DatabaseEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEdition

`func (o *MetadataSkuDatabaseType) SetDatabaseEdition(v string)`

SetDatabaseEdition sets DatabaseEdition field to given value.

### HasDatabaseEdition

`func (o *MetadataSkuDatabaseType) HasDatabaseEdition() bool`

HasDatabaseEdition returns a boolean if a field has been set.

### GetDatabaseEngine

`func (o *MetadataSkuDatabaseType) GetDatabaseEngine() string`

GetDatabaseEngine returns the DatabaseEngine field if non-nil, zero value otherwise.

### GetDatabaseEngineOk

`func (o *MetadataSkuDatabaseType) GetDatabaseEngineOk() (*string, bool)`

GetDatabaseEngineOk returns a tuple with the DatabaseEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEngine

`func (o *MetadataSkuDatabaseType) SetDatabaseEngine(v string)`

SetDatabaseEngine sets DatabaseEngine field to given value.

### HasDatabaseEngine

`func (o *MetadataSkuDatabaseType) HasDatabaseEngine() bool`

HasDatabaseEngine returns a boolean if a field has been set.

### GetLicenseModel

`func (o *MetadataSkuDatabaseType) GetLicenseModel() string`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *MetadataSkuDatabaseType) GetLicenseModelOk() (*string, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *MetadataSkuDatabaseType) SetLicenseModel(v string)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *MetadataSkuDatabaseType) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetNetworkPerformance

`func (o *MetadataSkuDatabaseType) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *MetadataSkuDatabaseType) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *MetadataSkuDatabaseType) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *MetadataSkuDatabaseType) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


