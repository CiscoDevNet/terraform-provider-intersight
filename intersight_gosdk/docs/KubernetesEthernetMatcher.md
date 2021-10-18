# KubernetesEthernetMatcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.EthernetMatcher"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.EthernetMatcher"]
**Type** | Pointer to **string** | Which property we should use to find the ethernet interface. * &#x60;Name&#x60; - A network interface name, e.g. eth0, eno9. * &#x60;MacAddress&#x60; - A network interface Mac Address. | [optional] [default to "Name"]
**Value** | Pointer to **string** | The value to match for the property specified by type. | [optional] 

## Methods

### NewKubernetesEthernetMatcher

`func NewKubernetesEthernetMatcher(classId string, objectType string, ) *KubernetesEthernetMatcher`

NewKubernetesEthernetMatcher instantiates a new KubernetesEthernetMatcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEthernetMatcherWithDefaults

`func NewKubernetesEthernetMatcherWithDefaults() *KubernetesEthernetMatcher`

NewKubernetesEthernetMatcherWithDefaults instantiates a new KubernetesEthernetMatcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesEthernetMatcher) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesEthernetMatcher) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesEthernetMatcher) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesEthernetMatcher) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesEthernetMatcher) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesEthernetMatcher) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetType

`func (o *KubernetesEthernetMatcher) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesEthernetMatcher) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesEthernetMatcher) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesEthernetMatcher) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *KubernetesEthernetMatcher) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesEthernetMatcher) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesEthernetMatcher) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesEthernetMatcher) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


