# RecoveryOnDemandBackupListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;recovery.OnDemandBackup&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]RecoveryOnDemandBackup**](RecoveryOnDemandBackup.md) | The array of &#39;recovery.OnDemandBackup&#39; resources matching the request. | [optional] 

## Methods

### NewRecoveryOnDemandBackupListAllOf

`func NewRecoveryOnDemandBackupListAllOf() *RecoveryOnDemandBackupListAllOf`

NewRecoveryOnDemandBackupListAllOf instantiates a new RecoveryOnDemandBackupListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryOnDemandBackupListAllOfWithDefaults

`func NewRecoveryOnDemandBackupListAllOfWithDefaults() *RecoveryOnDemandBackupListAllOf`

NewRecoveryOnDemandBackupListAllOfWithDefaults instantiates a new RecoveryOnDemandBackupListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RecoveryOnDemandBackupListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecoveryOnDemandBackupListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecoveryOnDemandBackupListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RecoveryOnDemandBackupListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *RecoveryOnDemandBackupListAllOf) GetResults() []RecoveryOnDemandBackup`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RecoveryOnDemandBackupListAllOf) GetResultsOk() (*[]RecoveryOnDemandBackup, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RecoveryOnDemandBackupListAllOf) SetResults(v []RecoveryOnDemandBackup)`

SetResults sets Results field to given value.

### HasResults

`func (o *RecoveryOnDemandBackupListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *RecoveryOnDemandBackupListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *RecoveryOnDemandBackupListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


