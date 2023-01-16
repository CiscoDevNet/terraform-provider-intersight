# NiatelemetryMdsNeighborsAllOf

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
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMdsNeighborsAllOf

`func NewNiatelemetryMdsNeighborsAllOf(classId string, objectType string, ) *NiatelemetryMdsNeighborsAllOf`

NewNiatelemetryMdsNeighborsAllOf instantiates a new NiatelemetryMdsNeighborsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMdsNeighborsAllOfWithDefaults

`func NewNiatelemetryMdsNeighborsAllOfWithDefaults() *NiatelemetryMdsNeighborsAllOf`

NewNiatelemetryMdsNeighborsAllOfWithDefaults instantiates a new NiatelemetryMdsNeighborsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMdsNeighborsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMdsNeighborsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMdsNeighborsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMdsNeighborsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceIp

`func (o *NiatelemetryMdsNeighborsAllOf) GetDeviceIp() string`

GetDeviceIp returns the DeviceIp field if non-nil, zero value otherwise.

### GetDeviceIpOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetDeviceIpOk() (*string, bool)`

GetDeviceIpOk returns a tuple with the DeviceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIp

`func (o *NiatelemetryMdsNeighborsAllOf) SetDeviceIp(v string)`

SetDeviceIp sets DeviceIp field to given value.

### HasDeviceIp

`func (o *NiatelemetryMdsNeighborsAllOf) HasDeviceIp() bool`

HasDeviceIp returns a boolean if a field has been set.

### GetDeviceName

`func (o *NiatelemetryMdsNeighborsAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryMdsNeighborsAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryMdsNeighborsAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceWwn

`func (o *NiatelemetryMdsNeighborsAllOf) GetDeviceWwn() string`

GetDeviceWwn returns the DeviceWwn field if non-nil, zero value otherwise.

### GetDeviceWwnOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetDeviceWwnOk() (*string, bool)`

GetDeviceWwnOk returns a tuple with the DeviceWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWwn

`func (o *NiatelemetryMdsNeighborsAllOf) SetDeviceWwn(v string)`

SetDeviceWwn sets DeviceWwn field to given value.

### HasDeviceWwn

`func (o *NiatelemetryMdsNeighborsAllOf) HasDeviceWwn() bool`

HasDeviceWwn returns a boolean if a field has been set.

### GetNeighborInfo

`func (o *NiatelemetryMdsNeighborsAllOf) GetNeighborInfo() []NiatelemetryMdsNeighborInfo`

GetNeighborInfo returns the NeighborInfo field if non-nil, zero value otherwise.

### GetNeighborInfoOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetNeighborInfoOk() (*[]NiatelemetryMdsNeighborInfo, bool)`

GetNeighborInfoOk returns a tuple with the NeighborInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborInfo

`func (o *NiatelemetryMdsNeighborsAllOf) SetNeighborInfo(v []NiatelemetryMdsNeighborInfo)`

SetNeighborInfo sets NeighborInfo field to given value.

### HasNeighborInfo

`func (o *NiatelemetryMdsNeighborsAllOf) HasNeighborInfo() bool`

HasNeighborInfo returns a boolean if a field has been set.

### SetNeighborInfoNil

`func (o *NiatelemetryMdsNeighborsAllOf) SetNeighborInfoNil(b bool)`

 SetNeighborInfoNil sets the value for NeighborInfo to be an explicit nil

### UnsetNeighborInfo
`func (o *NiatelemetryMdsNeighborsAllOf) UnsetNeighborInfo()`

UnsetNeighborInfo ensures that no value is present for NeighborInfo, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryMdsNeighborsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryMdsNeighborsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryMdsNeighborsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryMdsNeighborsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryMdsNeighborsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryMdsNeighborsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryMdsNeighborsAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryMdsNeighborsAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryMdsNeighborsAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMdsNeighborsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMdsNeighborsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMdsNeighborsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMdsNeighborsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


