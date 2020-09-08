# SmtpPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If enabled, controls the state of the SMTP client service on the managed device. | [optional] 
**MinSeverity** | Pointer to **string** | Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * &#x60;critical&#x60; - Minimum severity to report is critical. * &#x60;condition&#x60; - Minimum severity to report is informational. * &#x60;warning&#x60; - Minimum severity to report is warning. * &#x60;minor&#x60; - Minimum severity to report is minor. * &#x60;major&#x60; - Minimum severity to report is major. | [optional] [default to "critical"]
**SenderEmail** | Pointer to **string** | The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field. | [optional] 
**SmtpPort** | Pointer to **int64** | Port number used by the SMTP server for outgoing SMTP communication. | [optional] 
**SmtpRecipients** | Pointer to **[]string** |  | [optional] 
**SmtpServer** | Pointer to **string** | IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewSmtpPolicyAllOf

`func NewSmtpPolicyAllOf() *SmtpPolicyAllOf`

NewSmtpPolicyAllOf instantiates a new SmtpPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpPolicyAllOfWithDefaults

`func NewSmtpPolicyAllOfWithDefaults() *SmtpPolicyAllOf`

NewSmtpPolicyAllOfWithDefaults instantiates a new SmtpPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SmtpPolicyAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SmtpPolicyAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SmtpPolicyAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SmtpPolicyAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinSeverity

`func (o *SmtpPolicyAllOf) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *SmtpPolicyAllOf) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *SmtpPolicyAllOf) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *SmtpPolicyAllOf) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetSenderEmail

`func (o *SmtpPolicyAllOf) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *SmtpPolicyAllOf) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *SmtpPolicyAllOf) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *SmtpPolicyAllOf) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.

### GetSmtpPort

`func (o *SmtpPolicyAllOf) GetSmtpPort() int64`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *SmtpPolicyAllOf) GetSmtpPortOk() (*int64, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *SmtpPolicyAllOf) SetSmtpPort(v int64)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *SmtpPolicyAllOf) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpRecipients

`func (o *SmtpPolicyAllOf) GetSmtpRecipients() []string`

GetSmtpRecipients returns the SmtpRecipients field if non-nil, zero value otherwise.

### GetSmtpRecipientsOk

`func (o *SmtpPolicyAllOf) GetSmtpRecipientsOk() (*[]string, bool)`

GetSmtpRecipientsOk returns a tuple with the SmtpRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpRecipients

`func (o *SmtpPolicyAllOf) SetSmtpRecipients(v []string)`

SetSmtpRecipients sets SmtpRecipients field to given value.

### HasSmtpRecipients

`func (o *SmtpPolicyAllOf) HasSmtpRecipients() bool`

HasSmtpRecipients returns a boolean if a field has been set.

### GetSmtpServer

`func (o *SmtpPolicyAllOf) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *SmtpPolicyAllOf) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *SmtpPolicyAllOf) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *SmtpPolicyAllOf) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetOrganization

`func (o *SmtpPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SmtpPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SmtpPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SmtpPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *SmtpPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SmtpPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SmtpPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SmtpPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *SmtpPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *SmtpPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


