# CapabilitySwitchingModeCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchingModeCapability"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchingModeCapability"]
**SwitchingMode** | Pointer to **string** | Switching mode type (endhost, switch) of the switch. * &#x60;end-host&#x60; - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * &#x60;switch&#x60; - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. | [optional] [default to "end-host"]
**VpCompressionSupported** | Pointer to **bool** | VP Compression support on this switch. | [optional] 

## Methods

### NewCapabilitySwitchingModeCapability

`func NewCapabilitySwitchingModeCapability(classId string, objectType string, ) *CapabilitySwitchingModeCapability`

NewCapabilitySwitchingModeCapability instantiates a new CapabilitySwitchingModeCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchingModeCapabilityWithDefaults

`func NewCapabilitySwitchingModeCapabilityWithDefaults() *CapabilitySwitchingModeCapability`

NewCapabilitySwitchingModeCapabilityWithDefaults instantiates a new CapabilitySwitchingModeCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchingModeCapability) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchingModeCapability) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchingModeCapability) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchingModeCapability) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchingModeCapability) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchingModeCapability) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSwitchingMode

`func (o *CapabilitySwitchingModeCapability) GetSwitchingMode() string`

GetSwitchingMode returns the SwitchingMode field if non-nil, zero value otherwise.

### GetSwitchingModeOk

`func (o *CapabilitySwitchingModeCapability) GetSwitchingModeOk() (*string, bool)`

GetSwitchingModeOk returns a tuple with the SwitchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchingMode

`func (o *CapabilitySwitchingModeCapability) SetSwitchingMode(v string)`

SetSwitchingMode sets SwitchingMode field to given value.

### HasSwitchingMode

`func (o *CapabilitySwitchingModeCapability) HasSwitchingMode() bool`

HasSwitchingMode returns a boolean if a field has been set.

### GetVpCompressionSupported

`func (o *CapabilitySwitchingModeCapability) GetVpCompressionSupported() bool`

GetVpCompressionSupported returns the VpCompressionSupported field if non-nil, zero value otherwise.

### GetVpCompressionSupportedOk

`func (o *CapabilitySwitchingModeCapability) GetVpCompressionSupportedOk() (*bool, bool)`

GetVpCompressionSupportedOk returns a tuple with the VpCompressionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpCompressionSupported

`func (o *CapabilitySwitchingModeCapability) SetVpCompressionSupported(v bool)`

SetVpCompressionSupported sets VpCompressionSupported field to given value.

### HasVpCompressionSupported

`func (o *CapabilitySwitchingModeCapability) HasVpCompressionSupported() bool`

HasVpCompressionSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


