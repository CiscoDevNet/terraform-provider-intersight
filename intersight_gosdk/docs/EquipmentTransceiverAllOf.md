# EquipmentTransceiverAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Transceiver"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Transceiver"]
**AggregatePortId** | Pointer to **int64** | Breakout port member in the Fabric Interconnect. | [optional] [readonly] 
**CiscoExtendedIdNumber** | Pointer to **string** | The cisco extended Id number state of the pluggable SFP. | [optional] 
**InterfaceType** | Pointer to **string** | Interface type of transceiver copper or fiber. | [optional] [readonly] 
**ManufacturerPartNumber** | Pointer to **string** | The manufacturer part number of the pluggable SFP. | [optional] 
**ModuleId** | Pointer to **int64** | Fabric extender identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the pluggable transceiver. | [optional] 
**OperSpeed** | Pointer to **string** | Operational speed of the transceiver. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the transceiver. | [optional] [readonly] 
**OperStateQual** | Pointer to **string** | Reason for this transceiver&#39;s operational state. | [optional] [readonly] 
**PortId** | Pointer to **int64** | Switch physical port identifier. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Switch expansion slot module identifier. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the pluggable SFP. | [optional] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the transceiver. | [optional] [readonly] 
**EtherHostPort** | Pointer to [**EtherHostPortRelationship**](EtherHostPortRelationship.md) |  | [optional] 
**EtherPhysicalPort** | Pointer to [**EtherPhysicalPortRelationship**](EtherPhysicalPortRelationship.md) |  | [optional] 
**FcPhysicalPort** | Pointer to [**FcPhysicalPortRelationship**](FcPhysicalPortRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentTransceiverAllOf

`func NewEquipmentTransceiverAllOf(classId string, objectType string, ) *EquipmentTransceiverAllOf`

NewEquipmentTransceiverAllOf instantiates a new EquipmentTransceiverAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentTransceiverAllOfWithDefaults

`func NewEquipmentTransceiverAllOfWithDefaults() *EquipmentTransceiverAllOf`

NewEquipmentTransceiverAllOfWithDefaults instantiates a new EquipmentTransceiverAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentTransceiverAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentTransceiverAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentTransceiverAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentTransceiverAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentTransceiverAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentTransceiverAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregatePortId

`func (o *EquipmentTransceiverAllOf) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *EquipmentTransceiverAllOf) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *EquipmentTransceiverAllOf) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *EquipmentTransceiverAllOf) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetCiscoExtendedIdNumber

`func (o *EquipmentTransceiverAllOf) GetCiscoExtendedIdNumber() string`

GetCiscoExtendedIdNumber returns the CiscoExtendedIdNumber field if non-nil, zero value otherwise.

### GetCiscoExtendedIdNumberOk

`func (o *EquipmentTransceiverAllOf) GetCiscoExtendedIdNumberOk() (*string, bool)`

GetCiscoExtendedIdNumberOk returns a tuple with the CiscoExtendedIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoExtendedIdNumber

`func (o *EquipmentTransceiverAllOf) SetCiscoExtendedIdNumber(v string)`

SetCiscoExtendedIdNumber sets CiscoExtendedIdNumber field to given value.

### HasCiscoExtendedIdNumber

`func (o *EquipmentTransceiverAllOf) HasCiscoExtendedIdNumber() bool`

HasCiscoExtendedIdNumber returns a boolean if a field has been set.

### GetInterfaceType

`func (o *EquipmentTransceiverAllOf) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *EquipmentTransceiverAllOf) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *EquipmentTransceiverAllOf) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *EquipmentTransceiverAllOf) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetManufacturerPartNumber

`func (o *EquipmentTransceiverAllOf) GetManufacturerPartNumber() string`

GetManufacturerPartNumber returns the ManufacturerPartNumber field if non-nil, zero value otherwise.

### GetManufacturerPartNumberOk

`func (o *EquipmentTransceiverAllOf) GetManufacturerPartNumberOk() (*string, bool)`

GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerPartNumber

`func (o *EquipmentTransceiverAllOf) SetManufacturerPartNumber(v string)`

SetManufacturerPartNumber sets ManufacturerPartNumber field to given value.

### HasManufacturerPartNumber

`func (o *EquipmentTransceiverAllOf) HasManufacturerPartNumber() bool`

HasManufacturerPartNumber returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentTransceiverAllOf) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentTransceiverAllOf) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentTransceiverAllOf) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentTransceiverAllOf) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetName

`func (o *EquipmentTransceiverAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentTransceiverAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentTransceiverAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentTransceiverAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EquipmentTransceiverAllOf) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EquipmentTransceiverAllOf) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EquipmentTransceiverAllOf) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EquipmentTransceiverAllOf) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentTransceiverAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentTransceiverAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentTransceiverAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentTransceiverAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *EquipmentTransceiverAllOf) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *EquipmentTransceiverAllOf) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *EquipmentTransceiverAllOf) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *EquipmentTransceiverAllOf) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortId

`func (o *EquipmentTransceiverAllOf) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *EquipmentTransceiverAllOf) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *EquipmentTransceiverAllOf) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *EquipmentTransceiverAllOf) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentTransceiverAllOf) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentTransceiverAllOf) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentTransceiverAllOf) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentTransceiverAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetStatus

`func (o *EquipmentTransceiverAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EquipmentTransceiverAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EquipmentTransceiverAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EquipmentTransceiverAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentTransceiverAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentTransceiverAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentTransceiverAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentTransceiverAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetType

`func (o *EquipmentTransceiverAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EquipmentTransceiverAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EquipmentTransceiverAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EquipmentTransceiverAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEtherHostPort

`func (o *EquipmentTransceiverAllOf) GetEtherHostPort() EtherHostPortRelationship`

GetEtherHostPort returns the EtherHostPort field if non-nil, zero value otherwise.

### GetEtherHostPortOk

`func (o *EquipmentTransceiverAllOf) GetEtherHostPortOk() (*EtherHostPortRelationship, bool)`

GetEtherHostPortOk returns a tuple with the EtherHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherHostPort

`func (o *EquipmentTransceiverAllOf) SetEtherHostPort(v EtherHostPortRelationship)`

SetEtherHostPort sets EtherHostPort field to given value.

### HasEtherHostPort

`func (o *EquipmentTransceiverAllOf) HasEtherHostPort() bool`

HasEtherHostPort returns a boolean if a field has been set.

### GetEtherPhysicalPort

`func (o *EquipmentTransceiverAllOf) GetEtherPhysicalPort() EtherPhysicalPortRelationship`

GetEtherPhysicalPort returns the EtherPhysicalPort field if non-nil, zero value otherwise.

### GetEtherPhysicalPortOk

`func (o *EquipmentTransceiverAllOf) GetEtherPhysicalPortOk() (*EtherPhysicalPortRelationship, bool)`

GetEtherPhysicalPortOk returns a tuple with the EtherPhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherPhysicalPort

`func (o *EquipmentTransceiverAllOf) SetEtherPhysicalPort(v EtherPhysicalPortRelationship)`

SetEtherPhysicalPort sets EtherPhysicalPort field to given value.

### HasEtherPhysicalPort

`func (o *EquipmentTransceiverAllOf) HasEtherPhysicalPort() bool`

HasEtherPhysicalPort returns a boolean if a field has been set.

### GetFcPhysicalPort

`func (o *EquipmentTransceiverAllOf) GetFcPhysicalPort() FcPhysicalPortRelationship`

GetFcPhysicalPort returns the FcPhysicalPort field if non-nil, zero value otherwise.

### GetFcPhysicalPortOk

`func (o *EquipmentTransceiverAllOf) GetFcPhysicalPortOk() (*FcPhysicalPortRelationship, bool)`

GetFcPhysicalPortOk returns a tuple with the FcPhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPhysicalPort

`func (o *EquipmentTransceiverAllOf) SetFcPhysicalPort(v FcPhysicalPortRelationship)`

SetFcPhysicalPort sets FcPhysicalPort field to given value.

### HasFcPhysicalPort

`func (o *EquipmentTransceiverAllOf) HasFcPhysicalPort() bool`

HasFcPhysicalPort returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentTransceiverAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentTransceiverAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentTransceiverAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentTransceiverAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


