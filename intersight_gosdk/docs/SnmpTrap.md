# SnmpTrap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "snmp.Trap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "snmp.Trap"]
**Destination** | Pointer to **string** | Address to which the SNMP trap information is sent. | [optional] 
**Enabled** | Pointer to **bool** | Enables/disables the trap on the server If enabled, trap is active on the server. | [optional] [default to true]
**Port** | Pointer to **int64** | Port used by the server to communicate with the trap destination. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). | [optional] [default to 162]
**Type** | Pointer to **string** | Type of trap which decides whether to receive a notification when a trap is received at the destination. * &#x60;Trap&#x60; - Do not receive notifications when trap is sent to the destination. * &#x60;Inform&#x60; - Receive notifications when trap is sent to the destination. This option is valid only for V2 users. | [optional] [default to "Trap"]
**User** | Pointer to **string** | SNMP user for the trap. Applicable only to SNMPv3. | [optional] 
**Version** | Pointer to **string** | SNMP version used for the trap. * &#x60;V3&#x60; - SNMP v3 trap version notifications. * &#x60;V2&#x60; - SNMP v2 trap version notifications. | [optional] [default to "V3"]

## Methods

### NewSnmpTrap

`func NewSnmpTrap(classId string, objectType string, ) *SnmpTrap`

NewSnmpTrap instantiates a new SnmpTrap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpTrapWithDefaults

`func NewSnmpTrapWithDefaults() *SnmpTrap`

NewSnmpTrapWithDefaults instantiates a new SnmpTrap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SnmpTrap) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SnmpTrap) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SnmpTrap) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SnmpTrap) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SnmpTrap) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SnmpTrap) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestination

`func (o *SnmpTrap) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SnmpTrap) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SnmpTrap) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SnmpTrap) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpTrap) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpTrap) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpTrap) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpTrap) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPort

`func (o *SnmpTrap) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpTrap) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpTrap) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpTrap) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetType

`func (o *SnmpTrap) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnmpTrap) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnmpTrap) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnmpTrap) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *SnmpTrap) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnmpTrap) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnmpTrap) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SnmpTrap) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVersion

`func (o *SnmpTrap) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpTrap) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpTrap) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpTrap) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


