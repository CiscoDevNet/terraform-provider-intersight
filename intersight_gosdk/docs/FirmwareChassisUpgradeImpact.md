# FirmwareChassisUpgradeImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImpactDetail** | Pointer to [**[]FirmwareComponentImpact**](firmware.ComponentImpact.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the chassis that will be affected by the upgrade. | [optional] 
**UserLabel** | Pointer to **string** | Details for the chassis that will be impacted by the upgrade. | [optional] 

## Methods

### NewFirmwareChassisUpgradeImpact

`func NewFirmwareChassisUpgradeImpact() *FirmwareChassisUpgradeImpact`

NewFirmwareChassisUpgradeImpact instantiates a new FirmwareChassisUpgradeImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareChassisUpgradeImpactWithDefaults

`func NewFirmwareChassisUpgradeImpactWithDefaults() *FirmwareChassisUpgradeImpact`

NewFirmwareChassisUpgradeImpactWithDefaults instantiates a new FirmwareChassisUpgradeImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpactDetail

`func (o *FirmwareChassisUpgradeImpact) GetImpactDetail() []FirmwareComponentImpact`

GetImpactDetail returns the ImpactDetail field if non-nil, zero value otherwise.

### GetImpactDetailOk

`func (o *FirmwareChassisUpgradeImpact) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool)`

GetImpactDetailOk returns a tuple with the ImpactDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDetail

`func (o *FirmwareChassisUpgradeImpact) SetImpactDetail(v []FirmwareComponentImpact)`

SetImpactDetail sets ImpactDetail field to given value.

### HasImpactDetail

`func (o *FirmwareChassisUpgradeImpact) HasImpactDetail() bool`

HasImpactDetail returns a boolean if a field has been set.

### GetName

`func (o *FirmwareChassisUpgradeImpact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareChassisUpgradeImpact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareChassisUpgradeImpact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirmwareChassisUpgradeImpact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserLabel

`func (o *FirmwareChassisUpgradeImpact) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *FirmwareChassisUpgradeImpact) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *FirmwareChassisUpgradeImpact) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *FirmwareChassisUpgradeImpact) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


