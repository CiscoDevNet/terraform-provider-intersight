# PolicyPolicyErrorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.PolicyError"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.PolicyError"]
**Message** | Pointer to **string** | Localized message based on the locale setting of the user&#39;s context giving the policy error description. | [optional] [readonly] 
**Type** | Pointer to **string** | The error type indicating the point of failure like policy attach/detach, placement computation (LCP/SCP) or license validation (LCP/SCP Qos). Values  -- attachErr, detachErr, computationErr, licenseErr. * &#x60;none&#x60; - None is set as default value when there is no error and status is ok. * &#x60;attachErr&#x60; - The policy attach has failed. * &#x60;detachErr&#x60; - The policy detach has failed. * &#x60;computationErr&#x60; - The policy specific computation has failed. For example, vNIC placement on Network Adapter in LanConnectivityPolicy. * &#x60;licenseErr&#x60; - The policy specific license enforcement has failed. For example, Cos property in EthernetQoS policy which is a subpolicyof LanConnectivityPolicy requires Advantage License. But if the server attached to the Server Profile does not have anAdvantage License then the license enforcement check would fail during LanConnectivityPolicy validation. | [optional] [readonly] [default to "none"]

## Methods

### NewPolicyPolicyErrorAllOf

`func NewPolicyPolicyErrorAllOf(classId string, objectType string, ) *PolicyPolicyErrorAllOf`

NewPolicyPolicyErrorAllOf instantiates a new PolicyPolicyErrorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPolicyErrorAllOfWithDefaults

`func NewPolicyPolicyErrorAllOfWithDefaults() *PolicyPolicyErrorAllOf`

NewPolicyPolicyErrorAllOfWithDefaults instantiates a new PolicyPolicyErrorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyPolicyErrorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyPolicyErrorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyPolicyErrorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyPolicyErrorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyPolicyErrorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyPolicyErrorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *PolicyPolicyErrorAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyPolicyErrorAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyPolicyErrorAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PolicyPolicyErrorAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *PolicyPolicyErrorAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyPolicyErrorAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyPolicyErrorAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyPolicyErrorAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


