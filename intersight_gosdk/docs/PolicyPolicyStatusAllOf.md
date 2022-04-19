# PolicyPolicyStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.PolicyStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.PolicyStatus"]
**Moid** | Pointer to **string** | The object id of the policy being attached. | [optional] 
**Reason** | Pointer to **string** | The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. | [optional] 
**Status** | Pointer to **string** | Indicates if the policy attach/detach was successful or not. Values  -- ok, errored, validating. * &#x60;ok&#x60; - The policy attach/detach is successful. * &#x60;error&#x60; - The policy cannot be attached/detached due to an error. * &#x60;validating&#x60; - The policy preconfig validation is in progress. | [optional] [default to "ok"]
**Type** | Pointer to **string** | The object type of the policy being attached. | [optional] 

## Methods

### NewPolicyPolicyStatusAllOf

`func NewPolicyPolicyStatusAllOf(classId string, objectType string, ) *PolicyPolicyStatusAllOf`

NewPolicyPolicyStatusAllOf instantiates a new PolicyPolicyStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPolicyStatusAllOfWithDefaults

`func NewPolicyPolicyStatusAllOfWithDefaults() *PolicyPolicyStatusAllOf`

NewPolicyPolicyStatusAllOfWithDefaults instantiates a new PolicyPolicyStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyPolicyStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyPolicyStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyPolicyStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyPolicyStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyPolicyStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyPolicyStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoid

`func (o *PolicyPolicyStatusAllOf) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *PolicyPolicyStatusAllOf) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *PolicyPolicyStatusAllOf) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *PolicyPolicyStatusAllOf) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetReason

`func (o *PolicyPolicyStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PolicyPolicyStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PolicyPolicyStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PolicyPolicyStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *PolicyPolicyStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyPolicyStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyPolicyStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PolicyPolicyStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PolicyPolicyStatusAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyPolicyStatusAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyPolicyStatusAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyPolicyStatusAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


