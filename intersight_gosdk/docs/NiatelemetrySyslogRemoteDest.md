# NiatelemetrySyslogRemoteDest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SyslogRemoteDest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SyslogRemoteDest"]
**AdminState** | Pointer to **string** | Admin state of Syslog remote dest in APIC. | [optional] 
**CommonPolicy** | Pointer to **string** | Parent common policy for syslog src in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the Syslog remote dest in APIC. | [optional] 
**Host** | Pointer to **string** | Host of Syslog remote dest in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SyslogRsDestGrp** | Pointer to **string** | Dn of sys log src dest grp in APIC. | [optional] 
**SyslogSrc** | Pointer to **string** | Dn of parent syslog src for the syslog dest grp in APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySyslogRemoteDest

`func NewNiatelemetrySyslogRemoteDest(classId string, objectType string, ) *NiatelemetrySyslogRemoteDest`

NewNiatelemetrySyslogRemoteDest instantiates a new NiatelemetrySyslogRemoteDest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySyslogRemoteDestWithDefaults

`func NewNiatelemetrySyslogRemoteDestWithDefaults() *NiatelemetrySyslogRemoteDest`

NewNiatelemetrySyslogRemoteDestWithDefaults instantiates a new NiatelemetrySyslogRemoteDest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySyslogRemoteDest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySyslogRemoteDest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySyslogRemoteDest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySyslogRemoteDest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySyslogRemoteDest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySyslogRemoteDest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *NiatelemetrySyslogRemoteDest) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NiatelemetrySyslogRemoteDest) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NiatelemetrySyslogRemoteDest) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NiatelemetrySyslogRemoteDest) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetCommonPolicy

`func (o *NiatelemetrySyslogRemoteDest) GetCommonPolicy() string`

GetCommonPolicy returns the CommonPolicy field if non-nil, zero value otherwise.

### GetCommonPolicyOk

`func (o *NiatelemetrySyslogRemoteDest) GetCommonPolicyOk() (*string, bool)`

GetCommonPolicyOk returns a tuple with the CommonPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPolicy

`func (o *NiatelemetrySyslogRemoteDest) SetCommonPolicy(v string)`

SetCommonPolicy sets CommonPolicy field to given value.

### HasCommonPolicy

`func (o *NiatelemetrySyslogRemoteDest) HasCommonPolicy() bool`

HasCommonPolicy returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetrySyslogRemoteDest) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySyslogRemoteDest) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySyslogRemoteDest) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySyslogRemoteDest) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHost

`func (o *NiatelemetrySyslogRemoteDest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NiatelemetrySyslogRemoteDest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NiatelemetrySyslogRemoteDest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NiatelemetrySyslogRemoteDest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySyslogRemoteDest) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySyslogRemoteDest) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySyslogRemoteDest) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySyslogRemoteDest) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySyslogRemoteDest) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySyslogRemoteDest) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySyslogRemoteDest) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySyslogRemoteDest) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySyslogRemoteDest) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySyslogRemoteDest) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySyslogRemoteDest) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySyslogRemoteDest) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSyslogRsDestGrp

`func (o *NiatelemetrySyslogRemoteDest) GetSyslogRsDestGrp() string`

GetSyslogRsDestGrp returns the SyslogRsDestGrp field if non-nil, zero value otherwise.

### GetSyslogRsDestGrpOk

`func (o *NiatelemetrySyslogRemoteDest) GetSyslogRsDestGrpOk() (*string, bool)`

GetSyslogRsDestGrpOk returns a tuple with the SyslogRsDestGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogRsDestGrp

`func (o *NiatelemetrySyslogRemoteDest) SetSyslogRsDestGrp(v string)`

SetSyslogRsDestGrp sets SyslogRsDestGrp field to given value.

### HasSyslogRsDestGrp

`func (o *NiatelemetrySyslogRemoteDest) HasSyslogRsDestGrp() bool`

HasSyslogRsDestGrp returns a boolean if a field has been set.

### GetSyslogSrc

`func (o *NiatelemetrySyslogRemoteDest) GetSyslogSrc() string`

GetSyslogSrc returns the SyslogSrc field if non-nil, zero value otherwise.

### GetSyslogSrcOk

`func (o *NiatelemetrySyslogRemoteDest) GetSyslogSrcOk() (*string, bool)`

GetSyslogSrcOk returns a tuple with the SyslogSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSrc

`func (o *NiatelemetrySyslogRemoteDest) SetSyslogSrc(v string)`

SetSyslogSrc sets SyslogSrc field to given value.

### HasSyslogSrc

`func (o *NiatelemetrySyslogRemoteDest) HasSyslogSrc() bool`

HasSyslogSrc returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySyslogRemoteDest) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySyslogRemoteDest) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySyslogRemoteDest) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySyslogRemoteDest) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


