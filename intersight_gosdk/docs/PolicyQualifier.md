# PolicyQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.Qualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.Qualifier"]
**Disruption** | Pointer to **string** | Possible disruption the configuration change might cause. | [optional] 
**Policy** | Pointer to **string** | The name of the policy for which, the deployment is performed. | [optional] 

## Methods

### NewPolicyQualifier

`func NewPolicyQualifier(classId string, objectType string, ) *PolicyQualifier`

NewPolicyQualifier instantiates a new PolicyQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyQualifierWithDefaults

`func NewPolicyQualifierWithDefaults() *PolicyQualifier`

NewPolicyQualifierWithDefaults instantiates a new PolicyQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisruption

`func (o *PolicyQualifier) GetDisruption() string`

GetDisruption returns the Disruption field if non-nil, zero value otherwise.

### GetDisruptionOk

`func (o *PolicyQualifier) GetDisruptionOk() (*string, bool)`

GetDisruptionOk returns a tuple with the Disruption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruption

`func (o *PolicyQualifier) SetDisruption(v string)`

SetDisruption sets Disruption field to given value.

### HasDisruption

`func (o *PolicyQualifier) HasDisruption() bool`

HasDisruption returns a boolean if a field has been set.

### GetPolicy

`func (o *PolicyQualifier) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyQualifier) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyQualifier) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyQualifier) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


