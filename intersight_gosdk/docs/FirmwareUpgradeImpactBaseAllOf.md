# FirmwareUpgradeImpactBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Components** | Pointer to **[]string** |  | [optional] 
**ComputationState** | Pointer to **string** | Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint. * &#x60;Inprogress&#x60; - Upgrade impact calculation is in progress. * &#x60;Completed&#x60; - Upgrade impact calculation is completed. * &#x60;Unavailable&#x60; - Upgrade impact is not available since image is not present in FI. | [optional] [default to "Inprogress"]
**ExcludeComponents** | Pointer to **[]string** |  | [optional] 
**Impacts** | Pointer to [**[]FirmwareBaseImpact**](FirmwareBaseImpact.md) |  | [optional] 
**Summary** | Pointer to **string** | The summary on the component or components getting impacted by the upgrade. * &#x60;Basic&#x60; - Summary of a single instance involved in the upgrade operation. * &#x60;Detail&#x60; - Summary of the collection of single instances for a complex component involved in the upgrade operation. For example, in case of a server upgrade, a detailed summary indicates impact of all the single instances which are part of the server, such as storage controller, drives, and BIOS. | [optional] [default to "Basic"]

## Methods

### NewFirmwareUpgradeImpactBaseAllOf

`func NewFirmwareUpgradeImpactBaseAllOf(classId string, objectType string, ) *FirmwareUpgradeImpactBaseAllOf`

NewFirmwareUpgradeImpactBaseAllOf instantiates a new FirmwareUpgradeImpactBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeImpactBaseAllOfWithDefaults

`func NewFirmwareUpgradeImpactBaseAllOfWithDefaults() *FirmwareUpgradeImpactBaseAllOf`

NewFirmwareUpgradeImpactBaseAllOfWithDefaults instantiates a new FirmwareUpgradeImpactBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgradeImpactBaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeImpactBaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeImpactBaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgradeImpactBaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeImpactBaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeImpactBaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetComponentsNil

`func (o *FirmwareUpgradeImpactBaseAllOf) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *FirmwareUpgradeImpactBaseAllOf) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
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

### SetExcludeComponentsNil

`func (o *FirmwareUpgradeImpactBaseAllOf) SetExcludeComponentsNil(b bool)`

 SetExcludeComponentsNil sets the value for ExcludeComponents to be an explicit nil

### UnsetExcludeComponents
`func (o *FirmwareUpgradeImpactBaseAllOf) UnsetExcludeComponents()`

UnsetExcludeComponents ensures that no value is present for ExcludeComponents, not even an explicit nil
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

### SetImpactsNil

`func (o *FirmwareUpgradeImpactBaseAllOf) SetImpactsNil(b bool)`

 SetImpactsNil sets the value for Impacts to be an explicit nil

### UnsetImpacts
`func (o *FirmwareUpgradeImpactBaseAllOf) UnsetImpacts()`

UnsetImpacts ensures that no value is present for Impacts, not even an explicit nil
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


