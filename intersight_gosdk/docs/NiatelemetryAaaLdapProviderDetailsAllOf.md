# NiatelemetryAaaLdapProviderDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.AaaLdapProviderDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.AaaLdapProviderDetails"]
**BaseDn** | Pointer to **string** | Base dn of the AAA ldap provider in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the AAA ldap provider in APIC. | [optional] 
**Port** | Pointer to **string** | Port of the AAA ldap provider in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**RootDn** | Pointer to **string** | Root dn of the AAA ldap provider in APIC. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryAaaLdapProviderDetailsAllOf

`func NewNiatelemetryAaaLdapProviderDetailsAllOf(classId string, objectType string, ) *NiatelemetryAaaLdapProviderDetailsAllOf`

NewNiatelemetryAaaLdapProviderDetailsAllOf instantiates a new NiatelemetryAaaLdapProviderDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryAaaLdapProviderDetailsAllOfWithDefaults

`func NewNiatelemetryAaaLdapProviderDetailsAllOfWithDefaults() *NiatelemetryAaaLdapProviderDetailsAllOf`

NewNiatelemetryAaaLdapProviderDetailsAllOfWithDefaults instantiates a new NiatelemetryAaaLdapProviderDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBaseDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetPort

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRootDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRootDn() string`

GetRootDn returns the RootDn field if non-nil, zero value otherwise.

### GetRootDnOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRootDnOk() (*string, bool)`

GetRootDnOk returns a tuple with the RootDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetRootDn(v string)`

SetRootDn sets RootDn field to given value.

### HasRootDn

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasRootDn() bool`

HasRootDn returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryAaaLdapProviderDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


