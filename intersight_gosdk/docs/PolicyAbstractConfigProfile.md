# PolicyAbstractConfigProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. | [optional] 
**ConfigContext** | Pointer to [**PolicyConfigContext**](policy.ConfigContext.md) |  | [optional] 

## Methods

### NewPolicyAbstractConfigProfile

`func NewPolicyAbstractConfigProfile() *PolicyAbstractConfigProfile`

NewPolicyAbstractConfigProfile instantiates a new PolicyAbstractConfigProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigProfileWithDefaults

`func NewPolicyAbstractConfigProfileWithDefaults() *PolicyAbstractConfigProfile`

NewPolicyAbstractConfigProfileWithDefaults instantiates a new PolicyAbstractConfigProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicyAbstractConfigProfile) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyAbstractConfigProfile) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyAbstractConfigProfile) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PolicyAbstractConfigProfile) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConfigContext

`func (o *PolicyAbstractConfigProfile) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *PolicyAbstractConfigProfile) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *PolicyAbstractConfigProfile) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *PolicyAbstractConfigProfile) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


