# FaultInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to **string** | The user acknowledgement state of the fault. | [optional] [readonly] 
**AffectedDn** | Pointer to **string** | The Distinguished Name of the Managed object which was affected. | [optional] [readonly] 
**AffectedMoId** | Pointer to **string** | Managed object Id which was affected. | [optional] [readonly] 
**AffectedMoType** | Pointer to **string** | Managed object type which was affected. | [optional] [readonly] 
**AncestorMoId** | Pointer to **string** | Object Id of the parent of the Managed object which was affected. | [optional] [readonly] 
**AncestorMoType** | Pointer to **string** | Object type of the parent of the Managed object which was affected. | [optional] [readonly] 
**Code** | Pointer to **string** | Numerical fault code of the fault found. | [optional] [readonly] 
**CreationTime** | Pointer to **string** | The time of creation of the fault instance. | [optional] [readonly] 
**Description** | Pointer to **string** | Detailed message of the fault. | [optional] [readonly] 
**LastTransitionTime** | Pointer to **string** | Last transition time of the fault. | [optional] [readonly] 
**NumOccurrences** | Pointer to **int64** | The number of times this fault has occured. | [optional] [readonly] 
**OriginalSeverity** | Pointer to **string** | Current Severity of the fault found. | [optional] [readonly] 
**PreviousSeverity** | Pointer to **string** | The Severity of the fault prior to user update. | [optional] [readonly] 
**Rule** | Pointer to **string** | The rule that is responsible for generation of the fault. | [optional] [readonly] 
**Severity** | Pointer to **string** | Severity of the fault found. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewFaultInstanceAllOf

`func NewFaultInstanceAllOf() *FaultInstanceAllOf`

NewFaultInstanceAllOf instantiates a new FaultInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaultInstanceAllOfWithDefaults

`func NewFaultInstanceAllOfWithDefaults() *FaultInstanceAllOf`

NewFaultInstanceAllOfWithDefaults instantiates a new FaultInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *FaultInstanceAllOf) GetAcknowledged() string`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *FaultInstanceAllOf) GetAcknowledgedOk() (*string, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *FaultInstanceAllOf) SetAcknowledged(v string)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *FaultInstanceAllOf) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetAffectedDn

`func (o *FaultInstanceAllOf) GetAffectedDn() string`

GetAffectedDn returns the AffectedDn field if non-nil, zero value otherwise.

### GetAffectedDnOk

`func (o *FaultInstanceAllOf) GetAffectedDnOk() (*string, bool)`

GetAffectedDnOk returns a tuple with the AffectedDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedDn

`func (o *FaultInstanceAllOf) SetAffectedDn(v string)`

SetAffectedDn sets AffectedDn field to given value.

### HasAffectedDn

`func (o *FaultInstanceAllOf) HasAffectedDn() bool`

HasAffectedDn returns a boolean if a field has been set.

### GetAffectedMoId

`func (o *FaultInstanceAllOf) GetAffectedMoId() string`

GetAffectedMoId returns the AffectedMoId field if non-nil, zero value otherwise.

### GetAffectedMoIdOk

`func (o *FaultInstanceAllOf) GetAffectedMoIdOk() (*string, bool)`

GetAffectedMoIdOk returns a tuple with the AffectedMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoId

`func (o *FaultInstanceAllOf) SetAffectedMoId(v string)`

SetAffectedMoId sets AffectedMoId field to given value.

### HasAffectedMoId

`func (o *FaultInstanceAllOf) HasAffectedMoId() bool`

HasAffectedMoId returns a boolean if a field has been set.

### GetAffectedMoType

`func (o *FaultInstanceAllOf) GetAffectedMoType() string`

GetAffectedMoType returns the AffectedMoType field if non-nil, zero value otherwise.

### GetAffectedMoTypeOk

`func (o *FaultInstanceAllOf) GetAffectedMoTypeOk() (*string, bool)`

GetAffectedMoTypeOk returns a tuple with the AffectedMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoType

`func (o *FaultInstanceAllOf) SetAffectedMoType(v string)`

SetAffectedMoType sets AffectedMoType field to given value.

### HasAffectedMoType

`func (o *FaultInstanceAllOf) HasAffectedMoType() bool`

HasAffectedMoType returns a boolean if a field has been set.

### GetAncestorMoId

`func (o *FaultInstanceAllOf) GetAncestorMoId() string`

GetAncestorMoId returns the AncestorMoId field if non-nil, zero value otherwise.

### GetAncestorMoIdOk

`func (o *FaultInstanceAllOf) GetAncestorMoIdOk() (*string, bool)`

GetAncestorMoIdOk returns a tuple with the AncestorMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoId

`func (o *FaultInstanceAllOf) SetAncestorMoId(v string)`

SetAncestorMoId sets AncestorMoId field to given value.

### HasAncestorMoId

`func (o *FaultInstanceAllOf) HasAncestorMoId() bool`

HasAncestorMoId returns a boolean if a field has been set.

### GetAncestorMoType

`func (o *FaultInstanceAllOf) GetAncestorMoType() string`

GetAncestorMoType returns the AncestorMoType field if non-nil, zero value otherwise.

### GetAncestorMoTypeOk

`func (o *FaultInstanceAllOf) GetAncestorMoTypeOk() (*string, bool)`

GetAncestorMoTypeOk returns a tuple with the AncestorMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoType

`func (o *FaultInstanceAllOf) SetAncestorMoType(v string)`

SetAncestorMoType sets AncestorMoType field to given value.

### HasAncestorMoType

`func (o *FaultInstanceAllOf) HasAncestorMoType() bool`

HasAncestorMoType returns a boolean if a field has been set.

### GetCode

`func (o *FaultInstanceAllOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FaultInstanceAllOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FaultInstanceAllOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FaultInstanceAllOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreationTime

`func (o *FaultInstanceAllOf) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *FaultInstanceAllOf) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *FaultInstanceAllOf) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *FaultInstanceAllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *FaultInstanceAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FaultInstanceAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FaultInstanceAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FaultInstanceAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *FaultInstanceAllOf) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *FaultInstanceAllOf) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *FaultInstanceAllOf) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *FaultInstanceAllOf) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetNumOccurrences

`func (o *FaultInstanceAllOf) GetNumOccurrences() int64`

GetNumOccurrences returns the NumOccurrences field if non-nil, zero value otherwise.

### GetNumOccurrencesOk

`func (o *FaultInstanceAllOf) GetNumOccurrencesOk() (*int64, bool)`

GetNumOccurrencesOk returns a tuple with the NumOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOccurrences

`func (o *FaultInstanceAllOf) SetNumOccurrences(v int64)`

SetNumOccurrences sets NumOccurrences field to given value.

### HasNumOccurrences

`func (o *FaultInstanceAllOf) HasNumOccurrences() bool`

HasNumOccurrences returns a boolean if a field has been set.

### GetOriginalSeverity

`func (o *FaultInstanceAllOf) GetOriginalSeverity() string`

GetOriginalSeverity returns the OriginalSeverity field if non-nil, zero value otherwise.

### GetOriginalSeverityOk

`func (o *FaultInstanceAllOf) GetOriginalSeverityOk() (*string, bool)`

GetOriginalSeverityOk returns a tuple with the OriginalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSeverity

`func (o *FaultInstanceAllOf) SetOriginalSeverity(v string)`

SetOriginalSeverity sets OriginalSeverity field to given value.

### HasOriginalSeverity

`func (o *FaultInstanceAllOf) HasOriginalSeverity() bool`

HasOriginalSeverity returns a boolean if a field has been set.

### GetPreviousSeverity

`func (o *FaultInstanceAllOf) GetPreviousSeverity() string`

GetPreviousSeverity returns the PreviousSeverity field if non-nil, zero value otherwise.

### GetPreviousSeverityOk

`func (o *FaultInstanceAllOf) GetPreviousSeverityOk() (*string, bool)`

GetPreviousSeverityOk returns a tuple with the PreviousSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSeverity

`func (o *FaultInstanceAllOf) SetPreviousSeverity(v string)`

SetPreviousSeverity sets PreviousSeverity field to given value.

### HasPreviousSeverity

`func (o *FaultInstanceAllOf) HasPreviousSeverity() bool`

HasPreviousSeverity returns a boolean if a field has been set.

### GetRule

`func (o *FaultInstanceAllOf) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FaultInstanceAllOf) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FaultInstanceAllOf) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FaultInstanceAllOf) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSeverity

`func (o *FaultInstanceAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FaultInstanceAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FaultInstanceAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FaultInstanceAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *FaultInstanceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *FaultInstanceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *FaultInstanceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *FaultInstanceAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *FaultInstanceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FaultInstanceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FaultInstanceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FaultInstanceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


