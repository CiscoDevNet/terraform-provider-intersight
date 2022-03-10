# AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerMicrosoftSqlServerOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerMicrosoftSqlServerOptions"]
**BrowserServicePort** | Pointer to **int64** | Port that Microsoft SQL Server Browser listens for incoming requests for SQL Server resources and provides information about SQL Server instances that are installed on the computer. When this port is specified, Database will be communicated through the Browser Service with this port instead of default SQLServer port. | [optional] 
**FullDomainName** | Pointer to **string** | Active Directory domain, if required for this account. | [optional] 

## Methods

### NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf

`func NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf(classId string, objectType string, ) *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf`

NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf instantiates a new AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOfWithDefaults

`func NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOfWithDefaults() *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf`

NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBrowserServicePort

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetBrowserServicePort() int64`

GetBrowserServicePort returns the BrowserServicePort field if non-nil, zero value otherwise.

### GetBrowserServicePortOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetBrowserServicePortOk() (*int64, bool)`

GetBrowserServicePortOk returns a tuple with the BrowserServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserServicePort

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetBrowserServicePort(v int64)`

SetBrowserServicePort sets BrowserServicePort field to given value.

### HasBrowserServicePort

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) HasBrowserServicePort() bool`

HasBrowserServicePort returns a boolean if a field has been set.

### GetFullDomainName

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetFullDomainName() string`

GetFullDomainName returns the FullDomainName field if non-nil, zero value otherwise.

### GetFullDomainNameOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) GetFullDomainNameOk() (*string, bool)`

GetFullDomainNameOk returns a tuple with the FullDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDomainName

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) SetFullDomainName(v string)`

SetFullDomainName sets FullDomainName field to given value.

### HasFullDomainName

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptionsAllOf) HasFullDomainName() bool`

HasFullDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


