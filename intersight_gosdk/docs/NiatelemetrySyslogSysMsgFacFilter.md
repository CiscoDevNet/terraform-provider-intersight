# NiatelemetrySyslogSysMsgFacFilter

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
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySyslogSysMsgFacFilter

`func NewNiatelemetrySyslogSysMsgFacFilter(classId string, objectType string, ) *NiatelemetrySyslogSysMsgFacFilter`

NewNiatelemetrySyslogSysMsgFacFilter instantiates a new NiatelemetrySyslogSysMsgFacFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySyslogSysMsgFacFilterWithDefaults

`func NewNiatelemetrySyslogSysMsgFacFilterWithDefaults() *NiatelemetrySyslogSysMsgFacFilter`

NewNiatelemetrySyslogSysMsgFacFilterWithDefaults instantiates a new NiatelemetrySyslogSysMsgFacFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommonPolicy

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetCommonPolicy() string`

GetCommonPolicy returns the CommonPolicy field if non-nil, zero value otherwise.

### GetCommonPolicyOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetCommonPolicyOk() (*string, bool)`

GetCommonPolicyOk returns a tuple with the CommonPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPolicy

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetCommonPolicy(v string)`

SetCommonPolicy sets CommonPolicy field to given value.

### HasCommonPolicy

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasCommonPolicy() bool`

HasCommonPolicy returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFacility

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetMinSev

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetMinSev() string`

GetMinSev returns the MinSev field if non-nil, zero value otherwise.

### GetMinSevOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetMinSevOk() (*string, bool)`

GetMinSevOk returns a tuple with the MinSev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSev

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetMinSev(v string)`

SetMinSev sets MinSev field to given value.

### HasMinSev

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasMinSev() bool`

HasMinSev returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSyslogSysMsg

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetSyslogSysMsg() string`

GetSyslogSysMsg returns the SyslogSysMsg field if non-nil, zero value otherwise.

### GetSyslogSysMsgOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetSyslogSysMsgOk() (*string, bool)`

GetSyslogSysMsgOk returns a tuple with the SyslogSysMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSysMsg

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetSyslogSysMsg(v string)`

SetSyslogSysMsg sets SyslogSysMsg field to given value.

### HasSyslogSysMsg

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasSyslogSysMsg() bool`

HasSyslogSysMsg returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySyslogSysMsgFacFilter) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySyslogSysMsgFacFilter) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetrySyslogSysMsgFacFilter) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetrySyslogSysMsgFacFilter) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


