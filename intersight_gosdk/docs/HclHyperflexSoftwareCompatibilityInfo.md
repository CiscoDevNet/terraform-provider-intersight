# HclHyperflexSoftwareCompatibilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.HyperflexSoftwareCompatibilityInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.HyperflexSoftwareCompatibilityInfo"]
**Constraints** | Pointer to [**[]HclConstraint**](HclConstraint.md) |  | [optional] 
**HxdpVersion** | Pointer to **string** | HXDP component software version. | [optional] 
**HypervisorType** | Pointer to **string** | Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * &#x60;IWE&#x60; - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [default to "ESXi"]
**HypervisorVersion** | Pointer to **string** | Hypervisor component software version. | [optional] 
**IsMgmtBuild** | Pointer to **string** | Type of the HXDP bundle mgmt or full. | [optional] 
**MaxMgmtVersion** | Pointer to **string** | Maximum supported HyperFlex Data Platform build version. | [optional] 
**MinMgmtVersion** | Pointer to **string** | Minimum supported HyperFlex Data Platform build version. | [optional] 
**ServerFwVersion** | Pointer to **string** | UCS Server Firmware component software version. | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 

## Methods

### NewHclHyperflexSoftwareCompatibilityInfo

`func NewHclHyperflexSoftwareCompatibilityInfo(classId string, objectType string, ) *HclHyperflexSoftwareCompatibilityInfo`

NewHclHyperflexSoftwareCompatibilityInfo instantiates a new HclHyperflexSoftwareCompatibilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclHyperflexSoftwareCompatibilityInfoWithDefaults

`func NewHclHyperflexSoftwareCompatibilityInfoWithDefaults() *HclHyperflexSoftwareCompatibilityInfo`

NewHclHyperflexSoftwareCompatibilityInfoWithDefaults instantiates a new HclHyperflexSoftwareCompatibilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraints

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetConstraints() []HclConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetConstraintsOk() (*[]HclConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetConstraints(v []HclConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *HclHyperflexSoftwareCompatibilityInfo) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetHxdpVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetHxdpVersion() string`

GetHxdpVersion returns the HxdpVersion field if non-nil, zero value otherwise.

### GetHxdpVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetHxdpVersionOk() (*string, bool)`

GetHxdpVersionOk returns a tuple with the HxdpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetHxdpVersion(v string)`

SetHxdpVersion sets HxdpVersion field to given value.

### HasHxdpVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasHxdpVersion() bool`

HasHxdpVersion returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetIsMgmtBuild

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetIsMgmtBuild() string`

GetIsMgmtBuild returns the IsMgmtBuild field if non-nil, zero value otherwise.

### GetIsMgmtBuildOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetIsMgmtBuildOk() (*string, bool)`

GetIsMgmtBuildOk returns a tuple with the IsMgmtBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMgmtBuild

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetIsMgmtBuild(v string)`

SetIsMgmtBuild sets IsMgmtBuild field to given value.

### HasIsMgmtBuild

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasIsMgmtBuild() bool`

HasIsMgmtBuild returns a boolean if a field has been set.

### GetMaxMgmtVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetMaxMgmtVersion() string`

GetMaxMgmtVersion returns the MaxMgmtVersion field if non-nil, zero value otherwise.

### GetMaxMgmtVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetMaxMgmtVersionOk() (*string, bool)`

GetMaxMgmtVersionOk returns a tuple with the MaxMgmtVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMgmtVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetMaxMgmtVersion(v string)`

SetMaxMgmtVersion sets MaxMgmtVersion field to given value.

### HasMaxMgmtVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasMaxMgmtVersion() bool`

HasMaxMgmtVersion returns a boolean if a field has been set.

### GetMinMgmtVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetMinMgmtVersion() string`

GetMinMgmtVersion returns the MinMgmtVersion field if non-nil, zero value otherwise.

### GetMinMgmtVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetMinMgmtVersionOk() (*string, bool)`

GetMinMgmtVersionOk returns a tuple with the MinMgmtVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMgmtVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetMinMgmtVersion(v string)`

SetMinMgmtVersion sets MinMgmtVersion field to given value.

### HasMinMgmtVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasMinMgmtVersion() bool`

HasMinMgmtVersion returns a boolean if a field has been set.

### GetServerFwVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetServerFwVersion() string`

GetServerFwVersion returns the ServerFwVersion field if non-nil, zero value otherwise.

### GetServerFwVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetServerFwVersionOk() (*string, bool)`

GetServerFwVersionOk returns a tuple with the ServerFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFwVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetServerFwVersion(v string)`

SetServerFwVersion sets ServerFwVersion field to given value.

### HasServerFwVersion

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasServerFwVersion() bool`

HasServerFwVersion returns a boolean if a field has been set.

### GetAppCatalog

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HclHyperflexSoftwareCompatibilityInfo) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HclHyperflexSoftwareCompatibilityInfo) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HclHyperflexSoftwareCompatibilityInfo) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


