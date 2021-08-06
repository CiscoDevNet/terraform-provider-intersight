# CloudTfcOrganizationAllOf

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

### NewCloudTfcOrganizationAllOf

`func NewCloudTfcOrganizationAllOf(classId string, objectType string, ) *CloudTfcOrganizationAllOf`

NewCloudTfcOrganizationAllOf instantiates a new CloudTfcOrganizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTfcOrganizationAllOfWithDefaults

`func NewCloudTfcOrganizationAllOfWithDefaults() *CloudTfcOrganizationAllOf`

NewCloudTfcOrganizationAllOfWithDefaults instantiates a new CloudTfcOrganizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudTfcOrganizationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudTfcOrganizationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudTfcOrganizationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudTfcOrganizationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudTfcOrganizationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudTfcOrganizationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAgentCeiling

`func (o *CloudTfcOrganizationAllOf) GetAgentCeiling() int64`

GetAgentCeiling returns the AgentCeiling field if non-nil, zero value otherwise.

### GetAgentCeilingOk

`func (o *CloudTfcOrganizationAllOf) GetAgentCeilingOk() (*int64, bool)`

GetAgentCeilingOk returns a tuple with the AgentCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCeiling

`func (o *CloudTfcOrganizationAllOf) SetAgentCeiling(v int64)`

SetAgentCeiling sets AgentCeiling field to given value.

### HasAgentCeiling

`func (o *CloudTfcOrganizationAllOf) HasAgentCeiling() bool`

HasAgentCeiling returns a boolean if a field has been set.

### GetEmail

`func (o *CloudTfcOrganizationAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudTfcOrganizationAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudTfcOrganizationAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudTfcOrganizationAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudTfcOrganizationAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudTfcOrganizationAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudTfcOrganizationAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudTfcOrganizationAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *CloudTfcOrganizationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTfcOrganizationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTfcOrganizationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTfcOrganizationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumTeams

`func (o *CloudTfcOrganizationAllOf) GetNumTeams() int64`

GetNumTeams returns the NumTeams field if non-nil, zero value otherwise.

### GetNumTeamsOk

`func (o *CloudTfcOrganizationAllOf) GetNumTeamsOk() (*int64, bool)`

GetNumTeamsOk returns a tuple with the NumTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTeams

`func (o *CloudTfcOrganizationAllOf) SetNumTeams(v int64)`

SetNumTeams sets NumTeams field to given value.

### HasNumTeams

`func (o *CloudTfcOrganizationAllOf) HasNumTeams() bool`

HasNumTeams returns a boolean if a field has been set.

### GetNumUsers

`func (o *CloudTfcOrganizationAllOf) GetNumUsers() int64`

GetNumUsers returns the NumUsers field if non-nil, zero value otherwise.

### GetNumUsersOk

`func (o *CloudTfcOrganizationAllOf) GetNumUsersOk() (*int64, bool)`

GetNumUsersOk returns a tuple with the NumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsers

`func (o *CloudTfcOrganizationAllOf) SetNumUsers(v int64)`

SetNumUsers sets NumUsers field to given value.

### HasNumUsers

`func (o *CloudTfcOrganizationAllOf) HasNumUsers() bool`

HasNumUsers returns a boolean if a field has been set.

### GetRunCeiling

`func (o *CloudTfcOrganizationAllOf) GetRunCeiling() int64`

GetRunCeiling returns the RunCeiling field if non-nil, zero value otherwise.

### GetRunCeilingOk

`func (o *CloudTfcOrganizationAllOf) GetRunCeilingOk() (*int64, bool)`

GetRunCeilingOk returns a tuple with the RunCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCeiling

`func (o *CloudTfcOrganizationAllOf) SetRunCeiling(v int64)`

SetRunCeiling sets RunCeiling field to given value.

### HasRunCeiling

`func (o *CloudTfcOrganizationAllOf) HasRunCeiling() bool`

HasRunCeiling returns a boolean if a field has been set.

### GetVcsProviders

`func (o *CloudTfcOrganizationAllOf) GetVcsProviders() int64`

GetVcsProviders returns the VcsProviders field if non-nil, zero value otherwise.

### GetVcsProvidersOk

`func (o *CloudTfcOrganizationAllOf) GetVcsProvidersOk() (*int64, bool)`

GetVcsProvidersOk returns a tuple with the VcsProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsProviders

`func (o *CloudTfcOrganizationAllOf) SetVcsProviders(v int64)`

SetVcsProviders sets VcsProviders field to given value.

### HasVcsProviders

`func (o *CloudTfcOrganizationAllOf) HasVcsProviders() bool`

HasVcsProviders returns a boolean if a field has been set.

### GetTarget

`func (o *CloudTfcOrganizationAllOf) GetTarget() AssetTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudTfcOrganizationAllOf) GetTargetOk() (*AssetTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudTfcOrganizationAllOf) SetTarget(v AssetTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudTfcOrganizationAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


