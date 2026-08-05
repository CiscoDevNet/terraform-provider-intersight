# MgmtConfigObjectRestoreDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mgmt.ConfigObjectRestoreDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "mgmt.ConfigObjectRestoreDetails"]
**Description** | Pointer to **string** | Description of the configured object. | [optional] [readonly] 
**ImportIsIncomplete** | Pointer to **bool** | Import couldn&#39;t complete due to version incompatibility and user has to fix the error manually. | [optional] [readonly] 
**IncompletionReason** | Pointer to **string** | Reason for the incomplete import. | [optional] [readonly] 
**Message** | Pointer to **string** | Additional info about the restore status. | [optional] [readonly] 
**Moid** | Pointer to **string** | Moid of the configured object. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the configured object. | [optional] [readonly] 
**ObjectMeta** | Pointer to **interface{}** | Additional information about the type of configured object. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the config object restore operation. * &#x60;Pending&#x60; - Indicates that the request is yet to be processed. * &#x60;ObjPresenceCheckInProgress&#x60; - Indicates that the checking for object presence is in progress. * &#x60;ObjPresenceCheckInComplete&#x60; - Indicates that the request is being processed. * &#x60;ObjPresenceCheckFailed&#x60; - Indicates that the checking for object presence failed. * &#x60;Processing&#x60; - Indicates that the request is being processed. * &#x60;TimedOut&#x60; - Indicates that the request processing timed out. * &#x60;Failed&#x60; - Indicates that the request processing failed. * &#x60;Completed&#x60; - Indicates that the request processing is complete. * &#x60;Partial&#x60; - Indicates that the request is partially complete due to an error. * &#x60;Skipped&#x60; - Indicates that the request was skipped. | [optional] [readonly] [default to "Pending"]
**Type** | Pointer to **string** | ObjectType of the configured object. | [optional] [readonly] 

## Methods

### NewMgmtConfigObjectRestoreDetails

`func NewMgmtConfigObjectRestoreDetails(classId string, objectType string, ) *MgmtConfigObjectRestoreDetails`

NewMgmtConfigObjectRestoreDetails instantiates a new MgmtConfigObjectRestoreDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMgmtConfigObjectRestoreDetailsWithDefaults

`func NewMgmtConfigObjectRestoreDetailsWithDefaults() *MgmtConfigObjectRestoreDetails`

NewMgmtConfigObjectRestoreDetailsWithDefaults instantiates a new MgmtConfigObjectRestoreDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MgmtConfigObjectRestoreDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MgmtConfigObjectRestoreDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MgmtConfigObjectRestoreDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MgmtConfigObjectRestoreDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MgmtConfigObjectRestoreDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MgmtConfigObjectRestoreDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *MgmtConfigObjectRestoreDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MgmtConfigObjectRestoreDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MgmtConfigObjectRestoreDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MgmtConfigObjectRestoreDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImportIsIncomplete

`func (o *MgmtConfigObjectRestoreDetails) GetImportIsIncomplete() bool`

GetImportIsIncomplete returns the ImportIsIncomplete field if non-nil, zero value otherwise.

### GetImportIsIncompleteOk

`func (o *MgmtConfigObjectRestoreDetails) GetImportIsIncompleteOk() (*bool, bool)`

GetImportIsIncompleteOk returns a tuple with the ImportIsIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportIsIncomplete

`func (o *MgmtConfigObjectRestoreDetails) SetImportIsIncomplete(v bool)`

SetImportIsIncomplete sets ImportIsIncomplete field to given value.

### HasImportIsIncomplete

`func (o *MgmtConfigObjectRestoreDetails) HasImportIsIncomplete() bool`

HasImportIsIncomplete returns a boolean if a field has been set.

### GetIncompletionReason

`func (o *MgmtConfigObjectRestoreDetails) GetIncompletionReason() string`

GetIncompletionReason returns the IncompletionReason field if non-nil, zero value otherwise.

### GetIncompletionReasonOk

`func (o *MgmtConfigObjectRestoreDetails) GetIncompletionReasonOk() (*string, bool)`

GetIncompletionReasonOk returns a tuple with the IncompletionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompletionReason

`func (o *MgmtConfigObjectRestoreDetails) SetIncompletionReason(v string)`

SetIncompletionReason sets IncompletionReason field to given value.

### HasIncompletionReason

`func (o *MgmtConfigObjectRestoreDetails) HasIncompletionReason() bool`

HasIncompletionReason returns a boolean if a field has been set.

### GetMessage

`func (o *MgmtConfigObjectRestoreDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MgmtConfigObjectRestoreDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MgmtConfigObjectRestoreDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MgmtConfigObjectRestoreDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMoid

`func (o *MgmtConfigObjectRestoreDetails) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *MgmtConfigObjectRestoreDetails) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *MgmtConfigObjectRestoreDetails) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *MgmtConfigObjectRestoreDetails) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetName

`func (o *MgmtConfigObjectRestoreDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MgmtConfigObjectRestoreDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MgmtConfigObjectRestoreDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MgmtConfigObjectRestoreDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectMeta

`func (o *MgmtConfigObjectRestoreDetails) GetObjectMeta() interface{}`

GetObjectMeta returns the ObjectMeta field if non-nil, zero value otherwise.

### GetObjectMetaOk

`func (o *MgmtConfigObjectRestoreDetails) GetObjectMetaOk() (*interface{}, bool)`

GetObjectMetaOk returns a tuple with the ObjectMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMeta

`func (o *MgmtConfigObjectRestoreDetails) SetObjectMeta(v interface{})`

SetObjectMeta sets ObjectMeta field to given value.

### HasObjectMeta

`func (o *MgmtConfigObjectRestoreDetails) HasObjectMeta() bool`

HasObjectMeta returns a boolean if a field has been set.

### SetObjectMetaNil

`func (o *MgmtConfigObjectRestoreDetails) SetObjectMetaNil(b bool)`

 SetObjectMetaNil sets the value for ObjectMeta to be an explicit nil

### UnsetObjectMeta
`func (o *MgmtConfigObjectRestoreDetails) UnsetObjectMeta()`

UnsetObjectMeta ensures that no value is present for ObjectMeta, not even an explicit nil
### GetStatus

`func (o *MgmtConfigObjectRestoreDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MgmtConfigObjectRestoreDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MgmtConfigObjectRestoreDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MgmtConfigObjectRestoreDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MgmtConfigObjectRestoreDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MgmtConfigObjectRestoreDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MgmtConfigObjectRestoreDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MgmtConfigObjectRestoreDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


