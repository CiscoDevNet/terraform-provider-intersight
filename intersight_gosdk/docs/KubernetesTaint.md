# KubernetesTaint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Taint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Taint"]
**Effect** | Pointer to **string** | Effect of the taint is one of the following NoSchedule, PreferNoSchedule, NoExecute. | [optional] 
**Key** | Pointer to **string** | The key is any string, up to 253 characters. The key must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores. | [optional] 
**Value** | Pointer to **string** | The value is any string, up to 63 characters. The value must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores. | [optional] 

## Methods

### NewKubernetesTaint

`func NewKubernetesTaint(classId string, objectType string, ) *KubernetesTaint`

NewKubernetesTaint instantiates a new KubernetesTaint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTaintWithDefaults

`func NewKubernetesTaintWithDefaults() *KubernetesTaint`

NewKubernetesTaintWithDefaults instantiates a new KubernetesTaint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesTaint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesTaint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesTaint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesTaint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesTaint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesTaint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEffect

`func (o *KubernetesTaint) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *KubernetesTaint) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *KubernetesTaint) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *KubernetesTaint) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *KubernetesTaint) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KubernetesTaint) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KubernetesTaint) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KubernetesTaint) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *KubernetesTaint) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesTaint) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesTaint) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesTaint) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


