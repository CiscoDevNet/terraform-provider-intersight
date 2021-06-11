# NiatelemetrySshVersionTwoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SshVersionTwo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SshVersionTwo"]
**AdminSt** | Pointer to **string** | Admin state of SSH V2 in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of SSH V2 attribute in APIC. | [optional] 
**PasswordAuth** | Pointer to **string** | Password auth for SSH V2 in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SshCiphers** | Pointer to **string** | SSH Ciphers for SSH V2 in APIC. | [optional] 
**SshMacs** | Pointer to **string** | SSH MACS for SSH V2 in APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySshVersionTwoAllOf

`func NewNiatelemetrySshVersionTwoAllOf(classId string, objectType string, ) *NiatelemetrySshVersionTwoAllOf`

NewNiatelemetrySshVersionTwoAllOf instantiates a new NiatelemetrySshVersionTwoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySshVersionTwoAllOfWithDefaults

`func NewNiatelemetrySshVersionTwoAllOfWithDefaults() *NiatelemetrySshVersionTwoAllOf`

NewNiatelemetrySshVersionTwoAllOfWithDefaults instantiates a new NiatelemetrySshVersionTwoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySshVersionTwoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySshVersionTwoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySshVersionTwoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySshVersionTwoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSt

`func (o *NiatelemetrySshVersionTwoAllOf) GetAdminSt() string`

GetAdminSt returns the AdminSt field if non-nil, zero value otherwise.

### GetAdminStOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetAdminStOk() (*string, bool)`

GetAdminStOk returns a tuple with the AdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSt

`func (o *NiatelemetrySshVersionTwoAllOf) SetAdminSt(v string)`

SetAdminSt sets AdminSt field to given value.

### HasAdminSt

`func (o *NiatelemetrySshVersionTwoAllOf) HasAdminSt() bool`

HasAdminSt returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetrySshVersionTwoAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySshVersionTwoAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySshVersionTwoAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetPasswordAuth

`func (o *NiatelemetrySshVersionTwoAllOf) GetPasswordAuth() string`

GetPasswordAuth returns the PasswordAuth field if non-nil, zero value otherwise.

### GetPasswordAuthOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetPasswordAuthOk() (*string, bool)`

GetPasswordAuthOk returns a tuple with the PasswordAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAuth

`func (o *NiatelemetrySshVersionTwoAllOf) SetPasswordAuth(v string)`

SetPasswordAuth sets PasswordAuth field to given value.

### HasPasswordAuth

`func (o *NiatelemetrySshVersionTwoAllOf) HasPasswordAuth() bool`

HasPasswordAuth returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySshVersionTwoAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySshVersionTwoAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySshVersionTwoAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySshVersionTwoAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySshVersionTwoAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySshVersionTwoAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySshVersionTwoAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySshVersionTwoAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySshVersionTwoAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSshCiphers

`func (o *NiatelemetrySshVersionTwoAllOf) GetSshCiphers() string`

GetSshCiphers returns the SshCiphers field if non-nil, zero value otherwise.

### GetSshCiphersOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetSshCiphersOk() (*string, bool)`

GetSshCiphersOk returns a tuple with the SshCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCiphers

`func (o *NiatelemetrySshVersionTwoAllOf) SetSshCiphers(v string)`

SetSshCiphers sets SshCiphers field to given value.

### HasSshCiphers

`func (o *NiatelemetrySshVersionTwoAllOf) HasSshCiphers() bool`

HasSshCiphers returns a boolean if a field has been set.

### GetSshMacs

`func (o *NiatelemetrySshVersionTwoAllOf) GetSshMacs() string`

GetSshMacs returns the SshMacs field if non-nil, zero value otherwise.

### GetSshMacsOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetSshMacsOk() (*string, bool)`

GetSshMacsOk returns a tuple with the SshMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshMacs

`func (o *NiatelemetrySshVersionTwoAllOf) SetSshMacs(v string)`

SetSshMacs sets SshMacs field to given value.

### HasSshMacs

`func (o *NiatelemetrySshVersionTwoAllOf) HasSshMacs() bool`

HasSshMacs returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySshVersionTwoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySshVersionTwoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySshVersionTwoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySshVersionTwoAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


