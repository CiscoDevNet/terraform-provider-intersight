# NiatelemetryPodSnmpPolicies

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
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryPodSnmpPolicies

`func NewNiatelemetryPodSnmpPolicies(classId string, objectType string, ) *NiatelemetryPodSnmpPolicies`

NewNiatelemetryPodSnmpPolicies instantiates a new NiatelemetryPodSnmpPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPodSnmpPoliciesWithDefaults

`func NewNiatelemetryPodSnmpPoliciesWithDefaults() *NiatelemetryPodSnmpPolicies`

NewNiatelemetryPodSnmpPoliciesWithDefaults instantiates a new NiatelemetryPodSnmpPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPodSnmpPolicies) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPodSnmpPolicies) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPodSnmpPolicies) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPodSnmpPolicies) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPodSnmpPolicies) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPodSnmpPolicies) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSt

`func (o *NiatelemetryPodSnmpPolicies) GetAdminSt() string`

GetAdminSt returns the AdminSt field if non-nil, zero value otherwise.

### GetAdminStOk

`func (o *NiatelemetryPodSnmpPolicies) GetAdminStOk() (*string, bool)`

GetAdminStOk returns a tuple with the AdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSt

`func (o *NiatelemetryPodSnmpPolicies) SetAdminSt(v string)`

SetAdminSt sets AdminSt field to given value.

### HasAdminSt

`func (o *NiatelemetryPodSnmpPolicies) HasAdminSt() bool`

HasAdminSt returns a boolean if a field has been set.

### GetPolDn

`func (o *NiatelemetryPodSnmpPolicies) GetPolDn() string`

GetPolDn returns the PolDn field if non-nil, zero value otherwise.

### GetPolDnOk

`func (o *NiatelemetryPodSnmpPolicies) GetPolDnOk() (*string, bool)`

GetPolDnOk returns a tuple with the PolDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDn

`func (o *NiatelemetryPodSnmpPolicies) SetPolDn(v string)`

SetPolDn sets PolDn field to given value.

### HasPolDn

`func (o *NiatelemetryPodSnmpPolicies) HasPolDn() bool`

HasPolDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPodSnmpPolicies) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPodSnmpPolicies) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPodSnmpPolicies) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPodSnmpPolicies) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPodSnmpPolicies) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPodSnmpPolicies) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPodSnmpPolicies) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPodSnmpPolicies) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPodSnmpPolicies) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPodSnmpPolicies) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPodSnmpPolicies) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPodSnmpPolicies) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSnmpClientGrp

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpClientGrp() string`

GetSnmpClientGrp returns the SnmpClientGrp field if non-nil, zero value otherwise.

### GetSnmpClientGrpOk

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpClientGrpOk() (*string, bool)`

GetSnmpClientGrpOk returns a tuple with the SnmpClientGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpClientGrp

`func (o *NiatelemetryPodSnmpPolicies) SetSnmpClientGrp(v string)`

SetSnmpClientGrp sets SnmpClientGrp field to given value.

### HasSnmpClientGrp

`func (o *NiatelemetryPodSnmpPolicies) HasSnmpClientGrp() bool`

HasSnmpClientGrp returns a boolean if a field has been set.

### GetSnmpCommunity

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpCommunity() string`

GetSnmpCommunity returns the SnmpCommunity field if non-nil, zero value otherwise.

### GetSnmpCommunityOk

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpCommunityOk() (*string, bool)`

GetSnmpCommunityOk returns a tuple with the SnmpCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCommunity

`func (o *NiatelemetryPodSnmpPolicies) SetSnmpCommunity(v string)`

SetSnmpCommunity sets SnmpCommunity field to given value.

### HasSnmpCommunity

`func (o *NiatelemetryPodSnmpPolicies) HasSnmpCommunity() bool`

HasSnmpCommunity returns a boolean if a field has been set.

### GetSnmpTrapFwdServer

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpTrapFwdServer() string`

GetSnmpTrapFwdServer returns the SnmpTrapFwdServer field if non-nil, zero value otherwise.

### GetSnmpTrapFwdServerOk

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpTrapFwdServerOk() (*string, bool)`

GetSnmpTrapFwdServerOk returns a tuple with the SnmpTrapFwdServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapFwdServer

`func (o *NiatelemetryPodSnmpPolicies) SetSnmpTrapFwdServer(v string)`

SetSnmpTrapFwdServer sets SnmpTrapFwdServer field to given value.

### HasSnmpTrapFwdServer

`func (o *NiatelemetryPodSnmpPolicies) HasSnmpTrapFwdServer() bool`

HasSnmpTrapFwdServer returns a boolean if a field has been set.

### GetSnmpUser

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpUser() string`

GetSnmpUser returns the SnmpUser field if non-nil, zero value otherwise.

### GetSnmpUserOk

`func (o *NiatelemetryPodSnmpPolicies) GetSnmpUserOk() (*string, bool)`

GetSnmpUserOk returns a tuple with the SnmpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpUser

`func (o *NiatelemetryPodSnmpPolicies) SetSnmpUser(v string)`

SetSnmpUser sets SnmpUser field to given value.

### HasSnmpUser

`func (o *NiatelemetryPodSnmpPolicies) HasSnmpUser() bool`

HasSnmpUser returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPodSnmpPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPodSnmpPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPodSnmpPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPodSnmpPolicies) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryPodSnmpPolicies) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryPodSnmpPolicies) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


