# IaasLicenseUtilizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualUsed** | Pointer to **int64** | Number of licenses actually used for this feature. | [optional] [readonly] 
**Label** | Pointer to **string** | License label of UCSD license. | [optional] [readonly] 
**LicensedLimit** | Pointer to **string** | License limit for this license feature. | [optional] [readonly] 
**Sku** | Pointer to **string** | SKU for the UCSD license. | [optional] [readonly] 

## Methods

### NewIaasLicenseUtilizationInfo

`func NewIaasLicenseUtilizationInfo() *IaasLicenseUtilizationInfo`

NewIaasLicenseUtilizationInfo instantiates a new IaasLicenseUtilizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasLicenseUtilizationInfoWithDefaults

`func NewIaasLicenseUtilizationInfoWithDefaults() *IaasLicenseUtilizationInfo`

NewIaasLicenseUtilizationInfoWithDefaults instantiates a new IaasLicenseUtilizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualUsed

`func (o *IaasLicenseUtilizationInfo) GetActualUsed() int64`

GetActualUsed returns the ActualUsed field if non-nil, zero value otherwise.

### GetActualUsedOk

`func (o *IaasLicenseUtilizationInfo) GetActualUsedOk() (*int64, bool)`

GetActualUsedOk returns a tuple with the ActualUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualUsed

`func (o *IaasLicenseUtilizationInfo) SetActualUsed(v int64)`

SetActualUsed sets ActualUsed field to given value.

### HasActualUsed

`func (o *IaasLicenseUtilizationInfo) HasActualUsed() bool`

HasActualUsed returns a boolean if a field has been set.

### GetLabel

`func (o *IaasLicenseUtilizationInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IaasLicenseUtilizationInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IaasLicenseUtilizationInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IaasLicenseUtilizationInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicensedLimit

`func (o *IaasLicenseUtilizationInfo) GetLicensedLimit() string`

GetLicensedLimit returns the LicensedLimit field if non-nil, zero value otherwise.

### GetLicensedLimitOk

`func (o *IaasLicenseUtilizationInfo) GetLicensedLimitOk() (*string, bool)`

GetLicensedLimitOk returns a tuple with the LicensedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedLimit

`func (o *IaasLicenseUtilizationInfo) SetLicensedLimit(v string)`

SetLicensedLimit sets LicensedLimit field to given value.

### HasLicensedLimit

`func (o *IaasLicenseUtilizationInfo) HasLicensedLimit() bool`

HasLicensedLimit returns a boolean if a field has been set.

### GetSku

`func (o *IaasLicenseUtilizationInfo) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *IaasLicenseUtilizationInfo) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *IaasLicenseUtilizationInfo) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *IaasLicenseUtilizationInfo) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


