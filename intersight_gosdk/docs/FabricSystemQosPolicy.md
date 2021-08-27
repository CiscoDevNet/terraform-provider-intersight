# FabricSystemQosPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SystemQosPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SystemQosPolicy"]
**Classes** | Pointer to [**[]FabricQosClass**](FabricQosClass.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricSystemQosPolicy

`func NewFabricSystemQosPolicy(classId string, objectType string, ) *FabricSystemQosPolicy`

NewFabricSystemQosPolicy instantiates a new FabricSystemQosPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSystemQosPolicyWithDefaults

`func NewFabricSystemQosPolicyWithDefaults() *FabricSystemQosPolicy`

NewFabricSystemQosPolicyWithDefaults instantiates a new FabricSystemQosPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSystemQosPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSystemQosPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSystemQosPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSystemQosPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSystemQosPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSystemQosPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClasses

`func (o *FabricSystemQosPolicy) GetClasses() []FabricQosClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *FabricSystemQosPolicy) GetClassesOk() (*[]FabricQosClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *FabricSystemQosPolicy) SetClasses(v []FabricQosClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *FabricSystemQosPolicy) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### SetClassesNil

`func (o *FabricSystemQosPolicy) SetClassesNil(b bool)`

 SetClassesNil sets the value for Classes to be an explicit nil

### UnsetClasses
`func (o *FabricSystemQosPolicy) UnsetClasses()`

UnsetClasses ensures that no value is present for Classes, not even an explicit nil
### GetOrganization

`func (o *FabricSystemQosPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSystemQosPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSystemQosPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSystemQosPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *FabricSystemQosPolicy) GetProfiles() []FabricSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricSystemQosPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricSystemQosPolicy) SetProfiles(v []FabricSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FabricSystemQosPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FabricSystemQosPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FabricSystemQosPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


