# EquipmentIoCardBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** | Connectivity Status of FEX/IOM to Switch - A or B or AB. | [optional] 
**Description** | Pointer to **string** | This field is to provide description for the iocard module model. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | Module Identifier for the IO module. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of IO card or fabric extender. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part Number identifier for the IO module. | [optional] [readonly] 
**Pid** | Pointer to **string** | This field identifies the Product ID for the IO module. | [optional] [readonly] 
**Presence** | Pointer to **string** | This field identifies the Presence state of the IO card module. | [optional] [readonly] 
**ProductName** | Pointer to **string** | This field identifies the Product Name for the iocard module model. | [optional] [readonly] 
**Sku** | Pointer to **string** | This field identifies the Stock Keeping Unit for the IO card module. | [optional] [readonly] 
**Version** | Pointer to **string** | This field identifies the version of the IO card module. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the Vendor ID for the IO card module. | [optional] [readonly] 
**HostPorts** | Pointer to [**[]EtherHostPortRelationship**](ether.HostPort.Relationship.md) | An array of relationships to etherHostPort resources. | [optional] 
**MgmtController** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**NetworkPorts** | Pointer to [**[]EtherNetworkPortRelationship**](ether.NetworkPort.Relationship.md) | An array of relationships to etherNetworkPort resources. | [optional] 

## Methods

### NewEquipmentIoCardBaseAllOf

`func NewEquipmentIoCardBaseAllOf() *EquipmentIoCardBaseAllOf`

NewEquipmentIoCardBaseAllOf instantiates a new EquipmentIoCardBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardBaseAllOfWithDefaults

`func NewEquipmentIoCardBaseAllOfWithDefaults() *EquipmentIoCardBaseAllOf`

NewEquipmentIoCardBaseAllOfWithDefaults instantiates a new EquipmentIoCardBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *EquipmentIoCardBaseAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *EquipmentIoCardBaseAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *EquipmentIoCardBaseAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *EquipmentIoCardBaseAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDescription

`func (o *EquipmentIoCardBaseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquipmentIoCardBaseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquipmentIoCardBaseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquipmentIoCardBaseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentIoCardBaseAllOf) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentIoCardBaseAllOf) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentIoCardBaseAllOf) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentIoCardBaseAllOf) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentIoCardBaseAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentIoCardBaseAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentIoCardBaseAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentIoCardBaseAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentIoCardBaseAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentIoCardBaseAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentIoCardBaseAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentIoCardBaseAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPid

`func (o *EquipmentIoCardBaseAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *EquipmentIoCardBaseAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *EquipmentIoCardBaseAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *EquipmentIoCardBaseAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPresence

`func (o *EquipmentIoCardBaseAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *EquipmentIoCardBaseAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *EquipmentIoCardBaseAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *EquipmentIoCardBaseAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetProductName

`func (o *EquipmentIoCardBaseAllOf) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EquipmentIoCardBaseAllOf) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EquipmentIoCardBaseAllOf) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *EquipmentIoCardBaseAllOf) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSku

`func (o *EquipmentIoCardBaseAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EquipmentIoCardBaseAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EquipmentIoCardBaseAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *EquipmentIoCardBaseAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVersion

`func (o *EquipmentIoCardBaseAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EquipmentIoCardBaseAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EquipmentIoCardBaseAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EquipmentIoCardBaseAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentIoCardBaseAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentIoCardBaseAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentIoCardBaseAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentIoCardBaseAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetHostPorts

`func (o *EquipmentIoCardBaseAllOf) GetHostPorts() []EtherHostPortRelationship`

GetHostPorts returns the HostPorts field if non-nil, zero value otherwise.

### GetHostPortsOk

`func (o *EquipmentIoCardBaseAllOf) GetHostPortsOk() (*[]EtherHostPortRelationship, bool)`

GetHostPortsOk returns a tuple with the HostPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPorts

`func (o *EquipmentIoCardBaseAllOf) SetHostPorts(v []EtherHostPortRelationship)`

SetHostPorts sets HostPorts field to given value.

### HasHostPorts

`func (o *EquipmentIoCardBaseAllOf) HasHostPorts() bool`

HasHostPorts returns a boolean if a field has been set.

### SetHostPortsNil

`func (o *EquipmentIoCardBaseAllOf) SetHostPortsNil(b bool)`

 SetHostPortsNil sets the value for HostPorts to be an explicit nil

### UnsetHostPorts
`func (o *EquipmentIoCardBaseAllOf) UnsetHostPorts()`

UnsetHostPorts ensures that no value is present for HostPorts, not even an explicit nil
### GetMgmtController

`func (o *EquipmentIoCardBaseAllOf) GetMgmtController() ManagementControllerRelationship`

GetMgmtController returns the MgmtController field if non-nil, zero value otherwise.

### GetMgmtControllerOk

`func (o *EquipmentIoCardBaseAllOf) GetMgmtControllerOk() (*ManagementControllerRelationship, bool)`

GetMgmtControllerOk returns a tuple with the MgmtController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtController

`func (o *EquipmentIoCardBaseAllOf) SetMgmtController(v ManagementControllerRelationship)`

SetMgmtController sets MgmtController field to given value.

### HasMgmtController

`func (o *EquipmentIoCardBaseAllOf) HasMgmtController() bool`

HasMgmtController returns a boolean if a field has been set.

### GetNetworkPorts

`func (o *EquipmentIoCardBaseAllOf) GetNetworkPorts() []EtherNetworkPortRelationship`

GetNetworkPorts returns the NetworkPorts field if non-nil, zero value otherwise.

### GetNetworkPortsOk

`func (o *EquipmentIoCardBaseAllOf) GetNetworkPortsOk() (*[]EtherNetworkPortRelationship, bool)`

GetNetworkPortsOk returns a tuple with the NetworkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPorts

`func (o *EquipmentIoCardBaseAllOf) SetNetworkPorts(v []EtherNetworkPortRelationship)`

SetNetworkPorts sets NetworkPorts field to given value.

### HasNetworkPorts

`func (o *EquipmentIoCardBaseAllOf) HasNetworkPorts() bool`

HasNetworkPorts returns a boolean if a field has been set.

### SetNetworkPortsNil

`func (o *EquipmentIoCardBaseAllOf) SetNetworkPortsNil(b bool)`

 SetNetworkPortsNil sets the value for NetworkPorts to be an explicit nil

### UnsetNetworkPorts
`func (o *EquipmentIoCardBaseAllOf) UnsetNetworkPorts()`

UnsetNetworkPorts ensures that no value is present for NetworkPorts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


