# ApplianceCertRenewalPhaseAllOf

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

### NewApplianceCertRenewalPhaseAllOf

`func NewApplianceCertRenewalPhaseAllOf(classId string, objectType string, ) *ApplianceCertRenewalPhaseAllOf`

NewApplianceCertRenewalPhaseAllOf instantiates a new ApplianceCertRenewalPhaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceCertRenewalPhaseAllOfWithDefaults

`func NewApplianceCertRenewalPhaseAllOfWithDefaults() *ApplianceCertRenewalPhaseAllOf`

NewApplianceCertRenewalPhaseAllOfWithDefaults instantiates a new ApplianceCertRenewalPhaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceCertRenewalPhaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceCertRenewalPhaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceCertRenewalPhaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceCertRenewalPhaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceCertRenewalPhaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceCertRenewalPhaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndTime

`func (o *ApplianceCertRenewalPhaseAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceCertRenewalPhaseAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceCertRenewalPhaseAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceCertRenewalPhaseAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailed

`func (o *ApplianceCertRenewalPhaseAllOf) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ApplianceCertRenewalPhaseAllOf) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ApplianceCertRenewalPhaseAllOf) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ApplianceCertRenewalPhaseAllOf) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetMessage

`func (o *ApplianceCertRenewalPhaseAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApplianceCertRenewalPhaseAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApplianceCertRenewalPhaseAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApplianceCertRenewalPhaseAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *ApplianceCertRenewalPhaseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplianceCertRenewalPhaseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplianceCertRenewalPhaseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplianceCertRenewalPhaseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *ApplianceCertRenewalPhaseAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceCertRenewalPhaseAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceCertRenewalPhaseAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceCertRenewalPhaseAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


