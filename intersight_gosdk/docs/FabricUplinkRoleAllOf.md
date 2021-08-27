# FabricUplinkRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.UplinkRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.UplinkRole"]
**FlowControlPolicy** | Pointer to [**FabricFlowControlPolicyRelationship**](FabricFlowControlPolicyRelationship.md) |  | [optional] 
**LinkControlPolicy** | Pointer to [**FabricLinkControlPolicyRelationship**](FabricLinkControlPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricUplinkRoleAllOf

`func NewFabricUplinkRoleAllOf(classId string, objectType string, ) *FabricUplinkRoleAllOf`

NewFabricUplinkRoleAllOf instantiates a new FabricUplinkRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricUplinkRoleAllOfWithDefaults

`func NewFabricUplinkRoleAllOfWithDefaults() *FabricUplinkRoleAllOf`

NewFabricUplinkRoleAllOfWithDefaults instantiates a new FabricUplinkRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricUplinkRoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricUplinkRoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricUplinkRoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricUplinkRoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricUplinkRoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricUplinkRoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFlowControlPolicy

`func (o *FabricUplinkRoleAllOf) GetFlowControlPolicy() FabricFlowControlPolicyRelationship`

GetFlowControlPolicy returns the FlowControlPolicy field if non-nil, zero value otherwise.

### GetFlowControlPolicyOk

`func (o *FabricUplinkRoleAllOf) GetFlowControlPolicyOk() (*FabricFlowControlPolicyRelationship, bool)`

GetFlowControlPolicyOk returns a tuple with the FlowControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlPolicy

`func (o *FabricUplinkRoleAllOf) SetFlowControlPolicy(v FabricFlowControlPolicyRelationship)`

SetFlowControlPolicy sets FlowControlPolicy field to given value.

### HasFlowControlPolicy

`func (o *FabricUplinkRoleAllOf) HasFlowControlPolicy() bool`

HasFlowControlPolicy returns a boolean if a field has been set.

### GetLinkControlPolicy

`func (o *FabricUplinkRoleAllOf) GetLinkControlPolicy() FabricLinkControlPolicyRelationship`

GetLinkControlPolicy returns the LinkControlPolicy field if non-nil, zero value otherwise.

### GetLinkControlPolicyOk

`func (o *FabricUplinkRoleAllOf) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool)`

GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkControlPolicy

`func (o *FabricUplinkRoleAllOf) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship)`

SetLinkControlPolicy sets LinkControlPolicy field to given value.

### HasLinkControlPolicy

`func (o *FabricUplinkRoleAllOf) HasLinkControlPolicy() bool`

HasLinkControlPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


