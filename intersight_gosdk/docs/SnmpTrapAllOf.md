# SnmpTrapAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Address to which the SNMP trap information is sent. | [optional] 
**Enabled** | Pointer to **bool** | Enables/disables the trap on the server If enabled, trap is active on the server. | [optional] 
**Port** | Pointer to **int64** | Port used by the server to communicate with trap destination. Enter a value between 1-65535. | [optional] 
**Type** | Pointer to **string** | Type of trap which decides whether to receive a notification when a trap is received at the destination. * &#x60;Trap&#x60; - Do not receive notifications when trap is sent to the destination. * &#x60;Inform&#x60; - Receive notifications when trap is sent to the destination. This option is valid only for V2 users. | [optional] [default to "Trap"]
**User** | Pointer to **string** | SNMP user for the trap. Applicable only to SNMPv3. | [optional] 
**Version** | Pointer to **string** | SNMP version used for the trap. * &#x60;V3&#x60; - SNMP v3 trap version notifications. * &#x60;V2&#x60; - SNMP v2 trap version notifications. | [optional] [default to "V3"]

## Methods

### NewSnmpTrapAllOf

`func NewSnmpTrapAllOf() *SnmpTrapAllOf`

NewSnmpTrapAllOf instantiates a new SnmpTrapAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpTrapAllOfWithDefaults

`func NewSnmpTrapAllOfWithDefaults() *SnmpTrapAllOf`

NewSnmpTrapAllOfWithDefaults instantiates a new SnmpTrapAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


