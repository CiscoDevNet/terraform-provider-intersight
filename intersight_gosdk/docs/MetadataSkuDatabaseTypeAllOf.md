# MetadataSkuDatabaseTypeAllOf

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

### NewMetadataSkuDatabaseTypeAllOf

`func NewMetadataSkuDatabaseTypeAllOf(classId string, objectType string, ) *MetadataSkuDatabaseTypeAllOf`

NewMetadataSkuDatabaseTypeAllOf instantiates a new MetadataSkuDatabaseTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSkuDatabaseTypeAllOfWithDefaults

`func NewMetadataSkuDatabaseTypeAllOfWithDefaults() *MetadataSkuDatabaseTypeAllOf`

NewMetadataSkuDatabaseTypeAllOfWithDefaults instantiates a new MetadataSkuDatabaseTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetadataSkuDatabaseTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetadataSkuDatabaseTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetadataSkuDatabaseTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetadataSkuDatabaseTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetadataSkuDatabaseTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetadataSkuDatabaseTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatabaseEdition

`func (o *MetadataSkuDatabaseTypeAllOf) GetDatabaseEdition() string`

GetDatabaseEdition returns the DatabaseEdition field if non-nil, zero value otherwise.

### GetDatabaseEditionOk

`func (o *MetadataSkuDatabaseTypeAllOf) GetDatabaseEditionOk() (*string, bool)`

GetDatabaseEditionOk returns a tuple with the DatabaseEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEdition

`func (o *MetadataSkuDatabaseTypeAllOf) SetDatabaseEdition(v string)`

SetDatabaseEdition sets DatabaseEdition field to given value.

### HasDatabaseEdition

`func (o *MetadataSkuDatabaseTypeAllOf) HasDatabaseEdition() bool`

HasDatabaseEdition returns a boolean if a field has been set.

### GetDatabaseEngine

`func (o *MetadataSkuDatabaseTypeAllOf) GetDatabaseEngine() string`

GetDatabaseEngine returns the DatabaseEngine field if non-nil, zero value otherwise.

### GetDatabaseEngineOk

`func (o *MetadataSkuDatabaseTypeAllOf) GetDatabaseEngineOk() (*string, bool)`

GetDatabaseEngineOk returns a tuple with the DatabaseEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEngine

`func (o *MetadataSkuDatabaseTypeAllOf) SetDatabaseEngine(v string)`

SetDatabaseEngine sets DatabaseEngine field to given value.

### HasDatabaseEngine

`func (o *MetadataSkuDatabaseTypeAllOf) HasDatabaseEngine() bool`

HasDatabaseEngine returns a boolean if a field has been set.

### GetLicenseModel

`func (o *MetadataSkuDatabaseTypeAllOf) GetLicenseModel() string`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *MetadataSkuDatabaseTypeAllOf) GetLicenseModelOk() (*string, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *MetadataSkuDatabaseTypeAllOf) SetLicenseModel(v string)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *MetadataSkuDatabaseTypeAllOf) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetNetworkPerformance

`func (o *MetadataSkuDatabaseTypeAllOf) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *MetadataSkuDatabaseTypeAllOf) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *MetadataSkuDatabaseTypeAllOf) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *MetadataSkuDatabaseTypeAllOf) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


