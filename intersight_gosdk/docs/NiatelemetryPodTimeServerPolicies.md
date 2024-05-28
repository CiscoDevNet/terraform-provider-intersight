# NiatelemetryPodTimeServerPolicies

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
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryPodTimeServerPolicies

`func NewNiatelemetryPodTimeServerPolicies(classId string, objectType string, ) *NiatelemetryPodTimeServerPolicies`

NewNiatelemetryPodTimeServerPolicies instantiates a new NiatelemetryPodTimeServerPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPodTimeServerPoliciesWithDefaults

`func NewNiatelemetryPodTimeServerPoliciesWithDefaults() *NiatelemetryPodTimeServerPolicies`

NewNiatelemetryPodTimeServerPoliciesWithDefaults instantiates a new NiatelemetryPodTimeServerPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPodTimeServerPolicies) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPodTimeServerPolicies) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPodTimeServerPolicies) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPodTimeServerPolicies) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPodTimeServerPolicies) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPodTimeServerPolicies) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPolDn

`func (o *NiatelemetryPodTimeServerPolicies) GetPolDn() string`

GetPolDn returns the PolDn field if non-nil, zero value otherwise.

### GetPolDnOk

`func (o *NiatelemetryPodTimeServerPolicies) GetPolDnOk() (*string, bool)`

GetPolDnOk returns a tuple with the PolDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDn

`func (o *NiatelemetryPodTimeServerPolicies) SetPolDn(v string)`

SetPolDn sets PolDn field to given value.

### HasPolDn

`func (o *NiatelemetryPodTimeServerPolicies) HasPolDn() bool`

HasPolDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPodTimeServerPolicies) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPodTimeServerPolicies) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPodTimeServerPolicies) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPodTimeServerPolicies) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPodTimeServerPolicies) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPodTimeServerPolicies) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPodTimeServerPolicies) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPodTimeServerPolicies) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPodTimeServerPolicies) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPodTimeServerPolicies) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPodTimeServerPolicies) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPodTimeServerPolicies) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetTimeServerEnableSt

`func (o *NiatelemetryPodTimeServerPolicies) GetTimeServerEnableSt() string`

GetTimeServerEnableSt returns the TimeServerEnableSt field if non-nil, zero value otherwise.

### GetTimeServerEnableStOk

`func (o *NiatelemetryPodTimeServerPolicies) GetTimeServerEnableStOk() (*string, bool)`

GetTimeServerEnableStOk returns a tuple with the TimeServerEnableSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeServerEnableSt

`func (o *NiatelemetryPodTimeServerPolicies) SetTimeServerEnableSt(v string)`

SetTimeServerEnableSt sets TimeServerEnableSt field to given value.

### HasTimeServerEnableSt

`func (o *NiatelemetryPodTimeServerPolicies) HasTimeServerEnableSt() bool`

HasTimeServerEnableSt returns a boolean if a field has been set.

### GetTimeServers

`func (o *NiatelemetryPodTimeServerPolicies) GetTimeServers() string`

GetTimeServers returns the TimeServers field if non-nil, zero value otherwise.

### GetTimeServersOk

`func (o *NiatelemetryPodTimeServerPolicies) GetTimeServersOk() (*string, bool)`

GetTimeServersOk returns a tuple with the TimeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeServers

`func (o *NiatelemetryPodTimeServerPolicies) SetTimeServers(v string)`

SetTimeServers sets TimeServers field to given value.

### HasTimeServers

`func (o *NiatelemetryPodTimeServerPolicies) HasTimeServers() bool`

HasTimeServers returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPodTimeServerPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPodTimeServerPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPodTimeServerPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPodTimeServerPolicies) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryPodTimeServerPolicies) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryPodTimeServerPolicies) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


