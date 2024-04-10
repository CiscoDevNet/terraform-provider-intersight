# AdapterPhysicalNicModeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.PhysicalNicModeSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.PhysicalNicModeSettings"]
**PhyNicEnabled** | Pointer to **bool** | When Physical NIC Mode is enabled, up-link ports of the VIC are set to pass-through mode. This allows the host to transmit packets without any modification. When Physical NIC Mode is enabled, VLAN tagging of the packets will not happen. | [optional] [default to false]

## Methods

### NewAdapterPhysicalNicModeSettings

`func NewAdapterPhysicalNicModeSettings(classId string, objectType string, ) *AdapterPhysicalNicModeSettings`

NewAdapterPhysicalNicModeSettings instantiates a new AdapterPhysicalNicModeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterPhysicalNicModeSettingsWithDefaults

`func NewAdapterPhysicalNicModeSettingsWithDefaults() *AdapterPhysicalNicModeSettings`

NewAdapterPhysicalNicModeSettingsWithDefaults instantiates a new AdapterPhysicalNicModeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterPhysicalNicModeSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterPhysicalNicModeSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterPhysicalNicModeSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterPhysicalNicModeSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterPhysicalNicModeSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterPhysicalNicModeSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPhyNicEnabled

`func (o *AdapterPhysicalNicModeSettings) GetPhyNicEnabled() bool`

GetPhyNicEnabled returns the PhyNicEnabled field if non-nil, zero value otherwise.

### GetPhyNicEnabledOk

`func (o *AdapterPhysicalNicModeSettings) GetPhyNicEnabledOk() (*bool, bool)`

GetPhyNicEnabledOk returns a tuple with the PhyNicEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhyNicEnabled

`func (o *AdapterPhysicalNicModeSettings) SetPhyNicEnabled(v bool)`

SetPhyNicEnabled sets PhyNicEnabled field to given value.

### HasPhyNicEnabled

`func (o *AdapterPhysicalNicModeSettings) HasPhyNicEnabled() bool`

HasPhyNicEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


