# WorkflowOperationTypeDecommission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.OperationTypeDecommission"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.OperationTypeDecommission"]
**ServiceItemInstance** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkflowOperationTypeDecommission

`func NewWorkflowOperationTypeDecommission(classId string, objectType string, ) *WorkflowOperationTypeDecommission`

NewWorkflowOperationTypeDecommission instantiates a new WorkflowOperationTypeDecommission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOperationTypeDecommissionWithDefaults

`func NewWorkflowOperationTypeDecommissionWithDefaults() *WorkflowOperationTypeDecommission`

NewWorkflowOperationTypeDecommissionWithDefaults instantiates a new WorkflowOperationTypeDecommission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowOperationTypeDecommission) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowOperationTypeDecommission) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowOperationTypeDecommission) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowOperationTypeDecommission) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowOperationTypeDecommission) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowOperationTypeDecommission) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetServiceItemInstance

`func (o *WorkflowOperationTypeDecommission) GetServiceItemInstance() MoMoRef`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkflowOperationTypeDecommission) GetServiceItemInstanceOk() (*MoMoRef, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkflowOperationTypeDecommission) SetServiceItemInstance(v MoMoRef)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkflowOperationTypeDecommission) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


