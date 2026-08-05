# PowerPowerGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "power.PowerGroupMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "power.PowerGroupMember"]
**AllocatedPower** | Pointer to **int64** | Power allocated to the power group member by the power group in watts. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | Configuration State of the power group member. * &#x60;Ok&#x60; - The configuration state of the power group member is ok. * &#x60;Unknown&#x60; - The configuration state of the power group member is not known. * &#x60;PendingRemoval&#x60; - The power group member is pending removal. The member has been removed from the power group&#39;s InventoryDeviceMember list but the workflow to remove power limit from the device has not yet completed. Once the workflow succeeds, this member MO will be deleted. | [optional] [readonly] [default to "Ok"]
**DeviceRegistrationMoid** | Pointer to **string** | Moid of the asset.DeviceRegistration corresponding to this member&#39;s inventory device. | [optional] [readonly] 
**MaxRequiredPower** | Pointer to **int64** | Maximum power required by the power group member in watts. | [optional] [readonly] 
**MemberType** | Pointer to **string** | The type of the power group member - Chassis, Fex or Fabric Interconnect. * &#x60;Unknown&#x60; - The power group member type is unknown. * &#x60;Chassis&#x60; - The power group member is a chassis. * &#x60;Fex&#x60; - The power group member is a Fex. * &#x60;FabricInterconnect&#x60; - The power group member is a Fabric Interconnect. | [optional] [readonly] [default to "Unknown"]
**MinRequiredPower** | Pointer to **int64** | Minimum power required by the power group member in watts. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational State of the power group member. * &#x60;Ok&#x60; - The operational state of the power group member is ok. * &#x60;PowerConsumptionNearingLimit&#x60; - Operational state for power group member is PowerConsumptionNearingLimit. This may happen when the current power consumption of a chassis approaches the chassis power limit. | [optional] [readonly] [default to "Ok"]
**InventoryDevice** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**PowerGroupParent** | Pointer to [**NullablePowerPowerGroupRelationship**](PowerPowerGroupRelationship.md) |  | [optional] 

## Methods

### NewPowerPowerGroupMember

`func NewPowerPowerGroupMember(classId string, objectType string, ) *PowerPowerGroupMember`

NewPowerPowerGroupMember instantiates a new PowerPowerGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPowerGroupMemberWithDefaults

`func NewPowerPowerGroupMemberWithDefaults() *PowerPowerGroupMember`

NewPowerPowerGroupMemberWithDefaults instantiates a new PowerPowerGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PowerPowerGroupMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PowerPowerGroupMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PowerPowerGroupMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PowerPowerGroupMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPowerGroupMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPowerGroupMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllocatedPower

`func (o *PowerPowerGroupMember) GetAllocatedPower() int64`

GetAllocatedPower returns the AllocatedPower field if non-nil, zero value otherwise.

### GetAllocatedPowerOk

`func (o *PowerPowerGroupMember) GetAllocatedPowerOk() (*int64, bool)`

GetAllocatedPowerOk returns a tuple with the AllocatedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedPower

`func (o *PowerPowerGroupMember) SetAllocatedPower(v int64)`

SetAllocatedPower sets AllocatedPower field to given value.

### HasAllocatedPower

`func (o *PowerPowerGroupMember) HasAllocatedPower() bool`

HasAllocatedPower returns a boolean if a field has been set.

### GetConfigState

`func (o *PowerPowerGroupMember) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PowerPowerGroupMember) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PowerPowerGroupMember) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PowerPowerGroupMember) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetDeviceRegistrationMoid

`func (o *PowerPowerGroupMember) GetDeviceRegistrationMoid() string`

GetDeviceRegistrationMoid returns the DeviceRegistrationMoid field if non-nil, zero value otherwise.

### GetDeviceRegistrationMoidOk

`func (o *PowerPowerGroupMember) GetDeviceRegistrationMoidOk() (*string, bool)`

GetDeviceRegistrationMoidOk returns a tuple with the DeviceRegistrationMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationMoid

`func (o *PowerPowerGroupMember) SetDeviceRegistrationMoid(v string)`

SetDeviceRegistrationMoid sets DeviceRegistrationMoid field to given value.

### HasDeviceRegistrationMoid

`func (o *PowerPowerGroupMember) HasDeviceRegistrationMoid() bool`

HasDeviceRegistrationMoid returns a boolean if a field has been set.

### GetMaxRequiredPower

`func (o *PowerPowerGroupMember) GetMaxRequiredPower() int64`

GetMaxRequiredPower returns the MaxRequiredPower field if non-nil, zero value otherwise.

### GetMaxRequiredPowerOk

`func (o *PowerPowerGroupMember) GetMaxRequiredPowerOk() (*int64, bool)`

GetMaxRequiredPowerOk returns a tuple with the MaxRequiredPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequiredPower

`func (o *PowerPowerGroupMember) SetMaxRequiredPower(v int64)`

SetMaxRequiredPower sets MaxRequiredPower field to given value.

### HasMaxRequiredPower

`func (o *PowerPowerGroupMember) HasMaxRequiredPower() bool`

HasMaxRequiredPower returns a boolean if a field has been set.

### GetMemberType

`func (o *PowerPowerGroupMember) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *PowerPowerGroupMember) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *PowerPowerGroupMember) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *PowerPowerGroupMember) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetMinRequiredPower

`func (o *PowerPowerGroupMember) GetMinRequiredPower() int64`

GetMinRequiredPower returns the MinRequiredPower field if non-nil, zero value otherwise.

### GetMinRequiredPowerOk

`func (o *PowerPowerGroupMember) GetMinRequiredPowerOk() (*int64, bool)`

GetMinRequiredPowerOk returns a tuple with the MinRequiredPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRequiredPower

`func (o *PowerPowerGroupMember) SetMinRequiredPower(v int64)`

SetMinRequiredPower sets MinRequiredPower field to given value.

### HasMinRequiredPower

`func (o *PowerPowerGroupMember) HasMinRequiredPower() bool`

HasMinRequiredPower returns a boolean if a field has been set.

### GetOperState

`func (o *PowerPowerGroupMember) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PowerPowerGroupMember) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PowerPowerGroupMember) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PowerPowerGroupMember) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetInventoryDevice

`func (o *PowerPowerGroupMember) GetInventoryDevice() MoBaseMoRelationship`

GetInventoryDevice returns the InventoryDevice field if non-nil, zero value otherwise.

### GetInventoryDeviceOk

`func (o *PowerPowerGroupMember) GetInventoryDeviceOk() (*MoBaseMoRelationship, bool)`

GetInventoryDeviceOk returns a tuple with the InventoryDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDevice

`func (o *PowerPowerGroupMember) SetInventoryDevice(v MoBaseMoRelationship)`

SetInventoryDevice sets InventoryDevice field to given value.

### HasInventoryDevice

`func (o *PowerPowerGroupMember) HasInventoryDevice() bool`

HasInventoryDevice returns a boolean if a field has been set.

### SetInventoryDeviceNil

`func (o *PowerPowerGroupMember) SetInventoryDeviceNil(b bool)`

 SetInventoryDeviceNil sets the value for InventoryDevice to be an explicit nil

### UnsetInventoryDevice
`func (o *PowerPowerGroupMember) UnsetInventoryDevice()`

UnsetInventoryDevice ensures that no value is present for InventoryDevice, not even an explicit nil
### GetPowerGroupParent

`func (o *PowerPowerGroupMember) GetPowerGroupParent() PowerPowerGroupRelationship`

GetPowerGroupParent returns the PowerGroupParent field if non-nil, zero value otherwise.

### GetPowerGroupParentOk

`func (o *PowerPowerGroupMember) GetPowerGroupParentOk() (*PowerPowerGroupRelationship, bool)`

GetPowerGroupParentOk returns a tuple with the PowerGroupParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerGroupParent

`func (o *PowerPowerGroupMember) SetPowerGroupParent(v PowerPowerGroupRelationship)`

SetPowerGroupParent sets PowerGroupParent field to given value.

### HasPowerGroupParent

`func (o *PowerPowerGroupMember) HasPowerGroupParent() bool`

HasPowerGroupParent returns a boolean if a field has been set.

### SetPowerGroupParentNil

`func (o *PowerPowerGroupMember) SetPowerGroupParentNil(b bool)`

 SetPowerGroupParentNil sets the value for PowerGroupParent to be an explicit nil

### UnsetPowerGroupParent
`func (o *PowerPowerGroupMember) UnsetPowerGroupParent()`

UnsetPowerGroupParent ensures that no value is present for PowerGroupParent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


