# ApplianceDeviceCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.DeviceCertificate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.DeviceCertificate"]
**CaCertificate** | Pointer to **string** | The base64 encoded certificate in PEM format. | [optional] [readonly] 
**CaCertificateExpiryTime** | Pointer to **time.Time** | The expiry datetime of new ca certificate which need to be applied on device connector. | [optional] [readonly] 
**CertificateRenewalExpiryTime** | Pointer to **time.Time** | The date time allocated till cert renewal will be executed. This time used here will be based on cert renewal plan. | [optional] [readonly] 
**CompletedPhases** | Pointer to [**[]ApplianceCertRenewalPhase**](ApplianceCertRenewalPhase.md) |  | [optional] 
**ConfigurationMoId** | Pointer to **string** | The operation configuration MOId. | [optional] 
**CurrentPhase** | Pointer to [**NullableApplianceCertRenewalPhase**](ApplianceCertRenewalPhase.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | End date of the certificate renewal. | [optional] [readonly] 
**LastSuccessPollTime** | Pointer to **time.Time** | The last poll time when data collection was successfull. This time is used to collect data after this time in next cycle. | [optional] [readonly] 
**Messages** | Pointer to **[]string** |  | [optional] 
**StartTime** | Pointer to **time.Time** | Start date of the certificate renewal. | [optional] 
**Status** | Pointer to **string** | The status of ca certificate renewal. | [optional] 

## Methods

### NewApplianceDeviceCertificate

`func NewApplianceDeviceCertificate(classId string, objectType string, ) *ApplianceDeviceCertificate`

NewApplianceDeviceCertificate instantiates a new ApplianceDeviceCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDeviceCertificateWithDefaults

`func NewApplianceDeviceCertificateWithDefaults() *ApplianceDeviceCertificate`

NewApplianceDeviceCertificateWithDefaults instantiates a new ApplianceDeviceCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceDeviceCertificate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceDeviceCertificate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceDeviceCertificate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceDeviceCertificate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceDeviceCertificate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceDeviceCertificate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaCertificate

`func (o *ApplianceDeviceCertificate) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *ApplianceDeviceCertificate) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *ApplianceDeviceCertificate) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *ApplianceDeviceCertificate) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateExpiryTime

`func (o *ApplianceDeviceCertificate) GetCaCertificateExpiryTime() time.Time`

GetCaCertificateExpiryTime returns the CaCertificateExpiryTime field if non-nil, zero value otherwise.

### GetCaCertificateExpiryTimeOk

`func (o *ApplianceDeviceCertificate) GetCaCertificateExpiryTimeOk() (*time.Time, bool)`

GetCaCertificateExpiryTimeOk returns a tuple with the CaCertificateExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateExpiryTime

`func (o *ApplianceDeviceCertificate) SetCaCertificateExpiryTime(v time.Time)`

SetCaCertificateExpiryTime sets CaCertificateExpiryTime field to given value.

### HasCaCertificateExpiryTime

`func (o *ApplianceDeviceCertificate) HasCaCertificateExpiryTime() bool`

HasCaCertificateExpiryTime returns a boolean if a field has been set.

### GetCertificateRenewalExpiryTime

`func (o *ApplianceDeviceCertificate) GetCertificateRenewalExpiryTime() time.Time`

GetCertificateRenewalExpiryTime returns the CertificateRenewalExpiryTime field if non-nil, zero value otherwise.

### GetCertificateRenewalExpiryTimeOk

`func (o *ApplianceDeviceCertificate) GetCertificateRenewalExpiryTimeOk() (*time.Time, bool)`

GetCertificateRenewalExpiryTimeOk returns a tuple with the CertificateRenewalExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRenewalExpiryTime

`func (o *ApplianceDeviceCertificate) SetCertificateRenewalExpiryTime(v time.Time)`

SetCertificateRenewalExpiryTime sets CertificateRenewalExpiryTime field to given value.

### HasCertificateRenewalExpiryTime

`func (o *ApplianceDeviceCertificate) HasCertificateRenewalExpiryTime() bool`

HasCertificateRenewalExpiryTime returns a boolean if a field has been set.

### GetCompletedPhases

`func (o *ApplianceDeviceCertificate) GetCompletedPhases() []ApplianceCertRenewalPhase`

GetCompletedPhases returns the CompletedPhases field if non-nil, zero value otherwise.

### GetCompletedPhasesOk

`func (o *ApplianceDeviceCertificate) GetCompletedPhasesOk() (*[]ApplianceCertRenewalPhase, bool)`

GetCompletedPhasesOk returns a tuple with the CompletedPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedPhases

`func (o *ApplianceDeviceCertificate) SetCompletedPhases(v []ApplianceCertRenewalPhase)`

SetCompletedPhases sets CompletedPhases field to given value.

### HasCompletedPhases

`func (o *ApplianceDeviceCertificate) HasCompletedPhases() bool`

HasCompletedPhases returns a boolean if a field has been set.

### SetCompletedPhasesNil

`func (o *ApplianceDeviceCertificate) SetCompletedPhasesNil(b bool)`

 SetCompletedPhasesNil sets the value for CompletedPhases to be an explicit nil

### UnsetCompletedPhases
`func (o *ApplianceDeviceCertificate) UnsetCompletedPhases()`

UnsetCompletedPhases ensures that no value is present for CompletedPhases, not even an explicit nil
### GetConfigurationMoId

`func (o *ApplianceDeviceCertificate) GetConfigurationMoId() string`

GetConfigurationMoId returns the ConfigurationMoId field if non-nil, zero value otherwise.

### GetConfigurationMoIdOk

`func (o *ApplianceDeviceCertificate) GetConfigurationMoIdOk() (*string, bool)`

GetConfigurationMoIdOk returns a tuple with the ConfigurationMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationMoId

`func (o *ApplianceDeviceCertificate) SetConfigurationMoId(v string)`

SetConfigurationMoId sets ConfigurationMoId field to given value.

### HasConfigurationMoId

`func (o *ApplianceDeviceCertificate) HasConfigurationMoId() bool`

HasConfigurationMoId returns a boolean if a field has been set.

### GetCurrentPhase

`func (o *ApplianceDeviceCertificate) GetCurrentPhase() ApplianceCertRenewalPhase`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *ApplianceDeviceCertificate) GetCurrentPhaseOk() (*ApplianceCertRenewalPhase, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *ApplianceDeviceCertificate) SetCurrentPhase(v ApplianceCertRenewalPhase)`

SetCurrentPhase sets CurrentPhase field to given value.

### HasCurrentPhase

`func (o *ApplianceDeviceCertificate) HasCurrentPhase() bool`

HasCurrentPhase returns a boolean if a field has been set.

### SetCurrentPhaseNil

`func (o *ApplianceDeviceCertificate) SetCurrentPhaseNil(b bool)`

 SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil

### UnsetCurrentPhase
`func (o *ApplianceDeviceCertificate) UnsetCurrentPhase()`

UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
### GetEndTime

`func (o *ApplianceDeviceCertificate) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceDeviceCertificate) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceDeviceCertificate) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceDeviceCertificate) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLastSuccessPollTime

`func (o *ApplianceDeviceCertificate) GetLastSuccessPollTime() time.Time`

GetLastSuccessPollTime returns the LastSuccessPollTime field if non-nil, zero value otherwise.

### GetLastSuccessPollTimeOk

`func (o *ApplianceDeviceCertificate) GetLastSuccessPollTimeOk() (*time.Time, bool)`

GetLastSuccessPollTimeOk returns a tuple with the LastSuccessPollTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessPollTime

`func (o *ApplianceDeviceCertificate) SetLastSuccessPollTime(v time.Time)`

SetLastSuccessPollTime sets LastSuccessPollTime field to given value.

### HasLastSuccessPollTime

`func (o *ApplianceDeviceCertificate) HasLastSuccessPollTime() bool`

HasLastSuccessPollTime returns a boolean if a field has been set.

### GetMessages

`func (o *ApplianceDeviceCertificate) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ApplianceDeviceCertificate) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ApplianceDeviceCertificate) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ApplianceDeviceCertificate) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ApplianceDeviceCertificate) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ApplianceDeviceCertificate) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetStartTime

`func (o *ApplianceDeviceCertificate) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceDeviceCertificate) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceDeviceCertificate) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceDeviceCertificate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceDeviceCertificate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceDeviceCertificate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceDeviceCertificate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceDeviceCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


