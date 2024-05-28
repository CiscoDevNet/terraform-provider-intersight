# KubernetesVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Version"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Version"]
**EssentialAddons** | Pointer to [**[]KubernetesAddonVersionReference**](KubernetesAddonVersionReference.md) |  | [optional] 
**HelmOperatorVersion** | Pointer to **string** | Version of helm operator to use for this kubernetes version. | [optional] 
**IksUtilityContainer** | Pointer to **string** | The iks utility container to use for the kubernetes version. | [optional] 
**KubernetesVersion** | Pointer to **string** | Desired Kubernetes version. | [optional] 
**Name** | Pointer to **string** | The name of this IKS kubernetes version. | [optional] 
**BootIso** | Pointer to [**NullableSoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 
**Catalog** | Pointer to [**NullableKubernetesCatalogRelationship**](KubernetesCatalogRelationship.md) |  | [optional] 
**OvaImageTemplate** | Pointer to [**NullableSoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 
**Qcow2NodeTemplate** | Pointer to [**NullableSoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 
**Qcow2VirtualMachineTemplate** | Pointer to [**NullableSoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 

## Methods

### NewKubernetesVersion

`func NewKubernetesVersion(classId string, objectType string, ) *KubernetesVersion`

NewKubernetesVersion instantiates a new KubernetesVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVersionWithDefaults

`func NewKubernetesVersionWithDefaults() *KubernetesVersion`

NewKubernetesVersionWithDefaults instantiates a new KubernetesVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVersion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVersion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVersion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEssentialAddons

`func (o *KubernetesVersion) GetEssentialAddons() []KubernetesAddonVersionReference`

GetEssentialAddons returns the EssentialAddons field if non-nil, zero value otherwise.

### GetEssentialAddonsOk

`func (o *KubernetesVersion) GetEssentialAddonsOk() (*[]KubernetesAddonVersionReference, bool)`

GetEssentialAddonsOk returns a tuple with the EssentialAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssentialAddons

`func (o *KubernetesVersion) SetEssentialAddons(v []KubernetesAddonVersionReference)`

SetEssentialAddons sets EssentialAddons field to given value.

### HasEssentialAddons

`func (o *KubernetesVersion) HasEssentialAddons() bool`

HasEssentialAddons returns a boolean if a field has been set.

### SetEssentialAddonsNil

`func (o *KubernetesVersion) SetEssentialAddonsNil(b bool)`

 SetEssentialAddonsNil sets the value for EssentialAddons to be an explicit nil

### UnsetEssentialAddons
`func (o *KubernetesVersion) UnsetEssentialAddons()`

UnsetEssentialAddons ensures that no value is present for EssentialAddons, not even an explicit nil
### GetHelmOperatorVersion

`func (o *KubernetesVersion) GetHelmOperatorVersion() string`

GetHelmOperatorVersion returns the HelmOperatorVersion field if non-nil, zero value otherwise.

### GetHelmOperatorVersionOk

`func (o *KubernetesVersion) GetHelmOperatorVersionOk() (*string, bool)`

GetHelmOperatorVersionOk returns a tuple with the HelmOperatorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmOperatorVersion

`func (o *KubernetesVersion) SetHelmOperatorVersion(v string)`

SetHelmOperatorVersion sets HelmOperatorVersion field to given value.

### HasHelmOperatorVersion

`func (o *KubernetesVersion) HasHelmOperatorVersion() bool`

HasHelmOperatorVersion returns a boolean if a field has been set.

### GetIksUtilityContainer

`func (o *KubernetesVersion) GetIksUtilityContainer() string`

GetIksUtilityContainer returns the IksUtilityContainer field if non-nil, zero value otherwise.

### GetIksUtilityContainerOk

`func (o *KubernetesVersion) GetIksUtilityContainerOk() (*string, bool)`

GetIksUtilityContainerOk returns a tuple with the IksUtilityContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIksUtilityContainer

`func (o *KubernetesVersion) SetIksUtilityContainer(v string)`

SetIksUtilityContainer sets IksUtilityContainer field to given value.

### HasIksUtilityContainer

`func (o *KubernetesVersion) HasIksUtilityContainer() bool`

HasIksUtilityContainer returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *KubernetesVersion) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesVersion) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesVersion) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *KubernetesVersion) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetName

`func (o *KubernetesVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBootIso

`func (o *KubernetesVersion) GetBootIso() SoftwareSolutionDistributableRelationship`

GetBootIso returns the BootIso field if non-nil, zero value otherwise.

### GetBootIsoOk

`func (o *KubernetesVersion) GetBootIsoOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetBootIsoOk returns a tuple with the BootIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootIso

`func (o *KubernetesVersion) SetBootIso(v SoftwareSolutionDistributableRelationship)`

SetBootIso sets BootIso field to given value.

### HasBootIso

`func (o *KubernetesVersion) HasBootIso() bool`

HasBootIso returns a boolean if a field has been set.

### SetBootIsoNil

`func (o *KubernetesVersion) SetBootIsoNil(b bool)`

 SetBootIsoNil sets the value for BootIso to be an explicit nil

### UnsetBootIso
`func (o *KubernetesVersion) UnsetBootIso()`

UnsetBootIso ensures that no value is present for BootIso, not even an explicit nil
### GetCatalog

`func (o *KubernetesVersion) GetCatalog() KubernetesCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *KubernetesVersion) GetCatalogOk() (*KubernetesCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *KubernetesVersion) SetCatalog(v KubernetesCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *KubernetesVersion) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *KubernetesVersion) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *KubernetesVersion) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetOvaImageTemplate

`func (o *KubernetesVersion) GetOvaImageTemplate() SoftwareSolutionDistributableRelationship`

GetOvaImageTemplate returns the OvaImageTemplate field if non-nil, zero value otherwise.

### GetOvaImageTemplateOk

`func (o *KubernetesVersion) GetOvaImageTemplateOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetOvaImageTemplateOk returns a tuple with the OvaImageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvaImageTemplate

`func (o *KubernetesVersion) SetOvaImageTemplate(v SoftwareSolutionDistributableRelationship)`

SetOvaImageTemplate sets OvaImageTemplate field to given value.

### HasOvaImageTemplate

`func (o *KubernetesVersion) HasOvaImageTemplate() bool`

HasOvaImageTemplate returns a boolean if a field has been set.

### SetOvaImageTemplateNil

`func (o *KubernetesVersion) SetOvaImageTemplateNil(b bool)`

 SetOvaImageTemplateNil sets the value for OvaImageTemplate to be an explicit nil

### UnsetOvaImageTemplate
`func (o *KubernetesVersion) UnsetOvaImageTemplate()`

UnsetOvaImageTemplate ensures that no value is present for OvaImageTemplate, not even an explicit nil
### GetQcow2NodeTemplate

`func (o *KubernetesVersion) GetQcow2NodeTemplate() SoftwareSolutionDistributableRelationship`

GetQcow2NodeTemplate returns the Qcow2NodeTemplate field if non-nil, zero value otherwise.

### GetQcow2NodeTemplateOk

`func (o *KubernetesVersion) GetQcow2NodeTemplateOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetQcow2NodeTemplateOk returns a tuple with the Qcow2NodeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQcow2NodeTemplate

`func (o *KubernetesVersion) SetQcow2NodeTemplate(v SoftwareSolutionDistributableRelationship)`

SetQcow2NodeTemplate sets Qcow2NodeTemplate field to given value.

### HasQcow2NodeTemplate

`func (o *KubernetesVersion) HasQcow2NodeTemplate() bool`

HasQcow2NodeTemplate returns a boolean if a field has been set.

### SetQcow2NodeTemplateNil

`func (o *KubernetesVersion) SetQcow2NodeTemplateNil(b bool)`

 SetQcow2NodeTemplateNil sets the value for Qcow2NodeTemplate to be an explicit nil

### UnsetQcow2NodeTemplate
`func (o *KubernetesVersion) UnsetQcow2NodeTemplate()`

UnsetQcow2NodeTemplate ensures that no value is present for Qcow2NodeTemplate, not even an explicit nil
### GetQcow2VirtualMachineTemplate

`func (o *KubernetesVersion) GetQcow2VirtualMachineTemplate() SoftwareSolutionDistributableRelationship`

GetQcow2VirtualMachineTemplate returns the Qcow2VirtualMachineTemplate field if non-nil, zero value otherwise.

### GetQcow2VirtualMachineTemplateOk

`func (o *KubernetesVersion) GetQcow2VirtualMachineTemplateOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetQcow2VirtualMachineTemplateOk returns a tuple with the Qcow2VirtualMachineTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQcow2VirtualMachineTemplate

`func (o *KubernetesVersion) SetQcow2VirtualMachineTemplate(v SoftwareSolutionDistributableRelationship)`

SetQcow2VirtualMachineTemplate sets Qcow2VirtualMachineTemplate field to given value.

### HasQcow2VirtualMachineTemplate

`func (o *KubernetesVersion) HasQcow2VirtualMachineTemplate() bool`

HasQcow2VirtualMachineTemplate returns a boolean if a field has been set.

### SetQcow2VirtualMachineTemplateNil

`func (o *KubernetesVersion) SetQcow2VirtualMachineTemplateNil(b bool)`

 SetQcow2VirtualMachineTemplateNil sets the value for Qcow2VirtualMachineTemplate to be an explicit nil

### UnsetQcow2VirtualMachineTemplate
`func (o *KubernetesVersion) UnsetQcow2VirtualMachineTemplate()`

UnsetQcow2VirtualMachineTemplate ensures that no value is present for Qcow2VirtualMachineTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


