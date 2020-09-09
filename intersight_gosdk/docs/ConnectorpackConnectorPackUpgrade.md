# ConnectorpackConnectorPackUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorPackOpType** | Pointer to **string** | The type of operation to be performed on UCS Director. * &#x60;Install&#x60; - Installs the requisite connector packs on UCS Director. * &#x60;Push&#x60; - Pushes the requisite connector packs to UCS Director. | [optional] [default to "Install"]
**UcsdInfo** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewConnectorpackConnectorPackUpgrade

`func NewConnectorpackConnectorPackUpgrade() *ConnectorpackConnectorPackUpgrade`

NewConnectorpackConnectorPackUpgrade instantiates a new ConnectorpackConnectorPackUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorpackConnectorPackUpgradeWithDefaults

`func NewConnectorpackConnectorPackUpgradeWithDefaults() *ConnectorpackConnectorPackUpgrade`

NewConnectorpackConnectorPackUpgradeWithDefaults instantiates a new ConnectorpackConnectorPackUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorPackOpType

`func (o *ConnectorpackConnectorPackUpgrade) GetConnectorPackOpType() string`

GetConnectorPackOpType returns the ConnectorPackOpType field if non-nil, zero value otherwise.

### GetConnectorPackOpTypeOk

`func (o *ConnectorpackConnectorPackUpgrade) GetConnectorPackOpTypeOk() (*string, bool)`

GetConnectorPackOpTypeOk returns a tuple with the ConnectorPackOpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPackOpType

`func (o *ConnectorpackConnectorPackUpgrade) SetConnectorPackOpType(v string)`

SetConnectorPackOpType sets ConnectorPackOpType field to given value.

### HasConnectorPackOpType

`func (o *ConnectorpackConnectorPackUpgrade) HasConnectorPackOpType() bool`

HasConnectorPackOpType returns a boolean if a field has been set.

### GetUcsdInfo

`func (o *ConnectorpackConnectorPackUpgrade) GetUcsdInfo() IaasUcsdInfoRelationship`

GetUcsdInfo returns the UcsdInfo field if non-nil, zero value otherwise.

### GetUcsdInfoOk

`func (o *ConnectorpackConnectorPackUpgrade) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool)`

GetUcsdInfoOk returns a tuple with the UcsdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsdInfo

`func (o *ConnectorpackConnectorPackUpgrade) SetUcsdInfo(v IaasUcsdInfoRelationship)`

SetUcsdInfo sets UcsdInfo field to given value.

### HasUcsdInfo

`func (o *ConnectorpackConnectorPackUpgrade) HasUcsdInfo() bool`

HasUcsdInfo returns a boolean if a field has been set.

### GetWorkflow

`func (o *ConnectorpackConnectorPackUpgrade) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ConnectorpackConnectorPackUpgrade) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ConnectorpackConnectorPackUpgrade) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ConnectorpackConnectorPackUpgrade) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


