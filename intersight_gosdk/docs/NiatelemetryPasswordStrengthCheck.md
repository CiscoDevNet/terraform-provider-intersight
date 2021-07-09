# NiatelemetryPasswordStrengthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.PasswordStrengthCheck"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.PasswordStrengthCheck"]
**Dn** | Pointer to **string** | Dn for each registering user in APIC. | [optional] 
**PwdStrengthCheck** | Pointer to **string** | Check for password strength per user. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryPasswordStrengthCheck

`func NewNiatelemetryPasswordStrengthCheck(classId string, objectType string, ) *NiatelemetryPasswordStrengthCheck`

NewNiatelemetryPasswordStrengthCheck instantiates a new NiatelemetryPasswordStrengthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPasswordStrengthCheckWithDefaults

`func NewNiatelemetryPasswordStrengthCheckWithDefaults() *NiatelemetryPasswordStrengthCheck`

NewNiatelemetryPasswordStrengthCheckWithDefaults instantiates a new NiatelemetryPasswordStrengthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPasswordStrengthCheck) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPasswordStrengthCheck) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPasswordStrengthCheck) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPasswordStrengthCheck) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPasswordStrengthCheck) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPasswordStrengthCheck) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryPasswordStrengthCheck) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryPasswordStrengthCheck) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryPasswordStrengthCheck) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryPasswordStrengthCheck) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetPwdStrengthCheck

`func (o *NiatelemetryPasswordStrengthCheck) GetPwdStrengthCheck() string`

GetPwdStrengthCheck returns the PwdStrengthCheck field if non-nil, zero value otherwise.

### GetPwdStrengthCheckOk

`func (o *NiatelemetryPasswordStrengthCheck) GetPwdStrengthCheckOk() (*string, bool)`

GetPwdStrengthCheckOk returns a tuple with the PwdStrengthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdStrengthCheck

`func (o *NiatelemetryPasswordStrengthCheck) SetPwdStrengthCheck(v string)`

SetPwdStrengthCheck sets PwdStrengthCheck field to given value.

### HasPwdStrengthCheck

`func (o *NiatelemetryPasswordStrengthCheck) HasPwdStrengthCheck() bool`

HasPwdStrengthCheck returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPasswordStrengthCheck) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPasswordStrengthCheck) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPasswordStrengthCheck) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPasswordStrengthCheck) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPasswordStrengthCheck) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPasswordStrengthCheck) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPasswordStrengthCheck) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPasswordStrengthCheck) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPasswordStrengthCheck) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPasswordStrengthCheck) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPasswordStrengthCheck) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPasswordStrengthCheck) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPasswordStrengthCheck) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPasswordStrengthCheck) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPasswordStrengthCheck) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPasswordStrengthCheck) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


