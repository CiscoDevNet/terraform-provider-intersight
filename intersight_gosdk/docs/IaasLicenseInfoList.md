# IaasLicenseInfoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;iaas.LicenseInfo&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]IaasLicenseInfo**](IaasLicenseInfo.md) | The array of &#39;iaas.LicenseInfo&#39; resources matching the request. | [optional] 

## Methods

### NewIaasLicenseInfoList

`func NewIaasLicenseInfoList() *IaasLicenseInfoList`

NewIaasLicenseInfoList instantiates a new IaasLicenseInfoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasLicenseInfoListWithDefaults

`func NewIaasLicenseInfoListWithDefaults() *IaasLicenseInfoList`

NewIaasLicenseInfoListWithDefaults instantiates a new IaasLicenseInfoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IaasLicenseInfoList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IaasLicenseInfoList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IaasLicenseInfoList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IaasLicenseInfoList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *IaasLicenseInfoList) GetResults() []IaasLicenseInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IaasLicenseInfoList) GetResultsOk() (*[]IaasLicenseInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IaasLicenseInfoList) SetResults(v []IaasLicenseInfo)`

SetResults sets Results field to given value.

### HasResults

`func (o *IaasLicenseInfoList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *IaasLicenseInfoList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *IaasLicenseInfoList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


