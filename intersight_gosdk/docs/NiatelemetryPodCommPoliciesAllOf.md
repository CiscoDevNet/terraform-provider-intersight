# NiatelemetryPodCommPoliciesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.PodCommPolicies"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.PodCommPolicies"]
**CommHttpAdminSt** | Pointer to **string** | Comm Http Admin State of the Comm Pol in APIC. | [optional] 
**CommHttpsAdminSt** | Pointer to **string** | Comm Https Admin State of the Comm Pol in APIC. | [optional] 
**CommSshAdminSt** | Pointer to **string** | Comm Ssh Admin State of the Comm Pol in APIC. | [optional] 
**CommTelnetAdminSt** | Pointer to **string** | Comm Telnet Admin State of the Comm Pol in APIC. | [optional] 
**PolDn** | Pointer to **string** | Dn of the Comm Pol in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryPodCommPoliciesAllOf

`func NewNiatelemetryPodCommPoliciesAllOf(classId string, objectType string, ) *NiatelemetryPodCommPoliciesAllOf`

NewNiatelemetryPodCommPoliciesAllOf instantiates a new NiatelemetryPodCommPoliciesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPodCommPoliciesAllOfWithDefaults

`func NewNiatelemetryPodCommPoliciesAllOfWithDefaults() *NiatelemetryPodCommPoliciesAllOf`

NewNiatelemetryPodCommPoliciesAllOfWithDefaults instantiates a new NiatelemetryPodCommPoliciesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPodCommPoliciesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPodCommPoliciesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPodCommPoliciesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPodCommPoliciesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommHttpAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpAdminSt() string`

GetCommHttpAdminSt returns the CommHttpAdminSt field if non-nil, zero value otherwise.

### GetCommHttpAdminStOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpAdminStOk() (*string, bool)`

GetCommHttpAdminStOk returns a tuple with the CommHttpAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommHttpAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) SetCommHttpAdminSt(v string)`

SetCommHttpAdminSt sets CommHttpAdminSt field to given value.

### HasCommHttpAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) HasCommHttpAdminSt() bool`

HasCommHttpAdminSt returns a boolean if a field has been set.

### GetCommHttpsAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpsAdminSt() string`

GetCommHttpsAdminSt returns the CommHttpsAdminSt field if non-nil, zero value otherwise.

### GetCommHttpsAdminStOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpsAdminStOk() (*string, bool)`

GetCommHttpsAdminStOk returns a tuple with the CommHttpsAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommHttpsAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) SetCommHttpsAdminSt(v string)`

SetCommHttpsAdminSt sets CommHttpsAdminSt field to given value.

### HasCommHttpsAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) HasCommHttpsAdminSt() bool`

HasCommHttpsAdminSt returns a boolean if a field has been set.

### GetCommSshAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommSshAdminSt() string`

GetCommSshAdminSt returns the CommSshAdminSt field if non-nil, zero value otherwise.

### GetCommSshAdminStOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommSshAdminStOk() (*string, bool)`

GetCommSshAdminStOk returns a tuple with the CommSshAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommSshAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) SetCommSshAdminSt(v string)`

SetCommSshAdminSt sets CommSshAdminSt field to given value.

### HasCommSshAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) HasCommSshAdminSt() bool`

HasCommSshAdminSt returns a boolean if a field has been set.

### GetCommTelnetAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommTelnetAdminSt() string`

GetCommTelnetAdminSt returns the CommTelnetAdminSt field if non-nil, zero value otherwise.

### GetCommTelnetAdminStOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetCommTelnetAdminStOk() (*string, bool)`

GetCommTelnetAdminStOk returns a tuple with the CommTelnetAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommTelnetAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) SetCommTelnetAdminSt(v string)`

SetCommTelnetAdminSt sets CommTelnetAdminSt field to given value.

### HasCommTelnetAdminSt

`func (o *NiatelemetryPodCommPoliciesAllOf) HasCommTelnetAdminSt() bool`

HasCommTelnetAdminSt returns a boolean if a field has been set.

### GetPolDn

`func (o *NiatelemetryPodCommPoliciesAllOf) GetPolDn() string`

GetPolDn returns the PolDn field if non-nil, zero value otherwise.

### GetPolDnOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetPolDnOk() (*string, bool)`

GetPolDnOk returns a tuple with the PolDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDn

`func (o *NiatelemetryPodCommPoliciesAllOf) SetPolDn(v string)`

SetPolDn sets PolDn field to given value.

### HasPolDn

`func (o *NiatelemetryPodCommPoliciesAllOf) HasPolDn() bool`

HasPolDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPodCommPoliciesAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPodCommPoliciesAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPodCommPoliciesAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPodCommPoliciesAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPodCommPoliciesAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPodCommPoliciesAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPodCommPoliciesAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPodCommPoliciesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPodCommPoliciesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPodCommPoliciesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPodCommPoliciesAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


