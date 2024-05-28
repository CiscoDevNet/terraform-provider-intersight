# PartnerintegrationLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.Logs"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.Logs"]
**Stage** | Pointer to **string** | Stage in the build process these logs belong to. * &#x60;None&#x60; - Default value for the log stage. * &#x60;Backend&#x60; - Logs corresponding to backend build. * &#x60;Ui&#x60; - Logs corresponding to ui build stage. * &#x60;Apidocs&#x60; - Logs corresponding to the apidocs build stage. | [optional] [readonly] [default to "None"]
**Stderr** | Pointer to **[]string** |  | [optional] 
**Stdout** | Pointer to **[]string** |  | [optional] 
**Inventory** | Pointer to [**NullablePartnerintegrationInventoryRelationship**](PartnerintegrationInventoryRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationLogs

`func NewPartnerintegrationLogs(classId string, objectType string, ) *PartnerintegrationLogs`

NewPartnerintegrationLogs instantiates a new PartnerintegrationLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationLogsWithDefaults

`func NewPartnerintegrationLogsWithDefaults() *PartnerintegrationLogs`

NewPartnerintegrationLogsWithDefaults instantiates a new PartnerintegrationLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationLogs) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationLogs) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationLogs) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationLogs) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationLogs) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationLogs) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStage

`func (o *PartnerintegrationLogs) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PartnerintegrationLogs) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PartnerintegrationLogs) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PartnerintegrationLogs) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStderr

`func (o *PartnerintegrationLogs) GetStderr() []string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *PartnerintegrationLogs) GetStderrOk() (*[]string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *PartnerintegrationLogs) SetStderr(v []string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *PartnerintegrationLogs) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### SetStderrNil

`func (o *PartnerintegrationLogs) SetStderrNil(b bool)`

 SetStderrNil sets the value for Stderr to be an explicit nil

### UnsetStderr
`func (o *PartnerintegrationLogs) UnsetStderr()`

UnsetStderr ensures that no value is present for Stderr, not even an explicit nil
### GetStdout

`func (o *PartnerintegrationLogs) GetStdout() []string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *PartnerintegrationLogs) GetStdoutOk() (*[]string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *PartnerintegrationLogs) SetStdout(v []string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *PartnerintegrationLogs) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *PartnerintegrationLogs) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *PartnerintegrationLogs) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil
### GetInventory

`func (o *PartnerintegrationLogs) GetInventory() PartnerintegrationInventoryRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *PartnerintegrationLogs) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *PartnerintegrationLogs) SetInventory(v PartnerintegrationInventoryRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *PartnerintegrationLogs) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *PartnerintegrationLogs) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *PartnerintegrationLogs) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


