# NiatelemetryMdsAutoClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.MdsAutoClaim"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.MdsAutoClaim"]
**DeviceIp** | Pointer to **string** | Ip address of device being inventoried. | [optional] 
**DeviceName** | Pointer to **string** | Device name of device being inventoried. | [optional] 
**DeviceWwn** | Pointer to **string** | Device wwn of device being inventoried. | [optional] 
**NeighborInfo** | Pointer to [**[]NiatelemetryMdsNeighborInfo**](NiatelemetryMdsNeighborInfo.md) |  | [optional] 
**RecordType** | Pointer to **string** | Type of record MDS. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMdsAutoClaim

`func NewNiatelemetryMdsAutoClaim(classId string, objectType string, ) *NiatelemetryMdsAutoClaim`

NewNiatelemetryMdsAutoClaim instantiates a new NiatelemetryMdsAutoClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMdsAutoClaimWithDefaults

`func NewNiatelemetryMdsAutoClaimWithDefaults() *NiatelemetryMdsAutoClaim`

NewNiatelemetryMdsAutoClaimWithDefaults instantiates a new NiatelemetryMdsAutoClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMdsAutoClaim) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMdsAutoClaim) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMdsAutoClaim) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMdsAutoClaim) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMdsAutoClaim) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMdsAutoClaim) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceIp

`func (o *NiatelemetryMdsAutoClaim) GetDeviceIp() string`

GetDeviceIp returns the DeviceIp field if non-nil, zero value otherwise.

### GetDeviceIpOk

`func (o *NiatelemetryMdsAutoClaim) GetDeviceIpOk() (*string, bool)`

GetDeviceIpOk returns a tuple with the DeviceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIp

`func (o *NiatelemetryMdsAutoClaim) SetDeviceIp(v string)`

SetDeviceIp sets DeviceIp field to given value.

### HasDeviceIp

`func (o *NiatelemetryMdsAutoClaim) HasDeviceIp() bool`

HasDeviceIp returns a boolean if a field has been set.

### GetDeviceName

`func (o *NiatelemetryMdsAutoClaim) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryMdsAutoClaim) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryMdsAutoClaim) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryMdsAutoClaim) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceWwn

`func (o *NiatelemetryMdsAutoClaim) GetDeviceWwn() string`

GetDeviceWwn returns the DeviceWwn field if non-nil, zero value otherwise.

### GetDeviceWwnOk

`func (o *NiatelemetryMdsAutoClaim) GetDeviceWwnOk() (*string, bool)`

GetDeviceWwnOk returns a tuple with the DeviceWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWwn

`func (o *NiatelemetryMdsAutoClaim) SetDeviceWwn(v string)`

SetDeviceWwn sets DeviceWwn field to given value.

### HasDeviceWwn

`func (o *NiatelemetryMdsAutoClaim) HasDeviceWwn() bool`

HasDeviceWwn returns a boolean if a field has been set.

### GetNeighborInfo

`func (o *NiatelemetryMdsAutoClaim) GetNeighborInfo() []NiatelemetryMdsNeighborInfo`

GetNeighborInfo returns the NeighborInfo field if non-nil, zero value otherwise.

### GetNeighborInfoOk

`func (o *NiatelemetryMdsAutoClaim) GetNeighborInfoOk() (*[]NiatelemetryMdsNeighborInfo, bool)`

GetNeighborInfoOk returns a tuple with the NeighborInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborInfo

`func (o *NiatelemetryMdsAutoClaim) SetNeighborInfo(v []NiatelemetryMdsNeighborInfo)`

SetNeighborInfo sets NeighborInfo field to given value.

### HasNeighborInfo

`func (o *NiatelemetryMdsAutoClaim) HasNeighborInfo() bool`

HasNeighborInfo returns a boolean if a field has been set.

### SetNeighborInfoNil

`func (o *NiatelemetryMdsAutoClaim) SetNeighborInfoNil(b bool)`

 SetNeighborInfoNil sets the value for NeighborInfo to be an explicit nil

### UnsetNeighborInfo
`func (o *NiatelemetryMdsAutoClaim) UnsetNeighborInfo()`

UnsetNeighborInfo ensures that no value is present for NeighborInfo, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryMdsAutoClaim) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryMdsAutoClaim) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryMdsAutoClaim) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryMdsAutoClaim) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryMdsAutoClaim) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryMdsAutoClaim) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryMdsAutoClaim) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryMdsAutoClaim) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryMdsAutoClaim) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryMdsAutoClaim) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryMdsAutoClaim) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryMdsAutoClaim) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMdsAutoClaim) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMdsAutoClaim) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMdsAutoClaim) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMdsAutoClaim) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


