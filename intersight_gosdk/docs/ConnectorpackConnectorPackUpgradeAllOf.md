# ConnectorpackConnectorPackUpgradeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connectorpack.ConnectorPackUpgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connectorpack.ConnectorPackUpgrade"]
**ConnectorPackOpType** | Pointer to **string** | The type of operation to be performed on UCS Director. * &#x60;Install&#x60; - Installs the requisite connector packs on UCS Director. * &#x60;Push&#x60; - Pushes the requisite connector packs to UCS Director. | [optional] [default to "Install"]
**UcsdInfo** | Pointer to [**IaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewConnectorpackConnectorPackUpgradeAllOf

`func NewConnectorpackConnectorPackUpgradeAllOf(classId string, objectType string, ) *ConnectorpackConnectorPackUpgradeAllOf`

NewConnectorpackConnectorPackUpgradeAllOf instantiates a new ConnectorpackConnectorPackUpgradeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorpackConnectorPackUpgradeAllOfWithDefaults

`func NewConnectorpackConnectorPackUpgradeAllOfWithDefaults() *ConnectorpackConnectorPackUpgradeAllOf`

NewConnectorpackConnectorPackUpgradeAllOfWithDefaults instantiates a new ConnectorpackConnectorPackUpgradeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorpackConnectorPackUpgradeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorpackConnectorPackUpgradeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectorPackOpType

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetConnectorPackOpType() string`

GetConnectorPackOpType returns the ConnectorPackOpType field if non-nil, zero value otherwise.

### GetConnectorPackOpTypeOk

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetConnectorPackOpTypeOk() (*string, bool)`

GetConnectorPackOpTypeOk returns a tuple with the ConnectorPackOpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPackOpType

`func (o *ConnectorpackConnectorPackUpgradeAllOf) SetConnectorPackOpType(v string)`

SetConnectorPackOpType sets ConnectorPackOpType field to given value.

### HasConnectorPackOpType

`func (o *ConnectorpackConnectorPackUpgradeAllOf) HasConnectorPackOpType() bool`

HasConnectorPackOpType returns a boolean if a field has been set.

### GetUcsdInfo

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetUcsdInfo() IaasUcsdInfoRelationship`

GetUcsdInfo returns the UcsdInfo field if non-nil, zero value otherwise.

### GetUcsdInfoOk

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool)`

GetUcsdInfoOk returns a tuple with the UcsdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsdInfo

`func (o *ConnectorpackConnectorPackUpgradeAllOf) SetUcsdInfo(v IaasUcsdInfoRelationship)`

SetUcsdInfo sets UcsdInfo field to given value.

### HasUcsdInfo

`func (o *ConnectorpackConnectorPackUpgradeAllOf) HasUcsdInfo() bool`

HasUcsdInfo returns a boolean if a field has been set.

### GetWorkflow

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ConnectorpackConnectorPackUpgradeAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ConnectorpackConnectorPackUpgradeAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ConnectorpackConnectorPackUpgradeAllOf) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


