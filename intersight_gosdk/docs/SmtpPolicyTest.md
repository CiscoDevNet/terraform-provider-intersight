# SmtpPolicyTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "smtp.PolicyTest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "smtp.PolicyTest"]
**Enabled** | Pointer to **bool** | If enabled, an attempt will be made to send a test email notification to the configured recipients. | [optional] [default to true]
**LastStatusDetails** | Pointer to **string** | Information on the last status, including any encountered error specifics, to assist users in troubleshooting SMTP policy issues. | [optional] [readonly] 
**LastSuccessfulAttemptTime** | Pointer to **time.Time** | The date and time of the most recent successful attempt to send a test email notification. | [optional] [readonly] 
**LastSuccessfulPolicy** | Pointer to **interface{}** | The SMTP policy of the most recent successful attempt to send a test email notification. | [optional] [readonly] 
**LastTestStatus** | Pointer to **string** | Status of the last test email notification. * &#x60;&#x60; - Notification has not been tried. * &#x60;failed&#x60; - Sending notification failed. * &#x60;success&#x60; - Sending notification successful. | [optional] [readonly] [default to ""]
**Recipients** | Pointer to **[]string** |  | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Policy** | Pointer to [**NullableSmtpPolicyRelationship**](SmtpPolicyRelationship.md) |  | [optional] 

## Methods

### NewSmtpPolicyTest

`func NewSmtpPolicyTest(classId string, objectType string, ) *SmtpPolicyTest`

NewSmtpPolicyTest instantiates a new SmtpPolicyTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpPolicyTestWithDefaults

`func NewSmtpPolicyTestWithDefaults() *SmtpPolicyTest`

NewSmtpPolicyTestWithDefaults instantiates a new SmtpPolicyTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SmtpPolicyTest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SmtpPolicyTest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SmtpPolicyTest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SmtpPolicyTest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SmtpPolicyTest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SmtpPolicyTest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *SmtpPolicyTest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SmtpPolicyTest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SmtpPolicyTest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SmtpPolicyTest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastStatusDetails

`func (o *SmtpPolicyTest) GetLastStatusDetails() string`

GetLastStatusDetails returns the LastStatusDetails field if non-nil, zero value otherwise.

### GetLastStatusDetailsOk

`func (o *SmtpPolicyTest) GetLastStatusDetailsOk() (*string, bool)`

GetLastStatusDetailsOk returns a tuple with the LastStatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusDetails

`func (o *SmtpPolicyTest) SetLastStatusDetails(v string)`

SetLastStatusDetails sets LastStatusDetails field to given value.

### HasLastStatusDetails

`func (o *SmtpPolicyTest) HasLastStatusDetails() bool`

HasLastStatusDetails returns a boolean if a field has been set.

### GetLastSuccessfulAttemptTime

`func (o *SmtpPolicyTest) GetLastSuccessfulAttemptTime() time.Time`

GetLastSuccessfulAttemptTime returns the LastSuccessfulAttemptTime field if non-nil, zero value otherwise.

### GetLastSuccessfulAttemptTimeOk

`func (o *SmtpPolicyTest) GetLastSuccessfulAttemptTimeOk() (*time.Time, bool)`

GetLastSuccessfulAttemptTimeOk returns a tuple with the LastSuccessfulAttemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulAttemptTime

`func (o *SmtpPolicyTest) SetLastSuccessfulAttemptTime(v time.Time)`

SetLastSuccessfulAttemptTime sets LastSuccessfulAttemptTime field to given value.

### HasLastSuccessfulAttemptTime

`func (o *SmtpPolicyTest) HasLastSuccessfulAttemptTime() bool`

HasLastSuccessfulAttemptTime returns a boolean if a field has been set.

### GetLastSuccessfulPolicy

`func (o *SmtpPolicyTest) GetLastSuccessfulPolicy() interface{}`

GetLastSuccessfulPolicy returns the LastSuccessfulPolicy field if non-nil, zero value otherwise.

### GetLastSuccessfulPolicyOk

`func (o *SmtpPolicyTest) GetLastSuccessfulPolicyOk() (*interface{}, bool)`

GetLastSuccessfulPolicyOk returns a tuple with the LastSuccessfulPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulPolicy

`func (o *SmtpPolicyTest) SetLastSuccessfulPolicy(v interface{})`

SetLastSuccessfulPolicy sets LastSuccessfulPolicy field to given value.

### HasLastSuccessfulPolicy

`func (o *SmtpPolicyTest) HasLastSuccessfulPolicy() bool`

HasLastSuccessfulPolicy returns a boolean if a field has been set.

### SetLastSuccessfulPolicyNil

`func (o *SmtpPolicyTest) SetLastSuccessfulPolicyNil(b bool)`

 SetLastSuccessfulPolicyNil sets the value for LastSuccessfulPolicy to be an explicit nil

### UnsetLastSuccessfulPolicy
`func (o *SmtpPolicyTest) UnsetLastSuccessfulPolicy()`

UnsetLastSuccessfulPolicy ensures that no value is present for LastSuccessfulPolicy, not even an explicit nil
### GetLastTestStatus

`func (o *SmtpPolicyTest) GetLastTestStatus() string`

GetLastTestStatus returns the LastTestStatus field if non-nil, zero value otherwise.

### GetLastTestStatusOk

`func (o *SmtpPolicyTest) GetLastTestStatusOk() (*string, bool)`

GetLastTestStatusOk returns a tuple with the LastTestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestStatus

`func (o *SmtpPolicyTest) SetLastTestStatus(v string)`

SetLastTestStatus sets LastTestStatus field to given value.

### HasLastTestStatus

`func (o *SmtpPolicyTest) HasLastTestStatus() bool`

HasLastTestStatus returns a boolean if a field has been set.

### GetRecipients

`func (o *SmtpPolicyTest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *SmtpPolicyTest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *SmtpPolicyTest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *SmtpPolicyTest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *SmtpPolicyTest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *SmtpPolicyTest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetAccount

`func (o *SmtpPolicyTest) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SmtpPolicyTest) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SmtpPolicyTest) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SmtpPolicyTest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *SmtpPolicyTest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *SmtpPolicyTest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetPolicy

`func (o *SmtpPolicyTest) GetPolicy() SmtpPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SmtpPolicyTest) GetPolicyOk() (*SmtpPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SmtpPolicyTest) SetPolicy(v SmtpPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *SmtpPolicyTest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *SmtpPolicyTest) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *SmtpPolicyTest) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


