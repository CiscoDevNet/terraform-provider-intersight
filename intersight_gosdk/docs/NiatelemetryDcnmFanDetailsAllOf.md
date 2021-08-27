# NiatelemetryDcnmFanDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DcnmFanDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DcnmFanDetails"]
**Name** | Pointer to **string** | Name of the fan used in the switch. | [optional] 
**ProductId** | Pointer to **string** | Product ID of the fan used in the switch. | [optional] 
**RecordType** | Pointer to **string** | Type of record. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the fan used in the switch. | [optional] 
**VendorId** | Pointer to **string** | Vendor Id of the fan used in the switch. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryDcnmFanDetailsAllOf

`func NewNiatelemetryDcnmFanDetailsAllOf(classId string, objectType string, ) *NiatelemetryDcnmFanDetailsAllOf`

NewNiatelemetryDcnmFanDetailsAllOf instantiates a new NiatelemetryDcnmFanDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDcnmFanDetailsAllOfWithDefaults

`func NewNiatelemetryDcnmFanDetailsAllOfWithDefaults() *NiatelemetryDcnmFanDetailsAllOf`

NewNiatelemetryDcnmFanDetailsAllOfWithDefaults instantiates a new NiatelemetryDcnmFanDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryDcnmFanDetailsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductId

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *NiatelemetryDcnmFanDetailsAllOf) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryDcnmFanDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryDcnmFanDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryDcnmFanDetailsAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetVendorId

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *NiatelemetryDcnmFanDetailsAllOf) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryDcnmFanDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryDcnmFanDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryDcnmFanDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


