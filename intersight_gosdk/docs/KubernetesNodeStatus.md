# KubernetesNodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NodeStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NodeStatus"]
**Status** | Pointer to **string** | Statue of the node. Indicate if the node is kubernetes ready or not. | [optional] 
**Type** | Pointer to **string** | Type of the node. It can be either Master or Worker. | [optional] 

## Methods

### NewKubernetesNodeStatus

`func NewKubernetesNodeStatus(classId string, objectType string, ) *KubernetesNodeStatus`

NewKubernetesNodeStatus instantiates a new KubernetesNodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeStatusWithDefaults

`func NewKubernetesNodeStatusWithDefaults() *KubernetesNodeStatus`

NewKubernetesNodeStatusWithDefaults instantiates a new KubernetesNodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatus

`func (o *KubernetesNodeStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesNodeStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesNodeStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesNodeStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *KubernetesNodeStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodeStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodeStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesNodeStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


