# KubernetesAddonVersionReferenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.AddonVersionReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.AddonVersionReference"]
**Name** | Pointer to **string** | Name of the addon to lookup. | [optional] 
**Version** | Pointer to **string** | Version number to filter the addon with. | [optional] 

## Methods

### NewKubernetesAddonVersionReferenceAllOf

`func NewKubernetesAddonVersionReferenceAllOf(classId string, objectType string, ) *KubernetesAddonVersionReferenceAllOf`

NewKubernetesAddonVersionReferenceAllOf instantiates a new KubernetesAddonVersionReferenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAddonVersionReferenceAllOfWithDefaults

`func NewKubernetesAddonVersionReferenceAllOfWithDefaults() *KubernetesAddonVersionReferenceAllOf`

NewKubernetesAddonVersionReferenceAllOfWithDefaults instantiates a new KubernetesAddonVersionReferenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesAddonVersionReferenceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesAddonVersionReferenceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesAddonVersionReferenceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesAddonVersionReferenceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesAddonVersionReferenceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesAddonVersionReferenceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *KubernetesAddonVersionReferenceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesAddonVersionReferenceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesAddonVersionReferenceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesAddonVersionReferenceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *KubernetesAddonVersionReferenceAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesAddonVersionReferenceAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesAddonVersionReferenceAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KubernetesAddonVersionReferenceAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


