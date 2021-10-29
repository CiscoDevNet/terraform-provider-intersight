# NiatelemetrySnmpSrc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SnmpSrc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SnmpSrc"]
**CommonPolicy** | Pointer to **string** | Parent common policy for SNMP Source in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the SNMP Source in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SnmpTrapDest** | Pointer to **string** | List of SNMP trap destination for the above src. | [optional] 
**SnmpTrapDestGrp** | Pointer to **string** | SNMP trap destination group for the above src. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySnmpSrc

`func NewNiatelemetrySnmpSrc(classId string, objectType string, ) *NiatelemetrySnmpSrc`

NewNiatelemetrySnmpSrc instantiates a new NiatelemetrySnmpSrc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySnmpSrcWithDefaults

`func NewNiatelemetrySnmpSrcWithDefaults() *NiatelemetrySnmpSrc`

NewNiatelemetrySnmpSrcWithDefaults instantiates a new NiatelemetrySnmpSrc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySnmpSrc) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySnmpSrc) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySnmpSrc) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySnmpSrc) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySnmpSrc) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySnmpSrc) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommonPolicy

`func (o *NiatelemetrySnmpSrc) GetCommonPolicy() string`

GetCommonPolicy returns the CommonPolicy field if non-nil, zero value otherwise.

### GetCommonPolicyOk

`func (o *NiatelemetrySnmpSrc) GetCommonPolicyOk() (*string, bool)`

GetCommonPolicyOk returns a tuple with the CommonPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPolicy

`func (o *NiatelemetrySnmpSrc) SetCommonPolicy(v string)`

SetCommonPolicy sets CommonPolicy field to given value.

### HasCommonPolicy

`func (o *NiatelemetrySnmpSrc) HasCommonPolicy() bool`

HasCommonPolicy returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetrySnmpSrc) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySnmpSrc) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySnmpSrc) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySnmpSrc) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySnmpSrc) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySnmpSrc) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySnmpSrc) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySnmpSrc) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySnmpSrc) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySnmpSrc) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySnmpSrc) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySnmpSrc) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySnmpSrc) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySnmpSrc) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySnmpSrc) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySnmpSrc) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSnmpTrapDest

`func (o *NiatelemetrySnmpSrc) GetSnmpTrapDest() string`

GetSnmpTrapDest returns the SnmpTrapDest field if non-nil, zero value otherwise.

### GetSnmpTrapDestOk

`func (o *NiatelemetrySnmpSrc) GetSnmpTrapDestOk() (*string, bool)`

GetSnmpTrapDestOk returns a tuple with the SnmpTrapDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapDest

`func (o *NiatelemetrySnmpSrc) SetSnmpTrapDest(v string)`

SetSnmpTrapDest sets SnmpTrapDest field to given value.

### HasSnmpTrapDest

`func (o *NiatelemetrySnmpSrc) HasSnmpTrapDest() bool`

HasSnmpTrapDest returns a boolean if a field has been set.

### GetSnmpTrapDestGrp

`func (o *NiatelemetrySnmpSrc) GetSnmpTrapDestGrp() string`

GetSnmpTrapDestGrp returns the SnmpTrapDestGrp field if non-nil, zero value otherwise.

### GetSnmpTrapDestGrpOk

`func (o *NiatelemetrySnmpSrc) GetSnmpTrapDestGrpOk() (*string, bool)`

GetSnmpTrapDestGrpOk returns a tuple with the SnmpTrapDestGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapDestGrp

`func (o *NiatelemetrySnmpSrc) SetSnmpTrapDestGrp(v string)`

SetSnmpTrapDestGrp sets SnmpTrapDestGrp field to given value.

### HasSnmpTrapDestGrp

`func (o *NiatelemetrySnmpSrc) HasSnmpTrapDestGrp() bool`

HasSnmpTrapDestGrp returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySnmpSrc) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySnmpSrc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySnmpSrc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySnmpSrc) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


