# MerakiDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meraki.Device"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meraki.Device"]
**LanIp** | Pointer to **string** | LAN IP address of the device. | [optional] [readonly] 
**Mac** | Pointer to **string** | The MAC address of the device. | [optional] [readonly] 
**NetworkId** | Pointer to **string** | The network the device belongs to. | [optional] [readonly] 
**Network** | Pointer to [**NullableMerakiNetworkRelationship**](MerakiNetworkRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewMerakiDevice

`func NewMerakiDevice(classId string, objectType string, ) *MerakiDevice`

NewMerakiDevice instantiates a new MerakiDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerakiDeviceWithDefaults

`func NewMerakiDeviceWithDefaults() *MerakiDevice`

NewMerakiDeviceWithDefaults instantiates a new MerakiDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MerakiDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MerakiDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MerakiDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MerakiDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MerakiDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MerakiDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLanIp

`func (o *MerakiDevice) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *MerakiDevice) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *MerakiDevice) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *MerakiDevice) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetMac

`func (o *MerakiDevice) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MerakiDevice) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MerakiDevice) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MerakiDevice) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetworkId

`func (o *MerakiDevice) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *MerakiDevice) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *MerakiDevice) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *MerakiDevice) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetwork

`func (o *MerakiDevice) GetNetwork() MerakiNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MerakiDevice) GetNetworkOk() (*MerakiNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MerakiDevice) SetNetwork(v MerakiNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MerakiDevice) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *MerakiDevice) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *MerakiDevice) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetRegisteredDevice

`func (o *MerakiDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MerakiDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MerakiDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MerakiDevice) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *MerakiDevice) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *MerakiDevice) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


