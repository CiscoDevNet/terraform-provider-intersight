# KubernetesServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ServiceStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ServiceStatus"]
**LoadBalancer** | Pointer to [**NullableKubernetesLoadBalancer**](KubernetesLoadBalancer.md) |  | [optional] 

## Methods

### NewKubernetesServiceStatus

`func NewKubernetesServiceStatus(classId string, objectType string, ) *KubernetesServiceStatus`

NewKubernetesServiceStatus instantiates a new KubernetesServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServiceStatusWithDefaults

`func NewKubernetesServiceStatusWithDefaults() *KubernetesServiceStatus`

NewKubernetesServiceStatusWithDefaults instantiates a new KubernetesServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesServiceStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesServiceStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesServiceStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesServiceStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesServiceStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesServiceStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLoadBalancer

`func (o *KubernetesServiceStatus) GetLoadBalancer() KubernetesLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *KubernetesServiceStatus) GetLoadBalancerOk() (*KubernetesLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *KubernetesServiceStatus) SetLoadBalancer(v KubernetesLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *KubernetesServiceStatus) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### SetLoadBalancerNil

`func (o *KubernetesServiceStatus) SetLoadBalancerNil(b bool)`

 SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil

### UnsetLoadBalancer
`func (o *KubernetesServiceStatus) UnsetLoadBalancer()`

UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


