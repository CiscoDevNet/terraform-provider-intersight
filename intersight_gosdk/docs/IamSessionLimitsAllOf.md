# IamSessionLimitsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SessionLimits"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SessionLimits"]
**IdleTimeOut** | Pointer to **int64** | The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. The minimum value is 300 seconds and the maximum value is 18000 seconds (5 hours). The system default value is 1800 seconds. | [optional] [default to 1800]
**MaximumLimit** | Pointer to **int64** | The maximum number of sessions allowed in an account or permission. The default value is 128. | [optional] [default to 128]
**PerUserLimit** | Pointer to **int64** | The maximum number of sessions allowed per user. Default value is 32. | [optional] [default to 32]
**SessionTimeOut** | Pointer to **int64** | The session expiry duration in seconds. The minimum value is 350 seconds and the maximum value is 31536000 seconds (1 year). The system default value is 57600 seconds. | [optional] [default to 57600]
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 

## Methods

### NewIamSessionLimitsAllOf

`func NewIamSessionLimitsAllOf(classId string, objectType string, ) *IamSessionLimitsAllOf`

NewIamSessionLimitsAllOf instantiates a new IamSessionLimitsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionLimitsAllOfWithDefaults

`func NewIamSessionLimitsAllOfWithDefaults() *IamSessionLimitsAllOf`

NewIamSessionLimitsAllOfWithDefaults instantiates a new IamSessionLimitsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSessionLimitsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSessionLimitsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSessionLimitsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSessionLimitsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSessionLimitsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSessionLimitsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdleTimeOut

`func (o *IamSessionLimitsAllOf) GetIdleTimeOut() int64`

GetIdleTimeOut returns the IdleTimeOut field if non-nil, zero value otherwise.

### GetIdleTimeOutOk

`func (o *IamSessionLimitsAllOf) GetIdleTimeOutOk() (*int64, bool)`

GetIdleTimeOutOk returns a tuple with the IdleTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeOut

`func (o *IamSessionLimitsAllOf) SetIdleTimeOut(v int64)`

SetIdleTimeOut sets IdleTimeOut field to given value.

### HasIdleTimeOut

`func (o *IamSessionLimitsAllOf) HasIdleTimeOut() bool`

HasIdleTimeOut returns a boolean if a field has been set.

### GetMaximumLimit

`func (o *IamSessionLimitsAllOf) GetMaximumLimit() int64`

GetMaximumLimit returns the MaximumLimit field if non-nil, zero value otherwise.

### GetMaximumLimitOk

`func (o *IamSessionLimitsAllOf) GetMaximumLimitOk() (*int64, bool)`

GetMaximumLimitOk returns a tuple with the MaximumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLimit

`func (o *IamSessionLimitsAllOf) SetMaximumLimit(v int64)`

SetMaximumLimit sets MaximumLimit field to given value.

### HasMaximumLimit

`func (o *IamSessionLimitsAllOf) HasMaximumLimit() bool`

HasMaximumLimit returns a boolean if a field has been set.

### GetPerUserLimit

`func (o *IamSessionLimitsAllOf) GetPerUserLimit() int64`

GetPerUserLimit returns the PerUserLimit field if non-nil, zero value otherwise.

### GetPerUserLimitOk

`func (o *IamSessionLimitsAllOf) GetPerUserLimitOk() (*int64, bool)`

GetPerUserLimitOk returns a tuple with the PerUserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUserLimit

`func (o *IamSessionLimitsAllOf) SetPerUserLimit(v int64)`

SetPerUserLimit sets PerUserLimit field to given value.

### HasPerUserLimit

`func (o *IamSessionLimitsAllOf) HasPerUserLimit() bool`

HasPerUserLimit returns a boolean if a field has been set.

### GetSessionTimeOut

`func (o *IamSessionLimitsAllOf) GetSessionTimeOut() int64`

GetSessionTimeOut returns the SessionTimeOut field if non-nil, zero value otherwise.

### GetSessionTimeOutOk

`func (o *IamSessionLimitsAllOf) GetSessionTimeOutOk() (*int64, bool)`

GetSessionTimeOutOk returns a tuple with the SessionTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeOut

`func (o *IamSessionLimitsAllOf) SetSessionTimeOut(v int64)`

SetSessionTimeOut sets SessionTimeOut field to given value.

### HasSessionTimeOut

`func (o *IamSessionLimitsAllOf) HasSessionTimeOut() bool`

HasSessionTimeOut returns a boolean if a field has been set.

### GetAccount

`func (o *IamSessionLimitsAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamSessionLimitsAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamSessionLimitsAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamSessionLimitsAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPermission

`func (o *IamSessionLimitsAllOf) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamSessionLimitsAllOf) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamSessionLimitsAllOf) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamSessionLimitsAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


