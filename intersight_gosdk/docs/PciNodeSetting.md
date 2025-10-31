# PciNodeSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pci.NodeSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pci.NodeSetting"]
**AdminAction** | Pointer to **string** | User configured action on the PCIe Node. * &#x60;None&#x60; - Placeholder default value for PCIe Node admin state property. * &#x60;Reboot&#x60; - PCIe node reboot state property value. | [optional] [default to "None"]
**AdminLocatorLedState** | Pointer to **string** | User configured state of the locator LED for the PCIe Node. * &#x60;None&#x60; - No operation property for locator led. * &#x60;On&#x60; - The Locator Led is turned on. * &#x60;Off&#x60; - The Locator Led is turned off. | [optional] [default to "None"]
**ConfigState** | Pointer to **string** | The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Scheduled&#x60; - User configured settings are scheduled to be applied. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "Applied"]
**Name** | Pointer to **string** | The property used to identify the PCIe node it is associated with. | [optional] [readonly] 
**NodeOpStatus** | Pointer to [**[]ComputeServerOpStatus**](ComputeServerOpStatus.md) |  | [optional] 
**LocatorLed** | Pointer to [**NullableEquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**PciNode** | Pointer to [**NullablePciNodeRelationship**](PciNodeRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPciNodeSetting

`func NewPciNodeSetting(classId string, objectType string, ) *PciNodeSetting`

NewPciNodeSetting instantiates a new PciNodeSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciNodeSettingWithDefaults

`func NewPciNodeSettingWithDefaults() *PciNodeSetting`

NewPciNodeSettingWithDefaults instantiates a new PciNodeSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PciNodeSetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PciNodeSetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PciNodeSetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PciNodeSetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PciNodeSetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PciNodeSetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *PciNodeSetting) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *PciNodeSetting) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *PciNodeSetting) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *PciNodeSetting) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminLocatorLedState

`func (o *PciNodeSetting) GetAdminLocatorLedState() string`

GetAdminLocatorLedState returns the AdminLocatorLedState field if non-nil, zero value otherwise.

### GetAdminLocatorLedStateOk

`func (o *PciNodeSetting) GetAdminLocatorLedStateOk() (*string, bool)`

GetAdminLocatorLedStateOk returns a tuple with the AdminLocatorLedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedState

`func (o *PciNodeSetting) SetAdminLocatorLedState(v string)`

SetAdminLocatorLedState sets AdminLocatorLedState field to given value.

### HasAdminLocatorLedState

`func (o *PciNodeSetting) HasAdminLocatorLedState() bool`

HasAdminLocatorLedState returns a boolean if a field has been set.

### GetConfigState

`func (o *PciNodeSetting) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PciNodeSetting) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PciNodeSetting) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PciNodeSetting) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetName

`func (o *PciNodeSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PciNodeSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PciNodeSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PciNodeSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeOpStatus

`func (o *PciNodeSetting) GetNodeOpStatus() []ComputeServerOpStatus`

GetNodeOpStatus returns the NodeOpStatus field if non-nil, zero value otherwise.

### GetNodeOpStatusOk

`func (o *PciNodeSetting) GetNodeOpStatusOk() (*[]ComputeServerOpStatus, bool)`

GetNodeOpStatusOk returns a tuple with the NodeOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOpStatus

`func (o *PciNodeSetting) SetNodeOpStatus(v []ComputeServerOpStatus)`

SetNodeOpStatus sets NodeOpStatus field to given value.

### HasNodeOpStatus

`func (o *PciNodeSetting) HasNodeOpStatus() bool`

HasNodeOpStatus returns a boolean if a field has been set.

### SetNodeOpStatusNil

`func (o *PciNodeSetting) SetNodeOpStatusNil(b bool)`

 SetNodeOpStatusNil sets the value for NodeOpStatus to be an explicit nil

### UnsetNodeOpStatus
`func (o *PciNodeSetting) UnsetNodeOpStatus()`

UnsetNodeOpStatus ensures that no value is present for NodeOpStatus, not even an explicit nil
### GetLocatorLed

`func (o *PciNodeSetting) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *PciNodeSetting) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *PciNodeSetting) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *PciNodeSetting) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### SetLocatorLedNil

`func (o *PciNodeSetting) SetLocatorLedNil(b bool)`

 SetLocatorLedNil sets the value for LocatorLed to be an explicit nil

### UnsetLocatorLed
`func (o *PciNodeSetting) UnsetLocatorLed()`

UnsetLocatorLed ensures that no value is present for LocatorLed, not even an explicit nil
### GetPciNode

`func (o *PciNodeSetting) GetPciNode() PciNodeRelationship`

GetPciNode returns the PciNode field if non-nil, zero value otherwise.

### GetPciNodeOk

`func (o *PciNodeSetting) GetPciNodeOk() (*PciNodeRelationship, bool)`

GetPciNodeOk returns a tuple with the PciNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciNode

`func (o *PciNodeSetting) SetPciNode(v PciNodeRelationship)`

SetPciNode sets PciNode field to given value.

### HasPciNode

`func (o *PciNodeSetting) HasPciNode() bool`

HasPciNode returns a boolean if a field has been set.

### SetPciNodeNil

`func (o *PciNodeSetting) SetPciNodeNil(b bool)`

 SetPciNodeNil sets the value for PciNode to be an explicit nil

### UnsetPciNode
`func (o *PciNodeSetting) UnsetPciNode()`

UnsetPciNode ensures that no value is present for PciNode, not even an explicit nil
### GetRegisteredDevice

`func (o *PciNodeSetting) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciNodeSetting) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciNodeSetting) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciNodeSetting) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *PciNodeSetting) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *PciNodeSetting) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


