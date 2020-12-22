# FabricLldpSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.LldpSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.LldpSettings"]
**ReceiveEnabled** | Pointer to **bool** | Determines if the LLDP frames can be received by an interface on the switch. | [optional] [default to false]
**TransmitEnabled** | Pointer to **bool** | Determines if the LLDP frames can be transmitted by an interface on the switch. | [optional] [default to false]

## Methods

### NewFabricLldpSettingsAllOf

`func NewFabricLldpSettingsAllOf(classId string, objectType string, ) *FabricLldpSettingsAllOf`

NewFabricLldpSettingsAllOf instantiates a new FabricLldpSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricLldpSettingsAllOfWithDefaults

`func NewFabricLldpSettingsAllOfWithDefaults() *FabricLldpSettingsAllOf`

NewFabricLldpSettingsAllOfWithDefaults instantiates a new FabricLldpSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricLldpSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricLldpSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricLldpSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricLldpSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricLldpSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricLldpSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReceiveEnabled

`func (o *FabricLldpSettingsAllOf) GetReceiveEnabled() bool`

GetReceiveEnabled returns the ReceiveEnabled field if non-nil, zero value otherwise.

### GetReceiveEnabledOk

`func (o *FabricLldpSettingsAllOf) GetReceiveEnabledOk() (*bool, bool)`

GetReceiveEnabledOk returns a tuple with the ReceiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveEnabled

`func (o *FabricLldpSettingsAllOf) SetReceiveEnabled(v bool)`

SetReceiveEnabled sets ReceiveEnabled field to given value.

### HasReceiveEnabled

`func (o *FabricLldpSettingsAllOf) HasReceiveEnabled() bool`

HasReceiveEnabled returns a boolean if a field has been set.

### GetTransmitEnabled

`func (o *FabricLldpSettingsAllOf) GetTransmitEnabled() bool`

GetTransmitEnabled returns the TransmitEnabled field if non-nil, zero value otherwise.

### GetTransmitEnabledOk

`func (o *FabricLldpSettingsAllOf) GetTransmitEnabledOk() (*bool, bool)`

GetTransmitEnabledOk returns a tuple with the TransmitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitEnabled

`func (o *FabricLldpSettingsAllOf) SetTransmitEnabled(v bool)`

SetTransmitEnabled sets TransmitEnabled field to given value.

### HasTransmitEnabled

`func (o *FabricLldpSettingsAllOf) HasTransmitEnabled() bool`

HasTransmitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


