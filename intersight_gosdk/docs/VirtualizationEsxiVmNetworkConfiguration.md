# VirtualizationEsxiVmNetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiVmNetworkConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiVmNetworkConfiguration"]
**Interfaces** | Pointer to [**[]VirtualizationNetworkInterface**](VirtualizationNetworkInterface.md) |  | [optional] 

## Methods

### NewVirtualizationEsxiVmNetworkConfiguration

`func NewVirtualizationEsxiVmNetworkConfiguration(classId string, objectType string, ) *VirtualizationEsxiVmNetworkConfiguration`

NewVirtualizationEsxiVmNetworkConfiguration instantiates a new VirtualizationEsxiVmNetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiVmNetworkConfigurationWithDefaults

`func NewVirtualizationEsxiVmNetworkConfigurationWithDefaults() *VirtualizationEsxiVmNetworkConfiguration`

NewVirtualizationEsxiVmNetworkConfigurationWithDefaults instantiates a new VirtualizationEsxiVmNetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiVmNetworkConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiVmNetworkConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiVmNetworkConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiVmNetworkConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiVmNetworkConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiVmNetworkConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaces

`func (o *VirtualizationEsxiVmNetworkConfiguration) GetInterfaces() []VirtualizationNetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *VirtualizationEsxiVmNetworkConfiguration) GetInterfacesOk() (*[]VirtualizationNetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *VirtualizationEsxiVmNetworkConfiguration) SetInterfaces(v []VirtualizationNetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *VirtualizationEsxiVmNetworkConfiguration) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *VirtualizationEsxiVmNetworkConfiguration) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *VirtualizationEsxiVmNetworkConfiguration) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


