# IaasUcsdManagedInfraAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.UcsdManagedInfra"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.UcsdManagedInfra"]
**AdvancedCatalogCount** | Pointer to **int64** | Total advanced catalogs in UCSD. | [optional] [readonly] 
**BmCatalogCount** | Pointer to **int64** | Total bare metal catalogs in UCSD. | [optional] [readonly] 
**ContainerCatalogCount** | Pointer to **int64** | Total service container catalogs in UCSD. | [optional] [readonly] 
**EsxiHostCount** | Pointer to **int64** | Total ESXi hosts in UCSD. | [optional] [readonly] 
**ExternalGroupCount** | Pointer to **int64** | Total external (Ldap) groups in UCSD. | [optional] [readonly] 
**HypervHostCount** | Pointer to **int64** | Total HyperV hosts in UCSD. | [optional] [readonly] 
**LocalGroupCount** | Pointer to **int64** | Total local groups in UCSD. | [optional] [readonly] 
**StandardCatalogCount** | Pointer to **int64** | Total standard catalogs in UCSD. | [optional] [readonly] 
**UserCount** | Pointer to **int64** | Total user accounts in UCSD. | [optional] [readonly] 
**VdcCount** | Pointer to **int64** | Total virtual datacenters in UCSD. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | Total Virtual machines in UCSD. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewIaasUcsdManagedInfraAllOf

`func NewIaasUcsdManagedInfraAllOf(classId string, objectType string, ) *IaasUcsdManagedInfraAllOf`

NewIaasUcsdManagedInfraAllOf instantiates a new IaasUcsdManagedInfraAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasUcsdManagedInfraAllOfWithDefaults

`func NewIaasUcsdManagedInfraAllOfWithDefaults() *IaasUcsdManagedInfraAllOf`

NewIaasUcsdManagedInfraAllOfWithDefaults instantiates a new IaasUcsdManagedInfraAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasUcsdManagedInfraAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasUcsdManagedInfraAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasUcsdManagedInfraAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasUcsdManagedInfraAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasUcsdManagedInfraAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasUcsdManagedInfraAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdvancedCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) GetAdvancedCatalogCount() int64`

GetAdvancedCatalogCount returns the AdvancedCatalogCount field if non-nil, zero value otherwise.

### GetAdvancedCatalogCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetAdvancedCatalogCountOk() (*int64, bool)`

GetAdvancedCatalogCountOk returns a tuple with the AdvancedCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) SetAdvancedCatalogCount(v int64)`

SetAdvancedCatalogCount sets AdvancedCatalogCount field to given value.

### HasAdvancedCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) HasAdvancedCatalogCount() bool`

HasAdvancedCatalogCount returns a boolean if a field has been set.

### GetBmCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) GetBmCatalogCount() int64`

GetBmCatalogCount returns the BmCatalogCount field if non-nil, zero value otherwise.

### GetBmCatalogCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetBmCatalogCountOk() (*int64, bool)`

GetBmCatalogCountOk returns a tuple with the BmCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) SetBmCatalogCount(v int64)`

SetBmCatalogCount sets BmCatalogCount field to given value.

### HasBmCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) HasBmCatalogCount() bool`

HasBmCatalogCount returns a boolean if a field has been set.

### GetContainerCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) GetContainerCatalogCount() int64`

GetContainerCatalogCount returns the ContainerCatalogCount field if non-nil, zero value otherwise.

### GetContainerCatalogCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetContainerCatalogCountOk() (*int64, bool)`

GetContainerCatalogCountOk returns a tuple with the ContainerCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) SetContainerCatalogCount(v int64)`

SetContainerCatalogCount sets ContainerCatalogCount field to given value.

### HasContainerCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) HasContainerCatalogCount() bool`

HasContainerCatalogCount returns a boolean if a field has been set.

### GetEsxiHostCount

`func (o *IaasUcsdManagedInfraAllOf) GetEsxiHostCount() int64`

GetEsxiHostCount returns the EsxiHostCount field if non-nil, zero value otherwise.

### GetEsxiHostCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetEsxiHostCountOk() (*int64, bool)`

GetEsxiHostCountOk returns a tuple with the EsxiHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxiHostCount

`func (o *IaasUcsdManagedInfraAllOf) SetEsxiHostCount(v int64)`

SetEsxiHostCount sets EsxiHostCount field to given value.

### HasEsxiHostCount

`func (o *IaasUcsdManagedInfraAllOf) HasEsxiHostCount() bool`

HasEsxiHostCount returns a boolean if a field has been set.

### GetExternalGroupCount

`func (o *IaasUcsdManagedInfraAllOf) GetExternalGroupCount() int64`

GetExternalGroupCount returns the ExternalGroupCount field if non-nil, zero value otherwise.

### GetExternalGroupCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetExternalGroupCountOk() (*int64, bool)`

GetExternalGroupCountOk returns a tuple with the ExternalGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupCount

`func (o *IaasUcsdManagedInfraAllOf) SetExternalGroupCount(v int64)`

SetExternalGroupCount sets ExternalGroupCount field to given value.

### HasExternalGroupCount

`func (o *IaasUcsdManagedInfraAllOf) HasExternalGroupCount() bool`

HasExternalGroupCount returns a boolean if a field has been set.

### GetHypervHostCount

`func (o *IaasUcsdManagedInfraAllOf) GetHypervHostCount() int64`

GetHypervHostCount returns the HypervHostCount field if non-nil, zero value otherwise.

### GetHypervHostCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetHypervHostCountOk() (*int64, bool)`

GetHypervHostCountOk returns a tuple with the HypervHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervHostCount

`func (o *IaasUcsdManagedInfraAllOf) SetHypervHostCount(v int64)`

SetHypervHostCount sets HypervHostCount field to given value.

### HasHypervHostCount

`func (o *IaasUcsdManagedInfraAllOf) HasHypervHostCount() bool`

HasHypervHostCount returns a boolean if a field has been set.

### GetLocalGroupCount

`func (o *IaasUcsdManagedInfraAllOf) GetLocalGroupCount() int64`

GetLocalGroupCount returns the LocalGroupCount field if non-nil, zero value otherwise.

### GetLocalGroupCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetLocalGroupCountOk() (*int64, bool)`

GetLocalGroupCountOk returns a tuple with the LocalGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroupCount

`func (o *IaasUcsdManagedInfraAllOf) SetLocalGroupCount(v int64)`

SetLocalGroupCount sets LocalGroupCount field to given value.

### HasLocalGroupCount

`func (o *IaasUcsdManagedInfraAllOf) HasLocalGroupCount() bool`

HasLocalGroupCount returns a boolean if a field has been set.

### GetStandardCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) GetStandardCatalogCount() int64`

GetStandardCatalogCount returns the StandardCatalogCount field if non-nil, zero value otherwise.

### GetStandardCatalogCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetStandardCatalogCountOk() (*int64, bool)`

GetStandardCatalogCountOk returns a tuple with the StandardCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) SetStandardCatalogCount(v int64)`

SetStandardCatalogCount sets StandardCatalogCount field to given value.

### HasStandardCatalogCount

`func (o *IaasUcsdManagedInfraAllOf) HasStandardCatalogCount() bool`

HasStandardCatalogCount returns a boolean if a field has been set.

### GetUserCount

`func (o *IaasUcsdManagedInfraAllOf) GetUserCount() int64`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetUserCountOk() (*int64, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *IaasUcsdManagedInfraAllOf) SetUserCount(v int64)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *IaasUcsdManagedInfraAllOf) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetVdcCount

`func (o *IaasUcsdManagedInfraAllOf) GetVdcCount() int64`

GetVdcCount returns the VdcCount field if non-nil, zero value otherwise.

### GetVdcCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetVdcCountOk() (*int64, bool)`

GetVdcCountOk returns a tuple with the VdcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcCount

`func (o *IaasUcsdManagedInfraAllOf) SetVdcCount(v int64)`

SetVdcCount sets VdcCount field to given value.

### HasVdcCount

`func (o *IaasUcsdManagedInfraAllOf) HasVdcCount() bool`

HasVdcCount returns a boolean if a field has been set.

### GetVmCount

`func (o *IaasUcsdManagedInfraAllOf) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *IaasUcsdManagedInfraAllOf) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *IaasUcsdManagedInfraAllOf) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *IaasUcsdManagedInfraAllOf) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetGuid

`func (o *IaasUcsdManagedInfraAllOf) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasUcsdManagedInfraAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasUcsdManagedInfraAllOf) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasUcsdManagedInfraAllOf) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


