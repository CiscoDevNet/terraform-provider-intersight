# ChassisConfigImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.ConfigImport"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.ConfigImport"]
**Description** | Pointer to **string** | Description of the imported profile. | [optional] 
**PolicyPrefix** | Pointer to **string** | Policy prefix for the policies of the imported chassis profile. | [optional] 
**PolicyTypes** | Pointer to **[]string** |  | [optional] 
**ProfileName** | Pointer to **string** | Profile name for the imported chassis profile. | [optional] 
**Chassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**ChassisProfile** | Pointer to [**ChassisProfileRelationship**](ChassisProfileRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewChassisConfigImport

`func NewChassisConfigImport(classId string, objectType string, ) *ChassisConfigImport`

NewChassisConfigImport instantiates a new ChassisConfigImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisConfigImportWithDefaults

`func NewChassisConfigImportWithDefaults() *ChassisConfigImport`

NewChassisConfigImportWithDefaults instantiates a new ChassisConfigImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisConfigImport) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisConfigImport) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisConfigImport) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisConfigImport) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisConfigImport) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisConfigImport) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ChassisConfigImport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChassisConfigImport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChassisConfigImport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChassisConfigImport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyPrefix

`func (o *ChassisConfigImport) GetPolicyPrefix() string`

GetPolicyPrefix returns the PolicyPrefix field if non-nil, zero value otherwise.

### GetPolicyPrefixOk

`func (o *ChassisConfigImport) GetPolicyPrefixOk() (*string, bool)`

GetPolicyPrefixOk returns a tuple with the PolicyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPrefix

`func (o *ChassisConfigImport) SetPolicyPrefix(v string)`

SetPolicyPrefix sets PolicyPrefix field to given value.

### HasPolicyPrefix

`func (o *ChassisConfigImport) HasPolicyPrefix() bool`

HasPolicyPrefix returns a boolean if a field has been set.

### GetPolicyTypes

`func (o *ChassisConfigImport) GetPolicyTypes() []string`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *ChassisConfigImport) GetPolicyTypesOk() (*[]string, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *ChassisConfigImport) SetPolicyTypes(v []string)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *ChassisConfigImport) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.

### SetPolicyTypesNil

`func (o *ChassisConfigImport) SetPolicyTypesNil(b bool)`

 SetPolicyTypesNil sets the value for PolicyTypes to be an explicit nil

### UnsetPolicyTypes
`func (o *ChassisConfigImport) UnsetPolicyTypes()`

UnsetPolicyTypes ensures that no value is present for PolicyTypes, not even an explicit nil
### GetProfileName

`func (o *ChassisConfigImport) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ChassisConfigImport) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ChassisConfigImport) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ChassisConfigImport) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetChassis

`func (o *ChassisConfigImport) GetChassis() EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *ChassisConfigImport) GetChassisOk() (*EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *ChassisConfigImport) SetChassis(v EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *ChassisConfigImport) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetChassisProfile

`func (o *ChassisConfigImport) GetChassisProfile() ChassisProfileRelationship`

GetChassisProfile returns the ChassisProfile field if non-nil, zero value otherwise.

### GetChassisProfileOk

`func (o *ChassisConfigImport) GetChassisProfileOk() (*ChassisProfileRelationship, bool)`

GetChassisProfileOk returns a tuple with the ChassisProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisProfile

`func (o *ChassisConfigImport) SetChassisProfile(v ChassisProfileRelationship)`

SetChassisProfile sets ChassisProfile field to given value.

### HasChassisProfile

`func (o *ChassisConfigImport) HasChassisProfile() bool`

HasChassisProfile returns a boolean if a field has been set.

### GetOrganization

`func (o *ChassisConfigImport) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ChassisConfigImport) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ChassisConfigImport) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ChassisConfigImport) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


