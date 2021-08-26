# FirmwareSwitchUpgradeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.SwitchUpgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.SwitchUpgrade"]
**EnableFabricEvacuation** | Pointer to **bool** | The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen. | [optional] [default to true]
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**NetworkElements** | Pointer to [**[]NetworkElementRelationship**](NetworkElementRelationship.md) | An array of relationships to networkElement resources. | [optional] 

## Methods

### NewFirmwareSwitchUpgradeAllOf

`func NewFirmwareSwitchUpgradeAllOf(classId string, objectType string, ) *FirmwareSwitchUpgradeAllOf`

NewFirmwareSwitchUpgradeAllOf instantiates a new FirmwareSwitchUpgradeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareSwitchUpgradeAllOfWithDefaults

`func NewFirmwareSwitchUpgradeAllOfWithDefaults() *FirmwareSwitchUpgradeAllOf`

NewFirmwareSwitchUpgradeAllOfWithDefaults instantiates a new FirmwareSwitchUpgradeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareSwitchUpgradeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareSwitchUpgradeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareSwitchUpgradeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareSwitchUpgradeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareSwitchUpgradeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareSwitchUpgradeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableFabricEvacuation

`func (o *FirmwareSwitchUpgradeAllOf) GetEnableFabricEvacuation() bool`

GetEnableFabricEvacuation returns the EnableFabricEvacuation field if non-nil, zero value otherwise.

### GetEnableFabricEvacuationOk

`func (o *FirmwareSwitchUpgradeAllOf) GetEnableFabricEvacuationOk() (*bool, bool)`

GetEnableFabricEvacuationOk returns a tuple with the EnableFabricEvacuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFabricEvacuation

`func (o *FirmwareSwitchUpgradeAllOf) SetEnableFabricEvacuation(v bool)`

SetEnableFabricEvacuation sets EnableFabricEvacuation field to given value.

### HasEnableFabricEvacuation

`func (o *FirmwareSwitchUpgradeAllOf) HasEnableFabricEvacuation() bool`

HasEnableFabricEvacuation returns a boolean if a field has been set.

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


