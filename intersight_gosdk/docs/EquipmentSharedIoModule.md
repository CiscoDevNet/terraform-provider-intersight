# EquipmentSharedIoModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigState** | Pointer to **string** | This field identifies the configuration state for this SIOM Unit. | [optional] [readonly] 
**Discovery** | Pointer to **string** | This field identifies the discovery state of SIOM. | [optional] [readonly] 
**MacOfSharedIomAside** | Pointer to **string** | This field identifies the MAC of IOM-A side. | [optional] [readonly] 
**MacOfSharedIomBside** | Pointer to **string** | This field identifies the MAC of IOM-B side. | [optional] [readonly] 
**OperState** | Pointer to **string** | This field identifies the SIOM operational state. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | This field identifies the Part Number for this SIOM Unit. | [optional] [readonly] 
**Reachability** | Pointer to **string** | This field identifies the reachability to FI-A and B side. | [optional] [readonly] 
**UsrLbl** | Pointer to **string** | User label configured for the SIOM. | [optional] [readonly] 
**Vid** | Pointer to **string** | This field identifies the vendor id for this SIOM Unit. | [optional] [readonly] 
**Controller** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**EquipmentSystemIoController** | Pointer to [**EquipmentSystemIoControllerRelationship**](equipment.SystemIoController.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PortGroups** | Pointer to [**[]PortGroupRelationship**](port.Group.Relationship.md) | An array of relationships to portGroup resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentSharedIoModule

`func NewEquipmentSharedIoModule() *EquipmentSharedIoModule`

NewEquipmentSharedIoModule instantiates a new EquipmentSharedIoModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSharedIoModuleWithDefaults

`func NewEquipmentSharedIoModuleWithDefaults() *EquipmentSharedIoModule`

NewEquipmentSharedIoModuleWithDefaults instantiates a new EquipmentSharedIoModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigState

`func (o *EquipmentSharedIoModule) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentSharedIoModule) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentSharedIoModule) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentSharedIoModule) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetDiscovery

`func (o *EquipmentSharedIoModule) GetDiscovery() string`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *EquipmentSharedIoModule) GetDiscoveryOk() (*string, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *EquipmentSharedIoModule) SetDiscovery(v string)`

SetDiscovery sets Discovery field to given value.

### HasDiscovery

`func (o *EquipmentSharedIoModule) HasDiscovery() bool`

HasDiscovery returns a boolean if a field has been set.

### GetMacOfSharedIomAside

`func (o *EquipmentSharedIoModule) GetMacOfSharedIomAside() string`

GetMacOfSharedIomAside returns the MacOfSharedIomAside field if non-nil, zero value otherwise.

### GetMacOfSharedIomAsideOk

`func (o *EquipmentSharedIoModule) GetMacOfSharedIomAsideOk() (*string, bool)`

GetMacOfSharedIomAsideOk returns a tuple with the MacOfSharedIomAside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOfSharedIomAside

`func (o *EquipmentSharedIoModule) SetMacOfSharedIomAside(v string)`

SetMacOfSharedIomAside sets MacOfSharedIomAside field to given value.

### HasMacOfSharedIomAside

`func (o *EquipmentSharedIoModule) HasMacOfSharedIomAside() bool`

HasMacOfSharedIomAside returns a boolean if a field has been set.

### GetMacOfSharedIomBside

`func (o *EquipmentSharedIoModule) GetMacOfSharedIomBside() string`

GetMacOfSharedIomBside returns the MacOfSharedIomBside field if non-nil, zero value otherwise.

### GetMacOfSharedIomBsideOk

`func (o *EquipmentSharedIoModule) GetMacOfSharedIomBsideOk() (*string, bool)`

GetMacOfSharedIomBsideOk returns a tuple with the MacOfSharedIomBside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOfSharedIomBside

`func (o *EquipmentSharedIoModule) SetMacOfSharedIomBside(v string)`

SetMacOfSharedIomBside sets MacOfSharedIomBside field to given value.

### HasMacOfSharedIomBside

`func (o *EquipmentSharedIoModule) HasMacOfSharedIomBside() bool`

HasMacOfSharedIomBside returns a boolean if a field has been set.

### GetOperState

`func (o *EquipmentSharedIoModule) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EquipmentSharedIoModule) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EquipmentSharedIoModule) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EquipmentSharedIoModule) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPartNumber

`func (o *EquipmentSharedIoModule) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EquipmentSharedIoModule) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EquipmentSharedIoModule) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EquipmentSharedIoModule) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetReachability

`func (o *EquipmentSharedIoModule) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *EquipmentSharedIoModule) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *EquipmentSharedIoModule) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *EquipmentSharedIoModule) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetUsrLbl

`func (o *EquipmentSharedIoModule) GetUsrLbl() string`

GetUsrLbl returns the UsrLbl field if non-nil, zero value otherwise.

### GetUsrLblOk

`func (o *EquipmentSharedIoModule) GetUsrLblOk() (*string, bool)`

GetUsrLblOk returns a tuple with the UsrLbl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsrLbl

`func (o *EquipmentSharedIoModule) SetUsrLbl(v string)`

SetUsrLbl sets UsrLbl field to given value.

### HasUsrLbl

`func (o *EquipmentSharedIoModule) HasUsrLbl() bool`

HasUsrLbl returns a boolean if a field has been set.

### GetVid

`func (o *EquipmentSharedIoModule) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *EquipmentSharedIoModule) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *EquipmentSharedIoModule) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *EquipmentSharedIoModule) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetController

`func (o *EquipmentSharedIoModule) GetController() ManagementControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *EquipmentSharedIoModule) GetControllerOk() (*ManagementControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *EquipmentSharedIoModule) SetController(v ManagementControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *EquipmentSharedIoModule) HasController() bool`

HasController returns a boolean if a field has been set.

### GetEquipmentSystemIoController

`func (o *EquipmentSharedIoModule) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship`

GetEquipmentSystemIoController returns the EquipmentSystemIoController field if non-nil, zero value otherwise.

### GetEquipmentSystemIoControllerOk

`func (o *EquipmentSharedIoModule) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool)`

GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSystemIoController

`func (o *EquipmentSharedIoModule) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship)`

SetEquipmentSystemIoController sets EquipmentSystemIoController field to given value.

### HasEquipmentSystemIoController

`func (o *EquipmentSharedIoModule) HasEquipmentSystemIoController() bool`

HasEquipmentSystemIoController returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EquipmentSharedIoModule) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EquipmentSharedIoModule) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EquipmentSharedIoModule) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EquipmentSharedIoModule) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroups

`func (o *EquipmentSharedIoModule) GetPortGroups() []PortGroupRelationship`

GetPortGroups returns the PortGroups field if non-nil, zero value otherwise.

### GetPortGroupsOk

`func (o *EquipmentSharedIoModule) GetPortGroupsOk() (*[]PortGroupRelationship, bool)`

GetPortGroupsOk returns a tuple with the PortGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroups

`func (o *EquipmentSharedIoModule) SetPortGroups(v []PortGroupRelationship)`

SetPortGroups sets PortGroups field to given value.

### HasPortGroups

`func (o *EquipmentSharedIoModule) HasPortGroups() bool`

HasPortGroups returns a boolean if a field has been set.

### SetPortGroupsNil

`func (o *EquipmentSharedIoModule) SetPortGroupsNil(b bool)`

 SetPortGroupsNil sets the value for PortGroups to be an explicit nil

### UnsetPortGroups
`func (o *EquipmentSharedIoModule) UnsetPortGroups()`

UnsetPortGroups ensures that no value is present for PortGroups, not even an explicit nil
### GetRegisteredDevice

`func (o *EquipmentSharedIoModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EquipmentSharedIoModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EquipmentSharedIoModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EquipmentSharedIoModule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


