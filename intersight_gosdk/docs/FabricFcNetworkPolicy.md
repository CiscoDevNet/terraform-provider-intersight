# FabricFcNetworkPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcNetworkPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcNetworkPolicy"]
**EnableTrunking** | Pointer to **bool** | Enable or Disable Trunking on all of configured FC uplink ports. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricFcNetworkPolicy

`func NewFabricFcNetworkPolicy(classId string, objectType string, ) *FabricFcNetworkPolicy`

NewFabricFcNetworkPolicy instantiates a new FabricFcNetworkPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcNetworkPolicyWithDefaults

`func NewFabricFcNetworkPolicyWithDefaults() *FabricFcNetworkPolicy`

NewFabricFcNetworkPolicyWithDefaults instantiates a new FabricFcNetworkPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcNetworkPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcNetworkPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcNetworkPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcNetworkPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcNetworkPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcNetworkPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableTrunking

`func (o *FabricFcNetworkPolicy) GetEnableTrunking() bool`

GetEnableTrunking returns the EnableTrunking field if non-nil, zero value otherwise.

### GetEnableTrunkingOk

`func (o *FabricFcNetworkPolicy) GetEnableTrunkingOk() (*bool, bool)`

GetEnableTrunkingOk returns a tuple with the EnableTrunking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrunking

`func (o *FabricFcNetworkPolicy) SetEnableTrunking(v bool)`

SetEnableTrunking sets EnableTrunking field to given value.

### HasEnableTrunking

`func (o *FabricFcNetworkPolicy) HasEnableTrunking() bool`

HasEnableTrunking returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricFcNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricFcNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricFcNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricFcNetworkPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *FabricFcNetworkPolicy) GetProfiles() []FabricSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricFcNetworkPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricFcNetworkPolicy) SetProfiles(v []FabricSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FabricFcNetworkPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FabricFcNetworkPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FabricFcNetworkPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


