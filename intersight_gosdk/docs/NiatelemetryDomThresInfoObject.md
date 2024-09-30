# NiatelemetryDomThresInfoObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DomThresInfoObject"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DomThresInfoObject"]
**DomThresInfo** | Pointer to [**[]NiatelemetryDomThresInfo**](NiatelemetryDomThresInfo.md) |  | [optional] 
**RecordType** | Pointer to **string** | Type of record NEXUS - This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed - This determines what was the API version for data available from the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried - The serial number is unique per device. | [optional] 
**SlotId** | Pointer to **string** | Line card slot of device being inventoried - The linecard number is specific to serial of a device. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryDomThresInfoObject

`func NewNiatelemetryDomThresInfoObject(classId string, objectType string, ) *NiatelemetryDomThresInfoObject`

NewNiatelemetryDomThresInfoObject instantiates a new NiatelemetryDomThresInfoObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDomThresInfoObjectWithDefaults

`func NewNiatelemetryDomThresInfoObjectWithDefaults() *NiatelemetryDomThresInfoObject`

NewNiatelemetryDomThresInfoObjectWithDefaults instantiates a new NiatelemetryDomThresInfoObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDomThresInfoObject) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDomThresInfoObject) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDomThresInfoObject) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDomThresInfoObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDomThresInfoObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDomThresInfoObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDomThresInfo

`func (o *NiatelemetryDomThresInfoObject) GetDomThresInfo() []NiatelemetryDomThresInfo`

GetDomThresInfo returns the DomThresInfo field if non-nil, zero value otherwise.

### GetDomThresInfoOk

`func (o *NiatelemetryDomThresInfoObject) GetDomThresInfoOk() (*[]NiatelemetryDomThresInfo, bool)`

GetDomThresInfoOk returns a tuple with the DomThresInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomThresInfo

`func (o *NiatelemetryDomThresInfoObject) SetDomThresInfo(v []NiatelemetryDomThresInfo)`

SetDomThresInfo sets DomThresInfo field to given value.

### HasDomThresInfo

`func (o *NiatelemetryDomThresInfoObject) HasDomThresInfo() bool`

HasDomThresInfo returns a boolean if a field has been set.

### SetDomThresInfoNil

`func (o *NiatelemetryDomThresInfoObject) SetDomThresInfoNil(b bool)`

 SetDomThresInfoNil sets the value for DomThresInfo to be an explicit nil

### UnsetDomThresInfo
`func (o *NiatelemetryDomThresInfoObject) UnsetDomThresInfo()`

UnsetDomThresInfo ensures that no value is present for DomThresInfo, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryDomThresInfoObject) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryDomThresInfoObject) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryDomThresInfoObject) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryDomThresInfoObject) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryDomThresInfoObject) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryDomThresInfoObject) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryDomThresInfoObject) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryDomThresInfoObject) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryDomThresInfoObject) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryDomThresInfoObject) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryDomThresInfoObject) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryDomThresInfoObject) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSlotId

`func (o *NiatelemetryDomThresInfoObject) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *NiatelemetryDomThresInfoObject) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *NiatelemetryDomThresInfoObject) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *NiatelemetryDomThresInfoObject) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryDomThresInfoObject) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryDomThresInfoObject) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryDomThresInfoObject) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryDomThresInfoObject) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryDomThresInfoObject) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryDomThresInfoObject) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


