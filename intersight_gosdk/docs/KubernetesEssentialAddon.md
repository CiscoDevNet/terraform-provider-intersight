# KubernetesEssentialAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.EssentialAddon"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.EssentialAddon"]
**AddonConfiguration** | Pointer to [**NullableKubernetesAddonConfiguration**](KubernetesAddonConfiguration.md) |  | [optional] 
**AddonDefinition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Name** | Pointer to **string** | Name of addon to be installed on a Kubernetes cluster. | [optional] 

## Methods

### NewKubernetesEssentialAddon

`func NewKubernetesEssentialAddon(classId string, objectType string, ) *KubernetesEssentialAddon`

NewKubernetesEssentialAddon instantiates a new KubernetesEssentialAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEssentialAddonWithDefaults

`func NewKubernetesEssentialAddonWithDefaults() *KubernetesEssentialAddon`

NewKubernetesEssentialAddonWithDefaults instantiates a new KubernetesEssentialAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesEssentialAddon) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesEssentialAddon) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesEssentialAddon) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesEssentialAddon) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesEssentialAddon) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesEssentialAddon) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddonConfiguration

`func (o *KubernetesEssentialAddon) GetAddonConfiguration() KubernetesAddonConfiguration`

GetAddonConfiguration returns the AddonConfiguration field if non-nil, zero value otherwise.

### GetAddonConfigurationOk

`func (o *KubernetesEssentialAddon) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool)`

GetAddonConfigurationOk returns a tuple with the AddonConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonConfiguration

`func (o *KubernetesEssentialAddon) SetAddonConfiguration(v KubernetesAddonConfiguration)`

SetAddonConfiguration sets AddonConfiguration field to given value.

### HasAddonConfiguration

`func (o *KubernetesEssentialAddon) HasAddonConfiguration() bool`

HasAddonConfiguration returns a boolean if a field has been set.

### SetAddonConfigurationNil

`func (o *KubernetesEssentialAddon) SetAddonConfigurationNil(b bool)`

 SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil

### UnsetAddonConfiguration
`func (o *KubernetesEssentialAddon) UnsetAddonConfiguration()`

UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
### GetAddonDefinition

`func (o *KubernetesEssentialAddon) GetAddonDefinition() MoMoRef`

GetAddonDefinition returns the AddonDefinition field if non-nil, zero value otherwise.

### GetAddonDefinitionOk

`func (o *KubernetesEssentialAddon) GetAddonDefinitionOk() (*MoMoRef, bool)`

GetAddonDefinitionOk returns a tuple with the AddonDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonDefinition

`func (o *KubernetesEssentialAddon) SetAddonDefinition(v MoMoRef)`

SetAddonDefinition sets AddonDefinition field to given value.

### HasAddonDefinition

`func (o *KubernetesEssentialAddon) HasAddonDefinition() bool`

HasAddonDefinition returns a boolean if a field has been set.

### GetName

`func (o *KubernetesEssentialAddon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesEssentialAddon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesEssentialAddon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesEssentialAddon) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


