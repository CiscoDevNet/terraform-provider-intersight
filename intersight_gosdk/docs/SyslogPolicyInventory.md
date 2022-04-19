# SyslogPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "syslog.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "syslog.PolicyInventory"]
**LocalClients** | Pointer to [**[]SyslogLocalClientBase**](SyslogLocalClientBase.md) |  | [optional] 
**RemoteClients** | Pointer to [**[]SyslogRemoteClientBase**](SyslogRemoteClientBase.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewSyslogPolicyInventory

`func NewSyslogPolicyInventory(classId string, objectType string, ) *SyslogPolicyInventory`

NewSyslogPolicyInventory instantiates a new SyslogPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogPolicyInventoryWithDefaults

`func NewSyslogPolicyInventoryWithDefaults() *SyslogPolicyInventory`

NewSyslogPolicyInventoryWithDefaults instantiates a new SyslogPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SyslogPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SyslogPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SyslogPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SyslogPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SyslogPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SyslogPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocalClients

`func (o *SyslogPolicyInventory) GetLocalClients() []SyslogLocalClientBase`

GetLocalClients returns the LocalClients field if non-nil, zero value otherwise.

### GetLocalClientsOk

`func (o *SyslogPolicyInventory) GetLocalClientsOk() (*[]SyslogLocalClientBase, bool)`

GetLocalClientsOk returns a tuple with the LocalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalClients

`func (o *SyslogPolicyInventory) SetLocalClients(v []SyslogLocalClientBase)`

SetLocalClients sets LocalClients field to given value.

### HasLocalClients

`func (o *SyslogPolicyInventory) HasLocalClients() bool`

HasLocalClients returns a boolean if a field has been set.

### SetLocalClientsNil

`func (o *SyslogPolicyInventory) SetLocalClientsNil(b bool)`

 SetLocalClientsNil sets the value for LocalClients to be an explicit nil

### UnsetLocalClients
`func (o *SyslogPolicyInventory) UnsetLocalClients()`

UnsetLocalClients ensures that no value is present for LocalClients, not even an explicit nil
### GetRemoteClients

`func (o *SyslogPolicyInventory) GetRemoteClients() []SyslogRemoteClientBase`

GetRemoteClients returns the RemoteClients field if non-nil, zero value otherwise.

### GetRemoteClientsOk

`func (o *SyslogPolicyInventory) GetRemoteClientsOk() (*[]SyslogRemoteClientBase, bool)`

GetRemoteClientsOk returns a tuple with the RemoteClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClients

`func (o *SyslogPolicyInventory) SetRemoteClients(v []SyslogRemoteClientBase)`

SetRemoteClients sets RemoteClients field to given value.

### HasRemoteClients

`func (o *SyslogPolicyInventory) HasRemoteClients() bool`

HasRemoteClients returns a boolean if a field has been set.

### SetRemoteClientsNil

`func (o *SyslogPolicyInventory) SetRemoteClientsNil(b bool)`

 SetRemoteClientsNil sets the value for RemoteClients to be an explicit nil

### UnsetRemoteClients
`func (o *SyslogPolicyInventory) UnsetRemoteClients()`

UnsetRemoteClients ensures that no value is present for RemoteClients, not even an explicit nil
### GetTargetMo

`func (o *SyslogPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *SyslogPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *SyslogPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *SyslogPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


