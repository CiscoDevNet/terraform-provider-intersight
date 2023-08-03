# AssetWorkloadOptimizerMicrosoftAzureBillingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerMicrosoftAzureBillingOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerMicrosoftAzureBillingOptions"]
**BillingAccountId** | Pointer to **string** | The Microsoft Azure Billing Account ID. | [optional] 
**CostExportName** | Pointer to **string** | Name of the Cost Export Data that exports cost management data. | [optional] 
**TenantId** | Pointer to **string** | Id of the tenant used while authenticating the managed target. | [optional] 
**UseCostExport** | Pointer to **bool** | For larger topologies, use cost export to fetch billed cost data. | [optional] [default to false]

## Methods

### NewAssetWorkloadOptimizerMicrosoftAzureBillingOptions

`func NewAssetWorkloadOptimizerMicrosoftAzureBillingOptions(classId string, objectType string, ) *AssetWorkloadOptimizerMicrosoftAzureBillingOptions`

NewAssetWorkloadOptimizerMicrosoftAzureBillingOptions instantiates a new AssetWorkloadOptimizerMicrosoftAzureBillingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerMicrosoftAzureBillingOptionsWithDefaults

`func NewAssetWorkloadOptimizerMicrosoftAzureBillingOptionsWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureBillingOptions`

NewAssetWorkloadOptimizerMicrosoftAzureBillingOptionsWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureBillingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingAccountId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetCostExportName

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetCostExportName() string`

GetCostExportName returns the CostExportName field if non-nil, zero value otherwise.

### GetCostExportNameOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetCostExportNameOk() (*string, bool)`

GetCostExportNameOk returns a tuple with the CostExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostExportName

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) SetCostExportName(v string)`

SetCostExportName sets CostExportName field to given value.

### HasCostExportName

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) HasCostExportName() bool`

HasCostExportName returns a boolean if a field has been set.

### GetTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUseCostExport

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetUseCostExport() bool`

GetUseCostExport returns the UseCostExport field if non-nil, zero value otherwise.

### GetUseCostExportOk

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) GetUseCostExportOk() (*bool, bool)`

GetUseCostExportOk returns a tuple with the UseCostExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCostExport

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) SetUseCostExport(v bool)`

SetUseCostExport sets UseCostExport field to given value.

### HasUseCostExport

`func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptions) HasUseCostExport() bool`

HasUseCostExport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


