# NiatelemetryNiaLicenseState

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

### NewNiatelemetryNiaLicenseState

`func NewNiatelemetryNiaLicenseState(classId string, objectType string, ) *NiatelemetryNiaLicenseState`

NewNiatelemetryNiaLicenseState instantiates a new NiatelemetryNiaLicenseState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaLicenseStateWithDefaults

`func NewNiatelemetryNiaLicenseStateWithDefaults() *NiatelemetryNiaLicenseState`

NewNiatelemetryNiaLicenseStateWithDefaults instantiates a new NiatelemetryNiaLicenseState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaLicenseState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaLicenseState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaLicenseState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaLicenseState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaLicenseState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaLicenseState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatureActivated

`func (o *NiatelemetryNiaLicenseState) GetFeatureActivated() string`

GetFeatureActivated returns the FeatureActivated field if non-nil, zero value otherwise.

### GetFeatureActivatedOk

`func (o *NiatelemetryNiaLicenseState) GetFeatureActivatedOk() (*string, bool)`

GetFeatureActivatedOk returns a tuple with the FeatureActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureActivated

`func (o *NiatelemetryNiaLicenseState) SetFeatureActivated(v string)`

SetFeatureActivated sets FeatureActivated field to given value.

### HasFeatureActivated

`func (o *NiatelemetryNiaLicenseState) HasFeatureActivated() bool`

HasFeatureActivated returns a boolean if a field has been set.

### GetLicenseActivated

`func (o *NiatelemetryNiaLicenseState) GetLicenseActivated() string`

GetLicenseActivated returns the LicenseActivated field if non-nil, zero value otherwise.

### GetLicenseActivatedOk

`func (o *NiatelemetryNiaLicenseState) GetLicenseActivatedOk() (*string, bool)`

GetLicenseActivatedOk returns a tuple with the LicenseActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseActivated

`func (o *NiatelemetryNiaLicenseState) SetLicenseActivated(v string)`

SetLicenseActivated sets LicenseActivated field to given value.

### HasLicenseActivated

`func (o *NiatelemetryNiaLicenseState) HasLicenseActivated() bool`

HasLicenseActivated returns a boolean if a field has been set.

### GetPidType

`func (o *NiatelemetryNiaLicenseState) GetPidType() string`

GetPidType returns the PidType field if non-nil, zero value otherwise.

### GetPidTypeOk

`func (o *NiatelemetryNiaLicenseState) GetPidTypeOk() (*string, bool)`

GetPidTypeOk returns a tuple with the PidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPidType

`func (o *NiatelemetryNiaLicenseState) SetPidType(v string)`

SetPidType sets PidType field to given value.

### HasPidType

`func (o *NiatelemetryNiaLicenseState) HasPidType() bool`

HasPidType returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryNiaLicenseState) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNiaLicenseState) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNiaLicenseState) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNiaLicenseState) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryNiaLicenseState) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryNiaLicenseState) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryNiaLicenseState) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryNiaLicenseState) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaLicenseState) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaLicenseState) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaLicenseState) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaLicenseState) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaLicenseState) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaLicenseState) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaLicenseState) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaLicenseState) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetDevice

`func (o *NiatelemetryNiaLicenseState) GetDevice() NiatelemetryNiaInventoryRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NiatelemetryNiaLicenseState) GetDeviceOk() (*NiatelemetryNiaInventoryRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NiatelemetryNiaLicenseState) SetDevice(v NiatelemetryNiaInventoryRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *NiatelemetryNiaLicenseState) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


