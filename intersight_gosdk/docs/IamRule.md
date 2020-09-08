# IamRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpV6** | Pointer to **bool** | The flag represents if IP addresses in the rule is IPv4 or IPv6. | [optional] 
**RuleType** | Pointer to **string** | The type of the IP address. Currently three types are supported, ie IP, CIDR range and IP range. * &#x60;Ip&#x60; - The IP address rule type is IP. * &#x60;Cidr&#x60; - The IP address rule type is CIDR range. * &#x60;IpRange&#x60; - The IP address rule type is IP range. | [optional] [default to "Ip"]
**RuleValue** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamRule

`func NewIamRule() *IamRule`

NewIamRule instantiates a new IamRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRuleWithDefaults

`func NewIamRuleWithDefaults() *IamRule`

NewIamRuleWithDefaults instantiates a new IamRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpV6

`func (o *IamRule) GetIpV6() bool`

GetIpV6 returns the IpV6 field if non-nil, zero value otherwise.

### GetIpV6Ok

`func (o *IamRule) GetIpV6Ok() (*bool, bool)`

GetIpV6Ok returns a tuple with the IpV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6

`func (o *IamRule) SetIpV6(v bool)`

SetIpV6 sets IpV6 field to given value.

### HasIpV6

`func (o *IamRule) HasIpV6() bool`

HasIpV6 returns a boolean if a field has been set.

### GetRuleType

`func (o *IamRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *IamRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *IamRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *IamRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetRuleValue

`func (o *IamRule) GetRuleValue() []string`

GetRuleValue returns the RuleValue field if non-nil, zero value otherwise.

### GetRuleValueOk

`func (o *IamRule) GetRuleValueOk() (*[]string, bool)`

GetRuleValueOk returns a tuple with the RuleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleValue

`func (o *IamRule) SetRuleValue(v []string)`

SetRuleValue sets RuleValue field to given value.

### HasRuleValue

`func (o *IamRule) HasRuleValue() bool`

HasRuleValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


