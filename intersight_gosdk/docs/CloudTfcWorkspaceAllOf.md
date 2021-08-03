# CloudTfcWorkspaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.TfcWorkspace"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.TfcWorkspace"]
**AgentPoolId** | Pointer to **string** | The agent pool associated with this workspace. | [optional] [readonly] 
**ApplyMethod** | Pointer to **bool** | Whether or not Terraform Cloud should automatically apply a successful Terraform plan. | [optional] [readonly] 
**ExecutionMode** | Pointer to **string** | Indicates where the Terraform cloud should execute the runs. | [optional] [readonly] 
**Identity** | Pointer to **string** | The unique id for this workspace. | [optional] [readonly] 
**LastRunStatus** | Pointer to **string** | The status of the last executed run in this workspace. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the workspace. | [optional] [readonly] 
**WorkspaceVariables** | Pointer to [**[]CloudTfcWorkspaceVariables**](CloudTfcWorkspaceVariables.md) |  | [optional] 
**Organization** | Pointer to [**CloudTfcOrganizationRelationship**](cloud.TfcOrganization.Relationship.md) |  | [optional] 

## Methods

### NewCloudTfcWorkspaceAllOf

`func NewCloudTfcWorkspaceAllOf(classId string, objectType string, ) *CloudTfcWorkspaceAllOf`

NewCloudTfcWorkspaceAllOf instantiates a new CloudTfcWorkspaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTfcWorkspaceAllOfWithDefaults

`func NewCloudTfcWorkspaceAllOfWithDefaults() *CloudTfcWorkspaceAllOf`

NewCloudTfcWorkspaceAllOfWithDefaults instantiates a new CloudTfcWorkspaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudTfcWorkspaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudTfcWorkspaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudTfcWorkspaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudTfcWorkspaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudTfcWorkspaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudTfcWorkspaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAgentPoolId

`func (o *CloudTfcWorkspaceAllOf) GetAgentPoolId() string`

GetAgentPoolId returns the AgentPoolId field if non-nil, zero value otherwise.

### GetAgentPoolIdOk

`func (o *CloudTfcWorkspaceAllOf) GetAgentPoolIdOk() (*string, bool)`

GetAgentPoolIdOk returns a tuple with the AgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPoolId

`func (o *CloudTfcWorkspaceAllOf) SetAgentPoolId(v string)`

SetAgentPoolId sets AgentPoolId field to given value.

### HasAgentPoolId

`func (o *CloudTfcWorkspaceAllOf) HasAgentPoolId() bool`

HasAgentPoolId returns a boolean if a field has been set.

### GetApplyMethod

`func (o *CloudTfcWorkspaceAllOf) GetApplyMethod() bool`

GetApplyMethod returns the ApplyMethod field if non-nil, zero value otherwise.

### GetApplyMethodOk

`func (o *CloudTfcWorkspaceAllOf) GetApplyMethodOk() (*bool, bool)`

GetApplyMethodOk returns a tuple with the ApplyMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMethod

`func (o *CloudTfcWorkspaceAllOf) SetApplyMethod(v bool)`

SetApplyMethod sets ApplyMethod field to given value.

### HasApplyMethod

`func (o *CloudTfcWorkspaceAllOf) HasApplyMethod() bool`

HasApplyMethod returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CloudTfcWorkspaceAllOf) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CloudTfcWorkspaceAllOf) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CloudTfcWorkspaceAllOf) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CloudTfcWorkspaceAllOf) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudTfcWorkspaceAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudTfcWorkspaceAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudTfcWorkspaceAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudTfcWorkspaceAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLastRunStatus

`func (o *CloudTfcWorkspaceAllOf) GetLastRunStatus() string`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *CloudTfcWorkspaceAllOf) GetLastRunStatusOk() (*string, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *CloudTfcWorkspaceAllOf) SetLastRunStatus(v string)`

SetLastRunStatus sets LastRunStatus field to given value.

### HasLastRunStatus

`func (o *CloudTfcWorkspaceAllOf) HasLastRunStatus() bool`

HasLastRunStatus returns a boolean if a field has been set.

### GetName

`func (o *CloudTfcWorkspaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTfcWorkspaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTfcWorkspaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTfcWorkspaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWorkspaceVariables

`func (o *CloudTfcWorkspaceAllOf) GetWorkspaceVariables() []CloudTfcWorkspaceVariables`

GetWorkspaceVariables returns the WorkspaceVariables field if non-nil, zero value otherwise.

### GetWorkspaceVariablesOk

`func (o *CloudTfcWorkspaceAllOf) GetWorkspaceVariablesOk() (*[]CloudTfcWorkspaceVariables, bool)`

GetWorkspaceVariablesOk returns a tuple with the WorkspaceVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceVariables

`func (o *CloudTfcWorkspaceAllOf) SetWorkspaceVariables(v []CloudTfcWorkspaceVariables)`

SetWorkspaceVariables sets WorkspaceVariables field to given value.

### HasWorkspaceVariables

`func (o *CloudTfcWorkspaceAllOf) HasWorkspaceVariables() bool`

HasWorkspaceVariables returns a boolean if a field has been set.

### SetWorkspaceVariablesNil

`func (o *CloudTfcWorkspaceAllOf) SetWorkspaceVariablesNil(b bool)`

 SetWorkspaceVariablesNil sets the value for WorkspaceVariables to be an explicit nil

### UnsetWorkspaceVariables
`func (o *CloudTfcWorkspaceAllOf) UnsetWorkspaceVariables()`

UnsetWorkspaceVariables ensures that no value is present for WorkspaceVariables, not even an explicit nil
### GetOrganization

`func (o *CloudTfcWorkspaceAllOf) GetOrganization() CloudTfcOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CloudTfcWorkspaceAllOf) GetOrganizationOk() (*CloudTfcOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CloudTfcWorkspaceAllOf) SetOrganization(v CloudTfcOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CloudTfcWorkspaceAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


