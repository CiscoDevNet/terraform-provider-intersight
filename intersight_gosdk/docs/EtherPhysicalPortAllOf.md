# EtherPhysicalPortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSpeed** | Pointer to **string** | Administratively configured speed for this port. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port. | [optional] [readonly] 
**AggregatePortId** | Pointer to **int64** | Breakout port member in the Fabric Interconnect. | [optional] [readonly] 
**LicenseGrace** | Pointer to **string** | The number of days this port&#39;s license has been in Grace Period for. | [optional] [readonly] 
**LicenseState** | Pointer to **string** | The state of the port&#39;s licensing. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PortGroup** | Pointer to [**PortGroupRelationship**](port.Group.Relationship.md) |  | [optional] 
**PortSubGroup** | Pointer to [**PortSubGroupRelationship**](port.SubGroup.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEtherPhysicalPortAllOf

`func NewEtherPhysicalPortAllOf() *EtherPhysicalPortAllOf`

NewEtherPhysicalPortAllOf instantiates a new EtherPhysicalPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPhysicalPortAllOfWithDefaults

`func NewEtherPhysicalPortAllOfWithDefaults() *EtherPhysicalPortAllOf`

NewEtherPhysicalPortAllOfWithDefaults instantiates a new EtherPhysicalPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSpeed

`func (o *EtherPhysicalPortAllOf) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *EtherPhysicalPortAllOf) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *EtherPhysicalPortAllOf) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *EtherPhysicalPortAllOf) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetAdminState

`func (o *EtherPhysicalPortAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EtherPhysicalPortAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EtherPhysicalPortAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EtherPhysicalPortAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAggregatePortId

`func (o *EtherPhysicalPortAllOf) GetAggregatePortId() int64`

GetAggregatePortId returns the AggregatePortId field if non-nil, zero value otherwise.

### GetAggregatePortIdOk

`func (o *EtherPhysicalPortAllOf) GetAggregatePortIdOk() (*int64, bool)`

GetAggregatePortIdOk returns a tuple with the AggregatePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatePortId

`func (o *EtherPhysicalPortAllOf) SetAggregatePortId(v int64)`

SetAggregatePortId sets AggregatePortId field to given value.

### HasAggregatePortId

`func (o *EtherPhysicalPortAllOf) HasAggregatePortId() bool`

HasAggregatePortId returns a boolean if a field has been set.

### GetLicenseGrace

`func (o *EtherPhysicalPortAllOf) GetLicenseGrace() string`

GetLicenseGrace returns the LicenseGrace field if non-nil, zero value otherwise.

### GetLicenseGraceOk

`func (o *EtherPhysicalPortAllOf) GetLicenseGraceOk() (*string, bool)`

GetLicenseGraceOk returns a tuple with the LicenseGrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGrace

`func (o *EtherPhysicalPortAllOf) SetLicenseGrace(v string)`

SetLicenseGrace sets LicenseGrace field to given value.

### HasLicenseGrace

`func (o *EtherPhysicalPortAllOf) HasLicenseGrace() bool`

HasLicenseGrace returns a boolean if a field has been set.

### GetLicenseState

`func (o *EtherPhysicalPortAllOf) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *EtherPhysicalPortAllOf) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *EtherPhysicalPortAllOf) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *EtherPhysicalPortAllOf) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *EtherPhysicalPortAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *EtherPhysicalPortAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *EtherPhysicalPortAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *EtherPhysicalPortAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPortGroup

`func (o *EtherPhysicalPortAllOf) GetPortGroup() PortGroupRelationship`

GetPortGroup returns the PortGroup field if non-nil, zero value otherwise.

### GetPortGroupOk

`func (o *EtherPhysicalPortAllOf) GetPortGroupOk() (*PortGroupRelationship, bool)`

GetPortGroupOk returns a tuple with the PortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroup

`func (o *EtherPhysicalPortAllOf) SetPortGroup(v PortGroupRelationship)`

SetPortGroup sets PortGroup field to given value.

### HasPortGroup

`func (o *EtherPhysicalPortAllOf) HasPortGroup() bool`

HasPortGroup returns a boolean if a field has been set.

### GetPortSubGroup

`func (o *EtherPhysicalPortAllOf) GetPortSubGroup() PortSubGroupRelationship`

GetPortSubGroup returns the PortSubGroup field if non-nil, zero value otherwise.

### GetPortSubGroupOk

`func (o *EtherPhysicalPortAllOf) GetPortSubGroupOk() (*PortSubGroupRelationship, bool)`

GetPortSubGroupOk returns a tuple with the PortSubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSubGroup

`func (o *EtherPhysicalPortAllOf) SetPortSubGroup(v PortSubGroupRelationship)`

SetPortSubGroup sets PortSubGroup field to given value.

### HasPortSubGroup

`func (o *EtherPhysicalPortAllOf) HasPortSubGroup() bool`

HasPortSubGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EtherPhysicalPortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherPhysicalPortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherPhysicalPortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherPhysicalPortAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


