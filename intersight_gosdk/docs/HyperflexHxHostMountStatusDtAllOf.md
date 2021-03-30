# HyperflexHxHostMountStatusDtAllOf

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

### NewHyperflexHxHostMountStatusDtAllOf

`func NewHyperflexHxHostMountStatusDtAllOf(classId string, objectType string, ) *HyperflexHxHostMountStatusDtAllOf`

NewHyperflexHxHostMountStatusDtAllOf instantiates a new HyperflexHxHostMountStatusDtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxHostMountStatusDtAllOfWithDefaults

`func NewHyperflexHxHostMountStatusDtAllOfWithDefaults() *HyperflexHxHostMountStatusDtAllOf`

NewHyperflexHxHostMountStatusDtAllOfWithDefaults instantiates a new HyperflexHxHostMountStatusDtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxHostMountStatusDtAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxHostMountStatusDtAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxHostMountStatusDtAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxHostMountStatusDtAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxHostMountStatusDtAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxHostMountStatusDtAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessibility

`func (o *HyperflexHxHostMountStatusDtAllOf) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *HyperflexHxHostMountStatusDtAllOf) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *HyperflexHxHostMountStatusDtAllOf) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *HyperflexHxHostMountStatusDtAllOf) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetHostName

`func (o *HyperflexHxHostMountStatusDtAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexHxHostMountStatusDtAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexHxHostMountStatusDtAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexHxHostMountStatusDtAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetMounted

`func (o *HyperflexHxHostMountStatusDtAllOf) GetMounted() bool`

GetMounted returns the Mounted field if non-nil, zero value otherwise.

### GetMountedOk

`func (o *HyperflexHxHostMountStatusDtAllOf) GetMountedOk() (*bool, bool)`

GetMountedOk returns a tuple with the Mounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounted

`func (o *HyperflexHxHostMountStatusDtAllOf) SetMounted(v bool)`

SetMounted sets Mounted field to given value.

### HasMounted

`func (o *HyperflexHxHostMountStatusDtAllOf) HasMounted() bool`

HasMounted returns a boolean if a field has been set.

### GetReason

`func (o *HyperflexHxHostMountStatusDtAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HyperflexHxHostMountStatusDtAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HyperflexHxHostMountStatusDtAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HyperflexHxHostMountStatusDtAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


