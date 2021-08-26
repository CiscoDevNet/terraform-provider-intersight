# FabricVsanAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.Vsan"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.Vsan"]
**DefaultZoning** | Pointer to **string** | Enables or Disables the default zoning state. * &#x60;Enabled&#x60; - Admin configured Enabled State. * &#x60;Disabled&#x60; - Admin configured Disabled State. | [optional] [default to "Enabled"]
**FcZoneSharingMode** | Pointer to **string** | Logical grouping mode for fc ports. | [optional] 
**FcoeVlan** | Pointer to **int64** | FCOE Vlan associated to the VSAN configuration. | [optional] 
**Name** | Pointer to **string** | User given name for the VSAN configuration. | [optional] 
**VsanId** | Pointer to **int64** | Virtual San Identifier in the switch. | [optional] 
**FcNetworkPolicy** | Pointer to [**FabricFcNetworkPolicyRelationship**](FabricFcNetworkPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricVsanAllOf

`func NewFabricVsanAllOf(classId string, objectType string, ) *FabricVsanAllOf`

NewFabricVsanAllOf instantiates a new FabricVsanAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVsanAllOfWithDefaults

`func NewFabricVsanAllOfWithDefaults() *FabricVsanAllOf`

NewFabricVsanAllOfWithDefaults instantiates a new FabricVsanAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVsanAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVsanAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVsanAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVsanAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVsanAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVsanAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultZoning

`func (o *FabricVsanAllOf) GetDefaultZoning() string`

GetDefaultZoning returns the DefaultZoning field if non-nil, zero value otherwise.

### GetDefaultZoningOk

`func (o *FabricVsanAllOf) GetDefaultZoningOk() (*string, bool)`

GetDefaultZoningOk returns a tuple with the DefaultZoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZoning

`func (o *FabricVsanAllOf) SetDefaultZoning(v string)`

SetDefaultZoning sets DefaultZoning field to given value.

### HasDefaultZoning

`func (o *FabricVsanAllOf) HasDefaultZoning() bool`

HasDefaultZoning returns a boolean if a field has been set.

### GetFcZoneSharingMode

`func (o *FabricVsanAllOf) GetFcZoneSharingMode() string`

GetFcZoneSharingMode returns the FcZoneSharingMode field if non-nil, zero value otherwise.

### GetFcZoneSharingModeOk

`func (o *FabricVsanAllOf) GetFcZoneSharingModeOk() (*string, bool)`

GetFcZoneSharingModeOk returns a tuple with the FcZoneSharingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcZoneSharingMode

`func (o *FabricVsanAllOf) SetFcZoneSharingMode(v string)`

SetFcZoneSharingMode sets FcZoneSharingMode field to given value.

### HasFcZoneSharingMode

`func (o *FabricVsanAllOf) HasFcZoneSharingMode() bool`

HasFcZoneSharingMode returns a boolean if a field has been set.

### GetFcoeVlan

`func (o *FabricVsanAllOf) GetFcoeVlan() int64`

GetFcoeVlan returns the FcoeVlan field if non-nil, zero value otherwise.

### GetFcoeVlanOk

`func (o *FabricVsanAllOf) GetFcoeVlanOk() (*int64, bool)`

GetFcoeVlanOk returns a tuple with the FcoeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeVlan

`func (o *FabricVsanAllOf) SetFcoeVlan(v int64)`

SetFcoeVlan sets FcoeVlan field to given value.

### HasFcoeVlan

`func (o *FabricVsanAllOf) HasFcoeVlan() bool`

HasFcoeVlan returns a boolean if a field has been set.

### GetName

`func (o *FabricVsanAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVsanAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVsanAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVsanAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricVsanAllOf) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricVsanAllOf) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricVsanAllOf) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricVsanAllOf) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.

### GetFcNetworkPolicy

`func (o *FabricVsanAllOf) GetFcNetworkPolicy() FabricFcNetworkPolicyRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *FabricVsanAllOf) GetFcNetworkPolicyOk() (*FabricFcNetworkPolicyRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *FabricVsanAllOf) SetFcNetworkPolicy(v FabricFcNetworkPolicyRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *FabricVsanAllOf) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


