# KubernetesAddonAllOf

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

### NewKubernetesAddonAllOf

`func NewKubernetesAddonAllOf(classId string, objectType string, ) *KubernetesAddonAllOf`

NewKubernetesAddonAllOf instantiates a new KubernetesAddonAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonAllOfWithDefaults

`func NewKubernetesAddonAllOfWithDefaults() *KubernetesAddonAllOf`

NewKubernetesAddonAllOfWithDefaults instantiates a new KubernetesAddonAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInstallStrategy

`func (o *KubernetesAddonAllOf) GetInstallStrategy() string`

GetInstallStrategy returns the InstallStrategy field if non-nil, zero value otherwise.

### GetInstallStrategyOk

`func (o *KubernetesAddonAllOf) GetInstallStrategyOk() (*string, bool)`

GetInstallStrategyOk returns a tuple with the InstallStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStrategy

`func (o *KubernetesAddonAllOf) SetInstallStrategy(v string)`

SetInstallStrategy sets InstallStrategy field to given value.

### HasInstallStrategy

`func (o *KubernetesAddonAllOf) HasInstallStrategy() bool`

HasInstallStrategy returns a boolean if a field has been set.

### GetName

`func (o *KubernetesAddonAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesAddonAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesAddonAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesAddonAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideSets

`func (o *KubernetesAddonAllOf) GetOverrideSets() []KubernetesKeyValue`

GetOverrideSets returns the OverrideSets field if non-nil, zero value otherwise.

### GetOverrideSetsOk

`func (o *KubernetesAddonAllOf) GetOverrideSetsOk() (*[]KubernetesKeyValue, bool)`

GetOverrideSetsOk returns a tuple with the OverrideSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSets

`func (o *KubernetesAddonAllOf) SetOverrideSets(v []KubernetesKeyValue)`

SetOverrideSets sets OverrideSets field to given value.

### HasOverrideSets

`func (o *KubernetesAddonAllOf) HasOverrideSets() bool`

HasOverrideSets returns a boolean if a field has been set.

### SetOverrideSetsNil

`func (o *KubernetesAddonAllOf) SetOverrideSetsNil(b bool)`

 SetOverrideSetsNil sets the value for OverrideSets to be an explicit nil

### UnsetOverrideSets
`func (o *KubernetesAddonAllOf) UnsetOverrideSets()`

UnsetOverrideSets ensures that no value is present for OverrideSets, not even an explicit nil
### GetOverrides

`func (o *KubernetesAddonAllOf) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *KubernetesAddonAllOf) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *KubernetesAddonAllOf) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *KubernetesAddonAllOf) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetUpgradeStrategy

`func (o *KubernetesAddonAllOf) GetUpgradeStrategy() string`

GetUpgradeStrategy returns the UpgradeStrategy field if non-nil, zero value otherwise.

### GetUpgradeStrategyOk

`func (o *KubernetesAddonAllOf) GetUpgradeStrategyOk() (*string, bool)`

GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStrategy

`func (o *KubernetesAddonAllOf) SetUpgradeStrategy(v string)`

SetUpgradeStrategy sets UpgradeStrategy field to given value.

### HasUpgradeStrategy

`func (o *KubernetesAddonAllOf) HasUpgradeStrategy() bool`

HasUpgradeStrategy returns a boolean if a field has been set.

### GetAddonDefinition

`func (o *KubernetesAddonAllOf) GetAddonDefinition() KubernetesAddonDefinitionRelationship`

GetAddonDefinition returns the AddonDefinition field if non-nil, zero value otherwise.

### GetAddonDefinitionOk

`func (o *KubernetesAddonAllOf) GetAddonDefinitionOk() (*KubernetesAddonDefinitionRelationship, bool)`

GetAddonDefinitionOk returns a tuple with the AddonDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonDefinition

`func (o *KubernetesAddonAllOf) SetAddonDefinition(v KubernetesAddonDefinitionRelationship)`

SetAddonDefinition sets AddonDefinition field to given value.

### HasAddonDefinition

`func (o *KubernetesAddonAllOf) HasAddonDefinition() bool`

HasAddonDefinition returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesAddonAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAddonAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAddonAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAddonAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


