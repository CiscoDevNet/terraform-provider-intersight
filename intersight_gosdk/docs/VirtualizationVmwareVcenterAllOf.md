# VirtualizationVmwareVcenterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVcenter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVcenter"]
**ClusterCount** | Pointer to **int64** | Count of all Clusters associated with the vcenter. | [optional] [readonly] 
**DatacenterCount** | Pointer to **int64** | Count of all Datacenters in the vcenter. | [optional] [readonly] 
**DatastoreCount** | Pointer to **int64** | Count of all Datastores Templates associated with the vcenter. | [optional] [readonly] 
**DistributedVirtualSwitchCount** | Pointer to **int64** | Count of all Distributed Virtual Switches associated with vcenter. | [optional] [readonly] 
**DsClusterCount** | Pointer to **int64** | Count of all Datastore cluster associated with the vcenter. | [optional] [readonly] 
**ExternalIp** | Pointer to **string** | External Ip address fot the vcenter. | [optional] [readonly] 
**HostCount** | Pointer to **int64** | Count of all Hosts associated with the vcenter. | [optional] [readonly] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**TargetName** | Pointer to **string** | Name of th Target with which the vcenter was claimed. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | Count of all Virtual Machines associated with the vcenter. | [optional] [readonly] 
**VmTemplatesCount** | Pointer to **int64** | Count of all VM Templates associated with the vcenter. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareVcenterAllOf

`func NewVirtualizationVmwareVcenterAllOf(classId string, objectType string, ) *VirtualizationVmwareVcenterAllOf`

NewVirtualizationVmwareVcenterAllOf instantiates a new VirtualizationVmwareVcenterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVcenterAllOfWithDefaults

`func NewVirtualizationVmwareVcenterAllOfWithDefaults() *VirtualizationVmwareVcenterAllOf`

NewVirtualizationVmwareVcenterAllOfWithDefaults instantiates a new VirtualizationVmwareVcenterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVcenterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVcenterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVcenterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVcenterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVcenterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVcenterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterCount

`func (o *VirtualizationVmwareVcenterAllOf) GetClusterCount() int64`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetClusterCountOk() (*int64, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *VirtualizationVmwareVcenterAllOf) SetClusterCount(v int64)`

SetClusterCount sets ClusterCount field to given value.

### HasClusterCount

`func (o *VirtualizationVmwareVcenterAllOf) HasClusterCount() bool`

HasClusterCount returns a boolean if a field has been set.

### GetDatacenterCount

`func (o *VirtualizationVmwareVcenterAllOf) GetDatacenterCount() int64`

GetDatacenterCount returns the DatacenterCount field if non-nil, zero value otherwise.

### GetDatacenterCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetDatacenterCountOk() (*int64, bool)`

GetDatacenterCountOk returns a tuple with the DatacenterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterCount

`func (o *VirtualizationVmwareVcenterAllOf) SetDatacenterCount(v int64)`

SetDatacenterCount sets DatacenterCount field to given value.

### HasDatacenterCount

`func (o *VirtualizationVmwareVcenterAllOf) HasDatacenterCount() bool`

HasDatacenterCount returns a boolean if a field has been set.

### GetDatastoreCount

`func (o *VirtualizationVmwareVcenterAllOf) GetDatastoreCount() int64`

GetDatastoreCount returns the DatastoreCount field if non-nil, zero value otherwise.

### GetDatastoreCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetDatastoreCountOk() (*int64, bool)`

GetDatastoreCountOk returns a tuple with the DatastoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreCount

`func (o *VirtualizationVmwareVcenterAllOf) SetDatastoreCount(v int64)`

SetDatastoreCount sets DatastoreCount field to given value.

### HasDatastoreCount

`func (o *VirtualizationVmwareVcenterAllOf) HasDatastoreCount() bool`

HasDatastoreCount returns a boolean if a field has been set.

### GetDistributedVirtualSwitchCount

`func (o *VirtualizationVmwareVcenterAllOf) GetDistributedVirtualSwitchCount() int64`

GetDistributedVirtualSwitchCount returns the DistributedVirtualSwitchCount field if non-nil, zero value otherwise.

### GetDistributedVirtualSwitchCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetDistributedVirtualSwitchCountOk() (*int64, bool)`

GetDistributedVirtualSwitchCountOk returns a tuple with the DistributedVirtualSwitchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedVirtualSwitchCount

`func (o *VirtualizationVmwareVcenterAllOf) SetDistributedVirtualSwitchCount(v int64)`

SetDistributedVirtualSwitchCount sets DistributedVirtualSwitchCount field to given value.

### HasDistributedVirtualSwitchCount

`func (o *VirtualizationVmwareVcenterAllOf) HasDistributedVirtualSwitchCount() bool`

HasDistributedVirtualSwitchCount returns a boolean if a field has been set.

### GetDsClusterCount

`func (o *VirtualizationVmwareVcenterAllOf) GetDsClusterCount() int64`

GetDsClusterCount returns the DsClusterCount field if non-nil, zero value otherwise.

### GetDsClusterCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetDsClusterCountOk() (*int64, bool)`

GetDsClusterCountOk returns a tuple with the DsClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsClusterCount

`func (o *VirtualizationVmwareVcenterAllOf) SetDsClusterCount(v int64)`

SetDsClusterCount sets DsClusterCount field to given value.

### HasDsClusterCount

`func (o *VirtualizationVmwareVcenterAllOf) HasDsClusterCount() bool`

HasDsClusterCount returns a boolean if a field has been set.

### GetExternalIp

`func (o *VirtualizationVmwareVcenterAllOf) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *VirtualizationVmwareVcenterAllOf) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *VirtualizationVmwareVcenterAllOf) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *VirtualizationVmwareVcenterAllOf) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetHostCount

`func (o *VirtualizationVmwareVcenterAllOf) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *VirtualizationVmwareVcenterAllOf) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *VirtualizationVmwareVcenterAllOf) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationVmwareVcenterAllOf) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationVmwareVcenterAllOf) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationVmwareVcenterAllOf) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationVmwareVcenterAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VirtualizationVmwareVcenterAllOf) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VirtualizationVmwareVcenterAllOf) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetTargetName

`func (o *VirtualizationVmwareVcenterAllOf) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *VirtualizationVmwareVcenterAllOf) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *VirtualizationVmwareVcenterAllOf) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *VirtualizationVmwareVcenterAllOf) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetVmCount

`func (o *VirtualizationVmwareVcenterAllOf) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VirtualizationVmwareVcenterAllOf) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VirtualizationVmwareVcenterAllOf) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetVmTemplatesCount

`func (o *VirtualizationVmwareVcenterAllOf) GetVmTemplatesCount() int64`

GetVmTemplatesCount returns the VmTemplatesCount field if non-nil, zero value otherwise.

### GetVmTemplatesCountOk

`func (o *VirtualizationVmwareVcenterAllOf) GetVmTemplatesCountOk() (*int64, bool)`

GetVmTemplatesCountOk returns a tuple with the VmTemplatesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplatesCount

`func (o *VirtualizationVmwareVcenterAllOf) SetVmTemplatesCount(v int64)`

SetVmTemplatesCount sets VmTemplatesCount field to given value.

### HasVmTemplatesCount

`func (o *VirtualizationVmwareVcenterAllOf) HasVmTemplatesCount() bool`

HasVmTemplatesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


