# KubernetesAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Addon"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Addon"]
**AddonConfiguration** | Pointer to [**NullableKubernetesAddonConfiguration**](KubernetesAddonConfiguration.md) |  | [optional] 
**AddonPolicy** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Name** | Pointer to **string** | Name of addon to be installed on a Kubernetes cluster. | [optional] 

## Methods

### NewKubernetesAddon

`func NewKubernetesAddon(classId string, objectType string, ) *KubernetesAddon`

NewKubernetesAddon instantiates a new KubernetesAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonWithDefaults

`func NewKubernetesAddonWithDefaults() *KubernetesAddon`

NewKubernetesAddonWithDefaults instantiates a new KubernetesAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddon) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddon) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddon) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddon) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddon) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddon) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddonConfiguration

`func (o *KubernetesAddon) GetAddonConfiguration() KubernetesAddonConfiguration`

GetAddonConfiguration returns the AddonConfiguration field if non-nil, zero value otherwise.

### GetAddonConfigurationOk

`func (o *KubernetesAddon) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool)`

GetAddonConfigurationOk returns a tuple with the AddonConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonConfiguration

`func (o *KubernetesAddon) SetAddonConfiguration(v KubernetesAddonConfiguration)`

SetAddonConfiguration sets AddonConfiguration field to given value.

### HasAddonConfiguration

`func (o *KubernetesAddon) HasAddonConfiguration() bool`

HasAddonConfiguration returns a boolean if a field has been set.

### SetAddonConfigurationNil

`func (o *KubernetesAddon) SetAddonConfigurationNil(b bool)`

 SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil

### UnsetAddonConfiguration
`func (o *KubernetesAddon) UnsetAddonConfiguration()`

UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
### GetAddonPolicy

`func (o *KubernetesAddon) GetAddonPolicy() MoMoRef`

GetAddonPolicy returns the AddonPolicy field if non-nil, zero value otherwise.

### GetAddonPolicyOk

`func (o *KubernetesAddon) GetAddonPolicyOk() (*MoMoRef, bool)`

GetAddonPolicyOk returns a tuple with the AddonPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonPolicy

`func (o *KubernetesAddon) SetAddonPolicy(v MoMoRef)`

SetAddonPolicy sets AddonPolicy field to given value.

### HasAddonPolicy

`func (o *KubernetesAddon) HasAddonPolicy() bool`

HasAddonPolicy returns a boolean if a field has been set.

### GetName

`func (o *KubernetesAddon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesAddon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesAddon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesAddon) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


