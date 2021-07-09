# NiatelemetryHttpsAclFilterDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.HttpsAclFilterDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.HttpsAclFilterDetails"]
**DestFromPort** | Pointer to **string** | Destination From Port HTTPS ACL EPGs filter MO for APIC. | [optional] 
**DestToPort** | Pointer to **string** | Destination To Port HTTPS ACL EPGs filter MO for APIC. | [optional] 
**Dn** | Pointer to **string** | Dn of the HTTPS ACL EPGs filter MO for APIC. | [optional] 
**FilterName** | Pointer to **string** | Name of HTTPS ACL filter for APIC. | [optional] 
**Prot** | Pointer to **string** | Prot of the HTTPS ACL EPGs filter MO for APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SrcFromPort** | Pointer to **string** | Source From Port HTTPS ACL EPGs filter MO for APIC. | [optional] 
**SrcToPort** | Pointer to **string** | Source To Port HTTPS ACL EPGs filter MO for APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryHttpsAclFilterDetailsAllOf

`func NewNiatelemetryHttpsAclFilterDetailsAllOf(classId string, objectType string, ) *NiatelemetryHttpsAclFilterDetailsAllOf`

NewNiatelemetryHttpsAclFilterDetailsAllOf instantiates a new NiatelemetryHttpsAclFilterDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHttpsAclFilterDetailsAllOfWithDefaults

`func NewNiatelemetryHttpsAclFilterDetailsAllOfWithDefaults() *NiatelemetryHttpsAclFilterDetailsAllOf`

NewNiatelemetryHttpsAclFilterDetailsAllOfWithDefaults instantiates a new NiatelemetryHttpsAclFilterDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestFromPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestFromPort() string`

GetDestFromPort returns the DestFromPort field if non-nil, zero value otherwise.

### GetDestFromPortOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestFromPortOk() (*string, bool)`

GetDestFromPortOk returns a tuple with the DestFromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestFromPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetDestFromPort(v string)`

SetDestFromPort sets DestFromPort field to given value.

### HasDestFromPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasDestFromPort() bool`

HasDestFromPort returns a boolean if a field has been set.

### GetDestToPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestToPort() string`

GetDestToPort returns the DestToPort field if non-nil, zero value otherwise.

### GetDestToPortOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDestToPortOk() (*string, bool)`

GetDestToPortOk returns a tuple with the DestToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestToPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetDestToPort(v string)`

SetDestToPort sets DestToPort field to given value.

### HasDestToPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasDestToPort() bool`

HasDestToPort returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFilterName

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

### GetProt

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetProt() string`

GetProt returns the Prot field if non-nil, zero value otherwise.

### GetProtOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetProtOk() (*string, bool)`

GetProtOk returns a tuple with the Prot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProt

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetProt(v string)`

SetProt sets Prot field to given value.

### HasProt

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasProt() bool`

HasProt returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSrcFromPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcFromPort() string`

GetSrcFromPort returns the SrcFromPort field if non-nil, zero value otherwise.

### GetSrcFromPortOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcFromPortOk() (*string, bool)`

GetSrcFromPortOk returns a tuple with the SrcFromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcFromPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetSrcFromPort(v string)`

SetSrcFromPort sets SrcFromPort field to given value.

### HasSrcFromPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasSrcFromPort() bool`

HasSrcFromPort returns a boolean if a field has been set.

### GetSrcToPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcToPort() string`

GetSrcToPort returns the SrcToPort field if non-nil, zero value otherwise.

### GetSrcToPortOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetSrcToPortOk() (*string, bool)`

GetSrcToPortOk returns a tuple with the SrcToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcToPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetSrcToPort(v string)`

SetSrcToPort sets SrcToPort field to given value.

### HasSrcToPort

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasSrcToPort() bool`

HasSrcToPort returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryHttpsAclFilterDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


