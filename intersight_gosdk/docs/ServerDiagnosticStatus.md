# ServerDiagnosticStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.DiagnosticStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.DiagnosticStatus"]
**DiagnosticsType** | Pointer to **string** | Type of diagnostics to be performed on the server hardware components. * &#x60;Quick&#x60; - Perform fast and limited diagnostics on server hardware components. * &#x60;Comprehensive&#x60; - Perform slow and extensive diagnostics on server hardware components. | [optional] [default to "Quick"]
**Name** | Pointer to **string** | The name of the diagnostics being run. | [optional] [readonly] 
**OverallState** | Pointer to **string** | The overall state of the diagnostics being run. * &#x60;Queued&#x60; - The diagnostics are queued. * &#x60;InProgress&#x60; - The diagnostics are in progress. * &#x60;Completed&#x60; - The diagnostics are completed. * &#x60;Failed&#x60; - The diagnostics have failed. * &#x60;Terminated&#x60; - The diagnostics are terminated. | [optional] [readonly] [default to "Queued"]
**Progress** | Pointer to **int64** | The progress of the diagnostics being run. | [optional] [readonly] 
**Result** | Pointer to [**[]ServerDiagnosticResult**](ServerDiagnosticResult.md) |  | [optional] 
**Diagnostics** | Pointer to [**NullableServerDiagnosticsRelationship**](ServerDiagnosticsRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewServerDiagnosticStatus

`func NewServerDiagnosticStatus(classId string, objectType string, ) *ServerDiagnosticStatus`

NewServerDiagnosticStatus instantiates a new ServerDiagnosticStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDiagnosticStatusWithDefaults

`func NewServerDiagnosticStatusWithDefaults() *ServerDiagnosticStatus`

NewServerDiagnosticStatusWithDefaults instantiates a new ServerDiagnosticStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerDiagnosticStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerDiagnosticStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerDiagnosticStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerDiagnosticStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerDiagnosticStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerDiagnosticStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiagnosticsType

`func (o *ServerDiagnosticStatus) GetDiagnosticsType() string`

GetDiagnosticsType returns the DiagnosticsType field if non-nil, zero value otherwise.

### GetDiagnosticsTypeOk

`func (o *ServerDiagnosticStatus) GetDiagnosticsTypeOk() (*string, bool)`

GetDiagnosticsTypeOk returns a tuple with the DiagnosticsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsType

`func (o *ServerDiagnosticStatus) SetDiagnosticsType(v string)`

SetDiagnosticsType sets DiagnosticsType field to given value.

### HasDiagnosticsType

`func (o *ServerDiagnosticStatus) HasDiagnosticsType() bool`

HasDiagnosticsType returns a boolean if a field has been set.

### GetName

`func (o *ServerDiagnosticStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerDiagnosticStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerDiagnosticStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerDiagnosticStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverallState

`func (o *ServerDiagnosticStatus) GetOverallState() string`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *ServerDiagnosticStatus) GetOverallStateOk() (*string, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallState

`func (o *ServerDiagnosticStatus) SetOverallState(v string)`

SetOverallState sets OverallState field to given value.

### HasOverallState

`func (o *ServerDiagnosticStatus) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### GetProgress

`func (o *ServerDiagnosticStatus) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ServerDiagnosticStatus) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ServerDiagnosticStatus) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ServerDiagnosticStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResult

`func (o *ServerDiagnosticStatus) GetResult() []ServerDiagnosticResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ServerDiagnosticStatus) GetResultOk() (*[]ServerDiagnosticResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ServerDiagnosticStatus) SetResult(v []ServerDiagnosticResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ServerDiagnosticStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ServerDiagnosticStatus) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ServerDiagnosticStatus) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetDiagnostics

`func (o *ServerDiagnosticStatus) GetDiagnostics() ServerDiagnosticsRelationship`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *ServerDiagnosticStatus) GetDiagnosticsOk() (*ServerDiagnosticsRelationship, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *ServerDiagnosticStatus) SetDiagnostics(v ServerDiagnosticsRelationship)`

SetDiagnostics sets Diagnostics field to given value.

### HasDiagnostics

`func (o *ServerDiagnosticStatus) HasDiagnostics() bool`

HasDiagnostics returns a boolean if a field has been set.

### SetDiagnosticsNil

`func (o *ServerDiagnosticStatus) SetDiagnosticsNil(b bool)`

 SetDiagnosticsNil sets the value for Diagnostics to be an explicit nil

### UnsetDiagnostics
`func (o *ServerDiagnosticStatus) UnsetDiagnostics()`

UnsetDiagnostics ensures that no value is present for Diagnostics, not even an explicit nil
### GetServer

`func (o *ServerDiagnosticStatus) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerDiagnosticStatus) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerDiagnosticStatus) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ServerDiagnosticStatus) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *ServerDiagnosticStatus) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *ServerDiagnosticStatus) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetWorkflow

`func (o *ServerDiagnosticStatus) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ServerDiagnosticStatus) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ServerDiagnosticStatus) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ServerDiagnosticStatus) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *ServerDiagnosticStatus) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *ServerDiagnosticStatus) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


