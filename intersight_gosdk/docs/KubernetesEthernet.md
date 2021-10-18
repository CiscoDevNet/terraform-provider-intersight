# KubernetesEthernet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Ethernet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Ethernet"]
**Matcher** | Pointer to [**NullableKubernetesEthernetMatcher**](KubernetesEthernetMatcher.md) |  | [optional] 

## Methods

### NewKubernetesEthernet

`func NewKubernetesEthernet(classId string, objectType string, ) *KubernetesEthernet`

NewKubernetesEthernet instantiates a new KubernetesEthernet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEthernetWithDefaults

`func NewKubernetesEthernetWithDefaults() *KubernetesEthernet`

NewKubernetesEthernetWithDefaults instantiates a new KubernetesEthernet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesEthernet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesEthernet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesEthernet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesEthernet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesEthernet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesEthernet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMatcher

`func (o *KubernetesEthernet) GetMatcher() KubernetesEthernetMatcher`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *KubernetesEthernet) GetMatcherOk() (*KubernetesEthernetMatcher, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *KubernetesEthernet) SetMatcher(v KubernetesEthernetMatcher)`

SetMatcher sets Matcher field to given value.

### HasMatcher

`func (o *KubernetesEthernet) HasMatcher() bool`

HasMatcher returns a boolean if a field has been set.

### SetMatcherNil

`func (o *KubernetesEthernet) SetMatcherNil(b bool)`

 SetMatcherNil sets the value for Matcher to be an explicit nil

### UnsetMatcher
`func (o *KubernetesEthernet) UnsetMatcher()`

UnsetMatcher ensures that no value is present for Matcher, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


