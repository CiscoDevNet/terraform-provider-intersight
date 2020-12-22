# KubernetesDeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.DeploymentStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.DeploymentStatus"]
**AvailableReplicas** | Pointer to **int64** | Total number of available pods (ready for at least minReadySeconds) targeted by this deployment. | [optional] 
**ObservedGeneration** | Pointer to **int64** | The generation observed by the deployment controller. | [optional] 
**ReadyReplicas** | Pointer to **int64** | Total number of ready pods targeted by this deployment. | [optional] 
**Replicas** | Pointer to **int64** | Total number of non-terminated pods targeted by this deployment (their labels match the selector). | [optional] 
**UpdatedReplicas** | Pointer to **int64** | Total number of non-terminated pods targeted by this deployment that have the desired template spec. | [optional] 

## Methods

### NewKubernetesDeploymentStatus

`func NewKubernetesDeploymentStatus(classId string, objectType string, ) *KubernetesDeploymentStatus`

NewKubernetesDeploymentStatus instantiates a new KubernetesDeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDeploymentStatusWithDefaults

`func NewKubernetesDeploymentStatusWithDefaults() *KubernetesDeploymentStatus`

NewKubernetesDeploymentStatusWithDefaults instantiates a new KubernetesDeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesDeploymentStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesDeploymentStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesDeploymentStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesDeploymentStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesDeploymentStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesDeploymentStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailableReplicas

`func (o *KubernetesDeploymentStatus) GetAvailableReplicas() int64`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *KubernetesDeploymentStatus) GetAvailableReplicasOk() (*int64, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *KubernetesDeploymentStatus) SetAvailableReplicas(v int64)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *KubernetesDeploymentStatus) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *KubernetesDeploymentStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *KubernetesDeploymentStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *KubernetesDeploymentStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *KubernetesDeploymentStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *KubernetesDeploymentStatus) GetReadyReplicas() int64`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *KubernetesDeploymentStatus) GetReadyReplicasOk() (*int64, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *KubernetesDeploymentStatus) SetReadyReplicas(v int64)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *KubernetesDeploymentStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *KubernetesDeploymentStatus) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *KubernetesDeploymentStatus) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *KubernetesDeploymentStatus) SetReplicas(v int64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *KubernetesDeploymentStatus) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetUpdatedReplicas

`func (o *KubernetesDeploymentStatus) GetUpdatedReplicas() int64`

GetUpdatedReplicas returns the UpdatedReplicas field if non-nil, zero value otherwise.

### GetUpdatedReplicasOk

`func (o *KubernetesDeploymentStatus) GetUpdatedReplicasOk() (*int64, bool)`

GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReplicas

`func (o *KubernetesDeploymentStatus) SetUpdatedReplicas(v int64)`

SetUpdatedReplicas sets UpdatedReplicas field to given value.

### HasUpdatedReplicas

`func (o *KubernetesDeploymentStatus) HasUpdatedReplicas() bool`

HasUpdatedReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


