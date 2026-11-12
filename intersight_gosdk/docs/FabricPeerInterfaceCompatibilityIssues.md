# FabricPeerInterfaceCompatibilityIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.PeerInterfaceCompatibilityIssues"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.PeerInterfaceCompatibilityIssues"]
**EvalState** | Pointer to **string** | Freshness state for the latest peer interface compatibility evaluation. * &#x60;None&#x60; - No peer interface compatibility evaluation has been requested for this profile. * &#x60;InProgress&#x60; - A peer interface compatibility evaluation has been requested and the latest result has not been applied yet. * &#x60;Completed&#x60; - The stored peer interface compatibility issues were produced by the latest evaluation. * &#x60;Failed&#x60; - The latest peer interface compatibility evaluation did not produce a report. * &#x60;NoChange&#x60; - The latest evaluation determined there were no configuration changes affecting peer interface compatibility, so the previously stored issues remain valid. | [optional] [readonly] [default to "None"]
**FailureReason** | Pointer to **string** | Human-readable reason describing why the latest peer interface compatibility evaluation failed. Only populated when evalState is Failed. | [optional] [readonly] 
**Issues** | Pointer to [**[]FabricPeerInterfaceCompatibilityIssue**](FabricPeerInterfaceCompatibilityIssue.md) |  | [optional] 

## Methods

### NewFabricPeerInterfaceCompatibilityIssues

`func NewFabricPeerInterfaceCompatibilityIssues(classId string, objectType string, ) *FabricPeerInterfaceCompatibilityIssues`

NewFabricPeerInterfaceCompatibilityIssues instantiates a new FabricPeerInterfaceCompatibilityIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricPeerInterfaceCompatibilityIssuesWithDefaults

`func NewFabricPeerInterfaceCompatibilityIssuesWithDefaults() *FabricPeerInterfaceCompatibilityIssues`

NewFabricPeerInterfaceCompatibilityIssuesWithDefaults instantiates a new FabricPeerInterfaceCompatibilityIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricPeerInterfaceCompatibilityIssues) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricPeerInterfaceCompatibilityIssues) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricPeerInterfaceCompatibilityIssues) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricPeerInterfaceCompatibilityIssues) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricPeerInterfaceCompatibilityIssues) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricPeerInterfaceCompatibilityIssues) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEvalState

`func (o *FabricPeerInterfaceCompatibilityIssues) GetEvalState() string`

GetEvalState returns the EvalState field if non-nil, zero value otherwise.

### GetEvalStateOk

`func (o *FabricPeerInterfaceCompatibilityIssues) GetEvalStateOk() (*string, bool)`

GetEvalStateOk returns a tuple with the EvalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalState

`func (o *FabricPeerInterfaceCompatibilityIssues) SetEvalState(v string)`

SetEvalState sets EvalState field to given value.

### HasEvalState

`func (o *FabricPeerInterfaceCompatibilityIssues) HasEvalState() bool`

HasEvalState returns a boolean if a field has been set.

### GetFailureReason

`func (o *FabricPeerInterfaceCompatibilityIssues) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *FabricPeerInterfaceCompatibilityIssues) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *FabricPeerInterfaceCompatibilityIssues) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *FabricPeerInterfaceCompatibilityIssues) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetIssues

`func (o *FabricPeerInterfaceCompatibilityIssues) GetIssues() []FabricPeerInterfaceCompatibilityIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *FabricPeerInterfaceCompatibilityIssues) GetIssuesOk() (*[]FabricPeerInterfaceCompatibilityIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *FabricPeerInterfaceCompatibilityIssues) SetIssues(v []FabricPeerInterfaceCompatibilityIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *FabricPeerInterfaceCompatibilityIssues) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### SetIssuesNil

`func (o *FabricPeerInterfaceCompatibilityIssues) SetIssuesNil(b bool)`

 SetIssuesNil sets the value for Issues to be an explicit nil

### UnsetIssues
`func (o *FabricPeerInterfaceCompatibilityIssues) UnsetIssues()`

UnsetIssues ensures that no value is present for Issues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


