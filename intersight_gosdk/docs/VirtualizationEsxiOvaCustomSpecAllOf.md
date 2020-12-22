# VirtualizationEsxiOvaCustomSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiOvaCustomSpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiOvaCustomSpec"]
**OvaEnvSpec** | Pointer to **interface{}** | Specify the OVA Environment specification which can be configured on the virtual machine. | [optional] 

## Methods

### NewVirtualizationEsxiOvaCustomSpecAllOf

`func NewVirtualizationEsxiOvaCustomSpecAllOf(classId string, objectType string, ) *VirtualizationEsxiOvaCustomSpecAllOf`

NewVirtualizationEsxiOvaCustomSpecAllOf instantiates a new VirtualizationEsxiOvaCustomSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiOvaCustomSpecAllOfWithDefaults

`func NewVirtualizationEsxiOvaCustomSpecAllOfWithDefaults() *VirtualizationEsxiOvaCustomSpecAllOf`

NewVirtualizationEsxiOvaCustomSpecAllOfWithDefaults instantiates a new VirtualizationEsxiOvaCustomSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOvaEnvSpec

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetOvaEnvSpec() interface{}`

GetOvaEnvSpec returns the OvaEnvSpec field if non-nil, zero value otherwise.

### GetOvaEnvSpecOk

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) GetOvaEnvSpecOk() (*interface{}, bool)`

GetOvaEnvSpecOk returns a tuple with the OvaEnvSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvaEnvSpec

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) SetOvaEnvSpec(v interface{})`

SetOvaEnvSpec sets OvaEnvSpec field to given value.

### HasOvaEnvSpec

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) HasOvaEnvSpec() bool`

HasOvaEnvSpec returns a boolean if a field has been set.

### SetOvaEnvSpecNil

`func (o *VirtualizationEsxiOvaCustomSpecAllOf) SetOvaEnvSpecNil(b bool)`

 SetOvaEnvSpecNil sets the value for OvaEnvSpec to be an explicit nil

### UnsetOvaEnvSpec
`func (o *VirtualizationEsxiOvaCustomSpecAllOf) UnsetOvaEnvSpec()`

UnsetOvaEnvSpec ensures that no value is present for OvaEnvSpec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


