# VirtualizationCloudVmStorageConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.AwsVmStorageConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.AwsVmStorageConfiguration"]
**Volumes** | Pointer to [**[]VirtualizationVolumeInfo**](VirtualizationVolumeInfo.md) |  | [optional] 

## Methods

### NewVirtualizationCloudVmStorageConfigurationAllOf

`func NewVirtualizationCloudVmStorageConfigurationAllOf(classId string, objectType string, ) *VirtualizationCloudVmStorageConfigurationAllOf`

NewVirtualizationCloudVmStorageConfigurationAllOf instantiates a new VirtualizationCloudVmStorageConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationCloudVmStorageConfigurationAllOfWithDefaults

`func NewVirtualizationCloudVmStorageConfigurationAllOfWithDefaults() *VirtualizationCloudVmStorageConfigurationAllOf`

NewVirtualizationCloudVmStorageConfigurationAllOfWithDefaults instantiates a new VirtualizationCloudVmStorageConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVolumes

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetVolumes() []VirtualizationVolumeInfo`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) GetVolumesOk() (*[]VirtualizationVolumeInfo, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) SetVolumes(v []VirtualizationVolumeInfo)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *VirtualizationCloudVmStorageConfigurationAllOf) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *VirtualizationCloudVmStorageConfigurationAllOf) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


