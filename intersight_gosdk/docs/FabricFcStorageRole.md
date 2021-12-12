# FabricFcStorageRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcStorageRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcStorageRole"]
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. * &#x60;Auto&#x60; - Admin configurable speed AUTO ( default ). * &#x60;8Gbps&#x60; - Admin configurable speed 8Gbps. * &#x60;16Gbps&#x60; - Admin configurable speed 16Gbps. * &#x60;32Gbps&#x60; - Admin configurable speed 32Gbps. | [optional] [default to "Auto"]
**VsanId** | Pointer to **int64** | Virtual San Identifier associated to the FC port. | [optional] 

## Methods

### NewFabricFcStorageRole

`func NewFabricFcStorageRole(classId string, objectType string, ) *FabricFcStorageRole`

NewFabricFcStorageRole instantiates a new FabricFcStorageRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcStorageRoleWithDefaults

`func NewFabricFcStorageRoleWithDefaults() *FabricFcStorageRole`

NewFabricFcStorageRoleWithDefaults instantiates a new FabricFcStorageRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcStorageRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcStorageRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcStorageRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcStorageRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcStorageRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcStorageRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FabricFcStorageRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricFcStorageRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricFcStorageRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricFcStorageRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricFcStorageRole) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricFcStorageRole) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricFcStorageRole) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricFcStorageRole) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


