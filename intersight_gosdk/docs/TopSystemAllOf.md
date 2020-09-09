# TopSystemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Address** | Pointer to **string** | The IPv4 address of system. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | The IPv6 address of system. | [optional] [readonly] 
**Mode** | Pointer to **string** | The current mode of the system. | [optional] [readonly] 
**Name** | Pointer to **string** | The admin configured name of the system. | [optional] [readonly] 
**TimeZone** | Pointer to **string** | The operational timezone of the system, empty indicates no timezone has been set specifically. | [optional] 
**ComputeBlades** | Pointer to [**[]ComputeBladeRelationship**](compute.Blade.Relationship.md) | An array of relationships to computeBlade resources. | [optional] 
**ComputeRackUnits** | Pointer to [**[]ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) | An array of relationships to computeRackUnit resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**ManagementController** | Pointer to [**ManagementControllerRelationship**](management.Controller.Relationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](network.Element.Relationship.md) | An array of relationships to networkElement resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewTopSystemAllOf

`func NewTopSystemAllOf() *TopSystemAllOf`

NewTopSystemAllOf instantiates a new TopSystemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopSystemAllOfWithDefaults

`func NewTopSystemAllOfWithDefaults() *TopSystemAllOf`

NewTopSystemAllOfWithDefaults instantiates a new TopSystemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Address

`func (o *TopSystemAllOf) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *TopSystemAllOf) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *TopSystemAllOf) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *TopSystemAllOf) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *TopSystemAllOf) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *TopSystemAllOf) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *TopSystemAllOf) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *TopSystemAllOf) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetMode

`func (o *TopSystemAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TopSystemAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TopSystemAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TopSystemAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *TopSystemAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopSystemAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopSystemAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopSystemAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *TopSystemAllOf) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TopSystemAllOf) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TopSystemAllOf) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TopSystemAllOf) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetComputeBlades

`func (o *TopSystemAllOf) GetComputeBlades() []ComputeBladeRelationship`

GetComputeBlades returns the ComputeBlades field if non-nil, zero value otherwise.

### GetComputeBladesOk

`func (o *TopSystemAllOf) GetComputeBladesOk() (*[]ComputeBladeRelationship, bool)`

GetComputeBladesOk returns a tuple with the ComputeBlades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlades

`func (o *TopSystemAllOf) SetComputeBlades(v []ComputeBladeRelationship)`

SetComputeBlades sets ComputeBlades field to given value.

### HasComputeBlades

`func (o *TopSystemAllOf) HasComputeBlades() bool`

HasComputeBlades returns a boolean if a field has been set.

### SetComputeBladesNil

`func (o *TopSystemAllOf) SetComputeBladesNil(b bool)`

 SetComputeBladesNil sets the value for ComputeBlades to be an explicit nil

### UnsetComputeBlades
`func (o *TopSystemAllOf) UnsetComputeBlades()`

UnsetComputeBlades ensures that no value is present for ComputeBlades, not even an explicit nil
### GetComputeRackUnits

`func (o *TopSystemAllOf) GetComputeRackUnits() []ComputeRackUnitRelationship`

GetComputeRackUnits returns the ComputeRackUnits field if non-nil, zero value otherwise.

### GetComputeRackUnitsOk

`func (o *TopSystemAllOf) GetComputeRackUnitsOk() (*[]ComputeRackUnitRelationship, bool)`

GetComputeRackUnitsOk returns a tuple with the ComputeRackUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnits

`func (o *TopSystemAllOf) SetComputeRackUnits(v []ComputeRackUnitRelationship)`

SetComputeRackUnits sets ComputeRackUnits field to given value.

### HasComputeRackUnits

`func (o *TopSystemAllOf) HasComputeRackUnits() bool`

HasComputeRackUnits returns a boolean if a field has been set.

### SetComputeRackUnitsNil

`func (o *TopSystemAllOf) SetComputeRackUnitsNil(b bool)`

 SetComputeRackUnitsNil sets the value for ComputeRackUnits to be an explicit nil

### UnsetComputeRackUnits
`func (o *TopSystemAllOf) UnsetComputeRackUnits()`

UnsetComputeRackUnits ensures that no value is present for ComputeRackUnits, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *TopSystemAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *TopSystemAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *TopSystemAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *TopSystemAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementController

`func (o *TopSystemAllOf) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *TopSystemAllOf) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *TopSystemAllOf) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *TopSystemAllOf) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### GetNetworkElements

`func (o *TopSystemAllOf) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *TopSystemAllOf) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *TopSystemAllOf) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *TopSystemAllOf) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *TopSystemAllOf) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *TopSystemAllOf) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil
### GetRegisteredDevice

`func (o *TopSystemAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *TopSystemAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *TopSystemAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *TopSystemAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


