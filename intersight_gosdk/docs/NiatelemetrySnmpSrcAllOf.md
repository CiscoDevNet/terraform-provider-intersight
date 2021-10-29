# NiatelemetrySnmpSrcAllOf

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

### NewNiatelemetrySnmpSrcAllOf

`func NewNiatelemetrySnmpSrcAllOf(classId string, objectType string, ) *NiatelemetrySnmpSrcAllOf`

NewNiatelemetrySnmpSrcAllOf instantiates a new NiatelemetrySnmpSrcAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySnmpSrcAllOfWithDefaults

`func NewNiatelemetrySnmpSrcAllOfWithDefaults() *NiatelemetrySnmpSrcAllOf`

NewNiatelemetrySnmpSrcAllOfWithDefaults instantiates a new NiatelemetrySnmpSrcAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySnmpSrcAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySnmpSrcAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySnmpSrcAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySnmpSrcAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySnmpSrcAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySnmpSrcAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommonPolicy

`func (o *NiatelemetrySnmpSrcAllOf) GetCommonPolicy() string`

GetCommonPolicy returns the CommonPolicy field if non-nil, zero value otherwise.

### GetCommonPolicyOk

`func (o *NiatelemetrySnmpSrcAllOf) GetCommonPolicyOk() (*string, bool)`

GetCommonPolicyOk returns a tuple with the CommonPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPolicy

`func (o *NiatelemetrySnmpSrcAllOf) SetCommonPolicy(v string)`

SetCommonPolicy sets CommonPolicy field to given value.

### HasCommonPolicy

`func (o *NiatelemetrySnmpSrcAllOf) HasCommonPolicy() bool`

HasCommonPolicy returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetrySnmpSrcAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySnmpSrcAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySnmpSrcAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySnmpSrcAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySnmpSrcAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySnmpSrcAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySnmpSrcAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySnmpSrcAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySnmpSrcAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySnmpSrcAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySnmpSrcAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySnmpSrcAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySnmpSrcAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySnmpSrcAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySnmpSrcAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySnmpSrcAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSnmpTrapDest

`func (o *NiatelemetrySnmpSrcAllOf) GetSnmpTrapDest() string`

GetSnmpTrapDest returns the SnmpTrapDest field if non-nil, zero value otherwise.

### GetSnmpTrapDestOk

`func (o *NiatelemetrySnmpSrcAllOf) GetSnmpTrapDestOk() (*string, bool)`

GetSnmpTrapDestOk returns a tuple with the SnmpTrapDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapDest

`func (o *NiatelemetrySnmpSrcAllOf) SetSnmpTrapDest(v string)`

SetSnmpTrapDest sets SnmpTrapDest field to given value.

### HasSnmpTrapDest

`func (o *NiatelemetrySnmpSrcAllOf) HasSnmpTrapDest() bool`

HasSnmpTrapDest returns a boolean if a field has been set.

### GetSnmpTrapDestGrp

`func (o *NiatelemetrySnmpSrcAllOf) GetSnmpTrapDestGrp() string`

GetSnmpTrapDestGrp returns the SnmpTrapDestGrp field if non-nil, zero value otherwise.

### GetSnmpTrapDestGrpOk

`func (o *NiatelemetrySnmpSrcAllOf) GetSnmpTrapDestGrpOk() (*string, bool)`

GetSnmpTrapDestGrpOk returns a tuple with the SnmpTrapDestGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapDestGrp

`func (o *NiatelemetrySnmpSrcAllOf) SetSnmpTrapDestGrp(v string)`

SetSnmpTrapDestGrp sets SnmpTrapDestGrp field to given value.

### HasSnmpTrapDestGrp

`func (o *NiatelemetrySnmpSrcAllOf) HasSnmpTrapDestGrp() bool`

HasSnmpTrapDestGrp returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySnmpSrcAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySnmpSrcAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySnmpSrcAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySnmpSrcAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


