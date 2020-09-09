package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAssetManagedDevice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssetManagedDeviceRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"credential": {
				Description: "Credentials to manage Managed Device.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"is_password_set": {
							Description: "Indicates whether the value of the 'password' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Description: "Password for the Managed Device.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"username": {
							Description: "Username for the Managed Device. Format and restrictions are not enforced here but usually follow the ManagedDevice requirements.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"device_connector_manager": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"device_type": {
				Description: "Type of the Device such as VMware, Pure Storage supported by Intersight Assist.\n* `` - The device reported an empty or unrecognized platform type.\n* `APIC` - An Application Policy Infrastructure Controller cluster.\n* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.\n* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).\n* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.\n* `IMC` - A standalone UCS Server Integrated Management Controller.\n* `IMCM4` - A standalone UCS M4 Server.\n* `IMCM5` - A standalone UCS M5 server.\n* `UCSIOM` - An UCS Chassis IO module.\n* `HX` - A HyperFlex storage controller.\n* `HyperFlexAP` - A HyperFlex Application Platform.\n* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.\n* `IntersightAppliance` - Intersight on-premise appliance.\n* `PureStorageFlashArray` - A Pure Storage FlashArray device.\n* `NetAppOntap` - A NetApp ONTAP storage system.\n* `EmcScaleIo` - An EMC ScaleIO storage system.\n* `EmcVmax` - An EMC VMAX storage system.\n* `EmcVplex` - An EMC VPLEX storage system.\n* `EmcXtremIo` - An EMC XtremIO storage system.\n* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.\n* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.\n* `AppDynamics` - An AppDynamics controller that monitors applications.\n* `Dynatrace` - A Dynatrace controller that monitors applications.\n* `MicrosoftSqlServer` - A Microsoft SQL database server.\n* `Kubernetes` - A Kubernetes cluster that runs containerized applications.\n* `MicrosoftAzure` - A Microsoft Azure target.\n* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.\n* `IMCBlade` - An Intersight managed UCS Blade Server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ignore_cert": {
				Description: "Ignore Certificates with protocol https for connecting to the Managed Device. It is not used for other protocols.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_enabled": {
				Description: "Device is Enabled/Disabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"management_address": {
				Description: "Management address of the device. It can be IPv4, IPv6 or FQDN.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Description: "Port to use for connecting to the Managed Device. Port is optional. If not specified, default port for protocol is used.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"protocol": {
				Description: "Protocol to use for connecting to the Managed Device.\n* `https` - HTTPS communication is used for connecting to Device.\n* `http` - HTTP communication is used for connecting to Device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Description: "Status of communication releated to Managed Device.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"cloud_port": {
							Description: "Port used for the connection to the Cloud by the Device Connector for the Managed Device.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"connection_failure_reason": {
							Description: "Maintains the reason for the failure of connection to the Device in case of connection failure.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"connection_status": {
							Description: "Maintains the status of the connection to the Device.\n* `Unknown` - This represents Connection Status as unknown.\n* `Success` - This represents Connection Status as successful.\n* `Failure` - This represents Connection Status as failure.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"error_code": {
							Description: "Maintains code related to error from Device Connector, if any.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"error_reason": {
							Description: "Maintains the reason for the error from Device Connector, if any.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"process_id": {
							Description: "Maintains the process pid of the Device Connector for the Managed Device.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"server_port": {
							Description: "Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"state": {
							Description: "Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.\n* `New` - State when the Managed Device is created.\n* `StartSent` - State when Create Managed Device is sent successfully to Intersight Assist to start Managed Device Device Connector.\n* `StartSentFailure` - State when Create Managed Device failed to send start Device Connector to Intersight Assist.\n* `StartSuccess` - State when Device Connector for the Managed Device has been successfully started.\n* `StartFailure` - State when Device Connector for the Managed Device failed to start.\n* `UpdateSentFailure` - State when Update Managed Device request is failed to be sent to Device Connector for the Managed Device.\n* `UpdateSent` - State when Update Managed Device request is successfully sent to Device Connector for the Managed Device.\n* `DeleteSentFailure` - State when Delete Managed Device request is failed to be sent to Intersight Assist.\n* `DeleteInProgress` - State when Intersight Assist received delete Managed Device request successfully.\n* `DeleteFailure` - State when Intersight Assist failed to stop Device Connector for the Managed Device.\n* `DeleteSuccess` - State when Intersight Assist stopped Device Connector for the Managed Device successfully.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"workflow_info": {
				Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAssetManagedDeviceRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewAssetManagedDeviceWithDefaults()
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.SetDeviceType(x)
	}
	if v, ok := d.GetOk("ignore_cert"); ok {
		x := (v.(bool))
		o.SetIgnoreCert(x)
	}
	if v, ok := d.GetOk("is_enabled"); ok {
		x := (v.(bool))
		o.SetIsEnabled(x)
	}
	if v, ok := d.GetOk("management_address"); ok {
		x := (v.(string))
		o.SetManagementAddress(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("port"); ok {
		x := int64(v.(int))
		o.SetPort(x)
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.AssetApi.GetAssetManagedDeviceList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.AssetManagedDeviceList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to AssetManagedDevice: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewAssetManagedDeviceWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}

			if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Account: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("credential", flattenMapCommCredential(s.GetCredential(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Credential: %+v", err)
			}

			if err := d.Set("device_connector_manager", flattenMapAssetDeviceRegistrationRelationship(s.GetDeviceConnectorManager(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceConnectorManager: %+v", err)
			}
			if err := d.Set("device_type", (s.GetDeviceType())); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceType: %+v", err)
			}
			if err := d.Set("ignore_cert", (s.GetIgnoreCert())); err != nil {
				return fmt.Errorf("error occurred while setting property IgnoreCert: %+v", err)
			}
			if err := d.Set("is_enabled", (s.GetIsEnabled())); err != nil {
				return fmt.Errorf("error occurred while setting property IsEnabled: %+v", err)
			}
			if err := d.Set("management_address", (s.GetManagementAddress())); err != nil {
				return fmt.Errorf("error occurred while setting property ManagementAddress: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("port", (s.GetPort())); err != nil {
				return fmt.Errorf("error occurred while setting property Port: %+v", err)
			}
			if err := d.Set("protocol", (s.GetProtocol())); err != nil {
				return fmt.Errorf("error occurred while setting property Protocol: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}

			if err := d.Set("status", flattenMapAssetManagedDeviceStatus(s.GetStatus(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Status: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}

			if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property WorkflowInfo: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
