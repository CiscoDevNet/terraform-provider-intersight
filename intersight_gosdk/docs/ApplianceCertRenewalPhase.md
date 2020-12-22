# ApplianceCertRenewalPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.CertRenewalPhase"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.CertRenewalPhase"]
**EndTime** | Pointer to **time.Time** | End date of the cert renewal phase. | [optional] [readonly] 
**Failed** | Pointer to **bool** | Indicates if the cert renewal phase has failed or not. | [optional] [readonly] 
**Message** | Pointer to **string** | Status message set during the cert renewal phase. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the cert renewal phase phase. * &#x60;Init&#x60; - New certificate detected, cleanup the old process if any running. * &#x60;ScheduleCertificateAddOperation&#x60; - Certificate Add Operation Schedulled. * &#x60;WaitForCertCollectionByEndpoint&#x60; - Monitor cert collection by endpoint. * &#x60;Success&#x60; - Certificate Renewal Task Success. * &#x60;Fail&#x60; - Certificate Renewal Task Fail. * &#x60;Cancel&#x60; - Certificate Renewal Task Cancel. | [optional] [readonly] [default to "Init"]
**StartTime** | Pointer to **time.Time** | Start date of the cert renewal phase. | [optional] [readonly] 

## Methods

### NewApplianceCertRenewalPhase

`func NewApplianceCertRenewalPhase(classId string, objectType string, ) *ApplianceCertRenewalPhase`

NewApplianceCertRenewalPhase instantiates a new ApplianceCertRenewalPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceCertRenewalPhaseWithDefaults

`func NewApplianceCertRenewalPhaseWithDefaults() *ApplianceCertRenewalPhase`

NewApplianceCertRenewalPhaseWithDefaults instantiates a new ApplianceCertRenewalPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceCertRenewalPhase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceCertRenewalPhase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceCertRenewalPhase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceCertRenewalPhase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceCertRenewalPhase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceCertRenewalPhase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndTime

`func (o *ApplianceCertRenewalPhase) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceCertRenewalPhase) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceCertRenewalPhase) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceCertRenewalPhase) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailed

`func (o *ApplianceCertRenewalPhase) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ApplianceCertRenewalPhase) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ApplianceCertRenewalPhase) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ApplianceCertRenewalPhase) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetMessage

`func (o *ApplianceCertRenewalPhase) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApplianceCertRenewalPhase) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApplianceCertRenewalPhase) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApplianceCertRenewalPhase) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *ApplianceCertRenewalPhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplianceCertRenewalPhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplianceCertRenewalPhase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplianceCertRenewalPhase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *ApplianceCertRenewalPhase) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceCertRenewalPhase) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceCertRenewalPhase) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceCertRenewalPhase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


