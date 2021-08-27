# BiosBootDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.BootDevice"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.BootDevice"]
**DeviceName** | Pointer to **string** | Name of the Configured Boot Device. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | Type of the Configured Boot Device. | [optional] [readonly] 
**BiosSystemBootOrder** | Pointer to [**BiosSystemBootOrderRelationship**](BiosSystemBootOrderRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewBiosBootDevice

`func NewBiosBootDevice(classId string, objectType string, ) *BiosBootDevice`

NewBiosBootDevice instantiates a new BiosBootDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosBootDeviceWithDefaults

`func NewBiosBootDeviceWithDefaults() *BiosBootDevice`

NewBiosBootDeviceWithDefaults instantiates a new BiosBootDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosBootDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosBootDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosBootDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosBootDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosBootDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosBootDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


