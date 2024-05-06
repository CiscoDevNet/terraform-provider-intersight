# MotemplateActionParamAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "motemplate.ActionParam"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "motemplate.ActionParam"]
**Name** | Pointer to **string** | The action parameter identifier. The supported values are SyncType and SyncTimer for the template sync action. * &#x60;None&#x60; - The default parameter that implies that no action parameter is required for the template action. * &#x60;SyncType&#x60; - The parameter that describes the type of sync action such as SyncAll, SyncOne or SyncFailed supported on any template or derived object. * &#x60;SyncTimer&#x60; - The parameter for the initial delay in seconds after which the sync action must be executed. The supported range is from 0 to 60 seconds. | [optional] [default to "None"]
**Value** | Pointer to **string** | The action parameter value is based on the action parameter type. Supported action parameters and their values are- a) Name - SyncType, Supported Values - SyncAll, SyncFailed, SyncOne. b) Name - SyncTimer, Supported Values - 0 to 60 seconds. | [optional] 

## Methods

### NewMotemplateActionParamAllOf

`func NewMotemplateActionParamAllOf(classId string, objectType string, ) *MotemplateActionParamAllOf`

NewMotemplateActionParamAllOf instantiates a new MotemplateActionParamAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotemplateActionParamAllOfWithDefaults

`func NewMotemplateActionParamAllOfWithDefaults() *MotemplateActionParamAllOf`

NewMotemplateActionParamAllOfWithDefaults instantiates a new MotemplateActionParamAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MotemplateActionParamAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MotemplateActionParamAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MotemplateActionParamAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MotemplateActionParamAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MotemplateActionParamAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MotemplateActionParamAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *MotemplateActionParamAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MotemplateActionParamAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MotemplateActionParamAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MotemplateActionParamAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *MotemplateActionParamAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MotemplateActionParamAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MotemplateActionParamAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MotemplateActionParamAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


