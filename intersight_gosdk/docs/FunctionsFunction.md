# FunctionsFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "functions.Function"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "functions.Function"]
**Action** | Pointer to **string** | Action of the function such as build, deploy, undeploy, delete. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Build&#x60; - Build an instance of a Function. * &#x60;Deploy&#x60; - Deploy the build Function. * &#x60;Undeploy&#x60; - Undeploy a Function that was previously successfully deployed. * &#x60;Delete&#x60; - Delete a Function that has yet to be deployed or that was recently undeployed. | [optional] [default to "None"]
**Code** | Pointer to **string** | Custom function code for Function MO. | [optional] 
**CreateUser** | Pointer to **string** | The user identifier who created the Function. | [optional] [readonly] 
**DefaultVersion** | Pointer to **bool** | When true this function version will be used in functions table. The very first function created with a name will be set as the default version. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the function. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the function. Display name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**LastAction** | Pointer to [**NullableFunctionsFunctionLastAction**](FunctionsFunctionLastAction.md) |  | [optional] 
**ModUser** | Pointer to **string** | The user identifier who last updated the Function. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the function. Name can only contain letters (a-z), numbers (0-9), hyphen (-). | [optional] 
**State** | Pointer to **string** | Current representation of the Function MO state. * &#x60;Saved&#x60; - Function is saved, yet to be built and deployed. * &#x60;Building&#x60; - Function is currently being built. * &#x60;Built&#x60; - The Function has been built and can now be deployed. * &#x60;Deploying&#x60; - The built Function is currently being deployed. * &#x60;Deployed&#x60; - The Function has been deployed. * &#x60;Undeploying&#x60; - The deployed function is being Undeployed. * &#x60;Deleting&#x60; - The Function is being deleted. | [optional] [readonly] [default to "Saved"]
**Version** | Pointer to **int64** | The version of the function to support multiple versions. | [optional] [readonly] 
**ActionExecution** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Runtime** | Pointer to [**NullableFunctionsRuntimeRelationship**](FunctionsRuntimeRelationship.md) |  | [optional] 
**TaskDefinition** | Pointer to [**NullableWorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) |  | [optional] 

## Methods

### NewFunctionsFunction

`func NewFunctionsFunction(classId string, objectType string, ) *FunctionsFunction`

NewFunctionsFunction instantiates a new FunctionsFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsFunctionWithDefaults

`func NewFunctionsFunctionWithDefaults() *FunctionsFunction`

NewFunctionsFunctionWithDefaults instantiates a new FunctionsFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FunctionsFunction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FunctionsFunction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FunctionsFunction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FunctionsFunction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FunctionsFunction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FunctionsFunction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *FunctionsFunction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FunctionsFunction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FunctionsFunction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FunctionsFunction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCode

`func (o *FunctionsFunction) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FunctionsFunction) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FunctionsFunction) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FunctionsFunction) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateUser

`func (o *FunctionsFunction) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *FunctionsFunction) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *FunctionsFunction) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *FunctionsFunction) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultVersion

`func (o *FunctionsFunction) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *FunctionsFunction) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *FunctionsFunction) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *FunctionsFunction) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetDescription

`func (o *FunctionsFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionsFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionsFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionsFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FunctionsFunction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FunctionsFunction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FunctionsFunction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FunctionsFunction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastAction

`func (o *FunctionsFunction) GetLastAction() FunctionsFunctionLastAction`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *FunctionsFunction) GetLastActionOk() (*FunctionsFunctionLastAction, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *FunctionsFunction) SetLastAction(v FunctionsFunctionLastAction)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *FunctionsFunction) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### SetLastActionNil

`func (o *FunctionsFunction) SetLastActionNil(b bool)`

 SetLastActionNil sets the value for LastAction to be an explicit nil

### UnsetLastAction
`func (o *FunctionsFunction) UnsetLastAction()`

UnsetLastAction ensures that no value is present for LastAction, not even an explicit nil
### GetModUser

`func (o *FunctionsFunction) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *FunctionsFunction) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *FunctionsFunction) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *FunctionsFunction) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *FunctionsFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionsFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *FunctionsFunction) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FunctionsFunction) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FunctionsFunction) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FunctionsFunction) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *FunctionsFunction) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunctionsFunction) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunctionsFunction) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunctionsFunction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActionExecution

`func (o *FunctionsFunction) GetActionExecution() WorkflowWorkflowInfoRelationship`

GetActionExecution returns the ActionExecution field if non-nil, zero value otherwise.

### GetActionExecutionOk

`func (o *FunctionsFunction) GetActionExecutionOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetActionExecutionOk returns a tuple with the ActionExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionExecution

`func (o *FunctionsFunction) SetActionExecution(v WorkflowWorkflowInfoRelationship)`

SetActionExecution sets ActionExecution field to given value.

### HasActionExecution

`func (o *FunctionsFunction) HasActionExecution() bool`

HasActionExecution returns a boolean if a field has been set.

### SetActionExecutionNil

`func (o *FunctionsFunction) SetActionExecutionNil(b bool)`

 SetActionExecutionNil sets the value for ActionExecution to be an explicit nil

### UnsetActionExecution
`func (o *FunctionsFunction) UnsetActionExecution()`

UnsetActionExecution ensures that no value is present for ActionExecution, not even an explicit nil
### GetOrganization

`func (o *FunctionsFunction) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FunctionsFunction) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FunctionsFunction) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FunctionsFunction) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FunctionsFunction) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FunctionsFunction) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetRuntime

`func (o *FunctionsFunction) GetRuntime() FunctionsRuntimeRelationship`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *FunctionsFunction) GetRuntimeOk() (*FunctionsRuntimeRelationship, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *FunctionsFunction) SetRuntime(v FunctionsRuntimeRelationship)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *FunctionsFunction) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *FunctionsFunction) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *FunctionsFunction) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetTaskDefinition

`func (o *FunctionsFunction) GetTaskDefinition() WorkflowTaskDefinitionRelationship`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *FunctionsFunction) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *FunctionsFunction) SetTaskDefinition(v WorkflowTaskDefinitionRelationship)`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *FunctionsFunction) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.

### SetTaskDefinitionNil

`func (o *FunctionsFunction) SetTaskDefinitionNil(b bool)`

 SetTaskDefinitionNil sets the value for TaskDefinition to be an explicit nil

### UnsetTaskDefinition
`func (o *FunctionsFunction) UnsetTaskDefinition()`

UnsetTaskDefinition ensures that no value is present for TaskDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


