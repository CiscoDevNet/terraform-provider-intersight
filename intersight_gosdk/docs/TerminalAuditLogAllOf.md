# TerminalAuditLogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "terminal.AuditLog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "terminal.AuditLog"]
**EndTime** | Pointer to **time.Time** | The time the terminal was closed. If terminal has not closed, value is zero time. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | The time the terminal session was opened. | [optional] [readonly] 
**DeviceRegistration** | Pointer to [**AssetDeviceConnectionRelationship**](AssetDeviceConnectionRelationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewTerminalAuditLogAllOf

`func NewTerminalAuditLogAllOf(classId string, objectType string, ) *TerminalAuditLogAllOf`

NewTerminalAuditLogAllOf instantiates a new TerminalAuditLogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalAuditLogAllOfWithDefaults

`func NewTerminalAuditLogAllOfWithDefaults() *TerminalAuditLogAllOf`

NewTerminalAuditLogAllOfWithDefaults instantiates a new TerminalAuditLogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TerminalAuditLogAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TerminalAuditLogAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TerminalAuditLogAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TerminalAuditLogAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TerminalAuditLogAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TerminalAuditLogAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndTime

`func (o *TerminalAuditLogAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TerminalAuditLogAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TerminalAuditLogAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TerminalAuditLogAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *TerminalAuditLogAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TerminalAuditLogAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TerminalAuditLogAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TerminalAuditLogAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TerminalAuditLogAllOf) GetDeviceRegistration() AssetDeviceConnectionRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TerminalAuditLogAllOf) GetDeviceRegistrationOk() (*AssetDeviceConnectionRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TerminalAuditLogAllOf) SetDeviceRegistration(v AssetDeviceConnectionRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TerminalAuditLogAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetUser

`func (o *TerminalAuditLogAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TerminalAuditLogAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TerminalAuditLogAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *TerminalAuditLogAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


