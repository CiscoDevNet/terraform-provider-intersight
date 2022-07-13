# FabricFcZonePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcZonePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcZonePolicy"]
**FcTargetMembers** | Pointer to [**[]FabricFcZoneMember**](FabricFcZoneMember.md) |  | [optional] 
**FcTargetZoningType** | Pointer to **string** | Type of FC zoning. Allowed values are SIST, SIMT and None. * &#x60;SIST&#x60; - The system automatically creates one zone for each vHBA and storage port pair. Each zone has two members. * &#x60;SIMT&#x60; - The system automatically creates one zone for each vHBA. Configure this type of zoning if the number of zones created is likely to exceed the maximum supported number of zones. * &#x60;None&#x60; - FC zoning is not configured. | [optional] [default to "SIST"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricFcZonePolicy

`func NewFabricFcZonePolicy(classId string, objectType string, ) *FabricFcZonePolicy`

NewFabricFcZonePolicy instantiates a new FabricFcZonePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcZonePolicyWithDefaults

`func NewFabricFcZonePolicyWithDefaults() *FabricFcZonePolicy`

NewFabricFcZonePolicyWithDefaults instantiates a new FabricFcZonePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcZonePolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcZonePolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcZonePolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcZonePolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcZonePolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcZonePolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFcTargetMembers

`func (o *FabricFcZonePolicy) GetFcTargetMembers() []FabricFcZoneMember`

GetFcTargetMembers returns the FcTargetMembers field if non-nil, zero value otherwise.

### GetFcTargetMembersOk

`func (o *FabricFcZonePolicy) GetFcTargetMembersOk() (*[]FabricFcZoneMember, bool)`

GetFcTargetMembersOk returns a tuple with the FcTargetMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcTargetMembers

`func (o *FabricFcZonePolicy) SetFcTargetMembers(v []FabricFcZoneMember)`

SetFcTargetMembers sets FcTargetMembers field to given value.

### HasFcTargetMembers

`func (o *FabricFcZonePolicy) HasFcTargetMembers() bool`

HasFcTargetMembers returns a boolean if a field has been set.

### SetFcTargetMembersNil

`func (o *FabricFcZonePolicy) SetFcTargetMembersNil(b bool)`

 SetFcTargetMembersNil sets the value for FcTargetMembers to be an explicit nil

### UnsetFcTargetMembers
`func (o *FabricFcZonePolicy) UnsetFcTargetMembers()`

UnsetFcTargetMembers ensures that no value is present for FcTargetMembers, not even an explicit nil
### GetFcTargetZoningType

`func (o *FabricFcZonePolicy) GetFcTargetZoningType() string`

GetFcTargetZoningType returns the FcTargetZoningType field if non-nil, zero value otherwise.

### GetFcTargetZoningTypeOk

`func (o *FabricFcZonePolicy) GetFcTargetZoningTypeOk() (*string, bool)`

GetFcTargetZoningTypeOk returns a tuple with the FcTargetZoningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcTargetZoningType

`func (o *FabricFcZonePolicy) SetFcTargetZoningType(v string)`

SetFcTargetZoningType sets FcTargetZoningType field to given value.

### HasFcTargetZoningType

`func (o *FabricFcZonePolicy) HasFcTargetZoningType() bool`

HasFcTargetZoningType returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricFcZonePolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricFcZonePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricFcZonePolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricFcZonePolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


