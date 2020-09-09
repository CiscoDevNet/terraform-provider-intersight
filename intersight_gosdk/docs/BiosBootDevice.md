# BiosBootDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | Name of the Configured Boot Device. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | Type of the Configured Boot Device. | [optional] [readonly] 
**BiosSystemBootOrder** | Pointer to [**BiosSystemBootOrderRelationship**](bios.SystemBootOrder.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewBiosBootDevice

`func NewBiosBootDevice() *BiosBootDevice`

NewBiosBootDevice instantiates a new BiosBootDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosBootDeviceWithDefaults

`func NewBiosBootDeviceWithDefaults() *BiosBootDevice`

NewBiosBootDeviceWithDefaults instantiates a new BiosBootDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *BiosBootDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *BiosBootDevice) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *BiosBootDevice) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *BiosBootDevice) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *BiosBootDevice) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BiosBootDevice) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BiosBootDevice) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *BiosBootDevice) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetBiosSystemBootOrder

`func (o *BiosBootDevice) GetBiosSystemBootOrder() BiosSystemBootOrderRelationship`

GetBiosSystemBootOrder returns the BiosSystemBootOrder field if non-nil, zero value otherwise.

### GetBiosSystemBootOrderOk

`func (o *BiosBootDevice) GetBiosSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool)`

GetBiosSystemBootOrderOk returns a tuple with the BiosSystemBootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosSystemBootOrder

`func (o *BiosBootDevice) SetBiosSystemBootOrder(v BiosSystemBootOrderRelationship)`

SetBiosSystemBootOrder sets BiosSystemBootOrder field to given value.

### HasBiosSystemBootOrder

`func (o *BiosBootDevice) HasBiosSystemBootOrder() bool`

HasBiosSystemBootOrder returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BiosBootDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosBootDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosBootDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosBootDevice) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


