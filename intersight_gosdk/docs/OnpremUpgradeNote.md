# OnpremUpgradeNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.UpgradeNote"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.UpgradeNote"]
**Message** | Pointer to **string** | The change description, such as explanations of a new feature or defect resolution. | [optional] [readonly] 

## Methods

### NewOnpremUpgradeNote

`func NewOnpremUpgradeNote(classId string, objectType string, ) *OnpremUpgradeNote`

NewOnpremUpgradeNote instantiates a new OnpremUpgradeNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremUpgradeNoteWithDefaults

`func NewOnpremUpgradeNoteWithDefaults() *OnpremUpgradeNote`

NewOnpremUpgradeNoteWithDefaults instantiates a new OnpremUpgradeNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremUpgradeNote) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremUpgradeNote) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremUpgradeNote) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremUpgradeNote) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremUpgradeNote) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremUpgradeNote) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *OnpremUpgradeNote) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnpremUpgradeNote) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnpremUpgradeNote) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnpremUpgradeNote) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


