# ApplianceExternalSyslogSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ExternalSyslogSetting"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ExternalSyslogSetting"]
**Enabled** | Pointer to **bool** | Enable or disable External Syslog Server. | [optional] [default to false]
**ExportNginx** | Pointer to **bool** | Enable or disable exporting of Web Server access logs. | [optional] [default to false]
**Port** | Pointer to **int64** | External Syslog Server Port for connection establishment. | [optional] [default to 10514]
**Protocol** | Pointer to **string** | Protocol used to connect to external syslog server. * &#x60;TCP&#x60; - External Syslog messages sent over TCP. * &#x60;UDP&#x60; - External Syslog messages sent over UDP. * &#x60;TLS&#x60; - Secure External Syslog messages sent over TLS. | [optional] [default to "TCP"]
**Server** | Pointer to **string** | External Syslog Server Address, can be IP address or hostname. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewApplianceExternalSyslogSetting

`func NewApplianceExternalSyslogSetting(classId string, objectType string, ) *ApplianceExternalSyslogSetting`

NewApplianceExternalSyslogSetting instantiates a new ApplianceExternalSyslogSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceExternalSyslogSettingWithDefaults

`func NewApplianceExternalSyslogSettingWithDefaults() *ApplianceExternalSyslogSetting`

NewApplianceExternalSyslogSettingWithDefaults instantiates a new ApplianceExternalSyslogSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceExternalSyslogSetting) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceExternalSyslogSetting) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceExternalSyslogSetting) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceExternalSyslogSetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceExternalSyslogSetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceExternalSyslogSetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *ApplianceExternalSyslogSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplianceExternalSyslogSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplianceExternalSyslogSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplianceExternalSyslogSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExportNginx

`func (o *ApplianceExternalSyslogSetting) GetExportNginx() bool`

GetExportNginx returns the ExportNginx field if non-nil, zero value otherwise.

### GetExportNginxOk

`func (o *ApplianceExternalSyslogSetting) GetExportNginxOk() (*bool, bool)`

GetExportNginxOk returns a tuple with the ExportNginx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportNginx

`func (o *ApplianceExternalSyslogSetting) SetExportNginx(v bool)`

SetExportNginx sets ExportNginx field to given value.

### HasExportNginx

`func (o *ApplianceExternalSyslogSetting) HasExportNginx() bool`

HasExportNginx returns a boolean if a field has been set.

### GetPort

`func (o *ApplianceExternalSyslogSetting) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApplianceExternalSyslogSetting) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApplianceExternalSyslogSetting) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApplianceExternalSyslogSetting) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ApplianceExternalSyslogSetting) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplianceExternalSyslogSetting) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplianceExternalSyslogSetting) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplianceExternalSyslogSetting) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetServer

`func (o *ApplianceExternalSyslogSetting) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ApplianceExternalSyslogSetting) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ApplianceExternalSyslogSetting) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *ApplianceExternalSyslogSetting) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceExternalSyslogSetting) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceExternalSyslogSetting) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceExternalSyslogSetting) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceExternalSyslogSetting) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


