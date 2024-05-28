# ApicAciPod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.AciPod"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.AciPod"]
**Name** | Pointer to **string** | Name of a pod extracted from Dn within the Cisco Application Policy Infrastructure Controller (APIC). | [optional] 
**PodType** | Pointer to **string** | Pod type of an object within the Cisco Application Policy Infrastructure Controller (APIC). | [optional] 

## Methods

### NewApicAciPod

`func NewApicAciPod(classId string, objectType string, ) *ApicAciPod`

NewApicAciPod instantiates a new ApicAciPod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicAciPodWithDefaults

`func NewApicAciPodWithDefaults() *ApicAciPod`

NewApicAciPodWithDefaults instantiates a new ApicAciPod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicAciPod) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicAciPod) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicAciPod) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicAciPod) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicAciPod) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicAciPod) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *ApicAciPod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicAciPod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicAciPod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicAciPod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPodType

`func (o *ApicAciPod) GetPodType() string`

GetPodType returns the PodType field if non-nil, zero value otherwise.

### GetPodTypeOk

`func (o *ApicAciPod) GetPodTypeOk() (*string, bool)`

GetPodTypeOk returns a tuple with the PodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodType

`func (o *ApicAciPod) SetPodType(v string)`

SetPodType sets PodType field to given value.

### HasPodType

`func (o *ApicAciPod) HasPodType() bool`

HasPodType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


