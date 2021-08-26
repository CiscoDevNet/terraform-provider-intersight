# PolicyConfigChangeContextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ConfigChangeContext"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ConfigChangeContext"]
**ConfigChangeError** | Pointer to **string** | Indicates reason for failure state of configChangeState. | [optional] [readonly] 
**ConfigChangeState** | Pointer to **string** | Indicates a profile&#39;s configuration change state. Used for tracking pending-changes and out-of-synch states. * &#x60;Ok&#x60; - Config change state represents Validation for change/drift is successful or is not applicable. * &#x60;Validating-change&#x60; - Config change state represents policy changes are being validated. This state starts when policy is changed and becomes different from deployed changes (Pending-changes). * &#x60;Validating-drift&#x60; - Config change state represents endpoint configuration changes are being validated. This state starts when endpoint is changed and endpoint configuration becomes different from policy configured changes (Out-of-sync). * &#x60;Change-failed&#x60; - Config change state represents there is internal error in calculating policy change. * &#x60;Drift-failed&#x60; - Config change state represents there is internal error in calculating endpoint configuraion drift. | [optional] [readonly] [default to "Ok"]
**InitialConfigContext** | Pointer to [**NullablePolicyConfigContext**](PolicyConfigContext.md) |  | [optional] 

## Methods

### NewPolicyConfigChangeContextAllOf

`func NewPolicyConfigChangeContextAllOf(classId string, objectType string, ) *PolicyConfigChangeContextAllOf`

NewPolicyConfigChangeContextAllOf instantiates a new PolicyConfigChangeContextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigChangeContextAllOfWithDefaults

`func NewPolicyConfigChangeContextAllOfWithDefaults() *PolicyConfigChangeContextAllOf`

NewPolicyConfigChangeContextAllOfWithDefaults instantiates a new PolicyConfigChangeContextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyConfigChangeContextAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyConfigChangeContextAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyConfigChangeContextAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyConfigChangeContextAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyConfigChangeContextAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyConfigChangeContextAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigChangeError

`func (o *PolicyConfigChangeContextAllOf) GetConfigChangeError() string`

GetConfigChangeError returns the ConfigChangeError field if non-nil, zero value otherwise.

### GetConfigChangeErrorOk

`func (o *PolicyConfigChangeContextAllOf) GetConfigChangeErrorOk() (*string, bool)`

GetConfigChangeErrorOk returns a tuple with the ConfigChangeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeError

`func (o *PolicyConfigChangeContextAllOf) SetConfigChangeError(v string)`

SetConfigChangeError sets ConfigChangeError field to given value.

### HasConfigChangeError

`func (o *PolicyConfigChangeContextAllOf) HasConfigChangeError() bool`

HasConfigChangeError returns a boolean if a field has been set.

### GetConfigChangeState

`func (o *PolicyConfigChangeContextAllOf) GetConfigChangeState() string`

GetConfigChangeState returns the ConfigChangeState field if non-nil, zero value otherwise.

### GetConfigChangeStateOk

`func (o *PolicyConfigChangeContextAllOf) GetConfigChangeStateOk() (*string, bool)`

GetConfigChangeStateOk returns a tuple with the ConfigChangeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeState

`func (o *PolicyConfigChangeContextAllOf) SetConfigChangeState(v string)`

SetConfigChangeState sets ConfigChangeState field to given value.

### HasConfigChangeState

`func (o *PolicyConfigChangeContextAllOf) HasConfigChangeState() bool`

HasConfigChangeState returns a boolean if a field has been set.

### GetInitialConfigContext

`func (o *PolicyConfigChangeContextAllOf) GetInitialConfigContext() PolicyConfigContext`

GetInitialConfigContext returns the InitialConfigContext field if non-nil, zero value otherwise.

### GetInitialConfigContextOk

`func (o *PolicyConfigChangeContextAllOf) GetInitialConfigContextOk() (*PolicyConfigContext, bool)`

GetInitialConfigContextOk returns a tuple with the InitialConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConfigContext

`func (o *PolicyConfigChangeContextAllOf) SetInitialConfigContext(v PolicyConfigContext)`

SetInitialConfigContext sets InitialConfigContext field to given value.

### HasInitialConfigContext

`func (o *PolicyConfigChangeContextAllOf) HasInitialConfigContext() bool`

HasInitialConfigContext returns a boolean if a field has been set.

### SetInitialConfigContextNil

`func (o *PolicyConfigChangeContextAllOf) SetInitialConfigContextNil(b bool)`

 SetInitialConfigContextNil sets the value for InitialConfigContext to be an explicit nil

### UnsetInitialConfigContext
`func (o *PolicyConfigChangeContextAllOf) UnsetInitialConfigContext()`

UnsetInitialConfigContext ensures that no value is present for InitialConfigContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


