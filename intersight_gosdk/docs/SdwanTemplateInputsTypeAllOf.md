# SdwanTemplateInputsTypeAllOf

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

### NewSdwanTemplateInputsTypeAllOf

`func NewSdwanTemplateInputsTypeAllOf(classId string, objectType string, ) *SdwanTemplateInputsTypeAllOf`

NewSdwanTemplateInputsTypeAllOf instantiates a new SdwanTemplateInputsTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanTemplateInputsTypeAllOfWithDefaults

`func NewSdwanTemplateInputsTypeAllOfWithDefaults() *SdwanTemplateInputsTypeAllOf`

NewSdwanTemplateInputsTypeAllOfWithDefaults instantiates a new SdwanTemplateInputsTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdwanTemplateInputsTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanTemplateInputsTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanTemplateInputsTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdwanTemplateInputsTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanTemplateInputsTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanTemplateInputsTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEditable

`func (o *SdwanTemplateInputsTypeAllOf) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SdwanTemplateInputsTypeAllOf) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SdwanTemplateInputsTypeAllOf) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *SdwanTemplateInputsTypeAllOf) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetKey

`func (o *SdwanTemplateInputsTypeAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SdwanTemplateInputsTypeAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SdwanTemplateInputsTypeAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SdwanTemplateInputsTypeAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRequired

`func (o *SdwanTemplateInputsTypeAllOf) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SdwanTemplateInputsTypeAllOf) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SdwanTemplateInputsTypeAllOf) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SdwanTemplateInputsTypeAllOf) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTemplate

`func (o *SdwanTemplateInputsTypeAllOf) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SdwanTemplateInputsTypeAllOf) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SdwanTemplateInputsTypeAllOf) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SdwanTemplateInputsTypeAllOf) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *SdwanTemplateInputsTypeAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SdwanTemplateInputsTypeAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SdwanTemplateInputsTypeAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SdwanTemplateInputsTypeAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *SdwanTemplateInputsTypeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SdwanTemplateInputsTypeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SdwanTemplateInputsTypeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SdwanTemplateInputsTypeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SdwanTemplateInputsTypeAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SdwanTemplateInputsTypeAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SdwanTemplateInputsTypeAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SdwanTemplateInputsTypeAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


