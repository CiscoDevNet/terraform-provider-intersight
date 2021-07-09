# VirtualizationEsxiVmConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiVmConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiVmConfiguration"]
**Annotation** | Pointer to **string** | Specify annotation (optional) for the virtual machine. | [optional] 
**Compute** | Pointer to [**NullableVirtualizationEsxiVmComputeConfiguration**](virtualization.EsxiVmComputeConfiguration.md) |  | [optional] 
**Customspec** | Pointer to [**NullableVirtualizationBaseCustomSpec**](virtualization.BaseCustomSpec.md) |  | [optional] 
**Datacenter** | Pointer to **string** | Datacenter where virtual machine is deployed. | [optional] 
**Folder** | Pointer to **string** | Folder where virtual machine is deployed. | [optional] 
**Image** | Pointer to **string** | Image path of OVA (The image can be from any location). | [optional] 
**InventoryPath** | Pointer to **string** | The full inventory path as reported by virtual center (vCenter). Used by some of the operations to uniquely identify the VM. Inventory path is set internally based on notifications from the inventory service or some other internal channels. | [optional] [readonly] 
**Network** | Pointer to [**NullableVirtualizationEsxiVmNetworkConfiguration**](virtualization.EsxiVmNetworkConfiguration.md) |  | [optional] 
**Storage** | Pointer to [**NullableVirtualizationEsxiVmStorageConfiguration**](virtualization.EsxiVmStorageConfiguration.md) |  | [optional] 
**Template** | Pointer to **string** | Template to be used for clone. | [optional] 

## Methods

### NewVirtualizationEsxiVmConfiguration

`func NewVirtualizationEsxiVmConfiguration(classId string, objectType string, ) *VirtualizationEsxiVmConfiguration`

NewVirtualizationEsxiVmConfiguration instantiates a new VirtualizationEsxiVmConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiVmConfigurationWithDefaults

`func NewVirtualizationEsxiVmConfigurationWithDefaults() *VirtualizationEsxiVmConfiguration`

NewVirtualizationEsxiVmConfigurationWithDefaults instantiates a new VirtualizationEsxiVmConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiVmConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiVmConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiVmConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiVmConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiVmConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiVmConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnnotation

`func (o *VirtualizationEsxiVmConfiguration) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *VirtualizationEsxiVmConfiguration) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *VirtualizationEsxiVmConfiguration) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *VirtualizationEsxiVmConfiguration) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetCompute

`func (o *VirtualizationEsxiVmConfiguration) GetCompute() VirtualizationEsxiVmComputeConfiguration`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *VirtualizationEsxiVmConfiguration) GetComputeOk() (*VirtualizationEsxiVmComputeConfiguration, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *VirtualizationEsxiVmConfiguration) SetCompute(v VirtualizationEsxiVmComputeConfiguration)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *VirtualizationEsxiVmConfiguration) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *VirtualizationEsxiVmConfiguration) SetComputeNil(b bool)`

 SetComputeNil sets the value for Compute to be an explicit nil

### UnsetCompute
`func (o *VirtualizationEsxiVmConfiguration) UnsetCompute()`

UnsetCompute ensures that no value is present for Compute, not even an explicit nil
### GetCustomspec

`func (o *VirtualizationEsxiVmConfiguration) GetCustomspec() VirtualizationBaseCustomSpec`

GetCustomspec returns the Customspec field if non-nil, zero value otherwise.

### GetCustomspecOk

`func (o *VirtualizationEsxiVmConfiguration) GetCustomspecOk() (*VirtualizationBaseCustomSpec, bool)`

GetCustomspecOk returns a tuple with the Customspec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomspec

`func (o *VirtualizationEsxiVmConfiguration) SetCustomspec(v VirtualizationBaseCustomSpec)`

SetCustomspec sets Customspec field to given value.

### HasCustomspec

`func (o *VirtualizationEsxiVmConfiguration) HasCustomspec() bool`

HasCustomspec returns a boolean if a field has been set.

### SetCustomspecNil

`func (o *VirtualizationEsxiVmConfiguration) SetCustomspecNil(b bool)`

 SetCustomspecNil sets the value for Customspec to be an explicit nil

### UnsetCustomspec
`func (o *VirtualizationEsxiVmConfiguration) UnsetCustomspec()`

UnsetCustomspec ensures that no value is present for Customspec, not even an explicit nil
### GetDatacenter

`func (o *VirtualizationEsxiVmConfiguration) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationEsxiVmConfiguration) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationEsxiVmConfiguration) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationEsxiVmConfiguration) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetFolder

`func (o *VirtualizationEsxiVmConfiguration) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VirtualizationEsxiVmConfiguration) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VirtualizationEsxiVmConfiguration) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VirtualizationEsxiVmConfiguration) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetImage

`func (o *VirtualizationEsxiVmConfiguration) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VirtualizationEsxiVmConfiguration) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VirtualizationEsxiVmConfiguration) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VirtualizationEsxiVmConfiguration) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationEsxiVmConfiguration) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationEsxiVmConfiguration) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationEsxiVmConfiguration) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationEsxiVmConfiguration) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationEsxiVmConfiguration) GetNetwork() VirtualizationEsxiVmNetworkConfiguration`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationEsxiVmConfiguration) GetNetworkOk() (*VirtualizationEsxiVmNetworkConfiguration, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationEsxiVmConfiguration) SetNetwork(v VirtualizationEsxiVmNetworkConfiguration)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationEsxiVmConfiguration) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *VirtualizationEsxiVmConfiguration) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *VirtualizationEsxiVmConfiguration) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetStorage

`func (o *VirtualizationEsxiVmConfiguration) GetStorage() VirtualizationEsxiVmStorageConfiguration`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VirtualizationEsxiVmConfiguration) GetStorageOk() (*VirtualizationEsxiVmStorageConfiguration, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VirtualizationEsxiVmConfiguration) SetStorage(v VirtualizationEsxiVmStorageConfiguration)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *VirtualizationEsxiVmConfiguration) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *VirtualizationEsxiVmConfiguration) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *VirtualizationEsxiVmConfiguration) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetTemplate

`func (o *VirtualizationEsxiVmConfiguration) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *VirtualizationEsxiVmConfiguration) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *VirtualizationEsxiVmConfiguration) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *VirtualizationEsxiVmConfiguration) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


