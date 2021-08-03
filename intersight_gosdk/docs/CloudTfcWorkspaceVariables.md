# CloudTfcWorkspaceVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.TfcWorkspaceVariables"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.TfcWorkspaceVariables"]
**Category** | Pointer to **string** | Whether this is a Terraform environment variable. Valid values are \&quot;terraform\&quot; or \&quot;env\&quot;. | [optional] 
**Description** | Pointer to **string** | The description of the variable. | [optional] 
**Hcl** | Pointer to **bool** | Whether to evaluate the value of the variable as a string of HCL code. Has no effect for environment variables. | [optional] 
**Identity** | Pointer to **string** | The unique identifier for this variable. | [optional] 
**Key** | Pointer to **string** | The name of the variables. | [optional] 
**Sensitive** | Pointer to **bool** | Whether the value is sensitive. If true then variable is written once and not visible thereafter. | [optional] 
**Value** | Pointer to **string** | The value of the variables. | [optional] 

## Methods

### NewCloudTfcWorkspaceVariables

`func NewCloudTfcWorkspaceVariables(classId string, objectType string, ) *CloudTfcWorkspaceVariables`

NewCloudTfcWorkspaceVariables instantiates a new CloudTfcWorkspaceVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTfcWorkspaceVariablesWithDefaults

`func NewCloudTfcWorkspaceVariablesWithDefaults() *CloudTfcWorkspaceVariables`

NewCloudTfcWorkspaceVariablesWithDefaults instantiates a new CloudTfcWorkspaceVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudTfcWorkspaceVariables) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudTfcWorkspaceVariables) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudTfcWorkspaceVariables) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudTfcWorkspaceVariables) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudTfcWorkspaceVariables) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudTfcWorkspaceVariables) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *CloudTfcWorkspaceVariables) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudTfcWorkspaceVariables) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudTfcWorkspaceVariables) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudTfcWorkspaceVariables) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *CloudTfcWorkspaceVariables) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudTfcWorkspaceVariables) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudTfcWorkspaceVariables) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudTfcWorkspaceVariables) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHcl

`func (o *CloudTfcWorkspaceVariables) GetHcl() bool`

GetHcl returns the Hcl field if non-nil, zero value otherwise.

### GetHclOk

`func (o *CloudTfcWorkspaceVariables) GetHclOk() (*bool, bool)`

GetHclOk returns a tuple with the Hcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcl

`func (o *CloudTfcWorkspaceVariables) SetHcl(v bool)`

SetHcl sets Hcl field to given value.

### HasHcl

`func (o *CloudTfcWorkspaceVariables) HasHcl() bool`

HasHcl returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudTfcWorkspaceVariables) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudTfcWorkspaceVariables) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudTfcWorkspaceVariables) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudTfcWorkspaceVariables) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetKey

`func (o *CloudTfcWorkspaceVariables) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudTfcWorkspaceVariables) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudTfcWorkspaceVariables) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudTfcWorkspaceVariables) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSensitive

`func (o *CloudTfcWorkspaceVariables) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *CloudTfcWorkspaceVariables) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *CloudTfcWorkspaceVariables) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *CloudTfcWorkspaceVariables) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetValue

`func (o *CloudTfcWorkspaceVariables) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudTfcWorkspaceVariables) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudTfcWorkspaceVariables) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudTfcWorkspaceVariables) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


