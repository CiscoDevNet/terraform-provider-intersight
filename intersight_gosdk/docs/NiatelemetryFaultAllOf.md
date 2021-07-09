# NiatelemetryFaultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Fault"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Fault"]
**Cause** | Pointer to **string** | Cause of the fault present. | [optional] 
**Code** | Pointer to **string** | Code of the fault present. | [optional] 
**CreatedTime** | Pointer to **string** | Created time of the fault present. | [optional] 
**Description** | Pointer to **string** | Description of the fault present. | [optional] 
**Dn** | Pointer to **string** | Dn value for the fault present. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Severity** | Pointer to **string** | Severity of the fault present. | [optional] 
**SiteName** | Pointer to **string** | The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters/sites. | [optional] 
**Type** | Pointer to **string** | Type of the fault present. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryFaultAllOf

`func NewNiatelemetryFaultAllOf(classId string, objectType string, ) *NiatelemetryFaultAllOf`

NewNiatelemetryFaultAllOf instantiates a new NiatelemetryFaultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryFaultAllOfWithDefaults

`func NewNiatelemetryFaultAllOfWithDefaults() *NiatelemetryFaultAllOf`

NewNiatelemetryFaultAllOfWithDefaults instantiates a new NiatelemetryFaultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryFaultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryFaultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryFaultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryFaultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryFaultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryFaultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCause

`func (o *NiatelemetryFaultAllOf) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *NiatelemetryFaultAllOf) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *NiatelemetryFaultAllOf) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *NiatelemetryFaultAllOf) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCode

`func (o *NiatelemetryFaultAllOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NiatelemetryFaultAllOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NiatelemetryFaultAllOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NiatelemetryFaultAllOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NiatelemetryFaultAllOf) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NiatelemetryFaultAllOf) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NiatelemetryFaultAllOf) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NiatelemetryFaultAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *NiatelemetryFaultAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiatelemetryFaultAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiatelemetryFaultAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiatelemetryFaultAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryFaultAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryFaultAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryFaultAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryFaultAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryFaultAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryFaultAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryFaultAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryFaultAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryFaultAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryFaultAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryFaultAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryFaultAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSeverity

`func (o *NiatelemetryFaultAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NiatelemetryFaultAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NiatelemetryFaultAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NiatelemetryFaultAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryFaultAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryFaultAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryFaultAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryFaultAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetType

`func (o *NiatelemetryFaultAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiatelemetryFaultAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiatelemetryFaultAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiatelemetryFaultAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryFaultAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryFaultAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryFaultAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryFaultAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


