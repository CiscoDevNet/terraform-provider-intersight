# FirmwareSwitchUpgradeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](network.Element.Relationship.md) | An array of relationships to networkElement resources. | [optional] 

## Methods

### NewFirmwareSwitchUpgradeAllOf

`func NewFirmwareSwitchUpgradeAllOf() *FirmwareSwitchUpgradeAllOf`

NewFirmwareSwitchUpgradeAllOf instantiates a new FirmwareSwitchUpgradeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareSwitchUpgradeAllOfWithDefaults

`func NewFirmwareSwitchUpgradeAllOfWithDefaults() *FirmwareSwitchUpgradeAllOf`

NewFirmwareSwitchUpgradeAllOfWithDefaults instantiates a new FirmwareSwitchUpgradeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *FirmwareSwitchUpgradeAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareSwitchUpgradeAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareSwitchUpgradeAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareSwitchUpgradeAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetNetworkElements

`func (o *FirmwareSwitchUpgradeAllOf) GetNetworkElements() []NetworkElementRelationship`

GetNetworkElements returns the NetworkElements field if non-nil, zero value otherwise.

### GetNetworkElementsOk

`func (o *FirmwareSwitchUpgradeAllOf) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool)`

GetNetworkElementsOk returns a tuple with the NetworkElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElements

`func (o *FirmwareSwitchUpgradeAllOf) SetNetworkElements(v []NetworkElementRelationship)`

SetNetworkElements sets NetworkElements field to given value.

### HasNetworkElements

`func (o *FirmwareSwitchUpgradeAllOf) HasNetworkElements() bool`

HasNetworkElements returns a boolean if a field has been set.

### SetNetworkElementsNil

`func (o *FirmwareSwitchUpgradeAllOf) SetNetworkElementsNil(b bool)`

 SetNetworkElementsNil sets the value for NetworkElements to be an explicit nil

### UnsetNetworkElements
`func (o *FirmwareSwitchUpgradeAllOf) UnsetNetworkElements()`

UnsetNetworkElements ensures that no value is present for NetworkElements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


