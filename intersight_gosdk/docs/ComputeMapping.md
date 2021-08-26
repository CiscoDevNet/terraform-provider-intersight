# ComputeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.Mapping"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.Mapping"]
**FileLocation** | Pointer to **string** | Remote image location from where the image is uploaded to server. | [optional] 
**Identifier** | Pointer to **string** | The identity assigned to this Virtual Media Image by server. | [optional] [readonly] 
**ImageName** | Pointer to **string** | Image name of uploaded Virtual Media Image. | [optional] [readonly] 
**MediaTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name of Virtual Media mapping assigne by server. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vmedia** | Pointer to [**ComputeVmediaRelationship**](ComputeVmediaRelationship.md) |  | [optional] 

## Methods

### NewComputeMapping

`func NewComputeMapping(classId string, objectType string, ) *ComputeMapping`

NewComputeMapping instantiates a new ComputeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeMappingWithDefaults

`func NewComputeMappingWithDefaults() *ComputeMapping`

NewComputeMappingWithDefaults instantiates a new ComputeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeMapping) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeMapping) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeMapping) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeMapping) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeMapping) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeMapping) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileLocation

`func (o *ComputeMapping) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *ComputeMapping) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *ComputeMapping) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *ComputeMapping) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetIdentifier

`func (o *ComputeMapping) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ComputeMapping) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ComputeMapping) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ComputeMapping) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetImageName

`func (o *ComputeMapping) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ComputeMapping) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ComputeMapping) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ComputeMapping) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetMediaTypes

`func (o *ComputeMapping) GetMediaTypes() []string`

GetMediaTypes returns the MediaTypes field if non-nil, zero value otherwise.

### GetMediaTypesOk

`func (o *ComputeMapping) GetMediaTypesOk() (*[]string, bool)`

GetMediaTypesOk returns a tuple with the MediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypes

`func (o *ComputeMapping) SetMediaTypes(v []string)`

SetMediaTypes sets MediaTypes field to given value.

### HasMediaTypes

`func (o *ComputeMapping) HasMediaTypes() bool`

HasMediaTypes returns a boolean if a field has been set.

### SetMediaTypesNil

`func (o *ComputeMapping) SetMediaTypesNil(b bool)`

 SetMediaTypesNil sets the value for MediaTypes to be an explicit nil

### UnsetMediaTypes
`func (o *ComputeMapping) UnsetMediaTypes()`

UnsetMediaTypes ensures that no value is present for MediaTypes, not even an explicit nil
### GetName

`func (o *ComputeMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeMapping) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeMapping) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ComputeMapping) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ComputeMapping) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ComputeMapping) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ComputeMapping) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ComputeMapping) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ComputeMapping) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ComputeMapping) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ComputeMapping) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetVmedia

`func (o *ComputeMapping) GetVmedia() ComputeVmediaRelationship`

GetVmedia returns the Vmedia field if non-nil, zero value otherwise.

### GetVmediaOk

`func (o *ComputeMapping) GetVmediaOk() (*ComputeVmediaRelationship, bool)`

GetVmediaOk returns a tuple with the Vmedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmedia

`func (o *ComputeMapping) SetVmedia(v ComputeVmediaRelationship)`

SetVmedia sets Vmedia field to given value.

### HasVmedia

`func (o *ComputeMapping) HasVmedia() bool`

HasVmedia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


