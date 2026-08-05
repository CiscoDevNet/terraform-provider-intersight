# FabricNodeRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SecureRouterRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SecureRouterRole"]
**NodeSlotId** | Pointer to **int64** | Chassis slot identifier where the node is physically installed. | [optional] [default to 1]
**PortPolicy** | Pointer to [**NullableFabricPortPolicyRelationship**](FabricPortPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricNodeRole

`func NewFabricNodeRole(classId string, objectType string, ) *FabricNodeRole`

NewFabricNodeRole instantiates a new FabricNodeRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricNodeRoleWithDefaults

`func NewFabricNodeRoleWithDefaults() *FabricNodeRole`

NewFabricNodeRoleWithDefaults instantiates a new FabricNodeRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricNodeRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricNodeRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricNodeRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricNodeRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricNodeRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricNodeRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodeSlotId

`func (o *FabricNodeRole) GetNodeSlotId() int64`

GetNodeSlotId returns the NodeSlotId field if non-nil, zero value otherwise.

### GetNodeSlotIdOk

`func (o *FabricNodeRole) GetNodeSlotIdOk() (*int64, bool)`

GetNodeSlotIdOk returns a tuple with the NodeSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSlotId

`func (o *FabricNodeRole) SetNodeSlotId(v int64)`

SetNodeSlotId sets NodeSlotId field to given value.

### HasNodeSlotId

`func (o *FabricNodeRole) HasNodeSlotId() bool`

HasNodeSlotId returns a boolean if a field has been set.

### GetPortPolicy

`func (o *FabricNodeRole) GetPortPolicy() FabricPortPolicyRelationship`

GetPortPolicy returns the PortPolicy field if non-nil, zero value otherwise.

### GetPortPolicyOk

`func (o *FabricNodeRole) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool)`

GetPortPolicyOk returns a tuple with the PortPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPolicy

`func (o *FabricNodeRole) SetPortPolicy(v FabricPortPolicyRelationship)`

SetPortPolicy sets PortPolicy field to given value.

### HasPortPolicy

`func (o *FabricNodeRole) HasPortPolicy() bool`

HasPortPolicy returns a boolean if a field has been set.

### SetPortPolicyNil

`func (o *FabricNodeRole) SetPortPolicyNil(b bool)`

 SetPortPolicyNil sets the value for PortPolicy to be an explicit nil

### UnsetPortPolicy
`func (o *FabricNodeRole) UnsetPortPolicy()`

UnsetPortPolicy ensures that no value is present for PortPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


