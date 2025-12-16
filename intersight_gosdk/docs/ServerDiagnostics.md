# ServerDiagnostics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.Diagnostics"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.Diagnostics"]
**Action** | Pointer to **string** | The action to be performed on the server whether to start or to cancel the diagnostics. * &#x60;Start&#x60; - Mark the start of the diagnostics on the server. * &#x60;Cancel&#x60; - Mark the cancellation of the diagnostics on the server. * &#x60;Complete&#x60; - Mark the completion of the diagnostics on the server. | [optional] [default to "Start"]
**ComponentList** | Pointer to **[]string** |  | [optional] 
**DiagnosticsType** | Pointer to **string** | Type of diagnostics to be performed on the server hardware components. * &#x60;Quick&#x60; - Perform fast and limited diagnostics on server hardware components. * &#x60;Comprehensive&#x60; - Perform slow and extensive diagnostics on server hardware components. | [optional] [default to "Quick"]
**Name** | Pointer to **string** | The name of the diagnostics being run. | [optional] 
**DiagnosticStatus** | Pointer to [**NullableServerDiagnosticStatusRelationship**](ServerDiagnosticStatusRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewServerDiagnostics

`func NewServerDiagnostics(classId string, objectType string, ) *ServerDiagnostics`

NewServerDiagnostics instantiates a new ServerDiagnostics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDiagnosticsWithDefaults

`func NewServerDiagnosticsWithDefaults() *ServerDiagnostics`

NewServerDiagnosticsWithDefaults instantiates a new ServerDiagnostics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerDiagnostics) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerDiagnostics) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerDiagnostics) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerDiagnostics) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerDiagnostics) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerDiagnostics) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *ServerDiagnostics) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ServerDiagnostics) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ServerDiagnostics) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ServerDiagnostics) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComponentList

`func (o *ServerDiagnostics) GetComponentList() []string`

GetComponentList returns the ComponentList field if non-nil, zero value otherwise.

### GetComponentListOk

`func (o *ServerDiagnostics) GetComponentListOk() (*[]string, bool)`

GetComponentListOk returns a tuple with the ComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentList

`func (o *ServerDiagnostics) SetComponentList(v []string)`

SetComponentList sets ComponentList field to given value.

### HasComponentList

`func (o *ServerDiagnostics) HasComponentList() bool`

HasComponentList returns a boolean if a field has been set.

### SetComponentListNil

`func (o *ServerDiagnostics) SetComponentListNil(b bool)`

 SetComponentListNil sets the value for ComponentList to be an explicit nil

### UnsetComponentList
`func (o *ServerDiagnostics) UnsetComponentList()`

UnsetComponentList ensures that no value is present for ComponentList, not even an explicit nil
### GetDiagnosticsType

`func (o *ServerDiagnostics) GetDiagnosticsType() string`

GetDiagnosticsType returns the DiagnosticsType field if non-nil, zero value otherwise.

### GetDiagnosticsTypeOk

`func (o *ServerDiagnostics) GetDiagnosticsTypeOk() (*string, bool)`

GetDiagnosticsTypeOk returns a tuple with the DiagnosticsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsType

`func (o *ServerDiagnostics) SetDiagnosticsType(v string)`

SetDiagnosticsType sets DiagnosticsType field to given value.

### HasDiagnosticsType

`func (o *ServerDiagnostics) HasDiagnosticsType() bool`

HasDiagnosticsType returns a boolean if a field has been set.

### GetName

`func (o *ServerDiagnostics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerDiagnostics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerDiagnostics) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerDiagnostics) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDiagnosticStatus

`func (o *ServerDiagnostics) GetDiagnosticStatus() ServerDiagnosticStatusRelationship`

GetDiagnosticStatus returns the DiagnosticStatus field if non-nil, zero value otherwise.

### GetDiagnosticStatusOk

`func (o *ServerDiagnostics) GetDiagnosticStatusOk() (*ServerDiagnosticStatusRelationship, bool)`

GetDiagnosticStatusOk returns a tuple with the DiagnosticStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticStatus

`func (o *ServerDiagnostics) SetDiagnosticStatus(v ServerDiagnosticStatusRelationship)`

SetDiagnosticStatus sets DiagnosticStatus field to given value.

### HasDiagnosticStatus

`func (o *ServerDiagnostics) HasDiagnosticStatus() bool`

HasDiagnosticStatus returns a boolean if a field has been set.

### SetDiagnosticStatusNil

`func (o *ServerDiagnostics) SetDiagnosticStatusNil(b bool)`

 SetDiagnosticStatusNil sets the value for DiagnosticStatus to be an explicit nil

### UnsetDiagnosticStatus
`func (o *ServerDiagnostics) UnsetDiagnosticStatus()`

UnsetDiagnosticStatus ensures that no value is present for DiagnosticStatus, not even an explicit nil
### GetServer

`func (o *ServerDiagnostics) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerDiagnostics) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerDiagnostics) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *ServerDiagnostics) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *ServerDiagnostics) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *ServerDiagnostics) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


