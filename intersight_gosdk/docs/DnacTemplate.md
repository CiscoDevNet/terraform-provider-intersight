# DnacTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.Template"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.Template"]
**Composite** | Pointer to **string** | Value indicates the template is composite. | [optional] 
**ProjectId** | Pointer to **string** | Identity of the project template. | [optional] 
**ProjectName** | Pointer to **string** | Name of the project template. | [optional] 
**TemplateId** | Pointer to **string** | Unique identity of the template. | [optional] 
**TemplateName** | Pointer to **string** | Name assigned to the template. | [optional] 

## Methods

### NewDnacTemplate

`func NewDnacTemplate(classId string, objectType string, ) *DnacTemplate`

NewDnacTemplate instantiates a new DnacTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacTemplateWithDefaults

`func NewDnacTemplateWithDefaults() *DnacTemplate`

NewDnacTemplateWithDefaults instantiates a new DnacTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacTemplate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacTemplate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacTemplate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComposite

`func (o *DnacTemplate) GetComposite() string`

GetComposite returns the Composite field if non-nil, zero value otherwise.

### GetCompositeOk

`func (o *DnacTemplate) GetCompositeOk() (*string, bool)`

GetCompositeOk returns a tuple with the Composite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposite

`func (o *DnacTemplate) SetComposite(v string)`

SetComposite sets Composite field to given value.

### HasComposite

`func (o *DnacTemplate) HasComposite() bool`

HasComposite returns a boolean if a field has been set.

### GetProjectId

`func (o *DnacTemplate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DnacTemplate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DnacTemplate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DnacTemplate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *DnacTemplate) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DnacTemplate) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DnacTemplate) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *DnacTemplate) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetTemplateId

`func (o *DnacTemplate) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DnacTemplate) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DnacTemplate) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DnacTemplate) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *DnacTemplate) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *DnacTemplate) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *DnacTemplate) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *DnacTemplate) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


