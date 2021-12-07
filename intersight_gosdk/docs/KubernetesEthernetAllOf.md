# KubernetesEthernetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Ethernet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Ethernet"]
**Matcher** | Pointer to [**NullableKubernetesEthernetMatcher**](KubernetesEthernetMatcher.md) |  | [optional] 
**ProviderName** | Pointer to **string** | If the infrastructure network is selectable, this indicates which network to attach to the port. | [optional] 

## Methods

### NewKubernetesEthernetAllOf

`func NewKubernetesEthernetAllOf(classId string, objectType string, ) *KubernetesEthernetAllOf`

NewKubernetesEthernetAllOf instantiates a new KubernetesEthernetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEthernetAllOfWithDefaults

`func NewKubernetesEthernetAllOfWithDefaults() *KubernetesEthernetAllOf`

NewKubernetesEthernetAllOfWithDefaults instantiates a new KubernetesEthernetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesEthernetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesEthernetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesEthernetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesEthernetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesEthernetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesEthernetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMatcher

`func (o *KubernetesEthernetAllOf) GetMatcher() KubernetesEthernetMatcher`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *KubernetesEthernetAllOf) GetMatcherOk() (*KubernetesEthernetMatcher, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *KubernetesEthernetAllOf) SetMatcher(v KubernetesEthernetMatcher)`

SetMatcher sets Matcher field to given value.

### HasMatcher

`func (o *KubernetesEthernetAllOf) HasMatcher() bool`

HasMatcher returns a boolean if a field has been set.

### SetMatcherNil

`func (o *KubernetesEthernetAllOf) SetMatcherNil(b bool)`

 SetMatcherNil sets the value for Matcher to be an explicit nil

### UnsetMatcher
`func (o *KubernetesEthernetAllOf) UnsetMatcher()`

UnsetMatcher ensures that no value is present for Matcher, not even an explicit nil
### GetProviderName

`func (o *KubernetesEthernetAllOf) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *KubernetesEthernetAllOf) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *KubernetesEthernetAllOf) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *KubernetesEthernetAllOf) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


