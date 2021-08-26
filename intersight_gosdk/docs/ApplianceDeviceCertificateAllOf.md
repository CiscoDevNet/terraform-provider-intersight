# ApplianceDeviceCertificateAllOf

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

### NewApplianceDeviceCertificateAllOf

`func NewApplianceDeviceCertificateAllOf(classId string, objectType string, ) *ApplianceDeviceCertificateAllOf`

NewApplianceDeviceCertificateAllOf instantiates a new ApplianceDeviceCertificateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDeviceCertificateAllOfWithDefaults

`func NewApplianceDeviceCertificateAllOfWithDefaults() *ApplianceDeviceCertificateAllOf`

NewApplianceDeviceCertificateAllOfWithDefaults instantiates a new ApplianceDeviceCertificateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceDeviceCertificateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceDeviceCertificateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceDeviceCertificateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceDeviceCertificateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceDeviceCertificateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceDeviceCertificateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCaCertificate

`func (o *ApplianceDeviceCertificateAllOf) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *ApplianceDeviceCertificateAllOf) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *ApplianceDeviceCertificateAllOf) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *ApplianceDeviceCertificateAllOf) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCaCertificateExpiryTime

`func (o *ApplianceDeviceCertificateAllOf) GetCaCertificateExpiryTime() time.Time`

GetCaCertificateExpiryTime returns the CaCertificateExpiryTime field if non-nil, zero value otherwise.

### GetCaCertificateExpiryTimeOk

`func (o *ApplianceDeviceCertificateAllOf) GetCaCertificateExpiryTimeOk() (*time.Time, bool)`

GetCaCertificateExpiryTimeOk returns a tuple with the CaCertificateExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateExpiryTime

`func (o *ApplianceDeviceCertificateAllOf) SetCaCertificateExpiryTime(v time.Time)`

SetCaCertificateExpiryTime sets CaCertificateExpiryTime field to given value.

### HasCaCertificateExpiryTime

`func (o *ApplianceDeviceCertificateAllOf) HasCaCertificateExpiryTime() bool`

HasCaCertificateExpiryTime returns a boolean if a field has been set.

### GetCertificateRenewalExpiryTime

`func (o *ApplianceDeviceCertificateAllOf) GetCertificateRenewalExpiryTime() time.Time`

GetCertificateRenewalExpiryTime returns the CertificateRenewalExpiryTime field if non-nil, zero value otherwise.

### GetCertificateRenewalExpiryTimeOk

`func (o *ApplianceDeviceCertificateAllOf) GetCertificateRenewalExpiryTimeOk() (*time.Time, bool)`

GetCertificateRenewalExpiryTimeOk returns a tuple with the CertificateRenewalExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRenewalExpiryTime

`func (o *ApplianceDeviceCertificateAllOf) SetCertificateRenewalExpiryTime(v time.Time)`

SetCertificateRenewalExpiryTime sets CertificateRenewalExpiryTime field to given value.

### HasCertificateRenewalExpiryTime

`func (o *ApplianceDeviceCertificateAllOf) HasCertificateRenewalExpiryTime() bool`

HasCertificateRenewalExpiryTime returns a boolean if a field has been set.

### GetCompletedPhases

`func (o *ApplianceDeviceCertificateAllOf) GetCompletedPhases() []ApplianceCertRenewalPhase`

GetCompletedPhases returns the CompletedPhases field if non-nil, zero value otherwise.

### GetCompletedPhasesOk

`func (o *ApplianceDeviceCertificateAllOf) GetCompletedPhasesOk() (*[]ApplianceCertRenewalPhase, bool)`

GetCompletedPhasesOk returns a tuple with the CompletedPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedPhases

`func (o *ApplianceDeviceCertificateAllOf) SetCompletedPhases(v []ApplianceCertRenewalPhase)`

SetCompletedPhases sets CompletedPhases field to given value.

### HasCompletedPhases

`func (o *ApplianceDeviceCertificateAllOf) HasCompletedPhases() bool`

HasCompletedPhases returns a boolean if a field has been set.

### SetCompletedPhasesNil

`func (o *ApplianceDeviceCertificateAllOf) SetCompletedPhasesNil(b bool)`

 SetCompletedPhasesNil sets the value for CompletedPhases to be an explicit nil

### UnsetCompletedPhases
`func (o *ApplianceDeviceCertificateAllOf) UnsetCompletedPhases()`

UnsetCompletedPhases ensures that no value is present for CompletedPhases, not even an explicit nil
### GetConfigurationMoId

`func (o *ApplianceDeviceCertificateAllOf) GetConfigurationMoId() string`

GetConfigurationMoId returns the ConfigurationMoId field if non-nil, zero value otherwise.

### GetConfigurationMoIdOk

`func (o *ApplianceDeviceCertificateAllOf) GetConfigurationMoIdOk() (*string, bool)`

GetConfigurationMoIdOk returns a tuple with the ConfigurationMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationMoId

`func (o *ApplianceDeviceCertificateAllOf) SetConfigurationMoId(v string)`

SetConfigurationMoId sets ConfigurationMoId field to given value.

### HasConfigurationMoId

`func (o *ApplianceDeviceCertificateAllOf) HasConfigurationMoId() bool`

HasConfigurationMoId returns a boolean if a field has been set.

### GetCurrentPhase

`func (o *ApplianceDeviceCertificateAllOf) GetCurrentPhase() ApplianceCertRenewalPhase`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *ApplianceDeviceCertificateAllOf) GetCurrentPhaseOk() (*ApplianceCertRenewalPhase, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *ApplianceDeviceCertificateAllOf) SetCurrentPhase(v ApplianceCertRenewalPhase)`

SetCurrentPhase sets CurrentPhase field to given value.

### HasCurrentPhase

`func (o *ApplianceDeviceCertificateAllOf) HasCurrentPhase() bool`

HasCurrentPhase returns a boolean if a field has been set.

### SetCurrentPhaseNil

`func (o *ApplianceDeviceCertificateAllOf) SetCurrentPhaseNil(b bool)`

 SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil

### UnsetCurrentPhase
`func (o *ApplianceDeviceCertificateAllOf) UnsetCurrentPhase()`

UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
### GetEndTime

`func (o *ApplianceDeviceCertificateAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceDeviceCertificateAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceDeviceCertificateAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceDeviceCertificateAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLastSuccessPollTime

`func (o *ApplianceDeviceCertificateAllOf) GetLastSuccessPollTime() time.Time`

GetLastSuccessPollTime returns the LastSuccessPollTime field if non-nil, zero value otherwise.

### GetLastSuccessPollTimeOk

`func (o *ApplianceDeviceCertificateAllOf) GetLastSuccessPollTimeOk() (*time.Time, bool)`

GetLastSuccessPollTimeOk returns a tuple with the LastSuccessPollTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessPollTime

`func (o *ApplianceDeviceCertificateAllOf) SetLastSuccessPollTime(v time.Time)`

SetLastSuccessPollTime sets LastSuccessPollTime field to given value.

### HasLastSuccessPollTime

`func (o *ApplianceDeviceCertificateAllOf) HasLastSuccessPollTime() bool`

HasLastSuccessPollTime returns a boolean if a field has been set.

### GetMessages

`func (o *ApplianceDeviceCertificateAllOf) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ApplianceDeviceCertificateAllOf) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ApplianceDeviceCertificateAllOf) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ApplianceDeviceCertificateAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ApplianceDeviceCertificateAllOf) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ApplianceDeviceCertificateAllOf) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetStartTime

`func (o *ApplianceDeviceCertificateAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceDeviceCertificateAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceDeviceCertificateAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceDeviceCertificateAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceDeviceCertificateAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceDeviceCertificateAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceDeviceCertificateAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceDeviceCertificateAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


