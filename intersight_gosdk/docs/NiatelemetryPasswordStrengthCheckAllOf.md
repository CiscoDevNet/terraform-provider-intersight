# NiatelemetryPasswordStrengthCheckAllOf

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

### NewNiatelemetryPasswordStrengthCheckAllOf

`func NewNiatelemetryPasswordStrengthCheckAllOf(classId string, objectType string, ) *NiatelemetryPasswordStrengthCheckAllOf`

NewNiatelemetryPasswordStrengthCheckAllOf instantiates a new NiatelemetryPasswordStrengthCheckAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPasswordStrengthCheckAllOfWithDefaults

`func NewNiatelemetryPasswordStrengthCheckAllOfWithDefaults() *NiatelemetryPasswordStrengthCheckAllOf`

NewNiatelemetryPasswordStrengthCheckAllOfWithDefaults instantiates a new NiatelemetryPasswordStrengthCheckAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryPasswordStrengthCheckAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetPwdStrengthCheck

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetPwdStrengthCheck() string`

GetPwdStrengthCheck returns the PwdStrengthCheck field if non-nil, zero value otherwise.

### GetPwdStrengthCheckOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetPwdStrengthCheckOk() (*string, bool)`

GetPwdStrengthCheckOk returns a tuple with the PwdStrengthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdStrengthCheck

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetPwdStrengthCheck(v string)`

SetPwdStrengthCheck sets PwdStrengthCheck field to given value.

### HasPwdStrengthCheck

`func (o *NiatelemetryPasswordStrengthCheckAllOf) HasPwdStrengthCheck() bool`

HasPwdStrengthCheck returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPasswordStrengthCheckAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPasswordStrengthCheckAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPasswordStrengthCheckAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPasswordStrengthCheckAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPasswordStrengthCheckAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPasswordStrengthCheckAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


