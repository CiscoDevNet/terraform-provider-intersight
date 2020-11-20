# OsTemplateFileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.TemplateFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.TemplateFile"]
**Name** | Pointer to **string** | The name of the OS Template File that user uploads for unattended installation. | [optional] 
**Placeholders** | Pointer to **[]string** |  | [optional] 
**TemplateContent** | Pointer to **string** | The content of the entire template file is stored as value. The content can either be a static file content or a template content. The template is expected to conform to the golang template syntax.  The placeholders, if any, would be populated and the values provided would be  used to populate this template. | [optional] 

## Methods

### NewOsTemplateFileAllOf

`func NewOsTemplateFileAllOf(classId string, objectType string, ) *OsTemplateFileAllOf`

NewOsTemplateFileAllOf instantiates a new OsTemplateFileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsTemplateFileAllOfWithDefaults

`func NewOsTemplateFileAllOfWithDefaults() *OsTemplateFileAllOf`

NewOsTemplateFileAllOfWithDefaults instantiates a new OsTemplateFileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsTemplateFileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsTemplateFileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsTemplateFileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsTemplateFileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsTemplateFileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsTemplateFileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *OsTemplateFileAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsTemplateFileAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsTemplateFileAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsTemplateFileAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaceholders

`func (o *OsTemplateFileAllOf) GetPlaceholders() []string`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *OsTemplateFileAllOf) GetPlaceholdersOk() (*[]string, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *OsTemplateFileAllOf) SetPlaceholders(v []string)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *OsTemplateFileAllOf) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### SetPlaceholdersNil

`func (o *OsTemplateFileAllOf) SetPlaceholdersNil(b bool)`

 SetPlaceholdersNil sets the value for Placeholders to be an explicit nil

### UnsetPlaceholders
`func (o *OsTemplateFileAllOf) UnsetPlaceholders()`

UnsetPlaceholders ensures that no value is present for Placeholders, not even an explicit nil
### GetTemplateContent

`func (o *OsTemplateFileAllOf) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *OsTemplateFileAllOf) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *OsTemplateFileAllOf) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *OsTemplateFileAllOf) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


