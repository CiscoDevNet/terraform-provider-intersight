# FabricApplianceRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ApplianceRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ApplianceRole"]
**Mode** | Pointer to **string** | Port mode to be set on the appliance port. * &#x60;trunk&#x60; - Trunk Mode Switch Port Type. * &#x60;access&#x60; - Access Mode Switch Port Type. | [optional] [default to "trunk"]
**Priority** | Pointer to **string** | The &#39;name&#39; of the System QoS Class. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [default to "Best Effort"]
**EthNetworkControlPolicy** | Pointer to [**FabricEthNetworkControlPolicyRelationship**](FabricEthNetworkControlPolicyRelationship.md) |  | [optional] 
**EthNetworkGroupPolicy** | Pointer to [**FabricEthNetworkGroupPolicyRelationship**](FabricEthNetworkGroupPolicyRelationship.md) |  | [optional] 
**FlowControlPolicy** | Pointer to [**FabricFlowControlPolicyRelationship**](FabricFlowControlPolicyRelationship.md) |  | [optional] 
**LinkControlPolicy** | Pointer to [**FabricLinkControlPolicyRelationship**](FabricLinkControlPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricApplianceRoleAllOf

`func NewFabricApplianceRoleAllOf(classId string, objectType string, ) *FabricApplianceRoleAllOf`

NewFabricApplianceRoleAllOf instantiates a new FabricApplianceRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricApplianceRoleAllOfWithDefaults

`func NewFabricApplianceRoleAllOfWithDefaults() *FabricApplianceRoleAllOf`

NewFabricApplianceRoleAllOfWithDefaults instantiates a new FabricApplianceRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricApplianceRoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricApplianceRoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricApplianceRoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricApplianceRoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricApplianceRoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricApplianceRoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMode

`func (o *FabricApplianceRoleAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FabricApplianceRoleAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FabricApplianceRoleAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FabricApplianceRoleAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPriority

`func (o *FabricApplianceRoleAllOf) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *FabricApplianceRoleAllOf) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *FabricApplianceRoleAllOf) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *FabricApplianceRoleAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEthNetworkControlPolicy

`func (o *FabricApplianceRoleAllOf) GetEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship`

GetEthNetworkControlPolicy returns the EthNetworkControlPolicy field if non-nil, zero value otherwise.

### GetEthNetworkControlPolicyOk

`func (o *FabricApplianceRoleAllOf) GetEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool)`

GetEthNetworkControlPolicyOk returns a tuple with the EthNetworkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkControlPolicy

`func (o *FabricApplianceRoleAllOf) SetEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship)`

SetEthNetworkControlPolicy sets EthNetworkControlPolicy field to given value.

### HasEthNetworkControlPolicy

`func (o *FabricApplianceRoleAllOf) HasEthNetworkControlPolicy() bool`

HasEthNetworkControlPolicy returns a boolean if a field has been set.

### GetEthNetworkGroupPolicy

`func (o *FabricApplianceRoleAllOf) GetEthNetworkGroupPolicy() FabricEthNetworkGroupPolicyRelationship`

GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field if non-nil, zero value otherwise.

### GetEthNetworkGroupPolicyOk

`func (o *FabricApplianceRoleAllOf) GetEthNetworkGroupPolicyOk() (*FabricEthNetworkGroupPolicyRelationship, bool)`

GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthNetworkGroupPolicy

`func (o *FabricApplianceRoleAllOf) SetEthNetworkGroupPolicy(v FabricEthNetworkGroupPolicyRelationship)`

SetEthNetworkGroupPolicy sets EthNetworkGroupPolicy field to given value.

### HasEthNetworkGroupPolicy

`func (o *FabricApplianceRoleAllOf) HasEthNetworkGroupPolicy() bool`

HasEthNetworkGroupPolicy returns a boolean if a field has been set.

### GetFlowControlPolicy

`func (o *FabricApplianceRoleAllOf) GetFlowControlPolicy() FabricFlowControlPolicyRelationship`

GetFlowControlPolicy returns the FlowControlPolicy field if non-nil, zero value otherwise.

### GetFlowControlPolicyOk

`func (o *FabricApplianceRoleAllOf) GetFlowControlPolicyOk() (*FabricFlowControlPolicyRelationship, bool)`

GetFlowControlPolicyOk returns a tuple with the FlowControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControlPolicy

`func (o *FabricApplianceRoleAllOf) SetFlowControlPolicy(v FabricFlowControlPolicyRelationship)`

SetFlowControlPolicy sets FlowControlPolicy field to given value.

### HasFlowControlPolicy

`func (o *FabricApplianceRoleAllOf) HasFlowControlPolicy() bool`

HasFlowControlPolicy returns a boolean if a field has been set.

### GetLinkControlPolicy

`func (o *FabricApplianceRoleAllOf) GetLinkControlPolicy() FabricLinkControlPolicyRelationship`

GetLinkControlPolicy returns the LinkControlPolicy field if non-nil, zero value otherwise.

### GetLinkControlPolicyOk

`func (o *FabricApplianceRoleAllOf) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool)`

GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkControlPolicy

`func (o *FabricApplianceRoleAllOf) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship)`

SetLinkControlPolicy sets LinkControlPolicy field to given value.

### HasLinkControlPolicy

`func (o *FabricApplianceRoleAllOf) HasLinkControlPolicy() bool`

HasLinkControlPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


