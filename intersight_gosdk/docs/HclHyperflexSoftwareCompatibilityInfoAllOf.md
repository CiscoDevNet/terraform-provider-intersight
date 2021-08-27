# HclHyperflexSoftwareCompatibilityInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.HyperflexSoftwareCompatibilityInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.HyperflexSoftwareCompatibilityInfo"]
**Constraints** | Pointer to [**[]HclConstraint**](HclConstraint.md) |  | [optional] 
**HxdpVersion** | Pointer to **string** | HXDP component software version. | [optional] 
**HypervisorType** | Pointer to **string** | Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM. * &#x60;ESXi&#x60; - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * &#x60;HyperFlexAp&#x60; - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform. * &#x60;Hyper-V&#x60; - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * &#x60;Unknown&#x60; - The hypervisor running on the HyperFlex cluster is not known. | [optional] [default to "ESXi"]
**HypervisorVersion** | Pointer to **string** | Hypervisor component software version. | [optional] 
**ServerFwVersion** | Pointer to **string** | UCS Server Firmware component software version. | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 

## Methods

### NewHclHyperflexSoftwareCompatibilityInfoAllOf

`func NewHclHyperflexSoftwareCompatibilityInfoAllOf(classId string, objectType string, ) *HclHyperflexSoftwareCompatibilityInfoAllOf`

NewHclHyperflexSoftwareCompatibilityInfoAllOf instantiates a new HclHyperflexSoftwareCompatibilityInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults

`func NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults() *HclHyperflexSoftwareCompatibilityInfoAllOf`

NewHclHyperflexSoftwareCompatibilityInfoAllOfWithDefaults instantiates a new HclHyperflexSoftwareCompatibilityInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetConstraintsNil

`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *HclHyperflexSoftwareCompatibilityInfoAllOf) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
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


