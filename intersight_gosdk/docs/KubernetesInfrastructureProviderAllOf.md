# KubernetesInfrastructureProviderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**Description** | Pointer to **string** | Description for the infrastructure provider. | [optional] 
**Name** | Pointer to **string** | Name of an infrastructure provider. | [optional] 
**NodeGroups** | Pointer to [**[]KubernetesNodeGroupProfileRelationship**](KubernetesNodeGroupProfileRelationship.md) | An array of relationships to kubernetesNodeGroupProfile resources. | [optional] 

## Methods

### NewKubernetesInfrastructureProviderAllOf

`func NewKubernetesInfrastructureProviderAllOf(classId string, objectType string, ) *KubernetesInfrastructureProviderAllOf`

NewKubernetesInfrastructureProviderAllOf instantiates a new KubernetesInfrastructureProviderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInfrastructureProviderAllOfWithDefaults

`func NewKubernetesInfrastructureProviderAllOfWithDefaults() *KubernetesInfrastructureProviderAllOf`

NewKubernetesInfrastructureProviderAllOfWithDefaults instantiates a new KubernetesInfrastructureProviderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesInfrastructureProviderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesInfrastructureProviderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesInfrastructureProviderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesInfrastructureProviderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesInfrastructureProviderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesInfrastructureProviderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *KubernetesInfrastructureProviderAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesInfrastructureProviderAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesInfrastructureProviderAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesInfrastructureProviderAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *KubernetesInfrastructureProviderAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInfrastructureProviderAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInfrastructureProviderAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesInfrastructureProviderAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeGroups

`func (o *KubernetesInfrastructureProviderAllOf) GetNodeGroups() []KubernetesNodeGroupProfileRelationship`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *KubernetesInfrastructureProviderAllOf) GetNodeGroupsOk() (*[]KubernetesNodeGroupProfileRelationship, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *KubernetesInfrastructureProviderAllOf) SetNodeGroups(v []KubernetesNodeGroupProfileRelationship)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *KubernetesInfrastructureProviderAllOf) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### SetNodeGroupsNil

`func (o *KubernetesInfrastructureProviderAllOf) SetNodeGroupsNil(b bool)`

 SetNodeGroupsNil sets the value for NodeGroups to be an explicit nil

### UnsetNodeGroups
`func (o *KubernetesInfrastructureProviderAllOf) UnsetNodeGroups()`

UnsetNodeGroups ensures that no value is present for NodeGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


