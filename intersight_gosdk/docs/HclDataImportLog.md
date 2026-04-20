# HclDataImportLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.DataImportLog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.DataImportLog"]
**CheckSum** | Pointer to **string** | MD5 Checksum of the HCL Data file. | [optional] 
**HclIncrementalData** | Pointer to [**[]HclServerConfiguration**](HclServerConfiguration.md) |  | [optional] 
**ImportType** | Pointer to **string** | Type of the import whether it is full or a delta. * &#x60;Delta&#x60; - Imports only changes from the HCL tool into Intersight. * &#x60;Full&#x60; - Deletes the current data and does an full import of the data. | [optional] [default to "Delta"]
**InitiatorType** | Pointer to **string** | Type of the initiator whether it is manual or a automated periodic operation by system. The value will be set during DoPost to Manual. * &#x60;Auto&#x60; - Import is auto triggered during service startup or periodic poll. * &#x60;Manual&#x60; - Import is triggered externally by devops using API. | [optional] [readonly] [default to "Auto"]
**LastHclDataModifiedTime** | Pointer to **time.Time** | The timestamp of the last modified record in the HCL tool database. Used to query and get updated records. | [optional] 
**Status** | Pointer to **string** | Status of the import process. * &#x60;None&#x60; - Default import status when no activity is taking place with respect to import. * &#x60;InProgress&#x60; - Data import is in progress. * &#x60;Success&#x60; - Data import is successful. * &#x60;Started&#x60; - Data import process has started. * &#x60;Failed&#x60; - Data import process has failed. * &#x60;NoChange&#x60; - There is no change in the data. | [optional] [default to "None"]
**StatusDetails** | Pointer to **string** | More information on the status. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewHclDataImportLog

`func NewHclDataImportLog(classId string, objectType string, ) *HclDataImportLog`

NewHclDataImportLog instantiates a new HclDataImportLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclDataImportLogWithDefaults

`func NewHclDataImportLogWithDefaults() *HclDataImportLog`

NewHclDataImportLogWithDefaults instantiates a new HclDataImportLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclDataImportLog) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclDataImportLog) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclDataImportLog) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclDataImportLog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclDataImportLog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclDataImportLog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCheckSum

`func (o *HclDataImportLog) GetCheckSum() string`

GetCheckSum returns the CheckSum field if non-nil, zero value otherwise.

### GetCheckSumOk

`func (o *HclDataImportLog) GetCheckSumOk() (*string, bool)`

GetCheckSumOk returns a tuple with the CheckSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSum

`func (o *HclDataImportLog) SetCheckSum(v string)`

SetCheckSum sets CheckSum field to given value.

### HasCheckSum

`func (o *HclDataImportLog) HasCheckSum() bool`

HasCheckSum returns a boolean if a field has been set.

### GetHclIncrementalData

`func (o *HclDataImportLog) GetHclIncrementalData() []HclServerConfiguration`

GetHclIncrementalData returns the HclIncrementalData field if non-nil, zero value otherwise.

### GetHclIncrementalDataOk

`func (o *HclDataImportLog) GetHclIncrementalDataOk() (*[]HclServerConfiguration, bool)`

GetHclIncrementalDataOk returns a tuple with the HclIncrementalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHclIncrementalData

`func (o *HclDataImportLog) SetHclIncrementalData(v []HclServerConfiguration)`

SetHclIncrementalData sets HclIncrementalData field to given value.

### HasHclIncrementalData

`func (o *HclDataImportLog) HasHclIncrementalData() bool`

HasHclIncrementalData returns a boolean if a field has been set.

### SetHclIncrementalDataNil

`func (o *HclDataImportLog) SetHclIncrementalDataNil(b bool)`

 SetHclIncrementalDataNil sets the value for HclIncrementalData to be an explicit nil

### UnsetHclIncrementalData
`func (o *HclDataImportLog) UnsetHclIncrementalData()`

UnsetHclIncrementalData ensures that no value is present for HclIncrementalData, not even an explicit nil
### GetImportType

`func (o *HclDataImportLog) GetImportType() string`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *HclDataImportLog) GetImportTypeOk() (*string, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *HclDataImportLog) SetImportType(v string)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *HclDataImportLog) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### GetInitiatorType

`func (o *HclDataImportLog) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *HclDataImportLog) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *HclDataImportLog) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *HclDataImportLog) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetLastHclDataModifiedTime

`func (o *HclDataImportLog) GetLastHclDataModifiedTime() time.Time`

GetLastHclDataModifiedTime returns the LastHclDataModifiedTime field if non-nil, zero value otherwise.

### GetLastHclDataModifiedTimeOk

`func (o *HclDataImportLog) GetLastHclDataModifiedTimeOk() (*time.Time, bool)`

GetLastHclDataModifiedTimeOk returns a tuple with the LastHclDataModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHclDataModifiedTime

`func (o *HclDataImportLog) SetLastHclDataModifiedTime(v time.Time)`

SetLastHclDataModifiedTime sets LastHclDataModifiedTime field to given value.

### HasLastHclDataModifiedTime

`func (o *HclDataImportLog) HasLastHclDataModifiedTime() bool`

HasLastHclDataModifiedTime returns a boolean if a field has been set.

### GetStatus

`func (o *HclDataImportLog) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HclDataImportLog) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HclDataImportLog) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HclDataImportLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *HclDataImportLog) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *HclDataImportLog) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *HclDataImportLog) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *HclDataImportLog) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetAccount

`func (o *HclDataImportLog) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HclDataImportLog) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HclDataImportLog) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HclDataImportLog) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *HclDataImportLog) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *HclDataImportLog) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


