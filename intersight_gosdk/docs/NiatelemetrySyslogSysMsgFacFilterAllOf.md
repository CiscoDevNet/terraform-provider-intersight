# NiatelemetrySyslogSysMsgFacFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SyslogSysMsgFacFilter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SyslogSysMsgFacFilter"]
**CommonPolicy** | Pointer to **string** | Parent common policy for syslog system msg in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the Syslog System msg facility filter in APIC. | [optional] 
**Facility** | Pointer to **string** | Facility of Syslog System msg facility filter in APIC. | [optional] 
**MinSev** | Pointer to **string** | Minimum severity of Syslog System msg facility filter in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SyslogSysMsg** | Pointer to **string** | Parent syslog msg for syslog sys msg facility filter in APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySyslogSysMsgFacFilterAllOf

`func NewNiatelemetrySyslogSysMsgFacFilterAllOf(classId string, objectType string, ) *NiatelemetrySyslogSysMsgFacFilterAllOf`

NewNiatelemetrySyslogSysMsgFacFilterAllOf instantiates a new NiatelemetrySyslogSysMsgFacFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySyslogSysMsgFacFilterAllOfWithDefaults

`func NewNiatelemetrySyslogSysMsgFacFilterAllOfWithDefaults() *NiatelemetrySyslogSysMsgFacFilterAllOf`

NewNiatelemetrySyslogSysMsgFacFilterAllOfWithDefaults instantiates a new NiatelemetrySyslogSysMsgFacFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommonPolicy

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetCommonPolicy() string`

GetCommonPolicy returns the CommonPolicy field if non-nil, zero value otherwise.

### GetCommonPolicyOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetCommonPolicyOk() (*string, bool)`

GetCommonPolicyOk returns a tuple with the CommonPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPolicy

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetCommonPolicy(v string)`

SetCommonPolicy sets CommonPolicy field to given value.

### HasCommonPolicy

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasCommonPolicy() bool`

HasCommonPolicy returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFacility

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetMinSev

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetMinSev() string`

GetMinSev returns the MinSev field if non-nil, zero value otherwise.

### GetMinSevOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetMinSevOk() (*string, bool)`

GetMinSevOk returns a tuple with the MinSev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSev

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetMinSev(v string)`

SetMinSev sets MinSev field to given value.

### HasMinSev

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasMinSev() bool`

HasMinSev returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSyslogSysMsg

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSyslogSysMsg() string`

GetSyslogSysMsg returns the SyslogSysMsg field if non-nil, zero value otherwise.

### GetSyslogSysMsgOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetSyslogSysMsgOk() (*string, bool)`

GetSyslogSysMsgOk returns a tuple with the SyslogSysMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSysMsg

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetSyslogSysMsg(v string)`

SetSyslogSysMsg sets SyslogSysMsg field to given value.

### HasSyslogSysMsg

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasSyslogSysMsg() bool`

HasSyslogSysMsg returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySyslogSysMsgFacFilterAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


