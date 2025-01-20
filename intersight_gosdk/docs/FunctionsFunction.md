# FunctionsFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "functions.Function"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "functions.Function"]
**Action** | Pointer to **string** | Action of the function such as build, deploy, undeploy. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Publish&#x60; - Publish a Function that was saved or built. | [optional] [default to "None"]
**Code** | Pointer to **string** | Custom function code to create the first function version, mandatory in function creation payload. | [optional] 
**CreateUser** | Pointer to **string** | The user identifier who created the Function. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the function. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the function. Display name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**ModUser** | Pointer to **string** | The user identifier who last updated the Function. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the function. Name can only contain letters (a-z), numbers (0-9), hyphen (-). | [optional] 
**RuntimeMoid** | Pointer to **string** | Moid of runtime which is used to create the first function version, mandatory in function creation payload. | [optional] 
**Version** | Pointer to **int64** | The target version of the function, which is needed by action. | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFunctionsFunction

`func NewFunctionsFunction(classId string, objectType string, ) *FunctionsFunction`

NewFunctionsFunction instantiates a new FunctionsFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsFunctionWithDefaults

`func NewFunctionsFunctionWithDefaults() *FunctionsFunction`

NewFunctionsFunctionWithDefaults instantiates a new FunctionsFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FunctionsFunction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FunctionsFunction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FunctionsFunction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FunctionsFunction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FunctionsFunction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FunctionsFunction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *FunctionsFunction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FunctionsFunction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FunctionsFunction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FunctionsFunction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCode

`func (o *FunctionsFunction) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FunctionsFunction) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FunctionsFunction) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FunctionsFunction) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateUser

`func (o *FunctionsFunction) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *FunctionsFunction) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *FunctionsFunction) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *FunctionsFunction) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDescription

`func (o *FunctionsFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionsFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionsFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionsFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *FunctionsFunction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FunctionsFunction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FunctionsFunction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FunctionsFunction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetModUser

`func (o *FunctionsFunction) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *FunctionsFunction) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *FunctionsFunction) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *FunctionsFunction) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *FunctionsFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionsFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuntimeMoid

`func (o *FunctionsFunction) GetRuntimeMoid() string`

GetRuntimeMoid returns the RuntimeMoid field if non-nil, zero value otherwise.

### GetRuntimeMoidOk

`func (o *FunctionsFunction) GetRuntimeMoidOk() (*string, bool)`

GetRuntimeMoidOk returns a tuple with the RuntimeMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMoid

`func (o *FunctionsFunction) SetRuntimeMoid(v string)`

SetRuntimeMoid sets RuntimeMoid field to given value.

### HasRuntimeMoid

`func (o *FunctionsFunction) HasRuntimeMoid() bool`

HasRuntimeMoid returns a boolean if a field has been set.

### GetVersion

`func (o *FunctionsFunction) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunctionsFunction) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunctionsFunction) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunctionsFunction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrganization

`func (o *FunctionsFunction) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FunctionsFunction) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FunctionsFunction) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FunctionsFunction) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FunctionsFunction) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FunctionsFunction) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


