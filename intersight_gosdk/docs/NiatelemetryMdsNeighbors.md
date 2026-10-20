# NiatelemetryMdsNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.MdsNeighbors"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.MdsNeighbors"]
**DeviceIp** | Pointer to **string** | Ip address of device being inventoried. | [optional] 
**DeviceName** | Pointer to **string** | Device name of device being inventoried. | [optional] 
**DeviceWwn** | Pointer to **string** | Device wwn of device being inventoried. | [optional] 
**NeighborInfo** | Pointer to [**[]NiatelemetryMdsNeighborInfo**](NiatelemetryMdsNeighborInfo.md) |  | [optional] 
**RecordType** | Pointer to **string** | Type of record MDS. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMdsNeighbors

`func NewNiatelemetryMdsNeighbors(classId string, objectType string, ) *NiatelemetryMdsNeighbors`

NewNiatelemetryMdsNeighbors instantiates a new NiatelemetryMdsNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMdsNeighborsWithDefaults

`func NewNiatelemetryMdsNeighborsWithDefaults() *NiatelemetryMdsNeighbors`

NewNiatelemetryMdsNeighborsWithDefaults instantiates a new NiatelemetryMdsNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMdsNeighbors) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMdsNeighbors) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMdsNeighbors) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMdsNeighbors) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMdsNeighbors) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMdsNeighbors) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceIp

`func (o *NiatelemetryMdsNeighbors) GetDeviceIp() string`

GetDeviceIp returns the DeviceIp field if non-nil, zero value otherwise.

### GetDeviceIpOk

`func (o *NiatelemetryMdsNeighbors) GetDeviceIpOk() (*string, bool)`

GetDeviceIpOk returns a tuple with the DeviceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIp

`func (o *NiatelemetryMdsNeighbors) SetDeviceIp(v string)`

SetDeviceIp sets DeviceIp field to given value.

### HasDeviceIp

`func (o *NiatelemetryMdsNeighbors) HasDeviceIp() bool`

HasDeviceIp returns a boolean if a field has been set.

### GetDeviceName

`func (o *NiatelemetryMdsNeighbors) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryMdsNeighbors) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryMdsNeighbors) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryMdsNeighbors) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceWwn

`func (o *NiatelemetryMdsNeighbors) GetDeviceWwn() string`

GetDeviceWwn returns the DeviceWwn field if non-nil, zero value otherwise.

### GetDeviceWwnOk

`func (o *NiatelemetryMdsNeighbors) GetDeviceWwnOk() (*string, bool)`

GetDeviceWwnOk returns a tuple with the DeviceWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWwn

`func (o *NiatelemetryMdsNeighbors) SetDeviceWwn(v string)`

SetDeviceWwn sets DeviceWwn field to given value.

### HasDeviceWwn

`func (o *NiatelemetryMdsNeighbors) HasDeviceWwn() bool`

HasDeviceWwn returns a boolean if a field has been set.

### GetNeighborInfo

`func (o *NiatelemetryMdsNeighbors) GetNeighborInfo() []NiatelemetryMdsNeighborInfo`

GetNeighborInfo returns the NeighborInfo field if non-nil, zero value otherwise.

### GetNeighborInfoOk

`func (o *NiatelemetryMdsNeighbors) GetNeighborInfoOk() (*[]NiatelemetryMdsNeighborInfo, bool)`

GetNeighborInfoOk returns a tuple with the NeighborInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborInfo

`func (o *NiatelemetryMdsNeighbors) SetNeighborInfo(v []NiatelemetryMdsNeighborInfo)`

SetNeighborInfo sets NeighborInfo field to given value.

### HasNeighborInfo

`func (o *NiatelemetryMdsNeighbors) HasNeighborInfo() bool`

HasNeighborInfo returns a boolean if a field has been set.

### SetNeighborInfoNil

`func (o *NiatelemetryMdsNeighbors) SetNeighborInfoNil(b bool)`

 SetNeighborInfoNil sets the value for NeighborInfo to be an explicit nil

### UnsetNeighborInfo
`func (o *NiatelemetryMdsNeighbors) UnsetNeighborInfo()`

UnsetNeighborInfo ensures that no value is present for NeighborInfo, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryMdsNeighbors) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryMdsNeighbors) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryMdsNeighbors) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryMdsNeighbors) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryMdsNeighbors) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryMdsNeighbors) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryMdsNeighbors) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryMdsNeighbors) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryMdsNeighbors) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryMdsNeighbors) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryMdsNeighbors) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryMdsNeighbors) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMdsNeighbors) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMdsNeighbors) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMdsNeighbors) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMdsNeighbors) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryMdsNeighbors) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryMdsNeighbors) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


