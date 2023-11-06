# PartnerintegrationDcLogsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.DcLogs"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.DcLogs"]
**Stage** | Pointer to **string** | Stage in the build process these logs belong to. * &#x60;None&#x60; - Default value for the log stage. * &#x60;Backend&#x60; - Logs corresponding to backend build. * &#x60;Ui&#x60; - Logs corresponding to ui build stage. * &#x60;Apidocs&#x60; - Logs corresponding to the apidocs build stage. | [optional] [readonly] [default to "None"]
**Stderr** | Pointer to **[]string** |  | [optional] 
**Stdout** | Pointer to **[]string** |  | [optional] 
**DeviceConnector** | Pointer to [**PartnerintegrationDeviceConnectorRelationship**](PartnerintegrationDeviceConnectorRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationDcLogsAllOf

`func NewPartnerintegrationDcLogsAllOf(classId string, objectType string, ) *PartnerintegrationDcLogsAllOf`

NewPartnerintegrationDcLogsAllOf instantiates a new PartnerintegrationDcLogsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationDcLogsAllOfWithDefaults

`func NewPartnerintegrationDcLogsAllOfWithDefaults() *PartnerintegrationDcLogsAllOf`

NewPartnerintegrationDcLogsAllOfWithDefaults instantiates a new PartnerintegrationDcLogsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationDcLogsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationDcLogsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationDcLogsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationDcLogsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationDcLogsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationDcLogsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStage

`func (o *PartnerintegrationDcLogsAllOf) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PartnerintegrationDcLogsAllOf) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PartnerintegrationDcLogsAllOf) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PartnerintegrationDcLogsAllOf) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStderr

`func (o *PartnerintegrationDcLogsAllOf) GetStderr() []string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *PartnerintegrationDcLogsAllOf) GetStderrOk() (*[]string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *PartnerintegrationDcLogsAllOf) SetStderr(v []string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *PartnerintegrationDcLogsAllOf) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### SetStderrNil

`func (o *PartnerintegrationDcLogsAllOf) SetStderrNil(b bool)`

 SetStderrNil sets the value for Stderr to be an explicit nil

### UnsetStderr
`func (o *PartnerintegrationDcLogsAllOf) UnsetStderr()`

UnsetStderr ensures that no value is present for Stderr, not even an explicit nil
### GetStdout

`func (o *PartnerintegrationDcLogsAllOf) GetStdout() []string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *PartnerintegrationDcLogsAllOf) GetStdoutOk() (*[]string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *PartnerintegrationDcLogsAllOf) SetStdout(v []string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *PartnerintegrationDcLogsAllOf) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *PartnerintegrationDcLogsAllOf) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *PartnerintegrationDcLogsAllOf) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil
### GetDeviceConnector

`func (o *PartnerintegrationDcLogsAllOf) GetDeviceConnector() PartnerintegrationDeviceConnectorRelationship`

GetDeviceConnector returns the DeviceConnector field if non-nil, zero value otherwise.

### GetDeviceConnectorOk

`func (o *PartnerintegrationDcLogsAllOf) GetDeviceConnectorOk() (*PartnerintegrationDeviceConnectorRelationship, bool)`

GetDeviceConnectorOk returns a tuple with the DeviceConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConnector

`func (o *PartnerintegrationDcLogsAllOf) SetDeviceConnector(v PartnerintegrationDeviceConnectorRelationship)`

SetDeviceConnector sets DeviceConnector field to given value.

### HasDeviceConnector

`func (o *PartnerintegrationDcLogsAllOf) HasDeviceConnector() bool`

HasDeviceConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


