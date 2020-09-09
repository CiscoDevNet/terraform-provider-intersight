# IaasLicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseExpirationDate** | Pointer to **string** | UCS Director license expiration date. | [optional] [readonly] 
**LicenseKeysInfo** | Pointer to [**[]IaasLicenseKeysInfo**](iaas.LicenseKeysInfo.md) |  | [optional] 
**LicenseType** | Pointer to **string** | License type of UCSD whether it is EVAL/Permanent/Subscription.. | [optional] [readonly] 
**LicenseUtilizationInfo** | Pointer to [**[]IaasLicenseUtilizationInfo**](iaas.LicenseUtilizationInfo.md) |  | [optional] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 

## Methods

### NewIaasLicenseInfo

`func NewIaasLicenseInfo() *IaasLicenseInfo`

NewIaasLicenseInfo instantiates a new IaasLicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasLicenseInfoWithDefaults

`func NewIaasLicenseInfoWithDefaults() *IaasLicenseInfo`

NewIaasLicenseInfoWithDefaults instantiates a new IaasLicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseExpirationDate

`func (o *IaasLicenseInfo) GetLicenseExpirationDate() string`

GetLicenseExpirationDate returns the LicenseExpirationDate field if non-nil, zero value otherwise.

### GetLicenseExpirationDateOk

`func (o *IaasLicenseInfo) GetLicenseExpirationDateOk() (*string, bool)`

GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationDate

`func (o *IaasLicenseInfo) SetLicenseExpirationDate(v string)`

SetLicenseExpirationDate sets LicenseExpirationDate field to given value.

### HasLicenseExpirationDate

`func (o *IaasLicenseInfo) HasLicenseExpirationDate() bool`

HasLicenseExpirationDate returns a boolean if a field has been set.

### GetLicenseKeysInfo

`func (o *IaasLicenseInfo) GetLicenseKeysInfo() []IaasLicenseKeysInfo`

GetLicenseKeysInfo returns the LicenseKeysInfo field if non-nil, zero value otherwise.

### GetLicenseKeysInfoOk

`func (o *IaasLicenseInfo) GetLicenseKeysInfoOk() (*[]IaasLicenseKeysInfo, bool)`

GetLicenseKeysInfoOk returns a tuple with the LicenseKeysInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKeysInfo

`func (o *IaasLicenseInfo) SetLicenseKeysInfo(v []IaasLicenseKeysInfo)`

SetLicenseKeysInfo sets LicenseKeysInfo field to given value.

### HasLicenseKeysInfo

`func (o *IaasLicenseInfo) HasLicenseKeysInfo() bool`

HasLicenseKeysInfo returns a boolean if a field has been set.

### GetLicenseType

`func (o *IaasLicenseInfo) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *IaasLicenseInfo) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *IaasLicenseInfo) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *IaasLicenseInfo) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseUtilizationInfo

`func (o *IaasLicenseInfo) GetLicenseUtilizationInfo() []IaasLicenseUtilizationInfo`

GetLicenseUtilizationInfo returns the LicenseUtilizationInfo field if non-nil, zero value otherwise.

### GetLicenseUtilizationInfoOk

`func (o *IaasLicenseInfo) GetLicenseUtilizationInfoOk() (*[]IaasLicenseUtilizationInfo, bool)`

GetLicenseUtilizationInfoOk returns a tuple with the LicenseUtilizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUtilizationInfo

`func (o *IaasLicenseInfo) SetLicenseUtilizationInfo(v []IaasLicenseUtilizationInfo)`

SetLicenseUtilizationInfo sets LicenseUtilizationInfo field to given value.

### HasLicenseUtilizationInfo

`func (o *IaasLicenseInfo) HasLicenseUtilizationInfo() bool`

HasLicenseUtilizationInfo returns a boolean if a field has been set.

### GetGuid

`func (o *IaasLicenseInfo) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasLicenseInfo) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasLicenseInfo) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasLicenseInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


