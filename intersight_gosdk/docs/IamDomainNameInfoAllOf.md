# IamDomainNameInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.DomainNameInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.DomainNameInfo"]
**Action** | Pointer to **string** | Regenerate TXT record and validate TXT record. * &#x60;generate&#x60; - Generate TXT record for domain name ownership validation. * &#x60;verify&#x60; - Verify TXT record for domain name ownership validation. | [optional] [default to "generate"]
**DomainName** | Pointer to **string** | Email domain name. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication. | [optional] 
**FailureDetails** | Pointer to [**NullableIamFailureDetails**](IamFailureDetails.md) |  | [optional] 
**RecordExpiryTime** | Pointer to **time.Time** | Expiration time of TXT Record. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of Domain Ownership Verification. * &#x60;Pending&#x60; - Domain verification is pending. * &#x60;Failed&#x60; - Domain verification failed. Re-generate token and verify. * &#x60;Verified&#x60; - Domain verification succeeded. * &#x60;Expired&#x60; - TXT Record for Domain verification expired. | [optional] [readonly] [default to "Pending"]
**TxtRecord** | Pointer to **string** | Resource record used to verify Domain Ownership. Users need to verify the domain by adding the TXT Record in their DNS Host. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamDomainNameInfoAllOf

`func NewIamDomainNameInfoAllOf(classId string, objectType string, ) *IamDomainNameInfoAllOf`

NewIamDomainNameInfoAllOf instantiates a new IamDomainNameInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamDomainNameInfoAllOfWithDefaults

`func NewIamDomainNameInfoAllOfWithDefaults() *IamDomainNameInfoAllOf`

NewIamDomainNameInfoAllOfWithDefaults instantiates a new IamDomainNameInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamDomainNameInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamDomainNameInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamDomainNameInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamDomainNameInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamDomainNameInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamDomainNameInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *IamDomainNameInfoAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IamDomainNameInfoAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IamDomainNameInfoAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IamDomainNameInfoAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDomainName

`func (o *IamDomainNameInfoAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *IamDomainNameInfoAllOf) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *IamDomainNameInfoAllOf) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *IamDomainNameInfoAllOf) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetFailureDetails

`func (o *IamDomainNameInfoAllOf) GetFailureDetails() IamFailureDetails`

GetFailureDetails returns the FailureDetails field if non-nil, zero value otherwise.

### GetFailureDetailsOk

`func (o *IamDomainNameInfoAllOf) GetFailureDetailsOk() (*IamFailureDetails, bool)`

GetFailureDetailsOk returns a tuple with the FailureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDetails

`func (o *IamDomainNameInfoAllOf) SetFailureDetails(v IamFailureDetails)`

SetFailureDetails sets FailureDetails field to given value.

### HasFailureDetails

`func (o *IamDomainNameInfoAllOf) HasFailureDetails() bool`

HasFailureDetails returns a boolean if a field has been set.

### SetFailureDetailsNil

`func (o *IamDomainNameInfoAllOf) SetFailureDetailsNil(b bool)`

 SetFailureDetailsNil sets the value for FailureDetails to be an explicit nil

### UnsetFailureDetails
`func (o *IamDomainNameInfoAllOf) UnsetFailureDetails()`

UnsetFailureDetails ensures that no value is present for FailureDetails, not even an explicit nil
### GetRecordExpiryTime

`func (o *IamDomainNameInfoAllOf) GetRecordExpiryTime() time.Time`

GetRecordExpiryTime returns the RecordExpiryTime field if non-nil, zero value otherwise.

### GetRecordExpiryTimeOk

`func (o *IamDomainNameInfoAllOf) GetRecordExpiryTimeOk() (*time.Time, bool)`

GetRecordExpiryTimeOk returns a tuple with the RecordExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordExpiryTime

`func (o *IamDomainNameInfoAllOf) SetRecordExpiryTime(v time.Time)`

SetRecordExpiryTime sets RecordExpiryTime field to given value.

### HasRecordExpiryTime

`func (o *IamDomainNameInfoAllOf) HasRecordExpiryTime() bool`

HasRecordExpiryTime returns a boolean if a field has been set.

### GetStatus

`func (o *IamDomainNameInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamDomainNameInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamDomainNameInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamDomainNameInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTxtRecord

`func (o *IamDomainNameInfoAllOf) GetTxtRecord() string`

GetTxtRecord returns the TxtRecord field if non-nil, zero value otherwise.

### GetTxtRecordOk

`func (o *IamDomainNameInfoAllOf) GetTxtRecordOk() (*string, bool)`

GetTxtRecordOk returns a tuple with the TxtRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtRecord

`func (o *IamDomainNameInfoAllOf) SetTxtRecord(v string)`

SetTxtRecord sets TxtRecord field to given value.

### HasTxtRecord

`func (o *IamDomainNameInfoAllOf) HasTxtRecord() bool`

HasTxtRecord returns a boolean if a field has been set.

### GetAccount

`func (o *IamDomainNameInfoAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamDomainNameInfoAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamDomainNameInfoAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamDomainNameInfoAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


