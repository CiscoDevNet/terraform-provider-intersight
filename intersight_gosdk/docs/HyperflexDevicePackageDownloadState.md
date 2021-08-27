# HyperflexDevicePackageDownloadState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.DevicePackageDownloadState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.DevicePackageDownloadState"]
**Checksum** | Pointer to **string** | Checksum of HyperFlex health check Debian package installed on the HyperFlex Device. | [optional] 
**HxDeviceName** | Pointer to **string** | HyperFlex Device Name for which the package download state is tracked. | [optional] [readonly] 
**HxNodes** | Pointer to **[]string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp of the last health check Debian package installation on the HyperFlex Device. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexDevicePackageDownloadState

`func NewHyperflexDevicePackageDownloadState(classId string, objectType string, ) *HyperflexDevicePackageDownloadState`

NewHyperflexDevicePackageDownloadState instantiates a new HyperflexDevicePackageDownloadState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDevicePackageDownloadStateWithDefaults

`func NewHyperflexDevicePackageDownloadStateWithDefaults() *HyperflexDevicePackageDownloadState`

NewHyperflexDevicePackageDownloadStateWithDefaults instantiates a new HyperflexDevicePackageDownloadState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexDevicePackageDownloadState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexDevicePackageDownloadState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexDevicePackageDownloadState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexDevicePackageDownloadState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexDevicePackageDownloadState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexDevicePackageDownloadState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChecksum

`func (o *HyperflexDevicePackageDownloadState) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *HyperflexDevicePackageDownloadState) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *HyperflexDevicePackageDownloadState) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *HyperflexDevicePackageDownloadState) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetHxDeviceName

`func (o *HyperflexDevicePackageDownloadState) GetHxDeviceName() string`

GetHxDeviceName returns the HxDeviceName field if non-nil, zero value otherwise.

### GetHxDeviceNameOk

`func (o *HyperflexDevicePackageDownloadState) GetHxDeviceNameOk() (*string, bool)`

GetHxDeviceNameOk returns a tuple with the HxDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxDeviceName

`func (o *HyperflexDevicePackageDownloadState) SetHxDeviceName(v string)`

SetHxDeviceName sets HxDeviceName field to given value.

### HasHxDeviceName

`func (o *HyperflexDevicePackageDownloadState) HasHxDeviceName() bool`

HasHxDeviceName returns a boolean if a field has been set.

### GetHxNodes

`func (o *HyperflexDevicePackageDownloadState) GetHxNodes() []string`

GetHxNodes returns the HxNodes field if non-nil, zero value otherwise.

### GetHxNodesOk

`func (o *HyperflexDevicePackageDownloadState) GetHxNodesOk() (*[]string, bool)`

GetHxNodesOk returns a tuple with the HxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxNodes

`func (o *HyperflexDevicePackageDownloadState) SetHxNodes(v []string)`

SetHxNodes sets HxNodes field to given value.

### HasHxNodes

`func (o *HyperflexDevicePackageDownloadState) HasHxNodes() bool`

HasHxNodes returns a boolean if a field has been set.

### SetHxNodesNil

`func (o *HyperflexDevicePackageDownloadState) SetHxNodesNil(b bool)`

 SetHxNodesNil sets the value for HxNodes to be an explicit nil

### UnsetHxNodes
`func (o *HyperflexDevicePackageDownloadState) UnsetHxNodes()`

UnsetHxNodes ensures that no value is present for HxNodes, not even an explicit nil
### GetTimestamp

`func (o *HyperflexDevicePackageDownloadState) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HyperflexDevicePackageDownloadState) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HyperflexDevicePackageDownloadState) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HyperflexDevicePackageDownloadState) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexDevicePackageDownloadState) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexDevicePackageDownloadState) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexDevicePackageDownloadState) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexDevicePackageDownloadState) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


