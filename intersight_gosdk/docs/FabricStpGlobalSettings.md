# FabricStpGlobalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.StpGlobalSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.StpGlobalSettings"]
**StpMode** | Pointer to **string** | The Spanning Tree Protocol (STP) mode determines the specific version of STP that is used to prevent loops in a network topology. * &#x60;Disabled&#x60; - Spanning Tree Protocol is disabled, and the switch does not participate in STP calculations or operations. * &#x60;RPVST+&#x60; - Rapid Per-VLAN Spanning Tree (RPVST) is a Cisco proprietary protocol that improves STP by providing faster convergence and creating a separate spanning tree for each VLAN, enhancing network performance and redundancy. | [optional] [default to "Disabled"]

## Methods

### NewFabricStpGlobalSettings

`func NewFabricStpGlobalSettings(classId string, objectType string, ) *FabricStpGlobalSettings`

NewFabricStpGlobalSettings instantiates a new FabricStpGlobalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricStpGlobalSettingsWithDefaults

`func NewFabricStpGlobalSettingsWithDefaults() *FabricStpGlobalSettings`

NewFabricStpGlobalSettingsWithDefaults instantiates a new FabricStpGlobalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricStpGlobalSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricStpGlobalSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricStpGlobalSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricStpGlobalSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricStpGlobalSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricStpGlobalSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStpMode

`func (o *FabricStpGlobalSettings) GetStpMode() string`

GetStpMode returns the StpMode field if non-nil, zero value otherwise.

### GetStpModeOk

`func (o *FabricStpGlobalSettings) GetStpModeOk() (*string, bool)`

GetStpModeOk returns a tuple with the StpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpMode

`func (o *FabricStpGlobalSettings) SetStpMode(v string)`

SetStpMode sets StpMode field to given value.

### HasStpMode

`func (o *FabricStpGlobalSettings) HasStpMode() bool`

HasStpMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


