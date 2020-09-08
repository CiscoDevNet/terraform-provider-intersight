# FirmwareUpgradeImpactBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to **[]string** |  | [optional] 
**ComputationState** | Pointer to **string** | Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint. * &#x60;Inprogress&#x60; - Upgrade impact calculation is in progress. * &#x60;Completed&#x60; - Upgrade impact calculation is completed. * &#x60;Unavailable&#x60; - Upgrade impact is not available since image is not present in FI. | [optional] [default to "Inprogress"]
**ExcludeComponents** | Pointer to **[]string** |  | [optional] 
**Impacts** | Pointer to [**[]FirmwareBaseImpact**](firmware.BaseImpact.md) |  | [optional] 
**Summary** | Pointer to **string** | The summary on the component or components getting impacted by the upgrade. * &#x60;Basic&#x60; - Summary of a single instance involved in the upgrade operation. * &#x60;Detail&#x60; - Summary of the collection of single instances for a complex component involved in the upgrade operation. For example, in case of a server upgrade, a detailed summary indicates impact of all the single instances which are part of the server, such as storage controller, drives, and BIOS. | [optional] [default to "Basic"]

## Methods

### NewFirmwareUpgradeImpactBaseAllOf

`func NewFirmwareUpgradeImpactBaseAllOf() *FirmwareUpgradeImpactBaseAllOf`

NewFirmwareUpgradeImpactBaseAllOf instantiates a new FirmwareUpgradeImpactBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeImpactBaseAllOfWithDefaults

`func NewFirmwareUpgradeImpactBaseAllOfWithDefaults() *FirmwareUpgradeImpactBaseAllOf`

NewFirmwareUpgradeImpactBaseAllOfWithDefaults instantiates a new FirmwareUpgradeImpactBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *FirmwareUpgradeImpactBaseAllOf) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FirmwareUpgradeImpactBaseAllOf) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FirmwareUpgradeImpactBaseAllOf) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FirmwareUpgradeImpactBaseAllOf) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetComputationState

`func (o *FirmwareUpgradeImpactBaseAllOf) GetComputationState() string`

GetComputationState returns the ComputationState field if non-nil, zero value otherwise.

### GetComputationStateOk

`func (o *FirmwareUpgradeImpactBaseAllOf) GetComputationStateOk() (*string, bool)`

GetComputationStateOk returns a tuple with the ComputationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationState

`func (o *FirmwareUpgradeImpactBaseAllOf) SetComputationState(v string)`

SetComputationState sets ComputationState field to given value.

### HasComputationState

`func (o *FirmwareUpgradeImpactBaseAllOf) HasComputationState() bool`

HasComputationState returns a boolean if a field has been set.

### GetExcludeComponents

`func (o *FirmwareUpgradeImpactBaseAllOf) GetExcludeComponents() []string`

GetExcludeComponents returns the ExcludeComponents field if non-nil, zero value otherwise.

### GetExcludeComponentsOk

`func (o *FirmwareUpgradeImpactBaseAllOf) GetExcludeComponentsOk() (*[]string, bool)`

GetExcludeComponentsOk returns a tuple with the ExcludeComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponents

`func (o *FirmwareUpgradeImpactBaseAllOf) SetExcludeComponents(v []string)`

SetExcludeComponents sets ExcludeComponents field to given value.

### HasExcludeComponents

`func (o *FirmwareUpgradeImpactBaseAllOf) HasExcludeComponents() bool`

HasExcludeComponents returns a boolean if a field has been set.

### GetImpacts

`func (o *FirmwareUpgradeImpactBaseAllOf) GetImpacts() []FirmwareBaseImpact`

GetImpacts returns the Impacts field if non-nil, zero value otherwise.

### GetImpactsOk

`func (o *FirmwareUpgradeImpactBaseAllOf) GetImpactsOk() (*[]FirmwareBaseImpact, bool)`

GetImpactsOk returns a tuple with the Impacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpacts

`func (o *FirmwareUpgradeImpactBaseAllOf) SetImpacts(v []FirmwareBaseImpact)`

SetImpacts sets Impacts field to given value.

### HasImpacts

`func (o *FirmwareUpgradeImpactBaseAllOf) HasImpacts() bool`

HasImpacts returns a boolean if a field has been set.

### GetSummary

`func (o *FirmwareUpgradeImpactBaseAllOf) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FirmwareUpgradeImpactBaseAllOf) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FirmwareUpgradeImpactBaseAllOf) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *FirmwareUpgradeImpactBaseAllOf) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


