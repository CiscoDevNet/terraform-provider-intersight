# KubernetesConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Configuration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Configuration"]
**KubeConfig** | Pointer to **string** | Kubernetes configuration to connect to the cluster as an system administrator. | [optional] [readonly] 

## Methods

### NewKubernetesConfigurationAllOf

`func NewKubernetesConfigurationAllOf(classId string, objectType string, ) *KubernetesConfigurationAllOf`

NewKubernetesConfigurationAllOf instantiates a new KubernetesConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesConfigurationAllOfWithDefaults

`func NewKubernetesConfigurationAllOfWithDefaults() *KubernetesConfigurationAllOf`

NewKubernetesConfigurationAllOfWithDefaults instantiates a new KubernetesConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKubeConfig

`func (o *KubernetesConfigurationAllOf) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *KubernetesConfigurationAllOf) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *KubernetesConfigurationAllOf) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *KubernetesConfigurationAllOf) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


