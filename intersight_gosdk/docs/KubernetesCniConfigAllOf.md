# KubernetesCniConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.CalicoConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.CalicoConfig"]
**Registry** | Pointer to **string** | CNI Image registry location. | [optional] 
**Version** | Pointer to **string** | Container newtork interface plugin configuration. | [optional] 

## Methods

### NewKubernetesCniConfigAllOf

`func NewKubernetesCniConfigAllOf(classId string, objectType string, ) *KubernetesCniConfigAllOf`

NewKubernetesCniConfigAllOf instantiates a new KubernetesCniConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCniConfigAllOfWithDefaults

`func NewKubernetesCniConfigAllOfWithDefaults() *KubernetesCniConfigAllOf`

NewKubernetesCniConfigAllOfWithDefaults instantiates a new KubernetesCniConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesCniConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesCniConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesCniConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesCniConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesCniConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesCniConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRegistry

`func (o *KubernetesCniConfigAllOf) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *KubernetesCniConfigAllOf) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *KubernetesCniConfigAllOf) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *KubernetesCniConfigAllOf) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetVersion

`func (o *KubernetesCniConfigAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesCniConfigAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesCniConfigAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesCniConfigAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


