# HclHyperflexSoftwareCompatibilityInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]HclConstraint**](hcl.Constraint.md) |  | [optional] 
**HxdpVersion** | Pointer to **string** | HXDP component software version. | [optional] 
**HypervisorType** | Pointer to **string** | Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM. * &#x60;ESXi&#x60; - ESXi hypervisor as specified by the user. * &#x60;HYPERV&#x60; - Hyperv hypervisor as specified by the user. * &#x60;KVM&#x60; - KVM hypervisor as specified by the user. | [optional] [default to "ESXi"]
**HypervisorVersion** | Pointer to **string** | Hypervisor component software version. | [optional] 
**ServerFwVersion** | Pointer to **string** | UCS Server Firmware component software version. | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](hyperflex.AppCatalog.Relationship.md) |  | [optional] 

## Methods

### NewHclHyperflexSoftwareCompatibilityInfoAllOf

`func NewHclHyperflexSoftwareCompatibilityInfoAllOf() *HclHyperflexSoftwareCompatibilityInfoAllOf`

NewHclHyperflexSoftwareCompatibilityInfoAllOf instantiates a new HclHyperflexSoftwareCompatibilityInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults

`func NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults() *HclHyperflexSoftwareCompatibilityInfoAllOf`

NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults instantiates a new HclHyperflexSoftwareCompatibilityInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetConstraints() []HclConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetConstraintsOk() (*[]HclConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetConstraints(v []HclConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetHxdpVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHxdpVersion() string`

GetHxdpVersion returns the HxdpVersion field if non-nil, zero value otherwise.

### GetHxdpVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHxdpVersionOk() (*string, bool)`

GetHxdpVersionOk returns a tuple with the HxdpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetHxdpVersion(v string)`

SetHxdpVersion sets HxdpVersion field to given value.

### HasHxdpVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasHxdpVersion() bool`

HasHxdpVersion returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetServerFwVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetServerFwVersion() string`

GetServerFwVersion returns the ServerFwVersion field if non-nil, zero value otherwise.

### GetServerFwVersionOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetServerFwVersionOk() (*string, bool)`

GetServerFwVersionOk returns a tuple with the ServerFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFwVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetServerFwVersion(v string)`

SetServerFwVersion sets ServerFwVersion field to given value.

### HasServerFwVersion

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasServerFwVersion() bool`

HasServerFwVersion returns a boolean if a field has been set.

### GetAppCatalog

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


