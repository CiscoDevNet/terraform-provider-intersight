# NiatelemetryHttpsAclContractDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.HttpsAclContractDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.HttpsAclContractDetails"]
**ConsumerDn** | Pointer to **string** | Consumer Dn of the HTTPS ACL contract children MOs for APIC. | [optional] 
**ContractName** | Pointer to **string** | Name of HTTPS ACL contract for APIC. | [optional] 
**ProviderDn** | Pointer to **string** | Provider dn of the HTTPS ACL contract children MOs for APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryHttpsAclContractDetailsAllOf

`func NewNiatelemetryHttpsAclContractDetailsAllOf(classId string, objectType string, ) *NiatelemetryHttpsAclContractDetailsAllOf`

NewNiatelemetryHttpsAclContractDetailsAllOf instantiates a new NiatelemetryHttpsAclContractDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHttpsAclContractDetailsAllOfWithDefaults

`func NewNiatelemetryHttpsAclContractDetailsAllOfWithDefaults() *NiatelemetryHttpsAclContractDetailsAllOf`

NewNiatelemetryHttpsAclContractDetailsAllOfWithDefaults instantiates a new NiatelemetryHttpsAclContractDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConsumerDn

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetConsumerDn() string`

GetConsumerDn returns the ConsumerDn field if non-nil, zero value otherwise.

### GetConsumerDnOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetConsumerDnOk() (*string, bool)`

GetConsumerDnOk returns a tuple with the ConsumerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerDn

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetConsumerDn(v string)`

SetConsumerDn sets ConsumerDn field to given value.

### HasConsumerDn

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) HasConsumerDn() bool`

HasConsumerDn returns a boolean if a field has been set.

### GetContractName

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetProviderDn

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetProviderDn() string`

GetProviderDn returns the ProviderDn field if non-nil, zero value otherwise.

### GetProviderDnOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetProviderDnOk() (*string, bool)`

GetProviderDnOk returns a tuple with the ProviderDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDn

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetProviderDn(v string)`

SetProviderDn sets ProviderDn field to given value.

### HasProviderDn

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) HasProviderDn() bool`

HasProviderDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryHttpsAclContractDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


