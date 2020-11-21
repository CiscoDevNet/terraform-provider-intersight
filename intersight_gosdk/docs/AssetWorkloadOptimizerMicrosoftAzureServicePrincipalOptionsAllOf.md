# AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"]
**AzureCloudType** | Pointer to **string** | Azure has different endpoints for managing Germany subscriptions. Azure cloud type helps in differentiating between regular subscriptions and Germany subscriptions to manage the Azure services for workload optimization. Documentation for germany cloud [link](https://docs.microsoft.com/en-us/azure/germany/germany-manage-subscriptions). * &#x60;Global&#x60; - Global cloud type for Azure subscription. * &#x60;Germany&#x60; - Germany cloud type for Azure subscription. | [optional] [default to "Global"]
**TenantId** | Pointer to **string** | Id of the tenant used while authenticating the managed target. | [optional] 

## Methods

### NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf

`func NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf(classId string, objectType string, ) *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf`

NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf instantiates a new AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOfWithDefaults

`func NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOfWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf`

NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAzureCloudType

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetAzureCloudType() string`

GetAzureCloudType returns the AzureCloudType field if non-nil, zero value otherwise.

### GetAzureCloudTypeOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetAzureCloudTypeOk() (*string, bool)`

GetAzureCloudTypeOk returns a tuple with the AzureCloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCloudType

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetAzureCloudType(v string)`

SetAzureCloudType sets AzureCloudType field to given value.

### HasAzureCloudType

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) HasAzureCloudType() bool`

HasAzureCloudType returns a boolean if a field has been set.

### GetTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


