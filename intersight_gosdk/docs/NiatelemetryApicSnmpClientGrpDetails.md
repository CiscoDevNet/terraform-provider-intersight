# NiatelemetryApicSnmpClientGrpDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicSnmpClientGrpDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicSnmpClientGrpDetails"]
**Dn** | Pointer to **string** | Dn of the SNMP community in APIC. | [optional] 
**Name** | Pointer to **string** | Name of SNMP client grp in APIC. | [optional] 
**PolDn** | Pointer to **string** | Dn of the parent SNMP Policy in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**RestrictedClients** | Pointer to **string** | List of address of restricted clients for particular client grp. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicSnmpClientGrpDetails

`func NewNiatelemetryApicSnmpClientGrpDetails(classId string, objectType string, ) *NiatelemetryApicSnmpClientGrpDetails`

NewNiatelemetryApicSnmpClientGrpDetails instantiates a new NiatelemetryApicSnmpClientGrpDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicSnmpClientGrpDetailsWithDefaults

`func NewNiatelemetryApicSnmpClientGrpDetailsWithDefaults() *NiatelemetryApicSnmpClientGrpDetails`

NewNiatelemetryApicSnmpClientGrpDetailsWithDefaults instantiates a new NiatelemetryApicSnmpClientGrpDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolDn

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetPolDn() string`

GetPolDn returns the PolDn field if non-nil, zero value otherwise.

### GetPolDnOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetPolDnOk() (*string, bool)`

GetPolDnOk returns a tuple with the PolDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDn

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetPolDn(v string)`

SetPolDn sets PolDn field to given value.

### HasPolDn

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasPolDn() bool`

HasPolDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRestrictedClients

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRestrictedClients() string`

GetRestrictedClients returns the RestrictedClients field if non-nil, zero value otherwise.

### GetRestrictedClientsOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRestrictedClientsOk() (*string, bool)`

GetRestrictedClientsOk returns a tuple with the RestrictedClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedClients

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetRestrictedClients(v string)`

SetRestrictedClients sets RestrictedClients field to given value.

### HasRestrictedClients

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasRestrictedClients() bool`

HasRestrictedClients returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicSnmpClientGrpDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicSnmpClientGrpDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryApicSnmpClientGrpDetails) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryApicSnmpClientGrpDetails) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


