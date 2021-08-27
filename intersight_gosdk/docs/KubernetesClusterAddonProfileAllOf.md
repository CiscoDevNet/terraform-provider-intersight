# KubernetesClusterAddonProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.ClusterAddonProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.ClusterAddonProfile"]
**Addons** | Pointer to [**[]KubernetesAddon**](KubernetesAddon.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the cluster addon profile. | [optional] 
**AssociatedCluster** | Pointer to [**KubernetesClusterRelationship**](KubernetesClusterRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesClusterAddonProfileAllOf

`func NewKubernetesClusterAddonProfileAllOf(classId string, objectType string, ) *KubernetesClusterAddonProfileAllOf`

NewKubernetesClusterAddonProfileAllOf instantiates a new KubernetesClusterAddonProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterAddonProfileAllOfWithDefaults

`func NewKubernetesClusterAddonProfileAllOfWithDefaults() *KubernetesClusterAddonProfileAllOf`

NewKubernetesClusterAddonProfileAllOfWithDefaults instantiates a new KubernetesClusterAddonProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesClusterAddonProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesClusterAddonProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesClusterAddonProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesClusterAddonProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesClusterAddonProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesClusterAddonProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddons

`func (o *KubernetesClusterAddonProfileAllOf) GetAddons() []KubernetesAddon`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *KubernetesClusterAddonProfileAllOf) GetAddonsOk() (*[]KubernetesAddon, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *KubernetesClusterAddonProfileAllOf) SetAddons(v []KubernetesAddon)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *KubernetesClusterAddonProfileAllOf) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### SetAddonsNil

`func (o *KubernetesClusterAddonProfileAllOf) SetAddonsNil(b bool)`

 SetAddonsNil sets the value for Addons to be an explicit nil

### UnsetAddons
`func (o *KubernetesClusterAddonProfileAllOf) UnsetAddons()`

UnsetAddons ensures that no value is present for Addons, not even an explicit nil
### GetName

`func (o *KubernetesClusterAddonProfileAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterAddonProfileAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterAddonProfileAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesClusterAddonProfileAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssociatedCluster

`func (o *KubernetesClusterAddonProfileAllOf) GetAssociatedCluster() KubernetesClusterRelationship`

GetAssociatedCluster returns the AssociatedCluster field if non-nil, zero value otherwise.

### GetAssociatedClusterOk

`func (o *KubernetesClusterAddonProfileAllOf) GetAssociatedClusterOk() (*KubernetesClusterRelationship, bool)`

GetAssociatedClusterOk returns a tuple with the AssociatedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCluster

`func (o *KubernetesClusterAddonProfileAllOf) SetAssociatedCluster(v KubernetesClusterRelationship)`

SetAssociatedCluster sets AssociatedCluster field to given value.

### HasAssociatedCluster

`func (o *KubernetesClusterAddonProfileAllOf) HasAssociatedCluster() bool`

HasAssociatedCluster returns a boolean if a field has been set.

### GetOrganization

`func (o *KubernetesClusterAddonProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KubernetesClusterAddonProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KubernetesClusterAddonProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KubernetesClusterAddonProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


