# KubernetesVersionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Version"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Version"]
**KubernetesVersion** | Pointer to **string** | Desired Kubernetes version. | [optional] 
**Name** | Pointer to **string** | The name of this IKS kubernetes version. | [optional] 
**BootIso** | Pointer to [**SoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 
**Catalog** | Pointer to [**KubernetesCatalogRelationship**](KubernetesCatalogRelationship.md) |  | [optional] 
**OvaImageTemplate** | Pointer to [**SoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 
**Qcow2NodeTemplate** | Pointer to [**SoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 
**Qcow2VirtualMachineTemplate** | Pointer to [**SoftwareSolutionDistributableRelationship**](SoftwareSolutionDistributableRelationship.md) |  | [optional] 

## Methods

### NewKubernetesVersionAllOf

`func NewKubernetesVersionAllOf(classId string, objectType string, ) *KubernetesVersionAllOf`

NewKubernetesVersionAllOf instantiates a new KubernetesVersionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVersionAllOfWithDefaults

`func NewKubernetesVersionAllOfWithDefaults() *KubernetesVersionAllOf`

NewKubernetesVersionAllOfWithDefaults instantiates a new KubernetesVersionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesVersionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesVersionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesVersionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesVersionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVersionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVersionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKubernetesVersion

`func (o *KubernetesVersionAllOf) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesVersionAllOf) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesVersionAllOf) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *KubernetesVersionAllOf) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetName

`func (o *KubernetesVersionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesVersionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesVersionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesVersionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBootIso

`func (o *KubernetesVersionAllOf) GetBootIso() SoftwareSolutionDistributableRelationship`

GetBootIso returns the BootIso field if non-nil, zero value otherwise.

### GetBootIsoOk

`func (o *KubernetesVersionAllOf) GetBootIsoOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetBootIsoOk returns a tuple with the BootIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootIso

`func (o *KubernetesVersionAllOf) SetBootIso(v SoftwareSolutionDistributableRelationship)`

SetBootIso sets BootIso field to given value.

### HasBootIso

`func (o *KubernetesVersionAllOf) HasBootIso() bool`

HasBootIso returns a boolean if a field has been set.

### GetCatalog

`func (o *KubernetesVersionAllOf) GetCatalog() KubernetesCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *KubernetesVersionAllOf) GetCatalogOk() (*KubernetesCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *KubernetesVersionAllOf) SetCatalog(v KubernetesCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *KubernetesVersionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetOvaImageTemplate

`func (o *KubernetesVersionAllOf) GetOvaImageTemplate() SoftwareSolutionDistributableRelationship`

GetOvaImageTemplate returns the OvaImageTemplate field if non-nil, zero value otherwise.

### GetOvaImageTemplateOk

`func (o *KubernetesVersionAllOf) GetOvaImageTemplateOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetOvaImageTemplateOk returns a tuple with the OvaImageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvaImageTemplate

`func (o *KubernetesVersionAllOf) SetOvaImageTemplate(v SoftwareSolutionDistributableRelationship)`

SetOvaImageTemplate sets OvaImageTemplate field to given value.

### HasOvaImageTemplate

`func (o *KubernetesVersionAllOf) HasOvaImageTemplate() bool`

HasOvaImageTemplate returns a boolean if a field has been set.

### GetQcow2NodeTemplate

`func (o *KubernetesVersionAllOf) GetQcow2NodeTemplate() SoftwareSolutionDistributableRelationship`

GetQcow2NodeTemplate returns the Qcow2NodeTemplate field if non-nil, zero value otherwise.

### GetQcow2NodeTemplateOk

`func (o *KubernetesVersionAllOf) GetQcow2NodeTemplateOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetQcow2NodeTemplateOk returns a tuple with the Qcow2NodeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQcow2NodeTemplate

`func (o *KubernetesVersionAllOf) SetQcow2NodeTemplate(v SoftwareSolutionDistributableRelationship)`

SetQcow2NodeTemplate sets Qcow2NodeTemplate field to given value.

### HasQcow2NodeTemplate

`func (o *KubernetesVersionAllOf) HasQcow2NodeTemplate() bool`

HasQcow2NodeTemplate returns a boolean if a field has been set.

### GetQcow2VirtualMachineTemplate

`func (o *KubernetesVersionAllOf) GetQcow2VirtualMachineTemplate() SoftwareSolutionDistributableRelationship`

GetQcow2VirtualMachineTemplate returns the Qcow2VirtualMachineTemplate field if non-nil, zero value otherwise.

### GetQcow2VirtualMachineTemplateOk

`func (o *KubernetesVersionAllOf) GetQcow2VirtualMachineTemplateOk() (*SoftwareSolutionDistributableRelationship, bool)`

GetQcow2VirtualMachineTemplateOk returns a tuple with the Qcow2VirtualMachineTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQcow2VirtualMachineTemplate

`func (o *KubernetesVersionAllOf) SetQcow2VirtualMachineTemplate(v SoftwareSolutionDistributableRelationship)`

SetQcow2VirtualMachineTemplate sets Qcow2VirtualMachineTemplate field to given value.

### HasQcow2VirtualMachineTemplate

`func (o *KubernetesVersionAllOf) HasQcow2VirtualMachineTemplate() bool`

HasQcow2VirtualMachineTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


