# NiatelemetryHttpsAclEpgContractMapAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.HttpsAclEpgContractMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.HttpsAclEpgContractMap"]
**ContractName** | Pointer to **string** | Name of HTTPS ACL contract for APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the HTTPS ACL contract provider for APIC. | [optional] 
**EpgName** | Pointer to **string** | Name of EPGs of the HTTPS ACL contract for APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**TargetDn** | Pointer to **string** | TDn of the HTTPS ACL contract provider for APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryHttpsAclEpgContractMapAllOf

`func NewNiatelemetryHttpsAclEpgContractMapAllOf(classId string, objectType string, ) *NiatelemetryHttpsAclEpgContractMapAllOf`

NewNiatelemetryHttpsAclEpgContractMapAllOf instantiates a new NiatelemetryHttpsAclEpgContractMapAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHttpsAclEpgContractMapAllOfWithDefaults

`func NewNiatelemetryHttpsAclEpgContractMapAllOfWithDefaults() *NiatelemetryHttpsAclEpgContractMapAllOf`

NewNiatelemetryHttpsAclEpgContractMapAllOfWithDefaults instantiates a new NiatelemetryHttpsAclEpgContractMapAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContractName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEpgName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetEpgName() string`

GetEpgName returns the EpgName field if non-nil, zero value otherwise.

### GetEpgNameOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetEpgNameOk() (*string, bool)`

GetEpgNameOk returns a tuple with the EpgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetEpgName(v string)`

SetEpgName sets EpgName field to given value.

### HasEpgName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasEpgName() bool`

HasEpgName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetTargetDn

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetTargetDn() string`

GetTargetDn returns the TargetDn field if non-nil, zero value otherwise.

### GetTargetDnOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetTargetDnOk() (*string, bool)`

GetTargetDnOk returns a tuple with the TargetDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDn

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetTargetDn(v string)`

SetTargetDn sets TargetDn field to given value.

### HasTargetDn

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasTargetDn() bool`

HasTargetDn returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


