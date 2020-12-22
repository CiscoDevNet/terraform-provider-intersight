# VirtualizationEsxiVmNetworkConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiVmNetworkConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiVmNetworkConfiguration"]
**Interfaces** | Pointer to [**[]VirtualizationNetworkInterface**](VirtualizationNetworkInterface.md) |  | [optional] 

## Methods

### NewVirtualizationEsxiVmNetworkConfigurationAllOf

`func NewVirtualizationEsxiVmNetworkConfigurationAllOf(classId string, objectType string, ) *VirtualizationEsxiVmNetworkConfigurationAllOf`

NewVirtualizationEsxiVmNetworkConfigurationAllOf instantiates a new VirtualizationEsxiVmNetworkConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiVmNetworkConfigurationAllOfWithDefaults

`func NewVirtualizationEsxiVmNetworkConfigurationAllOfWithDefaults() *VirtualizationEsxiVmNetworkConfigurationAllOf`

NewVirtualizationEsxiVmNetworkConfigurationAllOfWithDefaults instantiates a new VirtualizationEsxiVmNetworkConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaces

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) GetInterfaces() []VirtualizationNetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) GetInterfacesOk() (*[]VirtualizationNetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) SetInterfaces(v []VirtualizationNetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *VirtualizationEsxiVmNetworkConfigurationAllOf) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


