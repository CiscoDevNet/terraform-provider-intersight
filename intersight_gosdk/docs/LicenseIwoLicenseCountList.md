# LicenseIwoLicenseCountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;license.IwoLicenseCount&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]LicenseIwoLicenseCount**](LicenseIwoLicenseCount.md) | The array of &#39;license.IwoLicenseCount&#39; resources matching the request. | [optional] 

## Methods

### NewLicenseIwoLicenseCountList

`func NewLicenseIwoLicenseCountList() *LicenseIwoLicenseCountList`

NewLicenseIwoLicenseCountList instantiates a new LicenseIwoLicenseCountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIwoLicenseCountListWithDefaults

`func NewLicenseIwoLicenseCountListWithDefaults() *LicenseIwoLicenseCountList`

NewLicenseIwoLicenseCountListWithDefaults instantiates a new LicenseIwoLicenseCountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *LicenseIwoLicenseCountList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LicenseIwoLicenseCountList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LicenseIwoLicenseCountList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *LicenseIwoLicenseCountList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *LicenseIwoLicenseCountList) GetResults() []LicenseIwoLicenseCount`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LicenseIwoLicenseCountList) GetResultsOk() (*[]LicenseIwoLicenseCount, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LicenseIwoLicenseCountList) SetResults(v []LicenseIwoLicenseCount)`

SetResults sets Results field to given value.

### HasResults

`func (o *LicenseIwoLicenseCountList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *LicenseIwoLicenseCountList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *LicenseIwoLicenseCountList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


