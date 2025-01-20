# FunctionsRuntime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "functions.Runtime"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "functions.Runtime"]
**Action** | Pointer to **string** | Action against the Runtime. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Disable&#x60; - Disable an instance of a Runtime. * &#x60;Deprecate&#x60; - Deprecate an instance of a Runtime. * &#x60;FlagInsecure&#x60; - Flag an instance of a Runtime as insecure. | [optional] [default to "None"]
**CodeFileName** | Pointer to **string** | Name of file containing function source code. | [optional] 
**CodeTemplate** | Pointer to **string** | Template to guide on how to compose code. | [optional] 
**Components** | Pointer to [**[]FunctionsRuntimeComponent**](FunctionsRuntimeComponent.md) |  | [optional] 
**CreateUser** | Pointer to **string** | The user identifier who created the language runtime. | [optional] [readonly] 
**Deprecated** | Pointer to **bool** | Indicate if this language runtime is deprecated. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the language runtime. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the language runtime. Display name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Enabled** | Pointer to **bool** | Indicate if this language runtime is enabled. If not, the runtime is not usable at all. | [optional] [readonly] 
**Insecure** | Pointer to **bool** | Indicate if this language runtime is insecure due to vulnerabilities. | [optional] [readonly] 
**LanguageName** | Pointer to **string** | The official name of the programming language. * &#x60;Python&#x60; - Python programming language. | [optional] [default to "Python"]
**LanguageVersion** | Pointer to **string** | The version of the programming language. | [optional] 
**ModUser** | Pointer to **string** | The user identifier who last updated the language runtime. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the language runtime. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**Note** | Pointer to **string** | A note to bring user&#39;s attention to the status of this language runtime. | [optional] 
**RuntimeFilePath** | Pointer to **string** | Path to the runtime file. | [optional] [readonly] 
**RuntimeUploadMoid** | Pointer to **string** | Moid of Upload which represents the uploaded runtime file. | [optional] 
**Catalog** | Pointer to [**NullableWorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewFunctionsRuntime

`func NewFunctionsRuntime(classId string, objectType string, ) *FunctionsRuntime`

NewFunctionsRuntime instantiates a new FunctionsRuntime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsRuntimeWithDefaults

`func NewFunctionsRuntimeWithDefaults() *FunctionsRuntime`

NewFunctionsRuntimeWithDefaults instantiates a new FunctionsRuntime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FunctionsRuntime) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FunctionsRuntime) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FunctionsRuntime) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FunctionsRuntime) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FunctionsRuntime) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FunctionsRuntime) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *FunctionsRuntime) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FunctionsRuntime) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FunctionsRuntime) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FunctionsRuntime) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCodeFileName

`func (o *FunctionsRuntime) GetCodeFileName() string`

GetCodeFileName returns the CodeFileName field if non-nil, zero value otherwise.

### GetCodeFileNameOk

`func (o *FunctionsRuntime) GetCodeFileNameOk() (*string, bool)`

GetCodeFileNameOk returns a tuple with the CodeFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeFileName

`func (o *FunctionsRuntime) SetCodeFileName(v string)`

SetCodeFileName sets CodeFileName field to given value.

### HasCodeFileName

`func (o *FunctionsRuntime) HasCodeFileName() bool`

HasCodeFileName returns a boolean if a field has been set.

### GetCodeTemplate

`func (o *FunctionsRuntime) GetCodeTemplate() string`

GetCodeTemplate returns the CodeTemplate field if non-nil, zero value otherwise.

### GetCodeTemplateOk

`func (o *FunctionsRuntime) GetCodeTemplateOk() (*string, bool)`

GetCodeTemplateOk returns a tuple with the CodeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeTemplate

`func (o *FunctionsRuntime) SetCodeTemplate(v string)`

SetCodeTemplate sets CodeTemplate field to given value.

### HasCodeTemplate

`func (o *FunctionsRuntime) HasCodeTemplate() bool`

HasCodeTemplate returns a boolean if a field has been set.

### GetComponents

`func (o *FunctionsRuntime) GetComponents() []FunctionsRuntimeComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *FunctionsRuntime) GetComponentsOk() (*[]FunctionsRuntimeComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *FunctionsRuntime) SetComponents(v []FunctionsRuntimeComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *FunctionsRuntime) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *FunctionsRuntime) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *FunctionsRuntime) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetCreateUser

`func (o *FunctionsRuntime) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *FunctionsRuntime) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *FunctionsRuntime) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *FunctionsRuntime) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDeprecated

`func (o *FunctionsRuntime) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FunctionsRuntime) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FunctionsRuntime) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *FunctionsRuntime) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *FunctionsRuntime) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionsRuntime) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionsRuntime) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionsRuntime) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FunctionsRuntime) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FunctionsRuntime) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FunctionsRuntime) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FunctionsRuntime) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *FunctionsRuntime) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FunctionsRuntime) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FunctionsRuntime) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FunctionsRuntime) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInsecure

`func (o *FunctionsRuntime) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *FunctionsRuntime) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *FunctionsRuntime) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *FunctionsRuntime) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetLanguageName

`func (o *FunctionsRuntime) GetLanguageName() string`

GetLanguageName returns the LanguageName field if non-nil, zero value otherwise.

### GetLanguageNameOk

`func (o *FunctionsRuntime) GetLanguageNameOk() (*string, bool)`

GetLanguageNameOk returns a tuple with the LanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageName

`func (o *FunctionsRuntime) SetLanguageName(v string)`

SetLanguageName sets LanguageName field to given value.

### HasLanguageName

`func (o *FunctionsRuntime) HasLanguageName() bool`

HasLanguageName returns a boolean if a field has been set.

### GetLanguageVersion

`func (o *FunctionsRuntime) GetLanguageVersion() string`

GetLanguageVersion returns the LanguageVersion field if non-nil, zero value otherwise.

### GetLanguageVersionOk

`func (o *FunctionsRuntime) GetLanguageVersionOk() (*string, bool)`

GetLanguageVersionOk returns a tuple with the LanguageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageVersion

`func (o *FunctionsRuntime) SetLanguageVersion(v string)`

SetLanguageVersion sets LanguageVersion field to given value.

### HasLanguageVersion

`func (o *FunctionsRuntime) HasLanguageVersion() bool`

HasLanguageVersion returns a boolean if a field has been set.

### GetModUser

`func (o *FunctionsRuntime) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *FunctionsRuntime) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *FunctionsRuntime) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *FunctionsRuntime) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *FunctionsRuntime) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsRuntime) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsRuntime) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionsRuntime) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *FunctionsRuntime) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FunctionsRuntime) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FunctionsRuntime) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *FunctionsRuntime) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRuntimeFilePath

`func (o *FunctionsRuntime) GetRuntimeFilePath() string`

GetRuntimeFilePath returns the RuntimeFilePath field if non-nil, zero value otherwise.

### GetRuntimeFilePathOk

`func (o *FunctionsRuntime) GetRuntimeFilePathOk() (*string, bool)`

GetRuntimeFilePathOk returns a tuple with the RuntimeFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeFilePath

`func (o *FunctionsRuntime) SetRuntimeFilePath(v string)`

SetRuntimeFilePath sets RuntimeFilePath field to given value.

### HasRuntimeFilePath

`func (o *FunctionsRuntime) HasRuntimeFilePath() bool`

HasRuntimeFilePath returns a boolean if a field has been set.

### GetRuntimeUploadMoid

`func (o *FunctionsRuntime) GetRuntimeUploadMoid() string`

GetRuntimeUploadMoid returns the RuntimeUploadMoid field if non-nil, zero value otherwise.

### GetRuntimeUploadMoidOk

`func (o *FunctionsRuntime) GetRuntimeUploadMoidOk() (*string, bool)`

GetRuntimeUploadMoidOk returns a tuple with the RuntimeUploadMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeUploadMoid

`func (o *FunctionsRuntime) SetRuntimeUploadMoid(v string)`

SetRuntimeUploadMoid sets RuntimeUploadMoid field to given value.

### HasRuntimeUploadMoid

`func (o *FunctionsRuntime) HasRuntimeUploadMoid() bool`

HasRuntimeUploadMoid returns a boolean if a field has been set.

### GetCatalog

`func (o *FunctionsRuntime) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FunctionsRuntime) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FunctionsRuntime) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FunctionsRuntime) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *FunctionsRuntime) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *FunctionsRuntime) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


