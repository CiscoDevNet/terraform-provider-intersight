# MemoryPersistentMemoryPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.PersistentMemoryPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.PersistentMemoryPolicy"]
**Goals** | Pointer to [**[]MemoryPersistentMemoryGoal**](MemoryPersistentMemoryGoal.md) |  | [optional] 
**LocalSecurity** | Pointer to [**NullableMemoryPersistentMemoryLocalSecurity**](MemoryPersistentMemoryLocalSecurity.md) |  | [optional] 
**LogicalNamespaces** | Pointer to [**[]MemoryPersistentMemoryLogicalNamespace**](MemoryPersistentMemoryLogicalNamespace.md) |  | [optional] 
**ManagementMode** | Pointer to **string** | Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System. * &#x60;configured-from-intersight&#x60; - The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy. * &#x60;configured-from-operating-system&#x60; - The Persistent Memory Modules are configured from operating system thorugh OS tools. | [optional] [default to "configured-from-intersight"]
**RetainNamespaces** | Pointer to **bool** | Persistent Memory Namespaces to be retained or not. | [optional] [default to true]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewMemoryPersistentMemoryPolicyAllOf

`func NewMemoryPersistentMemoryPolicyAllOf(classId string, objectType string, ) *MemoryPersistentMemoryPolicyAllOf`

NewMemoryPersistentMemoryPolicyAllOf instantiates a new MemoryPersistentMemoryPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryPolicyAllOfWithDefaults

`func NewMemoryPersistentMemoryPolicyAllOfWithDefaults() *MemoryPersistentMemoryPolicyAllOf`

NewMemoryPersistentMemoryPolicyAllOfWithDefaults instantiates a new MemoryPersistentMemoryPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPersistentMemoryPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPersistentMemoryPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGoals

`func (o *MemoryPersistentMemoryPolicyAllOf) GetGoals() []MemoryPersistentMemoryGoal`

GetGoals returns the Goals field if non-nil, zero value otherwise.

### GetGoalsOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetGoalsOk() (*[]MemoryPersistentMemoryGoal, bool)`

GetGoalsOk returns a tuple with the Goals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoals

`func (o *MemoryPersistentMemoryPolicyAllOf) SetGoals(v []MemoryPersistentMemoryGoal)`

SetGoals sets Goals field to given value.

### HasGoals

`func (o *MemoryPersistentMemoryPolicyAllOf) HasGoals() bool`

HasGoals returns a boolean if a field has been set.

### SetGoalsNil

`func (o *MemoryPersistentMemoryPolicyAllOf) SetGoalsNil(b bool)`

 SetGoalsNil sets the value for Goals to be an explicit nil

### UnsetGoals
`func (o *MemoryPersistentMemoryPolicyAllOf) UnsetGoals()`

UnsetGoals ensures that no value is present for Goals, not even an explicit nil
### GetLocalSecurity

`func (o *MemoryPersistentMemoryPolicyAllOf) GetLocalSecurity() MemoryPersistentMemoryLocalSecurity`

GetLocalSecurity returns the LocalSecurity field if non-nil, zero value otherwise.

### GetLocalSecurityOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetLocalSecurityOk() (*MemoryPersistentMemoryLocalSecurity, bool)`

GetLocalSecurityOk returns a tuple with the LocalSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSecurity

`func (o *MemoryPersistentMemoryPolicyAllOf) SetLocalSecurity(v MemoryPersistentMemoryLocalSecurity)`

SetLocalSecurity sets LocalSecurity field to given value.

### HasLocalSecurity

`func (o *MemoryPersistentMemoryPolicyAllOf) HasLocalSecurity() bool`

HasLocalSecurity returns a boolean if a field has been set.

### SetLocalSecurityNil

`func (o *MemoryPersistentMemoryPolicyAllOf) SetLocalSecurityNil(b bool)`

 SetLocalSecurityNil sets the value for LocalSecurity to be an explicit nil

### UnsetLocalSecurity
`func (o *MemoryPersistentMemoryPolicyAllOf) UnsetLocalSecurity()`

UnsetLocalSecurity ensures that no value is present for LocalSecurity, not even an explicit nil
### GetLogicalNamespaces

`func (o *MemoryPersistentMemoryPolicyAllOf) GetLogicalNamespaces() []MemoryPersistentMemoryLogicalNamespace`

GetLogicalNamespaces returns the LogicalNamespaces field if non-nil, zero value otherwise.

### GetLogicalNamespacesOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetLogicalNamespacesOk() (*[]MemoryPersistentMemoryLogicalNamespace, bool)`

GetLogicalNamespacesOk returns a tuple with the LogicalNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNamespaces

`func (o *MemoryPersistentMemoryPolicyAllOf) SetLogicalNamespaces(v []MemoryPersistentMemoryLogicalNamespace)`

SetLogicalNamespaces sets LogicalNamespaces field to given value.

### HasLogicalNamespaces

`func (o *MemoryPersistentMemoryPolicyAllOf) HasLogicalNamespaces() bool`

HasLogicalNamespaces returns a boolean if a field has been set.

### SetLogicalNamespacesNil

`func (o *MemoryPersistentMemoryPolicyAllOf) SetLogicalNamespacesNil(b bool)`

 SetLogicalNamespacesNil sets the value for LogicalNamespaces to be an explicit nil

### UnsetLogicalNamespaces
`func (o *MemoryPersistentMemoryPolicyAllOf) UnsetLogicalNamespaces()`

UnsetLogicalNamespaces ensures that no value is present for LogicalNamespaces, not even an explicit nil
### GetManagementMode

`func (o *MemoryPersistentMemoryPolicyAllOf) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *MemoryPersistentMemoryPolicyAllOf) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *MemoryPersistentMemoryPolicyAllOf) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetRetainNamespaces

`func (o *MemoryPersistentMemoryPolicyAllOf) GetRetainNamespaces() bool`

GetRetainNamespaces returns the RetainNamespaces field if non-nil, zero value otherwise.

### GetRetainNamespacesOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetRetainNamespacesOk() (*bool, bool)`

GetRetainNamespacesOk returns a tuple with the RetainNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainNamespaces

`func (o *MemoryPersistentMemoryPolicyAllOf) SetRetainNamespaces(v bool)`

SetRetainNamespaces sets RetainNamespaces field to given value.

### HasRetainNamespaces

`func (o *MemoryPersistentMemoryPolicyAllOf) HasRetainNamespaces() bool`

HasRetainNamespaces returns a boolean if a field has been set.

### GetOrganization

`func (o *MemoryPersistentMemoryPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MemoryPersistentMemoryPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MemoryPersistentMemoryPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *MemoryPersistentMemoryPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MemoryPersistentMemoryPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MemoryPersistentMemoryPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MemoryPersistentMemoryPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *MemoryPersistentMemoryPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *MemoryPersistentMemoryPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


