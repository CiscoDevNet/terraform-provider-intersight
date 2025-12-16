# PolicyPolicyContextHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.PolicyContextHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.PolicyContextHolder"]
**EndPointContext** | Pointer to **string** | Information about the endpoint to which it is applied. * &#x60;Server&#x60; - Configuration is applied to a server context. * &#x60;FI&#x60; - Configuration is applied to a Fabric Identifier (FI) context. * &#x60;IOM&#x60; - Configuration is applied to an Input/Output Module (IOM) context. | [optional] [readonly] [default to "Server"]
**Policy** | Pointer to **string** | The name of the policy for which entry is created. | [optional] [readonly] 

## Methods

### NewPolicyPolicyContextHolder

`func NewPolicyPolicyContextHolder(classId string, objectType string, ) *PolicyPolicyContextHolder`

NewPolicyPolicyContextHolder instantiates a new PolicyPolicyContextHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPolicyContextHolderWithDefaults

`func NewPolicyPolicyContextHolderWithDefaults() *PolicyPolicyContextHolder`

NewPolicyPolicyContextHolderWithDefaults instantiates a new PolicyPolicyContextHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyPolicyContextHolder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyPolicyContextHolder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyPolicyContextHolder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyPolicyContextHolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyPolicyContextHolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyPolicyContextHolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndPointContext

`func (o *PolicyPolicyContextHolder) GetEndPointContext() string`

GetEndPointContext returns the EndPointContext field if non-nil, zero value otherwise.

### GetEndPointContextOk

`func (o *PolicyPolicyContextHolder) GetEndPointContextOk() (*string, bool)`

GetEndPointContextOk returns a tuple with the EndPointContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointContext

`func (o *PolicyPolicyContextHolder) SetEndPointContext(v string)`

SetEndPointContext sets EndPointContext field to given value.

### HasEndPointContext

`func (o *PolicyPolicyContextHolder) HasEndPointContext() bool`

HasEndPointContext returns a boolean if a field has been set.

### GetPolicy

`func (o *PolicyPolicyContextHolder) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyPolicyContextHolder) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyPolicyContextHolder) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyPolicyContextHolder) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


