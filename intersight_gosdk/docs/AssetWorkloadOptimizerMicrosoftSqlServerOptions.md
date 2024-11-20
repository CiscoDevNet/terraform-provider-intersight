# AssetWorkloadOptimizerMicrosoftSqlServerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerMicrosoftSqlServerOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerMicrosoftSqlServerOptions"]
**BrowserServicePort** | Pointer to **int64** | Port that Microsoft SQL Server Browser listens for incoming requests for SQL Server resources and provides information about SQL Server instances that are installed on the computer. When this port is specified, Database will be communicated through the Browser Service with this port instead of default SQLServer port. | [optional] 
**DiscoveryPath** | Pointer to **string** | Discovery path to define if its scope target entities or hostname or IP address. | [optional] [default to "targetEntities"]
**FullDomainName** | Pointer to **string** | Active Directory domain, if required for this account. | [optional] 

## Methods

### NewAssetWorkloadOptimizerMicrosoftSqlServerOptions

`func NewAssetWorkloadOptimizerMicrosoftSqlServerOptions(classId string, objectType string, ) *AssetWorkloadOptimizerMicrosoftSqlServerOptions`

NewAssetWorkloadOptimizerMicrosoftSqlServerOptions instantiates a new AssetWorkloadOptimizerMicrosoftSqlServerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsWithDefaults

`func NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsWithDefaults() *AssetWorkloadOptimizerMicrosoftSqlServerOptions`

NewAssetWorkloadOptimizerMicrosoftSqlServerOptionsWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftSqlServerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBrowserServicePort

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetBrowserServicePort() int64`

GetBrowserServicePort returns the BrowserServicePort field if non-nil, zero value otherwise.

### GetBrowserServicePortOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetBrowserServicePortOk() (*int64, bool)`

GetBrowserServicePortOk returns a tuple with the BrowserServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserServicePort

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) SetBrowserServicePort(v int64)`

SetBrowserServicePort sets BrowserServicePort field to given value.

### HasBrowserServicePort

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) HasBrowserServicePort() bool`

HasBrowserServicePort returns a boolean if a field has been set.

### GetDiscoveryPath

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetDiscoveryPath() string`

GetDiscoveryPath returns the DiscoveryPath field if non-nil, zero value otherwise.

### GetDiscoveryPathOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetDiscoveryPathOk() (*string, bool)`

GetDiscoveryPathOk returns a tuple with the DiscoveryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryPath

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) SetDiscoveryPath(v string)`

SetDiscoveryPath sets DiscoveryPath field to given value.

### HasDiscoveryPath

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) HasDiscoveryPath() bool`

HasDiscoveryPath returns a boolean if a field has been set.

### GetFullDomainName

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetFullDomainName() string`

GetFullDomainName returns the FullDomainName field if non-nil, zero value otherwise.

### GetFullDomainNameOk

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) GetFullDomainNameOk() (*string, bool)`

GetFullDomainNameOk returns a tuple with the FullDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDomainName

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) SetFullDomainName(v string)`

SetFullDomainName sets FullDomainName field to given value.

### HasFullDomainName

`func (o *AssetWorkloadOptimizerMicrosoftSqlServerOptions) HasFullDomainName() bool`

HasFullDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


