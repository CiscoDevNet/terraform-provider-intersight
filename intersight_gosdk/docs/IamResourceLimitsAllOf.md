# IamResourceLimitsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerAccountUserLimit** | Pointer to **int64** | The maximum number of users allowed in an account. The default value is 200. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewIamResourceLimitsAllOf

`func NewIamResourceLimitsAllOf() *IamResourceLimitsAllOf`

NewIamResourceLimitsAllOf instantiates a new IamResourceLimitsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourceLimitsAllOfWithDefaults

`func NewIamResourceLimitsAllOfWithDefaults() *IamResourceLimitsAllOf`

NewIamResourceLimitsAllOfWithDefaults instantiates a new IamResourceLimitsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerAccountUserLimit

`func (o *IamResourceLimitsAllOf) GetPerAccountUserLimit() int64`

GetPerAccountUserLimit returns the PerAccountUserLimit field if non-nil, zero value otherwise.

### GetPerAccountUserLimitOk

`func (o *IamResourceLimitsAllOf) GetPerAccountUserLimitOk() (*int64, bool)`

GetPerAccountUserLimitOk returns a tuple with the PerAccountUserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerAccountUserLimit

`func (o *IamResourceLimitsAllOf) SetPerAccountUserLimit(v int64)`

SetPerAccountUserLimit sets PerAccountUserLimit field to given value.

### HasPerAccountUserLimit

`func (o *IamResourceLimitsAllOf) HasPerAccountUserLimit() bool`

HasPerAccountUserLimit returns a boolean if a field has been set.

### GetAccount

`func (o *IamResourceLimitsAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamResourceLimitsAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamResourceLimitsAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamResourceLimitsAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


