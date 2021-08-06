# IamSessionLimits

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

### NewIamSessionLimits

`func NewIamSessionLimits(classId string, objectType string, ) *IamSessionLimits`

NewIamSessionLimits instantiates a new IamSessionLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionLimitsWithDefaults

`func NewIamSessionLimitsWithDefaults() *IamSessionLimits`

NewIamSessionLimitsWithDefaults instantiates a new IamSessionLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSessionLimits) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSessionLimits) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSessionLimits) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSessionLimits) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSessionLimits) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSessionLimits) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


