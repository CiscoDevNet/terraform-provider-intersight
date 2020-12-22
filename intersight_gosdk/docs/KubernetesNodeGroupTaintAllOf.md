# KubernetesNodeGroupTaintAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NodeGroupTaint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NodeGroupTaint"]
**Effect** | Pointer to **string** | The effect to enforce when the key does not match the value. | [optional] 
**Key** | Pointer to **string** | The key for a Kubernetes taint. | [optional] 
**Value** | Pointer to **string** | If the key does not match this value, the specified effect is enforced. | [optional] 

## Methods

### NewKubernetesNodeGroupTaintAllOf

`func NewKubernetesNodeGroupTaintAllOf(classId string, objectType string, ) *KubernetesNodeGroupTaintAllOf`

NewKubernetesNodeGroupTaintAllOf instantiates a new KubernetesNodeGroupTaintAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeGroupTaintAllOfWithDefaults

`func NewKubernetesNodeGroupTaintAllOfWithDefaults() *KubernetesNodeGroupTaintAllOf`

NewKubernetesNodeGroupTaintAllOfWithDefaults instantiates a new KubernetesNodeGroupTaintAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeGroupTaintAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeGroupTaintAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeGroupTaintAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeGroupTaintAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeGroupTaintAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeGroupTaintAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEffect

`func (o *KubernetesNodeGroupTaintAllOf) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *KubernetesNodeGroupTaintAllOf) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *KubernetesNodeGroupTaintAllOf) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *KubernetesNodeGroupTaintAllOf) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *KubernetesNodeGroupTaintAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KubernetesNodeGroupTaintAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KubernetesNodeGroupTaintAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KubernetesNodeGroupTaintAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *KubernetesNodeGroupTaintAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesNodeGroupTaintAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesNodeGroupTaintAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesNodeGroupTaintAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


