# SnmpPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "snmp.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "snmp.Policy"]
**AccessCommunityString** | Pointer to **string** | The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long. | [optional] 
**CommunityAccess** | Pointer to **string** | Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users. * &#x60;Disabled&#x60; - Blocks access to the information in the inventory tables. * &#x60;Limited&#x60; - Partial access to read the information in the inventory tables. * &#x60;Full&#x60; - Full access to read the information in the inventory tables. | [optional] [default to "Disabled"]
**Enabled** | Pointer to **bool** | State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host. | [optional] [default to true]
**EngineId** | Pointer to **string** | User-defined unique identification of the static engine. | [optional] 
**SnmpPort** | Pointer to **int64** | Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). | [optional] [default to 161]
**SnmpTraps** | Pointer to [**[]SnmpTrap**](SnmpTrap.md) |  | [optional] 
**SnmpUsers** | Pointer to [**[]SnmpUser**](SnmpUser.md) |  | [optional] 
**SysContact** | Pointer to **string** | Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number. | [optional] 
**SysLocation** | Pointer to **string** | Location of host on which the SNMP agent (server) runs. | [optional] 
**TrapCommunity** | Pointer to **string** | SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users. | [optional] 
**V2Enabled** | Pointer to **bool** | State of the SNMP v2c on the endpoint. If enabled, the endpoint sends SNMP v2c properties to the designated host. | [optional] [default to true]
**V3Enabled** | Pointer to **bool** | State of the SNMP v3 on the endpoint. If enabled, the endpoint sends SNMP v3 properties to the designated host. | [optional] [default to true]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSnmpPolicy

`func NewSnmpPolicy(classId string, objectType string, ) *SnmpPolicy`

NewSnmpPolicy instantiates a new SnmpPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpPolicyWithDefaults

`func NewSnmpPolicyWithDefaults() *SnmpPolicy`

NewSnmpPolicyWithDefaults instantiates a new SnmpPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SnmpPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SnmpPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SnmpPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SnmpPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SnmpPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SnmpPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessCommunityString

`func (o *SnmpPolicy) GetAccessCommunityString() string`

GetAccessCommunityString returns the AccessCommunityString field if non-nil, zero value otherwise.

### GetAccessCommunityStringOk

`func (o *SnmpPolicy) GetAccessCommunityStringOk() (*string, bool)`

GetAccessCommunityStringOk returns a tuple with the AccessCommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCommunityString

`func (o *SnmpPolicy) SetAccessCommunityString(v string)`

SetAccessCommunityString sets AccessCommunityString field to given value.

### HasAccessCommunityString

`func (o *SnmpPolicy) HasAccessCommunityString() bool`

HasAccessCommunityString returns a boolean if a field has been set.

### GetCommunityAccess

`func (o *SnmpPolicy) GetCommunityAccess() string`

GetCommunityAccess returns the CommunityAccess field if non-nil, zero value otherwise.

### GetCommunityAccessOk

`func (o *SnmpPolicy) GetCommunityAccessOk() (*string, bool)`

GetCommunityAccessOk returns a tuple with the CommunityAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityAccess

`func (o *SnmpPolicy) SetCommunityAccess(v string)`

SetCommunityAccess sets CommunityAccess field to given value.

### HasCommunityAccess

`func (o *SnmpPolicy) HasCommunityAccess() bool`

HasCommunityAccess returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEngineId

`func (o *SnmpPolicy) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *SnmpPolicy) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *SnmpPolicy) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *SnmpPolicy) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetSnmpPort

`func (o *SnmpPolicy) GetSnmpPort() int64`

GetSnmpPort returns the SnmpPort field if non-nil, zero value otherwise.

### GetSnmpPortOk

`func (o *SnmpPolicy) GetSnmpPortOk() (*int64, bool)`

GetSnmpPortOk returns a tuple with the SnmpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpPort

`func (o *SnmpPolicy) SetSnmpPort(v int64)`

SetSnmpPort sets SnmpPort field to given value.

### HasSnmpPort

`func (o *SnmpPolicy) HasSnmpPort() bool`

HasSnmpPort returns a boolean if a field has been set.

### GetSnmpTraps

`func (o *SnmpPolicy) GetSnmpTraps() []SnmpTrap`

GetSnmpTraps returns the SnmpTraps field if non-nil, zero value otherwise.

### GetSnmpTrapsOk

`func (o *SnmpPolicy) GetSnmpTrapsOk() (*[]SnmpTrap, bool)`

GetSnmpTrapsOk returns a tuple with the SnmpTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTraps

`func (o *SnmpPolicy) SetSnmpTraps(v []SnmpTrap)`

SetSnmpTraps sets SnmpTraps field to given value.

### HasSnmpTraps

`func (o *SnmpPolicy) HasSnmpTraps() bool`

HasSnmpTraps returns a boolean if a field has been set.

### SetSnmpTrapsNil

`func (o *SnmpPolicy) SetSnmpTrapsNil(b bool)`

 SetSnmpTrapsNil sets the value for SnmpTraps to be an explicit nil

### UnsetSnmpTraps
`func (o *SnmpPolicy) UnsetSnmpTraps()`

UnsetSnmpTraps ensures that no value is present for SnmpTraps, not even an explicit nil
### GetSnmpUsers

`func (o *SnmpPolicy) GetSnmpUsers() []SnmpUser`

GetSnmpUsers returns the SnmpUsers field if non-nil, zero value otherwise.

### GetSnmpUsersOk

`func (o *SnmpPolicy) GetSnmpUsersOk() (*[]SnmpUser, bool)`

GetSnmpUsersOk returns a tuple with the SnmpUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpUsers

`func (o *SnmpPolicy) SetSnmpUsers(v []SnmpUser)`

SetSnmpUsers sets SnmpUsers field to given value.

### HasSnmpUsers

`func (o *SnmpPolicy) HasSnmpUsers() bool`

HasSnmpUsers returns a boolean if a field has been set.

### SetSnmpUsersNil

`func (o *SnmpPolicy) SetSnmpUsersNil(b bool)`

 SetSnmpUsersNil sets the value for SnmpUsers to be an explicit nil

### UnsetSnmpUsers
`func (o *SnmpPolicy) UnsetSnmpUsers()`

UnsetSnmpUsers ensures that no value is present for SnmpUsers, not even an explicit nil
### GetSysContact

`func (o *SnmpPolicy) GetSysContact() string`

GetSysContact returns the SysContact field if non-nil, zero value otherwise.

### GetSysContactOk

`func (o *SnmpPolicy) GetSysContactOk() (*string, bool)`

GetSysContactOk returns a tuple with the SysContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysContact

`func (o *SnmpPolicy) SetSysContact(v string)`

SetSysContact sets SysContact field to given value.

### HasSysContact

`func (o *SnmpPolicy) HasSysContact() bool`

HasSysContact returns a boolean if a field has been set.

### GetSysLocation

`func (o *SnmpPolicy) GetSysLocation() string`

GetSysLocation returns the SysLocation field if non-nil, zero value otherwise.

### GetSysLocationOk

`func (o *SnmpPolicy) GetSysLocationOk() (*string, bool)`

GetSysLocationOk returns a tuple with the SysLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysLocation

`func (o *SnmpPolicy) SetSysLocation(v string)`

SetSysLocation sets SysLocation field to given value.

### HasSysLocation

`func (o *SnmpPolicy) HasSysLocation() bool`

HasSysLocation returns a boolean if a field has been set.

### GetTrapCommunity

`func (o *SnmpPolicy) GetTrapCommunity() string`

GetTrapCommunity returns the TrapCommunity field if non-nil, zero value otherwise.

### GetTrapCommunityOk

`func (o *SnmpPolicy) GetTrapCommunityOk() (*string, bool)`

GetTrapCommunityOk returns a tuple with the TrapCommunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapCommunity

`func (o *SnmpPolicy) SetTrapCommunity(v string)`

SetTrapCommunity sets TrapCommunity field to given value.

### HasTrapCommunity

`func (o *SnmpPolicy) HasTrapCommunity() bool`

HasTrapCommunity returns a boolean if a field has been set.

### GetV2Enabled

`func (o *SnmpPolicy) GetV2Enabled() bool`

GetV2Enabled returns the V2Enabled field if non-nil, zero value otherwise.

### GetV2EnabledOk

`func (o *SnmpPolicy) GetV2EnabledOk() (*bool, bool)`

GetV2EnabledOk returns a tuple with the V2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2Enabled

`func (o *SnmpPolicy) SetV2Enabled(v bool)`

SetV2Enabled sets V2Enabled field to given value.

### HasV2Enabled

`func (o *SnmpPolicy) HasV2Enabled() bool`

HasV2Enabled returns a boolean if a field has been set.

### GetV3Enabled

`func (o *SnmpPolicy) GetV3Enabled() bool`

GetV3Enabled returns the V3Enabled field if non-nil, zero value otherwise.

### GetV3EnabledOk

`func (o *SnmpPolicy) GetV3EnabledOk() (*bool, bool)`

GetV3EnabledOk returns a tuple with the V3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3Enabled

`func (o *SnmpPolicy) SetV3Enabled(v bool)`

SetV3Enabled sets V3Enabled field to given value.

### HasV3Enabled

`func (o *SnmpPolicy) HasV3Enabled() bool`

HasV3Enabled returns a boolean if a field has been set.

### GetOrganization

`func (o *SnmpPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SnmpPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SnmpPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SnmpPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SnmpPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SnmpPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SnmpPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SnmpPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SnmpPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SnmpPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


