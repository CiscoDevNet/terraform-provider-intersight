# TamTextFsmTemplateDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.TextFsmTemplateDataSource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.TextFsmTemplateDataSource"]
**Cmd** | Pointer to **string** | Command used to gather data needed to evaluate field notice applicability. | [optional] 

## Methods

### NewTamTextFsmTemplateDataSource

`func NewTamTextFsmTemplateDataSource(classId string, objectType string, ) *TamTextFsmTemplateDataSource`

NewTamTextFsmTemplateDataSource instantiates a new TamTextFsmTemplateDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamTextFsmTemplateDataSourceWithDefaults

`func NewTamTextFsmTemplateDataSourceWithDefaults() *TamTextFsmTemplateDataSource`

NewTamTextFsmTemplateDataSourceWithDefaults instantiates a new TamTextFsmTemplateDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamTextFsmTemplateDataSource) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamTextFsmTemplateDataSource) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamTextFsmTemplateDataSource) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamTextFsmTemplateDataSource) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamTextFsmTemplateDataSource) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamTextFsmTemplateDataSource) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCmd

`func (o *TamTextFsmTemplateDataSource) GetCmd() string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *TamTextFsmTemplateDataSource) GetCmdOk() (*string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *TamTextFsmTemplateDataSource) SetCmd(v string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *TamTextFsmTemplateDataSource) HasCmd() bool`

HasCmd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


