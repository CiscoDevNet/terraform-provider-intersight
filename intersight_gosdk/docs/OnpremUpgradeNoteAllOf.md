# OnpremUpgradeNoteAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.UpgradeNote"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.UpgradeNote"]
**Message** | Pointer to **string** | The change description, such as explanations of a new feature or defect resolution. | [optional] [readonly] 

## Methods

### NewOnpremUpgradeNoteAllOf

`func NewOnpremUpgradeNoteAllOf(classId string, objectType string, ) *OnpremUpgradeNoteAllOf`

NewOnpremUpgradeNoteAllOf instantiates a new OnpremUpgradeNoteAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremUpgradeNoteAllOfWithDefaults

`func NewOnpremUpgradeNoteAllOfWithDefaults() *OnpremUpgradeNoteAllOf`

NewOnpremUpgradeNoteAllOfWithDefaults instantiates a new OnpremUpgradeNoteAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremUpgradeNoteAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremUpgradeNoteAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremUpgradeNoteAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremUpgradeNoteAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremUpgradeNoteAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremUpgradeNoteAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *OnpremUpgradeNoteAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnpremUpgradeNoteAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnpremUpgradeNoteAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnpremUpgradeNoteAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


