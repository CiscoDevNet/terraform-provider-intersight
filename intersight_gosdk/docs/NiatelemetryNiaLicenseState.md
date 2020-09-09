# NiatelemetryNiaLicenseState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureActivated** | Pointer to **string** | Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check. | [optional] 
**LicenseActivated** | Pointer to **string** | Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device. | [optional] 
**PidType** | Pointer to **string** | PID of device being inventoried. This determines the hardware model type of the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**Device** | Pointer to [**NiatelemetryNiaInventoryRelationship**](niatelemetry.NiaInventory.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaLicenseState

`func NewNiatelemetryNiaLicenseState() *NiatelemetryNiaLicenseState`

NewNiatelemetryNiaLicenseState instantiates a new NiatelemetryNiaLicenseState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaLicenseStateWithDefaults

`func NewNiatelemetryNiaLicenseStateWithDefaults() *NiatelemetryNiaLicenseState`

NewNiatelemetryNiaLicenseStateWithDefaults instantiates a new NiatelemetryNiaLicenseState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


