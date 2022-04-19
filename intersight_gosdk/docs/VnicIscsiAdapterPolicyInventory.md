# VnicIscsiAdapterPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.IscsiAdapterPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.IscsiAdapterPolicyInventory"]
**ConnectionTimeOut** | Pointer to **int64** | The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable. | [optional] [readonly] 
**DhcpTimeout** | Pointer to **int64** | The number of seconds to wait before the initiator assumes that the DHCP server is unavailable. | [optional] [readonly] 
**LunBusyRetryCount** | Pointer to **int64** | The number of times to retry the connection in case of a failure during iSCSI LUN discovery. | [optional] [readonly] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicIscsiAdapterPolicyInventory

`func NewVnicIscsiAdapterPolicyInventory(classId string, objectType string, ) *VnicIscsiAdapterPolicyInventory`

NewVnicIscsiAdapterPolicyInventory instantiates a new VnicIscsiAdapterPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiAdapterPolicyInventoryWithDefaults

`func NewVnicIscsiAdapterPolicyInventoryWithDefaults() *VnicIscsiAdapterPolicyInventory`

NewVnicIscsiAdapterPolicyInventoryWithDefaults instantiates a new VnicIscsiAdapterPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiAdapterPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiAdapterPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiAdapterPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiAdapterPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiAdapterPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiAdapterPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionTimeOut

`func (o *VnicIscsiAdapterPolicyInventory) GetConnectionTimeOut() int64`

GetConnectionTimeOut returns the ConnectionTimeOut field if non-nil, zero value otherwise.

### GetConnectionTimeOutOk

`func (o *VnicIscsiAdapterPolicyInventory) GetConnectionTimeOutOk() (*int64, bool)`

GetConnectionTimeOutOk returns a tuple with the ConnectionTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeOut

`func (o *VnicIscsiAdapterPolicyInventory) SetConnectionTimeOut(v int64)`

SetConnectionTimeOut sets ConnectionTimeOut field to given value.

### HasConnectionTimeOut

`func (o *VnicIscsiAdapterPolicyInventory) HasConnectionTimeOut() bool`

HasConnectionTimeOut returns a boolean if a field has been set.

### GetDhcpTimeout

`func (o *VnicIscsiAdapterPolicyInventory) GetDhcpTimeout() int64`

GetDhcpTimeout returns the DhcpTimeout field if non-nil, zero value otherwise.

### GetDhcpTimeoutOk

`func (o *VnicIscsiAdapterPolicyInventory) GetDhcpTimeoutOk() (*int64, bool)`

GetDhcpTimeoutOk returns a tuple with the DhcpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpTimeout

`func (o *VnicIscsiAdapterPolicyInventory) SetDhcpTimeout(v int64)`

SetDhcpTimeout sets DhcpTimeout field to given value.

### HasDhcpTimeout

`func (o *VnicIscsiAdapterPolicyInventory) HasDhcpTimeout() bool`

HasDhcpTimeout returns a boolean if a field has been set.

### GetLunBusyRetryCount

`func (o *VnicIscsiAdapterPolicyInventory) GetLunBusyRetryCount() int64`

GetLunBusyRetryCount returns the LunBusyRetryCount field if non-nil, zero value otherwise.

### GetLunBusyRetryCountOk

`func (o *VnicIscsiAdapterPolicyInventory) GetLunBusyRetryCountOk() (*int64, bool)`

GetLunBusyRetryCountOk returns a tuple with the LunBusyRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunBusyRetryCount

`func (o *VnicIscsiAdapterPolicyInventory) SetLunBusyRetryCount(v int64)`

SetLunBusyRetryCount sets LunBusyRetryCount field to given value.

### HasLunBusyRetryCount

`func (o *VnicIscsiAdapterPolicyInventory) HasLunBusyRetryCount() bool`

HasLunBusyRetryCount returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicIscsiAdapterPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicIscsiAdapterPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicIscsiAdapterPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicIscsiAdapterPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


