# KubernetesAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Addon"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Addon"]
**InstallStrategy** | Pointer to **string** | Addon install strategy to determine whether an addon is installed if not present. * &#x60;InstallOnly&#x60; - Only install in green field. No action in case of failure or removal. * &#x60;NoAction&#x60; - No install action performed. * &#x60;Always&#x60; - Attempt install if chart is not already installed. | [optional] [default to "InstallOnly"]
**Name** | Pointer to **string** | Name of addon to be installed on a Kubernetes cluster. | [optional] 
**OverrideSets** | Pointer to [**[]KubernetesKeyValue**](KubernetesKeyValue.md) |  | [optional] 
**Overrides** | Pointer to **string** | Properties that can be overridden for an addon. | [optional] 
**UpgradeStrategy** | Pointer to **string** | Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade. * &#x60;UpgradeOnly&#x60; - Attempt upgrade if chart or overrides options change, no action on upgrade failure. * &#x60;NoAction&#x60; - This choice enables No upgrades to be performed. * &#x60;ReinstallOnFailure&#x60; - Attempt upgrade first. Remove and install on upgrade failure. * &#x60;AlwaysReinstall&#x60; - Always remove older release and reinstall. | [optional] [default to "UpgradeOnly"]
**AddonDefinition** | Pointer to [**KubernetesAddonDefinitionRelationship**](kubernetes.AddonDefinition.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

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


### GetInstallStrategy

`func (o *KubernetesAddon) GetInstallStrategy() string`

GetInstallStrategy returns the InstallStrategy field if non-nil, zero value otherwise.

### GetInstallStrategyOk

`func (o *KubernetesAddon) GetInstallStrategyOk() (*string, bool)`

GetInstallStrategyOk returns a tuple with the InstallStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStrategy

`func (o *KubernetesAddon) SetInstallStrategy(v string)`

SetInstallStrategy sets InstallStrategy field to given value.

### HasInstallStrategy

`func (o *KubernetesAddon) HasInstallStrategy() bool`

HasInstallStrategy returns a boolean if a field has been set.

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

### GetOverrideSets

`func (o *KubernetesAddon) GetOverrideSets() []KubernetesKeyValue`

GetOverrideSets returns the OverrideSets field if non-nil, zero value otherwise.

### GetOverrideSetsOk

`func (o *KubernetesAddon) GetOverrideSetsOk() (*[]KubernetesKeyValue, bool)`

GetOverrideSetsOk returns a tuple with the OverrideSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSets

`func (o *KubernetesAddon) SetOverrideSets(v []KubernetesKeyValue)`

SetOverrideSets sets OverrideSets field to given value.

### HasOverrideSets

`func (o *KubernetesAddon) HasOverrideSets() bool`

HasOverrideSets returns a boolean if a field has been set.

### SetOverrideSetsNil

`func (o *KubernetesAddon) SetOverrideSetsNil(b bool)`

 SetOverrideSetsNil sets the value for OverrideSets to be an explicit nil

### UnsetOverrideSets
`func (o *KubernetesAddon) UnsetOverrideSets()`

UnsetOverrideSets ensures that no value is present for OverrideSets, not even an explicit nil
### GetOverrides

`func (o *KubernetesAddon) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *KubernetesAddon) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *KubernetesAddon) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *KubernetesAddon) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *KubernetesAddon) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *KubernetesAddon) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *KubernetesAddon) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *KubernetesAddon) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.

### GetAddonDefinition

`func (o *KubernetesAddon) GetAddonDefinition() KubernetesAddonDefinitionRelationship`

GetAddonDefinition returns the AddonDefinition field if non-nil, zero value otherwise.

### GetAddonDefinitionOk

`func (o *KubernetesAddon) GetAddonDefinitionOk() (*KubernetesAddonDefinitionRelationship, bool)`

GetAddonDefinitionOk returns a tuple with the AddonDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonDefinition

`func (o *KubernetesAddon) SetAddonDefinition(v KubernetesAddonDefinitionRelationship)`

SetAddonDefinition sets AddonDefinition field to given value.

### HasAddonDefinition

`func (o *KubernetesAddon) HasAddonDefinition() bool`

HasAddonDefinition returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesAddon) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAddon) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAddon) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAddon) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


