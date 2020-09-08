# IamResourceLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerAccountUserLimit** | Pointer to **int64** | The maximum number of users allowed in an account. The default value is 200. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewIamResourceLimits

`func NewIamResourceLimits() *IamResourceLimits`

NewIamResourceLimits instantiates a new IamResourceLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourceLimitsWithDefaults

`func NewIamResourceLimitsWithDefaults() *IamResourceLimits`

NewIamResourceLimitsWithDefaults instantiates a new IamResourceLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerAccountUserLimit

`func (o *IamResourceLimits) GetPerAccountUserLimit() int64`

GetPerAccountUserLimit returns the PerAccountUserLimit field if non-nil, zero value otherwise.

### GetPerAccountUserLimitOk

`func (o *IamResourceLimits) GetPerAccountUserLimitOk() (*int64, bool)`

GetPerAccountUserLimitOk returns a tuple with the PerAccountUserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerAccountUserLimit

`func (o *IamResourceLimits) SetPerAccountUserLimit(v int64)`

SetPerAccountUserLimit sets PerAccountUserLimit field to given value.

### HasPerAccountUserLimit

`func (o *IamResourceLimits) HasPerAccountUserLimit() bool`

HasPerAccountUserLimit returns a boolean if a field has been set.

### GetAccount

`func (o *IamResourceLimits) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamResourceLimits) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamResourceLimits) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamResourceLimits) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


