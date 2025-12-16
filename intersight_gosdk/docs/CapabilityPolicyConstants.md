# CapabilityPolicyConstants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.PolicyConstants"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.PolicyConstants"]
**BootDeviceReservedKeywords** | Pointer to **string** | List of reserved keywords for boot device policies. | [optional] [readonly] 

## Methods

### NewCapabilityPolicyConstants

`func NewCapabilityPolicyConstants(classId string, objectType string, ) *CapabilityPolicyConstants`

NewCapabilityPolicyConstants instantiates a new CapabilityPolicyConstants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityPolicyConstantsWithDefaults

`func NewCapabilityPolicyConstantsWithDefaults() *CapabilityPolicyConstants`

NewCapabilityPolicyConstantsWithDefaults instantiates a new CapabilityPolicyConstants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityPolicyConstants) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityPolicyConstants) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityPolicyConstants) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityPolicyConstants) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityPolicyConstants) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityPolicyConstants) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootDeviceReservedKeywords

`func (o *CapabilityPolicyConstants) GetBootDeviceReservedKeywords() string`

GetBootDeviceReservedKeywords returns the BootDeviceReservedKeywords field if non-nil, zero value otherwise.

### GetBootDeviceReservedKeywordsOk

`func (o *CapabilityPolicyConstants) GetBootDeviceReservedKeywordsOk() (*string, bool)`

GetBootDeviceReservedKeywordsOk returns a tuple with the BootDeviceReservedKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceReservedKeywords

`func (o *CapabilityPolicyConstants) SetBootDeviceReservedKeywords(v string)`

SetBootDeviceReservedKeywords sets BootDeviceReservedKeywords field to given value.

### HasBootDeviceReservedKeywords

`func (o *CapabilityPolicyConstants) HasBootDeviceReservedKeywords() bool`

HasBootDeviceReservedKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


