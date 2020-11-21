# HyperflexVmInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmInterface"]
**Bridge** | Pointer to **string** | Virtual machine brige name of interface. | [optional] [readonly] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**MacAddress** | Pointer to **string** | Virtual machine interface mac address. | [optional] [readonly] 
**Model** | Pointer to **string** | Virtual machine interface model. | [optional] [readonly] 

## Methods

### NewHyperflexVmInterface

`func NewHyperflexVmInterface(classId string, objectType string, ) *HyperflexVmInterface`

NewHyperflexVmInterface instantiates a new HyperflexVmInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmInterfaceWithDefaults

`func NewHyperflexVmInterfaceWithDefaults() *HyperflexVmInterface`

NewHyperflexVmInterfaceWithDefaults instantiates a new HyperflexVmInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBridge

`func (o *HyperflexVmInterface) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *HyperflexVmInterface) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *HyperflexVmInterface) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *HyperflexVmInterface) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetIpAddress

`func (o *HyperflexVmInterface) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HyperflexVmInterface) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HyperflexVmInterface) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HyperflexVmInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *HyperflexVmInterface) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *HyperflexVmInterface) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMacAddress

`func (o *HyperflexVmInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HyperflexVmInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HyperflexVmInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HyperflexVmInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetModel

`func (o *HyperflexVmInterface) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HyperflexVmInterface) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HyperflexVmInterface) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HyperflexVmInterface) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


