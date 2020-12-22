# SdwanTemplateInputsType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdwan.TemplateInputsType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdwan.TemplateInputsType"]
**Editable** | Pointer to **bool** | Defines if the input is editable. | [optional] [default to false]
**Key** | Pointer to **string** | Name of the dynamic input key specified in the vManage template. | [optional] 
**Required** | Pointer to **bool** | Defines if the input is optional or required. | [optional] [default to false]
**Template** | Pointer to **string** | Refers to the name of the vManage template that this inputs belongs to. | [optional] [readonly] 
**Title** | Pointer to **string** | Label for the property being saved in the current instance of template Input. | [optional] 
**Type** | Pointer to **string** | Defines the object type for the input. | [optional] [default to "string"]
**Value** | Pointer to **string** | Value of the dynamic input key specfied in the vManage template. | [optional] 

## Methods

### NewSdwanTemplateInputsType

`func NewSdwanTemplateInputsType(classId string, objectType string, ) *SdwanTemplateInputsType`

NewSdwanTemplateInputsType instantiates a new SdwanTemplateInputsType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanTemplateInputsTypeWithDefaults

`func NewSdwanTemplateInputsTypeWithDefaults() *SdwanTemplateInputsType`

NewSdwanTemplateInputsTypeWithDefaults instantiates a new SdwanTemplateInputsType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdwanTemplateInputsType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanTemplateInputsType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanTemplateInputsType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdwanTemplateInputsType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanTemplateInputsType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanTemplateInputsType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEditable

`func (o *SdwanTemplateInputsType) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SdwanTemplateInputsType) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SdwanTemplateInputsType) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *SdwanTemplateInputsType) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetKey

`func (o *SdwanTemplateInputsType) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SdwanTemplateInputsType) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SdwanTemplateInputsType) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SdwanTemplateInputsType) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRequired

`func (o *SdwanTemplateInputsType) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SdwanTemplateInputsType) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SdwanTemplateInputsType) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SdwanTemplateInputsType) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTemplate

`func (o *SdwanTemplateInputsType) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SdwanTemplateInputsType) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SdwanTemplateInputsType) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SdwanTemplateInputsType) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *SdwanTemplateInputsType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SdwanTemplateInputsType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SdwanTemplateInputsType) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SdwanTemplateInputsType) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *SdwanTemplateInputsType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SdwanTemplateInputsType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SdwanTemplateInputsType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SdwanTemplateInputsType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SdwanTemplateInputsType) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SdwanTemplateInputsType) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SdwanTemplateInputsType) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SdwanTemplateInputsType) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


