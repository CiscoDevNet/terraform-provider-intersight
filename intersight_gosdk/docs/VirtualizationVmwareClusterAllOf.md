# VirtualizationVmwareClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreCount** | Pointer to **int64** | Count of all datastores associated with this cluster. | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareClusterAllOf

`func NewVirtualizationVmwareClusterAllOf() *VirtualizationVmwareClusterAllOf`

NewVirtualizationVmwareClusterAllOf instantiates a new VirtualizationVmwareClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareClusterAllOfWithDefaults

`func NewVirtualizationVmwareClusterAllOfWithDefaults() *VirtualizationVmwareClusterAllOf`

NewVirtualizationVmwareClusterAllOfWithDefaults instantiates a new VirtualizationVmwareClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreCount

`func (o *VirtualizationVmwareClusterAllOf) GetDatastoreCount() int64`

GetDatastoreCount returns the DatastoreCount field if non-nil, zero value otherwise.

### GetDatastoreCountOk

`func (o *VirtualizationVmwareClusterAllOf) GetDatastoreCountOk() (*int64, bool)`

GetDatastoreCountOk returns a tuple with the DatastoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreCount

`func (o *VirtualizationVmwareClusterAllOf) SetDatastoreCount(v int64)`

SetDatastoreCount sets DatastoreCount field to given value.

### HasDatastoreCount

`func (o *VirtualizationVmwareClusterAllOf) HasDatastoreCount() bool`

HasDatastoreCount returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareClusterAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareClusterAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareClusterAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareClusterAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


