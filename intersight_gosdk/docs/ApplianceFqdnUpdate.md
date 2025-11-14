# ApplianceFqdnUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.FqdnUpdate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.FqdnUpdate"]
**EndTime** | Pointer to **time.Time** | End date of the appliance FQDN change. | [optional] [readonly] 
**Fqdn** | Pointer to **string** | The FQDN (fully qualified domain name) of the destination appliance for target migration. | [optional] 
**StartTime** | Pointer to **time.Time** | Start date of the appliance FQDN change. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the FQDN update operation. * &#x60;NotStarted&#x60; - Appliance FQDN update operation has not started. * &#x60;Started&#x60; - Appliance FQDN update operation has started. * &#x60;Failed&#x60; - Appliance FQDN update operation has failed. * &#x60;Completed&#x60; - Appliance FQDN update operation has completed. | [optional] [readonly] [default to "NotStarted"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceFqdnUpdate

`func NewApplianceFqdnUpdate(classId string, objectType string, ) *ApplianceFqdnUpdate`

NewApplianceFqdnUpdate instantiates a new ApplianceFqdnUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceFqdnUpdateWithDefaults

`func NewApplianceFqdnUpdateWithDefaults() *ApplianceFqdnUpdate`

NewApplianceFqdnUpdateWithDefaults instantiates a new ApplianceFqdnUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceFqdnUpdate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceFqdnUpdate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceFqdnUpdate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceFqdnUpdate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceFqdnUpdate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceFqdnUpdate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndTime

`func (o *ApplianceFqdnUpdate) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceFqdnUpdate) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceFqdnUpdate) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceFqdnUpdate) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFqdn

`func (o *ApplianceFqdnUpdate) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ApplianceFqdnUpdate) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ApplianceFqdnUpdate) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ApplianceFqdnUpdate) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetStartTime

`func (o *ApplianceFqdnUpdate) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceFqdnUpdate) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceFqdnUpdate) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceFqdnUpdate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceFqdnUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceFqdnUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceFqdnUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceFqdnUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceFqdnUpdate) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceFqdnUpdate) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceFqdnUpdate) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceFqdnUpdate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ApplianceFqdnUpdate) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ApplianceFqdnUpdate) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


