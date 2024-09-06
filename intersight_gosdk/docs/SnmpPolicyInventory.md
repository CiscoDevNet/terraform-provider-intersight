# SnmpPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "snmp.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "snmp.PolicyInventory"]
**AccessCommunityString** | Pointer to **string** | The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 32 characters long. | [optional] [readonly] 
**CommunityAccess** | Pointer to **string** | Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users. * &#x60;Disabled&#x60; - Blocks access to the information in the inventory tables. * &#x60;Limited&#x60; - Partial access to read the information in the inventory tables. * &#x60;Full&#x60; - Full access to read the information in the inventory tables. | [optional] [readonly] [default to "Disabled"]
**Enabled** | Pointer to **bool** | State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host. | [optional] [readonly] [default to true]
**EngineId** | Pointer to **string** | User-defined unique identification of the static engine. | [optional] [readonly] 
**SnmpPort** | Pointer to **int64** | Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). | [optional] [readonly] [default to 161]
**SnmpTraps** | Pointer to [**[]SnmpTrap**](SnmpTrap.md) |  | [optional] 
**SnmpUsers** | Pointer to [**[]SnmpUser**](SnmpUser.md) |  | [optional] 
**SysContact** | Pointer to **string** | Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number. | [optional] [readonly] 
**SysLocation** | Pointer to **string** | Location of host on which the SNMP agent (server) runs. | [optional] [readonly] 
**TrapCommunity** | Pointer to **string** | SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users. | [optional] [readonly] 
**V2Enabled** | Pointer to **bool** | State of the SNMP v2c on the endpoint. If enabled, the endpoint sends SNMP v2c properties to the designated host. | [optional] [readonly] [default to true]
**V3Enabled** | Pointer to **bool** | State of the SNMP v3 on the endpoint. If enabled, the endpoint sends SNMP v3 properties to the designated host. | [optional] [readonly] [default to true]
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewSnmpPolicyInventory

`func NewSnmpPolicyInventory(classId string, objectType string, ) *SnmpPolicyInventory`

NewSnmpPolicyInventory instantiates a new SnmpPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpPolicyInventoryWithDefaults

`func NewSnmpPolicyInventoryWithDefaults() *SnmpPolicyInventory`

NewSnmpPolicyInventoryWithDefaults instantiates a new SnmpPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SnmpPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SnmpPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SnmpPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SnmpPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SnmpPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SnmpPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessCommunityString

`func (o *SnmpPolicyInventory) GetAccessCommunityString() string`

GetAccessCommunityString returns the AccessCommunityString field if non-nil, zero value otherwise.

### GetAccessCommunityStringOk

`func (o *SnmpPolicyInventory) GetAccessCommunityStringOk() (*string, bool)`

GetAccessCommunityStringOk returns a tuple with the AccessCommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCommunityString

`func (o *SnmpPolicyInventory) SetAccessCommunityString(v string)`

SetAccessCommunityString sets AccessCommunityString field to given value.

### HasAccessCommunityString

`func (o *SnmpPolicyInventory) HasAccessCommunityString() bool`

HasAccessCommunityString returns a boolean if a field has been set.

### GetCommunityAccess

`func (o *SnmpPolicyInventory) GetCommunityAccess() string`

GetCommunityAccess returns the CommunityAccess field if non-nil, zero value otherwise.

### GetCommunityAccessOk

`func (o *SnmpPolicyInventory) GetCommunityAccessOk() (*string, bool)`

GetCommunityAccessOk returns a tuple with the CommunityAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityAccess

`func (o *SnmpPolicyInventory) SetCommunityAccess(v string)`

SetCommunityAccess sets CommunityAccess field to given value.

### HasCommunityAccess

`func (o *SnmpPolicyInventory) HasCommunityAccess() bool`

HasCommunityAccess returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpPolicyInventory) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpPolicyInventory) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpPolicyInventory) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpPolicyInventory) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEngineId

`func (o *SnmpPolicyInventory) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *SnmpPolicyInventory) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *SnmpPolicyInventory) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *SnmpPolicyInventory) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSnmpPort

`func (o *SnmpPolicyInventory) GetSnmpPort() int64`

GetSnmpPort returns the SnmpPort field if non-nil, zero value otherwise.

### GetSnmpPortOk

`func (o *SnmpPolicyInventory) GetSnmpPortOk() (*int64, bool)`

GetSnmpPortOk returns a tuple with the SnmpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpPort

`func (o *SnmpPolicyInventory) SetSnmpPort(v int64)`

SetSnmpPort sets SnmpPort field to given value.

### HasSnmpPort

`func (o *SnmpPolicyInventory) HasSnmpPort() bool`

HasSnmpPort returns a boolean if a field has been set.

### GetSnmpTraps

`func (o *SnmpPolicyInventory) GetSnmpTraps() []SnmpTrap`

GetSnmpTraps returns the SnmpTraps field if non-nil, zero value otherwise.

### GetSnmpTrapsOk

`func (o *SnmpPolicyInventory) GetSnmpTrapsOk() (*[]SnmpTrap, bool)`

GetSnmpTrapsOk returns a tuple with the SnmpTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTraps

`func (o *SnmpPolicyInventory) SetSnmpTraps(v []SnmpTrap)`

SetSnmpTraps sets SnmpTraps field to given value.

### HasSnmpTraps

`func (o *SnmpPolicyInventory) HasSnmpTraps() bool`

HasSnmpTraps returns a boolean if a field has been set.

### SetSnmpTrapsNil

`func (o *SnmpPolicyInventory) SetSnmpTrapsNil(b bool)`

 SetSnmpTrapsNil sets the value for SnmpTraps to be an explicit nil

### UnsetSnmpTraps
`func (o *SnmpPolicyInventory) UnsetSnmpTraps()`

UnsetSnmpTraps ensures that no value is present for SnmpTraps, not even an explicit nil
### GetSnmpUsers

`func (o *SnmpPolicyInventory) GetSnmpUsers() []SnmpUser`

GetSnmpUsers returns the SnmpUsers field if non-nil, zero value otherwise.

### GetSnmpUsersOk

`func (o *SnmpPolicyInventory) GetSnmpUsersOk() (*[]SnmpUser, bool)`

GetSnmpUsersOk returns a tuple with the SnmpUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpUsers

`func (o *SnmpPolicyInventory) SetSnmpUsers(v []SnmpUser)`

SetSnmpUsers sets SnmpUsers field to given value.

### HasSnmpUsers

`func (o *SnmpPolicyInventory) HasSnmpUsers() bool`

HasSnmpUsers returns a boolean if a field has been set.

### SetSnmpUsersNil

`func (o *SnmpPolicyInventory) SetSnmpUsersNil(b bool)`

 SetSnmpUsersNil sets the value for SnmpUsers to be an explicit nil

### UnsetSnmpUsers
`func (o *SnmpPolicyInventory) UnsetSnmpUsers()`

UnsetSnmpUsers ensures that no value is present for SnmpUsers, not even an explicit nil
### GetSysContact

`func (o *SnmpPolicyInventory) GetSysContact() string`

GetSysContact returns the SysContact field if non-nil, zero value otherwise.

### GetSysContactOk

`func (o *SnmpPolicyInventory) GetSysContactOk() (*string, bool)`

GetSysContactOk returns a tuple with the SysContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysContact

`func (o *SnmpPolicyInventory) SetSysContact(v string)`

SetSysContact sets SysContact field to given value.

### HasSysContact

`func (o *SnmpPolicyInventory) HasSysContact() bool`

HasSysContact returns a boolean if a field has been set.

### GetSysLocation

`func (o *SnmpPolicyInventory) GetSysLocation() string`

GetSysLocation returns the SysLocation field if non-nil, zero value otherwise.

### GetSysLocationOk

`func (o *SnmpPolicyInventory) GetSysLocationOk() (*string, bool)`

GetSysLocationOk returns a tuple with the SysLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysLocation

`func (o *SnmpPolicyInventory) SetSysLocation(v string)`

SetSysLocation sets SysLocation field to given value.

### HasSysLocation

`func (o *SnmpPolicyInventory) HasSysLocation() bool`

HasSysLocation returns a boolean if a field has been set.

### GetTrapCommunity

`func (o *SnmpPolicyInventory) GetTrapCommunity() string`

GetTrapCommunity returns the TrapCommunity field if non-nil, zero value otherwise.

### GetTrapCommunityOk

`func (o *SnmpPolicyInventory) GetTrapCommunityOk() (*string, bool)`

GetTrapCommunityOk returns a tuple with the TrapCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapCommunity

`func (o *SnmpPolicyInventory) SetTrapCommunity(v string)`

SetTrapCommunity sets TrapCommunity field to given value.

### HasTrapCommunity

`func (o *SnmpPolicyInventory) HasTrapCommunity() bool`

HasTrapCommunity returns a boolean if a field has been set.

### GetV2Enabled

`func (o *SnmpPolicyInventory) GetV2Enabled() bool`

GetV2Enabled returns the V2Enabled field if non-nil, zero value otherwise.

### GetV2EnabledOk

`func (o *SnmpPolicyInventory) GetV2EnabledOk() (*bool, bool)`

GetV2EnabledOk returns a tuple with the V2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2Enabled

`func (o *SnmpPolicyInventory) SetV2Enabled(v bool)`

SetV2Enabled sets V2Enabled field to given value.

### HasV2Enabled

`func (o *SnmpPolicyInventory) HasV2Enabled() bool`

HasV2Enabled returns a boolean if a field has been set.

### GetV3Enabled

`func (o *SnmpPolicyInventory) GetV3Enabled() bool`

GetV3Enabled returns the V3Enabled field if non-nil, zero value otherwise.

### GetV3EnabledOk

`func (o *SnmpPolicyInventory) GetV3EnabledOk() (*bool, bool)`

GetV3EnabledOk returns a tuple with the V3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Enabled

`func (o *SnmpPolicyInventory) SetV3Enabled(v bool)`

SetV3Enabled sets V3Enabled field to given value.

### HasV3Enabled

`func (o *SnmpPolicyInventory) HasV3Enabled() bool`

HasV3Enabled returns a boolean if a field has been set.

### GetTargetMo

`func (o *SnmpPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *SnmpPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *SnmpPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *SnmpPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *SnmpPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *SnmpPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


