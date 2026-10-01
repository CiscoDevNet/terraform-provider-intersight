# IamRoutingRuleReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.RoutingRuleReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.RoutingRuleReference"]
**RuleId** | Pointer to **string** | Stable external identity-provider routing-rule identifier. | [optional] 
**RuleName** | Pointer to **string** | Display name of the external identity-provider routing rule. | [optional] 

## Methods

### NewIamRoutingRuleReference

`func NewIamRoutingRuleReference(classId string, objectType string, ) *IamRoutingRuleReference`

NewIamRoutingRuleReference instantiates a new IamRoutingRuleReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoutingRuleReferenceWithDefaults

`func NewIamRoutingRuleReferenceWithDefaults() *IamRoutingRuleReference`

NewIamRoutingRuleReferenceWithDefaults instantiates a new IamRoutingRuleReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamRoutingRuleReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamRoutingRuleReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamRoutingRuleReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamRoutingRuleReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamRoutingRuleReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamRoutingRuleReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRuleId

`func (o *IamRoutingRuleReference) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *IamRoutingRuleReference) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *IamRoutingRuleReference) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *IamRoutingRuleReference) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *IamRoutingRuleReference) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *IamRoutingRuleReference) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *IamRoutingRuleReference) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *IamRoutingRuleReference) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


