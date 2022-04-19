# FabricSanPinGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SanPinGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SanPinGroup"]
**PinTargetInterfaceRole** | Pointer to [**FabricAbstractInterfaceRoleRelationship**](FabricAbstractInterfaceRoleRelationship.md) |  | [optional] 

## Methods

### NewFabricSanPinGroup

`func NewFabricSanPinGroup(classId string, objectType string, ) *FabricSanPinGroup`

NewFabricSanPinGroup instantiates a new FabricSanPinGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSanPinGroupWithDefaults

`func NewFabricSanPinGroupWithDefaults() *FabricSanPinGroup`

NewFabricSanPinGroupWithDefaults instantiates a new FabricSanPinGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSanPinGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSanPinGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSanPinGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSanPinGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSanPinGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSanPinGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPinTargetInterfaceRole

`func (o *FabricSanPinGroup) GetPinTargetInterfaceRole() FabricAbstractInterfaceRoleRelationship`

GetPinTargetInterfaceRole returns the PinTargetInterfaceRole field if non-nil, zero value otherwise.

### GetPinTargetInterfaceRoleOk

`func (o *FabricSanPinGroup) GetPinTargetInterfaceRoleOk() (*FabricAbstractInterfaceRoleRelationship, bool)`

GetPinTargetInterfaceRoleOk returns a tuple with the PinTargetInterfaceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinTargetInterfaceRole

`func (o *FabricSanPinGroup) SetPinTargetInterfaceRole(v FabricAbstractInterfaceRoleRelationship)`

SetPinTargetInterfaceRole sets PinTargetInterfaceRole field to given value.

### HasPinTargetInterfaceRole

`func (o *FabricSanPinGroup) HasPinTargetInterfaceRole() bool`

HasPinTargetInterfaceRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


