# HyperflexAppSettingConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.AppSettingConstraint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.AppSettingConstraint"]
**DeploymentType** | Pointer to **string** | The deployment type of the cluster. * &#x60;NA&#x60; - The deployment type of the HyperFlex cluster is not available. * &#x60;Datacenter&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * &#x60;Stretched Cluster&#x60; - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * &#x60;Edge&#x60; - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes. | [optional] [default to "NA"]
**HxdpVersion** | Pointer to **string** | The supported HyperFlex Data Platform version in regex format. | [optional] 
**HypervisorType** | Pointer to **string** | The hypervisor type for the HyperFlex cluster. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * &#x60;IWE&#x60; - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [default to "ESXi"]
**MgmtPlatform** | Pointer to **string** | The supported management platform for the HyperFlex Cluster. * &#x60;FI&#x60; - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * &#x60;EDGE&#x60; - The host servers used in the cluster deployment are standalone severs. | [optional] [default to "FI"]
**ServerModel** | Pointer to **string** | The supported server models in regex format. | [optional] 

## Methods

### NewHyperflexAppSettingConstraint

`func NewHyperflexAppSettingConstraint(classId string, objectType string, ) *HyperflexAppSettingConstraint`

NewHyperflexAppSettingConstraint instantiates a new HyperflexAppSettingConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAppSettingConstraintWithDefaults

`func NewHyperflexAppSettingConstraintWithDefaults() *HyperflexAppSettingConstraint`

NewHyperflexAppSettingConstraintWithDefaults instantiates a new HyperflexAppSettingConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexAppSettingConstraint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexAppSettingConstraint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexAppSettingConstraint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexAppSettingConstraint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexAppSettingConstraint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexAppSettingConstraint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeploymentType

`func (o *HyperflexAppSettingConstraint) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *HyperflexAppSettingConstraint) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *HyperflexAppSettingConstraint) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *HyperflexAppSettingConstraint) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetHxdpVersion

`func (o *HyperflexAppSettingConstraint) GetHxdpVersion() string`

GetHxdpVersion returns the HxdpVersion field if non-nil, zero value otherwise.

### GetHxdpVersionOk

`func (o *HyperflexAppSettingConstraint) GetHxdpVersionOk() (*string, bool)`

GetHxdpVersionOk returns a tuple with the HxdpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersion

`func (o *HyperflexAppSettingConstraint) SetHxdpVersion(v string)`

SetHxdpVersion sets HxdpVersion field to given value.

### HasHxdpVersion

`func (o *HyperflexAppSettingConstraint) HasHxdpVersion() bool`

HasHxdpVersion returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HyperflexAppSettingConstraint) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HyperflexAppSettingConstraint) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HyperflexAppSettingConstraint) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HyperflexAppSettingConstraint) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetMgmtPlatform

`func (o *HyperflexAppSettingConstraint) GetMgmtPlatform() string`

GetMgmtPlatform returns the MgmtPlatform field if non-nil, zero value otherwise.

### GetMgmtPlatformOk

`func (o *HyperflexAppSettingConstraint) GetMgmtPlatformOk() (*string, bool)`

GetMgmtPlatformOk returns a tuple with the MgmtPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPlatform

`func (o *HyperflexAppSettingConstraint) SetMgmtPlatform(v string)`

SetMgmtPlatform sets MgmtPlatform field to given value.

### HasMgmtPlatform

`func (o *HyperflexAppSettingConstraint) HasMgmtPlatform() bool`

HasMgmtPlatform returns a boolean if a field has been set.

### GetServerModel

`func (o *HyperflexAppSettingConstraint) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HyperflexAppSettingConstraint) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HyperflexAppSettingConstraint) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HyperflexAppSettingConstraint) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


