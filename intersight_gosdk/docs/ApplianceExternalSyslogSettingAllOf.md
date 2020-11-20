# ApplianceExternalSyslogSettingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ExternalSyslogSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ExternalSyslogSetting"]
**Enabled** | Pointer to **bool** | Enable or disable External Syslog Server. | [optional] 
**Port** | Pointer to **int64** | External Syslog Server Port for connection establishment. | [optional] 
**Server** | Pointer to **string** | External Syslog Server Address, can be IP address or hostname. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewApplianceExternalSyslogSettingAllOf

`func NewApplianceExternalSyslogSettingAllOf(classId string, objectType string, ) *ApplianceExternalSyslogSettingAllOf`

NewApplianceExternalSyslogSettingAllOf instantiates a new ApplianceExternalSyslogSettingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceExternalSyslogSettingAllOfWithDefaults

`func NewApplianceExternalSyslogSettingAllOfWithDefaults() *ApplianceExternalSyslogSettingAllOf`

NewApplianceExternalSyslogSettingAllOfWithDefaults instantiates a new ApplianceExternalSyslogSettingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceExternalSyslogSettingAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceExternalSyslogSettingAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceExternalSyslogSettingAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceExternalSyslogSettingAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceExternalSyslogSettingAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceExternalSyslogSettingAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *ApplianceExternalSyslogSettingAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplianceExternalSyslogSettingAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplianceExternalSyslogSettingAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplianceExternalSyslogSettingAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPort

`func (o *ApplianceExternalSyslogSettingAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApplianceExternalSyslogSettingAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApplianceExternalSyslogSettingAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApplianceExternalSyslogSettingAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServer

`func (o *ApplianceExternalSyslogSettingAllOf) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ApplianceExternalSyslogSettingAllOf) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ApplianceExternalSyslogSettingAllOf) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *ApplianceExternalSyslogSettingAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceExternalSyslogSettingAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceExternalSyslogSettingAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceExternalSyslogSettingAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceExternalSyslogSettingAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


