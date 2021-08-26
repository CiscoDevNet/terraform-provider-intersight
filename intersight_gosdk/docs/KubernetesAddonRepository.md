# KubernetesAddonRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AddonRepository"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AddonRepository"]
**CaCert** | Pointer to [**NullableX509Certificate**](X509Certificate.md) |  | [optional] 
**InsecureSkipVerification** | Pointer to **bool** | Allow connecting to http registries or https registries which do not have certificate signed by a well known CA. | [optional] [default to false]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**IsTokenSet** | Pointer to **bool** | Indicates whether the value of the &#39;token&#39; property has been set. | [optional] [readonly] [default to false]
**Name** | Pointer to **string** | Name of the addon repository or registry. | [optional] 
**RepositoryUrl** | Pointer to **string** | URL for the repository where the addon is hosted. | [optional] 
**Username** | Pointer to **string** | Username to authenticate to the addon registry. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesAddonRepository

`func NewKubernetesAddonRepository(classId string, objectType string, ) *KubernetesAddonRepository`

NewKubernetesAddonRepository instantiates a new KubernetesAddonRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonRepositoryWithDefaults

`func NewKubernetesAddonRepositoryWithDefaults() *KubernetesAddonRepository`

NewKubernetesAddonRepositoryWithDefaults instantiates a new KubernetesAddonRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonRepository) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonRepository) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonRepository) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonRepository) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonRepository) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonRepository) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaCert

`func (o *KubernetesAddonRepository) GetCaCert() X509Certificate`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *KubernetesAddonRepository) GetCaCertOk() (*X509Certificate, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *KubernetesAddonRepository) SetCaCert(v X509Certificate)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *KubernetesAddonRepository) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *KubernetesAddonRepository) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *KubernetesAddonRepository) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetInsecureSkipVerification

`func (o *KubernetesAddonRepository) GetInsecureSkipVerification() bool`

GetInsecureSkipVerification returns the InsecureSkipVerification field if non-nil, zero value otherwise.

### GetInsecureSkipVerificationOk

`func (o *KubernetesAddonRepository) GetInsecureSkipVerificationOk() (*bool, bool)`

GetInsecureSkipVerificationOk returns a tuple with the InsecureSkipVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipVerification

`func (o *KubernetesAddonRepository) SetInsecureSkipVerification(v bool)`

SetInsecureSkipVerification sets InsecureSkipVerification field to given value.

### HasInsecureSkipVerification

`func (o *KubernetesAddonRepository) HasInsecureSkipVerification() bool`

HasInsecureSkipVerification returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *KubernetesAddonRepository) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *KubernetesAddonRepository) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *KubernetesAddonRepository) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *KubernetesAddonRepository) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetIsTokenSet

`func (o *KubernetesAddonRepository) GetIsTokenSet() bool`

GetIsTokenSet returns the IsTokenSet field if non-nil, zero value otherwise.

### GetIsTokenSetOk

`func (o *KubernetesAddonRepository) GetIsTokenSetOk() (*bool, bool)`

GetIsTokenSetOk returns a tuple with the IsTokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenSet

`func (o *KubernetesAddonRepository) SetIsTokenSet(v bool)`

SetIsTokenSet sets IsTokenSet field to given value.

### HasIsTokenSet

`func (o *KubernetesAddonRepository) HasIsTokenSet() bool`

HasIsTokenSet returns a boolean if a field has been set.

### GetName

`func (o *KubernetesAddonRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesAddonRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesAddonRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesAddonRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *KubernetesAddonRepository) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *KubernetesAddonRepository) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *KubernetesAddonRepository) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *KubernetesAddonRepository) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetUsername

`func (o *KubernetesAddonRepository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KubernetesAddonRepository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KubernetesAddonRepository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KubernetesAddonRepository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCatalog

`func (o *KubernetesAddonRepository) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *KubernetesAddonRepository) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *KubernetesAddonRepository) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *KubernetesAddonRepository) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *KubernetesAddonRepository) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesAddonRepository) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesAddonRepository) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesAddonRepository) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


