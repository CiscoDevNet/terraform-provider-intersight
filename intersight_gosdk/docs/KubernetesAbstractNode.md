# KubernetesAbstractNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.Node"]
**Annotations** | Pointer to **interface{}** | Kubernetes metadata annotations for a Node. | [optional] 
**Labels** | Pointer to **interface{}** | Kubernetes metadata labels for a Node. | [optional] 

## Methods

### NewKubernetesAbstractNode

`func NewKubernetesAbstractNode(classId string, objectType string, ) *KubernetesAbstractNode`

NewKubernetesAbstractNode instantiates a new KubernetesAbstractNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAbstractNodeWithDefaults

`func NewKubernetesAbstractNodeWithDefaults() *KubernetesAbstractNode`

NewKubernetesAbstractNodeWithDefaults instantiates a new KubernetesAbstractNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAbstractNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAbstractNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAbstractNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAbstractNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAbstractNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAbstractNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnnotations

`func (o *KubernetesAbstractNode) GetAnnotations() interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesAbstractNode) GetAnnotationsOk() (*interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesAbstractNode) SetAnnotations(v interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesAbstractNode) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *KubernetesAbstractNode) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *KubernetesAbstractNode) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetLabels

`func (o *KubernetesAbstractNode) GetLabels() interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesAbstractNode) GetLabelsOk() (*interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesAbstractNode) SetLabels(v interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesAbstractNode) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *KubernetesAbstractNode) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *KubernetesAbstractNode) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


