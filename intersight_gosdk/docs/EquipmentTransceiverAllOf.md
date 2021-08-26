# EquipmentTransceiverAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.Transceiver"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.Transceiver"]
**OperSpeed** | Pointer to **string** | Operational speed of the transceiver. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the transceiver. | [optional] [readonly] 
**OperStateQual** | Pointer to **string** | Reason for this transceiver&#39;s operational state. | [optional] [readonly] 
**PortId** | Pointer to **int64** | Switch physical port identifier. | [optional] [readonly] 
**SlotId** | Pointer to **int64** | Switch expansion slot module identifier. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the transceiver. | [optional] [readonly] 
**EtherPhysicalPort** | Pointer to [**EtherPhysicalPortRelationship**](EtherPhysicalPortRelationship.md) |  | [optional] 
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


