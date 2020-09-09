# EquipmentPsuControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterState** | Pointer to **string** | This field identifies the cluster state of the psu redundancy. | [optional] [readonly] 
**InputPowerState** | Pointer to **string** | This field identifies the input power state of the psus. | [optional] [readonly] 
**Name** | Pointer to **string** | This field identifies the name of psu control object. | [optional] [readonly] 
**OperQualifier** | Pointer to **string** | This field identifies the operational qualifier for the psu redundancy. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the operational state of the psu redundancy. | [optional] [readonly] 
**OutputPowerState** | Pointer to **string** | This field identifies the output power state of the psus. | [optional] [readonly] 
**Redundancy** | Pointer to **string** | This field identifies the redundancy state of the psus. | [optional] [readonly] 
**EquipmentChassis** | Pointer to [**EquipmentChassisRelationship**](equipment.Chassis.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentPsuControl

`func NewEquipmentPsuControl() *EquipmentPsuControl`

NewEquipmentPsuControl instantiates a new EquipmentPsuControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPsuControlWithDefaults

`func NewEquipmentPsuControlWithDefaults() *EquipmentPsuControl`

NewEquipmentPsuControlWithDefaults instantiates a new EquipmentPsuControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterState

`func (o *EquipmentPsuControl) GetClusterState() string`

GetClusterState returns the ClusterState field if non-nil, zero value otherwise.

### GetClusterStateOk

`func (o *EquipmentPsuControl) GetClusterStateOk() (*string, bool)`

GetClusterStateOk returns a tuple with the ClusterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterState

`func (o *EquipmentPsuControl) SetClusterState(v string)`

SetClusterState sets ClusterState field to given value.

### HasClusterState

`func (o *EquipmentPsuControl) HasClusterState() bool`

HasClusterState returns a boolean if a field has been set.

### GetInputPowerState

`func (o *EquipmentPsuControl) GetInputPowerState() string`

GetInputPowerState returns the InputPowerState field if non-nil, zero value otherwise.

### GetInputPowerStateOk

`func (o *EquipmentPsuControl) GetInputPowerStateOk() (*string, bool)`

GetInputPowerStateOk returns a tuple with the InputPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPowerState

`func (o *EquipmentPsuControl) SetInputPowerState(v string)`

SetInputPowerState sets InputPowerState field to given value.

### HasInputPowerState

`func (o *EquipmentPsuControl) HasInputPowerState() bool`

HasInputPowerState returns a boolean if a field has been set.

### GetName

`func (o *EquipmentPsuControl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentPsuControl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentPsuControl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentPsuControl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperQualifier

`func (o *EquipmentPsuControl) GetOperQualifier() string`

GetOperQualifier returns the OperQualifier field if non-nil, zero value otherwise.

### GetOperQualifierOk

`func (o *EquipmentPsuControl) GetOperQualifierOk() (*string, bool)`

GetOperQualifierOk returns a tuple with the OperQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperQualifier

`func (o *EquipmentPsuControl) SetOperQualifier(v string)`

SetOperQualifier sets OperQualifier field to given value.

### HasOperQualifier

`func (o *EquipmentPsuControl) HasOperQualifier() bool`

HasOperQualifier returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentPsuControl) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentPsuControl) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentPsuControl) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentPsuControl) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOutputPowerState

`func (o *EquipmentPsuControl) GetOutputPowerState() string`

GetOutputPowerState returns the OutputPowerState field if non-nil, zero value otherwise.

### GetOutputPowerStateOk

`func (o *EquipmentPsuControl) GetOutputPowerStateOk() (*string, bool)`

GetOutputPowerStateOk returns a tuple with the OutputPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPowerState

`func (o *EquipmentPsuControl) SetOutputPowerState(v string)`

SetOutputPowerState sets OutputPowerState field to given value.

### HasOutputPowerState

`func (o *EquipmentPsuControl) HasOutputPowerState() bool`

HasOutputPowerState returns a boolean if a field has been set.

### GetRedundancy

`func (o *EquipmentPsuControl) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *EquipmentPsuControl) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *EquipmentPsuControl) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *EquipmentPsuControl) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentPsuControl) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentPsuControl) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentPsuControl) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentPsuControl) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentPsuControl) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentPsuControl) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentPsuControl) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentPsuControl) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentPsuControl) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentPsuControl) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentPsuControl) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentPsuControl) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


