# SmtpPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "smtp.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "smtp.PolicyInventory"]
**AuthPassword** | Pointer to **string** | Authorization password for the process. | [optional] 
**EnableAuth** | Pointer to **bool** | If enabled, lets user input username and password. | [optional] [readonly] [default to false]
**EnableTls** | Pointer to **bool** | If enabled, lets user input valid CA certificates for authorization. | [optional] [readonly] [default to false]
**Enabled** | Pointer to **bool** | If enabled, controls the state of the SMTP client service on the managed device. | [optional] [readonly] [default to true]
**IsAuthPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;authPassword&#39; property has been set. | [optional] [readonly] [default to false]
**MinSeverity** | Pointer to **string** | Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * &#x60;critical&#x60; - Minimum severity to report is critical. * &#x60;condition&#x60; - Minimum severity to report is informational. * &#x60;warning&#x60; - Minimum severity to report is warning. * &#x60;minor&#x60; - Minimum severity to report is minor. * &#x60;major&#x60; - Minimum severity to report is major. | [optional] [readonly] [default to "critical"]
**SenderEmail** | Pointer to **string** | The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field. | [optional] [readonly] 
**SmtpPort** | Pointer to **int64** | Port number used by the SMTP server for outgoing SMTP communication. | [optional] [readonly] [default to 25]
**SmtpRecipients** | Pointer to **[]string** |  | [optional] 
**SmtpServer** | Pointer to **string** | IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications. | [optional] [readonly] 
**UserName** | Pointer to **string** | SMTP username from which email notification is sent. | [optional] [readonly] 
**Certificate** | Pointer to [**NullableIamTrustPointRelationship**](IamTrustPointRelationship.md) |  | [optional] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewSmtpPolicyInventory

`func NewSmtpPolicyInventory(classId string, objectType string, ) *SmtpPolicyInventory`

NewSmtpPolicyInventory instantiates a new SmtpPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpPolicyInventoryWithDefaults

`func NewSmtpPolicyInventoryWithDefaults() *SmtpPolicyInventory`

NewSmtpPolicyInventoryWithDefaults instantiates a new SmtpPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SmtpPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SmtpPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SmtpPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SmtpPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SmtpPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SmtpPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthPassword

`func (o *SmtpPolicyInventory) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *SmtpPolicyInventory) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *SmtpPolicyInventory) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *SmtpPolicyInventory) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetEnableAuth

`func (o *SmtpPolicyInventory) GetEnableAuth() bool`

GetEnableAuth returns the EnableAuth field if non-nil, zero value otherwise.

### GetEnableAuthOk

`func (o *SmtpPolicyInventory) GetEnableAuthOk() (*bool, bool)`

GetEnableAuthOk returns a tuple with the EnableAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuth

`func (o *SmtpPolicyInventory) SetEnableAuth(v bool)`

SetEnableAuth sets EnableAuth field to given value.

### HasEnableAuth

`func (o *SmtpPolicyInventory) HasEnableAuth() bool`

HasEnableAuth returns a boolean if a field has been set.

### GetEnableTls

`func (o *SmtpPolicyInventory) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *SmtpPolicyInventory) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *SmtpPolicyInventory) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *SmtpPolicyInventory) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetEnabled

`func (o *SmtpPolicyInventory) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SmtpPolicyInventory) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SmtpPolicyInventory) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SmtpPolicyInventory) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsAuthPasswordSet

`func (o *SmtpPolicyInventory) GetIsAuthPasswordSet() bool`

GetIsAuthPasswordSet returns the IsAuthPasswordSet field if non-nil, zero value otherwise.

### GetIsAuthPasswordSetOk

`func (o *SmtpPolicyInventory) GetIsAuthPasswordSetOk() (*bool, bool)`

GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthPasswordSet

`func (o *SmtpPolicyInventory) SetIsAuthPasswordSet(v bool)`

SetIsAuthPasswordSet sets IsAuthPasswordSet field to given value.

### HasIsAuthPasswordSet

`func (o *SmtpPolicyInventory) HasIsAuthPasswordSet() bool`

HasIsAuthPasswordSet returns a boolean if a field has been set.

### GetMinSeverity

`func (o *SmtpPolicyInventory) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *SmtpPolicyInventory) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *SmtpPolicyInventory) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *SmtpPolicyInventory) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetSenderEmail

`func (o *SmtpPolicyInventory) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *SmtpPolicyInventory) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *SmtpPolicyInventory) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *SmtpPolicyInventory) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### GetSmtpPort

`func (o *SmtpPolicyInventory) GetSmtpPort() int64`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *SmtpPolicyInventory) GetSmtpPortOk() (*int64, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *SmtpPolicyInventory) SetSmtpPort(v int64)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *SmtpPolicyInventory) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpRecipients

`func (o *SmtpPolicyInventory) GetSmtpRecipients() []string`

GetSmtpRecipients returns the SmtpRecipients field if non-nil, zero value otherwise.

### GetSmtpRecipientsOk

`func (o *SmtpPolicyInventory) GetSmtpRecipientsOk() (*[]string, bool)`

GetSmtpRecipientsOk returns a tuple with the SmtpRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpRecipients

`func (o *SmtpPolicyInventory) SetSmtpRecipients(v []string)`

SetSmtpRecipients sets SmtpRecipients field to given value.

### HasSmtpRecipients

`func (o *SmtpPolicyInventory) HasSmtpRecipients() bool`

HasSmtpRecipients returns a boolean if a field has been set.

### SetSmtpRecipientsNil

`func (o *SmtpPolicyInventory) SetSmtpRecipientsNil(b bool)`

 SetSmtpRecipientsNil sets the value for SmtpRecipients to be an explicit nil

### UnsetSmtpRecipients
`func (o *SmtpPolicyInventory) UnsetSmtpRecipients()`

UnsetSmtpRecipients ensures that no value is present for SmtpRecipients, not even an explicit nil
### GetSmtpServer

`func (o *SmtpPolicyInventory) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *SmtpPolicyInventory) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *SmtpPolicyInventory) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *SmtpPolicyInventory) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetUserName

`func (o *SmtpPolicyInventory) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SmtpPolicyInventory) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SmtpPolicyInventory) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SmtpPolicyInventory) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetCertificate

`func (o *SmtpPolicyInventory) GetCertificate() IamTrustPointRelationship`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SmtpPolicyInventory) GetCertificateOk() (*IamTrustPointRelationship, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SmtpPolicyInventory) SetCertificate(v IamTrustPointRelationship)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SmtpPolicyInventory) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *SmtpPolicyInventory) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *SmtpPolicyInventory) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetTargetMo

`func (o *SmtpPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *SmtpPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *SmtpPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *SmtpPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *SmtpPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *SmtpPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


