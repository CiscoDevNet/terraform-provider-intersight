# CloudSkuDatabaseTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.SkuDatabaseType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.SkuDatabaseType"]
**DatabaseEdition** | Pointer to **string** | The database edition. For e.g. standard or enterprise. | [optional] 
**DatabaseEngine** | Pointer to **string** | The database engine. For e.g. SQL Server, Oracle, PostgreSQL. | [optional] 
**LicenseModel** | Pointer to **string** | The licensing option for the database. For e.g. license required or not. | [optional] 
**NetworkPerformance** | Pointer to **string** | Network performance of this instance type. | [optional] 

## Methods

### NewCloudSkuDatabaseTypeAllOf

`func NewCloudSkuDatabaseTypeAllOf(classId string, objectType string, ) *CloudSkuDatabaseTypeAllOf`

NewCloudSkuDatabaseTypeAllOf instantiates a new CloudSkuDatabaseTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSkuDatabaseTypeAllOfWithDefaults

`func NewCloudSkuDatabaseTypeAllOfWithDefaults() *CloudSkuDatabaseTypeAllOf`

NewCloudSkuDatabaseTypeAllOfWithDefaults instantiates a new CloudSkuDatabaseTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudSkuDatabaseTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudSkuDatabaseTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudSkuDatabaseTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudSkuDatabaseTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudSkuDatabaseTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudSkuDatabaseTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatabaseEdition

`func (o *CloudSkuDatabaseTypeAllOf) GetDatabaseEdition() string`

GetDatabaseEdition returns the DatabaseEdition field if non-nil, zero value otherwise.

### GetDatabaseEditionOk

`func (o *CloudSkuDatabaseTypeAllOf) GetDatabaseEditionOk() (*string, bool)`

GetDatabaseEditionOk returns a tuple with the DatabaseEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEdition

`func (o *CloudSkuDatabaseTypeAllOf) SetDatabaseEdition(v string)`

SetDatabaseEdition sets DatabaseEdition field to given value.

### HasDatabaseEdition

`func (o *CloudSkuDatabaseTypeAllOf) HasDatabaseEdition() bool`

HasDatabaseEdition returns a boolean if a field has been set.

### GetDatabaseEngine

`func (o *CloudSkuDatabaseTypeAllOf) GetDatabaseEngine() string`

GetDatabaseEngine returns the DatabaseEngine field if non-nil, zero value otherwise.

### GetDatabaseEngineOk

`func (o *CloudSkuDatabaseTypeAllOf) GetDatabaseEngineOk() (*string, bool)`

GetDatabaseEngineOk returns a tuple with the DatabaseEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEngine

`func (o *CloudSkuDatabaseTypeAllOf) SetDatabaseEngine(v string)`

SetDatabaseEngine sets DatabaseEngine field to given value.

### HasDatabaseEngine

`func (o *CloudSkuDatabaseTypeAllOf) HasDatabaseEngine() bool`

HasDatabaseEngine returns a boolean if a field has been set.

### GetLicenseModel

`func (o *CloudSkuDatabaseTypeAllOf) GetLicenseModel() string`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *CloudSkuDatabaseTypeAllOf) GetLicenseModelOk() (*string, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *CloudSkuDatabaseTypeAllOf) SetLicenseModel(v string)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *CloudSkuDatabaseTypeAllOf) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetNetworkPerformance

`func (o *CloudSkuDatabaseTypeAllOf) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *CloudSkuDatabaseTypeAllOf) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *CloudSkuDatabaseTypeAllOf) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *CloudSkuDatabaseTypeAllOf) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


