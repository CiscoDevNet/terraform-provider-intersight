# TechsupportmanagementTechSupportBundleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIdentifier** | Pointer to **string** | The device identifier used to uniquely identify an individual device. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | The device type obtained from the inventory. | [optional] [readonly] 
**Pid** | Pointer to **string** | Product identification of the device. | [optional] 
**PlatformParam** | Pointer to [**ConnectorPlatformParamBase**](connector.PlatformParamBase.md) |  | [optional] 
**PlatformType** | Pointer to **string** | The platform type of the device. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - Intersight on-premise appliance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft HyperV system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A Dynatrace controller that monitors applications. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;MicrosoftAzure&#x60; - A Microsoft Azure target. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. | [optional] [default to ""]
**Serial** | Pointer to **string** | Serial number of the device. | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**TargetResource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**TechSupportStatus** | Pointer to [**TechsupportmanagementTechSupportStatusRelationship**](techsupportmanagement.TechSupportStatus.Relationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementTechSupportBundleAllOf

`func NewTechsupportmanagementTechSupportBundleAllOf() *TechsupportmanagementTechSupportBundleAllOf`

NewTechsupportmanagementTechSupportBundleAllOf instantiates a new TechsupportmanagementTechSupportBundleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportBundleAllOfWithDefaults

`func NewTechsupportmanagementTechSupportBundleAllOfWithDefaults() *TechsupportmanagementTechSupportBundleAllOf`

NewTechsupportmanagementTechSupportBundleAllOfWithDefaults instantiates a new TechsupportmanagementTechSupportBundleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetDeviceType

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetPid

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformParam

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformParam() ConnectorPlatformParamBase`

GetPlatformParam returns the PlatformParam field if non-nil, zero value otherwise.

### GetPlatformParamOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformParamOk() (*ConnectorPlatformParamBase, bool)`

GetPlatformParamOk returns a tuple with the PlatformParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformParam

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetPlatformParam(v ConnectorPlatformParamBase)`

SetPlatformParam sets PlatformParam field to given value.

### HasPlatformParam

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasPlatformParam() bool`

HasPlatformParam returns a boolean if a field has been set.

### GetPlatformType

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetSerial

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetTargetResource

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTargetResource() MoBaseMoRelationship`

GetTargetResource returns the TargetResource field if non-nil, zero value otherwise.

### GetTargetResourceOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTargetResourceOk() (*MoBaseMoRelationship, bool)`

GetTargetResourceOk returns a tuple with the TargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResource

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetTargetResource(v MoBaseMoRelationship)`

SetTargetResource sets TargetResource field to given value.

### HasTargetResource

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasTargetResource() bool`

HasTargetResource returns a boolean if a field has been set.

### GetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship`

GetTechSupportStatus returns the TechSupportStatus field if non-nil, zero value otherwise.

### GetTechSupportStatusOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool)`

GetTechSupportStatusOk returns a tuple with the TechSupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship)`

SetTechSupportStatus sets TechSupportStatus field to given value.

### HasTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasTechSupportStatus() bool`

HasTechSupportStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


