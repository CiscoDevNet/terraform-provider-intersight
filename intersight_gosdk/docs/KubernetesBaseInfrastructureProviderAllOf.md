# KubernetesBaseInfrastructureProviderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.VirtualMachineInfrastructureProvider"]
**Description** | Pointer to **string** | Description for the infrastructure provider. | [optional] 
**Name** | Pointer to **string** | Name of an infrastructure provider. | [optional] 
**NodeGroup** | Pointer to [**KubernetesNodeGroupProfileRelationship**](KubernetesNodeGroupProfileRelationship.md) |  | [optional] 

## Methods

### NewKubernetesBaseInfrastructureProviderAllOf

`func NewKubernetesBaseInfrastructureProviderAllOf(classId string, objectType string, ) *KubernetesBaseInfrastructureProviderAllOf`

NewKubernetesBaseInfrastructureProviderAllOf instantiates a new KubernetesBaseInfrastructureProviderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesBaseInfrastructureProviderAllOfWithDefaults

`func NewKubernetesBaseInfrastructureProviderAllOfWithDefaults() *KubernetesBaseInfrastructureProviderAllOf`

NewKubernetesBaseInfrastructureProviderAllOfWithDefaults instantiates a new KubernetesBaseInfrastructureProviderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesBaseInfrastructureProviderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesBaseInfrastructureProviderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesBaseInfrastructureProviderAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesBaseInfrastructureProviderAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesBaseInfrastructureProviderAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesBaseInfrastructureProviderAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeGroup

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetNodeGroup() KubernetesNodeGroupProfileRelationship`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *KubernetesBaseInfrastructureProviderAllOf) GetNodeGroupOk() (*KubernetesNodeGroupProfileRelationship, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *KubernetesBaseInfrastructureProviderAllOf) SetNodeGroup(v KubernetesNodeGroupProfileRelationship)`

SetNodeGroup sets NodeGroup field to given value.

### HasNodeGroup

`func (o *KubernetesBaseInfrastructureProviderAllOf) HasNodeGroup() bool`

HasNodeGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


