# NiatelemetryPodCommPolicies

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

### NewNiatelemetryPodCommPolicies

`func NewNiatelemetryPodCommPolicies(classId string, objectType string, ) *NiatelemetryPodCommPolicies`

NewNiatelemetryPodCommPolicies instantiates a new NiatelemetryPodCommPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryPodCommPoliciesWithDefaults

`func NewNiatelemetryPodCommPoliciesWithDefaults() *NiatelemetryPodCommPolicies`

NewNiatelemetryPodCommPoliciesWithDefaults instantiates a new NiatelemetryPodCommPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryPodCommPolicies) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryPodCommPolicies) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryPodCommPolicies) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryPodCommPolicies) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryPodCommPolicies) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryPodCommPolicies) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommHttpAdminSt

`func (o *NiatelemetryPodCommPolicies) GetCommHttpAdminSt() string`

GetCommHttpAdminSt returns the CommHttpAdminSt field if non-nil, zero value otherwise.

### GetCommHttpAdminStOk

`func (o *NiatelemetryPodCommPolicies) GetCommHttpAdminStOk() (*string, bool)`

GetCommHttpAdminStOk returns a tuple with the CommHttpAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommHttpAdminSt

`func (o *NiatelemetryPodCommPolicies) SetCommHttpAdminSt(v string)`

SetCommHttpAdminSt sets CommHttpAdminSt field to given value.

### HasCommHttpAdminSt

`func (o *NiatelemetryPodCommPolicies) HasCommHttpAdminSt() bool`

HasCommHttpAdminSt returns a boolean if a field has been set.

### GetCommHttpsAdminSt

`func (o *NiatelemetryPodCommPolicies) GetCommHttpsAdminSt() string`

GetCommHttpsAdminSt returns the CommHttpsAdminSt field if non-nil, zero value otherwise.

### GetCommHttpsAdminStOk

`func (o *NiatelemetryPodCommPolicies) GetCommHttpsAdminStOk() (*string, bool)`

GetCommHttpsAdminStOk returns a tuple with the CommHttpsAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommHttpsAdminSt

`func (o *NiatelemetryPodCommPolicies) SetCommHttpsAdminSt(v string)`

SetCommHttpsAdminSt sets CommHttpsAdminSt field to given value.

### HasCommHttpsAdminSt

`func (o *NiatelemetryPodCommPolicies) HasCommHttpsAdminSt() bool`

HasCommHttpsAdminSt returns a boolean if a field has been set.

### GetCommSshAdminSt

`func (o *NiatelemetryPodCommPolicies) GetCommSshAdminSt() string`

GetCommSshAdminSt returns the CommSshAdminSt field if non-nil, zero value otherwise.

### GetCommSshAdminStOk

`func (o *NiatelemetryPodCommPolicies) GetCommSshAdminStOk() (*string, bool)`

GetCommSshAdminStOk returns a tuple with the CommSshAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommSshAdminSt

`func (o *NiatelemetryPodCommPolicies) SetCommSshAdminSt(v string)`

SetCommSshAdminSt sets CommSshAdminSt field to given value.

### HasCommSshAdminSt

`func (o *NiatelemetryPodCommPolicies) HasCommSshAdminSt() bool`

HasCommSshAdminSt returns a boolean if a field has been set.

### GetCommTelnetAdminSt

`func (o *NiatelemetryPodCommPolicies) GetCommTelnetAdminSt() string`

GetCommTelnetAdminSt returns the CommTelnetAdminSt field if non-nil, zero value otherwise.

### GetCommTelnetAdminStOk

`func (o *NiatelemetryPodCommPolicies) GetCommTelnetAdminStOk() (*string, bool)`

GetCommTelnetAdminStOk returns a tuple with the CommTelnetAdminSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommTelnetAdminSt

`func (o *NiatelemetryPodCommPolicies) SetCommTelnetAdminSt(v string)`

SetCommTelnetAdminSt sets CommTelnetAdminSt field to given value.

### HasCommTelnetAdminSt

`func (o *NiatelemetryPodCommPolicies) HasCommTelnetAdminSt() bool`

HasCommTelnetAdminSt returns a boolean if a field has been set.

### GetPolDn

`func (o *NiatelemetryPodCommPolicies) GetPolDn() string`

GetPolDn returns the PolDn field if non-nil, zero value otherwise.

### GetPolDnOk

`func (o *NiatelemetryPodCommPolicies) GetPolDnOk() (*string, bool)`

GetPolDnOk returns a tuple with the PolDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolDn

`func (o *NiatelemetryPodCommPolicies) SetPolDn(v string)`

SetPolDn sets PolDn field to given value.

### HasPolDn

`func (o *NiatelemetryPodCommPolicies) HasPolDn() bool`

HasPolDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryPodCommPolicies) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryPodCommPolicies) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryPodCommPolicies) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryPodCommPolicies) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryPodCommPolicies) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryPodCommPolicies) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryPodCommPolicies) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryPodCommPolicies) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryPodCommPolicies) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryPodCommPolicies) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryPodCommPolicies) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryPodCommPolicies) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryPodCommPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryPodCommPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryPodCommPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryPodCommPolicies) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


