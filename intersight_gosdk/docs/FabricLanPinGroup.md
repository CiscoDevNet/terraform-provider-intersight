# FabricLanPinGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.LanPinGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.LanPinGroup"]
**PinTargetInterfaceRole** | Pointer to [**FabricAbstractInterfaceRoleRelationship**](FabricAbstractInterfaceRoleRelationship.md) |  | [optional] 

## Methods

### NewFabricLanPinGroup

`func NewFabricLanPinGroup(classId string, objectType string, ) *FabricLanPinGroup`

NewFabricLanPinGroup instantiates a new FabricLanPinGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricLanPinGroupWithDefaults

`func NewFabricLanPinGroupWithDefaults() *FabricLanPinGroup`

NewFabricLanPinGroupWithDefaults instantiates a new FabricLanPinGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricLanPinGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricLanPinGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricLanPinGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricLanPinGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricLanPinGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricLanPinGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPinTargetInterfaceRole

`func (o *FabricLanPinGroup) GetPinTargetInterfaceRole() FabricAbstractInterfaceRoleRelationship`

GetPinTargetInterfaceRole returns the PinTargetInterfaceRole field if non-nil, zero value otherwise.

### GetPinTargetInterfaceRoleOk

`func (o *FabricLanPinGroup) GetPinTargetInterfaceRoleOk() (*FabricAbstractInterfaceRoleRelationship, bool)`

GetPinTargetInterfaceRoleOk returns a tuple with the PinTargetInterfaceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinTargetInterfaceRole

`func (o *FabricLanPinGroup) SetPinTargetInterfaceRole(v FabricAbstractInterfaceRoleRelationship)`

SetPinTargetInterfaceRole sets PinTargetInterfaceRole field to given value.

### HasPinTargetInterfaceRole

`func (o *FabricLanPinGroup) HasPinTargetInterfaceRole() bool`

HasPinTargetInterfaceRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


