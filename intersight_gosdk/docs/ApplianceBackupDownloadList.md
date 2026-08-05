# ApplianceBackupDownloadList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;appliance.BackupDownload&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]ApplianceBackupDownload**](ApplianceBackupDownload.md) | The array of &#39;appliance.BackupDownload&#39; resources matching the request. | [optional] 

## Methods

### NewApplianceBackupDownloadList

`func NewApplianceBackupDownloadList() *ApplianceBackupDownloadList`

NewApplianceBackupDownloadList instantiates a new ApplianceBackupDownloadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupDownloadListWithDefaults

`func NewApplianceBackupDownloadListWithDefaults() *ApplianceBackupDownloadList`

NewApplianceBackupDownloadListWithDefaults instantiates a new ApplianceBackupDownloadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ApplianceBackupDownloadList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApplianceBackupDownloadList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApplianceBackupDownloadList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ApplianceBackupDownloadList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *ApplianceBackupDownloadList) GetResults() []ApplianceBackupDownload`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApplianceBackupDownloadList) GetResultsOk() (*[]ApplianceBackupDownload, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApplianceBackupDownloadList) SetResults(v []ApplianceBackupDownload)`

SetResults sets Results field to given value.

### HasResults

`func (o *ApplianceBackupDownloadList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *ApplianceBackupDownloadList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *ApplianceBackupDownloadList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


