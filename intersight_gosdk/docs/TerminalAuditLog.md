# TerminalAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time the terminal was closed. If terminal has not closed, value is zero time. | [optional] [readonly] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time the terminal session was opened. | [optional] [readonly] 
**DeviceRegistration** | Pointer to [**AssetDeviceConnectionRelationship**](asset.DeviceConnection.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewTerminalAuditLog

`func NewTerminalAuditLog() *TerminalAuditLog`

NewTerminalAuditLog instantiates a new TerminalAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalAuditLogWithDefaults

`func NewTerminalAuditLogWithDefaults() *TerminalAuditLog`

NewTerminalAuditLogWithDefaults instantiates a new TerminalAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *TerminalAuditLog) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TerminalAuditLog) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TerminalAuditLog) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TerminalAuditLog) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *TerminalAuditLog) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TerminalAuditLog) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TerminalAuditLog) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TerminalAuditLog) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TerminalAuditLog) GetDeviceRegistration() AssetDeviceConnectionRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TerminalAuditLog) GetDeviceRegistrationOk() (*AssetDeviceConnectionRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TerminalAuditLog) SetDeviceRegistration(v AssetDeviceConnectionRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TerminalAuditLog) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetUser

`func (o *TerminalAuditLog) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TerminalAuditLog) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TerminalAuditLog) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *TerminalAuditLog) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


