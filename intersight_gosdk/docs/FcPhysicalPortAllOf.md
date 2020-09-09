# FcPhysicalPortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSpeed** | Pointer to **string** | Administrator configured Speed applied on the port. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port. | [optional] [readonly] 
**B2bCredit** | Pointer to **int64** | Buffer to Buffer credits of FC port. | [optional] [readonly] 
**MaxSpeed** | Pointer to **string** | Maximum Speed with which the port operates. | [optional] [readonly] 
**Mode** | Pointer to **string** | Mode information N_proxy, F or E associated to the Fibre Channel port. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | Operational Speed with which the port operates. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerDn for fibre channel physical port. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Port channel id of FC port channel created on FI switch. | [optional] [readonly] 
**TransceiverType** | Pointer to **string** | Transceiver type of a Fibre Channel port. | [optional] [readonly] 
**Vsan** | Pointer to **int64** | Virtual San that is associated to the port. | [optional] [readonly] 
**Wwn** | Pointer to **string** | World Wide Name of a Fibre Channel port. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PortGroup** | Pointer to [**PortGroupRelationship**](port.Group.Relationship.md) |  | [optional] 
**PortSubGroup** | Pointer to [**PortSubGroupRelationship**](port.SubGroup.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewFcPhysicalPortAllOf

`func NewFcPhysicalPortAllOf() *FcPhysicalPortAllOf`

NewFcPhysicalPortAllOf instantiates a new FcPhysicalPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcPhysicalPortAllOfWithDefaults

`func NewFcPhysicalPortAllOfWithDefaults() *FcPhysicalPortAllOf`

NewFcPhysicalPortAllOfWithDefaults instantiates a new FcPhysicalPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSpeed

`func (o *FcPhysicalPortAllOf) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FcPhysicalPortAllOf) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FcPhysicalPortAllOf) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FcPhysicalPortAllOf) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *FcPhysicalPortAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FcPhysicalPortAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FcPhysicalPortAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FcPhysicalPortAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetB2bCredit

`func (o *FcPhysicalPortAllOf) GetB2bCredit() int64`

GetB2bCredit returns the B2bCredit field if non-nil, zero value otherwise.

### GetB2bCreditOk

`func (o *FcPhysicalPortAllOf) GetB2bCreditOk() (*int64, bool)`

GetB2bCreditOk returns a tuple with the B2bCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2bCredit

`func (o *FcPhysicalPortAllOf) SetB2bCredit(v int64)`

SetB2bCredit sets B2bCredit field to given value.

### HasB2bCredit

`func (o *FcPhysicalPortAllOf) HasB2bCredit() bool`

HasB2bCredit returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *FcPhysicalPortAllOf) GetMaxSpeed() string`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *FcPhysicalPortAllOf) GetMaxSpeedOk() (*string, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *FcPhysicalPortAllOf) SetMaxSpeed(v string)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *FcPhysicalPortAllOf) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetMode

`func (o *FcPhysicalPortAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FcPhysicalPortAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FcPhysicalPortAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FcPhysicalPortAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperSpeed

`func (o *FcPhysicalPortAllOf) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *FcPhysicalPortAllOf) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *FcPhysicalPortAllOf) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *FcPhysicalPortAllOf) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetPeerDn

`func (o *FcPhysicalPortAllOf) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *FcPhysicalPortAllOf) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *FcPhysicalPortAllOf) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *FcPhysicalPortAllOf) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortChannelId

`func (o *FcPhysicalPortAllOf) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *FcPhysicalPortAllOf) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *FcPhysicalPortAllOf) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *FcPhysicalPortAllOf) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetTransceiverType

`func (o *FcPhysicalPortAllOf) GetTransceiverType() string`

GetTransceiverType returns the TransceiverType field if non-nil, zero value otherwise.

### GetTransceiverTypeOk

`func (o *FcPhysicalPortAllOf) GetTransceiverTypeOk() (*string, bool)`

GetTransceiverTypeOk returns a tuple with the TransceiverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverType

`func (o *FcPhysicalPortAllOf) SetTransceiverType(v string)`

SetTransceiverType sets TransceiverType field to given value.

### HasTransceiverType

`func (o *FcPhysicalPortAllOf) HasTransceiverType() bool`

HasTransceiverType returns a boolean if a field has been set.

### GetVsan

`func (o *FcPhysicalPortAllOf) GetVsan() int64`

GetVsan returns the Vsan field if non-nil, zero value otherwise.

### GetVsanOk

`func (o *FcPhysicalPortAllOf) GetVsanOk() (*int64, bool)`

GetVsanOk returns a tuple with the Vsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsan

`func (o *FcPhysicalPortAllOf) SetVsan(v int64)`

SetVsan sets Vsan field to given value.

### HasVsan

`func (o *FcPhysicalPortAllOf) HasVsan() bool`

HasVsan returns a boolean if a field has been set.

### GetWwn

`func (o *FcPhysicalPortAllOf) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *FcPhysicalPortAllOf) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *FcPhysicalPortAllOf) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *FcPhysicalPortAllOf) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *FcPhysicalPortAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *FcPhysicalPortAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *FcPhysicalPortAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *FcPhysicalPortAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroup

`func (o *FcPhysicalPortAllOf) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *FcPhysicalPortAllOf) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *FcPhysicalPortAllOf) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *FcPhysicalPortAllOf) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### GetPortSubGroup

`func (o *FcPhysicalPortAllOf) GetPortSubGroup() PortSubGroupRelationship`

GetPortSubGroup returns the PortSubGroup field if non-nil, zero value otherwise.

### GetPortSubGroupOk

`func (o *FcPhysicalPortAllOf) GetPortSubGroupOk() (*PortSubGroupRelationship, bool)`

GetPortSubGroupOk returns a tuple with the PortSubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSubGroup

`func (o *FcPhysicalPortAllOf) SetPortSubGroup(v PortSubGroupRelationship)`

SetPortSubGroup sets PortSubGroup field to given value.

### HasPortSubGroup

`func (o *FcPhysicalPortAllOf) HasPortSubGroup() bool`

HasPortSubGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *FcPhysicalPortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FcPhysicalPortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FcPhysicalPortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FcPhysicalPortAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


