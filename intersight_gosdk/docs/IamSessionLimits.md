# IamSessionLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleTimeOut** | Pointer to **int64** | The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. | [optional] 
**MaximumLimit** | Pointer to **int64** | The maximum number of sessions allowed in an account. The default value is 128. | [optional] [readonly] 
**PerUserLimit** | Pointer to **int64** | The maximum number of sessions allowed per user. Default value is 32. | [optional] [readonly] 
**SessionTimeOut** | Pointer to **int64** | The session expiry duration in seconds. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 

## Methods

### NewIamSessionLimits

`func NewIamSessionLimits() *IamSessionLimits`

NewIamSessionLimits instantiates a new IamSessionLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionLimitsWithDefaults

`func NewIamSessionLimitsWithDefaults() *IamSessionLimits`

NewIamSessionLimitsWithDefaults instantiates a new IamSessionLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleTimeOut

`func (o *IamSessionLimits) GetIdleTimeOut() int64`

GetIdleTimeOut returns the IdleTimeOut field if non-nil, zero value otherwise.

### GetIdleTimeOutOk

`func (o *IamSessionLimits) GetIdleTimeOutOk() (*int64, bool)`

GetIdleTimeOutOk returns a tuple with the IdleTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeOut

`func (o *IamSessionLimits) SetIdleTimeOut(v int64)`

SetIdleTimeOut sets IdleTimeOut field to given value.

### HasIdleTimeOut

`func (o *IamSessionLimits) HasIdleTimeOut() bool`

HasIdleTimeOut returns a boolean if a field has been set.

### GetMaximumLimit

`func (o *IamSessionLimits) GetMaximumLimit() int64`

GetMaximumLimit returns the MaximumLimit field if non-nil, zero value otherwise.

### GetMaximumLimitOk

`func (o *IamSessionLimits) GetMaximumLimitOk() (*int64, bool)`

GetMaximumLimitOk returns a tuple with the MaximumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLimit

`func (o *IamSessionLimits) SetMaximumLimit(v int64)`

SetMaximumLimit sets MaximumLimit field to given value.

### HasMaximumLimit

`func (o *IamSessionLimits) HasMaximumLimit() bool`

HasMaximumLimit returns a boolean if a field has been set.

### GetPerUserLimit

`func (o *IamSessionLimits) GetPerUserLimit() int64`

GetPerUserLimit returns the PerUserLimit field if non-nil, zero value otherwise.

### GetPerUserLimitOk

`func (o *IamSessionLimits) GetPerUserLimitOk() (*int64, bool)`

GetPerUserLimitOk returns a tuple with the PerUserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUserLimit

`func (o *IamSessionLimits) SetPerUserLimit(v int64)`

SetPerUserLimit sets PerUserLimit field to given value.

### HasPerUserLimit

`func (o *IamSessionLimits) HasPerUserLimit() bool`

HasPerUserLimit returns a boolean if a field has been set.

### GetSessionTimeOut

`func (o *IamSessionLimits) GetSessionTimeOut() int64`

GetSessionTimeOut returns the SessionTimeOut field if non-nil, zero value otherwise.

### GetSessionTimeOutOk

`func (o *IamSessionLimits) GetSessionTimeOutOk() (*int64, bool)`

GetSessionTimeOutOk returns a tuple with the SessionTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeOut

`func (o *IamSessionLimits) SetSessionTimeOut(v int64)`

SetSessionTimeOut sets SessionTimeOut field to given value.

### HasSessionTimeOut

`func (o *IamSessionLimits) HasSessionTimeOut() bool`

HasSessionTimeOut returns a boolean if a field has been set.

### GetAccount

`func (o *IamSessionLimits) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamSessionLimits) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamSessionLimits) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamSessionLimits) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPermission

`func (o *IamSessionLimits) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamSessionLimits) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamSessionLimits) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamSessionLimits) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


