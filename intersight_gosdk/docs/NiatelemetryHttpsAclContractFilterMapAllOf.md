# NiatelemetryHttpsAclContractFilterMapAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.HttpsAclContractFilterMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.HttpsAclContractFilterMap"]
**ContractName** | Pointer to **string** | Name of HTTPS ACL contract for APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the HTTPS ACL EPGs subject filter relation MO for APIC. | [optional] 
**FilterName** | Pointer to **string** | Name of HTTPS ACL filter for APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryHttpsAclContractFilterMapAllOf

`func NewNiatelemetryHttpsAclContractFilterMapAllOf(classId string, objectType string, ) *NiatelemetryHttpsAclContractFilterMapAllOf`

NewNiatelemetryHttpsAclContractFilterMapAllOf instantiates a new NiatelemetryHttpsAclContractFilterMapAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHttpsAclContractFilterMapAllOfWithDefaults

`func NewNiatelemetryHttpsAclContractFilterMapAllOfWithDefaults() *NiatelemetryHttpsAclContractFilterMapAllOf`

NewNiatelemetryHttpsAclContractFilterMapAllOfWithDefaults instantiates a new NiatelemetryHttpsAclContractFilterMapAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContractName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFilterName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryHttpsAclContractFilterMapAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


