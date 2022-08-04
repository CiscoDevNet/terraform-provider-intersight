# VmediaPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vmedia.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vmedia.PolicyInventory"]
**Enabled** | Pointer to **bool** | State of the Virtual Media service on the endpoint. | [optional] [readonly] [default to true]
**Encryption** | Pointer to **bool** | If enabled, allows encryption of all Virtual Media communications. Please note that this can no longer be disabled for servers running versions 4.2 and above. | [optional] [readonly] [default to true]
**LowPowerUsb** | Pointer to **bool** | If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. | [optional] [readonly] [default to true]
**Mappings** | Pointer to [**[]VmediaMapping**](VmediaMapping.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVmediaPolicyInventoryAllOf

`func NewVmediaPolicyInventoryAllOf(classId string, objectType string, ) *VmediaPolicyInventoryAllOf`

NewVmediaPolicyInventoryAllOf instantiates a new VmediaPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmediaPolicyInventoryAllOfWithDefaults

`func NewVmediaPolicyInventoryAllOfWithDefaults() *VmediaPolicyInventoryAllOf`

NewVmediaPolicyInventoryAllOfWithDefaults instantiates a new VmediaPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VmediaPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VmediaPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VmediaPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VmediaPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VmediaPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VmediaPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *VmediaPolicyInventoryAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VmediaPolicyInventoryAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VmediaPolicyInventoryAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VmediaPolicyInventoryAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncryption

`func (o *VmediaPolicyInventoryAllOf) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *VmediaPolicyInventoryAllOf) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *VmediaPolicyInventoryAllOf) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *VmediaPolicyInventoryAllOf) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetLowPowerUsb

`func (o *VmediaPolicyInventoryAllOf) GetLowPowerUsb() bool`

GetLowPowerUsb returns the LowPowerUsb field if non-nil, zero value otherwise.

### GetLowPowerUsbOk

`func (o *VmediaPolicyInventoryAllOf) GetLowPowerUsbOk() (*bool, bool)`

GetLowPowerUsbOk returns a tuple with the LowPowerUsb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPowerUsb

`func (o *VmediaPolicyInventoryAllOf) SetLowPowerUsb(v bool)`

SetLowPowerUsb sets LowPowerUsb field to given value.

### HasLowPowerUsb

`func (o *VmediaPolicyInventoryAllOf) HasLowPowerUsb() bool`

HasLowPowerUsb returns a boolean if a field has been set.

### GetMappings

`func (o *VmediaPolicyInventoryAllOf) GetMappings() []VmediaMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *VmediaPolicyInventoryAllOf) GetMappingsOk() (*[]VmediaMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *VmediaPolicyInventoryAllOf) SetMappings(v []VmediaMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *VmediaPolicyInventoryAllOf) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *VmediaPolicyInventoryAllOf) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *VmediaPolicyInventoryAllOf) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetTargetMo

`func (o *VmediaPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VmediaPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VmediaPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VmediaPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


