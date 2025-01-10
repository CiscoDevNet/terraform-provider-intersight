# FunctionsFunctionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "functions.FunctionVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "functions.FunctionVersion"]
**Code** | Pointer to **string** | Custom function code for Function MO. | [optional] 
**CreateUser** | Pointer to **string** | The user identifier who created the Function. | [optional] [readonly] 
**DefaultVersion** | Pointer to **bool** | When true this function version will be used in functions table. The very first function created with a name will be set as the default version. | [optional] [readonly] 
**LastAction** | Pointer to [**NullableFunctionsFunctionLastAction**](FunctionsFunctionLastAction.md) |  | [optional] 
**ModUser** | Pointer to **string** | The user identifier who last updated the Function. | [optional] [readonly] 
**State** | Pointer to **string** | Current representation of the Function MO state. * &#x60;Saved&#x60; - Function is saved, yet to be built and deployed. * &#x60;Building&#x60; - Function is currently being built. * &#x60;Built&#x60; - The Function has been built and can now be deployed. * &#x60;Deploying&#x60; - The built Function is currently being deployed. * &#x60;Deployed&#x60; - The Function has been deployed. * &#x60;Undeploying&#x60; - The deployed function is being Undeployed. * &#x60;Deleting&#x60; - The Function is being deleted. | [optional] [readonly] [default to "Saved"]
**Version** | Pointer to **int64** | The version of the function to support multiple versions. | [optional] [readonly] 
**ActionExecution** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 
**Function** | Pointer to [**NullableFunctionsFunctionRelationship**](FunctionsFunctionRelationship.md) |  | [optional] 
**Runtime** | Pointer to [**NullableFunctionsRuntimeRelationship**](FunctionsRuntimeRelationship.md) |  | [optional] 

## Methods

### NewFunctionsFunctionVersion

`func NewFunctionsFunctionVersion(classId string, objectType string, ) *FunctionsFunctionVersion`

NewFunctionsFunctionVersion instantiates a new FunctionsFunctionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsFunctionVersionWithDefaults

`func NewFunctionsFunctionVersionWithDefaults() *FunctionsFunctionVersion`

NewFunctionsFunctionVersionWithDefaults instantiates a new FunctionsFunctionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FunctionsFunctionVersion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FunctionsFunctionVersion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FunctionsFunctionVersion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FunctionsFunctionVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FunctionsFunctionVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FunctionsFunctionVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCode

`func (o *FunctionsFunctionVersion) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FunctionsFunctionVersion) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FunctionsFunctionVersion) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FunctionsFunctionVersion) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateUser

`func (o *FunctionsFunctionVersion) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *FunctionsFunctionVersion) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *FunctionsFunctionVersion) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *FunctionsFunctionVersion) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultVersion

`func (o *FunctionsFunctionVersion) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *FunctionsFunctionVersion) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *FunctionsFunctionVersion) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *FunctionsFunctionVersion) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetLastAction

`func (o *FunctionsFunctionVersion) GetLastAction() FunctionsFunctionLastAction`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *FunctionsFunctionVersion) GetLastActionOk() (*FunctionsFunctionLastAction, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *FunctionsFunctionVersion) SetLastAction(v FunctionsFunctionLastAction)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *FunctionsFunctionVersion) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### SetLastActionNil

`func (o *FunctionsFunctionVersion) SetLastActionNil(b bool)`

 SetLastActionNil sets the value for LastAction to be an explicit nil

### UnsetLastAction
`func (o *FunctionsFunctionVersion) UnsetLastAction()`

UnsetLastAction ensures that no value is present for LastAction, not even an explicit nil
### GetModUser

`func (o *FunctionsFunctionVersion) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *FunctionsFunctionVersion) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *FunctionsFunctionVersion) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *FunctionsFunctionVersion) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetState

`func (o *FunctionsFunctionVersion) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FunctionsFunctionVersion) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FunctionsFunctionVersion) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FunctionsFunctionVersion) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *FunctionsFunctionVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunctionsFunctionVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunctionsFunctionVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunctionsFunctionVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActionExecution

`func (o *FunctionsFunctionVersion) GetActionExecution() WorkflowWorkflowInfoRelationship`

GetActionExecution returns the ActionExecution field if non-nil, zero value otherwise.

### GetActionExecutionOk

`func (o *FunctionsFunctionVersion) GetActionExecutionOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetActionExecutionOk returns a tuple with the ActionExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionExecution

`func (o *FunctionsFunctionVersion) SetActionExecution(v WorkflowWorkflowInfoRelationship)`

SetActionExecution sets ActionExecution field to given value.

### HasActionExecution

`func (o *FunctionsFunctionVersion) HasActionExecution() bool`

HasActionExecution returns a boolean if a field has been set.

### SetActionExecutionNil

`func (o *FunctionsFunctionVersion) SetActionExecutionNil(b bool)`

 SetActionExecutionNil sets the value for ActionExecution to be an explicit nil

### UnsetActionExecution
`func (o *FunctionsFunctionVersion) UnsetActionExecution()`

UnsetActionExecution ensures that no value is present for ActionExecution, not even an explicit nil
### GetFunction

`func (o *FunctionsFunctionVersion) GetFunction() FunctionsFunctionRelationship`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *FunctionsFunctionVersion) GetFunctionOk() (*FunctionsFunctionRelationship, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *FunctionsFunctionVersion) SetFunction(v FunctionsFunctionRelationship)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *FunctionsFunctionVersion) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### SetFunctionNil

`func (o *FunctionsFunctionVersion) SetFunctionNil(b bool)`

 SetFunctionNil sets the value for Function to be an explicit nil

### UnsetFunction
`func (o *FunctionsFunctionVersion) UnsetFunction()`

UnsetFunction ensures that no value is present for Function, not even an explicit nil
### GetRuntime

`func (o *FunctionsFunctionVersion) GetRuntime() FunctionsRuntimeRelationship`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *FunctionsFunctionVersion) GetRuntimeOk() (*FunctionsRuntimeRelationship, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *FunctionsFunctionVersion) SetRuntime(v FunctionsRuntimeRelationship)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *FunctionsFunctionVersion) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *FunctionsFunctionVersion) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *FunctionsFunctionVersion) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


