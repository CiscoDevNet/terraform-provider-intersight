# MemoryPersistentMemoryPolicy

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

### NewMemoryPersistentMemoryPolicy

`func NewMemoryPersistentMemoryPolicy(classId string, objectType string, ) *MemoryPersistentMemoryPolicy`

NewMemoryPersistentMemoryPolicy instantiates a new MemoryPersistentMemoryPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryPolicyWithDefaults

`func NewMemoryPersistentMemoryPolicyWithDefaults() *MemoryPersistentMemoryPolicy`

NewMemoryPersistentMemoryPolicyWithDefaults instantiates a new MemoryPersistentMemoryPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPersistentMemoryPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPersistentMemoryPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGoals

`func (o *MemoryPersistentMemoryPolicy) GetGoals() []MemoryPersistentMemoryGoal`

GetGoals returns the Goals field if non-nil, zero value otherwise.

### GetGoalsOk

`func (o *MemoryPersistentMemoryPolicy) GetGoalsOk() (*[]MemoryPersistentMemoryGoal, bool)`

GetGoalsOk returns a tuple with the Goals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoals

`func (o *MemoryPersistentMemoryPolicy) SetGoals(v []MemoryPersistentMemoryGoal)`

SetGoals sets Goals field to given value.

### HasGoals

`func (o *MemoryPersistentMemoryPolicy) HasGoals() bool`

HasGoals returns a boolean if a field has been set.

### SetGoalsNil

`func (o *MemoryPersistentMemoryPolicy) SetGoalsNil(b bool)`

 SetGoalsNil sets the value for Goals to be an explicit nil

### UnsetGoals
`func (o *MemoryPersistentMemoryPolicy) UnsetGoals()`

UnsetGoals ensures that no value is present for Goals, not even an explicit nil
### GetLocalSecurity

`func (o *MemoryPersistentMemoryPolicy) GetLocalSecurity() MemoryPersistentMemoryLocalSecurity`

GetLocalSecurity returns the LocalSecurity field if non-nil, zero value otherwise.

### GetLocalSecurityOk

`func (o *MemoryPersistentMemoryPolicy) GetLocalSecurityOk() (*MemoryPersistentMemoryLocalSecurity, bool)`

GetLocalSecurityOk returns a tuple with the LocalSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSecurity

`func (o *MemoryPersistentMemoryPolicy) SetLocalSecurity(v MemoryPersistentMemoryLocalSecurity)`

SetLocalSecurity sets LocalSecurity field to given value.

### HasLocalSecurity

`func (o *MemoryPersistentMemoryPolicy) HasLocalSecurity() bool`

HasLocalSecurity returns a boolean if a field has been set.

### SetLocalSecurityNil

`func (o *MemoryPersistentMemoryPolicy) SetLocalSecurityNil(b bool)`

 SetLocalSecurityNil sets the value for LocalSecurity to be an explicit nil

### UnsetLocalSecurity
`func (o *MemoryPersistentMemoryPolicy) UnsetLocalSecurity()`

UnsetLocalSecurity ensures that no value is present for LocalSecurity, not even an explicit nil
### GetLogicalNamespaces

`func (o *MemoryPersistentMemoryPolicy) GetLogicalNamespaces() []MemoryPersistentMemoryLogicalNamespace`

GetLogicalNamespaces returns the LogicalNamespaces field if non-nil, zero value otherwise.

### GetLogicalNamespacesOk

`func (o *MemoryPersistentMemoryPolicy) GetLogicalNamespacesOk() (*[]MemoryPersistentMemoryLogicalNamespace, bool)`

GetLogicalNamespacesOk returns a tuple with the LogicalNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNamespaces

`func (o *MemoryPersistentMemoryPolicy) SetLogicalNamespaces(v []MemoryPersistentMemoryLogicalNamespace)`

SetLogicalNamespaces sets LogicalNamespaces field to given value.

### HasLogicalNamespaces

`func (o *MemoryPersistentMemoryPolicy) HasLogicalNamespaces() bool`

HasLogicalNamespaces returns a boolean if a field has been set.

### SetLogicalNamespacesNil

`func (o *MemoryPersistentMemoryPolicy) SetLogicalNamespacesNil(b bool)`

 SetLogicalNamespacesNil sets the value for LogicalNamespaces to be an explicit nil

### UnsetLogicalNamespaces
`func (o *MemoryPersistentMemoryPolicy) UnsetLogicalNamespaces()`

UnsetLogicalNamespaces ensures that no value is present for LogicalNamespaces, not even an explicit nil
### GetManagementMode

`func (o *MemoryPersistentMemoryPolicy) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *MemoryPersistentMemoryPolicy) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *MemoryPersistentMemoryPolicy) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *MemoryPersistentMemoryPolicy) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetRetainNamespaces

`func (o *MemoryPersistentMemoryPolicy) GetRetainNamespaces() bool`

GetRetainNamespaces returns the RetainNamespaces field if non-nil, zero value otherwise.

### GetRetainNamespacesOk

`func (o *MemoryPersistentMemoryPolicy) GetRetainNamespacesOk() (*bool, bool)`

GetRetainNamespacesOk returns a tuple with the RetainNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainNamespaces

`func (o *MemoryPersistentMemoryPolicy) SetRetainNamespaces(v bool)`

SetRetainNamespaces sets RetainNamespaces field to given value.

### HasRetainNamespaces

`func (o *MemoryPersistentMemoryPolicy) HasRetainNamespaces() bool`

HasRetainNamespaces returns a boolean if a field has been set.

### GetOrganization

`func (o *MemoryPersistentMemoryPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MemoryPersistentMemoryPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MemoryPersistentMemoryPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MemoryPersistentMemoryPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *MemoryPersistentMemoryPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MemoryPersistentMemoryPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MemoryPersistentMemoryPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MemoryPersistentMemoryPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *MemoryPersistentMemoryPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *MemoryPersistentMemoryPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


