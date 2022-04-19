# VnicIscsiStaticTargetPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.IscsiStaticTargetPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.IscsiStaticTargetPolicyInventory"]
**IpAddress** | Pointer to **string** | The IPv4 address assigned to the iSCSI target. | [optional] [readonly] 
**Lun** | Pointer to [**NullableVnicLun**](VnicLun.md) |  | [optional] 
**Port** | Pointer to **int64** | The port associated with the iSCSI target. | [optional] [readonly] 
**TargetName** | Pointer to **string** | Qualified Name (IQN) or Extended Unique Identifier (EUI) name of the iSCSI target. | [optional] [readonly] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicIscsiStaticTargetPolicyInventory

`func NewVnicIscsiStaticTargetPolicyInventory(classId string, objectType string, ) *VnicIscsiStaticTargetPolicyInventory`

NewVnicIscsiStaticTargetPolicyInventory instantiates a new VnicIscsiStaticTargetPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiStaticTargetPolicyInventoryWithDefaults

`func NewVnicIscsiStaticTargetPolicyInventoryWithDefaults() *VnicIscsiStaticTargetPolicyInventory`

NewVnicIscsiStaticTargetPolicyInventoryWithDefaults instantiates a new VnicIscsiStaticTargetPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiStaticTargetPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiStaticTargetPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiStaticTargetPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiStaticTargetPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiStaticTargetPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiStaticTargetPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpAddress

`func (o *VnicIscsiStaticTargetPolicyInventory) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VnicIscsiStaticTargetPolicyInventory) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VnicIscsiStaticTargetPolicyInventory) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VnicIscsiStaticTargetPolicyInventory) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLun

`func (o *VnicIscsiStaticTargetPolicyInventory) GetLun() VnicLun`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *VnicIscsiStaticTargetPolicyInventory) GetLunOk() (*VnicLun, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *VnicIscsiStaticTargetPolicyInventory) SetLun(v VnicLun)`

SetLun sets Lun field to given value.

### HasLun

`func (o *VnicIscsiStaticTargetPolicyInventory) HasLun() bool`

HasLun returns a boolean if a field has been set.

### SetLunNil

`func (o *VnicIscsiStaticTargetPolicyInventory) SetLunNil(b bool)`

 SetLunNil sets the value for Lun to be an explicit nil

### UnsetLun
`func (o *VnicIscsiStaticTargetPolicyInventory) UnsetLun()`

UnsetLun ensures that no value is present for Lun, not even an explicit nil
### GetPort

`func (o *VnicIscsiStaticTargetPolicyInventory) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VnicIscsiStaticTargetPolicyInventory) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VnicIscsiStaticTargetPolicyInventory) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *VnicIscsiStaticTargetPolicyInventory) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTargetName

`func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *VnicIscsiStaticTargetPolicyInventory) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *VnicIscsiStaticTargetPolicyInventory) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicIscsiStaticTargetPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicIscsiStaticTargetPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


