# HciNutanixClusterCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.NutanixClusterCheck"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.NutanixClusterCheck"]
**Action** | Pointer to **string** | The action to perform. Possible values: - Run: Trigger NCC execution. - None: No action (default after workflow starts). * &#x60;None&#x60; - No action. Default state after a workflow starts or when no operation is requested. * &#x60;Run&#x60; - Trigger NCC check execution. | [optional] [default to "None"]
**AdditionalRecipients** | Pointer to **[]string** |  | [optional] 
**AnonymizeOutput** | Pointer to **bool** | Anonymize sensitive information in the check results. Always set to true for security. | [optional] [default to true]
**CheckIds** | Pointer to **[]string** |  | [optional] 
**RunAllChecks** | Pointer to **bool** | When true, runs all available NCC checks. Mutually exclusive with checkIds. | [optional] [default to false]
**SendToConfiguredRecipients** | Pointer to **bool** | Send check run summary to cluster&#39;s configured email addresses. | [optional] [default to true]

## Methods

### NewHciNutanixClusterCheck

`func NewHciNutanixClusterCheck(classId string, objectType string, ) *HciNutanixClusterCheck`

NewHciNutanixClusterCheck instantiates a new HciNutanixClusterCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciNutanixClusterCheckWithDefaults

`func NewHciNutanixClusterCheckWithDefaults() *HciNutanixClusterCheck`

NewHciNutanixClusterCheckWithDefaults instantiates a new HciNutanixClusterCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciNutanixClusterCheck) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciNutanixClusterCheck) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciNutanixClusterCheck) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciNutanixClusterCheck) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciNutanixClusterCheck) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciNutanixClusterCheck) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *HciNutanixClusterCheck) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HciNutanixClusterCheck) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HciNutanixClusterCheck) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HciNutanixClusterCheck) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAdditionalRecipients

`func (o *HciNutanixClusterCheck) GetAdditionalRecipients() []string`

GetAdditionalRecipients returns the AdditionalRecipients field if non-nil, zero value otherwise.

### GetAdditionalRecipientsOk

`func (o *HciNutanixClusterCheck) GetAdditionalRecipientsOk() (*[]string, bool)`

GetAdditionalRecipientsOk returns a tuple with the AdditionalRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRecipients

`func (o *HciNutanixClusterCheck) SetAdditionalRecipients(v []string)`

SetAdditionalRecipients sets AdditionalRecipients field to given value.

### HasAdditionalRecipients

`func (o *HciNutanixClusterCheck) HasAdditionalRecipients() bool`

HasAdditionalRecipients returns a boolean if a field has been set.

### SetAdditionalRecipientsNil

`func (o *HciNutanixClusterCheck) SetAdditionalRecipientsNil(b bool)`

 SetAdditionalRecipientsNil sets the value for AdditionalRecipients to be an explicit nil

### UnsetAdditionalRecipients
`func (o *HciNutanixClusterCheck) UnsetAdditionalRecipients()`

UnsetAdditionalRecipients ensures that no value is present for AdditionalRecipients, not even an explicit nil
### GetAnonymizeOutput

`func (o *HciNutanixClusterCheck) GetAnonymizeOutput() bool`

GetAnonymizeOutput returns the AnonymizeOutput field if non-nil, zero value otherwise.

### GetAnonymizeOutputOk

`func (o *HciNutanixClusterCheck) GetAnonymizeOutputOk() (*bool, bool)`

GetAnonymizeOutputOk returns a tuple with the AnonymizeOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeOutput

`func (o *HciNutanixClusterCheck) SetAnonymizeOutput(v bool)`

SetAnonymizeOutput sets AnonymizeOutput field to given value.

### HasAnonymizeOutput

`func (o *HciNutanixClusterCheck) HasAnonymizeOutput() bool`

HasAnonymizeOutput returns a boolean if a field has been set.

### GetCheckIds

`func (o *HciNutanixClusterCheck) GetCheckIds() []string`

GetCheckIds returns the CheckIds field if non-nil, zero value otherwise.

### GetCheckIdsOk

`func (o *HciNutanixClusterCheck) GetCheckIdsOk() (*[]string, bool)`

GetCheckIdsOk returns a tuple with the CheckIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIds

`func (o *HciNutanixClusterCheck) SetCheckIds(v []string)`

SetCheckIds sets CheckIds field to given value.

### HasCheckIds

`func (o *HciNutanixClusterCheck) HasCheckIds() bool`

HasCheckIds returns a boolean if a field has been set.

### SetCheckIdsNil

`func (o *HciNutanixClusterCheck) SetCheckIdsNil(b bool)`

 SetCheckIdsNil sets the value for CheckIds to be an explicit nil

### UnsetCheckIds
`func (o *HciNutanixClusterCheck) UnsetCheckIds()`

UnsetCheckIds ensures that no value is present for CheckIds, not even an explicit nil
### GetRunAllChecks

`func (o *HciNutanixClusterCheck) GetRunAllChecks() bool`

GetRunAllChecks returns the RunAllChecks field if non-nil, zero value otherwise.

### GetRunAllChecksOk

`func (o *HciNutanixClusterCheck) GetRunAllChecksOk() (*bool, bool)`

GetRunAllChecksOk returns a tuple with the RunAllChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAllChecks

`func (o *HciNutanixClusterCheck) SetRunAllChecks(v bool)`

SetRunAllChecks sets RunAllChecks field to given value.

### HasRunAllChecks

`func (o *HciNutanixClusterCheck) HasRunAllChecks() bool`

HasRunAllChecks returns a boolean if a field has been set.

### GetSendToConfiguredRecipients

`func (o *HciNutanixClusterCheck) GetSendToConfiguredRecipients() bool`

GetSendToConfiguredRecipients returns the SendToConfiguredRecipients field if non-nil, zero value otherwise.

### GetSendToConfiguredRecipientsOk

`func (o *HciNutanixClusterCheck) GetSendToConfiguredRecipientsOk() (*bool, bool)`

GetSendToConfiguredRecipientsOk returns a tuple with the SendToConfiguredRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToConfiguredRecipients

`func (o *HciNutanixClusterCheck) SetSendToConfiguredRecipients(v bool)`

SetSendToConfiguredRecipients sets SendToConfiguredRecipients field to given value.

### HasSendToConfiguredRecipients

`func (o *HciNutanixClusterCheck) HasSendToConfiguredRecipients() bool`

HasSendToConfiguredRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


