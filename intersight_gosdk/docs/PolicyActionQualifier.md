# PolicyActionQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ActionQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ActionQualifier"]
**Qualifiers** | Pointer to [**[]PolicyQualifier**](PolicyQualifier.md) |  | [optional] 
**Scope** | Pointer to **string** | Specify if profile action is targeted for a specific scope. * &#x60;selective&#x60; - Profile deploy action to be limited to modified policies or to the policy list passed. * &#x60;all&#x60; - Profile deploy action to perform deploy of all policies associated with profile. | [optional] [default to "selective"]

## Methods

### NewPolicyActionQualifier

`func NewPolicyActionQualifier(classId string, objectType string, ) *PolicyActionQualifier`

NewPolicyActionQualifier instantiates a new PolicyActionQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyActionQualifierWithDefaults

`func NewPolicyActionQualifierWithDefaults() *PolicyActionQualifier`

NewPolicyActionQualifierWithDefaults instantiates a new PolicyActionQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyActionQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyActionQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyActionQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyActionQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyActionQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyActionQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetQualifiers

`func (o *PolicyActionQualifier) GetQualifiers() []PolicyQualifier`

GetQualifiers returns the Qualifiers field if non-nil, zero value otherwise.

### GetQualifiersOk

`func (o *PolicyActionQualifier) GetQualifiersOk() (*[]PolicyQualifier, bool)`

GetQualifiersOk returns a tuple with the Qualifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiers

`func (o *PolicyActionQualifier) SetQualifiers(v []PolicyQualifier)`

SetQualifiers sets Qualifiers field to given value.

### HasQualifiers

`func (o *PolicyActionQualifier) HasQualifiers() bool`

HasQualifiers returns a boolean if a field has been set.

### SetQualifiersNil

`func (o *PolicyActionQualifier) SetQualifiersNil(b bool)`

 SetQualifiersNil sets the value for Qualifiers to be an explicit nil

### UnsetQualifiers
`func (o *PolicyActionQualifier) UnsetQualifiers()`

UnsetQualifiers ensures that no value is present for Qualifiers, not even an explicit nil
### GetScope

`func (o *PolicyActionQualifier) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PolicyActionQualifier) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PolicyActionQualifier) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PolicyActionQualifier) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


