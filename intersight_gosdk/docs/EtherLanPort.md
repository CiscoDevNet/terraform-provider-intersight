# EtherLanPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.LanPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.LanPort"]
**PortName** | Pointer to **string** | The name of the LAN port. | [optional] [readonly] 
**NetworkSecureRouter** | Pointer to [**NullableNetworkSecureRouterRelationship**](NetworkSecureRouterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEtherLanPort

`func NewEtherLanPort(classId string, objectType string, ) *EtherLanPort`

NewEtherLanPort instantiates a new EtherLanPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherLanPortWithDefaults

`func NewEtherLanPortWithDefaults() *EtherLanPort`

NewEtherLanPortWithDefaults instantiates a new EtherLanPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherLanPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherLanPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherLanPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherLanPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherLanPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherLanPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPortName

`func (o *EtherLanPort) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *EtherLanPort) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *EtherLanPort) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *EtherLanPort) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetNetworkSecureRouter

`func (o *EtherLanPort) GetNetworkSecureRouter() NetworkSecureRouterRelationship`

GetNetworkSecureRouter returns the NetworkSecureRouter field if non-nil, zero value otherwise.

### GetNetworkSecureRouterOk

`func (o *EtherLanPort) GetNetworkSecureRouterOk() (*NetworkSecureRouterRelationship, bool)`

GetNetworkSecureRouterOk returns a tuple with the NetworkSecureRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecureRouter

`func (o *EtherLanPort) SetNetworkSecureRouter(v NetworkSecureRouterRelationship)`

SetNetworkSecureRouter sets NetworkSecureRouter field to given value.

### HasNetworkSecureRouter

`func (o *EtherLanPort) HasNetworkSecureRouter() bool`

HasNetworkSecureRouter returns a boolean if a field has been set.

### SetNetworkSecureRouterNil

`func (o *EtherLanPort) SetNetworkSecureRouterNil(b bool)`

 SetNetworkSecureRouterNil sets the value for NetworkSecureRouter to be an explicit nil

### UnsetNetworkSecureRouter
`func (o *EtherLanPort) UnsetNetworkSecureRouter()`

UnsetNetworkSecureRouter ensures that no value is present for NetworkSecureRouter, not even an explicit nil
### GetRegisteredDevice

`func (o *EtherLanPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherLanPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherLanPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherLanPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *EtherLanPort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EtherLanPort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


