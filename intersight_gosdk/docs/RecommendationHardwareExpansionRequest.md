# RecommendationHardwareExpansionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.HardwareExpansionRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.HardwareExpansionRequest"]
**Action** | Pointer to **string** | Action to be triggered for computing recommendation. Default value is None. * &#x60;None&#x60; - The Enum value None represents that no action is triggered on the forecast Instance managed object. * &#x60;Evaluate&#x60; - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered. | [optional] [default to "None"]
**Message** | Pointer to **string** | User visible error message for the Hardware Expansion request. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the Hardware Expansion request. * &#x60;None&#x60; - The Enum value None represents the default status value before any API call is made. * &#x60;Success&#x60; - The Enum value Success represents that the API call returned with success. * &#x60;Fail&#x60; - The Enum value Fail represents that the API call returned with a failure. | [optional] [readonly] [default to "None"]
**HwExpansionRequestItems** | Pointer to [**[]RecommendationHardwareExpansionRequestItemRelationship**](RecommendationHardwareExpansionRequestItemRelationship.md) | An array of relationships to recommendationHardwareExpansionRequestItem resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewRecommendationHardwareExpansionRequest

`func NewRecommendationHardwareExpansionRequest(classId string, objectType string, ) *RecommendationHardwareExpansionRequest`

NewRecommendationHardwareExpansionRequest instantiates a new RecommendationHardwareExpansionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationHardwareExpansionRequestWithDefaults

`func NewRecommendationHardwareExpansionRequestWithDefaults() *RecommendationHardwareExpansionRequest`

NewRecommendationHardwareExpansionRequestWithDefaults instantiates a new RecommendationHardwareExpansionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationHardwareExpansionRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationHardwareExpansionRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationHardwareExpansionRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationHardwareExpansionRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationHardwareExpansionRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationHardwareExpansionRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *RecommendationHardwareExpansionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RecommendationHardwareExpansionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RecommendationHardwareExpansionRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RecommendationHardwareExpansionRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMessage

`func (o *RecommendationHardwareExpansionRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RecommendationHardwareExpansionRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RecommendationHardwareExpansionRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RecommendationHardwareExpansionRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *RecommendationHardwareExpansionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecommendationHardwareExpansionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecommendationHardwareExpansionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecommendationHardwareExpansionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHwExpansionRequestItems

`func (o *RecommendationHardwareExpansionRequest) GetHwExpansionRequestItems() []RecommendationHardwareExpansionRequestItemRelationship`

GetHwExpansionRequestItems returns the HwExpansionRequestItems field if non-nil, zero value otherwise.

### GetHwExpansionRequestItemsOk

`func (o *RecommendationHardwareExpansionRequest) GetHwExpansionRequestItemsOk() (*[]RecommendationHardwareExpansionRequestItemRelationship, bool)`

GetHwExpansionRequestItemsOk returns a tuple with the HwExpansionRequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwExpansionRequestItems

`func (o *RecommendationHardwareExpansionRequest) SetHwExpansionRequestItems(v []RecommendationHardwareExpansionRequestItemRelationship)`

SetHwExpansionRequestItems sets HwExpansionRequestItems field to given value.

### HasHwExpansionRequestItems

`func (o *RecommendationHardwareExpansionRequest) HasHwExpansionRequestItems() bool`

HasHwExpansionRequestItems returns a boolean if a field has been set.

### SetHwExpansionRequestItemsNil

`func (o *RecommendationHardwareExpansionRequest) SetHwExpansionRequestItemsNil(b bool)`

 SetHwExpansionRequestItemsNil sets the value for HwExpansionRequestItems to be an explicit nil

### UnsetHwExpansionRequestItems
`func (o *RecommendationHardwareExpansionRequest) UnsetHwExpansionRequestItems()`

UnsetHwExpansionRequestItems ensures that no value is present for HwExpansionRequestItems, not even an explicit nil
### GetRegisteredDevice

`func (o *RecommendationHardwareExpansionRequest) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *RecommendationHardwareExpansionRequest) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *RecommendationHardwareExpansionRequest) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *RecommendationHardwareExpansionRequest) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


