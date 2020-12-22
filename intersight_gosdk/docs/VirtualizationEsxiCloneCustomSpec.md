# VirtualizationEsxiCloneCustomSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.EsxiCloneCustomSpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.EsxiCloneCustomSpec"]
**ExtraConfig** | Pointer to **interface{}** | Specify the Extra Config specification which can be configured on virtual machine. | [optional] 
**OvaEnvSpec** | Pointer to **interface{}** | Specify the OVA Environment specification which can be configured on virtual machine. | [optional] 

## Methods

### NewVirtualizationEsxiCloneCustomSpec

`func NewVirtualizationEsxiCloneCustomSpec(classId string, objectType string, ) *VirtualizationEsxiCloneCustomSpec`

NewVirtualizationEsxiCloneCustomSpec instantiates a new VirtualizationEsxiCloneCustomSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationEsxiCloneCustomSpecWithDefaults

`func NewVirtualizationEsxiCloneCustomSpecWithDefaults() *VirtualizationEsxiCloneCustomSpec`

NewVirtualizationEsxiCloneCustomSpecWithDefaults instantiates a new VirtualizationEsxiCloneCustomSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationEsxiCloneCustomSpec) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationEsxiCloneCustomSpec) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationEsxiCloneCustomSpec) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationEsxiCloneCustomSpec) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationEsxiCloneCustomSpec) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationEsxiCloneCustomSpec) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExtraConfig

`func (o *VirtualizationEsxiCloneCustomSpec) GetExtraConfig() interface{}`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *VirtualizationEsxiCloneCustomSpec) GetExtraConfigOk() (*interface{}, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *VirtualizationEsxiCloneCustomSpec) SetExtraConfig(v interface{})`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *VirtualizationEsxiCloneCustomSpec) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *VirtualizationEsxiCloneCustomSpec) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *VirtualizationEsxiCloneCustomSpec) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetOvaEnvSpec

`func (o *VirtualizationEsxiCloneCustomSpec) GetOvaEnvSpec() interface{}`

GetOvaEnvSpec returns the OvaEnvSpec field if non-nil, zero value otherwise.

### GetOvaEnvSpecOk

`func (o *VirtualizationEsxiCloneCustomSpec) GetOvaEnvSpecOk() (*interface{}, bool)`

GetOvaEnvSpecOk returns a tuple with the OvaEnvSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvaEnvSpec

`func (o *VirtualizationEsxiCloneCustomSpec) SetOvaEnvSpec(v interface{})`

SetOvaEnvSpec sets OvaEnvSpec field to given value.

### HasOvaEnvSpec

`func (o *VirtualizationEsxiCloneCustomSpec) HasOvaEnvSpec() bool`

HasOvaEnvSpec returns a boolean if a field has been set.

### SetOvaEnvSpecNil

`func (o *VirtualizationEsxiCloneCustomSpec) SetOvaEnvSpecNil(b bool)`

 SetOvaEnvSpecNil sets the value for OvaEnvSpec to be an explicit nil

### UnsetOvaEnvSpec
`func (o *VirtualizationEsxiCloneCustomSpec) UnsetOvaEnvSpec()`

UnsetOvaEnvSpec ensures that no value is present for OvaEnvSpec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


