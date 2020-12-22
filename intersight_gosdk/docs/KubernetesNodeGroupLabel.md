# KubernetesNodeGroupLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.NodeGroupLabel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.NodeGroupLabel"]
**Key** | Pointer to **string** | The key for a Kubernetes label for a node. | [optional] 
**Value** | Pointer to **string** | The value for a Kubernetes label for a node. | [optional] 

## Methods

### NewKubernetesNodeGroupLabel

`func NewKubernetesNodeGroupLabel(classId string, objectType string, ) *KubernetesNodeGroupLabel`

NewKubernetesNodeGroupLabel instantiates a new KubernetesNodeGroupLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeGroupLabelWithDefaults

`func NewKubernetesNodeGroupLabelWithDefaults() *KubernetesNodeGroupLabel`

NewKubernetesNodeGroupLabelWithDefaults instantiates a new KubernetesNodeGroupLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesNodeGroupLabel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesNodeGroupLabel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesNodeGroupLabel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesNodeGroupLabel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesNodeGroupLabel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesNodeGroupLabel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKey

`func (o *KubernetesNodeGroupLabel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KubernetesNodeGroupLabel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KubernetesNodeGroupLabel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KubernetesNodeGroupLabel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *KubernetesNodeGroupLabel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesNodeGroupLabel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesNodeGroupLabel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesNodeGroupLabel) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


