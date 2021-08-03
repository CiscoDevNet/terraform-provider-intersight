# PolicyAbstractConfigProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Action** | Pointer to **string** | User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. | [optional] [default to "No-op"]
**ConfigContext** | Pointer to [**NullablePolicyConfigContext**](policy.ConfigContext.md) |  | [optional] 
**PolicyBucket** | Pointer to [**[]PolicyAbstractPolicyRelationship**](PolicyAbstractPolicyRelationship.md) | An array of relationships to policyAbstractPolicy resources. | [optional] 

## Methods

### NewPolicyAbstractConfigProfileAllOf

`func NewPolicyAbstractConfigProfileAllOf(classId string, objectType string, ) *PolicyAbstractConfigProfileAllOf`

NewPolicyAbstractConfigProfileAllOf instantiates a new PolicyAbstractConfigProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigProfileAllOfWithDefaults

`func NewPolicyAbstractConfigProfileAllOfWithDefaults() *PolicyAbstractConfigProfileAllOf`

NewPolicyAbstractConfigProfileAllOfWithDefaults instantiates a new PolicyAbstractConfigProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractConfigProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractConfigProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractConfigProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractConfigProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractConfigProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractConfigProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *PolicyAbstractConfigProfileAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyAbstractConfigProfileAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyAbstractConfigProfileAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PolicyAbstractConfigProfileAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConfigContext

`func (o *PolicyAbstractConfigProfileAllOf) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *PolicyAbstractConfigProfileAllOf) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *PolicyAbstractConfigProfileAllOf) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *PolicyAbstractConfigProfileAllOf) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### SetConfigContextNil

`func (o *PolicyAbstractConfigProfileAllOf) SetConfigContextNil(b bool)`

 SetConfigContextNil sets the value for ConfigContext to be an explicit nil

### UnsetConfigContext
`func (o *PolicyAbstractConfigProfileAllOf) UnsetConfigContext()`

UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
### GetPolicyBucket

`func (o *PolicyAbstractConfigProfileAllOf) GetPolicyBucket() []PolicyAbstractPolicyRelationship`

GetPolicyBucket returns the PolicyBucket field if non-nil, zero value otherwise.

### GetPolicyBucketOk

`func (o *PolicyAbstractConfigProfileAllOf) GetPolicyBucketOk() (*[]PolicyAbstractPolicyRelationship, bool)`

GetPolicyBucketOk returns a tuple with the PolicyBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBucket

`func (o *PolicyAbstractConfigProfileAllOf) SetPolicyBucket(v []PolicyAbstractPolicyRelationship)`

SetPolicyBucket sets PolicyBucket field to given value.

### HasPolicyBucket

`func (o *PolicyAbstractConfigProfileAllOf) HasPolicyBucket() bool`

HasPolicyBucket returns a boolean if a field has been set.

### SetPolicyBucketNil

`func (o *PolicyAbstractConfigProfileAllOf) SetPolicyBucketNil(b bool)`

 SetPolicyBucketNil sets the value for PolicyBucket to be an explicit nil

### UnsetPolicyBucket
`func (o *PolicyAbstractConfigProfileAllOf) UnsetPolicyBucket()`

UnsetPolicyBucket ensures that no value is present for PolicyBucket, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


