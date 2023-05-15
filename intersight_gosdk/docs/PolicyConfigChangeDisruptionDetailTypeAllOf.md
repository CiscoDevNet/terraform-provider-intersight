# PolicyConfigChangeDisruptionDetailTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ConfigChangeDisruptionDetailType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ConfigChangeDisruptionDetailType"]
**Disruptions** | Pointer to **[]string** |  | [optional] 
**PolicyName** | Pointer to **string** | Name of the policy that, when modified, causes the disruption. | [optional] 
**PolicyPendingAction** | Pointer to **string** | Name of the action which is pending on this policy. Example, if policy is not yet activated we mark this field as not-activated. Currently we support two actions, not-deployed and not-activated. | [optional] 

## Methods

### NewPolicyConfigChangeDisruptionDetailTypeAllOf

`func NewPolicyConfigChangeDisruptionDetailTypeAllOf(classId string, objectType string, ) *PolicyConfigChangeDisruptionDetailTypeAllOf`

NewPolicyConfigChangeDisruptionDetailTypeAllOf instantiates a new PolicyConfigChangeDisruptionDetailTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigChangeDisruptionDetailTypeAllOfWithDefaults

`func NewPolicyConfigChangeDisruptionDetailTypeAllOfWithDefaults() *PolicyConfigChangeDisruptionDetailTypeAllOf`

NewPolicyConfigChangeDisruptionDetailTypeAllOfWithDefaults instantiates a new PolicyConfigChangeDisruptionDetailTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisruptions

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetDisruptions() []string`

GetDisruptions returns the Disruptions field if non-nil, zero value otherwise.

### GetDisruptionsOk

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetDisruptionsOk() (*[]string, bool)`

GetDisruptionsOk returns a tuple with the Disruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptions

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) SetDisruptions(v []string)`

SetDisruptions sets Disruptions field to given value.

### HasDisruptions

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) HasDisruptions() bool`

HasDisruptions returns a boolean if a field has been set.

### SetDisruptionsNil

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) SetDisruptionsNil(b bool)`

 SetDisruptionsNil sets the value for Disruptions to be an explicit nil

### UnsetDisruptions
`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) UnsetDisruptions()`

UnsetDisruptions ensures that no value is present for Disruptions, not even an explicit nil
### GetPolicyName

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyPendingAction

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetPolicyPendingAction() string`

GetPolicyPendingAction returns the PolicyPendingAction field if non-nil, zero value otherwise.

### GetPolicyPendingActionOk

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) GetPolicyPendingActionOk() (*string, bool)`

GetPolicyPendingActionOk returns a tuple with the PolicyPendingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPendingAction

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) SetPolicyPendingAction(v string)`

SetPolicyPendingAction sets PolicyPendingAction field to given value.

### HasPolicyPendingAction

`func (o *PolicyConfigChangeDisruptionDetailTypeAllOf) HasPolicyPendingAction() bool`

HasPolicyPendingAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


