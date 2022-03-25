# EquipmentTransceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Transceiver"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Transceiver"]
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

### NewEquipmentTransceiver

`func NewEquipmentTransceiver(classId string, objectType string, ) *EquipmentTransceiver`

NewEquipmentTransceiver instantiates a new EquipmentTransceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentTransceiverWithDefaults

`func NewEquipmentTransceiverWithDefaults() *EquipmentTransceiver`

NewEquipmentTransceiverWithDefaults instantiates a new EquipmentTransceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentTransceiver) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentTransceiver) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentTransceiver) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentTransceiver) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentTransceiver) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentTransceiver) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCiscoExtendedIdNumber

`func (o *EquipmentTransceiver) GetCiscoExtendedIdNumber() string`

GetCiscoExtendedIdNumber returns the CiscoExtendedIdNumber field if non-nil, zero value otherwise.

### GetCiscoExtendedIdNumberOk

`func (o *EquipmentTransceiver) GetCiscoExtendedIdNumberOk() (*string, bool)`

GetCiscoExtendedIdNumberOk returns a tuple with the CiscoExtendedIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoExtendedIdNumber

`func (o *EquipmentTransceiver) SetCiscoExtendedIdNumber(v string)`

SetCiscoExtendedIdNumber sets CiscoExtendedIdNumber field to given value.

### HasCiscoExtendedIdNumber

`func (o *EquipmentTransceiver) HasCiscoExtendedIdNumber() bool`

HasCiscoExtendedIdNumber returns a boolean if a field has been set.

### GetInterfaceType

`func (o *EquipmentTransceiver) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *EquipmentTransceiver) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *EquipmentTransceiver) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *EquipmentTransceiver) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetManufacturerPartNumber

`func (o *EquipmentTransceiver) GetManufacturerPartNumber() string`

GetManufacturerPartNumber returns the ManufacturerPartNumber field if non-nil, zero value otherwise.

### GetManufacturerPartNumberOk

`func (o *EquipmentTransceiver) GetManufacturerPartNumberOk() (*string, bool)`

GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerPartNumber

`func (o *EquipmentTransceiver) SetManufacturerPartNumber(v string)`

SetManufacturerPartNumber sets ManufacturerPartNumber field to given value.

### HasManufacturerPartNumber

`func (o *EquipmentTransceiver) HasManufacturerPartNumber() bool`

HasManufacturerPartNumber returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentTransceiver) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentTransceiver) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentTransceiver) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentTransceiver) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetName

`func (o *EquipmentTransceiver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentTransceiver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentTransceiver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentTransceiver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EquipmentTransceiver) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EquipmentTransceiver) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EquipmentTransceiver) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EquipmentTransceiver) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentTransceiver) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentTransceiver) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentTransceiver) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentTransceiver) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *EquipmentTransceiver) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *EquipmentTransceiver) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *EquipmentTransceiver) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *EquipmentTransceiver) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortId

`func (o *EquipmentTransceiver) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *EquipmentTransceiver) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *EquipmentTransceiver) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *EquipmentTransceiver) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetSlotId

`func (o *EquipmentTransceiver) GetSlotId() int64`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentTransceiver) GetSlotIdOk() (*int64, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentTransceiver) SetSlotId(v int64)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentTransceiver) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetStatus

`func (o *EquipmentTransceiver) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EquipmentTransceiver) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EquipmentTransceiver) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EquipmentTransceiver) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentTransceiver) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentTransceiver) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentTransceiver) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentTransceiver) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetType

`func (o *EquipmentTransceiver) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EquipmentTransceiver) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EquipmentTransceiver) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EquipmentTransceiver) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEtherHostPort

`func (o *EquipmentTransceiver) GetEtherHostPort() EtherHostPortRelationship`

GetEtherHostPort returns the EtherHostPort field if non-nil, zero value otherwise.

### GetEtherHostPortOk

`func (o *EquipmentTransceiver) GetEtherHostPortOk() (*EtherHostPortRelationship, bool)`

GetEtherHostPortOk returns a tuple with the EtherHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherHostPort

`func (o *EquipmentTransceiver) SetEtherHostPort(v EtherHostPortRelationship)`

SetEtherHostPort sets EtherHostPort field to given value.

### HasEtherHostPort

`func (o *EquipmentTransceiver) HasEtherHostPort() bool`

HasEtherHostPort returns a boolean if a field has been set.

### GetEtherPhysicalPort

`func (o *EquipmentTransceiver) GetEtherPhysicalPort() EtherPhysicalPortRelationship`

GetEtherPhysicalPort returns the EtherPhysicalPort field if non-nil, zero value otherwise.

### GetEtherPhysicalPortOk

`func (o *EquipmentTransceiver) GetEtherPhysicalPortOk() (*EtherPhysicalPortRelationship, bool)`

GetEtherPhysicalPortOk returns a tuple with the EtherPhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherPhysicalPort

`func (o *EquipmentTransceiver) SetEtherPhysicalPort(v EtherPhysicalPortRelationship)`

SetEtherPhysicalPort sets EtherPhysicalPort field to given value.

### HasEtherPhysicalPort

`func (o *EquipmentTransceiver) HasEtherPhysicalPort() bool`

HasEtherPhysicalPort returns a boolean if a field has been set.

### GetFcPhysicalPort

`func (o *EquipmentTransceiver) GetFcPhysicalPort() FcPhysicalPortRelationship`

GetFcPhysicalPort returns the FcPhysicalPort field if non-nil, zero value otherwise.

### GetFcPhysicalPortOk

`func (o *EquipmentTransceiver) GetFcPhysicalPortOk() (*FcPhysicalPortRelationship, bool)`

GetFcPhysicalPortOk returns a tuple with the FcPhysicalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPhysicalPort

`func (o *EquipmentTransceiver) SetFcPhysicalPort(v FcPhysicalPortRelationship)`

SetFcPhysicalPort sets FcPhysicalPort field to given value.

### HasFcPhysicalPort

`func (o *EquipmentTransceiver) HasFcPhysicalPort() bool`

HasFcPhysicalPort returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EquipmentTransceiver) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentTransceiver) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentTransceiver) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentTransceiver) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


