# CloudVolumeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.VolumeType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.VolumeType"]
**Name** | Pointer to **string** | The type of the volume.Types vary from cloud to cloud. | [optional] [readonly] 
**VolumeTypeId** | Pointer to **string** | Unique identity of the volume type, assigned by the cloud provider. | [optional] [readonly] 

## Methods

### NewCloudVolumeType

`func NewCloudVolumeType(classId string, objectType string, ) *CloudVolumeType`

NewCloudVolumeType instantiates a new CloudVolumeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeTypeWithDefaults

`func NewCloudVolumeTypeWithDefaults() *CloudVolumeType`

NewCloudVolumeTypeWithDefaults instantiates a new CloudVolumeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudVolumeType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudVolumeType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudVolumeType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudVolumeType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudVolumeType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudVolumeType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *CloudVolumeType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudVolumeType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudVolumeType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudVolumeType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVolumeTypeId

`func (o *CloudVolumeType) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *CloudVolumeType) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *CloudVolumeType) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.

### HasVolumeTypeId

`func (o *CloudVolumeType) HasVolumeTypeId() bool`

HasVolumeTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


