# SyslogPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "syslog.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "syslog.PolicyInventory"]
**LocalClients** | Pointer to [**[]SyslogLocalClientBase**](SyslogLocalClientBase.md) |  | [optional] 
**RemoteClients** | Pointer to [**[]SyslogRemoteClientBase**](SyslogRemoteClientBase.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewSyslogPolicyInventoryAllOf

`func NewSyslogPolicyInventoryAllOf(classId string, objectType string, ) *SyslogPolicyInventoryAllOf`

NewSyslogPolicyInventoryAllOf instantiates a new SyslogPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogPolicyInventoryAllOfWithDefaults

`func NewSyslogPolicyInventoryAllOfWithDefaults() *SyslogPolicyInventoryAllOf`

NewSyslogPolicyInventoryAllOfWithDefaults instantiates a new SyslogPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SyslogPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SyslogPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SyslogPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SyslogPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SyslogPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SyslogPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocalClients

`func (o *SyslogPolicyInventoryAllOf) GetLocalClients() []SyslogLocalClientBase`

GetLocalClients returns the LocalClients field if non-nil, zero value otherwise.

### GetLocalClientsOk

`func (o *SyslogPolicyInventoryAllOf) GetLocalClientsOk() (*[]SyslogLocalClientBase, bool)`

GetLocalClientsOk returns a tuple with the LocalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalClients

`func (o *SyslogPolicyInventoryAllOf) SetLocalClients(v []SyslogLocalClientBase)`

SetLocalClients sets LocalClients field to given value.

### HasLocalClients

`func (o *SyslogPolicyInventoryAllOf) HasLocalClients() bool`

HasLocalClients returns a boolean if a field has been set.

### SetLocalClientsNil

`func (o *SyslogPolicyInventoryAllOf) SetLocalClientsNil(b bool)`

 SetLocalClientsNil sets the value for LocalClients to be an explicit nil

### UnsetLocalClients
`func (o *SyslogPolicyInventoryAllOf) UnsetLocalClients()`

UnsetLocalClients ensures that no value is present for LocalClients, not even an explicit nil
### GetRemoteClients

`func (o *SyslogPolicyInventoryAllOf) GetRemoteClients() []SyslogRemoteClientBase`

GetRemoteClients returns the RemoteClients field if non-nil, zero value otherwise.

### GetRemoteClientsOk

`func (o *SyslogPolicyInventoryAllOf) GetRemoteClientsOk() (*[]SyslogRemoteClientBase, bool)`

GetRemoteClientsOk returns a tuple with the RemoteClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClients

`func (o *SyslogPolicyInventoryAllOf) SetRemoteClients(v []SyslogRemoteClientBase)`

SetRemoteClients sets RemoteClients field to given value.

### HasRemoteClients

`func (o *SyslogPolicyInventoryAllOf) HasRemoteClients() bool`

HasRemoteClients returns a boolean if a field has been set.

### SetRemoteClientsNil

`func (o *SyslogPolicyInventoryAllOf) SetRemoteClientsNil(b bool)`

 SetRemoteClientsNil sets the value for RemoteClients to be an explicit nil

### UnsetRemoteClients
`func (o *SyslogPolicyInventoryAllOf) UnsetRemoteClients()`

UnsetRemoteClients ensures that no value is present for RemoteClients, not even an explicit nil
### GetTargetMo

`func (o *SyslogPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *SyslogPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *SyslogPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *SyslogPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


