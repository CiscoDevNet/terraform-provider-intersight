# IamEndPointRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the end point role. | [optional] [readonly] 
**RoleType** | Pointer to **string** | User specified tags to roles like as ep-admin or ep-readonly. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the end point like Cisco UCS Fabric Interconnect or Cisco IMC. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - Intersight on-premise appliance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft HyperV system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A Dynatrace controller that monitors applications. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;MicrosoftAzure&#x60; - A Microsoft Azure target. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. | [optional] [readonly] [default to ""]
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**EndPointPrivileges** | Pointer to [**[]IamEndPointPrivilegeRelationship**](iam.EndPointPrivilege.Relationship.md) | An array of relationships to iamEndPointPrivilege resources. | [optional] [readonly] 
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewIamEndPointRole

`func NewIamEndPointRole() *IamEndPointRole`

NewIamEndPointRole instantiates a new IamEndPointRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointRoleWithDefaults

`func NewIamEndPointRoleWithDefaults() *IamEndPointRole`

NewIamEndPointRoleWithDefaults instantiates a new IamEndPointRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamEndPointRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamEndPointRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamEndPointRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamEndPointRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleType

`func (o *IamEndPointRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *IamEndPointRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *IamEndPointRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *IamEndPointRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetType

`func (o *IamEndPointRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamEndPointRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamEndPointRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamEndPointRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *IamEndPointRole) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamEndPointRole) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamEndPointRole) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamEndPointRole) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEndPointPrivileges

`func (o *IamEndPointRole) GetEndPointPrivileges() []IamEndPointPrivilegeRelationship`

GetEndPointPrivileges returns the EndPointPrivileges field if non-nil, zero value otherwise.

### GetEndPointPrivilegesOk

`func (o *IamEndPointRole) GetEndPointPrivilegesOk() (*[]IamEndPointPrivilegeRelationship, bool)`

GetEndPointPrivilegesOk returns a tuple with the EndPointPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointPrivileges

`func (o *IamEndPointRole) SetEndPointPrivileges(v []IamEndPointPrivilegeRelationship)`

SetEndPointPrivileges sets EndPointPrivileges field to given value.

### HasEndPointPrivileges

`func (o *IamEndPointRole) HasEndPointPrivileges() bool`

HasEndPointPrivileges returns a boolean if a field has been set.

### SetEndPointPrivilegesNil

`func (o *IamEndPointRole) SetEndPointPrivilegesNil(b bool)`

 SetEndPointPrivilegesNil sets the value for EndPointPrivileges to be an explicit nil

### UnsetEndPointPrivileges
`func (o *IamEndPointRole) UnsetEndPointPrivileges()`

UnsetEndPointPrivileges ensures that no value is present for EndPointPrivileges, not even an explicit nil
### GetSystem

`func (o *IamEndPointRole) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamEndPointRole) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamEndPointRole) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamEndPointRole) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


