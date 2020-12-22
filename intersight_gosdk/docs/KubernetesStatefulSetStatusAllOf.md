# KubernetesStatefulSetStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.StatefulSetStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.StatefulSetStatus"]
**AvailableReplicas** | Pointer to **int64** | AvailableReplicas indicates the current avaliable replicas running. | [optional] 
**CollisionCount** | Pointer to **int64** | CollisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision. | [optional] 
**CurrentRevision** | Pointer to **string** | CurrentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods. | [optional] 
**ObservedGeneration** | Pointer to **int64** | ObservedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet&#39;s generation, which is updated on mutation by the API Server. | [optional] 
**ReadyReplicas** | Pointer to **int64** | ReadyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition. | [optional] 
**Replicas** | Pointer to **int64** | Number of replicas the statefulset desired to have. | [optional] 
**UpdateRevision** | Pointer to **string** | UpdateRevision, if not empty, indicates the version of the StatefulSet used to generate the pods that have yet to be updated, i.e. pod number &lt;replicas&gt; - &lt;updatedReplicas&gt;, until pod number &lt;replicas&gt;. | [optional] 
**UpdatedReplicas** | Pointer to **int64** | UpdatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision. | [optional] 

## Methods

### NewKubernetesStatefulSetStatusAllOf

`func NewKubernetesStatefulSetStatusAllOf(classId string, objectType string, ) *KubernetesStatefulSetStatusAllOf`

NewKubernetesStatefulSetStatusAllOf instantiates a new KubernetesStatefulSetStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesStatefulSetStatusAllOfWithDefaults

`func NewKubernetesStatefulSetStatusAllOfWithDefaults() *KubernetesStatefulSetStatusAllOf`

NewKubernetesStatefulSetStatusAllOfWithDefaults instantiates a new KubernetesStatefulSetStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesStatefulSetStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesStatefulSetStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesStatefulSetStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesStatefulSetStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesStatefulSetStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesStatefulSetStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailableReplicas

`func (o *KubernetesStatefulSetStatusAllOf) GetAvailableReplicas() int64`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *KubernetesStatefulSetStatusAllOf) GetAvailableReplicasOk() (*int64, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *KubernetesStatefulSetStatusAllOf) SetAvailableReplicas(v int64)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *KubernetesStatefulSetStatusAllOf) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetCollisionCount

`func (o *KubernetesStatefulSetStatusAllOf) GetCollisionCount() int64`

GetCollisionCount returns the CollisionCount field if non-nil, zero value otherwise.

### GetCollisionCountOk

`func (o *KubernetesStatefulSetStatusAllOf) GetCollisionCountOk() (*int64, bool)`

GetCollisionCountOk returns a tuple with the CollisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollisionCount

`func (o *KubernetesStatefulSetStatusAllOf) SetCollisionCount(v int64)`

SetCollisionCount sets CollisionCount field to given value.

### HasCollisionCount

`func (o *KubernetesStatefulSetStatusAllOf) HasCollisionCount() bool`

HasCollisionCount returns a boolean if a field has been set.

### GetCurrentRevision

`func (o *KubernetesStatefulSetStatusAllOf) GetCurrentRevision() string`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *KubernetesStatefulSetStatusAllOf) GetCurrentRevisionOk() (*string, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *KubernetesStatefulSetStatusAllOf) SetCurrentRevision(v string)`

SetCurrentRevision sets CurrentRevision field to given value.

### HasCurrentRevision

`func (o *KubernetesStatefulSetStatusAllOf) HasCurrentRevision() bool`

HasCurrentRevision returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *KubernetesStatefulSetStatusAllOf) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *KubernetesStatefulSetStatusAllOf) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *KubernetesStatefulSetStatusAllOf) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *KubernetesStatefulSetStatusAllOf) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *KubernetesStatefulSetStatusAllOf) GetReadyReplicas() int64`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *KubernetesStatefulSetStatusAllOf) GetReadyReplicasOk() (*int64, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *KubernetesStatefulSetStatusAllOf) SetReadyReplicas(v int64)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *KubernetesStatefulSetStatusAllOf) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *KubernetesStatefulSetStatusAllOf) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *KubernetesStatefulSetStatusAllOf) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *KubernetesStatefulSetStatusAllOf) SetReplicas(v int64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *KubernetesStatefulSetStatusAllOf) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetUpdateRevision

`func (o *KubernetesStatefulSetStatusAllOf) GetUpdateRevision() string`

GetUpdateRevision returns the UpdateRevision field if non-nil, zero value otherwise.

### GetUpdateRevisionOk

`func (o *KubernetesStatefulSetStatusAllOf) GetUpdateRevisionOk() (*string, bool)`

GetUpdateRevisionOk returns a tuple with the UpdateRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRevision

`func (o *KubernetesStatefulSetStatusAllOf) SetUpdateRevision(v string)`

SetUpdateRevision sets UpdateRevision field to given value.

### HasUpdateRevision

`func (o *KubernetesStatefulSetStatusAllOf) HasUpdateRevision() bool`

HasUpdateRevision returns a boolean if a field has been set.

### GetUpdatedReplicas

`func (o *KubernetesStatefulSetStatusAllOf) GetUpdatedReplicas() int64`

GetUpdatedReplicas returns the UpdatedReplicas field if non-nil, zero value otherwise.

### GetUpdatedReplicasOk

`func (o *KubernetesStatefulSetStatusAllOf) GetUpdatedReplicasOk() (*int64, bool)`

GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReplicas

`func (o *KubernetesStatefulSetStatusAllOf) SetUpdatedReplicas(v int64)`

SetUpdatedReplicas sets UpdatedReplicas field to given value.

### HasUpdatedReplicas

`func (o *KubernetesStatefulSetStatusAllOf) HasUpdatedReplicas() bool`

HasUpdatedReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


