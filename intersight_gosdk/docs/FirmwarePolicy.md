# FirmwarePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.Policy"]
**ExcludeComponentList** | Pointer to **[]string** |  | [optional] 
**ModelBundleCombo** | Pointer to [**[]FirmwareModelBundleVersion**](FirmwareModelBundleVersion.md) |  | [optional] 
**TargetPlatform** | Pointer to **string** | The target platform on which the policy to be applied. Either standalone or connected. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "Standalone"]
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewFirmwarePolicy

`func NewFirmwarePolicy(classId string, objectType string, ) *FirmwarePolicy`

NewFirmwarePolicy instantiates a new FirmwarePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwarePolicyWithDefaults

`func NewFirmwarePolicyWithDefaults() *FirmwarePolicy`

NewFirmwarePolicyWithDefaults instantiates a new FirmwarePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwarePolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwarePolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwarePolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwarePolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwarePolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwarePolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeComponentList

`func (o *FirmwarePolicy) GetExcludeComponentList() []string`

GetExcludeComponentList returns the ExcludeComponentList field if non-nil, zero value otherwise.

### GetExcludeComponentListOk

`func (o *FirmwarePolicy) GetExcludeComponentListOk() (*[]string, bool)`

GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentList

`func (o *FirmwarePolicy) SetExcludeComponentList(v []string)`

SetExcludeComponentList sets ExcludeComponentList field to given value.

### HasExcludeComponentList

`func (o *FirmwarePolicy) HasExcludeComponentList() bool`

HasExcludeComponentList returns a boolean if a field has been set.

### SetExcludeComponentListNil

`func (o *FirmwarePolicy) SetExcludeComponentListNil(b bool)`

 SetExcludeComponentListNil sets the value for ExcludeComponentList to be an explicit nil

### UnsetExcludeComponentList
`func (o *FirmwarePolicy) UnsetExcludeComponentList()`

UnsetExcludeComponentList ensures that no value is present for ExcludeComponentList, not even an explicit nil
### GetModelBundleCombo

`func (o *FirmwarePolicy) GetModelBundleCombo() []FirmwareModelBundleVersion`

GetModelBundleCombo returns the ModelBundleCombo field if non-nil, zero value otherwise.

### GetModelBundleComboOk

`func (o *FirmwarePolicy) GetModelBundleComboOk() (*[]FirmwareModelBundleVersion, bool)`

GetModelBundleComboOk returns a tuple with the ModelBundleCombo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelBundleCombo

`func (o *FirmwarePolicy) SetModelBundleCombo(v []FirmwareModelBundleVersion)`

SetModelBundleCombo sets ModelBundleCombo field to given value.

### HasModelBundleCombo

`func (o *FirmwarePolicy) HasModelBundleCombo() bool`

HasModelBundleCombo returns a boolean if a field has been set.

### SetModelBundleComboNil

`func (o *FirmwarePolicy) SetModelBundleComboNil(b bool)`

 SetModelBundleComboNil sets the value for ModelBundleCombo to be an explicit nil

### UnsetModelBundleCombo
`func (o *FirmwarePolicy) UnsetModelBundleCombo()`

UnsetModelBundleCombo ensures that no value is present for ModelBundleCombo, not even an explicit nil
### GetTargetPlatform

`func (o *FirmwarePolicy) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *FirmwarePolicy) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *FirmwarePolicy) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *FirmwarePolicy) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetOrganization

`func (o *FirmwarePolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FirmwarePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FirmwarePolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FirmwarePolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FirmwarePolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FirmwarePolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *FirmwarePolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FirmwarePolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FirmwarePolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FirmwarePolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FirmwarePolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FirmwarePolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


