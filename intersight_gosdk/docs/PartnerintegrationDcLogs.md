# PartnerintegrationDcLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.DcLogs"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.DcLogs"]
**Stage** | Pointer to **string** | Stage in the build process these logs belong to. * &#x60;None&#x60; - Default value for the log stage. * &#x60;Backend&#x60; - Logs corresponding to backend build. * &#x60;Ui&#x60; - Logs corresponding to ui build stage. * &#x60;Apidocs&#x60; - Logs corresponding to the apidocs build stage. * &#x60;MetricsCollectorBackend&#x60; - Logs corresponding to Metrics Collector backend build. * &#x60;MetricsCollectorDependentBackend&#x60; - Logs corresponding to Metrics Collector backend build. * &#x60;MetricsCollectorUi&#x60; - Logs corresponding to Metrics Collector ui build stage. | [optional] [readonly] [default to "None"]
**Stderr** | Pointer to **[]string** |  | [optional] 
**Stdout** | Pointer to **[]string** |  | [optional] 
**DeviceConnector** | Pointer to [**NullablePartnerintegrationDeviceConnectorRelationship**](PartnerintegrationDeviceConnectorRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationDcLogs

`func NewPartnerintegrationDcLogs(classId string, objectType string, ) *PartnerintegrationDcLogs`

NewPartnerintegrationDcLogs instantiates a new PartnerintegrationDcLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationDcLogsWithDefaults

`func NewPartnerintegrationDcLogsWithDefaults() *PartnerintegrationDcLogs`

NewPartnerintegrationDcLogsWithDefaults instantiates a new PartnerintegrationDcLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationDcLogs) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationDcLogs) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationDcLogs) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationDcLogs) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationDcLogs) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationDcLogs) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStage

`func (o *PartnerintegrationDcLogs) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PartnerintegrationDcLogs) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PartnerintegrationDcLogs) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PartnerintegrationDcLogs) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStderr

`func (o *PartnerintegrationDcLogs) GetStderr() []string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *PartnerintegrationDcLogs) GetStderrOk() (*[]string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *PartnerintegrationDcLogs) SetStderr(v []string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *PartnerintegrationDcLogs) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### SetStderrNil

`func (o *PartnerintegrationDcLogs) SetStderrNil(b bool)`

 SetStderrNil sets the value for Stderr to be an explicit nil

### UnsetStderr
`func (o *PartnerintegrationDcLogs) UnsetStderr()`

UnsetStderr ensures that no value is present for Stderr, not even an explicit nil
### GetStdout

`func (o *PartnerintegrationDcLogs) GetStdout() []string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *PartnerintegrationDcLogs) GetStdoutOk() (*[]string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *PartnerintegrationDcLogs) SetStdout(v []string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *PartnerintegrationDcLogs) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *PartnerintegrationDcLogs) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *PartnerintegrationDcLogs) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil
### GetDeviceConnector

`func (o *PartnerintegrationDcLogs) GetDeviceConnector() PartnerintegrationDeviceConnectorRelationship`

GetDeviceConnector returns the DeviceConnector field if non-nil, zero value otherwise.

### GetDeviceConnectorOk

`func (o *PartnerintegrationDcLogs) GetDeviceConnectorOk() (*PartnerintegrationDeviceConnectorRelationship, bool)`

GetDeviceConnectorOk returns a tuple with the DeviceConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConnector

`func (o *PartnerintegrationDcLogs) SetDeviceConnector(v PartnerintegrationDeviceConnectorRelationship)`

SetDeviceConnector sets DeviceConnector field to given value.

### HasDeviceConnector

`func (o *PartnerintegrationDcLogs) HasDeviceConnector() bool`

HasDeviceConnector returns a boolean if a field has been set.

### SetDeviceConnectorNil

`func (o *PartnerintegrationDcLogs) SetDeviceConnectorNil(b bool)`

 SetDeviceConnectorNil sets the value for DeviceConnector to be an explicit nil

### UnsetDeviceConnector
`func (o *PartnerintegrationDcLogs) UnsetDeviceConnector()`

UnsetDeviceConnector ensures that no value is present for DeviceConnector, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


