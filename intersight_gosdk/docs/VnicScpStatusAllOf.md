# VnicScpStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.ScpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.ScpStatus"]
**Reason** | Pointer to **string** | The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. | [optional] 
**Status** | Pointer to **string** | Indicates if the LCP is ready for Deploy or not. * &#x60;ok&#x60; - No issues with the LCP/SCP/VIF. * &#x60;error&#x60; - The LCP/SCP/VIF cannot be deployed due to error. * &#x60;validating&#x60; - Validation in progress for the LCP. | [optional] [default to "ok"]
**VhbaInfo** | Pointer to [**[]VnicVifStatus**](VnicVifStatus.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 

## Methods

### NewVnicScpStatusAllOf

`func NewVnicScpStatusAllOf(classId string, objectType string, ) *VnicScpStatusAllOf`

NewVnicScpStatusAllOf instantiates a new VnicScpStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicScpStatusAllOfWithDefaults

`func NewVnicScpStatusAllOfWithDefaults() *VnicScpStatusAllOf`

NewVnicScpStatusAllOfWithDefaults instantiates a new VnicScpStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicScpStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicScpStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicScpStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicScpStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicScpStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicScpStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReason

`func (o *VnicScpStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VnicScpStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VnicScpStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VnicScpStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *VnicScpStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VnicScpStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VnicScpStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VnicScpStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVhbaInfo

`func (o *VnicScpStatusAllOf) GetVhbaInfo() []VnicVifStatus`

GetVhbaInfo returns the VhbaInfo field if non-nil, zero value otherwise.

### GetVhbaInfoOk

`func (o *VnicScpStatusAllOf) GetVhbaInfoOk() (*[]VnicVifStatus, bool)`

GetVhbaInfoOk returns a tuple with the VhbaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhbaInfo

`func (o *VnicScpStatusAllOf) SetVhbaInfo(v []VnicVifStatus)`

SetVhbaInfo sets VhbaInfo field to given value.

### HasVhbaInfo

`func (o *VnicScpStatusAllOf) HasVhbaInfo() bool`

HasVhbaInfo returns a boolean if a field has been set.

### SetVhbaInfoNil

`func (o *VnicScpStatusAllOf) SetVhbaInfoNil(b bool)`

 SetVhbaInfoNil sets the value for VhbaInfo to be an explicit nil

### UnsetVhbaInfo
`func (o *VnicScpStatusAllOf) UnsetVhbaInfo()`

UnsetVhbaInfo ensures that no value is present for VhbaInfo, not even an explicit nil
### GetProfile

`func (o *VnicScpStatusAllOf) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicScpStatusAllOf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicScpStatusAllOf) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicScpStatusAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


