# CapabilitySwitchControlPolicyConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchControlPolicyConstraints"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchControlPolicyConstraints"]
**MinBundleVersionForStpModeRpvst** | Pointer to **string** | Minimum bundle version required to support RPVST mode for Spanning Tree Protocol. | [optional] [readonly] 
**MinSwitchVersionForStpModeRpvst** | Pointer to **string** | Minimum switch firmware version required to support RPVST mode for Spanning Tree Protocol. | [optional] [readonly] 

## Methods

### NewCapabilitySwitchControlPolicyConstraints

`func NewCapabilitySwitchControlPolicyConstraints(classId string, objectType string, ) *CapabilitySwitchControlPolicyConstraints`

NewCapabilitySwitchControlPolicyConstraints instantiates a new CapabilitySwitchControlPolicyConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchControlPolicyConstraintsWithDefaults

`func NewCapabilitySwitchControlPolicyConstraintsWithDefaults() *CapabilitySwitchControlPolicyConstraints`

NewCapabilitySwitchControlPolicyConstraintsWithDefaults instantiates a new CapabilitySwitchControlPolicyConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchControlPolicyConstraints) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchControlPolicyConstraints) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchControlPolicyConstraints) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchControlPolicyConstraints) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchControlPolicyConstraints) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchControlPolicyConstraints) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMinBundleVersionForStpModeRpvst

`func (o *CapabilitySwitchControlPolicyConstraints) GetMinBundleVersionForStpModeRpvst() string`

GetMinBundleVersionForStpModeRpvst returns the MinBundleVersionForStpModeRpvst field if non-nil, zero value otherwise.

### GetMinBundleVersionForStpModeRpvstOk

`func (o *CapabilitySwitchControlPolicyConstraints) GetMinBundleVersionForStpModeRpvstOk() (*string, bool)`

GetMinBundleVersionForStpModeRpvstOk returns a tuple with the MinBundleVersionForStpModeRpvst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBundleVersionForStpModeRpvst

`func (o *CapabilitySwitchControlPolicyConstraints) SetMinBundleVersionForStpModeRpvst(v string)`

SetMinBundleVersionForStpModeRpvst sets MinBundleVersionForStpModeRpvst field to given value.

### HasMinBundleVersionForStpModeRpvst

`func (o *CapabilitySwitchControlPolicyConstraints) HasMinBundleVersionForStpModeRpvst() bool`

HasMinBundleVersionForStpModeRpvst returns a boolean if a field has been set.

### GetMinSwitchVersionForStpModeRpvst

`func (o *CapabilitySwitchControlPolicyConstraints) GetMinSwitchVersionForStpModeRpvst() string`

GetMinSwitchVersionForStpModeRpvst returns the MinSwitchVersionForStpModeRpvst field if non-nil, zero value otherwise.

### GetMinSwitchVersionForStpModeRpvstOk

`func (o *CapabilitySwitchControlPolicyConstraints) GetMinSwitchVersionForStpModeRpvstOk() (*string, bool)`

GetMinSwitchVersionForStpModeRpvstOk returns a tuple with the MinSwitchVersionForStpModeRpvst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSwitchVersionForStpModeRpvst

`func (o *CapabilitySwitchControlPolicyConstraints) SetMinSwitchVersionForStpModeRpvst(v string)`

SetMinSwitchVersionForStpModeRpvst sets MinSwitchVersionForStpModeRpvst field to given value.

### HasMinSwitchVersionForStpModeRpvst

`func (o *CapabilitySwitchControlPolicyConstraints) HasMinSwitchVersionForStpModeRpvst() bool`

HasMinSwitchVersionForStpModeRpvst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


