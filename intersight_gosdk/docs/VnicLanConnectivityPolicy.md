# VnicLanConnectivityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.LanConnectivityPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.LanConnectivityPolicy"]
**AzureQosEnabled** | Pointer to **bool** | Enabling AzureStack-Host QoS on an adapter allows the user to carve out traffic classes for RDMA traffic which ensures that a desired portion of the bandwidth is allocated to it. | [optional] [default to false]
**IqnAllocationType** | Pointer to **string** | Allocation Type of iSCSI Qualified Name - Static/Pool/None. * &#x60;None&#x60; - Type indicates that there is no IQN associated to an interface. * &#x60;Static&#x60; - Type represents that static IQN is associated to an interface. * &#x60;Pool&#x60; - Type indicates that IQN value is sourced from an associated pool. | [optional] [default to "None"]
**PlacementMode** | Pointer to **string** | The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * &#x60;custom&#x60; - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * &#x60;auto&#x60; - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. | [optional] [default to "custom"]
**StaticIqnName** | Pointer to **string** | User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain. | [optional] 
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "Standalone"]
**EthIfs** | Pointer to [**[]VnicEthIfRelationship**](VnicEthIfRelationship.md) | An array of relationships to vnicEthIf resources. | [optional] 
**IqnPool** | Pointer to [**IqnpoolPoolRelationship**](iqnpool.Pool.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewVnicLanConnectivityPolicy

`func NewVnicLanConnectivityPolicy(classId string, objectType string, ) *VnicLanConnectivityPolicy`

NewVnicLanConnectivityPolicy instantiates a new VnicLanConnectivityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicLanConnectivityPolicyWithDefaults

`func NewVnicLanConnectivityPolicyWithDefaults() *VnicLanConnectivityPolicy`

NewVnicLanConnectivityPolicyWithDefaults instantiates a new VnicLanConnectivityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicLanConnectivityPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicLanConnectivityPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicLanConnectivityPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicLanConnectivityPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicLanConnectivityPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicLanConnectivityPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAzureQosEnabled

`func (o *VnicLanConnectivityPolicy) GetAzureQosEnabled() bool`

GetAzureQosEnabled returns the AzureQosEnabled field if non-nil, zero value otherwise.

### GetAzureQosEnabledOk

`func (o *VnicLanConnectivityPolicy) GetAzureQosEnabledOk() (*bool, bool)`

GetAzureQosEnabledOk returns a tuple with the AzureQosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureQosEnabled

`func (o *VnicLanConnectivityPolicy) SetAzureQosEnabled(v bool)`

SetAzureQosEnabled sets AzureQosEnabled field to given value.

### HasAzureQosEnabled

`func (o *VnicLanConnectivityPolicy) HasAzureQosEnabled() bool`

HasAzureQosEnabled returns a boolean if a field has been set.

### GetIqnAllocationType

`func (o *VnicLanConnectivityPolicy) GetIqnAllocationType() string`

GetIqnAllocationType returns the IqnAllocationType field if non-nil, zero value otherwise.

### GetIqnAllocationTypeOk

`func (o *VnicLanConnectivityPolicy) GetIqnAllocationTypeOk() (*string, bool)`

GetIqnAllocationTypeOk returns a tuple with the IqnAllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnAllocationType

`func (o *VnicLanConnectivityPolicy) SetIqnAllocationType(v string)`

SetIqnAllocationType sets IqnAllocationType field to given value.

### HasIqnAllocationType

`func (o *VnicLanConnectivityPolicy) HasIqnAllocationType() bool`

HasIqnAllocationType returns a boolean if a field has been set.

### GetPlacementMode

`func (o *VnicLanConnectivityPolicy) GetPlacementMode() string`

GetPlacementMode returns the PlacementMode field if non-nil, zero value otherwise.

### GetPlacementModeOk

`func (o *VnicLanConnectivityPolicy) GetPlacementModeOk() (*string, bool)`

GetPlacementModeOk returns a tuple with the PlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementMode

`func (o *VnicLanConnectivityPolicy) SetPlacementMode(v string)`

SetPlacementMode sets PlacementMode field to given value.

### HasPlacementMode

`func (o *VnicLanConnectivityPolicy) HasPlacementMode() bool`

HasPlacementMode returns a boolean if a field has been set.

### GetStaticIqnName

`func (o *VnicLanConnectivityPolicy) GetStaticIqnName() string`

GetStaticIqnName returns the StaticIqnName field if non-nil, zero value otherwise.

### GetStaticIqnNameOk

`func (o *VnicLanConnectivityPolicy) GetStaticIqnNameOk() (*string, bool)`

GetStaticIqnNameOk returns a tuple with the StaticIqnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIqnName

`func (o *VnicLanConnectivityPolicy) SetStaticIqnName(v string)`

SetStaticIqnName sets StaticIqnName field to given value.

### HasStaticIqnName

`func (o *VnicLanConnectivityPolicy) HasStaticIqnName() bool`

HasStaticIqnName returns a boolean if a field has been set.

### GetTargetPlatform

`func (o *VnicLanConnectivityPolicy) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *VnicLanConnectivityPolicy) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *VnicLanConnectivityPolicy) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *VnicLanConnectivityPolicy) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetEthIfs

`func (o *VnicLanConnectivityPolicy) GetEthIfs() []VnicEthIfRelationship`

GetEthIfs returns the EthIfs field if non-nil, zero value otherwise.

### GetEthIfsOk

`func (o *VnicLanConnectivityPolicy) GetEthIfsOk() (*[]VnicEthIfRelationship, bool)`

GetEthIfsOk returns a tuple with the EthIfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthIfs

`func (o *VnicLanConnectivityPolicy) SetEthIfs(v []VnicEthIfRelationship)`

SetEthIfs sets EthIfs field to given value.

### HasEthIfs

`func (o *VnicLanConnectivityPolicy) HasEthIfs() bool`

HasEthIfs returns a boolean if a field has been set.

### SetEthIfsNil

`func (o *VnicLanConnectivityPolicy) SetEthIfsNil(b bool)`

 SetEthIfsNil sets the value for EthIfs to be an explicit nil

### UnsetEthIfs
`func (o *VnicLanConnectivityPolicy) UnsetEthIfs()`

UnsetEthIfs ensures that no value is present for EthIfs, not even an explicit nil
### GetIqnPool

`func (o *VnicLanConnectivityPolicy) GetIqnPool() IqnpoolPoolRelationship`

GetIqnPool returns the IqnPool field if non-nil, zero value otherwise.

### GetIqnPoolOk

`func (o *VnicLanConnectivityPolicy) GetIqnPoolOk() (*IqnpoolPoolRelationship, bool)`

GetIqnPoolOk returns a tuple with the IqnPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnPool

`func (o *VnicLanConnectivityPolicy) SetIqnPool(v IqnpoolPoolRelationship)`

SetIqnPool sets IqnPool field to given value.

### HasIqnPool

`func (o *VnicLanConnectivityPolicy) HasIqnPool() bool`

HasIqnPool returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicLanConnectivityPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicLanConnectivityPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicLanConnectivityPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicLanConnectivityPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *VnicLanConnectivityPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *VnicLanConnectivityPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *VnicLanConnectivityPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *VnicLanConnectivityPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *VnicLanConnectivityPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *VnicLanConnectivityPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


