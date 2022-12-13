# AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerOracleDatabaseServerOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerOracleDatabaseServerOptions"]
**DatabaseId** | Pointer to **string** | A unique name for an Oracle database instance on a specific host. | [optional] 

## Methods

### NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf

`func NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf(classId string, objectType string, ) *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf`

NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf instantiates a new AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOfWithDefaults

`func NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOfWithDefaults() *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf`

NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatabaseId

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetDatabaseId() string`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetDatabaseIdOk() (*string, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) SetDatabaseId(v string)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


