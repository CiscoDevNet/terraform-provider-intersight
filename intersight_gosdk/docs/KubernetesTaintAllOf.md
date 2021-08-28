# KubernetesTaintAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Taint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Taint"]
**Effect** | Pointer to **string** | Effect of the taint is one of the following NoSchedule, PreferNoSchedule, NoExecute. | [optional] 
**Key** | Pointer to **string** | The key is any string, up to 253 characters. The key must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores. | [optional] 
**Value** | Pointer to **string** | The value is any string, up to 63 characters. The value must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores. | [optional] 

## Methods

### NewKubernetesTaintAllOf

`func NewKubernetesTaintAllOf(classId string, objectType string, ) *KubernetesTaintAllOf`

NewKubernetesTaintAllOf instantiates a new KubernetesTaintAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTaintAllOfWithDefaults

`func NewKubernetesTaintAllOfWithDefaults() *KubernetesTaintAllOf`

NewKubernetesTaintAllOfWithDefaults instantiates a new KubernetesTaintAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesTaintAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesTaintAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesTaintAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesTaintAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesTaintAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesTaintAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEffect

`func (o *KubernetesTaintAllOf) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *KubernetesTaintAllOf) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *KubernetesTaintAllOf) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *KubernetesTaintAllOf) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *KubernetesTaintAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KubernetesTaintAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KubernetesTaintAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KubernetesTaintAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *KubernetesTaintAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesTaintAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesTaintAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesTaintAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


