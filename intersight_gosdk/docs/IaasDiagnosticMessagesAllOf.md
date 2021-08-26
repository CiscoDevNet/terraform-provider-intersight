# IaasDiagnosticMessagesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.DiagnosticMessages"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.DiagnosticMessages"]
**Category** | Pointer to **string** | Message category of the alerts. | [optional] [readonly] 
**Guid** | Pointer to **string** | Unique ID of UCS Director getting registerd with Intersight. | [optional] [readonly] 
**Item** | Pointer to **string** | Message target type of the alerts. | [optional] [readonly] 
**LastChecked** | Pointer to **string** | Last checked time of the alerts. | [optional] [readonly] 
**Message** | Pointer to **string** | Detailed info about the alert. | [optional] [readonly] 
**Recommendation** | Pointer to **string** | Recommended fix for the alert. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the given alert. | [optional] [readonly] 
**StatusId** | Pointer to **string** | Status Id of the given alert. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewIaasDiagnosticMessagesAllOf

`func NewIaasDiagnosticMessagesAllOf(classId string, objectType string, ) *IaasDiagnosticMessagesAllOf`

NewIaasDiagnosticMessagesAllOf instantiates a new IaasDiagnosticMessagesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasDiagnosticMessagesAllOfWithDefaults

`func NewIaasDiagnosticMessagesAllOfWithDefaults() *IaasDiagnosticMessagesAllOf`

NewIaasDiagnosticMessagesAllOfWithDefaults instantiates a new IaasDiagnosticMessagesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasDiagnosticMessagesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasDiagnosticMessagesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasDiagnosticMessagesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasDiagnosticMessagesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasDiagnosticMessagesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasDiagnosticMessagesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *IaasDiagnosticMessagesAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IaasDiagnosticMessagesAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IaasDiagnosticMessagesAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IaasDiagnosticMessagesAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGuid

`func (o *IaasDiagnosticMessagesAllOf) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasDiagnosticMessagesAllOf) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasDiagnosticMessagesAllOf) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasDiagnosticMessagesAllOf) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetItem

`func (o *IaasDiagnosticMessagesAllOf) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *IaasDiagnosticMessagesAllOf) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *IaasDiagnosticMessagesAllOf) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *IaasDiagnosticMessagesAllOf) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLastChecked

`func (o *IaasDiagnosticMessagesAllOf) GetLastChecked() string`

GetLastChecked returns the LastChecked field if non-nil, zero value otherwise.

### GetLastCheckedOk

`func (o *IaasDiagnosticMessagesAllOf) GetLastCheckedOk() (*string, bool)`

GetLastCheckedOk returns a tuple with the LastChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChecked

`func (o *IaasDiagnosticMessagesAllOf) SetLastChecked(v string)`

SetLastChecked sets LastChecked field to given value.

### HasLastChecked

`func (o *IaasDiagnosticMessagesAllOf) HasLastChecked() bool`

HasLastChecked returns a boolean if a field has been set.

### GetMessage

`func (o *IaasDiagnosticMessagesAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IaasDiagnosticMessagesAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IaasDiagnosticMessagesAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IaasDiagnosticMessagesAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecommendation

`func (o *IaasDiagnosticMessagesAllOf) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *IaasDiagnosticMessagesAllOf) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *IaasDiagnosticMessagesAllOf) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *IaasDiagnosticMessagesAllOf) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetStatus

`func (o *IaasDiagnosticMessagesAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IaasDiagnosticMessagesAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IaasDiagnosticMessagesAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IaasDiagnosticMessagesAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusId

`func (o *IaasDiagnosticMessagesAllOf) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *IaasDiagnosticMessagesAllOf) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *IaasDiagnosticMessagesAllOf) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *IaasDiagnosticMessagesAllOf) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *IaasDiagnosticMessagesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasDiagnosticMessagesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasDiagnosticMessagesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasDiagnosticMessagesAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


