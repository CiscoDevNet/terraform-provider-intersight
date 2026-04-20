# CapabilityPortPropertyConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.PortPropertyConstraints"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.PortPropertyConstraints"]
**LagSuspendIndividualPortSupported** | Pointer to **bool** | Indicates whether LAG suspend individual port setting is supported on this platform. | [optional] [readonly] [default to true]
**MinBundleVersionFor25gSpeed** | Pointer to **string** | Minimum bundle version required to support 25G speed on ports. | [optional] [readonly] 
**MinSwitchVersionFor25gSpeed** | Pointer to **string** | Minimum switch firmware version required to support 25G speed on ports. | [optional] [readonly] 

## Methods

### NewCapabilityPortPropertyConstraints

`func NewCapabilityPortPropertyConstraints(classId string, objectType string, ) *CapabilityPortPropertyConstraints`

NewCapabilityPortPropertyConstraints instantiates a new CapabilityPortPropertyConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityPortPropertyConstraintsWithDefaults

`func NewCapabilityPortPropertyConstraintsWithDefaults() *CapabilityPortPropertyConstraints`

NewCapabilityPortPropertyConstraintsWithDefaults instantiates a new CapabilityPortPropertyConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityPortPropertyConstraints) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityPortPropertyConstraints) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityPortPropertyConstraints) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityPortPropertyConstraints) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityPortPropertyConstraints) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityPortPropertyConstraints) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLagSuspendIndividualPortSupported

`func (o *CapabilityPortPropertyConstraints) GetLagSuspendIndividualPortSupported() bool`

GetLagSuspendIndividualPortSupported returns the LagSuspendIndividualPortSupported field if non-nil, zero value otherwise.

### GetLagSuspendIndividualPortSupportedOk

`func (o *CapabilityPortPropertyConstraints) GetLagSuspendIndividualPortSupportedOk() (*bool, bool)`

GetLagSuspendIndividualPortSupportedOk returns a tuple with the LagSuspendIndividualPortSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSuspendIndividualPortSupported

`func (o *CapabilityPortPropertyConstraints) SetLagSuspendIndividualPortSupported(v bool)`

SetLagSuspendIndividualPortSupported sets LagSuspendIndividualPortSupported field to given value.

### HasLagSuspendIndividualPortSupported

`func (o *CapabilityPortPropertyConstraints) HasLagSuspendIndividualPortSupported() bool`

HasLagSuspendIndividualPortSupported returns a boolean if a field has been set.

### GetMinBundleVersionFor25gSpeed

`func (o *CapabilityPortPropertyConstraints) GetMinBundleVersionFor25gSpeed() string`

GetMinBundleVersionFor25gSpeed returns the MinBundleVersionFor25gSpeed field if non-nil, zero value otherwise.

### GetMinBundleVersionFor25gSpeedOk

`func (o *CapabilityPortPropertyConstraints) GetMinBundleVersionFor25gSpeedOk() (*string, bool)`

GetMinBundleVersionFor25gSpeedOk returns a tuple with the MinBundleVersionFor25gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBundleVersionFor25gSpeed

`func (o *CapabilityPortPropertyConstraints) SetMinBundleVersionFor25gSpeed(v string)`

SetMinBundleVersionFor25gSpeed sets MinBundleVersionFor25gSpeed field to given value.

### HasMinBundleVersionFor25gSpeed

`func (o *CapabilityPortPropertyConstraints) HasMinBundleVersionFor25gSpeed() bool`

HasMinBundleVersionFor25gSpeed returns a boolean if a field has been set.

### GetMinSwitchVersionFor25gSpeed

`func (o *CapabilityPortPropertyConstraints) GetMinSwitchVersionFor25gSpeed() string`

GetMinSwitchVersionFor25gSpeed returns the MinSwitchVersionFor25gSpeed field if non-nil, zero value otherwise.

### GetMinSwitchVersionFor25gSpeedOk

`func (o *CapabilityPortPropertyConstraints) GetMinSwitchVersionFor25gSpeedOk() (*string, bool)`

GetMinSwitchVersionFor25gSpeedOk returns a tuple with the MinSwitchVersionFor25gSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSwitchVersionFor25gSpeed

`func (o *CapabilityPortPropertyConstraints) SetMinSwitchVersionFor25gSpeed(v string)`

SetMinSwitchVersionFor25gSpeed sets MinSwitchVersionFor25gSpeed field to given value.

### HasMinSwitchVersionFor25gSpeed

`func (o *CapabilityPortPropertyConstraints) HasMinSwitchVersionFor25gSpeed() bool`

HasMinSwitchVersionFor25gSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


