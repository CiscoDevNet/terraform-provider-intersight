# FabricFcZoneMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcZoneMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcZoneMember"]
**Name** | Pointer to **string** | Name given to the target member. | [optional] 
**SwitchId** | Pointer to **string** | Unique identifier for the Fabric object. * &#x60;A&#x60; - Switch Identifier of Fabric Interconnect A. * &#x60;B&#x60; - Switch Identifier of Fabric Interconnect B. | [optional] [default to "A"]
**VsanId** | Pointer to **int64** | VSAN with scope defined as Storage in the VSAN policy. | [optional] 
**Wwpn** | Pointer to **string** | WWPN that is a member of the FC zone. | [optional] 

## Methods

### NewFabricFcZoneMember

`func NewFabricFcZoneMember(classId string, objectType string, ) *FabricFcZoneMember`

NewFabricFcZoneMember instantiates a new FabricFcZoneMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcZoneMemberWithDefaults

`func NewFabricFcZoneMemberWithDefaults() *FabricFcZoneMember`

NewFabricFcZoneMemberWithDefaults instantiates a new FabricFcZoneMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcZoneMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcZoneMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcZoneMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcZoneMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcZoneMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcZoneMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *FabricFcZoneMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricFcZoneMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricFcZoneMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricFcZoneMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSwitchId

`func (o *FabricFcZoneMember) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FabricFcZoneMember) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FabricFcZoneMember) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FabricFcZoneMember) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricFcZoneMember) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricFcZoneMember) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricFcZoneMember) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricFcZoneMember) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.

### GetWwpn

`func (o *FabricFcZoneMember) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *FabricFcZoneMember) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *FabricFcZoneMember) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *FabricFcZoneMember) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


