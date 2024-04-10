# CapabilityAdapterUpdateConstraintMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.AdapterUpdateConstraintMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.AdapterUpdateConstraintMeta"]
**SupportedPlatform** | Pointer to **string** | Platform for which the constraint is to be enforced. | [optional] [readonly] 

## Methods

### NewCapabilityAdapterUpdateConstraintMeta

`func NewCapabilityAdapterUpdateConstraintMeta(classId string, objectType string, ) *CapabilityAdapterUpdateConstraintMeta`

NewCapabilityAdapterUpdateConstraintMeta instantiates a new CapabilityAdapterUpdateConstraintMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityAdapterUpdateConstraintMetaWithDefaults

`func NewCapabilityAdapterUpdateConstraintMetaWithDefaults() *CapabilityAdapterUpdateConstraintMeta`

NewCapabilityAdapterUpdateConstraintMetaWithDefaults instantiates a new CapabilityAdapterUpdateConstraintMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityAdapterUpdateConstraintMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityAdapterUpdateConstraintMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityAdapterUpdateConstraintMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityAdapterUpdateConstraintMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityAdapterUpdateConstraintMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityAdapterUpdateConstraintMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSupportedPlatform

`func (o *CapabilityAdapterUpdateConstraintMeta) GetSupportedPlatform() string`

GetSupportedPlatform returns the SupportedPlatform field if non-nil, zero value otherwise.

### GetSupportedPlatformOk

`func (o *CapabilityAdapterUpdateConstraintMeta) GetSupportedPlatformOk() (*string, bool)`

GetSupportedPlatformOk returns a tuple with the SupportedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatform

`func (o *CapabilityAdapterUpdateConstraintMeta) SetSupportedPlatform(v string)`

SetSupportedPlatform sets SupportedPlatform field to given value.

### HasSupportedPlatform

`func (o *CapabilityAdapterUpdateConstraintMeta) HasSupportedPlatform() bool`

HasSupportedPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


