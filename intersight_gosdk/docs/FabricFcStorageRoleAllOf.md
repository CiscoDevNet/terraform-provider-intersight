# FabricFcStorageRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcStorageRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcStorageRole"]
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. * &#x60;Auto&#x60; - Admin configurable speed AUTO ( default ). * &#x60;8Gbps&#x60; - Admin configurable speed 8Gbps. * &#x60;16Gbps&#x60; - Admin configurable speed 16Gbps. * &#x60;32Gbps&#x60; - Admin configurable speed 32Gbps. | [optional] [default to "Auto"]
**VsanId** | Pointer to **int64** | Virtual San Identifier associated to the FC port. | [optional] 

## Methods

### NewFabricFcStorageRoleAllOf

`func NewFabricFcStorageRoleAllOf(classId string, objectType string, ) *FabricFcStorageRoleAllOf`

NewFabricFcStorageRoleAllOf instantiates a new FabricFcStorageRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcStorageRoleAllOfWithDefaults

`func NewFabricFcStorageRoleAllOfWithDefaults() *FabricFcStorageRoleAllOf`

NewFabricFcStorageRoleAllOfWithDefaults instantiates a new FabricFcStorageRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcStorageRoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcStorageRoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcStorageRoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcStorageRoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcStorageRoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcStorageRoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FabricFcStorageRoleAllOf) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricFcStorageRoleAllOf) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricFcStorageRoleAllOf) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricFcStorageRoleAllOf) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricFcStorageRoleAllOf) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricFcStorageRoleAllOf) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricFcStorageRoleAllOf) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricFcStorageRoleAllOf) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


