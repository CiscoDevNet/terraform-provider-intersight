# PolicyAbstractPendingChangeEval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "server.ProfilePendingChangeEval"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "server.ProfilePendingChangeEval"]
**ChangeId** | Pointer to **string** | Change id for which the evaluation is pending. | [optional] [readonly] 
**ChangedPolicy** | Pointer to [**NullablePolicyAbstractPolicyRelationship**](PolicyAbstractPolicyRelationship.md) |  | [optional] 

## Methods

### NewPolicyAbstractPendingChangeEval

`func NewPolicyAbstractPendingChangeEval(classId string, objectType string, ) *PolicyAbstractPendingChangeEval`

NewPolicyAbstractPendingChangeEval instantiates a new PolicyAbstractPendingChangeEval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractPendingChangeEvalWithDefaults

`func NewPolicyAbstractPendingChangeEvalWithDefaults() *PolicyAbstractPendingChangeEval`

NewPolicyAbstractPendingChangeEvalWithDefaults instantiates a new PolicyAbstractPendingChangeEval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractPendingChangeEval) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractPendingChangeEval) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractPendingChangeEval) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractPendingChangeEval) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractPendingChangeEval) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractPendingChangeEval) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChangeId

`func (o *PolicyAbstractPendingChangeEval) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *PolicyAbstractPendingChangeEval) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *PolicyAbstractPendingChangeEval) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *PolicyAbstractPendingChangeEval) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetChangedPolicy

`func (o *PolicyAbstractPendingChangeEval) GetChangedPolicy() PolicyAbstractPolicyRelationship`

GetChangedPolicy returns the ChangedPolicy field if non-nil, zero value otherwise.

### GetChangedPolicyOk

`func (o *PolicyAbstractPendingChangeEval) GetChangedPolicyOk() (*PolicyAbstractPolicyRelationship, bool)`

GetChangedPolicyOk returns a tuple with the ChangedPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedPolicy

`func (o *PolicyAbstractPendingChangeEval) SetChangedPolicy(v PolicyAbstractPolicyRelationship)`

SetChangedPolicy sets ChangedPolicy field to given value.

### HasChangedPolicy

`func (o *PolicyAbstractPendingChangeEval) HasChangedPolicy() bool`

HasChangedPolicy returns a boolean if a field has been set.

### SetChangedPolicyNil

`func (o *PolicyAbstractPendingChangeEval) SetChangedPolicyNil(b bool)`

 SetChangedPolicyNil sets the value for ChangedPolicy to be an explicit nil

### UnsetChangedPolicy
`func (o *PolicyAbstractPendingChangeEval) UnsetChangedPolicy()`

UnsetChangedPolicy ensures that no value is present for ChangedPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


