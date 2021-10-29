# NiatelemetryPodTimeServerPoliciesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.PodTimeServerPolicies"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.PodTimeServerPolicies"]
**PolDn** | Pointer to **string** | Dn of the Time server Pol in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**TimeServerEnableSt** | Pointer to **string** | Admin State of the time server Pol in APIC. | [optional] 
**TimeServers** | Pointer to **string** | Time server of the time server Pol in APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryPodTimeServerPoliciesAllOf

`func NewNiatelemetryPodTimeServerPoliciesAllOf(classId string, objectType string, ) *NiatelemetryPodTimeServerPoliciesAllOf`

NewNiatelemetryPodTimeServerPoliciesAllOf instantiates a new NiatelemetryPodTimeServerPoliciesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPodTimeServerPoliciesAllOfWithDefaults

`func NewNiatelemetryPodTimeServerPoliciesAllOfWithDefaults() *NiatelemetryPodTimeServerPoliciesAllOf`

NewNiatelemetryPodTimeServerPoliciesAllOfWithDefaults instantiates a new NiatelemetryPodTimeServerPoliciesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPolDn

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetPolDn() string`

GetPolDn returns the PolDn field if non-nil, zero value otherwise.

### GetPolDnOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetPolDnOk() (*string, bool)`

GetPolDnOk returns a tuple with the PolDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDn

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetPolDn(v string)`

SetPolDn sets PolDn field to given value.

### HasPolDn

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasPolDn() bool`

HasPolDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetTimeServerEnableSt

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServerEnableSt() string`

GetTimeServerEnableSt returns the TimeServerEnableSt field if non-nil, zero value otherwise.

### GetTimeServerEnableStOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServerEnableStOk() (*string, bool)`

GetTimeServerEnableStOk returns a tuple with the TimeServerEnableSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeServerEnableSt

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetTimeServerEnableSt(v string)`

SetTimeServerEnableSt sets TimeServerEnableSt field to given value.

### HasTimeServerEnableSt

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasTimeServerEnableSt() bool`

HasTimeServerEnableSt returns a boolean if a field has been set.

### GetTimeServers

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServers() string`

GetTimeServers returns the TimeServers field if non-nil, zero value otherwise.

### GetTimeServersOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetTimeServersOk() (*string, bool)`

GetTimeServersOk returns a tuple with the TimeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeServers

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetTimeServers(v string)`

SetTimeServers sets TimeServers field to given value.

### HasTimeServers

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasTimeServers() bool`

HasTimeServers returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPodTimeServerPoliciesAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


