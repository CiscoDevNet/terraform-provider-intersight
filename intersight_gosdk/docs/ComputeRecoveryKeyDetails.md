# ComputeRecoveryKeyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.RecoveryKeyDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.RecoveryKeyDetails"]
**RecoveryKeyEndPointState** | Pointer to **int32** | Recovery key state as per the last endpoint notification. * &#x60;0&#x60; - Migration key retrieval unsupported on the platform, typically due to incompatible server firmware. * &#x60;1&#x60; - Recovery key collection timed out. Review support documentation, fix any issue, and re-trigger Server Profile deployment. * &#x60;2&#x60; - Recovery key collection failed. Review support documentation, fix any issue, and re-trigger Server Profile deployment. * &#x60;3&#x60; - Recovery key rotation timed out. Review support documentation, fix any issue, and re-trigger Server Profile deployment. * &#x60;4&#x60; - Recovery key rotation failed. Review support documentation, fix any issue, and re-trigger Server Profile deployment. * &#x60;5&#x60; - Recovery key is not purged at enpoint. * &#x60;6&#x60; - Recovery key available for retrieval from the endpoint. | [optional] [readonly] [default to 0]
**RecoveryKeyInfo** | Pointer to [**NullableComputeMigrationKeyInfo**](ComputeMigrationKeyInfo.md) |  | [optional] 
**RecoveryKeyState** | Pointer to **string** | Recovery key state is deduced based on the endpoint info and its associated server profile. * &#x60;KeyNotApplicable&#x60; - Recovery key retrieval not applicable on this platform, typically because recovery key collection is not enabled. * &#x60;KeyNotSupported&#x60; - Recovery key retrieval not supported on this platform, typically because server firmware does not support key retrieval. * &#x60;KeyNotAvailable&#x60; - Recovery key has not been collected from end-point. Review support documentation for any issues. * &#x60;WaitingForKey&#x60; - Recovery key rotation in progress. The server boots with the old key, then stores the new key in Intersight. After completion, the state changes to KeyAvailable. * &#x60;KeyOutOfSync&#x60; - Server assignment changed after key availability, causing a recovery key mismatch. Run Server Profile Activate to start key rotation. * &#x60;KeyAvailable&#x60; - Recovery key successfully fetched and securely stored in Intersight. | [optional] [readonly] [default to "KeyNotApplicable"]

## Methods

### NewComputeRecoveryKeyDetails

`func NewComputeRecoveryKeyDetails(classId string, objectType string, ) *ComputeRecoveryKeyDetails`

NewComputeRecoveryKeyDetails instantiates a new ComputeRecoveryKeyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeRecoveryKeyDetailsWithDefaults

`func NewComputeRecoveryKeyDetailsWithDefaults() *ComputeRecoveryKeyDetails`

NewComputeRecoveryKeyDetailsWithDefaults instantiates a new ComputeRecoveryKeyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeRecoveryKeyDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeRecoveryKeyDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeRecoveryKeyDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeRecoveryKeyDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeRecoveryKeyDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeRecoveryKeyDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRecoveryKeyEndPointState

`func (o *ComputeRecoveryKeyDetails) GetRecoveryKeyEndPointState() int32`

GetRecoveryKeyEndPointState returns the RecoveryKeyEndPointState field if non-nil, zero value otherwise.

### GetRecoveryKeyEndPointStateOk

`func (o *ComputeRecoveryKeyDetails) GetRecoveryKeyEndPointStateOk() (*int32, bool)`

GetRecoveryKeyEndPointStateOk returns a tuple with the RecoveryKeyEndPointState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeyEndPointState

`func (o *ComputeRecoveryKeyDetails) SetRecoveryKeyEndPointState(v int32)`

SetRecoveryKeyEndPointState sets RecoveryKeyEndPointState field to given value.

### HasRecoveryKeyEndPointState

`func (o *ComputeRecoveryKeyDetails) HasRecoveryKeyEndPointState() bool`

HasRecoveryKeyEndPointState returns a boolean if a field has been set.

### GetRecoveryKeyInfo

`func (o *ComputeRecoveryKeyDetails) GetRecoveryKeyInfo() ComputeMigrationKeyInfo`

GetRecoveryKeyInfo returns the RecoveryKeyInfo field if non-nil, zero value otherwise.

### GetRecoveryKeyInfoOk

`func (o *ComputeRecoveryKeyDetails) GetRecoveryKeyInfoOk() (*ComputeMigrationKeyInfo, bool)`

GetRecoveryKeyInfoOk returns a tuple with the RecoveryKeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeyInfo

`func (o *ComputeRecoveryKeyDetails) SetRecoveryKeyInfo(v ComputeMigrationKeyInfo)`

SetRecoveryKeyInfo sets RecoveryKeyInfo field to given value.

### HasRecoveryKeyInfo

`func (o *ComputeRecoveryKeyDetails) HasRecoveryKeyInfo() bool`

HasRecoveryKeyInfo returns a boolean if a field has been set.

### SetRecoveryKeyInfoNil

`func (o *ComputeRecoveryKeyDetails) SetRecoveryKeyInfoNil(b bool)`

 SetRecoveryKeyInfoNil sets the value for RecoveryKeyInfo to be an explicit nil

### UnsetRecoveryKeyInfo
`func (o *ComputeRecoveryKeyDetails) UnsetRecoveryKeyInfo()`

UnsetRecoveryKeyInfo ensures that no value is present for RecoveryKeyInfo, not even an explicit nil
### GetRecoveryKeyState

`func (o *ComputeRecoveryKeyDetails) GetRecoveryKeyState() string`

GetRecoveryKeyState returns the RecoveryKeyState field if non-nil, zero value otherwise.

### GetRecoveryKeyStateOk

`func (o *ComputeRecoveryKeyDetails) GetRecoveryKeyStateOk() (*string, bool)`

GetRecoveryKeyStateOk returns a tuple with the RecoveryKeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeyState

`func (o *ComputeRecoveryKeyDetails) SetRecoveryKeyState(v string)`

SetRecoveryKeyState sets RecoveryKeyState field to given value.

### HasRecoveryKeyState

`func (o *ComputeRecoveryKeyDetails) HasRecoveryKeyState() bool`

HasRecoveryKeyState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


