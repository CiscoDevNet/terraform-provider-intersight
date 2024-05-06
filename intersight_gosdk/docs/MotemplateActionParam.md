# MotemplateActionParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "motemplate.ActionParam"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "motemplate.ActionParam"]
**Name** | Pointer to **string** | The action parameter identifier. The supported values are SyncType and SyncTimer for the template sync action. * &#x60;None&#x60; - The default parameter that implies that no action parameter is required for the template action. * &#x60;SyncType&#x60; - The parameter that describes the type of sync action such as SyncAll, SyncOne or SyncFailed supported on any template or derived object. * &#x60;SyncTimer&#x60; - The parameter for the initial delay in seconds after which the sync action must be executed. The supported range is from 0 to 60 seconds. | [optional] [default to "None"]
**Value** | Pointer to **string** | The action parameter value is based on the action parameter type. Supported action parameters and their values are- a) Name - SyncType, Supported Values - SyncAll, SyncFailed, SyncOne. b) Name - SyncTimer, Supported Values - 0 to 60 seconds. | [optional] 

## Methods

### NewMotemplateActionParam

`func NewMotemplateActionParam(classId string, objectType string, ) *MotemplateActionParam`

NewMotemplateActionParam instantiates a new MotemplateActionParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotemplateActionParamWithDefaults

`func NewMotemplateActionParamWithDefaults() *MotemplateActionParam`

NewMotemplateActionParamWithDefaults instantiates a new MotemplateActionParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MotemplateActionParam) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MotemplateActionParam) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MotemplateActionParam) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MotemplateActionParam) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MotemplateActionParam) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MotemplateActionParam) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *MotemplateActionParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MotemplateActionParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MotemplateActionParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MotemplateActionParam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *MotemplateActionParam) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MotemplateActionParam) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MotemplateActionParam) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MotemplateActionParam) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


