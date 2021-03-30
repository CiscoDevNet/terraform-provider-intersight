# KubernetesAddonConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AddonConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AddonConfiguration"]
**InstallStrategy** | Pointer to **string** | Addon install strategy to determine whether an addon is installed if not present. * &#x60;None&#x60; - Unspecified install strategy. * &#x60;NoAction&#x60; - No install action performed. * &#x60;InstallOnly&#x60; - Only install in green field. No action in case of failure or removal. * &#x60;Always&#x60; - Attempt install if chart is not already installed. | [optional] [default to "None"]
**OverrideSets** | Pointer to [**[]KubernetesKeyValue**](KubernetesKeyValue.md) |  | [optional] 
**Overrides** | Pointer to **string** | Properties that can be overridden for an addon. | [optional] 
**ReleaseName** | Pointer to **string** | Name for the helm release. | [optional] 
**ReleaseNamespace** | Pointer to **string** | Namespace for the helm release. | [optional] 
**UpgradeStrategy** | Pointer to **string** | Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade. * &#x60;None&#x60; - Unspecified upgrade strategy. * &#x60;NoAction&#x60; - This choice enables No upgrades to be performed. * &#x60;UpgradeOnly&#x60; - Attempt upgrade if chart or overrides options change, no action on upgrade failure. * &#x60;ReinstallOnFailure&#x60; - Attempt upgrade first. Remove and install on upgrade failure. * &#x60;AlwaysReinstall&#x60; - Always remove older release and reinstall. | [optional] [default to "None"]

## Methods

### NewKubernetesAddonConfiguration

`func NewKubernetesAddonConfiguration(classId string, objectType string, ) *KubernetesAddonConfiguration`

NewKubernetesAddonConfiguration instantiates a new KubernetesAddonConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonConfigurationWithDefaults

`func NewKubernetesAddonConfigurationWithDefaults() *KubernetesAddonConfiguration`

NewKubernetesAddonConfigurationWithDefaults instantiates a new KubernetesAddonConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInstallStrategy

`func (o *KubernetesAddonConfiguration) GetInstallStrategy() string`

GetInstallStrategy returns the InstallStrategy field if non-nil, zero value otherwise.

### GetInstallStrategyOk

`func (o *KubernetesAddonConfiguration) GetInstallStrategyOk() (*string, bool)`

GetInstallStrategyOk returns a tuple with the InstallStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStrategy

`func (o *KubernetesAddonConfiguration) SetInstallStrategy(v string)`

SetInstallStrategy sets InstallStrategy field to given value.

### HasInstallStrategy

`func (o *KubernetesAddonConfiguration) HasInstallStrategy() bool`

HasInstallStrategy returns a boolean if a field has been set.

### GetOverrideSets

`func (o *KubernetesAddonConfiguration) GetOverrideSets() []KubernetesKeyValue`

GetOverrideSets returns the OverrideSets field if non-nil, zero value otherwise.

### GetOverrideSetsOk

`func (o *KubernetesAddonConfiguration) GetOverrideSetsOk() (*[]KubernetesKeyValue, bool)`

GetOverrideSetsOk returns a tuple with the OverrideSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSets

`func (o *KubernetesAddonConfiguration) SetOverrideSets(v []KubernetesKeyValue)`

SetOverrideSets sets OverrideSets field to given value.

### HasOverrideSets

`func (o *KubernetesAddonConfiguration) HasOverrideSets() bool`

HasOverrideSets returns a boolean if a field has been set.

### SetOverrideSetsNil

`func (o *KubernetesAddonConfiguration) SetOverrideSetsNil(b bool)`

 SetOverrideSetsNil sets the value for OverrideSets to be an explicit nil

### UnsetOverrideSets
`func (o *KubernetesAddonConfiguration) UnsetOverrideSets()`

UnsetOverrideSets ensures that no value is present for OverrideSets, not even an explicit nil
### GetOverrides

`func (o *KubernetesAddonConfiguration) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *KubernetesAddonConfiguration) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *KubernetesAddonConfiguration) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *KubernetesAddonConfiguration) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetReleaseName

`func (o *KubernetesAddonConfiguration) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *KubernetesAddonConfiguration) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *KubernetesAddonConfiguration) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *KubernetesAddonConfiguration) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetReleaseNamespace

`func (o *KubernetesAddonConfiguration) GetReleaseNamespace() string`

GetReleaseNamespace returns the ReleaseNamespace field if non-nil, zero value otherwise.

### GetReleaseNamespaceOk

`func (o *KubernetesAddonConfiguration) GetReleaseNamespaceOk() (*string, bool)`

GetReleaseNamespaceOk returns a tuple with the ReleaseNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNamespace

`func (o *KubernetesAddonConfiguration) SetReleaseNamespace(v string)`

SetReleaseNamespace sets ReleaseNamespace field to given value.

### HasReleaseNamespace

`func (o *KubernetesAddonConfiguration) HasReleaseNamespace() bool`

HasReleaseNamespace returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *KubernetesAddonConfiguration) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *KubernetesAddonConfiguration) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *KubernetesAddonConfiguration) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *KubernetesAddonConfiguration) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


