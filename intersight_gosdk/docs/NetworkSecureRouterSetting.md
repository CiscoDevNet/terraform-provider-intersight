# NetworkSecureRouterSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.SecureRouterSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.SecureRouterSetting"]
**AdminAction** | Pointer to **string** | User configured action on the Secure Router. * &#x60;None&#x60; - Placeholder default value for Secure Router admin state property. * &#x60;Reboot&#x60; - Secure Router reboot state property value. | [optional] [default to "None"]
**AdminLocatorLedState** | Pointer to **string** | User configured state of the locator LED for the PCIe Node. * &#x60;None&#x60; - No operation property for locator led. * &#x60;On&#x60; - The Locator Led is turned on. * &#x60;Off&#x60; - The Locator Led is turned off. | [optional] [default to "None"]
**ConfigState** | Pointer to **string** | The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Scheduled&#x60; - User configured settings are scheduled to be applied. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "Applied"]
**Name** | Pointer to **string** | The property used to identify the Secure Router it is associated with. | [optional] [readonly] 
**RouterOpStatus** | Pointer to [**[]ComputeServerOpStatus**](ComputeServerOpStatus.md) |  | [optional] 
**LocatorLed** | Pointer to [**NullableEquipmentLocatorLedRelationship**](EquipmentLocatorLedRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SecureRouter** | Pointer to [**NullableNetworkSecureRouterRelationship**](NetworkSecureRouterRelationship.md) |  | [optional] 

## Methods

### NewNetworkSecureRouterSetting

`func NewNetworkSecureRouterSetting(classId string, objectType string, ) *NetworkSecureRouterSetting`

NewNetworkSecureRouterSetting instantiates a new NetworkSecureRouterSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecureRouterSettingWithDefaults

`func NewNetworkSecureRouterSettingWithDefaults() *NetworkSecureRouterSetting`

NewNetworkSecureRouterSettingWithDefaults instantiates a new NetworkSecureRouterSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkSecureRouterSetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkSecureRouterSetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkSecureRouterSetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkSecureRouterSetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkSecureRouterSetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkSecureRouterSetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminAction

`func (o *NetworkSecureRouterSetting) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *NetworkSecureRouterSetting) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *NetworkSecureRouterSetting) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *NetworkSecureRouterSetting) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetAdminLocatorLedState

`func (o *NetworkSecureRouterSetting) GetAdminLocatorLedState() string`

GetAdminLocatorLedState returns the AdminLocatorLedState field if non-nil, zero value otherwise.

### GetAdminLocatorLedStateOk

`func (o *NetworkSecureRouterSetting) GetAdminLocatorLedStateOk() (*string, bool)`

GetAdminLocatorLedStateOk returns a tuple with the AdminLocatorLedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedState

`func (o *NetworkSecureRouterSetting) SetAdminLocatorLedState(v string)`

SetAdminLocatorLedState sets AdminLocatorLedState field to given value.

### HasAdminLocatorLedState

`func (o *NetworkSecureRouterSetting) HasAdminLocatorLedState() bool`

HasAdminLocatorLedState returns a boolean if a field has been set.

### GetConfigState

`func (o *NetworkSecureRouterSetting) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *NetworkSecureRouterSetting) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *NetworkSecureRouterSetting) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *NetworkSecureRouterSetting) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetName

`func (o *NetworkSecureRouterSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkSecureRouterSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkSecureRouterSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkSecureRouterSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRouterOpStatus

`func (o *NetworkSecureRouterSetting) GetRouterOpStatus() []ComputeServerOpStatus`

GetRouterOpStatus returns the RouterOpStatus field if non-nil, zero value otherwise.

### GetRouterOpStatusOk

`func (o *NetworkSecureRouterSetting) GetRouterOpStatusOk() (*[]ComputeServerOpStatus, bool)`

GetRouterOpStatusOk returns a tuple with the RouterOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterOpStatus

`func (o *NetworkSecureRouterSetting) SetRouterOpStatus(v []ComputeServerOpStatus)`

SetRouterOpStatus sets RouterOpStatus field to given value.

### HasRouterOpStatus

`func (o *NetworkSecureRouterSetting) HasRouterOpStatus() bool`

HasRouterOpStatus returns a boolean if a field has been set.

### SetRouterOpStatusNil

`func (o *NetworkSecureRouterSetting) SetRouterOpStatusNil(b bool)`

 SetRouterOpStatusNil sets the value for RouterOpStatus to be an explicit nil

### UnsetRouterOpStatus
`func (o *NetworkSecureRouterSetting) UnsetRouterOpStatus()`

UnsetRouterOpStatus ensures that no value is present for RouterOpStatus, not even an explicit nil
### GetLocatorLed

`func (o *NetworkSecureRouterSetting) GetLocatorLed() EquipmentLocatorLedRelationship`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *NetworkSecureRouterSetting) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *NetworkSecureRouterSetting) SetLocatorLed(v EquipmentLocatorLedRelationship)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *NetworkSecureRouterSetting) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### SetLocatorLedNil

`func (o *NetworkSecureRouterSetting) SetLocatorLedNil(b bool)`

 SetLocatorLedNil sets the value for LocatorLed to be an explicit nil

### UnsetLocatorLed
`func (o *NetworkSecureRouterSetting) UnsetLocatorLed()`

UnsetLocatorLed ensures that no value is present for LocatorLed, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkSecureRouterSetting) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkSecureRouterSetting) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkSecureRouterSetting) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkSecureRouterSetting) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkSecureRouterSetting) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkSecureRouterSetting) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSecureRouter

`func (o *NetworkSecureRouterSetting) GetSecureRouter() NetworkSecureRouterRelationship`

GetSecureRouter returns the SecureRouter field if non-nil, zero value otherwise.

### GetSecureRouterOk

`func (o *NetworkSecureRouterSetting) GetSecureRouterOk() (*NetworkSecureRouterRelationship, bool)`

GetSecureRouterOk returns a tuple with the SecureRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRouter

`func (o *NetworkSecureRouterSetting) SetSecureRouter(v NetworkSecureRouterRelationship)`

SetSecureRouter sets SecureRouter field to given value.

### HasSecureRouter

`func (o *NetworkSecureRouterSetting) HasSecureRouter() bool`

HasSecureRouter returns a boolean if a field has been set.

### SetSecureRouterNil

`func (o *NetworkSecureRouterSetting) SetSecureRouterNil(b bool)`

 SetSecureRouterNil sets the value for SecureRouter to be an explicit nil

### UnsetSecureRouter
`func (o *NetworkSecureRouterSetting) UnsetSecureRouter()`

UnsetSecureRouter ensures that no value is present for SecureRouter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


