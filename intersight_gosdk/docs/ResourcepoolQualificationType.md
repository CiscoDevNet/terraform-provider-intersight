# ResourcepoolQualificationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.QualificationType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.QualificationType"]
**IsStaticAddition** | Pointer to **bool** | The resource is added statically or dynamically. If the value is true, it signifies a static addition, implying that the resource was added with fixed specifications. Conversely, if the value is false, it suggests a dynamic addition, indicating that the resource was incorporated based on certain conditions or qualifiers. | [optional] [readonly] [default to false]
**QualificationPolicies** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewResourcepoolQualificationType

`func NewResourcepoolQualificationType(classId string, objectType string, ) *ResourcepoolQualificationType`

NewResourcepoolQualificationType instantiates a new ResourcepoolQualificationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolQualificationTypeWithDefaults

`func NewResourcepoolQualificationTypeWithDefaults() *ResourcepoolQualificationType`

NewResourcepoolQualificationTypeWithDefaults instantiates a new ResourcepoolQualificationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolQualificationType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolQualificationType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolQualificationType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolQualificationType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolQualificationType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolQualificationType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsStaticAddition

`func (o *ResourcepoolQualificationType) GetIsStaticAddition() bool`

GetIsStaticAddition returns the IsStaticAddition field if non-nil, zero value otherwise.

### GetIsStaticAdditionOk

`func (o *ResourcepoolQualificationType) GetIsStaticAdditionOk() (*bool, bool)`

GetIsStaticAdditionOk returns a tuple with the IsStaticAddition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaticAddition

`func (o *ResourcepoolQualificationType) SetIsStaticAddition(v bool)`

SetIsStaticAddition sets IsStaticAddition field to given value.

### HasIsStaticAddition

`func (o *ResourcepoolQualificationType) HasIsStaticAddition() bool`

HasIsStaticAddition returns a boolean if a field has been set.

### GetQualificationPolicies

`func (o *ResourcepoolQualificationType) GetQualificationPolicies() []MoMoRef`

GetQualificationPolicies returns the QualificationPolicies field if non-nil, zero value otherwise.

### GetQualificationPoliciesOk

`func (o *ResourcepoolQualificationType) GetQualificationPoliciesOk() (*[]MoMoRef, bool)`

GetQualificationPoliciesOk returns a tuple with the QualificationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualificationPolicies

`func (o *ResourcepoolQualificationType) SetQualificationPolicies(v []MoMoRef)`

SetQualificationPolicies sets QualificationPolicies field to given value.

### HasQualificationPolicies

`func (o *ResourcepoolQualificationType) HasQualificationPolicies() bool`

HasQualificationPolicies returns a boolean if a field has been set.

### SetQualificationPoliciesNil

`func (o *ResourcepoolQualificationType) SetQualificationPoliciesNil(b bool)`

 SetQualificationPoliciesNil sets the value for QualificationPolicies to be an explicit nil

### UnsetQualificationPolicies
`func (o *ResourcepoolQualificationType) UnsetQualificationPolicies()`

UnsetQualificationPolicies ensures that no value is present for QualificationPolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


