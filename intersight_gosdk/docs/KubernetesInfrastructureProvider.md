# KubernetesInfrastructureProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**Description** | Pointer to **string** | Description for the infrastructure provider. | [optional] 
**Name** | Pointer to **string** | Name of an infrastructure provider. | [optional] 
**NodeGroups** | Pointer to [**[]KubernetesNodeGroupProfileRelationship**](KubernetesNodeGroupProfileRelationship.md) | An array of relationships to kubernetesNodeGroupProfile resources. | [optional] 

## Methods

### NewKubernetesInfrastructureProvider

`func NewKubernetesInfrastructureProvider(classId string, objectType string, ) *KubernetesInfrastructureProvider`

NewKubernetesInfrastructureProvider instantiates a new KubernetesInfrastructureProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInfrastructureProviderWithDefaults

`func NewKubernetesInfrastructureProviderWithDefaults() *KubernetesInfrastructureProvider`

NewKubernetesInfrastructureProviderWithDefaults instantiates a new KubernetesInfrastructureProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesInfrastructureProvider) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesInfrastructureProvider) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesInfrastructureProvider) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesInfrastructureProvider) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesInfrastructureProvider) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesInfrastructureProvider) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *KubernetesInfrastructureProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesInfrastructureProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesInfrastructureProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesInfrastructureProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *KubernetesInfrastructureProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesInfrastructureProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesInfrastructureProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesInfrastructureProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeGroups

`func (o *KubernetesInfrastructureProvider) GetNodeGroups() []KubernetesNodeGroupProfileRelationship`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *KubernetesInfrastructureProvider) GetNodeGroupsOk() (*[]KubernetesNodeGroupProfileRelationship, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *KubernetesInfrastructureProvider) SetNodeGroups(v []KubernetesNodeGroupProfileRelationship)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *KubernetesInfrastructureProvider) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### SetNodeGroupsNil

`func (o *KubernetesInfrastructureProvider) SetNodeGroupsNil(b bool)`

 SetNodeGroupsNil sets the value for NodeGroups to be an explicit nil

### UnsetNodeGroups
`func (o *KubernetesInfrastructureProvider) UnsetNodeGroups()`

UnsetNodeGroups ensures that no value is present for NodeGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


