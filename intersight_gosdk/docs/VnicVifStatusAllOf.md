# VnicVifStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.VifStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.VifStatus"]
**Name** | Pointer to **string** | Name of the vNIC for which the status is reported. | [optional] 
**Reason** | Pointer to **string** | The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. | [optional] 
**Status** | Pointer to **string** | Indicates if the vNIC / vHBA is ready for deploy or not. * &#x60;ok&#x60; - No issues with the LCP/SCP/VIF. * &#x60;error&#x60; - The LCP/SCP/VIF cannot be deployed due to error. * &#x60;validating&#x60; - Validation in progress for the LCP. | [optional] [default to "ok"]

## Methods

### NewVnicVifStatusAllOf

`func NewVnicVifStatusAllOf(classId string, objectType string, ) *VnicVifStatusAllOf`

NewVnicVifStatusAllOf instantiates a new VnicVifStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVifStatusAllOfWithDefaults

`func NewVnicVifStatusAllOfWithDefaults() *VnicVifStatusAllOf`

NewVnicVifStatusAllOfWithDefaults instantiates a new VnicVifStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicVifStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicVifStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicVifStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicVifStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicVifStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicVifStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *VnicVifStatusAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicVifStatusAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicVifStatusAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicVifStatusAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *VnicVifStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VnicVifStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VnicVifStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VnicVifStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *VnicVifStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VnicVifStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VnicVifStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VnicVifStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


