# RecoveryBackupProfileListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;recovery.BackupProfile&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]RecoveryBackupProfile**](RecoveryBackupProfile.md) | The array of &#39;recovery.BackupProfile&#39; resources matching the request. | [optional] 

## Methods

### NewRecoveryBackupProfileListAllOf

`func NewRecoveryBackupProfileListAllOf() *RecoveryBackupProfileListAllOf`

NewRecoveryBackupProfileListAllOf instantiates a new RecoveryBackupProfileListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupProfileListAllOfWithDefaults

`func NewRecoveryBackupProfileListAllOfWithDefaults() *RecoveryBackupProfileListAllOf`

NewRecoveryBackupProfileListAllOfWithDefaults instantiates a new RecoveryBackupProfileListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RecoveryBackupProfileListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecoveryBackupProfileListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecoveryBackupProfileListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RecoveryBackupProfileListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *RecoveryBackupProfileListAllOf) GetResults() []RecoveryBackupProfile`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RecoveryBackupProfileListAllOf) GetResultsOk() (*[]RecoveryBackupProfile, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RecoveryBackupProfileListAllOf) SetResults(v []RecoveryBackupProfile)`

SetResults sets Results field to given value.

### HasResults

`func (o *RecoveryBackupProfileListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *RecoveryBackupProfileListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *RecoveryBackupProfileListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


