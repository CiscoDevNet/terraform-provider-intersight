# ApplianceReleaseNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ReleaseNote"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ReleaseNote"]
**Notes** | Pointer to [**[]OnpremUpgradeNote**](OnpremUpgradeNote.md) |  | [optional] 
**Version** | Pointer to **string** | Version number of the pending upgrade. | [optional] [readonly] 

## Methods

### NewApplianceReleaseNote

`func NewApplianceReleaseNote(classId string, objectType string, ) *ApplianceReleaseNote`

NewApplianceReleaseNote instantiates a new ApplianceReleaseNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceReleaseNoteWithDefaults

`func NewApplianceReleaseNoteWithDefaults() *ApplianceReleaseNote`

NewApplianceReleaseNoteWithDefaults instantiates a new ApplianceReleaseNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceReleaseNote) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceReleaseNote) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceReleaseNote) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceReleaseNote) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceReleaseNote) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceReleaseNote) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNotes

`func (o *ApplianceReleaseNote) GetNotes() []OnpremUpgradeNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApplianceReleaseNote) GetNotesOk() (*[]OnpremUpgradeNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApplianceReleaseNote) SetNotes(v []OnpremUpgradeNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApplianceReleaseNote) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ApplianceReleaseNote) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ApplianceReleaseNote) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetVersion

`func (o *ApplianceReleaseNote) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceReleaseNote) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceReleaseNote) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceReleaseNote) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


