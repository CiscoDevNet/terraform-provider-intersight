# PolicyAbstractPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Description** | Pointer to **string** | Description of the policy. | [optional] 
**Name** | Pointer to **string** | Name of the concrete policy. | [optional] 

## Methods

### NewPolicyAbstractPolicyAllOf

`func NewPolicyAbstractPolicyAllOf(classId string, objectType string, ) *PolicyAbstractPolicyAllOf`

NewPolicyAbstractPolicyAllOf instantiates a new PolicyAbstractPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractPolicyAllOfWithDefaults

`func NewPolicyAbstractPolicyAllOfWithDefaults() *PolicyAbstractPolicyAllOf`

NewPolicyAbstractPolicyAllOfWithDefaults instantiates a new PolicyAbstractPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *PolicyAbstractPolicyAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyAbstractPolicyAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyAbstractPolicyAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyAbstractPolicyAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyAbstractPolicyAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAbstractPolicyAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAbstractPolicyAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyAbstractPolicyAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


