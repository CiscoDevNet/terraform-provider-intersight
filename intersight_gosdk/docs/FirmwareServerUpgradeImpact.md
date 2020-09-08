# FirmwareServerUpgradeImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImpactDetail** | Pointer to [**[]FirmwareComponentImpact**](firmware.ComponentImpact.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the server impacted by the upgrade. | [optional] 
**UserLabel** | Pointer to **string** | Details about the server which will be impacted by the upgrade. | [optional] 

## Methods

### NewFirmwareServerUpgradeImpact

`func NewFirmwareServerUpgradeImpact() *FirmwareServerUpgradeImpact`

NewFirmwareServerUpgradeImpact instantiates a new FirmwareServerUpgradeImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareServerUpgradeImpactWithDefaults

`func NewFirmwareServerUpgradeImpactWithDefaults() *FirmwareServerUpgradeImpact`

NewFirmwareServerUpgradeImpactWithDefaults instantiates a new FirmwareServerUpgradeImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpactDetail

`func (o *FirmwareServerUpgradeImpact) GetImpactDetail() []FirmwareComponentImpact`

GetImpactDetail returns the ImpactDetail field if non-nil, zero value otherwise.

### GetImpactDetailOk

`func (o *FirmwareServerUpgradeImpact) GetImpactDetailOk() (*[]FirmwareComponentImpact, bool)`

GetImpactDetailOk returns a tuple with the ImpactDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDetail

`func (o *FirmwareServerUpgradeImpact) SetImpactDetail(v []FirmwareComponentImpact)`

SetImpactDetail sets ImpactDetail field to given value.

### HasImpactDetail

`func (o *FirmwareServerUpgradeImpact) HasImpactDetail() bool`

HasImpactDetail returns a boolean if a field has been set.

### GetName

`func (o *FirmwareServerUpgradeImpact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareServerUpgradeImpact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareServerUpgradeImpact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirmwareServerUpgradeImpact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserLabel

`func (o *FirmwareServerUpgradeImpact) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *FirmwareServerUpgradeImpact) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *FirmwareServerUpgradeImpact) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *FirmwareServerUpgradeImpact) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


