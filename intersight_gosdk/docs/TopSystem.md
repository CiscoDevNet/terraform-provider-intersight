# TopSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "top.System"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "top.System"]
**Ipv4Address** | Pointer to **string** | The IPv4 address of system. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | The IPv6 address of system. | [optional] [readonly] 
**Mode** | Pointer to **string** | The current mode of the system. | [optional] [readonly] 
**Name** | Pointer to **string** | The admin configured name of the system. | [optional] [readonly] 
**TimeZone** | Pointer to **string** | The operational timezone of the system, empty indicates no timezone has been set specifically. | [optional] 
**ComputeBlades** | Pointer to [**[]ComputeBladeRelationship**](ComputeBladeRelationship.md) | An array of relationships to computeBlade resources. | [optional] 
**ComputeRackUnits** | Pointer to [**[]ComputeRackUnitRelationship**](ComputeRackUnitRelationship.md) | An array of relationships to computeRackUnit resources. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**ManagementController** | Pointer to [**ManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](NetworkElementRelationship.md) | An array of relationships to networkElement resources. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewTopSystem

`func NewTopSystem(classId string, objectType string, ) *TopSystem`

NewTopSystem instantiates a new TopSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopSystemWithDefaults

`func NewTopSystemWithDefaults() *TopSystem`

NewTopSystemWithDefaults instantiates a new TopSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TopSystem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TopSystem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TopSystem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TopSystem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TopSystem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TopSystem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpv4Address

`func (o *TopSystem) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *TopSystem) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *TopSystem) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *TopSystem) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *TopSystem) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *TopSystem) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *TopSystem) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *TopSystem) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetMode

`func (o *TopSystem) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TopSystem) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TopSystem) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TopSystem) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *TopSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *TopSystem) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TopSystem) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TopSystem) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TopSystem) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetComputeBlades

`func (o *TopSystem) GetComputeBlades() []ComputeBladeRelationship`

GetComputeBlades returns the ComputeBlades field if non-nil, zero value otherwise.

### GetComputeBladesOk

`func (o *TopSystem) GetComputeBladesOk() (*[]ComputeBladeRelationship, bool)`

GetComputeBladesOk returns a tuple with the ComputeBlades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlades

`func (o *TopSystem) SetComputeBlades(v []ComputeBladeRelationship)`

SetComputeBlades sets ComputeBlades field to given value.

### HasComputeBlades

`func (o *TopSystem) HasComputeBlades() bool`

HasComputeBlades returns a boolean if a field has been set.

### SetComputeBladesNil

`func (o *TopSystem) SetComputeBladesNil(b bool)`

 SetComputeBladesNil sets the value for ComputeBlades to be an explicit nil

### UnsetComputeBlades
`func (o *TopSystem) UnsetComputeBlades()`

UnsetComputeBlades ensures that no value is present for ComputeBlades, not even an explicit nil
### GetComputeRackUnits

`func (o *TopSystem) GetComputeRackUnits() []ComputeRackUnitRelationship`

GetComputeRackUnits returns the ComputeRackUnits field if non-nil, zero value otherwise.

### GetComputeRackUnitsOk

`func (o *TopSystem) GetComputeRackUnitsOk() (*[]ComputeRackUnitRelationship, bool)`

GetComputeRackUnitsOk returns a tuple with the ComputeRackUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnits

`func (o *TopSystem) SetComputeRackUnits(v []ComputeRackUnitRelationship)`

SetComputeRackUnits sets ComputeRackUnits field to given value.

### HasComputeRackUnits

`func (o *TopSystem) HasComputeRackUnits() bool`

HasComputeRackUnits returns a boolean if a field has been set.

### SetComputeRackUnitsNil

`func (o *TopSystem) SetComputeRackUnitsNil(b bool)`

 SetComputeRackUnitsNil sets the value for ComputeRackUnits to be an explicit nil

### UnsetComputeRackUnits
`func (o *TopSystem) UnsetComputeRackUnits()`

UnsetComputeRackUnits ensures that no value is present for ComputeRackUnits, not even an explicit nil
### GetInventoryDeviceInfo

`func (o *TopSystem) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *TopSystem) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *TopSystem) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *TopSystem) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementController

`func (o *TopSystem) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *TopSystem) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *TopSystem) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *TopSystem) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### GetNetworkElements

`func (o *TopSystem) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *TopSystem) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *TopSystem) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *TopSystem) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *TopSystem) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *TopSystem) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil
### GetRegisteredDevice

`func (o *TopSystem) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *TopSystem) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *TopSystem) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *TopSystem) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


