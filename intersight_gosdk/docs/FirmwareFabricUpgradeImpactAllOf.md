# FirmwareFabricUpgradeImpactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImpactDetail** | Pointer to [**[]FirmwareComponentImpact**](firmware.ComponentImpact.md) |  | [optional] 
**Serial** | Pointer to **string** | Details for the Fabric Interconnect that will be impacted by the upgrade. | [optional] 

## Methods

### NewFirmwareFabricUpgradeImpactAllOf

`func NewFirmwareFabricUpgradeImpactAllOf() *FirmwareFabricUpgradeImpactAllOf`

NewFirmwareFabricUpgradeImpactAllOf instantiates a new FirmwareFabricUpgradeImpactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareFabricUpgradeImpactAllOfWithDefaults

`func NewFirmwareFabricUpgradeImpactAllOfWithDefaults() *FirmwareFabricUpgradeImpactAllOf`

NewFirmwareFabricUpgradeImpactAllOfWithDefaults instantiates a new FirmwareFabricUpgradeImpactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpactDetail

`func (o *FirmwareFabricUpgradeImpactAllOf) GetImpactDetail() []FirmwareComponentImpact`

GetImpactDetail returns the ImpactDetail field if non-nil, zero value otherwise.

### GetImpactDetailOk

`func (o *FirmwareFabricUpgradeImpactAllOf) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool)`

GetImpactDetailOk returns a tuple with the ImpactDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDetail

`func (o *FirmwareFabricUpgradeImpactAllOf) SetImpactDetail(v []FirmwareComponentImpact)`

SetImpactDetail sets ImpactDetail field to given value.

### HasImpactDetail

`func (o *FirmwareFabricUpgradeImpactAllOf) HasImpactDetail() bool`

HasImpactDetail returns a boolean if a field has been set.

### GetSerial

`func (o *FirmwareFabricUpgradeImpactAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *FirmwareFabricUpgradeImpactAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *FirmwareFabricUpgradeImpactAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *FirmwareFabricUpgradeImpactAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


