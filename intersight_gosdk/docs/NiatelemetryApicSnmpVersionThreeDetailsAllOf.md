# NiatelemetryApicSnmpVersionThreeDetailsAllOf

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

### NewNiatelemetryApicSnmpVersionThreeDetailsAllOf

`func NewNiatelemetryApicSnmpVersionThreeDetailsAllOf(classId string, objectType string, ) *NiatelemetryApicSnmpVersionThreeDetailsAllOf`

NewNiatelemetryApicSnmpVersionThreeDetailsAllOf instantiates a new NiatelemetryApicSnmpVersionThreeDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicSnmpVersionThreeDetailsAllOfWithDefaults

`func NewNiatelemetryApicSnmpVersionThreeDetailsAllOfWithDefaults() *NiatelemetryApicSnmpVersionThreeDetailsAllOf`

NewNiatelemetryApicSnmpVersionThreeDetailsAllOfWithDefaults instantiates a new NiatelemetryApicSnmpVersionThreeDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetPrivType() string`

GetPrivType returns the PrivType field if non-nil, zero value otherwise.

### GetPrivTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetPrivTypeOk() (*string, bool)`

GetPrivTypeOk returns a tuple with the PrivType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetPrivType(v string)`

SetPrivType sets PrivType field to given value.

### HasPrivType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasPrivType() bool`

HasPrivType returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicSnmpVersionThreeDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


