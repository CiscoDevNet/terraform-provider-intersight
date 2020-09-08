# FirmwareUpgradeImpactBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to **[]string** |  | [optional] 
**ComputationState** | Pointer to **string** | Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint. * &#x60;Inprogress&#x60; - Upgrade impact calculation is in progress. * &#x60;Completed&#x60; - Upgrade impact calculation is completed. * &#x60;Unavailable&#x60; - Upgrade impact is not available since image is not present in FI. | [optional] [default to "Inprogress"]
**ExcludeComponents** | Pointer to **[]string** |  | [optional] 
**Impacts** | Pointer to [**[]FirmwareBaseImpact**](firmware.BaseImpact.md) |  | [optional] 
**Summary** | Pointer to **string** | The summary on the component or components getting impacted by the upgrade. * &#x60;Basic&#x60; - Summary of a single instance involved in the upgrade operation. * &#x60;Detail&#x60; - Summary of the collection of single instances for a complex component involved in the upgrade operation. For example, in case of a server upgrade, a detailed summary indicates impact of all the single instances which are part of the server, such as storage controller, drives, and BIOS. | [optional] [default to "Basic"]

## Methods

### NewFirmwareUpgradeImpactBase

`func NewFirmwareUpgradeImpactBase() *FirmwareUpgradeImpactBase`

NewFirmwareUpgradeImpactBase instantiates a new FirmwareUpgradeImpactBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeImpactBaseWithDefaults

`func NewFirmwareUpgradeImpactBaseWithDefaults() *FirmwareUpgradeImpactBase`

NewFirmwareUpgradeImpactBaseWithDefaults instantiates a new FirmwareUpgradeImpactBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *FirmwareUpgradeImpactBase) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FirmwareUpgradeImpactBase) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FirmwareUpgradeImpactBase) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FirmwareUpgradeImpactBase) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetComputationState

`func (o *FirmwareUpgradeImpactBase) GetComputationState() string`

GetComputationState returns the ComputationState field if non-nil, zero value otherwise.

### GetComputationStateOk

`func (o *FirmwareUpgradeImpactBase) GetComputationStateOk() (*string, bool)`

GetComputationStateOk returns a tuple with the ComputationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationState

`func (o *FirmwareUpgradeImpactBase) SetComputationState(v string)`

SetComputationState sets ComputationState field to given value.

### HasComputationState

`func (o *FirmwareUpgradeImpactBase) HasComputationState() bool`

HasComputationState returns a boolean if a field has been set.

### GetExcludeComponents

`func (o *FirmwareUpgradeImpactBase) GetExcludeComponents() []string`

GetExcludeComponents returns the ExcludeComponents field if non-nil, zero value otherwise.

### GetExcludeComponentsOk

`func (o *FirmwareUpgradeImpactBase) GetExcludeComponentsOk() (*[]string, bool)`

GetExcludeComponentsOk returns a tuple with the ExcludeComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponents

`func (o *FirmwareUpgradeImpactBase) SetExcludeComponents(v []string)`

SetExcludeComponents sets ExcludeComponents field to given value.

### HasExcludeComponents

`func (o *FirmwareUpgradeImpactBase) HasExcludeComponents() bool`

HasExcludeComponents returns a boolean if a field has been set.

### GetImpacts

`func (o *FirmwareUpgradeImpactBase) GetImpacts() []FirmwareBaseImpact`

GetImpacts returns the Impacts field if non-nil, zero value otherwise.

### GetImpactsOk

`func (o *FirmwareUpgradeImpactBase) GetImpactsOk() (*[]FirmwareBaseImpact, bool)`

GetImpactsOk returns a tuple with the Impacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacts

`func (o *FirmwareUpgradeImpactBase) SetImpacts(v []FirmwareBaseImpact)`

SetImpacts sets Impacts field to given value.

### HasImpacts

`func (o *FirmwareUpgradeImpactBase) HasImpacts() bool`

HasImpacts returns a boolean if a field has been set.

### GetSummary

`func (o *FirmwareUpgradeImpactBase) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FirmwareUpgradeImpactBase) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FirmwareUpgradeImpactBase) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *FirmwareUpgradeImpactBase) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


