# WorkflowAnsiblePlaySessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.AnsiblePlaySession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.AnsiblePlaySession"]
**CommandLineArguments** | Pointer to **string** | The command line arguments for running the Ansible playbook against the given endpoint. Escape character backslash needs to be used when the command line arguments contain double quotes in them. | [optional] 
**HostInventory** | Pointer to **string** | The path of the host inventory file that resides on the Ansible Endpoint target or the comma separated list of hosts on which the Ansible playbook is to be run. Make sure to suffix a comma when the list of hosts is provided as input, even if the list has only one value. | [optional] 
**PlaybookPath** | Pointer to **string** | The path of the Ansible playbook that resides on the Ansible Endpoint target. | [optional] 
**SshOpTimeout** | Pointer to **string** | SSH operation timeout value in seconds. Value provided should be string representation of an interger. | [optional] 

## Methods

### NewWorkflowAnsiblePlaySessionAllOf

`func NewWorkflowAnsiblePlaySessionAllOf(classId string, objectType string, ) *WorkflowAnsiblePlaySessionAllOf`

NewWorkflowAnsiblePlaySessionAllOf instantiates a new WorkflowAnsiblePlaySessionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAnsiblePlaySessionAllOfWithDefaults

`func NewWorkflowAnsiblePlaySessionAllOfWithDefaults() *WorkflowAnsiblePlaySessionAllOf`

NewWorkflowAnsiblePlaySessionAllOfWithDefaults instantiates a new WorkflowAnsiblePlaySessionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowAnsiblePlaySessionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowAnsiblePlaySessionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowAnsiblePlaySessionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowAnsiblePlaySessionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAnsiblePlaySessionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAnsiblePlaySessionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommandLineArguments

`func (o *WorkflowAnsiblePlaySessionAllOf) GetCommandLineArguments() string`

GetCommandLineArguments returns the CommandLineArguments field if non-nil, zero value otherwise.

### GetCommandLineArgumentsOk

`func (o *WorkflowAnsiblePlaySessionAllOf) GetCommandLineArgumentsOk() (*string, bool)`

GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArguments

`func (o *WorkflowAnsiblePlaySessionAllOf) SetCommandLineArguments(v string)`

SetCommandLineArguments sets CommandLineArguments field to given value.

### HasCommandLineArguments

`func (o *WorkflowAnsiblePlaySessionAllOf) HasCommandLineArguments() bool`

HasCommandLineArguments returns a boolean if a field has been set.

### GetHostInventory

`func (o *WorkflowAnsiblePlaySessionAllOf) GetHostInventory() string`

GetHostInventory returns the HostInventory field if non-nil, zero value otherwise.

### GetHostInventoryOk

`func (o *WorkflowAnsiblePlaySessionAllOf) GetHostInventoryOk() (*string, bool)`

GetHostInventoryOk returns a tuple with the HostInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInventory

`func (o *WorkflowAnsiblePlaySessionAllOf) SetHostInventory(v string)`

SetHostInventory sets HostInventory field to given value.

### HasHostInventory

`func (o *WorkflowAnsiblePlaySessionAllOf) HasHostInventory() bool`

HasHostInventory returns a boolean if a field has been set.

### GetPlaybookPath

`func (o *WorkflowAnsiblePlaySessionAllOf) GetPlaybookPath() string`

GetPlaybookPath returns the PlaybookPath field if non-nil, zero value otherwise.

### GetPlaybookPathOk

`func (o *WorkflowAnsiblePlaySessionAllOf) GetPlaybookPathOk() (*string, bool)`

GetPlaybookPathOk returns a tuple with the PlaybookPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybookPath

`func (o *WorkflowAnsiblePlaySessionAllOf) SetPlaybookPath(v string)`

SetPlaybookPath sets PlaybookPath field to given value.

### HasPlaybookPath

`func (o *WorkflowAnsiblePlaySessionAllOf) HasPlaybookPath() bool`

HasPlaybookPath returns a boolean if a field has been set.

### GetSshOpTimeout

`func (o *WorkflowAnsiblePlaySessionAllOf) GetSshOpTimeout() string`

GetSshOpTimeout returns the SshOpTimeout field if non-nil, zero value otherwise.

### GetSshOpTimeoutOk

`func (o *WorkflowAnsiblePlaySessionAllOf) GetSshOpTimeoutOk() (*string, bool)`

GetSshOpTimeoutOk returns a tuple with the SshOpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshOpTimeout

`func (o *WorkflowAnsiblePlaySessionAllOf) SetSshOpTimeout(v string)`

SetSshOpTimeout sets SshOpTimeout field to given value.

### HasSshOpTimeout

`func (o *WorkflowAnsiblePlaySessionAllOf) HasSshOpTimeout() bool`

HasSshOpTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


