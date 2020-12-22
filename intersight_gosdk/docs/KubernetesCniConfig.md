# KubernetesCniConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.CalicoConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.CalicoConfig"]
**Registry** | Pointer to **string** | CNI Image registry location. | [optional] 
**Version** | Pointer to **string** | Container newtork interface plugin configuration. | [optional] 

## Methods

### NewKubernetesCniConfig

`func NewKubernetesCniConfig(classId string, objectType string, ) *KubernetesCniConfig`

NewKubernetesCniConfig instantiates a new KubernetesCniConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCniConfigWithDefaults

`func NewKubernetesCniConfigWithDefaults() *KubernetesCniConfig`

NewKubernetesCniConfigWithDefaults instantiates a new KubernetesCniConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesCniConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesCniConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesCniConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesCniConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesCniConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesCniConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRegistry

`func (o *KubernetesCniConfig) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *KubernetesCniConfig) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *KubernetesCniConfig) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *KubernetesCniConfig) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetVersion

`func (o *KubernetesCniConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesCniConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesCniConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesCniConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


