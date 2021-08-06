# CloudTfcOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.TfcOrganization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.TfcOrganization"]
**AgentCeiling** | Pointer to **int64** | The max number of active agents allowed in this organization. | [optional] [readonly] 
**Email** | Pointer to **string** | The admin email address associated with a user under this organization. | [optional] [readonly] 
**Identity** | Pointer to **string** | The ID of the organization. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the organization. | [optional] [readonly] 
**NumTeams** | Pointer to **int64** | The number of teams under this organization. | [optional] [readonly] 
**NumUsers** | Pointer to **int64** | The number of users in this organization. | [optional] [readonly] 
**RunCeiling** | Pointer to **int64** | The max number of simultaneous runs allowed in this organization. | [optional] [readonly] 
**VcsProviders** | Pointer to **int64** | Total number of VCS providers in the organization. | [optional] [readonly] 
**Target** | Pointer to [**AssetTargetRelationship**](asset.Target.Relationship.md) |  | [optional] 

## Methods

### NewCloudTfcOrganization

`func NewCloudTfcOrganization(classId string, objectType string, ) *CloudTfcOrganization`

NewCloudTfcOrganization instantiates a new CloudTfcOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTfcOrganizationWithDefaults

`func NewCloudTfcOrganizationWithDefaults() *CloudTfcOrganization`

NewCloudTfcOrganizationWithDefaults instantiates a new CloudTfcOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudTfcOrganization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudTfcOrganization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudTfcOrganization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudTfcOrganization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudTfcOrganization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudTfcOrganization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAgentCeiling

`func (o *CloudTfcOrganization) GetAgentCeiling() int64`

GetAgentCeiling returns the AgentCeiling field if non-nil, zero value otherwise.

### GetAgentCeilingOk

`func (o *CloudTfcOrganization) GetAgentCeilingOk() (*int64, bool)`

GetAgentCeilingOk returns a tuple with the AgentCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCeiling

`func (o *CloudTfcOrganization) SetAgentCeiling(v int64)`

SetAgentCeiling sets AgentCeiling field to given value.

### HasAgentCeiling

`func (o *CloudTfcOrganization) HasAgentCeiling() bool`

HasAgentCeiling returns a boolean if a field has been set.

### GetEmail

`func (o *CloudTfcOrganization) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudTfcOrganization) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudTfcOrganization) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudTfcOrganization) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudTfcOrganization) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudTfcOrganization) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudTfcOrganization) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudTfcOrganization) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *CloudTfcOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTfcOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTfcOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTfcOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumTeams

`func (o *CloudTfcOrganization) GetNumTeams() int64`

GetNumTeams returns the NumTeams field if non-nil, zero value otherwise.

### GetNumTeamsOk

`func (o *CloudTfcOrganization) GetNumTeamsOk() (*int64, bool)`

GetNumTeamsOk returns a tuple with the NumTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTeams

`func (o *CloudTfcOrganization) SetNumTeams(v int64)`

SetNumTeams sets NumTeams field to given value.

### HasNumTeams

`func (o *CloudTfcOrganization) HasNumTeams() bool`

HasNumTeams returns a boolean if a field has been set.

### GetNumUsers

`func (o *CloudTfcOrganization) GetNumUsers() int64`

GetNumUsers returns the NumUsers field if non-nil, zero value otherwise.

### GetNumUsersOk

`func (o *CloudTfcOrganization) GetNumUsersOk() (*int64, bool)`

GetNumUsersOk returns a tuple with the NumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsers

`func (o *CloudTfcOrganization) SetNumUsers(v int64)`

SetNumUsers sets NumUsers field to given value.

### HasNumUsers

`func (o *CloudTfcOrganization) HasNumUsers() bool`

HasNumUsers returns a boolean if a field has been set.

### GetRunCeiling

`func (o *CloudTfcOrganization) GetRunCeiling() int64`

GetRunCeiling returns the RunCeiling field if non-nil, zero value otherwise.

### GetRunCeilingOk

`func (o *CloudTfcOrganization) GetRunCeilingOk() (*int64, bool)`

GetRunCeilingOk returns a tuple with the RunCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCeiling

`func (o *CloudTfcOrganization) SetRunCeiling(v int64)`

SetRunCeiling sets RunCeiling field to given value.

### HasRunCeiling

`func (o *CloudTfcOrganization) HasRunCeiling() bool`

HasRunCeiling returns a boolean if a field has been set.

### GetVcsProviders

`func (o *CloudTfcOrganization) GetVcsProviders() int64`

GetVcsProviders returns the VcsProviders field if non-nil, zero value otherwise.

### GetVcsProvidersOk

`func (o *CloudTfcOrganization) GetVcsProvidersOk() (*int64, bool)`

GetVcsProvidersOk returns a tuple with the VcsProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsProviders

`func (o *CloudTfcOrganization) SetVcsProviders(v int64)`

SetVcsProviders sets VcsProviders field to given value.

### HasVcsProviders

`func (o *CloudTfcOrganization) HasVcsProviders() bool`

HasVcsProviders returns a boolean if a field has been set.

### GetTarget

`func (o *CloudTfcOrganization) GetTarget() AssetTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudTfcOrganization) GetTargetOk() (*AssetTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudTfcOrganization) SetTarget(v AssetTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudTfcOrganization) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


