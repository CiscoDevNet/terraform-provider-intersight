# NiatelemetryPodSnmpPoliciesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.PodSnmpPolicies"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.PodSnmpPolicies"]
**AdminSt** | Pointer to **string** | Admin State of the Snmp Pol in APIC. | [optional] 
**PolDn** | Pointer to **string** | Dn of the Snmp Pol in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SnmpClientGrp** | Pointer to **string** | List of Dn of the SNMP Client grp in APIC. | [optional] 
**SnmpCommunity** | Pointer to **string** | List of Dn of the SNMP Community in APIC. | [optional] 
**SnmpTrapFwdServer** | Pointer to **string** | List of Dn of the SNMP Trap Fwd Server in APIC. | [optional] 
**SnmpUser** | Pointer to **string** | List of Dn of the SNMP user in APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryPodSnmpPoliciesAllOf

`func NewNiatelemetryPodSnmpPoliciesAllOf(classId string, objectType string, ) *NiatelemetryPodSnmpPoliciesAllOf`

NewNiatelemetryPodSnmpPoliciesAllOf instantiates a new NiatelemetryPodSnmpPoliciesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPodSnmpPoliciesAllOfWithDefaults

`func NewNiatelemetryPodSnmpPoliciesAllOfWithDefaults() *NiatelemetryPodSnmpPoliciesAllOf`

NewNiatelemetryPodSnmpPoliciesAllOfWithDefaults instantiates a new NiatelemetryPodSnmpPoliciesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSt

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetAdminSt() string`

GetAdminSt returns the AdminSt field if non-nil, zero value otherwise.

### GetAdminStOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetAdminStOk() (*string, bool)`

GetAdminStOk returns a tuple with the AdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSt

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetAdminSt(v string)`

SetAdminSt sets AdminSt field to given value.

### HasAdminSt

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasAdminSt() bool`

HasAdminSt returns a boolean if a field has been set.

### GetPolDn

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetPolDn() string`

GetPolDn returns the PolDn field if non-nil, zero value otherwise.

### GetPolDnOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetPolDnOk() (*string, bool)`

GetPolDnOk returns a tuple with the PolDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDn

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetPolDn(v string)`

SetPolDn sets PolDn field to given value.

### HasPolDn

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasPolDn() bool`

HasPolDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSnmpClientGrp

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpClientGrp() string`

GetSnmpClientGrp returns the SnmpClientGrp field if non-nil, zero value otherwise.

### GetSnmpClientGrpOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpClientGrpOk() (*string, bool)`

GetSnmpClientGrpOk returns a tuple with the SnmpClientGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpClientGrp

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpClientGrp(v string)`

SetSnmpClientGrp sets SnmpClientGrp field to given value.

### HasSnmpClientGrp

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpClientGrp() bool`

HasSnmpClientGrp returns a boolean if a field has been set.

### GetSnmpCommunity

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpCommunity() string`

GetSnmpCommunity returns the SnmpCommunity field if non-nil, zero value otherwise.

### GetSnmpCommunityOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpCommunityOk() (*string, bool)`

GetSnmpCommunityOk returns a tuple with the SnmpCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCommunity

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpCommunity(v string)`

SetSnmpCommunity sets SnmpCommunity field to given value.

### HasSnmpCommunity

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpCommunity() bool`

HasSnmpCommunity returns a boolean if a field has been set.

### GetSnmpTrapFwdServer

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpTrapFwdServer() string`

GetSnmpTrapFwdServer returns the SnmpTrapFwdServer field if non-nil, zero value otherwise.

### GetSnmpTrapFwdServerOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpTrapFwdServerOk() (*string, bool)`

GetSnmpTrapFwdServerOk returns a tuple with the SnmpTrapFwdServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapFwdServer

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpTrapFwdServer(v string)`

SetSnmpTrapFwdServer sets SnmpTrapFwdServer field to given value.

### HasSnmpTrapFwdServer

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpTrapFwdServer() bool`

HasSnmpTrapFwdServer returns a boolean if a field has been set.

### GetSnmpUser

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpUser() string`

GetSnmpUser returns the SnmpUser field if non-nil, zero value otherwise.

### GetSnmpUserOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetSnmpUserOk() (*string, bool)`

GetSnmpUserOk returns a tuple with the SnmpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpUser

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetSnmpUser(v string)`

SetSnmpUser sets SnmpUser field to given value.

### HasSnmpUser

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasSnmpUser() bool`

HasSnmpUser returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPodSnmpPoliciesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPodSnmpPoliciesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPodSnmpPoliciesAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


