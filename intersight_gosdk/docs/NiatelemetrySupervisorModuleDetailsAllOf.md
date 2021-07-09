# NiatelemetrySupervisorModuleDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SupervisorModuleDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SupervisorModuleDetails"]
**Dn** | Pointer to **string** | Dn of the supervisor module in APIC. | [optional] 
**HwVer** | Pointer to **string** | Hardware version of supervisor module. | [optional] 
**Model** | Pointer to **string** | Model of the supervisor module. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of the supervisor module. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySupervisorModuleDetailsAllOf

`func NewNiatelemetrySupervisorModuleDetailsAllOf(classId string, objectType string, ) *NiatelemetrySupervisorModuleDetailsAllOf`

NewNiatelemetrySupervisorModuleDetailsAllOf instantiates a new NiatelemetrySupervisorModuleDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySupervisorModuleDetailsAllOfWithDefaults

`func NewNiatelemetrySupervisorModuleDetailsAllOfWithDefaults() *NiatelemetrySupervisorModuleDetailsAllOf`

NewNiatelemetrySupervisorModuleDetailsAllOfWithDefaults instantiates a new NiatelemetrySupervisorModuleDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHwVer

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetHwVer() string`

GetHwVer returns the HwVer field if non-nil, zero value otherwise.

### GetHwVerOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetHwVerOk() (*string, bool)`

GetHwVerOk returns a tuple with the HwVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwVer

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetHwVer(v string)`

SetHwVer sets HwVer field to given value.

### HasHwVer

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasHwVer() bool`

HasHwVer returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySupervisorModuleDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


