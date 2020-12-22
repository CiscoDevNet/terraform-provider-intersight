# KubernetesActionInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ActionInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ActionInfo"]
**FailureReason** | Pointer to **string** | Description of failure i.e. derived from the workflow failure message. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Action performed on a resource like VM, Disk etc. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the Action like InProgress, Success, Failure etc. * &#x60;None&#x60; - A place holder for the default value. * &#x60;InProgress&#x60; - Action triggered on the resource is still running. * &#x60;Success&#x60; - Action triggered on the resource is completed successfully. * &#x60;Failure&#x60; - Action triggered on the resource is failed. | [optional] [readonly] [default to "None"]

## Methods

### NewKubernetesActionInfoAllOf

`func NewKubernetesActionInfoAllOf(classId string, objectType string, ) *KubernetesActionInfoAllOf`

NewKubernetesActionInfoAllOf instantiates a new KubernetesActionInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesActionInfoAllOfWithDefaults

`func NewKubernetesActionInfoAllOfWithDefaults() *KubernetesActionInfoAllOf`

NewKubernetesActionInfoAllOfWithDefaults instantiates a new KubernetesActionInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesActionInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesActionInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesActionInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesActionInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesActionInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesActionInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFailureReason

`func (o *KubernetesActionInfoAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *KubernetesActionInfoAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *KubernetesActionInfoAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *KubernetesActionInfoAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetName

`func (o *KubernetesActionInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesActionInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesActionInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesActionInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *KubernetesActionInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesActionInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesActionInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesActionInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


