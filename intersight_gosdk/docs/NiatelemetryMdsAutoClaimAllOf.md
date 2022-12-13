# NiatelemetryMdsAutoClaimAllOf

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

### NewNiatelemetryMdsAutoClaimAllOf

`func NewNiatelemetryMdsAutoClaimAllOf(classId string, objectType string, ) *NiatelemetryMdsAutoClaimAllOf`

NewNiatelemetryMdsAutoClaimAllOf instantiates a new NiatelemetryMdsAutoClaimAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMdsAutoClaimAllOfWithDefaults

`func NewNiatelemetryMdsAutoClaimAllOfWithDefaults() *NiatelemetryMdsAutoClaimAllOf`

NewNiatelemetryMdsAutoClaimAllOfWithDefaults instantiates a new NiatelemetryMdsAutoClaimAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMdsAutoClaimAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMdsAutoClaimAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMdsAutoClaimAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMdsAutoClaimAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceIp

`func (o *NiatelemetryMdsAutoClaimAllOf) GetDeviceIp() string`

GetDeviceIp returns the DeviceIp field if non-nil, zero value otherwise.

### GetDeviceIpOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetDeviceIpOk() (*string, bool)`

GetDeviceIpOk returns a tuple with the DeviceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIp

`func (o *NiatelemetryMdsAutoClaimAllOf) SetDeviceIp(v string)`

SetDeviceIp sets DeviceIp field to given value.

### HasDeviceIp

`func (o *NiatelemetryMdsAutoClaimAllOf) HasDeviceIp() bool`

HasDeviceIp returns a boolean if a field has been set.

### GetDeviceName

`func (o *NiatelemetryMdsAutoClaimAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryMdsAutoClaimAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryMdsAutoClaimAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceWwn

`func (o *NiatelemetryMdsAutoClaimAllOf) GetDeviceWwn() string`

GetDeviceWwn returns the DeviceWwn field if non-nil, zero value otherwise.

### GetDeviceWwnOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetDeviceWwnOk() (*string, bool)`

GetDeviceWwnOk returns a tuple with the DeviceWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWwn

`func (o *NiatelemetryMdsAutoClaimAllOf) SetDeviceWwn(v string)`

SetDeviceWwn sets DeviceWwn field to given value.

### HasDeviceWwn

`func (o *NiatelemetryMdsAutoClaimAllOf) HasDeviceWwn() bool`

HasDeviceWwn returns a boolean if a field has been set.

### GetNeighborInfo

`func (o *NiatelemetryMdsAutoClaimAllOf) GetNeighborInfo() []NiatelemetryMdsNeighborInfo`

GetNeighborInfo returns the NeighborInfo field if non-nil, zero value otherwise.

### GetNeighborInfoOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetNeighborInfoOk() (*[]NiatelemetryMdsNeighborInfo, bool)`

GetNeighborInfoOk returns a tuple with the NeighborInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborInfo

`func (o *NiatelemetryMdsAutoClaimAllOf) SetNeighborInfo(v []NiatelemetryMdsNeighborInfo)`

SetNeighborInfo sets NeighborInfo field to given value.

### HasNeighborInfo

`func (o *NiatelemetryMdsAutoClaimAllOf) HasNeighborInfo() bool`

HasNeighborInfo returns a boolean if a field has been set.

### SetNeighborInfoNil

`func (o *NiatelemetryMdsAutoClaimAllOf) SetNeighborInfoNil(b bool)`

 SetNeighborInfoNil sets the value for NeighborInfo to be an explicit nil

### UnsetNeighborInfo
`func (o *NiatelemetryMdsAutoClaimAllOf) UnsetNeighborInfo()`

UnsetNeighborInfo ensures that no value is present for NeighborInfo, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryMdsAutoClaimAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryMdsAutoClaimAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryMdsAutoClaimAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryMdsAutoClaimAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryMdsAutoClaimAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryMdsAutoClaimAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryMdsAutoClaimAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryMdsAutoClaimAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryMdsAutoClaimAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMdsAutoClaimAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMdsAutoClaimAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMdsAutoClaimAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMdsAutoClaimAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


