# KubernetesAddonDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AddonDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AddonDefinition"]
**ChartUrl** | Pointer to **string** | Description of the addon component. | [optional] 
**DefaultInstallStrategy** | Pointer to **string** | Default installation strategy for the release. * &#x60;None&#x60; - Unspecified install strategy. * &#x60;NoAction&#x60; - No install action performed. * &#x60;InstallOnly&#x60; - Only install in green field. No action in case of failure or removal. * &#x60;Always&#x60; - Attempt install if chart is not already installed. | [optional] [default to "None"]
**DefaultNamespace** | Pointer to **string** | Default namespace to install the release. | [optional] 
**DefaultUpgradeStrategy** | Pointer to **string** | Default upgrade strategy for the release. * &#x60;None&#x60; - Unspecified upgrade strategy. * &#x60;NoAction&#x60; - This choice enables No upgrades to be performed. * &#x60;UpgradeOnly&#x60; - Attempt upgrade if chart or overrides options change, no action on upgrade failure. * &#x60;ReinstallOnFailure&#x60; - Attempt upgrade first. Remove and install on upgrade failure. * &#x60;AlwaysReinstall&#x60; - Always remove older release and reinstall. | [optional] [default to "None"]
**Description** | Pointer to **string** | Description of the addon component. | [optional] 
**Digest** | Pointer to **string** | Digest used to verify the integrity of an addon. | [optional] 
**IconUrl** | Pointer to **string** | Icon used to represent the addon in UI. | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Name of an addon component. | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** | Version of the addon component. | [optional] 
**Catalog** | Pointer to [**KubernetesCatalogRelationship**](kubernetes.Catalog.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewKubernetesAddonDefinitionAllOf

`func NewKubernetesAddonDefinitionAllOf(classId string, objectType string, ) *KubernetesAddonDefinitionAllOf`

NewKubernetesAddonDefinitionAllOf instantiates a new KubernetesAddonDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonDefinitionAllOfWithDefaults

`func NewKubernetesAddonDefinitionAllOfWithDefaults() *KubernetesAddonDefinitionAllOf`

NewKubernetesAddonDefinitionAllOfWithDefaults instantiates a new KubernetesAddonDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChartUrl

`func (o *KubernetesAddonDefinitionAllOf) GetChartUrl() string`

GetChartUrl returns the ChartUrl field if non-nil, zero value otherwise.

### GetChartUrlOk

`func (o *KubernetesAddonDefinitionAllOf) GetChartUrlOk() (*string, bool)`

GetChartUrlOk returns a tuple with the ChartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartUrl

`func (o *KubernetesAddonDefinitionAllOf) SetChartUrl(v string)`

SetChartUrl sets ChartUrl field to given value.

### HasChartUrl

`func (o *KubernetesAddonDefinitionAllOf) HasChartUrl() bool`

HasChartUrl returns a boolean if a field has been set.

### GetDefaultInstallStrategy

`func (o *KubernetesAddonDefinitionAllOf) GetDefaultInstallStrategy() string`

GetDefaultInstallStrategy returns the DefaultInstallStrategy field if non-nil, zero value otherwise.

### GetDefaultInstallStrategyOk

`func (o *KubernetesAddonDefinitionAllOf) GetDefaultInstallStrategyOk() (*string, bool)`

GetDefaultInstallStrategyOk returns a tuple with the DefaultInstallStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstallStrategy

`func (o *KubernetesAddonDefinitionAllOf) SetDefaultInstallStrategy(v string)`

SetDefaultInstallStrategy sets DefaultInstallStrategy field to given value.

### HasDefaultInstallStrategy

`func (o *KubernetesAddonDefinitionAllOf) HasDefaultInstallStrategy() bool`

HasDefaultInstallStrategy returns a boolean if a field has been set.

### GetDefaultNamespace

`func (o *KubernetesAddonDefinitionAllOf) GetDefaultNamespace() string`

GetDefaultNamespace returns the DefaultNamespace field if non-nil, zero value otherwise.

### GetDefaultNamespaceOk

`func (o *KubernetesAddonDefinitionAllOf) GetDefaultNamespaceOk() (*string, bool)`

GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNamespace

`func (o *KubernetesAddonDefinitionAllOf) SetDefaultNamespace(v string)`

SetDefaultNamespace sets DefaultNamespace field to given value.

### HasDefaultNamespace

`func (o *KubernetesAddonDefinitionAllOf) HasDefaultNamespace() bool`

HasDefaultNamespace returns a boolean if a field has been set.

### GetDefaultUpgradeStrategy

`func (o *KubernetesAddonDefinitionAllOf) GetDefaultUpgradeStrategy() string`

GetDefaultUpgradeStrategy returns the DefaultUpgradeStrategy field if non-nil, zero value otherwise.

### GetDefaultUpgradeStrategyOk

`func (o *KubernetesAddonDefinitionAllOf) GetDefaultUpgradeStrategyOk() (*string, bool)`

GetDefaultUpgradeStrategyOk returns a tuple with the DefaultUpgradeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUpgradeStrategy

`func (o *KubernetesAddonDefinitionAllOf) SetDefaultUpgradeStrategy(v string)`

SetDefaultUpgradeStrategy sets DefaultUpgradeStrategy field to given value.

### HasDefaultUpgradeStrategy

`func (o *KubernetesAddonDefinitionAllOf) HasDefaultUpgradeStrategy() bool`

HasDefaultUpgradeStrategy returns a boolean if a field has been set.

### GetDescription

`func (o *KubernetesAddonDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesAddonDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesAddonDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesAddonDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigest

`func (o *KubernetesAddonDefinitionAllOf) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *KubernetesAddonDefinitionAllOf) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *KubernetesAddonDefinitionAllOf) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *KubernetesAddonDefinitionAllOf) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetIconUrl

`func (o *KubernetesAddonDefinitionAllOf) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *KubernetesAddonDefinitionAllOf) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *KubernetesAddonDefinitionAllOf) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *KubernetesAddonDefinitionAllOf) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesAddonDefinitionAllOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesAddonDefinitionAllOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesAddonDefinitionAllOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesAddonDefinitionAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *KubernetesAddonDefinitionAllOf) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *KubernetesAddonDefinitionAllOf) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetName

`func (o *KubernetesAddonDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesAddonDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesAddonDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesAddonDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatforms

`func (o *KubernetesAddonDefinitionAllOf) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *KubernetesAddonDefinitionAllOf) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *KubernetesAddonDefinitionAllOf) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *KubernetesAddonDefinitionAllOf) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *KubernetesAddonDefinitionAllOf) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *KubernetesAddonDefinitionAllOf) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetVersion

`func (o *KubernetesAddonDefinitionAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesAddonDefinitionAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesAddonDefinitionAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesAddonDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *KubernetesAddonDefinitionAllOf) GetCatalog() KubernetesCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *KubernetesAddonDefinitionAllOf) GetCatalogOk() (*KubernetesCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *KubernetesAddonDefinitionAllOf) SetCatalog(v KubernetesCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *KubernetesAddonDefinitionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesAddonDefinitionAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesAddonDefinitionAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesAddonDefinitionAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesAddonDefinitionAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


