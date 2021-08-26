# FabricFlowControlPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FlowControlPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FlowControlPolicy"]
**PriorityFlowControlMode** | Pointer to **string** | Configure PFC on a per-port basis to enable the no-drop behavior for the CoS as defined by the active network qos policy. * &#x60;auto&#x60; - Enables the no-drop CoS values to be advertised by the DCBXP and negotiated with the peer.A successful negotiation enables PFC on the no-drop CoS.Any failures because of a mismatch in the capability of peers causes the PFC not to be enabled. * &#x60;on&#x60; - Enables PFC on the local port regardless of the capability of the peers. | [optional] [default to "auto"]
**ReceiveDirection** | Pointer to **string** | Link-level Flow Control configured in the receive direction. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**SendDirection** | Pointer to **string** | Link-level Flow Control configured in the send direction. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricFlowControlPolicyAllOf

`func NewFabricFlowControlPolicyAllOf(classId string, objectType string, ) *FabricFlowControlPolicyAllOf`

NewFabricFlowControlPolicyAllOf instantiates a new FabricFlowControlPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFlowControlPolicyAllOfWithDefaults

`func NewFabricFlowControlPolicyAllOfWithDefaults() *FabricFlowControlPolicyAllOf`

NewFabricFlowControlPolicyAllOfWithDefaults instantiates a new FabricFlowControlPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFlowControlPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFlowControlPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFlowControlPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFlowControlPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFlowControlPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFlowControlPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPriorityFlowControlMode

`func (o *FabricFlowControlPolicyAllOf) GetPriorityFlowControlMode() string`

GetPriorityFlowControlMode returns the PriorityFlowControlMode field if non-nil, zero value otherwise.

### GetPriorityFlowControlModeOk

`func (o *FabricFlowControlPolicyAllOf) GetPriorityFlowControlModeOk() (*string, bool)`

GetPriorityFlowControlModeOk returns a tuple with the PriorityFlowControlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityFlowControlMode

`func (o *FabricFlowControlPolicyAllOf) SetPriorityFlowControlMode(v string)`

SetPriorityFlowControlMode sets PriorityFlowControlMode field to given value.

### HasPriorityFlowControlMode

`func (o *FabricFlowControlPolicyAllOf) HasPriorityFlowControlMode() bool`

HasPriorityFlowControlMode returns a boolean if a field has been set.

### GetReceiveDirection

`func (o *FabricFlowControlPolicyAllOf) GetReceiveDirection() string`

GetReceiveDirection returns the ReceiveDirection field if non-nil, zero value otherwise.

### GetReceiveDirectionOk

`func (o *FabricFlowControlPolicyAllOf) GetReceiveDirectionOk() (*string, bool)`

GetReceiveDirectionOk returns a tuple with the ReceiveDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveDirection

`func (o *FabricFlowControlPolicyAllOf) SetReceiveDirection(v string)`

SetReceiveDirection sets ReceiveDirection field to given value.

### HasReceiveDirection

`func (o *FabricFlowControlPolicyAllOf) HasReceiveDirection() bool`

HasReceiveDirection returns a boolean if a field has been set.

### GetSendDirection

`func (o *FabricFlowControlPolicyAllOf) GetSendDirection() string`

GetSendDirection returns the SendDirection field if non-nil, zero value otherwise.

### GetSendDirectionOk

`func (o *FabricFlowControlPolicyAllOf) GetSendDirectionOk() (*string, bool)`

GetSendDirectionOk returns a tuple with the SendDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDirection

`func (o *FabricFlowControlPolicyAllOf) SetSendDirection(v string)`

SetSendDirection sets SendDirection field to given value.

### HasSendDirection

`func (o *FabricFlowControlPolicyAllOf) HasSendDirection() bool`

HasSendDirection returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricFlowControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricFlowControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricFlowControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricFlowControlPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


