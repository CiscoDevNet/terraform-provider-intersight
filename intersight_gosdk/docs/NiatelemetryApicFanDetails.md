# NiatelemetryApicFanDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicFanDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicFanDetails"]
**Dn** | Pointer to **string** | Dn for the fan in the inventory. | [optional] 
**ModelNumber** | Pointer to **string** | Model number of the fan in APIC. | [optional] 
**NodeId** | Pointer to **int64** | Node id for the fan in the inventory. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the fan in APIC. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**VendorName** | Pointer to **string** | Vendor name of the fan in APIC. | [optional] 
**Vid** | Pointer to **string** | VID for the fan in the inventory. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicFanDetails

`func NewNiatelemetryApicFanDetails(classId string, objectType string, ) *NiatelemetryApicFanDetails`

NewNiatelemetryApicFanDetails instantiates a new NiatelemetryApicFanDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicFanDetailsWithDefaults

`func NewNiatelemetryApicFanDetailsWithDefaults() *NiatelemetryApicFanDetails`

NewNiatelemetryApicFanDetailsWithDefaults instantiates a new NiatelemetryApicFanDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicFanDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicFanDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicFanDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicFanDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicFanDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicFanDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryApicFanDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicFanDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicFanDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicFanDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetModelNumber

`func (o *NiatelemetryApicFanDetails) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *NiatelemetryApicFanDetails) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *NiatelemetryApicFanDetails) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *NiatelemetryApicFanDetails) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetNodeId

`func (o *NiatelemetryApicFanDetails) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NiatelemetryApicFanDetails) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NiatelemetryApicFanDetails) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NiatelemetryApicFanDetails) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicFanDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicFanDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicFanDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicFanDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicFanDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicFanDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicFanDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicFanDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerialNumber

`func (o *NiatelemetryApicFanDetails) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *NiatelemetryApicFanDetails) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *NiatelemetryApicFanDetails) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *NiatelemetryApicFanDetails) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicFanDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicFanDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicFanDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicFanDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetVendorName

`func (o *NiatelemetryApicFanDetails) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *NiatelemetryApicFanDetails) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *NiatelemetryApicFanDetails) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *NiatelemetryApicFanDetails) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetVid

`func (o *NiatelemetryApicFanDetails) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *NiatelemetryApicFanDetails) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *NiatelemetryApicFanDetails) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *NiatelemetryApicFanDetails) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicFanDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicFanDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicFanDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicFanDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


