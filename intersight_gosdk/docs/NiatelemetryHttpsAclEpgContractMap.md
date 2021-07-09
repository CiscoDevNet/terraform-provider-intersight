# NiatelemetryHttpsAclEpgContractMap

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

### NewNiatelemetryHttpsAclEpgContractMap

`func NewNiatelemetryHttpsAclEpgContractMap(classId string, objectType string, ) *NiatelemetryHttpsAclEpgContractMap`

NewNiatelemetryHttpsAclEpgContractMap instantiates a new NiatelemetryHttpsAclEpgContractMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHttpsAclEpgContractMapWithDefaults

`func NewNiatelemetryHttpsAclEpgContractMapWithDefaults() *NiatelemetryHttpsAclEpgContractMap`

NewNiatelemetryHttpsAclEpgContractMapWithDefaults instantiates a new NiatelemetryHttpsAclEpgContractMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryHttpsAclEpgContractMap) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryHttpsAclEpgContractMap) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryHttpsAclEpgContractMap) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryHttpsAclEpgContractMap) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContractName

`func (o *NiatelemetryHttpsAclEpgContractMap) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *NiatelemetryHttpsAclEpgContractMap) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *NiatelemetryHttpsAclEpgContractMap) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryHttpsAclEpgContractMap) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryHttpsAclEpgContractMap) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryHttpsAclEpgContractMap) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEpgName

`func (o *NiatelemetryHttpsAclEpgContractMap) GetEpgName() string`

GetEpgName returns the EpgName field if non-nil, zero value otherwise.

### GetEpgNameOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetEpgNameOk() (*string, bool)`

GetEpgNameOk returns a tuple with the EpgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgName

`func (o *NiatelemetryHttpsAclEpgContractMap) SetEpgName(v string)`

SetEpgName sets EpgName field to given value.

### HasEpgName

`func (o *NiatelemetryHttpsAclEpgContractMap) HasEpgName() bool`

HasEpgName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryHttpsAclEpgContractMap) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryHttpsAclEpgContractMap) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryHttpsAclEpgContractMap) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryHttpsAclEpgContractMap) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryHttpsAclEpgContractMap) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryHttpsAclEpgContractMap) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryHttpsAclEpgContractMap) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetTargetDn

`func (o *NiatelemetryHttpsAclEpgContractMap) GetTargetDn() string`

GetTargetDn returns the TargetDn field if non-nil, zero value otherwise.

### GetTargetDnOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetTargetDnOk() (*string, bool)`

GetTargetDnOk returns a tuple with the TargetDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDn

`func (o *NiatelemetryHttpsAclEpgContractMap) SetTargetDn(v string)`

SetTargetDn sets TargetDn field to given value.

### HasTargetDn

`func (o *NiatelemetryHttpsAclEpgContractMap) HasTargetDn() bool`

HasTargetDn returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryHttpsAclEpgContractMap) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryHttpsAclEpgContractMap) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryHttpsAclEpgContractMap) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryHttpsAclEpgContractMap) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


