# IssueOdataConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "issue.OdataCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "issue.OdataCondition"]
**DeviceTags** | Pointer to [**[]IssueDeviceTag**](IssueDeviceTag.md) |  | [optional] 
**Motype** | Pointer to **string** | The Intersight object type on which the condition is to be applied. The object type may be either a concrete type such as compute.RackUnit or a shared parent type such as compute.Physical. | [optional] 
**Query** | Pointer to **string** | An Odata query string interpreted to be the value portion of a $filter expression. For example, for the function query $filter&#x3D;Serial eq &#39;ABC&#39;, the &#39;query&#39; field should be assigned the string Serial eq &#39;ABC. | [optional] 

## Methods

### NewIssueOdataConditionAllOf

`func NewIssueOdataConditionAllOf(classId string, objectType string, ) *IssueOdataConditionAllOf`

NewIssueOdataConditionAllOf instantiates a new IssueOdataConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueOdataConditionAllOfWithDefaults

`func NewIssueOdataConditionAllOfWithDefaults() *IssueOdataConditionAllOf`

NewIssueOdataConditionAllOfWithDefaults instantiates a new IssueOdataConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IssueOdataConditionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IssueOdataConditionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IssueOdataConditionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IssueOdataConditionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IssueOdataConditionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IssueOdataConditionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceTags

`func (o *IssueOdataConditionAllOf) GetDeviceTags() []IssueDeviceTag`

GetDeviceTags returns the DeviceTags field if non-nil, zero value otherwise.

### GetDeviceTagsOk

`func (o *IssueOdataConditionAllOf) GetDeviceTagsOk() (*[]IssueDeviceTag, bool)`

GetDeviceTagsOk returns a tuple with the DeviceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTags

`func (o *IssueOdataConditionAllOf) SetDeviceTags(v []IssueDeviceTag)`

SetDeviceTags sets DeviceTags field to given value.

### HasDeviceTags

`func (o *IssueOdataConditionAllOf) HasDeviceTags() bool`

HasDeviceTags returns a boolean if a field has been set.

### SetDeviceTagsNil

`func (o *IssueOdataConditionAllOf) SetDeviceTagsNil(b bool)`

 SetDeviceTagsNil sets the value for DeviceTags to be an explicit nil

### UnsetDeviceTags
`func (o *IssueOdataConditionAllOf) UnsetDeviceTags()`

UnsetDeviceTags ensures that no value is present for DeviceTags, not even an explicit nil
### GetMotype

`func (o *IssueOdataConditionAllOf) GetMotype() string`

GetMotype returns the Motype field if non-nil, zero value otherwise.

### GetMotypeOk

`func (o *IssueOdataConditionAllOf) GetMotypeOk() (*string, bool)`

GetMotypeOk returns a tuple with the Motype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotype

`func (o *IssueOdataConditionAllOf) SetMotype(v string)`

SetMotype sets Motype field to given value.

### HasMotype

`func (o *IssueOdataConditionAllOf) HasMotype() bool`

HasMotype returns a boolean if a field has been set.

### GetQuery

`func (o *IssueOdataConditionAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IssueOdataConditionAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IssueOdataConditionAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IssueOdataConditionAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


