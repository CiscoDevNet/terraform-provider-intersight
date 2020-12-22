# FabricFcoeUplinkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcoeUplinkRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcoeUplinkRole"]
**UdldAdminState** | Pointer to **string** | Admin configured state for UDLD for this port. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]

## Methods

### NewFabricFcoeUplinkRole

`func NewFabricFcoeUplinkRole(classId string, objectType string, ) *FabricFcoeUplinkRole`

NewFabricFcoeUplinkRole instantiates a new FabricFcoeUplinkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcoeUplinkRoleWithDefaults

`func NewFabricFcoeUplinkRoleWithDefaults() *FabricFcoeUplinkRole`

NewFabricFcoeUplinkRoleWithDefaults instantiates a new FabricFcoeUplinkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcoeUplinkRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcoeUplinkRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcoeUplinkRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcoeUplinkRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcoeUplinkRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcoeUplinkRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUdldAdminState

`func (o *FabricFcoeUplinkRole) GetUdldAdminState() string`

GetUdldAdminState returns the UdldAdminState field if non-nil, zero value otherwise.

### GetUdldAdminStateOk

`func (o *FabricFcoeUplinkRole) GetUdldAdminStateOk() (*string, bool)`

GetUdldAdminStateOk returns a tuple with the UdldAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdldAdminState

`func (o *FabricFcoeUplinkRole) SetUdldAdminState(v string)`

SetUdldAdminState sets UdldAdminState field to given value.

### HasUdldAdminState

`func (o *FabricFcoeUplinkRole) HasUdldAdminState() bool`

HasUdldAdminState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


