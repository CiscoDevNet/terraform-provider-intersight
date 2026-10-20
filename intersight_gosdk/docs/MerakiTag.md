# MerakiTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meraki.Tag"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meraki.Tag"]
**Name** | Pointer to **string** | The name of the tag on different meraki entities. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the tag on different meraki entities. * &#x60;unknown&#x60; - The tag type is unknown, it is used as default value. * &#x60;Network&#x60; - The tag belongs to Meraki network. * &#x60;Device&#x60; - The tag belongs to Meraki device. * &#x60;SwitchPort&#x60; - The tag belongs to Meraki switch port. | [optional] [readonly] [default to "unknown"]
**Device** | Pointer to [**NullableMerakiDeviceRelationship**](MerakiDeviceRelationship.md) |  | [optional] 
**Network** | Pointer to [**NullableMerakiNetworkRelationship**](MerakiNetworkRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewMerakiTag

`func NewMerakiTag(classId string, objectType string, ) *MerakiTag`

NewMerakiTag instantiates a new MerakiTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerakiTagWithDefaults

`func NewMerakiTagWithDefaults() *MerakiTag`

NewMerakiTagWithDefaults instantiates a new MerakiTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MerakiTag) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MerakiTag) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MerakiTag) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MerakiTag) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MerakiTag) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MerakiTag) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *MerakiTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerakiTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerakiTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerakiTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MerakiTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerakiTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerakiTag) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MerakiTag) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *MerakiTag) GetDevice() MerakiDeviceRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *MerakiTag) GetDeviceOk() (*MerakiDeviceRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *MerakiTag) SetDevice(v MerakiDeviceRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *MerakiTag) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *MerakiTag) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *MerakiTag) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetNetwork

`func (o *MerakiTag) GetNetwork() MerakiNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MerakiTag) GetNetworkOk() (*MerakiNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MerakiTag) SetNetwork(v MerakiNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MerakiTag) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *MerakiTag) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *MerakiTag) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetRegisteredDevice

`func (o *MerakiTag) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MerakiTag) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MerakiTag) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MerakiTag) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *MerakiTag) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *MerakiTag) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


