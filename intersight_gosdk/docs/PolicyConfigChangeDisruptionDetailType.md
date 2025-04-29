# PolicyConfigChangeDisruptionDetailType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ConfigChangeDisruptionDetailType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ConfigChangeDisruptionDetailType"]
**Disruptions** | Pointer to **[]string** |  | [optional] 
**IsOnlyRequiredByOtherPolicies** | Pointer to **bool** | The current policy has to be redeployed only because there are other policy changes that require this. | [optional] 
**PolicyName** | Pointer to **string** | Name of the policy that, when modified, causes the disruption. | [optional] 
**PolicyPendingAction** | Pointer to **string** | Name of the action which is pending on this policy. Example, if policy is not yet activated we mark this field as not-activated. Currently we support two actions, not-deployed and not-activated. | [optional] 
**RequiredByPolicies** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyConfigChangeDisruptionDetailType

`func NewPolicyConfigChangeDisruptionDetailType(classId string, objectType string, ) *PolicyConfigChangeDisruptionDetailType`

NewPolicyConfigChangeDisruptionDetailType instantiates a new PolicyConfigChangeDisruptionDetailType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigChangeDisruptionDetailTypeWithDefaults

`func NewPolicyConfigChangeDisruptionDetailTypeWithDefaults() *PolicyConfigChangeDisruptionDetailType`

NewPolicyConfigChangeDisruptionDetailTypeWithDefaults instantiates a new PolicyConfigChangeDisruptionDetailType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyConfigChangeDisruptionDetailType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyConfigChangeDisruptionDetailType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyConfigChangeDisruptionDetailType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyConfigChangeDisruptionDetailType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyConfigChangeDisruptionDetailType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyConfigChangeDisruptionDetailType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisruptions

`func (o *PolicyConfigChangeDisruptionDetailType) GetDisruptions() []string`

GetDisruptions returns the Disruptions field if non-nil, zero value otherwise.

### GetDisruptionsOk

`func (o *PolicyConfigChangeDisruptionDetailType) GetDisruptionsOk() (*[]string, bool)`

GetDisruptionsOk returns a tuple with the Disruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptions

`func (o *PolicyConfigChangeDisruptionDetailType) SetDisruptions(v []string)`

SetDisruptions sets Disruptions field to given value.

### HasDisruptions

`func (o *PolicyConfigChangeDisruptionDetailType) HasDisruptions() bool`

HasDisruptions returns a boolean if a field has been set.

### SetDisruptionsNil

`func (o *PolicyConfigChangeDisruptionDetailType) SetDisruptionsNil(b bool)`

 SetDisruptionsNil sets the value for Disruptions to be an explicit nil

### UnsetDisruptions
`func (o *PolicyConfigChangeDisruptionDetailType) UnsetDisruptions()`

UnsetDisruptions ensures that no value is present for Disruptions, not even an explicit nil
### GetIsOnlyRequiredByOtherPolicies

`func (o *PolicyConfigChangeDisruptionDetailType) GetIsOnlyRequiredByOtherPolicies() bool`

GetIsOnlyRequiredByOtherPolicies returns the IsOnlyRequiredByOtherPolicies field if non-nil, zero value otherwise.

### GetIsOnlyRequiredByOtherPoliciesOk

`func (o *PolicyConfigChangeDisruptionDetailType) GetIsOnlyRequiredByOtherPoliciesOk() (*bool, bool)`

GetIsOnlyRequiredByOtherPoliciesOk returns a tuple with the IsOnlyRequiredByOtherPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnlyRequiredByOtherPolicies

`func (o *PolicyConfigChangeDisruptionDetailType) SetIsOnlyRequiredByOtherPolicies(v bool)`

SetIsOnlyRequiredByOtherPolicies sets IsOnlyRequiredByOtherPolicies field to given value.

### HasIsOnlyRequiredByOtherPolicies

`func (o *PolicyConfigChangeDisruptionDetailType) HasIsOnlyRequiredByOtherPolicies() bool`

HasIsOnlyRequiredByOtherPolicies returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyConfigChangeDisruptionDetailType) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyConfigChangeDisruptionDetailType) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyPendingAction

`func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyPendingAction() string`

GetPolicyPendingAction returns the PolicyPendingAction field if non-nil, zero value otherwise.

### GetPolicyPendingActionOk

`func (o *PolicyConfigChangeDisruptionDetailType) GetPolicyPendingActionOk() (*string, bool)`

GetPolicyPendingActionOk returns a tuple with the PolicyPendingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPendingAction

`func (o *PolicyConfigChangeDisruptionDetailType) SetPolicyPendingAction(v string)`

SetPolicyPendingAction sets PolicyPendingAction field to given value.

### HasPolicyPendingAction

`func (o *PolicyConfigChangeDisruptionDetailType) HasPolicyPendingAction() bool`

HasPolicyPendingAction returns a boolean if a field has been set.

### GetRequiredByPolicies

`func (o *PolicyConfigChangeDisruptionDetailType) GetRequiredByPolicies() []string`

GetRequiredByPolicies returns the RequiredByPolicies field if non-nil, zero value otherwise.

### GetRequiredByPoliciesOk

`func (o *PolicyConfigChangeDisruptionDetailType) GetRequiredByPoliciesOk() (*[]string, bool)`

GetRequiredByPoliciesOk returns a tuple with the RequiredByPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredByPolicies

`func (o *PolicyConfigChangeDisruptionDetailType) SetRequiredByPolicies(v []string)`

SetRequiredByPolicies sets RequiredByPolicies field to given value.

### HasRequiredByPolicies

`func (o *PolicyConfigChangeDisruptionDetailType) HasRequiredByPolicies() bool`

HasRequiredByPolicies returns a boolean if a field has been set.

### SetRequiredByPoliciesNil

`func (o *PolicyConfigChangeDisruptionDetailType) SetRequiredByPoliciesNil(b bool)`

 SetRequiredByPoliciesNil sets the value for RequiredByPolicies to be an explicit nil

### UnsetRequiredByPolicies
`func (o *PolicyConfigChangeDisruptionDetailType) UnsetRequiredByPolicies()`

UnsetRequiredByPolicies ensures that no value is present for RequiredByPolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


