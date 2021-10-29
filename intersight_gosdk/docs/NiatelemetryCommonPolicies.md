# NiatelemetryCommonPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.CommonPolicies"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.CommonPolicies"]
**Dn** | Pointer to **string** | Dn of the Common Policy in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SnmpSrc** | Pointer to **string** | List of Dn of SNMP Src for the above common pol. | [optional] 
**SyslogSrc** | Pointer to **string** | List of Dn of Syslog Src for the above common pol. | [optional] 
**SyslogSysMsg** | Pointer to **string** | List of Dn of Syslog Sys Msg the above common pol. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryCommonPolicies

`func NewNiatelemetryCommonPolicies(classId string, objectType string, ) *NiatelemetryCommonPolicies`

NewNiatelemetryCommonPolicies instantiates a new NiatelemetryCommonPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryCommonPoliciesWithDefaults

`func NewNiatelemetryCommonPoliciesWithDefaults() *NiatelemetryCommonPolicies`

NewNiatelemetryCommonPoliciesWithDefaults instantiates a new NiatelemetryCommonPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryCommonPolicies) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryCommonPolicies) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryCommonPolicies) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryCommonPolicies) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryCommonPolicies) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryCommonPolicies) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryCommonPolicies) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryCommonPolicies) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryCommonPolicies) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryCommonPolicies) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryCommonPolicies) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryCommonPolicies) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryCommonPolicies) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryCommonPolicies) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryCommonPolicies) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryCommonPolicies) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryCommonPolicies) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryCommonPolicies) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryCommonPolicies) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryCommonPolicies) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryCommonPolicies) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryCommonPolicies) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSnmpSrc

`func (o *NiatelemetryCommonPolicies) GetSnmpSrc() string`

GetSnmpSrc returns the SnmpSrc field if non-nil, zero value otherwise.

### GetSnmpSrcOk

`func (o *NiatelemetryCommonPolicies) GetSnmpSrcOk() (*string, bool)`

GetSnmpSrcOk returns a tuple with the SnmpSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpSrc

`func (o *NiatelemetryCommonPolicies) SetSnmpSrc(v string)`

SetSnmpSrc sets SnmpSrc field to given value.

### HasSnmpSrc

`func (o *NiatelemetryCommonPolicies) HasSnmpSrc() bool`

HasSnmpSrc returns a boolean if a field has been set.

### GetSyslogSrc

`func (o *NiatelemetryCommonPolicies) GetSyslogSrc() string`

GetSyslogSrc returns the SyslogSrc field if non-nil, zero value otherwise.

### GetSyslogSrcOk

`func (o *NiatelemetryCommonPolicies) GetSyslogSrcOk() (*string, bool)`

GetSyslogSrcOk returns a tuple with the SyslogSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSrc

`func (o *NiatelemetryCommonPolicies) SetSyslogSrc(v string)`

SetSyslogSrc sets SyslogSrc field to given value.

### HasSyslogSrc

`func (o *NiatelemetryCommonPolicies) HasSyslogSrc() bool`

HasSyslogSrc returns a boolean if a field has been set.

### GetSyslogSysMsg

`func (o *NiatelemetryCommonPolicies) GetSyslogSysMsg() string`

GetSyslogSysMsg returns the SyslogSysMsg field if non-nil, zero value otherwise.

### GetSyslogSysMsgOk

`func (o *NiatelemetryCommonPolicies) GetSyslogSysMsgOk() (*string, bool)`

GetSyslogSysMsgOk returns a tuple with the SyslogSysMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSysMsg

`func (o *NiatelemetryCommonPolicies) SetSyslogSysMsg(v string)`

SetSyslogSysMsg sets SyslogSysMsg field to given value.

### HasSyslogSysMsg

`func (o *NiatelemetryCommonPolicies) HasSyslogSysMsg() bool`

HasSyslogSysMsg returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryCommonPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryCommonPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryCommonPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryCommonPolicies) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


