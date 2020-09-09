# EquipmentPsuControlAllOf

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

### NewEquipmentPsuControlAllOf

`func NewEquipmentPsuControlAllOf() *EquipmentPsuControlAllOf`

NewEquipmentPsuControlAllOf instantiates a new EquipmentPsuControlAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPsuControlAllOfWithDefaults

`func NewEquipmentPsuControlAllOfWithDefaults() *EquipmentPsuControlAllOf`

NewEquipmentPsuControlAllOfWithDefaults instantiates a new EquipmentPsuControlAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterState

`func (o *EquipmentPsuControlAllOf) GetClusterState() string`

GetClusterState returns the ClusterState field if non-nil, zero value otherwise.

### GetClusterStateOk

`func (o *EquipmentPsuControlAllOf) GetClusterStateOk() (*string, bool)`

GetClusterStateOk returns a tuple with the ClusterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterState

`func (o *EquipmentPsuControlAllOf) SetClusterState(v string)`

SetClusterState sets ClusterState field to given value.

### HasClusterState

`func (o *EquipmentPsuControlAllOf) HasClusterState() bool`

HasClusterState returns a boolean if a field has been set.

### GetInputPowerState

`func (o *EquipmentPsuControlAllOf) GetInputPowerState() string`

GetInputPowerState returns the InputPowerState field if non-nil, zero value otherwise.

### GetInputPowerStateOk

`func (o *EquipmentPsuControlAllOf) GetInputPowerStateOk() (*string, bool)`

GetInputPowerStateOk returns a tuple with the InputPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPowerState

`func (o *EquipmentPsuControlAllOf) SetInputPowerState(v string)`

SetInputPowerState sets InputPowerState field to given value.

### HasInputPowerState

`func (o *EquipmentPsuControlAllOf) HasInputPowerState() bool`

HasInputPowerState returns a boolean if a field has been set.

### GetName

`func (o *EquipmentPsuControlAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentPsuControlAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentPsuControlAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentPsuControlAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperQualifier

`func (o *EquipmentPsuControlAllOf) GetOperQualifier() string`

GetOperQualifier returns the OperQualifier field if non-nil, zero value otherwise.

### GetOperQualifierOk

`func (o *EquipmentPsuControlAllOf) GetOperQualifierOk() (*string, bool)`

GetOperQualifierOk returns a tuple with the OperQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperQualifier

`func (o *EquipmentPsuControlAllOf) SetOperQualifier(v string)`

SetOperQualifier sets OperQualifier field to given value.

### HasOperQualifier

`func (o *EquipmentPsuControlAllOf) HasOperQualifier() bool`

HasOperQualifier returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentPsuControlAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentPsuControlAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentPsuControlAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentPsuControlAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOutputPowerState

`func (o *EquipmentPsuControlAllOf) GetOutputPowerState() string`

GetOutputPowerState returns the OutputPowerState field if non-nil, zero value otherwise.

### GetOutputPowerStateOk

`func (o *EquipmentPsuControlAllOf) GetOutputPowerStateOk() (*string, bool)`

GetOutputPowerStateOk returns a tuple with the OutputPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPowerState

`func (o *EquipmentPsuControlAllOf) SetOutputPowerState(v string)`

SetOutputPowerState sets OutputPowerState field to given value.

### HasOutputPowerState

`func (o *EquipmentPsuControlAllOf) HasOutputPowerState() bool`

HasOutputPowerState returns a boolean if a field has been set.

### GetRedundancy

`func (o *EquipmentPsuControlAllOf) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *EquipmentPsuControlAllOf) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *EquipmentPsuControlAllOf) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *EquipmentPsuControlAllOf) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetEquipmentChassis

`func (o *EquipmentPsuControlAllOf) GetEquipmentChassis() EquipmentChassisRelationship`

GetEquipmentChassis returns the EquipmentChassis field if non-nil, zero value otherwise.

### GetEquipmentChassisOk

`func (o *EquipmentPsuControlAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool)`

GetEquipmentChassisOk returns a tuple with the EquipmentChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentChassis

`func (o *EquipmentPsuControlAllOf) SetEquipmentChassis(v EquipmentChassisRelationship)`

SetEquipmentChassis sets EquipmentChassis field to given value.

### HasEquipmentChassis

`func (o *EquipmentPsuControlAllOf) HasEquipmentChassis() bool`

HasEquipmentChassis returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentPsuControlAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentPsuControlAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentPsuControlAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentPsuControlAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentPsuControlAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentPsuControlAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentPsuControlAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentPsuControlAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


