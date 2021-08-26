# NiatelemetryNiaLicenseStateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaLicenseState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaLicenseState"]
**FeatureActivated** | Pointer to **string** | Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check. | [optional] 
**LicenseActivated** | Pointer to **string** | Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device. | [optional] 
**PidType** | Pointer to **string** | PID of device being inventoried. This determines the hardware model type of the device. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**SiteName** | Pointer to **string** | Name of fabric domain of the controller. | [optional] 
**Device** | Pointer to [**NiatelemetryNiaInventoryRelationship**](NiatelemetryNiaInventoryRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaLicenseStateAllOf

`func NewNiatelemetryNiaLicenseStateAllOf(classId string, objectType string, ) *NiatelemetryNiaLicenseStateAllOf`

NewNiatelemetryNiaLicenseStateAllOf instantiates a new NiatelemetryNiaLicenseStateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaLicenseStateAllOfWithDefaults

`func NewNiatelemetryNiaLicenseStateAllOfWithDefaults() *NiatelemetryNiaLicenseStateAllOf`

NewNiatelemetryNiaLicenseStateAllOfWithDefaults instantiates a new NiatelemetryNiaLicenseStateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaLicenseStateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaLicenseStateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaLicenseStateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaLicenseStateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatureActivated

`func (o *NiatelemetryNiaLicenseStateAllOf) GetFeatureActivated() string`

GetFeatureActivated returns the FeatureActivated field if non-nil, zero value otherwise.

### GetFeatureActivatedOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetFeatureActivatedOk() (*string, bool)`

GetFeatureActivatedOk returns a tuple with the FeatureActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureActivated

`func (o *NiatelemetryNiaLicenseStateAllOf) SetFeatureActivated(v string)`

SetFeatureActivated sets FeatureActivated field to given value.

### HasFeatureActivated

`func (o *NiatelemetryNiaLicenseStateAllOf) HasFeatureActivated() bool`

HasFeatureActivated returns a boolean if a field has been set.

### GetLicenseActivated

`func (o *NiatelemetryNiaLicenseStateAllOf) GetLicenseActivated() string`

GetLicenseActivated returns the LicenseActivated field if non-nil, zero value otherwise.

### GetLicenseActivatedOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetLicenseActivatedOk() (*string, bool)`

GetLicenseActivatedOk returns a tuple with the LicenseActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseActivated

`func (o *NiatelemetryNiaLicenseStateAllOf) SetLicenseActivated(v string)`

SetLicenseActivated sets LicenseActivated field to given value.

### HasLicenseActivated

`func (o *NiatelemetryNiaLicenseStateAllOf) HasLicenseActivated() bool`

HasLicenseActivated returns a boolean if a field has been set.

### GetPidType

`func (o *NiatelemetryNiaLicenseStateAllOf) GetPidType() string`

GetPidType returns the PidType field if non-nil, zero value otherwise.

### GetPidTypeOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetPidTypeOk() (*string, bool)`

GetPidTypeOk returns a tuple with the PidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPidType

`func (o *NiatelemetryNiaLicenseStateAllOf) SetPidType(v string)`

SetPidType sets PidType field to given value.

### HasPidType

`func (o *NiatelemetryNiaLicenseStateAllOf) HasPidType() bool`

HasPidType returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNiaLicenseStateAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNiaLicenseStateAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryNiaLicenseStateAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryNiaLicenseStateAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaLicenseStateAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaLicenseStateAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaLicenseStateAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaLicenseStateAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaLicenseStateAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaLicenseStateAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetDevice

`func (o *NiatelemetryNiaLicenseStateAllOf) GetDevice() NiatelemetryNiaInventoryRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NiatelemetryNiaLicenseStateAllOf) GetDeviceOk() (*NiatelemetryNiaInventoryRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NiatelemetryNiaLicenseStateAllOf) SetDevice(v NiatelemetryNiaInventoryRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *NiatelemetryNiaLicenseStateAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


