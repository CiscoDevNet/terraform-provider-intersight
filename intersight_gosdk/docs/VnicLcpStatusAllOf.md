# VnicLcpStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.LcpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.LcpStatus"]
**Reason** | Pointer to **string** | The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. | [optional] 
**Status** | Pointer to **string** | Indicates if the LCP is ready for Deploy or not. * &#x60;ok&#x60; - No issues with the LCP/SCP/VIF. * &#x60;error&#x60; - The LCP/SCP/VIF cannot be deployed due to error. * &#x60;validating&#x60; - Validation in progress for the LCP. | [optional] [default to "ok"]
**VnicInfo** | Pointer to [**[]VnicVifStatus**](VnicVifStatus.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 

## Methods

### NewVnicLcpStatusAllOf

`func NewVnicLcpStatusAllOf(classId string, objectType string, ) *VnicLcpStatusAllOf`

NewVnicLcpStatusAllOf instantiates a new VnicLcpStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicLcpStatusAllOfWithDefaults

`func NewVnicLcpStatusAllOfWithDefaults() *VnicLcpStatusAllOf`

NewVnicLcpStatusAllOfWithDefaults instantiates a new VnicLcpStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicLcpStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicLcpStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicLcpStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicLcpStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicLcpStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicLcpStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReason

`func (o *VnicLcpStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VnicLcpStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VnicLcpStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VnicLcpStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *VnicLcpStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VnicLcpStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VnicLcpStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VnicLcpStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVnicInfo

`func (o *VnicLcpStatusAllOf) GetVnicInfo() []VnicVifStatus`

GetVnicInfo returns the VnicInfo field if non-nil, zero value otherwise.

### GetVnicInfoOk

`func (o *VnicLcpStatusAllOf) GetVnicInfoOk() (*[]VnicVifStatus, bool)`

GetVnicInfoOk returns a tuple with the VnicInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicInfo

`func (o *VnicLcpStatusAllOf) SetVnicInfo(v []VnicVifStatus)`

SetVnicInfo sets VnicInfo field to given value.

### HasVnicInfo

`func (o *VnicLcpStatusAllOf) HasVnicInfo() bool`

HasVnicInfo returns a boolean if a field has been set.

### SetVnicInfoNil

`func (o *VnicLcpStatusAllOf) SetVnicInfoNil(b bool)`

 SetVnicInfoNil sets the value for VnicInfo to be an explicit nil

### UnsetVnicInfo
`func (o *VnicLcpStatusAllOf) UnsetVnicInfo()`

UnsetVnicInfo ensures that no value is present for VnicInfo, not even an explicit nil
### GetProfile

`func (o *VnicLcpStatusAllOf) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicLcpStatusAllOf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicLcpStatusAllOf) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicLcpStatusAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


