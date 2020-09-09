# ApplianceReleaseNoteAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to [**[]OnpremUpgradeNote**](onprem.UpgradeNote.md) |  | [optional] 
**Version** | Pointer to **string** | Version number of the pending upgrade. | [optional] [readonly] 

## Methods

### NewApplianceReleaseNoteAllOf

`func NewApplianceReleaseNoteAllOf() *ApplianceReleaseNoteAllOf`

NewApplianceReleaseNoteAllOf instantiates a new ApplianceReleaseNoteAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceReleaseNoteAllOfWithDefaults

`func NewApplianceReleaseNoteAllOfWithDefaults() *ApplianceReleaseNoteAllOf`

NewApplianceReleaseNoteAllOfWithDefaults instantiates a new ApplianceReleaseNoteAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *ApplianceReleaseNoteAllOf) GetNotes() []OnpremUpgradeNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApplianceReleaseNoteAllOf) GetNotesOk() (*[]OnpremUpgradeNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApplianceReleaseNoteAllOf) SetNotes(v []OnpremUpgradeNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApplianceReleaseNoteAllOf) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceReleaseNoteAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceReleaseNoteAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceReleaseNoteAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceReleaseNoteAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


