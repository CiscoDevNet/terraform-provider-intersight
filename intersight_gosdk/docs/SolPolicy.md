# SolPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaudRate** | Pointer to **int32** | Baud Rate used for Serial Over LAN communication. * &#x60;9600&#x60; - Use baud rate 9600 for communication. * &#x60;19200&#x60; - Use baud rate 19200 for communication. * &#x60;38400&#x60; - Use baud rate 38400 for communication. * &#x60;57600&#x60; - Use baud rate 57600 for communication. * &#x60;115200&#x60; - Use baud rate 115200 for communication. | [optional] [default to 9600]
**ComPort** | Pointer to **string** | Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default. * &#x60;com0&#x60; - Use serial port com0 for communication. * &#x60;com1&#x60; - Use serial port com1 for communication. | [optional] [default to "com0"]
**Enabled** | Pointer to **bool** | State of Serial Over LAN service on the endpoint. | [optional] 
**SshPort** | Pointer to **int64** | SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSolPolicy

`func NewSolPolicy() *SolPolicy`

NewSolPolicy instantiates a new SolPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolPolicyWithDefaults

`func NewSolPolicyWithDefaults() *SolPolicy`

NewSolPolicyWithDefaults instantiates a new SolPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaudRate

`func (o *SolPolicy) GetBaudRate() int32`

GetBaudRate returns the BaudRate field if non-nil, zero value otherwise.

### GetBaudRateOk

`func (o *SolPolicy) GetBaudRateOk() (*int32, bool)`

GetBaudRateOk returns a tuple with the BaudRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaudRate

`func (o *SolPolicy) SetBaudRate(v int32)`

SetBaudRate sets BaudRate field to given value.

### HasBaudRate

`func (o *SolPolicy) HasBaudRate() bool`

HasBaudRate returns a boolean if a field has been set.

### GetComPort

`func (o *SolPolicy) GetComPort() string`

GetComPort returns the ComPort field if non-nil, zero value otherwise.

### GetComPortOk

`func (o *SolPolicy) GetComPortOk() (*string, bool)`

GetComPortOk returns a tuple with the ComPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComPort

`func (o *SolPolicy) SetComPort(v string)`

SetComPort sets ComPort field to given value.

### HasComPort

`func (o *SolPolicy) HasComPort() bool`

HasComPort returns a boolean if a field has been set.

### GetEnabled

`func (o *SolPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SolPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SolPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SolPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSshPort

`func (o *SolPolicy) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *SolPolicy) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *SolPolicy) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *SolPolicy) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetOrganization

`func (o *SolPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SolPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SolPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SolPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SolPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SolPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SolPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SolPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SolPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SolPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


