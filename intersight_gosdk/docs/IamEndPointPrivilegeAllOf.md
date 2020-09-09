# IamEndPointPrivilegeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The functionality of this privilege. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the end point privilege. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - Intersight on-premise appliance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft HyperV system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A Dynatrace controller that monitors applications. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;MicrosoftAzure&#x60; - A Microsoft Azure target. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. | [optional] [readonly] [default to ""]
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewIamEndPointPrivilegeAllOf

`func NewIamEndPointPrivilegeAllOf() *IamEndPointPrivilegeAllOf`

NewIamEndPointPrivilegeAllOf instantiates a new IamEndPointPrivilegeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointPrivilegeAllOfWithDefaults

`func NewIamEndPointPrivilegeAllOfWithDefaults() *IamEndPointPrivilegeAllOf`

NewIamEndPointPrivilegeAllOfWithDefaults instantiates a new IamEndPointPrivilegeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IamEndPointPrivilegeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamEndPointPrivilegeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamEndPointPrivilegeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamEndPointPrivilegeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamEndPointPrivilegeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamEndPointPrivilegeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamEndPointPrivilegeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamEndPointPrivilegeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IamEndPointPrivilegeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamEndPointPrivilegeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamEndPointPrivilegeAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamEndPointPrivilegeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSystem

`func (o *IamEndPointPrivilegeAllOf) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamEndPointPrivilegeAllOf) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamEndPointPrivilegeAllOf) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamEndPointPrivilegeAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


