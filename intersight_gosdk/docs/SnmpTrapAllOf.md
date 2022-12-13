# SnmpTrapAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "snmp.Trap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "snmp.Trap"]
**Community** | Pointer to **string** | SNMP community group used for sending SNMP trap to other devices. Applicable only for SNMP v2c. | [optional] 
**Destination** | Pointer to **string** | Address to which the SNMP trap information is sent. | [optional] 
**Enabled** | Pointer to **bool** | Enables/disables the trap on the server If enabled, trap is active on the server. | [optional] [default to true]
**Port** | Pointer to **int64** | Port used by the server to communicate with the trap destination. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). | [optional] [default to 162]
**SecurityLevel** | Pointer to **string** | Security level of the trap receiver used for communication. * &#x60;AuthPriv&#x60; - The user requires both an authorization password and a privacy password. * &#x60;NoAuthNoPriv&#x60; - The user does not require an authorization or privacy password. * &#x60;AuthNoPriv&#x60; - The user requires an authorization password but not a privacy password. | [optional] [readonly] [default to "AuthPriv"]
**Type** | Pointer to **string** | Type of trap which decides whether to receive a notification when a trap is received at the destination. * &#x60;Trap&#x60; - Do not receive notifications when trap is sent to the destination. * &#x60;Inform&#x60; - Receive notifications when trap is sent to the destination. This option is valid only for V2 users. | [optional] [default to "Trap"]
**User** | Pointer to **string** | SNMP user for the trap. Applicable only to SNMPv3. | [optional] 
**Version** | Pointer to **string** | SNMP version used for the trap. * &#x60;V3&#x60; - SNMP v3 trap version notifications. * &#x60;V1&#x60; - SNMP v1 trap version notifications. * &#x60;V2&#x60; - SNMP v2 trap version notifications. | [optional] [default to "V3"]
**VrfName** | Pointer to **string** | VRF name of the SNMP server. | [optional] [readonly] 

## Methods

### NewSnmpTrapAllOf

`func NewSnmpTrapAllOf(classId string, objectType string, ) *SnmpTrapAllOf`

NewSnmpTrapAllOf instantiates a new SnmpTrapAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpTrapAllOfWithDefaults

`func NewSnmpTrapAllOfWithDefaults() *SnmpTrapAllOf`

NewSnmpTrapAllOfWithDefaults instantiates a new SnmpTrapAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SnmpTrapAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SnmpTrapAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SnmpTrapAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SnmpTrapAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SnmpTrapAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SnmpTrapAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCommunity

`func (o *SnmpTrapAllOf) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *SnmpTrapAllOf) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *SnmpTrapAllOf) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *SnmpTrapAllOf) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetDestination

`func (o *SnmpTrapAllOf) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SnmpTrapAllOf) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SnmpTrapAllOf) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SnmpTrapAllOf) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpTrapAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpTrapAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpTrapAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpTrapAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPort

`func (o *SnmpTrapAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpTrapAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpTrapAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpTrapAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SnmpTrapAllOf) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpTrapAllOf) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpTrapAllOf) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpTrapAllOf) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetType

`func (o *SnmpTrapAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnmpTrapAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnmpTrapAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnmpTrapAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *SnmpTrapAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnmpTrapAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnmpTrapAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SnmpTrapAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *SnmpTrapAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpTrapAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpTrapAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpTrapAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVrfName

`func (o *SnmpTrapAllOf) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *SnmpTrapAllOf) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *SnmpTrapAllOf) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *SnmpTrapAllOf) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


