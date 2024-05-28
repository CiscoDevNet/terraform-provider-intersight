# ConsoleConsoleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "console.ConsoleConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "console.ConsoleConfig"]
**DataBits** | Pointer to **int64** | The databits numbers in console. | [optional] 
**Parity** | Pointer to **string** | The parity bit of the console. * &#x60;none&#x60; - If Parity is none, parity checking is not performed and the parity bit is not transmitted. * &#x60;odd&#x60; - If Parity is odd, the number of mark bits (1s) in the data is counted, and the parity bit is asserted or unasserted to obtain an odd number of mark bits. * &#x60;even&#x60; - If Parity is even, the number of mark bits in the data is counted, and the parity bit is asserted or unasserted to obtain an even number of mark bits. * &#x60;mark&#x60; - If Parity is mark, the parity bit is asserted. * &#x60;space&#x60; - If Parity is space, the parity bit is unasserted. | [optional] [default to "none"]
**Speed** | Pointer to **int64** | The speed of the console. | [optional] 
**StopBits** | Pointer to **int64** | The stopbits of async line in console. | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewConsoleConsoleConfig

`func NewConsoleConsoleConfig(classId string, objectType string, ) *ConsoleConsoleConfig`

NewConsoleConsoleConfig instantiates a new ConsoleConsoleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleConsoleConfigWithDefaults

`func NewConsoleConsoleConfigWithDefaults() *ConsoleConsoleConfig`

NewConsoleConsoleConfigWithDefaults instantiates a new ConsoleConsoleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConsoleConsoleConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConsoleConsoleConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConsoleConsoleConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConsoleConsoleConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConsoleConsoleConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConsoleConsoleConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataBits

`func (o *ConsoleConsoleConfig) GetDataBits() int64`

GetDataBits returns the DataBits field if non-nil, zero value otherwise.

### GetDataBitsOk

`func (o *ConsoleConsoleConfig) GetDataBitsOk() (*int64, bool)`

GetDataBitsOk returns a tuple with the DataBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBits

`func (o *ConsoleConsoleConfig) SetDataBits(v int64)`

SetDataBits sets DataBits field to given value.

### HasDataBits

`func (o *ConsoleConsoleConfig) HasDataBits() bool`

HasDataBits returns a boolean if a field has been set.

### GetParity

`func (o *ConsoleConsoleConfig) GetParity() string`

GetParity returns the Parity field if non-nil, zero value otherwise.

### GetParityOk

`func (o *ConsoleConsoleConfig) GetParityOk() (*string, bool)`

GetParityOk returns a tuple with the Parity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParity

`func (o *ConsoleConsoleConfig) SetParity(v string)`

SetParity sets Parity field to given value.

### HasParity

`func (o *ConsoleConsoleConfig) HasParity() bool`

HasParity returns a boolean if a field has been set.

### GetSpeed

`func (o *ConsoleConsoleConfig) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ConsoleConsoleConfig) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ConsoleConsoleConfig) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ConsoleConsoleConfig) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStopBits

`func (o *ConsoleConsoleConfig) GetStopBits() int64`

GetStopBits returns the StopBits field if non-nil, zero value otherwise.

### GetStopBitsOk

`func (o *ConsoleConsoleConfig) GetStopBitsOk() (*int64, bool)`

GetStopBitsOk returns a tuple with the StopBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopBits

`func (o *ConsoleConsoleConfig) SetStopBits(v int64)`

SetStopBits sets StopBits field to given value.

### HasStopBits

`func (o *ConsoleConsoleConfig) HasStopBits() bool`

HasStopBits returns a boolean if a field has been set.

### GetNetworkElement

`func (o *ConsoleConsoleConfig) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *ConsoleConsoleConfig) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *ConsoleConsoleConfig) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *ConsoleConsoleConfig) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *ConsoleConsoleConfig) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *ConsoleConsoleConfig) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *ConsoleConsoleConfig) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ConsoleConsoleConfig) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ConsoleConsoleConfig) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ConsoleConsoleConfig) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ConsoleConsoleConfig) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ConsoleConsoleConfig) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


