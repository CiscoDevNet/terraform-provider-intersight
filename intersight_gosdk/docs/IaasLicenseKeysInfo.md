# IaasLicenseKeysInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Number of licenses available for the UCSD PID (Product ID). | [optional] [readonly] 
**ExpirationDate** | Pointer to **string** | Expiration date for the license. | [optional] [readonly] 
**LicenseId** | Pointer to **string** | UCS Director Unique license ID. | [optional] [readonly] 
**Pid** | Pointer to **string** | PID (Product ID) for UCSD License. | [optional] [readonly] 

## Methods

### NewIaasLicenseKeysInfo

`func NewIaasLicenseKeysInfo() *IaasLicenseKeysInfo`

NewIaasLicenseKeysInfo instantiates a new IaasLicenseKeysInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasLicenseKeysInfoWithDefaults

`func NewIaasLicenseKeysInfoWithDefaults() *IaasLicenseKeysInfo`

NewIaasLicenseKeysInfoWithDefaults instantiates a new IaasLicenseKeysInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IaasLicenseKeysInfo) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IaasLicenseKeysInfo) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IaasLicenseKeysInfo) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *IaasLicenseKeysInfo) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExpirationDate

`func (o *IaasLicenseKeysInfo) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *IaasLicenseKeysInfo) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *IaasLicenseKeysInfo) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *IaasLicenseKeysInfo) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicenseId

`func (o *IaasLicenseKeysInfo) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *IaasLicenseKeysInfo) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *IaasLicenseKeysInfo) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *IaasLicenseKeysInfo) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetPid

`func (o *IaasLicenseKeysInfo) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *IaasLicenseKeysInfo) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *IaasLicenseKeysInfo) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *IaasLicenseKeysInfo) HasPid() bool`

HasPid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


