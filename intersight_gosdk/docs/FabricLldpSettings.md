# FabricLldpSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.LldpSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.LldpSettings"]
**ReceiveEnabled** | Pointer to **bool** | Determines if the LLDP frames can be received by an interface on the switch. | [optional] [default to false]
**TransmitEnabled** | Pointer to **bool** | Determines if the LLDP frames can be transmitted by an interface on the switch. | [optional] [default to false]

## Methods

### NewFabricLldpSettings

`func NewFabricLldpSettings(classId string, objectType string, ) *FabricLldpSettings`

NewFabricLldpSettings instantiates a new FabricLldpSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricLldpSettingsWithDefaults

`func NewFabricLldpSettingsWithDefaults() *FabricLldpSettings`

NewFabricLldpSettingsWithDefaults instantiates a new FabricLldpSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricLldpSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricLldpSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricLldpSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricLldpSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricLldpSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricLldpSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReceiveEnabled

`func (o *FabricLldpSettings) GetReceiveEnabled() bool`

GetReceiveEnabled returns the ReceiveEnabled field if non-nil, zero value otherwise.

### GetReceiveEnabledOk

`func (o *FabricLldpSettings) GetReceiveEnabledOk() (*bool, bool)`

GetReceiveEnabledOk returns a tuple with the ReceiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveEnabled

`func (o *FabricLldpSettings) SetReceiveEnabled(v bool)`

SetReceiveEnabled sets ReceiveEnabled field to given value.

### HasReceiveEnabled

`func (o *FabricLldpSettings) HasReceiveEnabled() bool`

HasReceiveEnabled returns a boolean if a field has been set.

### GetTransmitEnabled

`func (o *FabricLldpSettings) GetTransmitEnabled() bool`

GetTransmitEnabled returns the TransmitEnabled field if non-nil, zero value otherwise.

### GetTransmitEnabledOk

`func (o *FabricLldpSettings) GetTransmitEnabledOk() (*bool, bool)`

GetTransmitEnabledOk returns a tuple with the TransmitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitEnabled

`func (o *FabricLldpSettings) SetTransmitEnabled(v bool)`

SetTransmitEnabled sets TransmitEnabled field to given value.

### HasTransmitEnabled

`func (o *FabricLldpSettings) HasTransmitEnabled() bool`

HasTransmitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


