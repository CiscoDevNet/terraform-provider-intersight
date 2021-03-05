# CloudSkuDatabaseType

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

### NewCloudSkuDatabaseType

`func NewCloudSkuDatabaseType(classId string, objectType string, ) *CloudSkuDatabaseType`

NewCloudSkuDatabaseType instantiates a new CloudSkuDatabaseType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSkuDatabaseTypeWithDefaults

`func NewCloudSkuDatabaseTypeWithDefaults() *CloudSkuDatabaseType`

NewCloudSkuDatabaseTypeWithDefaults instantiates a new CloudSkuDatabaseType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudSkuDatabaseType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudSkuDatabaseType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudSkuDatabaseType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudSkuDatabaseType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudSkuDatabaseType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudSkuDatabaseType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatabaseEdition

`func (o *CloudSkuDatabaseType) GetDatabaseEdition() string`

GetDatabaseEdition returns the DatabaseEdition field if non-nil, zero value otherwise.

### GetDatabaseEditionOk

`func (o *CloudSkuDatabaseType) GetDatabaseEditionOk() (*string, bool)`

GetDatabaseEditionOk returns a tuple with the DatabaseEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEdition

`func (o *CloudSkuDatabaseType) SetDatabaseEdition(v string)`

SetDatabaseEdition sets DatabaseEdition field to given value.

### HasDatabaseEdition

`func (o *CloudSkuDatabaseType) HasDatabaseEdition() bool`

HasDatabaseEdition returns a boolean if a field has been set.

### GetDatabaseEngine

`func (o *CloudSkuDatabaseType) GetDatabaseEngine() string`

GetDatabaseEngine returns the DatabaseEngine field if non-nil, zero value otherwise.

### GetDatabaseEngineOk

`func (o *CloudSkuDatabaseType) GetDatabaseEngineOk() (*string, bool)`

GetDatabaseEngineOk returns a tuple with the DatabaseEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEngine

`func (o *CloudSkuDatabaseType) SetDatabaseEngine(v string)`

SetDatabaseEngine sets DatabaseEngine field to given value.

### HasDatabaseEngine

`func (o *CloudSkuDatabaseType) HasDatabaseEngine() bool`

HasDatabaseEngine returns a boolean if a field has been set.

### GetLicenseModel

`func (o *CloudSkuDatabaseType) GetLicenseModel() string`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *CloudSkuDatabaseType) GetLicenseModelOk() (*string, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *CloudSkuDatabaseType) SetLicenseModel(v string)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *CloudSkuDatabaseType) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetNetworkPerformance

`func (o *CloudSkuDatabaseType) GetNetworkPerformance() string`

GetNetworkPerformance returns the NetworkPerformance field if non-nil, zero value otherwise.

### GetNetworkPerformanceOk

`func (o *CloudSkuDatabaseType) GetNetworkPerformanceOk() (*string, bool)`

GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPerformance

`func (o *CloudSkuDatabaseType) SetNetworkPerformance(v string)`

SetNetworkPerformance sets NetworkPerformance field to given value.

### HasNetworkPerformance

`func (o *CloudSkuDatabaseType) HasNetworkPerformance() bool`

HasNetworkPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


