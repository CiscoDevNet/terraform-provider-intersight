# CapabilityNetworkEquipmentPowerDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.NetworkEquipmentPowerDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.NetworkEquipmentPowerDef"]
**DeviceType** | Pointer to **string** | Device type - FabricInterconnect or Fex. * &#x60;Unknown&#x60; - The device type is unknown. * &#x60;FabricInterconnect&#x60; - The device is a Fabric Interconnect. * &#x60;Fex&#x60; - The device is a Fabric Extender (FEX). | [optional] [default to "Unknown"]
**MaxPower** | Pointer to **int64** | Maximum power consumption in watts. | [optional] 
**Pid** | Pointer to **string** | Product Identifier for the network equipment (e.g., UCS-FI-6454). | [optional] 
**TypicalPower** | Pointer to **int64** | Typical power consumption in watts under normal operating conditions. | [optional] 

## Methods

### NewCapabilityNetworkEquipmentPowerDef

`func NewCapabilityNetworkEquipmentPowerDef(classId string, objectType string, ) *CapabilityNetworkEquipmentPowerDef`

NewCapabilityNetworkEquipmentPowerDef instantiates a new CapabilityNetworkEquipmentPowerDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityNetworkEquipmentPowerDefWithDefaults

`func NewCapabilityNetworkEquipmentPowerDefWithDefaults() *CapabilityNetworkEquipmentPowerDef`

NewCapabilityNetworkEquipmentPowerDefWithDefaults instantiates a new CapabilityNetworkEquipmentPowerDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityNetworkEquipmentPowerDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityNetworkEquipmentPowerDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityNetworkEquipmentPowerDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityNetworkEquipmentPowerDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityNetworkEquipmentPowerDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityNetworkEquipmentPowerDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceType

`func (o *CapabilityNetworkEquipmentPowerDef) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CapabilityNetworkEquipmentPowerDef) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CapabilityNetworkEquipmentPowerDef) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CapabilityNetworkEquipmentPowerDef) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetMaxPower

`func (o *CapabilityNetworkEquipmentPowerDef) GetMaxPower() int64`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *CapabilityNetworkEquipmentPowerDef) GetMaxPowerOk() (*int64, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *CapabilityNetworkEquipmentPowerDef) SetMaxPower(v int64)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *CapabilityNetworkEquipmentPowerDef) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetPid

`func (o *CapabilityNetworkEquipmentPowerDef) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CapabilityNetworkEquipmentPowerDef) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CapabilityNetworkEquipmentPowerDef) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *CapabilityNetworkEquipmentPowerDef) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetTypicalPower

`func (o *CapabilityNetworkEquipmentPowerDef) GetTypicalPower() int64`

GetTypicalPower returns the TypicalPower field if non-nil, zero value otherwise.

### GetTypicalPowerOk

`func (o *CapabilityNetworkEquipmentPowerDef) GetTypicalPowerOk() (*int64, bool)`

GetTypicalPowerOk returns a tuple with the TypicalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypicalPower

`func (o *CapabilityNetworkEquipmentPowerDef) SetTypicalPower(v int64)`

SetTypicalPower sets TypicalPower field to given value.

### HasTypicalPower

`func (o *CapabilityNetworkEquipmentPowerDef) HasTypicalPower() bool`

HasTypicalPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


