# ResourceAbstractResourceQualificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Qualifiers** | Pointer to [**[]ResourceResourceQualifier**](ResourceResourceQualifier.md) |  | [optional] 

## Methods

### NewResourceAbstractResourceQualificationPolicy

`func NewResourceAbstractResourceQualificationPolicy(classId string, objectType string, ) *ResourceAbstractResourceQualificationPolicy`

NewResourceAbstractResourceQualificationPolicy instantiates a new ResourceAbstractResourceQualificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAbstractResourceQualificationPolicyWithDefaults

`func NewResourceAbstractResourceQualificationPolicyWithDefaults() *ResourceAbstractResourceQualificationPolicy`

NewResourceAbstractResourceQualificationPolicyWithDefaults instantiates a new ResourceAbstractResourceQualificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceAbstractResourceQualificationPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceAbstractResourceQualificationPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceAbstractResourceQualificationPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceAbstractResourceQualificationPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceAbstractResourceQualificationPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceAbstractResourceQualificationPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetQualifiers

`func (o *ResourceAbstractResourceQualificationPolicy) GetQualifiers() []ResourceResourceQualifier`

GetQualifiers returns the Qualifiers field if non-nil, zero value otherwise.

### GetQualifiersOk

`func (o *ResourceAbstractResourceQualificationPolicy) GetQualifiersOk() (*[]ResourceResourceQualifier, bool)`

GetQualifiersOk returns a tuple with the Qualifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiers

`func (o *ResourceAbstractResourceQualificationPolicy) SetQualifiers(v []ResourceResourceQualifier)`

SetQualifiers sets Qualifiers field to given value.

### HasQualifiers

`func (o *ResourceAbstractResourceQualificationPolicy) HasQualifiers() bool`

HasQualifiers returns a boolean if a field has been set.

### SetQualifiersNil

`func (o *ResourceAbstractResourceQualificationPolicy) SetQualifiersNil(b bool)`

 SetQualifiersNil sets the value for Qualifiers to be an explicit nil

### UnsetQualifiers
`func (o *ResourceAbstractResourceQualificationPolicy) UnsetQualifiers()`

UnsetQualifiers ensures that no value is present for Qualifiers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


