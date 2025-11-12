# FirmwareChassisUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.ChassisUpgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.ChassisUpgrade"]
**ExcludeComponentList** | Pointer to **[]string** |  | [optional] 
**XfmUpgradeOption** | Pointer to **string** | XFM upgrade option Full or Partial Disruption. * &#x60;none&#x60; - If no option is selected for exclusion. * &#x60;full-shutdown&#x60; - PSX Switch in XFM will be upgraded in single action. * &#x60;partial-shutdown&#x60; - PSX Switch in XFM will be upgraded one after other. | [optional] [default to "none"]
**Chassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**Device** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewFirmwareChassisUpgrade

`func NewFirmwareChassisUpgrade(classId string, objectType string, ) *FirmwareChassisUpgrade`

NewFirmwareChassisUpgrade instantiates a new FirmwareChassisUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareChassisUpgradeWithDefaults

`func NewFirmwareChassisUpgradeWithDefaults() *FirmwareChassisUpgrade`

NewFirmwareChassisUpgradeWithDefaults instantiates a new FirmwareChassisUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareChassisUpgrade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareChassisUpgrade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareChassisUpgrade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareChassisUpgrade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareChassisUpgrade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareChassisUpgrade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeComponentList

`func (o *FirmwareChassisUpgrade) GetExcludeComponentList() []string`

GetExcludeComponentList returns the ExcludeComponentList field if non-nil, zero value otherwise.

### GetExcludeComponentListOk

`func (o *FirmwareChassisUpgrade) GetExcludeComponentListOk() (*[]string, bool)`

GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentList

`func (o *FirmwareChassisUpgrade) SetExcludeComponentList(v []string)`

SetExcludeComponentList sets ExcludeComponentList field to given value.

### HasExcludeComponentList

`func (o *FirmwareChassisUpgrade) HasExcludeComponentList() bool`

HasExcludeComponentList returns a boolean if a field has been set.

### SetExcludeComponentListNil

`func (o *FirmwareChassisUpgrade) SetExcludeComponentListNil(b bool)`

 SetExcludeComponentListNil sets the value for ExcludeComponentList to be an explicit nil

### UnsetExcludeComponentList
`func (o *FirmwareChassisUpgrade) UnsetExcludeComponentList()`

UnsetExcludeComponentList ensures that no value is present for ExcludeComponentList, not even an explicit nil
### GetXfmUpgradeOption

`func (o *FirmwareChassisUpgrade) GetXfmUpgradeOption() string`

GetXfmUpgradeOption returns the XfmUpgradeOption field if non-nil, zero value otherwise.

### GetXfmUpgradeOptionOk

`func (o *FirmwareChassisUpgrade) GetXfmUpgradeOptionOk() (*string, bool)`

GetXfmUpgradeOptionOk returns a tuple with the XfmUpgradeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXfmUpgradeOption

`func (o *FirmwareChassisUpgrade) SetXfmUpgradeOption(v string)`

SetXfmUpgradeOption sets XfmUpgradeOption field to given value.

### HasXfmUpgradeOption

`func (o *FirmwareChassisUpgrade) HasXfmUpgradeOption() bool`

HasXfmUpgradeOption returns a boolean if a field has been set.

### GetChassis

`func (o *FirmwareChassisUpgrade) GetChassis() EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *FirmwareChassisUpgrade) GetChassisOk() (*EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *FirmwareChassisUpgrade) SetChassis(v EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *FirmwareChassisUpgrade) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### SetChassisNil

`func (o *FirmwareChassisUpgrade) SetChassisNil(b bool)`

 SetChassisNil sets the value for Chassis to be an explicit nil

### UnsetChassis
`func (o *FirmwareChassisUpgrade) UnsetChassis()`

UnsetChassis ensures that no value is present for Chassis, not even an explicit nil
### GetDevice

`func (o *FirmwareChassisUpgrade) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareChassisUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareChassisUpgrade) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareChassisUpgrade) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *FirmwareChassisUpgrade) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *FirmwareChassisUpgrade) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


