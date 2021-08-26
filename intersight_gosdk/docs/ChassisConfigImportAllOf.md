# ChassisConfigImportAllOf

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

### NewChassisConfigImportAllOf

`func NewChassisConfigImportAllOf(classId string, objectType string, ) *ChassisConfigImportAllOf`

NewChassisConfigImportAllOf instantiates a new ChassisConfigImportAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisConfigImportAllOfWithDefaults

`func NewChassisConfigImportAllOfWithDefaults() *ChassisConfigImportAllOf`

NewChassisConfigImportAllOfWithDefaults instantiates a new ChassisConfigImportAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisConfigImportAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisConfigImportAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisConfigImportAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisConfigImportAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisConfigImportAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisConfigImportAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *ChassisConfigImportAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChassisConfigImportAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChassisConfigImportAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChassisConfigImportAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyPrefix

`func (o *ChassisConfigImportAllOf) GetPolicyPrefix() string`

GetPolicyPrefix returns the PolicyPrefix field if non-nil, zero value otherwise.

### GetPolicyPrefixOk

`func (o *ChassisConfigImportAllOf) GetPolicyPrefixOk() (*string, bool)`

GetPolicyPrefixOk returns a tuple with the PolicyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyPrefix

`func (o *ChassisConfigImportAllOf) SetPolicyPrefix(v string)`

SetPolicyPrefix sets PolicyPrefix field to given value.

### HasPolicyPrefix

`func (o *ChassisConfigImportAllOf) HasPolicyPrefix() bool`

HasPolicyPrefix returns a boolean if a field has been set.

### GetPolicyTypes

`func (o *ChassisConfigImportAllOf) GetPolicyTypes() []string`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *ChassisConfigImportAllOf) GetPolicyTypesOk() (*[]string, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *ChassisConfigImportAllOf) SetPolicyTypes(v []string)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *ChassisConfigImportAllOf) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.

### SetPolicyTypesNil

`func (o *ChassisConfigImportAllOf) SetPolicyTypesNil(b bool)`

 SetPolicyTypesNil sets the value for PolicyTypes to be an explicit nil

### UnsetPolicyTypes
`func (o *ChassisConfigImportAllOf) UnsetPolicyTypes()`

UnsetPolicyTypes ensures that no value is present for PolicyTypes, not even an explicit nil
### GetProfileName

`func (o *ChassisConfigImportAllOf) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ChassisConfigImportAllOf) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ChassisConfigImportAllOf) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ChassisConfigImportAllOf) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetChassis

`func (o *ChassisConfigImportAllOf) GetChassis() EquipmentChassisRelationship`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *ChassisConfigImportAllOf) GetChassisOk() (*EquipmentChassisRelationship, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *ChassisConfigImportAllOf) SetChassis(v EquipmentChassisRelationship)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *ChassisConfigImportAllOf) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetChassisProfile

`func (o *ChassisConfigImportAllOf) GetChassisProfile() ChassisProfileRelationship`

GetChassisProfile returns the ChassisProfile field if non-nil, zero value otherwise.

### GetChassisProfileOk

`func (o *ChassisConfigImportAllOf) GetChassisProfileOk() (*ChassisProfileRelationship, bool)`

GetChassisProfileOk returns a tuple with the ChassisProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisProfile

`func (o *ChassisConfigImportAllOf) SetChassisProfile(v ChassisProfileRelationship)`

SetChassisProfile sets ChassisProfile field to given value.

### HasChassisProfile

`func (o *ChassisConfigImportAllOf) HasChassisProfile() bool`

HasChassisProfile returns a boolean if a field has been set.

### GetOrganization

`func (o *ChassisConfigImportAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ChassisConfigImportAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ChassisConfigImportAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ChassisConfigImportAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


