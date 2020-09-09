# ApplianceBackupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElapsedTime** | Pointer to **int64** | Elapsed time in seconds since the backup process has started. | [optional] [readonly] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | End date and time of the backup process. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**Password** | Pointer to **string** | Password to authenticate the fileserver. | [optional] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Start date and time of the backup process. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the backup managed object. * &#x60;Started&#x60; - Backup or restore process has started. * &#x60;Created&#x60; - Backup or restore is in created state. * &#x60;Failed&#x60; - Backup or restore process has failed. * &#x60;Completed&#x60; - Backup or restore process has completed. * &#x60;Copied&#x60; - Backup file has been copied. | [optional] [readonly] [default to "Started"]
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewApplianceBackupAllOf

`func NewApplianceBackupAllOf() *ApplianceBackupAllOf`

NewApplianceBackupAllOf instantiates a new ApplianceBackupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceBackupAllOfWithDefaults

`func NewApplianceBackupAllOfWithDefaults() *ApplianceBackupAllOf`

NewApplianceBackupAllOfWithDefaults instantiates a new ApplianceBackupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsedTime

`func (o *ApplianceBackupAllOf) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ApplianceBackupAllOf) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ApplianceBackupAllOf) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ApplianceBackupAllOf) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceBackupAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceBackupAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceBackupAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceBackupAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *ApplianceBackupAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ApplianceBackupAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ApplianceBackupAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ApplianceBackupAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMessages

`func (o *ApplianceBackupAllOf) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ApplianceBackupAllOf) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ApplianceBackupAllOf) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ApplianceBackupAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetPassword

`func (o *ApplianceBackupAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApplianceBackupAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApplianceBackupAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApplianceBackupAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStartTime

`func (o *ApplianceBackupAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceBackupAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceBackupAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceBackupAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceBackupAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceBackupAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceBackupAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceBackupAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceBackupAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceBackupAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceBackupAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceBackupAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


