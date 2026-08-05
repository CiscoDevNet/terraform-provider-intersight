# FirmwareSecureRouterUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.SecureRouterUpgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.SecureRouterUpgrade"]
**Device** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**NetworkSecureRouter** | Pointer to [**NullableNetworkSecureRouterRelationship**](NetworkSecureRouterRelationship.md) |  | [optional] 

## Methods

### NewFirmwareSecureRouterUpgrade

`func NewFirmwareSecureRouterUpgrade(classId string, objectType string, ) *FirmwareSecureRouterUpgrade`

NewFirmwareSecureRouterUpgrade instantiates a new FirmwareSecureRouterUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareSecureRouterUpgradeWithDefaults

`func NewFirmwareSecureRouterUpgradeWithDefaults() *FirmwareSecureRouterUpgrade`

NewFirmwareSecureRouterUpgradeWithDefaults instantiates a new FirmwareSecureRouterUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareSecureRouterUpgrade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareSecureRouterUpgrade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareSecureRouterUpgrade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareSecureRouterUpgrade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareSecureRouterUpgrade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareSecureRouterUpgrade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDevice

`func (o *FirmwareSecureRouterUpgrade) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareSecureRouterUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareSecureRouterUpgrade) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareSecureRouterUpgrade) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *FirmwareSecureRouterUpgrade) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *FirmwareSecureRouterUpgrade) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetNetworkSecureRouter

`func (o *FirmwareSecureRouterUpgrade) GetNetworkSecureRouter() NetworkSecureRouterRelationship`

GetNetworkSecureRouter returns the NetworkSecureRouter field if non-nil, zero value otherwise.

### GetNetworkSecureRouterOk

`func (o *FirmwareSecureRouterUpgrade) GetNetworkSecureRouterOk() (*NetworkSecureRouterRelationship, bool)`

GetNetworkSecureRouterOk returns a tuple with the NetworkSecureRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecureRouter

`func (o *FirmwareSecureRouterUpgrade) SetNetworkSecureRouter(v NetworkSecureRouterRelationship)`

SetNetworkSecureRouter sets NetworkSecureRouter field to given value.

### HasNetworkSecureRouter

`func (o *FirmwareSecureRouterUpgrade) HasNetworkSecureRouter() bool`

HasNetworkSecureRouter returns a boolean if a field has been set.

### SetNetworkSecureRouterNil

`func (o *FirmwareSecureRouterUpgrade) SetNetworkSecureRouterNil(b bool)`

 SetNetworkSecureRouterNil sets the value for NetworkSecureRouter to be an explicit nil

### UnsetNetworkSecureRouter
`func (o *FirmwareSecureRouterUpgrade) UnsetNetworkSecureRouter()`

UnsetNetworkSecureRouter ensures that no value is present for NetworkSecureRouter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


