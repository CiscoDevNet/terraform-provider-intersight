# NiatelemetrySyslogSysMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SyslogSysMsg"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SyslogSysMsg"]
**CommonPolicy** | Pointer to **string** | Parent common policy for syslog msg in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the Syslog System msg in APIC. | [optional] 
**FacilityFilter** | Pointer to **string** | List of Facility filter DN of the Syslog System msg in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySyslogSysMsg

`func NewNiatelemetrySyslogSysMsg(classId string, objectType string, ) *NiatelemetrySyslogSysMsg`

NewNiatelemetrySyslogSysMsg instantiates a new NiatelemetrySyslogSysMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySyslogSysMsgWithDefaults

`func NewNiatelemetrySyslogSysMsgWithDefaults() *NiatelemetrySyslogSysMsg`

NewNiatelemetrySyslogSysMsgWithDefaults instantiates a new NiatelemetrySyslogSysMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySyslogSysMsg) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySyslogSysMsg) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySyslogSysMsg) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySyslogSysMsg) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySyslogSysMsg) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySyslogSysMsg) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommonPolicy

`func (o *NiatelemetrySyslogSysMsg) GetCommonPolicy() string`

GetCommonPolicy returns the CommonPolicy field if non-nil, zero value otherwise.

### GetCommonPolicyOk

`func (o *NiatelemetrySyslogSysMsg) GetCommonPolicyOk() (*string, bool)`

GetCommonPolicyOk returns a tuple with the CommonPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPolicy

`func (o *NiatelemetrySyslogSysMsg) SetCommonPolicy(v string)`

SetCommonPolicy sets CommonPolicy field to given value.

### HasCommonPolicy

`func (o *NiatelemetrySyslogSysMsg) HasCommonPolicy() bool`

HasCommonPolicy returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetrySyslogSysMsg) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySyslogSysMsg) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySyslogSysMsg) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySyslogSysMsg) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFacilityFilter

`func (o *NiatelemetrySyslogSysMsg) GetFacilityFilter() string`

GetFacilityFilter returns the FacilityFilter field if non-nil, zero value otherwise.

### GetFacilityFilterOk

`func (o *NiatelemetrySyslogSysMsg) GetFacilityFilterOk() (*string, bool)`

GetFacilityFilterOk returns a tuple with the FacilityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityFilter

`func (o *NiatelemetrySyslogSysMsg) SetFacilityFilter(v string)`

SetFacilityFilter sets FacilityFilter field to given value.

### HasFacilityFilter

`func (o *NiatelemetrySyslogSysMsg) HasFacilityFilter() bool`

HasFacilityFilter returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySyslogSysMsg) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySyslogSysMsg) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySyslogSysMsg) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySyslogSysMsg) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySyslogSysMsg) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySyslogSysMsg) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySyslogSysMsg) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySyslogSysMsg) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySyslogSysMsg) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySyslogSysMsg) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySyslogSysMsg) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySyslogSysMsg) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySyslogSysMsg) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySyslogSysMsg) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySyslogSysMsg) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySyslogSysMsg) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


