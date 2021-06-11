# NiatelemetryApicSnmpVersionThreeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicSnmpVersionThreeDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicSnmpVersionThreeDetails"]
**AuthType** | Pointer to **string** | AuthType of SNMP V3 in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of SNMP V3 attribute in APIC. | [optional] 
**Name** | Pointer to **string** | Name of SNMP V3 attribute in APIC. | [optional] 
**PrivType** | Pointer to **string** | PrivType of SNMP V3 in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicSnmpVersionThreeDetails

`func NewNiatelemetryApicSnmpVersionThreeDetails(classId string, objectType string, ) *NiatelemetryApicSnmpVersionThreeDetails`

NewNiatelemetryApicSnmpVersionThreeDetails instantiates a new NiatelemetryApicSnmpVersionThreeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicSnmpVersionThreeDetailsWithDefaults

`func NewNiatelemetryApicSnmpVersionThreeDetailsWithDefaults() *NiatelemetryApicSnmpVersionThreeDetails`

NewNiatelemetryApicSnmpVersionThreeDetailsWithDefaults instantiates a new NiatelemetryApicSnmpVersionThreeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetPrivType() string`

GetPrivType returns the PrivType field if non-nil, zero value otherwise.

### GetPrivTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetPrivTypeOk() (*string, bool)`

GetPrivTypeOk returns a tuple with the PrivType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetPrivType(v string)`

SetPrivType sets PrivType field to given value.

### HasPrivType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasPrivType() bool`

HasPrivType returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicSnmpVersionThreeDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicSnmpVersionThreeDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicSnmpVersionThreeDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


