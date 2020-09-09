# AdapterUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdapterId** | Pointer to **string** | Unique Identifier of an adapter Unit within a Rack Interface. | [optional] [readonly] 
**BaseMacAddress** | Pointer to **string** | Original Base Mac address of an adapter unit. | [optional] [readonly] 
**ConnectionStatus** | Pointer to **string** | Connectivity Status of adapter - A or B or AB. | [optional] [readonly] 
**Integrated** | Pointer to **string** | Cisco Integrated adapter or other type. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of an adapter unit. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability state of an adapter unit. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Part number of an adapter unit. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | PCIe slot of the adapter in the server. | [optional] [readonly] 
**Power** | Pointer to **string** | Power state of an adapter unit. | [optional] [readonly] 
**Presence** | Pointer to **string** | Adapter Unit presence or absence. | [optional] [readonly] 
**Thermal** | Pointer to **string** | Thermal state of an adapter unit. | [optional] [readonly] 
**Vid** | Pointer to **string** | Virtual Id of the adapter in the server. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**Controller** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**ExtEthIfs** | Pointer to [**[]AdapterExtEthInterfaceRelationship**](adapter.ExtEthInterface.Relationship.md) | An array of relationships to adapterExtEthInterface resources. | [optional] [readonly] 
**HostEthIfs** | Pointer to [**[]AdapterHostEthInterfaceRelationship**](adapter.HostEthInterface.Relationship.md) | An array of relationships to adapterHostEthInterface resources. | [optional] [readonly] 
**HostFcIfs** | Pointer to [**[]AdapterHostFcInterfaceRelationship**](adapter.HostFcInterface.Relationship.md) | An array of relationships to adapterHostFcInterface resources. | [optional] [readonly] 
**HostIscsiIfs** | Pointer to [**[]AdapterHostIscsiInterfaceRelationship**](adapter.HostIscsiInterface.Relationship.md) | An array of relationships to adapterHostIscsiInterface resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterUnitAllOf

`func NewAdapterUnitAllOf() *AdapterUnitAllOf`

NewAdapterUnitAllOf instantiates a new AdapterUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterUnitAllOfWithDefaults

`func NewAdapterUnitAllOfWithDefaults() *AdapterUnitAllOf`

NewAdapterUnitAllOfWithDefaults instantiates a new AdapterUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapterId

`func (o *AdapterUnitAllOf) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *AdapterUnitAllOf) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *AdapterUnitAllOf) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *AdapterUnitAllOf) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### GetBaseMacAddress

`func (o *AdapterUnitAllOf) GetBaseMacAddress() string`

GetBaseMacAddress returns the BaseMacAddress field if non-nil, zero value otherwise.

### GetBaseMacAddressOk

`func (o *AdapterUnitAllOf) GetBaseMacAddressOk() (*string, bool)`

GetBaseMacAddressOk returns a tuple with the BaseMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMacAddress

`func (o *AdapterUnitAllOf) SetBaseMacAddress(v string)`

SetBaseMacAddress sets BaseMacAddress field to given value.

### HasBaseMacAddress

`func (o *AdapterUnitAllOf) HasBaseMacAddress() bool`

HasBaseMacAddress returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AdapterUnitAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AdapterUnitAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AdapterUnitAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AdapterUnitAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetIntegrated

`func (o *AdapterUnitAllOf) GetIntegrated() string`

GetIntegrated returns the Integrated field if non-nil, zero value otherwise.

### GetIntegratedOk

`func (o *AdapterUnitAllOf) GetIntegratedOk() (*string, bool)`

GetIntegratedOk returns a tuple with the Integrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrated

`func (o *AdapterUnitAllOf) SetIntegrated(v string)`

SetIntegrated sets Integrated field to given value.

### HasIntegrated

`func (o *AdapterUnitAllOf) HasIntegrated() bool`

HasIntegrated returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterUnitAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterUnitAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterUnitAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterUnitAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterUnitAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterUnitAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterUnitAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterUnitAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPartNumber

`func (o *AdapterUnitAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *AdapterUnitAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *AdapterUnitAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *AdapterUnitAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPciSlot

`func (o *AdapterUnitAllOf) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *AdapterUnitAllOf) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *AdapterUnitAllOf) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *AdapterUnitAllOf) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPower

`func (o *AdapterUnitAllOf) GetPower() string`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *AdapterUnitAllOf) GetPowerOk() (*string, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *AdapterUnitAllOf) SetPower(v string)`

SetPower sets Power field to given value.

### HasPower

`func (o *AdapterUnitAllOf) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPresence

`func (o *AdapterUnitAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *AdapterUnitAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *AdapterUnitAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *AdapterUnitAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetThermal

`func (o *AdapterUnitAllOf) GetThermal() string`

GetThermal returns the Thermal field if non-nil, zero value otherwise.

### GetThermalOk

`func (o *AdapterUnitAllOf) GetThermalOk() (*string, bool)`

GetThermalOk returns a tuple with the Thermal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermal

`func (o *AdapterUnitAllOf) SetThermal(v string)`

SetThermal sets Thermal field to given value.

### HasThermal

`func (o *AdapterUnitAllOf) HasThermal() bool`

HasThermal returns a boolean if a field has been set.

### GetVid

`func (o *AdapterUnitAllOf) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *AdapterUnitAllOf) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *AdapterUnitAllOf) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *AdapterUnitAllOf) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetComputeBlade

`func (o *AdapterUnitAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *AdapterUnitAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *AdapterUnitAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *AdapterUnitAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *AdapterUnitAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *AdapterUnitAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *AdapterUnitAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *AdapterUnitAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetController

`func (o *AdapterUnitAllOf) GetController() ManagementControllerRelationship`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AdapterUnitAllOf) GetControllerOk() (*ManagementControllerRelationship, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AdapterUnitAllOf) SetController(v ManagementControllerRelationship)`

SetController sets Controller field to given value.

### HasController

`func (o *AdapterUnitAllOf) HasController() bool`

HasController returns a boolean if a field has been set.

### GetExtEthIfs

`func (o *AdapterUnitAllOf) GetExtEthIfs() []AdapterExtEthInterfaceRelationship`

GetExtEthIfs returns the ExtEthIfs field if non-nil, zero value otherwise.

### GetExtEthIfsOk

`func (o *AdapterUnitAllOf) GetExtEthIfsOk() (*[]AdapterExtEthInterfaceRelationship, bool)`

GetExtEthIfsOk returns a tuple with the ExtEthIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtEthIfs

`func (o *AdapterUnitAllOf) SetExtEthIfs(v []AdapterExtEthInterfaceRelationship)`

SetExtEthIfs sets ExtEthIfs field to given value.

### HasExtEthIfs

`func (o *AdapterUnitAllOf) HasExtEthIfs() bool`

HasExtEthIfs returns a boolean if a field has been set.

### SetExtEthIfsNil

`func (o *AdapterUnitAllOf) SetExtEthIfsNil(b bool)`

 SetExtEthIfsNil sets the value for ExtEthIfs to be an explicit nil

### UnsetExtEthIfs
`func (o *AdapterUnitAllOf) UnsetExtEthIfs()`

UnsetExtEthIfs ensures that no value is present for ExtEthIfs, not even an explicit nil
### GetHostEthIfs

`func (o *AdapterUnitAllOf) GetHostEthIfs() []AdapterHostEthInterfaceRelationship`

GetHostEthIfs returns the HostEthIfs field if non-nil, zero value otherwise.

### GetHostEthIfsOk

`func (o *AdapterUnitAllOf) GetHostEthIfsOk() (*[]AdapterHostEthInterfaceRelationship, bool)`

GetHostEthIfsOk returns a tuple with the HostEthIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEthIfs

`func (o *AdapterUnitAllOf) SetHostEthIfs(v []AdapterHostEthInterfaceRelationship)`

SetHostEthIfs sets HostEthIfs field to given value.

### HasHostEthIfs

`func (o *AdapterUnitAllOf) HasHostEthIfs() bool`

HasHostEthIfs returns a boolean if a field has been set.

### SetHostEthIfsNil

`func (o *AdapterUnitAllOf) SetHostEthIfsNil(b bool)`

 SetHostEthIfsNil sets the value for HostEthIfs to be an explicit nil

### UnsetHostEthIfs
`func (o *AdapterUnitAllOf) UnsetHostEthIfs()`

UnsetHostEthIfs ensures that no value is present for HostEthIfs, not even an explicit nil
### GetHostFcIfs

`func (o *AdapterUnitAllOf) GetHostFcIfs() []AdapterHostFcInterfaceRelationship`

GetHostFcIfs returns the HostFcIfs field if non-nil, zero value otherwise.

### GetHostFcIfsOk

`func (o *AdapterUnitAllOf) GetHostFcIfsOk() (*[]AdapterHostFcInterfaceRelationship, bool)`

GetHostFcIfsOk returns a tuple with the HostFcIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFcIfs

`func (o *AdapterUnitAllOf) SetHostFcIfs(v []AdapterHostFcInterfaceRelationship)`

SetHostFcIfs sets HostFcIfs field to given value.

### HasHostFcIfs

`func (o *AdapterUnitAllOf) HasHostFcIfs() bool`

HasHostFcIfs returns a boolean if a field has been set.

### SetHostFcIfsNil

`func (o *AdapterUnitAllOf) SetHostFcIfsNil(b bool)`

 SetHostFcIfsNil sets the value for HostFcIfs to be an explicit nil

### UnsetHostFcIfs
`func (o *AdapterUnitAllOf) UnsetHostFcIfs()`

UnsetHostFcIfs ensures that no value is present for HostFcIfs, not even an explicit nil
### GetHostIscsiIfs

`func (o *AdapterUnitAllOf) GetHostIscsiIfs() []AdapterHostIscsiInterfaceRelationship`

GetHostIscsiIfs returns the HostIscsiIfs field if non-nil, zero value otherwise.

### GetHostIscsiIfsOk

`func (o *AdapterUnitAllOf) GetHostIscsiIfsOk() (*[]AdapterHostIscsiInterfaceRelationship, bool)`

GetHostIscsiIfsOk returns a tuple with the HostIscsiIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIscsiIfs

`func (o *AdapterUnitAllOf) SetHostIscsiIfs(v []AdapterHostIscsiInterfaceRelationship)`

SetHostIscsiIfs sets HostIscsiIfs field to given value.

### HasHostIscsiIfs

`func (o *AdapterUnitAllOf) HasHostIscsiIfs() bool`

HasHostIscsiIfs returns a boolean if a field has been set.

### SetHostIscsiIfsNil

`func (o *AdapterUnitAllOf) SetHostIscsiIfsNil(b bool)`

 SetHostIscsiIfsNil sets the value for HostIscsiIfs to be an explicit nil

### UnsetHostIscsiIfs
`func (o *AdapterUnitAllOf) UnsetHostIscsiIfs()`

UnsetHostIscsiIfs ensures that no value is present for HostIscsiIfs, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *AdapterUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterUnitAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


