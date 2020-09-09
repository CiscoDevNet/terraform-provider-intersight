# FabricVsan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultZoning** | Pointer to **string** | Enables or Disables the default zoning state. * &#x60;Enabled&#x60; - Admin configured Enabled State. * &#x60;Disabled&#x60; - Admin configured Disabled State. | [optional] [default to "Enabled"]
**FcZoneSharingMode** | Pointer to **string** | Logical grouping mode for fc ports. | [optional] 
**FcoeVlan** | Pointer to **int64** | FCOE Vlan associated to the VSAN configuration. | [optional] 
**Name** | Pointer to **string** | User given name for the VSAN configuration. | [optional] 
**VsanId** | Pointer to **int64** | Virtual San Identifier in the switch. | [optional] 
**FcNetworkPolicy** | Pointer to [**FabricFcNetworkPolicyRelationship**](fabric.FcNetworkPolicy.Relationship.md) |  | [optional] 

## Methods

### NewFabricVsan

`func NewFabricVsan() *FabricVsan`

NewFabricVsan instantiates a new FabricVsan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVsanWithDefaults

`func NewFabricVsanWithDefaults() *FabricVsan`

NewFabricVsanWithDefaults instantiates a new FabricVsan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultZoning

`func (o *FabricVsan) GetDefaultZoning() string`

GetDefaultZoning returns the DefaultZoning field if non-nil, zero value otherwise.

### GetDefaultZoningOk

`func (o *FabricVsan) GetDefaultZoningOk() (*string, bool)`

GetDefaultZoningOk returns a tuple with the DefaultZoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZoning

`func (o *FabricVsan) SetDefaultZoning(v string)`

SetDefaultZoning sets DefaultZoning field to given value.

### HasDefaultZoning

`func (o *FabricVsan) HasDefaultZoning() bool`

HasDefaultZoning returns a boolean if a field has been set.

### GetFcZoneSharingMode

`func (o *FabricVsan) GetFcZoneSharingMode() string`

GetFcZoneSharingMode returns the FcZoneSharingMode field if non-nil, zero value otherwise.

### GetFcZoneSharingModeOk

`func (o *FabricVsan) GetFcZoneSharingModeOk() (*string, bool)`

GetFcZoneSharingModeOk returns a tuple with the FcZoneSharingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcZoneSharingMode

`func (o *FabricVsan) SetFcZoneSharingMode(v string)`

SetFcZoneSharingMode sets FcZoneSharingMode field to given value.

### HasFcZoneSharingMode

`func (o *FabricVsan) HasFcZoneSharingMode() bool`

HasFcZoneSharingMode returns a boolean if a field has been set.

### GetFcoeVlan

`func (o *FabricVsan) GetFcoeVlan() int64`

GetFcoeVlan returns the FcoeVlan field if non-nil, zero value otherwise.

### GetFcoeVlanOk

`func (o *FabricVsan) GetFcoeVlanOk() (*int64, bool)`

GetFcoeVlanOk returns a tuple with the FcoeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeVlan

`func (o *FabricVsan) SetFcoeVlan(v int64)`

SetFcoeVlan sets FcoeVlan field to given value.

### HasFcoeVlan

`func (o *FabricVsan) HasFcoeVlan() bool`

HasFcoeVlan returns a boolean if a field has been set.

### GetName

`func (o *FabricVsan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVsan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVsan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVsan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricVsan) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricVsan) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricVsan) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricVsan) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.

### GetFcNetworkPolicy

`func (o *FabricVsan) GetFcNetworkPolicy() FabricFcNetworkPolicyRelationship`

GetFcNetworkPolicy returns the FcNetworkPolicy field if non-nil, zero value otherwise.

### GetFcNetworkPolicyOk

`func (o *FabricVsan) GetFcNetworkPolicyOk() (*FabricFcNetworkPolicyRelationship, bool)`

GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNetworkPolicy

`func (o *FabricVsan) SetFcNetworkPolicy(v FabricFcNetworkPolicyRelationship)`

SetFcNetworkPolicy sets FcNetworkPolicy field to given value.

### HasFcNetworkPolicy

`func (o *FabricVsan) HasFcNetworkPolicy() bool`

HasFcNetworkPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


