# BiosBootDeviceAllOf

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

### NewBiosBootDeviceAllOf

`func NewBiosBootDeviceAllOf(classId string, objectType string, ) *BiosBootDeviceAllOf`

NewBiosBootDeviceAllOf instantiates a new BiosBootDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosBootDeviceAllOfWithDefaults

`func NewBiosBootDeviceAllOfWithDefaults() *BiosBootDeviceAllOf`

NewBiosBootDeviceAllOfWithDefaults instantiates a new BiosBootDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosBootDeviceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosBootDeviceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosBootDeviceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosBootDeviceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosBootDeviceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosBootDeviceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceName

`func (o *BiosBootDeviceAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *BiosBootDeviceAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *BiosBootDeviceAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *BiosBootDeviceAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *BiosBootDeviceAllOf) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BiosBootDeviceAllOf) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BiosBootDeviceAllOf) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *BiosBootDeviceAllOf) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetBiosSystemBootOrder

`func (o *BiosBootDeviceAllOf) GetBiosSystemBootOrder() BiosSystemBootOrderRelationship`

GetBiosSystemBootOrder returns the BiosSystemBootOrder field if non-nil, zero value otherwise.

### GetBiosSystemBootOrderOk

`func (o *BiosBootDeviceAllOf) GetBiosSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool)`

GetBiosSystemBootOrderOk returns a tuple with the BiosSystemBootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosSystemBootOrder

`func (o *BiosBootDeviceAllOf) SetBiosSystemBootOrder(v BiosSystemBootOrderRelationship)`

SetBiosSystemBootOrder sets BiosSystemBootOrder field to given value.

### HasBiosSystemBootOrder

`func (o *BiosBootDeviceAllOf) HasBiosSystemBootOrder() bool`

HasBiosSystemBootOrder returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BiosBootDeviceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BiosBootDeviceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BiosBootDeviceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BiosBootDeviceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


