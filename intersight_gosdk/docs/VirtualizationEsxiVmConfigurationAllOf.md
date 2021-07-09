# VirtualizationEsxiVmConfigurationAllOf

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

### NewVirtualizationEsxiVmConfigurationAllOf

`func NewVirtualizationEsxiVmConfigurationAllOf(classId string, objectType string, ) *VirtualizationEsxiVmConfigurationAllOf`

NewVirtualizationEsxiVmConfigurationAllOf instantiates a new VirtualizationEsxiVmConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiVmConfigurationAllOfWithDefaults

`func NewVirtualizationEsxiVmConfigurationAllOfWithDefaults() *VirtualizationEsxiVmConfigurationAllOf`

NewVirtualizationEsxiVmConfigurationAllOfWithDefaults instantiates a new VirtualizationEsxiVmConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnnotation

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetCompute

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetCompute() VirtualizationEsxiVmComputeConfiguration`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetComputeOk() (*VirtualizationEsxiVmComputeConfiguration, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetCompute(v VirtualizationEsxiVmComputeConfiguration)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetComputeNil(b bool)`

 SetComputeNil sets the value for Compute to be an explicit nil

### UnsetCompute
`func (o *VirtualizationEsxiVmConfigurationAllOf) UnsetCompute()`

UnsetCompute ensures that no value is present for Compute, not even an explicit nil
### GetCustomspec

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetCustomspec() VirtualizationBaseCustomSpec`

GetCustomspec returns the Customspec field if non-nil, zero value otherwise.

### GetCustomspecOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetCustomspecOk() (*VirtualizationBaseCustomSpec, bool)`

GetCustomspecOk returns a tuple with the Customspec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomspec

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetCustomspec(v VirtualizationBaseCustomSpec)`

SetCustomspec sets Customspec field to given value.

### HasCustomspec

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasCustomspec() bool`

HasCustomspec returns a boolean if a field has been set.

### SetCustomspecNil

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetCustomspecNil(b bool)`

 SetCustomspecNil sets the value for Customspec to be an explicit nil

### UnsetCustomspec
`func (o *VirtualizationEsxiVmConfigurationAllOf) UnsetCustomspec()`

UnsetCustomspec ensures that no value is present for Customspec, not even an explicit nil
### GetDatacenter

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetFolder

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetImage

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetNetwork() VirtualizationEsxiVmNetworkConfiguration`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetNetworkOk() (*VirtualizationEsxiVmNetworkConfiguration, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetNetwork(v VirtualizationEsxiVmNetworkConfiguration)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *VirtualizationEsxiVmConfigurationAllOf) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetStorage

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetStorage() VirtualizationEsxiVmStorageConfiguration`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetStorageOk() (*VirtualizationEsxiVmStorageConfiguration, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetStorage(v VirtualizationEsxiVmStorageConfiguration)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *VirtualizationEsxiVmConfigurationAllOf) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetTemplate

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *VirtualizationEsxiVmConfigurationAllOf) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *VirtualizationEsxiVmConfigurationAllOf) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *VirtualizationEsxiVmConfigurationAllOf) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


