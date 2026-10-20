# NiatelemetryDomInfoObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DomInfoObject"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DomInfoObject"]
**CollectionId** | Pointer to **string** | Collection id is for index of one of 4 records in the timestamp interval for the particular dom threshold info. | [optional] 
**DomInfo** | Pointer to [**[]NiatelemetryDomInfo**](NiatelemetryDomInfo.md) |  | [optional] 
**RecordType** | Pointer to **string** | Type of record NEXUS. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**SlotId** | Pointer to **string** | Line card slot of device being inventoried - The linecard number is specific to serial of a device. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryDomInfoObject

`func NewNiatelemetryDomInfoObject(classId string, objectType string, ) *NiatelemetryDomInfoObject`

NewNiatelemetryDomInfoObject instantiates a new NiatelemetryDomInfoObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDomInfoObjectWithDefaults

`func NewNiatelemetryDomInfoObjectWithDefaults() *NiatelemetryDomInfoObject`

NewNiatelemetryDomInfoObjectWithDefaults instantiates a new NiatelemetryDomInfoObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDomInfoObject) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDomInfoObject) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDomInfoObject) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDomInfoObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDomInfoObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDomInfoObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectionId

`func (o *NiatelemetryDomInfoObject) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *NiatelemetryDomInfoObject) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *NiatelemetryDomInfoObject) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *NiatelemetryDomInfoObject) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetDomInfo

`func (o *NiatelemetryDomInfoObject) GetDomInfo() []NiatelemetryDomInfo`

GetDomInfo returns the DomInfo field if non-nil, zero value otherwise.

### GetDomInfoOk

`func (o *NiatelemetryDomInfoObject) GetDomInfoOk() (*[]NiatelemetryDomInfo, bool)`

GetDomInfoOk returns a tuple with the DomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomInfo

`func (o *NiatelemetryDomInfoObject) SetDomInfo(v []NiatelemetryDomInfo)`

SetDomInfo sets DomInfo field to given value.

### HasDomInfo

`func (o *NiatelemetryDomInfoObject) HasDomInfo() bool`

HasDomInfo returns a boolean if a field has been set.

### SetDomInfoNil

`func (o *NiatelemetryDomInfoObject) SetDomInfoNil(b bool)`

 SetDomInfoNil sets the value for DomInfo to be an explicit nil

### UnsetDomInfo
`func (o *NiatelemetryDomInfoObject) UnsetDomInfo()`

UnsetDomInfo ensures that no value is present for DomInfo, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryDomInfoObject) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryDomInfoObject) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryDomInfoObject) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryDomInfoObject) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryDomInfoObject) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryDomInfoObject) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryDomInfoObject) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryDomInfoObject) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryDomInfoObject) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryDomInfoObject) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryDomInfoObject) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryDomInfoObject) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSlotId

`func (o *NiatelemetryDomInfoObject) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *NiatelemetryDomInfoObject) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *NiatelemetryDomInfoObject) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *NiatelemetryDomInfoObject) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryDomInfoObject) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryDomInfoObject) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryDomInfoObject) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryDomInfoObject) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryDomInfoObject) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryDomInfoObject) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


