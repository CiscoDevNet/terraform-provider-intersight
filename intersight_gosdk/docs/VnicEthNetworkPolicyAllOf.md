# VnicEthNetworkPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthNetworkPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthNetworkPolicy"]
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "Standalone"]
**VlanSettings** | Pointer to [**NullableVnicVlanSettings**](VnicVlanSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicEthNetworkPolicyAllOf

`func NewVnicEthNetworkPolicyAllOf(classId string, objectType string, ) *VnicEthNetworkPolicyAllOf`

NewVnicEthNetworkPolicyAllOf instantiates a new VnicEthNetworkPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthNetworkPolicyAllOfWithDefaults

`func NewVnicEthNetworkPolicyAllOfWithDefaults() *VnicEthNetworkPolicyAllOf`

NewVnicEthNetworkPolicyAllOfWithDefaults instantiates a new VnicEthNetworkPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthNetworkPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthNetworkPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthNetworkPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthNetworkPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthNetworkPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthNetworkPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetPlatform

`func (o *VnicEthNetworkPolicyAllOf) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *VnicEthNetworkPolicyAllOf) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *VnicEthNetworkPolicyAllOf) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *VnicEthNetworkPolicyAllOf) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetVlanSettings

`func (o *VnicEthNetworkPolicyAllOf) GetVlanSettings() VnicVlanSettings`

GetVlanSettings returns the VlanSettings field if non-nil, zero value otherwise.

### GetVlanSettingsOk

`func (o *VnicEthNetworkPolicyAllOf) GetVlanSettingsOk() (*VnicVlanSettings, bool)`

GetVlanSettingsOk returns a tuple with the VlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSettings

`func (o *VnicEthNetworkPolicyAllOf) SetVlanSettings(v VnicVlanSettings)`

SetVlanSettings sets VlanSettings field to given value.

### HasVlanSettings

`func (o *VnicEthNetworkPolicyAllOf) HasVlanSettings() bool`

HasVlanSettings returns a boolean if a field has been set.

### SetVlanSettingsNil

`func (o *VnicEthNetworkPolicyAllOf) SetVlanSettingsNil(b bool)`

 SetVlanSettingsNil sets the value for VlanSettings to be an explicit nil

### UnsetVlanSettings
`func (o *VnicEthNetworkPolicyAllOf) UnsetVlanSettings()`

UnsetVlanSettings ensures that no value is present for VlanSettings, not even an explicit nil
### GetOrganization

`func (o *VnicEthNetworkPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicEthNetworkPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicEthNetworkPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicEthNetworkPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


