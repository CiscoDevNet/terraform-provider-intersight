# FcPhysicalPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fc.PhysicalPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fc.PhysicalPort"]
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
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PortGroup** | Pointer to [**PortGroupRelationship**](PortGroupRelationship.md) |  | [optional] 
**PortSubGroup** | Pointer to [**PortSubGroupRelationship**](PortSubGroupRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewFcPhysicalPort

`func NewFcPhysicalPort(classId string, objectType string, ) *FcPhysicalPort`

NewFcPhysicalPort instantiates a new FcPhysicalPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcPhysicalPortWithDefaults

`func NewFcPhysicalPortWithDefaults() *FcPhysicalPort`

NewFcPhysicalPortWithDefaults instantiates a new FcPhysicalPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcPhysicalPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcPhysicalPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcPhysicalPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcPhysicalPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcPhysicalPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcPhysicalPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FcPhysicalPort) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FcPhysicalPort) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FcPhysicalPort) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FcPhysicalPort) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *FcPhysicalPort) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FcPhysicalPort) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FcPhysicalPort) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FcPhysicalPort) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetB2bCredit

`func (o *FcPhysicalPort) GetB2bCredit() int64`

GetB2bCredit returns the B2bCredit field if non-nil, zero value otherwise.

### GetB2bCreditOk

`func (o *FcPhysicalPort) GetB2bCreditOk() (*int64, bool)`

GetB2bCreditOk returns a tuple with the B2bCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2bCredit

`func (o *FcPhysicalPort) SetB2bCredit(v int64)`

SetB2bCredit sets B2bCredit field to given value.

### HasB2bCredit

`func (o *FcPhysicalPort) HasB2bCredit() bool`

HasB2bCredit returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *FcPhysicalPort) GetMaxSpeed() string`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *FcPhysicalPort) GetMaxSpeedOk() (*string, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *FcPhysicalPort) SetMaxSpeed(v string)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *FcPhysicalPort) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetMode

`func (o *FcPhysicalPort) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FcPhysicalPort) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FcPhysicalPort) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FcPhysicalPort) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperSpeed

`func (o *FcPhysicalPort) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *FcPhysicalPort) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *FcPhysicalPort) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *FcPhysicalPort) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetPeerDn

`func (o *FcPhysicalPort) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *FcPhysicalPort) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *FcPhysicalPort) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *FcPhysicalPort) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPortChannelId

`func (o *FcPhysicalPort) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *FcPhysicalPort) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *FcPhysicalPort) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *FcPhysicalPort) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetTransceiverType

`func (o *FcPhysicalPort) GetTransceiverType() string`

GetTransceiverType returns the TransceiverType field if non-nil, zero value otherwise.

### GetTransceiverTypeOk

`func (o *FcPhysicalPort) GetTransceiverTypeOk() (*string, bool)`

GetTransceiverTypeOk returns a tuple with the TransceiverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverType

`func (o *FcPhysicalPort) SetTransceiverType(v string)`

SetTransceiverType sets TransceiverType field to given value.

### HasTransceiverType

`func (o *FcPhysicalPort) HasTransceiverType() bool`

HasTransceiverType returns a boolean if a field has been set.

### GetVsan

`func (o *FcPhysicalPort) GetVsan() int64`

GetVsan returns the Vsan field if non-nil, zero value otherwise.

### GetVsanOk

`func (o *FcPhysicalPort) GetVsanOk() (*int64, bool)`

GetVsanOk returns a tuple with the Vsan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsan

`func (o *FcPhysicalPort) SetVsan(v int64)`

SetVsan sets Vsan field to given value.

### HasVsan

`func (o *FcPhysicalPort) HasVsan() bool`

HasVsan returns a boolean if a field has been set.

### GetWwn

`func (o *FcPhysicalPort) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *FcPhysicalPort) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *FcPhysicalPort) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *FcPhysicalPort) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *FcPhysicalPort) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *FcPhysicalPort) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *FcPhysicalPort) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *FcPhysicalPort) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroup

`func (o *FcPhysicalPort) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *FcPhysicalPort) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *FcPhysicalPort) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *FcPhysicalPort) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### GetPortSubGroup

`func (o *FcPhysicalPort) GetPortSubGroup() PortSubGroupRelationship`

GetPortSubGroup returns the PortSubGroup field if non-nil, zero value otherwise.

### GetPortSubGroupOk

`func (o *FcPhysicalPort) GetPortSubGroupOk() (*PortSubGroupRelationship, bool)`

GetPortSubGroupOk returns a tuple with the PortSubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSubGroup

`func (o *FcPhysicalPort) SetPortSubGroup(v PortSubGroupRelationship)`

SetPortSubGroup sets PortSubGroup field to given value.

### HasPortSubGroup

`func (o *FcPhysicalPort) HasPortSubGroup() bool`

HasPortSubGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *FcPhysicalPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FcPhysicalPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FcPhysicalPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FcPhysicalPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


