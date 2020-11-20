# IamRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Rule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Rule"]
**IpV6** | Pointer to **bool** | The flag represents if IP addresses in the rule is IPv4 or IPv6. | [optional] 
**RuleType** | Pointer to **string** | The type of the IP address. Currently three types are supported, ie IP, CIDR range and IP range. * &#x60;Ip&#x60; - The IP address rule type is IP. * &#x60;Cidr&#x60; - The IP address rule type is CIDR range. * &#x60;IpRange&#x60; - The IP address rule type is IP range. | [optional] [default to "Ip"]
**RuleValue** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamRuleAllOf

`func NewIamRuleAllOf(classId string, objectType string, ) *IamRuleAllOf`

NewIamRuleAllOf instantiates a new IamRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRuleAllOfWithDefaults

`func NewIamRuleAllOfWithDefaults() *IamRuleAllOf`

NewIamRuleAllOfWithDefaults instantiates a new IamRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamRuleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamRuleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamRuleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamRuleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamRuleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamRuleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpV6

`func (o *IamRuleAllOf) GetIpV6() bool`

GetIpV6 returns the IpV6 field if non-nil, zero value otherwise.

### GetIpV6Ok

`func (o *IamRuleAllOf) GetIpV6Ok() (*bool, bool)`

GetIpV6Ok returns a tuple with the IpV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6

`func (o *IamRuleAllOf) SetIpV6(v bool)`

SetIpV6 sets IpV6 field to given value.

### HasIpV6

`func (o *IamRuleAllOf) HasIpV6() bool`

HasIpV6 returns a boolean if a field has been set.

### GetRuleType

`func (o *IamRuleAllOf) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *IamRuleAllOf) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *IamRuleAllOf) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *IamRuleAllOf) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetRuleValue

`func (o *IamRuleAllOf) GetRuleValue() []string`

GetRuleValue returns the RuleValue field if non-nil, zero value otherwise.

### GetRuleValueOk

`func (o *IamRuleAllOf) GetRuleValueOk() (*[]string, bool)`

GetRuleValueOk returns a tuple with the RuleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleValue

`func (o *IamRuleAllOf) SetRuleValue(v []string)`

SetRuleValue sets RuleValue field to given value.

### HasRuleValue

`func (o *IamRuleAllOf) HasRuleValue() bool`

HasRuleValue returns a boolean if a field has been set.

### SetRuleValueNil

`func (o *IamRuleAllOf) SetRuleValueNil(b bool)`

 SetRuleValueNil sets the value for RuleValue to be an explicit nil

### UnsetRuleValue
`func (o *IamRuleAllOf) UnsetRuleValue()`

UnsetRuleValue ensures that no value is present for RuleValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


