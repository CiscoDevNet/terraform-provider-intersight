# SolPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sol.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sol.PolicyInventory"]
**BaudRate** | Pointer to **int32** | Baud Rate used for Serial Over LAN communication. * &#x60;9600&#x60; - Use baud rate 9600 for communication. * &#x60;19200&#x60; - Use baud rate 19200 for communication. * &#x60;38400&#x60; - Use baud rate 38400 for communication. * &#x60;57600&#x60; - Use baud rate 57600 for communication. * &#x60;115200&#x60; - Use baud rate 115200 for communication. | [optional] [readonly] [default to 9600]
**ComPort** | Pointer to **string** | Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default. * &#x60;com0&#x60; - Use serial port com0 for communication. * &#x60;com1&#x60; - Use serial port com1 for communication. | [optional] [readonly] [default to "com0"]
**Enabled** | Pointer to **bool** | State of Serial Over LAN service on the endpoint. | [optional] [readonly] [default to true]
**SshPort** | Pointer to **int64** | SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN. | [optional] [readonly] [default to 2400]
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewSolPolicyInventoryAllOf

`func NewSolPolicyInventoryAllOf(classId string, objectType string, ) *SolPolicyInventoryAllOf`

NewSolPolicyInventoryAllOf instantiates a new SolPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolPolicyInventoryAllOfWithDefaults

`func NewSolPolicyInventoryAllOfWithDefaults() *SolPolicyInventoryAllOf`

NewSolPolicyInventoryAllOfWithDefaults instantiates a new SolPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SolPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SolPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SolPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SolPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SolPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SolPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBaudRate

`func (o *SolPolicyInventoryAllOf) GetBaudRate() int32`

GetBaudRate returns the BaudRate field if non-nil, zero value otherwise.

### GetBaudRateOk

`func (o *SolPolicyInventoryAllOf) GetBaudRateOk() (*int32, bool)`

GetBaudRateOk returns a tuple with the BaudRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaudRate

`func (o *SolPolicyInventoryAllOf) SetBaudRate(v int32)`

SetBaudRate sets BaudRate field to given value.

### HasBaudRate

`func (o *SolPolicyInventoryAllOf) HasBaudRate() bool`

HasBaudRate returns a boolean if a field has been set.

### GetComPort

`func (o *SolPolicyInventoryAllOf) GetComPort() string`

GetComPort returns the ComPort field if non-nil, zero value otherwise.

### GetComPortOk

`func (o *SolPolicyInventoryAllOf) GetComPortOk() (*string, bool)`

GetComPortOk returns a tuple with the ComPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComPort

`func (o *SolPolicyInventoryAllOf) SetComPort(v string)`

SetComPort sets ComPort field to given value.

### HasComPort

`func (o *SolPolicyInventoryAllOf) HasComPort() bool`

HasComPort returns a boolean if a field has been set.

### GetEnabled

`func (o *SolPolicyInventoryAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SolPolicyInventoryAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SolPolicyInventoryAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SolPolicyInventoryAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSshPort

`func (o *SolPolicyInventoryAllOf) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *SolPolicyInventoryAllOf) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *SolPolicyInventoryAllOf) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *SolPolicyInventoryAllOf) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetTargetMo

`func (o *SolPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *SolPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *SolPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *SolPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


