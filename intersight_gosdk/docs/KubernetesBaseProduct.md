# KubernetesBaseProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.NvidiaGpuProduct"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "kubernetes.NvidiaGpuProduct"]
**Description** | Pointer to **string** | Optional description of a product. | [optional] 
**DeviceId** | Pointer to **int64** | Device Id of a product, which is unique within a vendor. | [optional] 
**Name** | Pointer to **string** | Display Name of a product. | [optional] 
**VendorId** | Pointer to **int64** | Vendor Id of a product. Each vendor has a globally unique Id, for example 0x10DE for Nvidia. | [optional] 
**Catalog** | Pointer to [**NullableKubernetesCatalogRelationship**](KubernetesCatalogRelationship.md) |  | [optional] 

## Methods

### NewKubernetesBaseProduct

`func NewKubernetesBaseProduct(classId string, objectType string, ) *KubernetesBaseProduct`

NewKubernetesBaseProduct instantiates a new KubernetesBaseProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesBaseProductWithDefaults

`func NewKubernetesBaseProductWithDefaults() *KubernetesBaseProduct`

NewKubernetesBaseProductWithDefaults instantiates a new KubernetesBaseProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesBaseProduct) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesBaseProduct) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesBaseProduct) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesBaseProduct) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesBaseProduct) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesBaseProduct) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *KubernetesBaseProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesBaseProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesBaseProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesBaseProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceId

`func (o *KubernetesBaseProduct) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *KubernetesBaseProduct) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *KubernetesBaseProduct) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *KubernetesBaseProduct) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetName

`func (o *KubernetesBaseProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesBaseProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesBaseProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesBaseProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendorId

`func (o *KubernetesBaseProduct) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *KubernetesBaseProduct) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *KubernetesBaseProduct) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *KubernetesBaseProduct) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetCatalog

`func (o *KubernetesBaseProduct) GetCatalog() KubernetesCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *KubernetesBaseProduct) GetCatalogOk() (*KubernetesCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *KubernetesBaseProduct) SetCatalog(v KubernetesCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *KubernetesBaseProduct) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *KubernetesBaseProduct) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *KubernetesBaseProduct) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


