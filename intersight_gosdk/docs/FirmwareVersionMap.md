# FirmwareVersionMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.VersionMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.VersionMap"]
**BundleVersion** | Pointer to **string** | Bundle version. Usually the first released bundle containing the specific device firmware version. | [optional] [readonly] 
**DeviceFirmwareVersion** | Pointer to **string** | Bundled device firmware version. | [optional] [readonly] 

## Methods

### NewFirmwareVersionMap

`func NewFirmwareVersionMap(classId string, objectType string, ) *FirmwareVersionMap`

NewFirmwareVersionMap instantiates a new FirmwareVersionMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareVersionMapWithDefaults

`func NewFirmwareVersionMapWithDefaults() *FirmwareVersionMap`

NewFirmwareVersionMapWithDefaults instantiates a new FirmwareVersionMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareVersionMap) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareVersionMap) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareVersionMap) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareVersionMap) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareVersionMap) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareVersionMap) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBundleVersion

`func (o *FirmwareVersionMap) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *FirmwareVersionMap) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *FirmwareVersionMap) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.

### HasBundleVersion

`func (o *FirmwareVersionMap) HasBundleVersion() bool`

HasBundleVersion returns a boolean if a field has been set.

### GetDeviceFirmwareVersion

`func (o *FirmwareVersionMap) GetDeviceFirmwareVersion() string`

GetDeviceFirmwareVersion returns the DeviceFirmwareVersion field if non-nil, zero value otherwise.

### GetDeviceFirmwareVersionOk

`func (o *FirmwareVersionMap) GetDeviceFirmwareVersionOk() (*string, bool)`

GetDeviceFirmwareVersionOk returns a tuple with the DeviceFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFirmwareVersion

`func (o *FirmwareVersionMap) SetDeviceFirmwareVersion(v string)`

SetDeviceFirmwareVersion sets DeviceFirmwareVersion field to given value.

### HasDeviceFirmwareVersion

`func (o *FirmwareVersionMap) HasDeviceFirmwareVersion() bool`

HasDeviceFirmwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


