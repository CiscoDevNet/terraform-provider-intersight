# HyperflexHxHostMountStatusDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxHostMountStatusDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxHostMountStatusDt"]
**Accessibility** | Pointer to **string** | Accessibility of HyperFlex datastore. * &#x60;ACCESSIBLE&#x60; - The HyperFlex datastore accessibility state is accessible. * &#x60;NOT_ACCESSIBLE&#x60; - The HyperFlex datastore accessibility state is not accessible. * &#x60;PARTIALLY_ACCESSIBLE&#x60; - The HyperFlex datastore accessibility state is partially accessible. * &#x60;READONLY&#x60; - The HyperFlex datastore accessibility state is read-only. * &#x60;HOST_NOT_REACHABLE&#x60; - The HyperFlex datastore accessibility state is host not reachable. * &#x60;UNKNOWN&#x60; - The HyperFlex datastore accessibility state is unknown. | [optional] [readonly] [default to "ACCESSIBLE"]
**HostName** | Pointer to **string** | HyperFlex name of host for this datastore. | [optional] [readonly] 
**Mounted** | Pointer to **bool** | Is the HyperFlex datastore mounted or not. | [optional] [readonly] 
**Reason** | Pointer to **string** | Reason for inaccessibility for this datastore. | [optional] [readonly] 

## Methods

### NewHyperflexHxHostMountStatusDt

`func NewHyperflexHxHostMountStatusDt(classId string, objectType string, ) *HyperflexHxHostMountStatusDt`

NewHyperflexHxHostMountStatusDt instantiates a new HyperflexHxHostMountStatusDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxHostMountStatusDtWithDefaults

`func NewHyperflexHxHostMountStatusDtWithDefaults() *HyperflexHxHostMountStatusDt`

NewHyperflexHxHostMountStatusDtWithDefaults instantiates a new HyperflexHxHostMountStatusDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxHostMountStatusDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxHostMountStatusDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxHostMountStatusDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxHostMountStatusDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxHostMountStatusDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxHostMountStatusDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessibility

`func (o *HyperflexHxHostMountStatusDt) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *HyperflexHxHostMountStatusDt) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *HyperflexHxHostMountStatusDt) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *HyperflexHxHostMountStatusDt) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetHostName

`func (o *HyperflexHxHostMountStatusDt) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexHxHostMountStatusDt) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexHxHostMountStatusDt) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexHxHostMountStatusDt) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetMounted

`func (o *HyperflexHxHostMountStatusDt) GetMounted() bool`

GetMounted returns the Mounted field if non-nil, zero value otherwise.

### GetMountedOk

`func (o *HyperflexHxHostMountStatusDt) GetMountedOk() (*bool, bool)`

GetMountedOk returns a tuple with the Mounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounted

`func (o *HyperflexHxHostMountStatusDt) SetMounted(v bool)`

SetMounted sets Mounted field to given value.

### HasMounted

`func (o *HyperflexHxHostMountStatusDt) HasMounted() bool`

HasMounted returns a boolean if a field has been set.

### GetReason

`func (o *HyperflexHxHostMountStatusDt) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HyperflexHxHostMountStatusDt) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HyperflexHxHostMountStatusDt) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HyperflexHxHostMountStatusDt) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


