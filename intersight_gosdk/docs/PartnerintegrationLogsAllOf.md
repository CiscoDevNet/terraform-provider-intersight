# PartnerintegrationLogsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.Logs"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.Logs"]
**Stage** | Pointer to **string** | Stage in the build process these logs belong to. * &#x60;None&#x60; - Default value for the log stage. * &#x60;Backend&#x60; - Logs corresponding to backend build. * &#x60;Ui&#x60; - Logs corresponding to ui build stage. * &#x60;Apidocs&#x60; - Logs corresponding to the apidocs build stage. | [optional] [readonly] [default to "None"]
**Stderr** | Pointer to **[]string** |  | [optional] 
**Stdout** | Pointer to **[]string** |  | [optional] 
**Inventory** | Pointer to [**PartnerintegrationInventoryRelationship**](PartnerintegrationInventoryRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationLogsAllOf

`func NewPartnerintegrationLogsAllOf(classId string, objectType string, ) *PartnerintegrationLogsAllOf`

NewPartnerintegrationLogsAllOf instantiates a new PartnerintegrationLogsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationLogsAllOfWithDefaults

`func NewPartnerintegrationLogsAllOfWithDefaults() *PartnerintegrationLogsAllOf`

NewPartnerintegrationLogsAllOfWithDefaults instantiates a new PartnerintegrationLogsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationLogsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationLogsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationLogsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationLogsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationLogsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationLogsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStage

`func (o *PartnerintegrationLogsAllOf) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PartnerintegrationLogsAllOf) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PartnerintegrationLogsAllOf) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PartnerintegrationLogsAllOf) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStderr

`func (o *PartnerintegrationLogsAllOf) GetStderr() []string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *PartnerintegrationLogsAllOf) GetStderrOk() (*[]string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *PartnerintegrationLogsAllOf) SetStderr(v []string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *PartnerintegrationLogsAllOf) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### SetStderrNil

`func (o *PartnerintegrationLogsAllOf) SetStderrNil(b bool)`

 SetStderrNil sets the value for Stderr to be an explicit nil

### UnsetStderr
`func (o *PartnerintegrationLogsAllOf) UnsetStderr()`

UnsetStderr ensures that no value is present for Stderr, not even an explicit nil
### GetStdout

`func (o *PartnerintegrationLogsAllOf) GetStdout() []string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *PartnerintegrationLogsAllOf) GetStdoutOk() (*[]string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *PartnerintegrationLogsAllOf) SetStdout(v []string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *PartnerintegrationLogsAllOf) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *PartnerintegrationLogsAllOf) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *PartnerintegrationLogsAllOf) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil
### GetInventory

`func (o *PartnerintegrationLogsAllOf) GetInventory() PartnerintegrationInventoryRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PartnerintegrationLogsAllOf) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PartnerintegrationLogsAllOf) SetInventory(v PartnerintegrationInventoryRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *PartnerintegrationLogsAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


