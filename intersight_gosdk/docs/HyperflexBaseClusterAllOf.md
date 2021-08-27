# HyperflexBaseClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AlarmSummary** | Pointer to [**NullableHyperflexAlarmSummary**](HyperflexAlarmSummary.md) |  | [optional] 
**CapacityRunway** | Pointer to **int64** | The number of days remaining before the cluster&#39;s storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \&quot;Unknown\&quot; for a cluster that is not connected or with not sufficient data. | [optional] [readonly] [default to 2147483647]
**ClusterName** | Pointer to **string** | The name of this HyperFlex cluster. | [optional] [readonly] 
**ClusterPurpose** | Pointer to **string** | This can be a Storage or Compute cluster. A storage cluster contains storage nodes that are used to persist data. A compute cluster contains compute nodes that are used for executing business logic. * &#x60;Storage&#x60; - Cluster of storage nodes used to persist data. * &#x60;Compute&#x60; - Cluster of compute nodes used to execute business logic. * &#x60;Unknown&#x60; - This cluster type is Unknown. Expect Compute or Storage as valid values. | [optional] [default to "Storage"]
**ComputeNodeCount** | Pointer to **int64** | The number of compute nodes that belong to this cluster. | [optional] [readonly] 
**ConvergedNodeCount** | Pointer to **int64** | The number of converged nodes that belong to this cluster. | [optional] [readonly] 
**DeploymentType** | Pointer to **string** | The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as &#39;NA&#39; (not available). * &#x60;NA&#x60; - The deployment type of the HyperFlex cluster is not available. * &#x60;Datacenter&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * &#x60;Stretched Cluster&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * &#x60;Edge&#x60; - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes. | [optional] [readonly] [default to "NA"]
**DriveType** | Pointer to **string** | The type of the drives used for storage in this cluster. * &#x60;NA&#x60; - The drive type of the HyperFlex cluster is not available. * &#x60;All-Flash&#x60; - Indicates that this HyperFlex cluster contains flash drives only. * &#x60;Hybrid&#x60; - Indicates that this HyperFlex cluster contains both flash and hard disk drives. | [optional] [readonly] [default to "NA"]
**HxVersion** | Pointer to **string** | The HyperFlex Data or Application Platform version of this cluster. | [optional] [readonly] 
**HypervisorVersion** | Pointer to **string** | The version of hypervisor running on this cluster. | [optional] [readonly] 
**StorageCapacity** | Pointer to **int64** | The storage capacity in this cluster. | [optional] [readonly] 
**StorageNodeCount** | Pointer to **int64** | The number of storage nodes that belong to this cluster. | [optional] [readonly] 
**StorageUtilization** | Pointer to **float32** | The storage utilization is computed based on total capacity and current capacity utilization. | [optional] [readonly] 
**UtilizationPercentage** | Pointer to **float32** | The storage utilization percentage is computed based on total capacity and current capacity utilization. | [optional] [readonly] 
**UtilizationTrendPercentage** | Pointer to **float32** | The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data. | [optional] [readonly] 
**AssociatedProfile** | Pointer to [**PolicyAbstractProfileRelationship**](PolicyAbstractProfileRelationship.md) |  | [optional] 
**ChildClusters** | Pointer to [**[]HyperflexBaseClusterRelationship**](HyperflexBaseClusterRelationship.md) | An array of relationships to hyperflexBaseCluster resources. | [optional] 
**ParentCluster** | Pointer to [**HyperflexBaseClusterRelationship**](HyperflexBaseClusterRelationship.md) |  | [optional] 

## Methods

### NewHyperflexBaseClusterAllOf

`func NewHyperflexBaseClusterAllOf(classId string, objectType string, ) *HyperflexBaseClusterAllOf`

NewHyperflexBaseClusterAllOf instantiates a new HyperflexBaseClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexBaseClusterAllOfWithDefaults

`func NewHyperflexBaseClusterAllOfWithDefaults() *HyperflexBaseClusterAllOf`

NewHyperflexBaseClusterAllOfWithDefaults instantiates a new HyperflexBaseClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexBaseClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexBaseClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexBaseClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexBaseClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexBaseClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexBaseClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmSummary

`func (o *HyperflexBaseClusterAllOf) GetAlarmSummary() HyperflexAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *HyperflexBaseClusterAllOf) GetAlarmSummaryOk() (*HyperflexAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *HyperflexBaseClusterAllOf) SetAlarmSummary(v HyperflexAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *HyperflexBaseClusterAllOf) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *HyperflexBaseClusterAllOf) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *HyperflexBaseClusterAllOf) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetCapacityRunway

`func (o *HyperflexBaseClusterAllOf) GetCapacityRunway() int64`

GetCapacityRunway returns the CapacityRunway field if non-nil, zero value otherwise.

### GetCapacityRunwayOk

`func (o *HyperflexBaseClusterAllOf) GetCapacityRunwayOk() (*int64, bool)`

GetCapacityRunwayOk returns a tuple with the CapacityRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRunway

`func (o *HyperflexBaseClusterAllOf) SetCapacityRunway(v int64)`

SetCapacityRunway sets CapacityRunway field to given value.

### HasCapacityRunway

`func (o *HyperflexBaseClusterAllOf) HasCapacityRunway() bool`

HasCapacityRunway returns a boolean if a field has been set.

### GetClusterName

`func (o *HyperflexBaseClusterAllOf) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HyperflexBaseClusterAllOf) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HyperflexBaseClusterAllOf) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HyperflexBaseClusterAllOf) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterPurpose

`func (o *HyperflexBaseClusterAllOf) GetClusterPurpose() string`

GetClusterPurpose returns the ClusterPurpose field if non-nil, zero value otherwise.

### GetClusterPurposeOk

`func (o *HyperflexBaseClusterAllOf) GetClusterPurposeOk() (*string, bool)`

GetClusterPurposeOk returns a tuple with the ClusterPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPurpose

`func (o *HyperflexBaseClusterAllOf) SetClusterPurpose(v string)`

SetClusterPurpose sets ClusterPurpose field to given value.

### HasClusterPurpose

`func (o *HyperflexBaseClusterAllOf) HasClusterPurpose() bool`

HasClusterPurpose returns a boolean if a field has been set.

### GetComputeNodeCount

`func (o *HyperflexBaseClusterAllOf) GetComputeNodeCount() int64`

GetComputeNodeCount returns the ComputeNodeCount field if non-nil, zero value otherwise.

### GetComputeNodeCountOk

`func (o *HyperflexBaseClusterAllOf) GetComputeNodeCountOk() (*int64, bool)`

GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodeCount

`func (o *HyperflexBaseClusterAllOf) SetComputeNodeCount(v int64)`

SetComputeNodeCount sets ComputeNodeCount field to given value.

### HasComputeNodeCount

`func (o *HyperflexBaseClusterAllOf) HasComputeNodeCount() bool`

HasComputeNodeCount returns a boolean if a field has been set.

### GetConvergedNodeCount

`func (o *HyperflexBaseClusterAllOf) GetConvergedNodeCount() int64`

GetConvergedNodeCount returns the ConvergedNodeCount field if non-nil, zero value otherwise.

### GetConvergedNodeCountOk

`func (o *HyperflexBaseClusterAllOf) GetConvergedNodeCountOk() (*int64, bool)`

GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergedNodeCount

`func (o *HyperflexBaseClusterAllOf) SetConvergedNodeCount(v int64)`

SetConvergedNodeCount sets ConvergedNodeCount field to given value.

### HasConvergedNodeCount

`func (o *HyperflexBaseClusterAllOf) HasConvergedNodeCount() bool`

HasConvergedNodeCount returns a boolean if a field has been set.

### GetDeploymentType

`func (o *HyperflexBaseClusterAllOf) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *HyperflexBaseClusterAllOf) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *HyperflexBaseClusterAllOf) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *HyperflexBaseClusterAllOf) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetDriveType

`func (o *HyperflexBaseClusterAllOf) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *HyperflexBaseClusterAllOf) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *HyperflexBaseClusterAllOf) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *HyperflexBaseClusterAllOf) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetHxVersion

`func (o *HyperflexBaseClusterAllOf) GetHxVersion() string`

GetHxVersion returns the HxVersion field if non-nil, zero value otherwise.

### GetHxVersionOk

`func (o *HyperflexBaseClusterAllOf) GetHxVersionOk() (*string, bool)`

GetHxVersionOk returns a tuple with the HxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxVersion

`func (o *HyperflexBaseClusterAllOf) SetHxVersion(v string)`

SetHxVersion sets HxVersion field to given value.

### HasHxVersion

`func (o *HyperflexBaseClusterAllOf) HasHxVersion() bool`

HasHxVersion returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HyperflexBaseClusterAllOf) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HyperflexBaseClusterAllOf) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HyperflexBaseClusterAllOf) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HyperflexBaseClusterAllOf) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetStorageCapacity

`func (o *HyperflexBaseClusterAllOf) GetStorageCapacity() int64`

GetStorageCapacity returns the StorageCapacity field if non-nil, zero value otherwise.

### GetStorageCapacityOk

`func (o *HyperflexBaseClusterAllOf) GetStorageCapacityOk() (*int64, bool)`

GetStorageCapacityOk returns a tuple with the StorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacity

`func (o *HyperflexBaseClusterAllOf) SetStorageCapacity(v int64)`

SetStorageCapacity sets StorageCapacity field to given value.

### HasStorageCapacity

`func (o *HyperflexBaseClusterAllOf) HasStorageCapacity() bool`

HasStorageCapacity returns a boolean if a field has been set.

### GetStorageNodeCount

`func (o *HyperflexBaseClusterAllOf) GetStorageNodeCount() int64`

GetStorageNodeCount returns the StorageNodeCount field if non-nil, zero value otherwise.

### GetStorageNodeCountOk

`func (o *HyperflexBaseClusterAllOf) GetStorageNodeCountOk() (*int64, bool)`

GetStorageNodeCountOk returns a tuple with the StorageNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageNodeCount

`func (o *HyperflexBaseClusterAllOf) SetStorageNodeCount(v int64)`

SetStorageNodeCount sets StorageNodeCount field to given value.

### HasStorageNodeCount

`func (o *HyperflexBaseClusterAllOf) HasStorageNodeCount() bool`

HasStorageNodeCount returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *HyperflexBaseClusterAllOf) GetStorageUtilization() float32`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *HyperflexBaseClusterAllOf) GetStorageUtilizationOk() (*float32, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *HyperflexBaseClusterAllOf) SetStorageUtilization(v float32)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *HyperflexBaseClusterAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetUtilizationPercentage

`func (o *HyperflexBaseClusterAllOf) GetUtilizationPercentage() float32`

GetUtilizationPercentage returns the UtilizationPercentage field if non-nil, zero value otherwise.

### GetUtilizationPercentageOk

`func (o *HyperflexBaseClusterAllOf) GetUtilizationPercentageOk() (*float32, bool)`

GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercentage

`func (o *HyperflexBaseClusterAllOf) SetUtilizationPercentage(v float32)`

SetUtilizationPercentage sets UtilizationPercentage field to given value.

### HasUtilizationPercentage

`func (o *HyperflexBaseClusterAllOf) HasUtilizationPercentage() bool`

HasUtilizationPercentage returns a boolean if a field has been set.

### GetUtilizationTrendPercentage

`func (o *HyperflexBaseClusterAllOf) GetUtilizationTrendPercentage() float32`

GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field if non-nil, zero value otherwise.

### GetUtilizationTrendPercentageOk

`func (o *HyperflexBaseClusterAllOf) GetUtilizationTrendPercentageOk() (*float32, bool)`

GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationTrendPercentage

`func (o *HyperflexBaseClusterAllOf) SetUtilizationTrendPercentage(v float32)`

SetUtilizationTrendPercentage sets UtilizationTrendPercentage field to given value.

### HasUtilizationTrendPercentage

`func (o *HyperflexBaseClusterAllOf) HasUtilizationTrendPercentage() bool`

HasUtilizationTrendPercentage returns a boolean if a field has been set.

### GetAssociatedProfile

`func (o *HyperflexBaseClusterAllOf) GetAssociatedProfile() PolicyAbstractProfileRelationship`

GetAssociatedProfile returns the AssociatedProfile field if non-nil, zero value otherwise.

### GetAssociatedProfileOk

`func (o *HyperflexBaseClusterAllOf) GetAssociatedProfileOk() (*PolicyAbstractProfileRelationship, bool)`

GetAssociatedProfileOk returns a tuple with the AssociatedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedProfile

`func (o *HyperflexBaseClusterAllOf) SetAssociatedProfile(v PolicyAbstractProfileRelationship)`

SetAssociatedProfile sets AssociatedProfile field to given value.

### HasAssociatedProfile

`func (o *HyperflexBaseClusterAllOf) HasAssociatedProfile() bool`

HasAssociatedProfile returns a boolean if a field has been set.

### GetChildClusters

`func (o *HyperflexBaseClusterAllOf) GetChildClusters() []HyperflexBaseClusterRelationship`

GetChildClusters returns the ChildClusters field if non-nil, zero value otherwise.

### GetChildClustersOk

`func (o *HyperflexBaseClusterAllOf) GetChildClustersOk() (*[]HyperflexBaseClusterRelationship, bool)`

GetChildClustersOk returns a tuple with the ChildClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildClusters

`func (o *HyperflexBaseClusterAllOf) SetChildClusters(v []HyperflexBaseClusterRelationship)`

SetChildClusters sets ChildClusters field to given value.

### HasChildClusters

`func (o *HyperflexBaseClusterAllOf) HasChildClusters() bool`

HasChildClusters returns a boolean if a field has been set.

### SetChildClustersNil

`func (o *HyperflexBaseClusterAllOf) SetChildClustersNil(b bool)`

 SetChildClustersNil sets the value for ChildClusters to be an explicit nil

### UnsetChildClusters
`func (o *HyperflexBaseClusterAllOf) UnsetChildClusters()`

UnsetChildClusters ensures that no value is present for ChildClusters, not even an explicit nil
### GetParentCluster

`func (o *HyperflexBaseClusterAllOf) GetParentCluster() HyperflexBaseClusterRelationship`

GetParentCluster returns the ParentCluster field if non-nil, zero value otherwise.

### GetParentClusterOk

`func (o *HyperflexBaseClusterAllOf) GetParentClusterOk() (*HyperflexBaseClusterRelationship, bool)`

GetParentClusterOk returns a tuple with the ParentCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCluster

`func (o *HyperflexBaseClusterAllOf) SetParentCluster(v HyperflexBaseClusterRelationship)`

SetParentCluster sets ParentCluster field to given value.

### HasParentCluster

`func (o *HyperflexBaseClusterAllOf) HasParentCluster() bool`

HasParentCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


