# VirtualizationVmwareVcenter

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

### NewVirtualizationVmwareVcenter

`func NewVirtualizationVmwareVcenter(classId string, objectType string, ) *VirtualizationVmwareVcenter`

NewVirtualizationVmwareVcenter instantiates a new VirtualizationVmwareVcenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVcenterWithDefaults

`func NewVirtualizationVmwareVcenterWithDefaults() *VirtualizationVmwareVcenter`

NewVirtualizationVmwareVcenterWithDefaults instantiates a new VirtualizationVmwareVcenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVcenter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVcenter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVcenter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVcenter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVcenter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVcenter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterCount

`func (o *VirtualizationVmwareVcenter) GetClusterCount() int64`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *VirtualizationVmwareVcenter) GetClusterCountOk() (*int64, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *VirtualizationVmwareVcenter) SetClusterCount(v int64)`

SetClusterCount sets ClusterCount field to given value.

### HasClusterCount

`func (o *VirtualizationVmwareVcenter) HasClusterCount() bool`

HasClusterCount returns a boolean if a field has been set.

### GetDatacenterCount

`func (o *VirtualizationVmwareVcenter) GetDatacenterCount() int64`

GetDatacenterCount returns the DatacenterCount field if non-nil, zero value otherwise.

### GetDatacenterCountOk

`func (o *VirtualizationVmwareVcenter) GetDatacenterCountOk() (*int64, bool)`

GetDatacenterCountOk returns a tuple with the DatacenterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterCount

`func (o *VirtualizationVmwareVcenter) SetDatacenterCount(v int64)`

SetDatacenterCount sets DatacenterCount field to given value.

### HasDatacenterCount

`func (o *VirtualizationVmwareVcenter) HasDatacenterCount() bool`

HasDatacenterCount returns a boolean if a field has been set.

### GetDatastoreCount

`func (o *VirtualizationVmwareVcenter) GetDatastoreCount() int64`

GetDatastoreCount returns the DatastoreCount field if non-nil, zero value otherwise.

### GetDatastoreCountOk

`func (o *VirtualizationVmwareVcenter) GetDatastoreCountOk() (*int64, bool)`

GetDatastoreCountOk returns a tuple with the DatastoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreCount

`func (o *VirtualizationVmwareVcenter) SetDatastoreCount(v int64)`

SetDatastoreCount sets DatastoreCount field to given value.

### HasDatastoreCount

`func (o *VirtualizationVmwareVcenter) HasDatastoreCount() bool`

HasDatastoreCount returns a boolean if a field has been set.

### GetDistributedVirtualSwitchCount

`func (o *VirtualizationVmwareVcenter) GetDistributedVirtualSwitchCount() int64`

GetDistributedVirtualSwitchCount returns the DistributedVirtualSwitchCount field if non-nil, zero value otherwise.

### GetDistributedVirtualSwitchCountOk

`func (o *VirtualizationVmwareVcenter) GetDistributedVirtualSwitchCountOk() (*int64, bool)`

GetDistributedVirtualSwitchCountOk returns a tuple with the DistributedVirtualSwitchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedVirtualSwitchCount

`func (o *VirtualizationVmwareVcenter) SetDistributedVirtualSwitchCount(v int64)`

SetDistributedVirtualSwitchCount sets DistributedVirtualSwitchCount field to given value.

### HasDistributedVirtualSwitchCount

`func (o *VirtualizationVmwareVcenter) HasDistributedVirtualSwitchCount() bool`

HasDistributedVirtualSwitchCount returns a boolean if a field has been set.

### GetDsClusterCount

`func (o *VirtualizationVmwareVcenter) GetDsClusterCount() int64`

GetDsClusterCount returns the DsClusterCount field if non-nil, zero value otherwise.

### GetDsClusterCountOk

`func (o *VirtualizationVmwareVcenter) GetDsClusterCountOk() (*int64, bool)`

GetDsClusterCountOk returns a tuple with the DsClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsClusterCount

`func (o *VirtualizationVmwareVcenter) SetDsClusterCount(v int64)`

SetDsClusterCount sets DsClusterCount field to given value.

### HasDsClusterCount

`func (o *VirtualizationVmwareVcenter) HasDsClusterCount() bool`

HasDsClusterCount returns a boolean if a field has been set.

### GetExternalIp

`func (o *VirtualizationVmwareVcenter) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *VirtualizationVmwareVcenter) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *VirtualizationVmwareVcenter) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *VirtualizationVmwareVcenter) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetHostCount

`func (o *VirtualizationVmwareVcenter) GetHostCount() int64`

GetHostCount returns the HostCount field if non-nil, zero value otherwise.

### GetHostCountOk

`func (o *VirtualizationVmwareVcenter) GetHostCountOk() (*int64, bool)`

GetHostCountOk returns a tuple with the HostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCount

`func (o *VirtualizationVmwareVcenter) SetHostCount(v int64)`

SetHostCount sets HostCount field to given value.

### HasHostCount

`func (o *VirtualizationVmwareVcenter) HasHostCount() bool`

HasHostCount returns a boolean if a field has been set.

### GetIpAddress

`func (o *VirtualizationVmwareVcenter) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VirtualizationVmwareVcenter) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VirtualizationVmwareVcenter) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VirtualizationVmwareVcenter) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VirtualizationVmwareVcenter) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VirtualizationVmwareVcenter) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetTargetName

`func (o *VirtualizationVmwareVcenter) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *VirtualizationVmwareVcenter) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *VirtualizationVmwareVcenter) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *VirtualizationVmwareVcenter) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetVmCount

`func (o *VirtualizationVmwareVcenter) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *VirtualizationVmwareVcenter) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *VirtualizationVmwareVcenter) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *VirtualizationVmwareVcenter) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetVmTemplatesCount

`func (o *VirtualizationVmwareVcenter) GetVmTemplatesCount() int64`

GetVmTemplatesCount returns the VmTemplatesCount field if non-nil, zero value otherwise.

### GetVmTemplatesCountOk

`func (o *VirtualizationVmwareVcenter) GetVmTemplatesCountOk() (*int64, bool)`

GetVmTemplatesCountOk returns a tuple with the VmTemplatesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplatesCount

`func (o *VirtualizationVmwareVcenter) SetVmTemplatesCount(v int64)`

SetVmTemplatesCount sets VmTemplatesCount field to given value.

### HasVmTemplatesCount

`func (o *VirtualizationVmwareVcenter) HasVmTemplatesCount() bool`

HasVmTemplatesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


