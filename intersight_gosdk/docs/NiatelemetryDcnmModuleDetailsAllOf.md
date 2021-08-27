# NiatelemetryDcnmModuleDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DcnmModuleDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DcnmModuleDetails"]
**Name** | Pointer to **string** | Name of the hardware module in the fabric inventory. | [optional] 
**ProductId** | Pointer to **string** | Product ID of the hardware module in the fabric inventory. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the hardware module in the fabric inventory. | [optional] 
**VendorId** | Pointer to **string** | Vendor Id of the hardware module in the fabric inventory. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryDcnmModuleDetailsAllOf

`func NewNiatelemetryDcnmModuleDetailsAllOf(classId string, objectType string, ) *NiatelemetryDcnmModuleDetailsAllOf`

NewNiatelemetryDcnmModuleDetailsAllOf instantiates a new NiatelemetryDcnmModuleDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDcnmModuleDetailsAllOfWithDefaults

`func NewNiatelemetryDcnmModuleDetailsAllOfWithDefaults() *NiatelemetryDcnmModuleDetailsAllOf`

NewNiatelemetryDcnmModuleDetailsAllOfWithDefaults instantiates a new NiatelemetryDcnmModuleDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryDcnmModuleDetailsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryDcnmModuleDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryDcnmModuleDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryDcnmModuleDetailsAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetVendorId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *NiatelemetryDcnmModuleDetailsAllOf) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryDcnmModuleDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryDcnmModuleDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryDcnmModuleDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


