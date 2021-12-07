# VirtualizationEsxiHostConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiHostConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiHostConfiguration"]
**Datacenter** | Pointer to **string** | Datacenter where host is deployed. | [optional] 

## Methods

### NewVirtualizationEsxiHostConfigurationAllOf

`func NewVirtualizationEsxiHostConfigurationAllOf(classId string, objectType string, ) *VirtualizationEsxiHostConfigurationAllOf`

NewVirtualizationEsxiHostConfigurationAllOf instantiates a new VirtualizationEsxiHostConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiHostConfigurationAllOfWithDefaults

`func NewVirtualizationEsxiHostConfigurationAllOfWithDefaults() *VirtualizationEsxiHostConfigurationAllOf`

NewVirtualizationEsxiHostConfigurationAllOfWithDefaults instantiates a new VirtualizationEsxiHostConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiHostConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiHostConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiHostConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiHostConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiHostConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiHostConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatacenter

`func (o *VirtualizationEsxiHostConfigurationAllOf) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationEsxiHostConfigurationAllOf) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationEsxiHostConfigurationAllOf) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationEsxiHostConfigurationAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


