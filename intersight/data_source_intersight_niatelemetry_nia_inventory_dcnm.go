package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getNiatelemetryNiaInventoryDcnmSchema() map[string]*schema.Schema {
	var schemaMap = make(map[string]*schema.Schema)
	schemaMap = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"controller_health": {
			Description: "Health of controller on DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dev": {
			Description: "Returns the value of the dev Field.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"epld_image_count": {
			Description: "Number of EPLD images uploaded to DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"golden_image_details": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"image_name": {
						Description: "Returns name of the image on controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Returns name of the image on controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Returns version of the image on controller.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"ha_enabled": {
			Description: "Returns the value of the haEnabled field.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"ha_replication_status": {
			Description: "Returns the value of the haReplicationStatus field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"install": {
			Description: "Returns the value of the install field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"installation_type": {
			Description: "Installation type of controller on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"installation_type_description": {
			Description: "Installation type description of controller on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_isn_configured": {
			Description: "Returns true if ISN is configured.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_media_controller": {
			Description: "Returns the value of the isMediaController field.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"is_smart_license_enabled": {
			Description: "Returns true if the Smart license is enabled and is in use.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mode": {
			Description: "Mode of controller on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ndfc_fabric_name": {
			Description: "NDFC name information of the setup.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ndfc_oper_state": {
			Description: "NDFC status information for the setup.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"network_info": {
			Description: "Installation type description of controller on DCNM.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"active_node": {
						Description: "Returns details of active node.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"hostname": {
									Description: "Returns hostname of the node.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"managementt_ip": {
									Description: "Returns management IP of the node.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"outofband_ip": {
									Description: "Returns out of band IP of the node.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hostname": {
						Description: "Returns hostname of the network.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"managementt_ip": {
						Description: "Returns management IP of the network.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"outofband_ip": {
						Description: "Returns out of band IP of the network.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"standby_node": {
						Description: "Returns details of standby node.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"hostname": {
									Description: "Returns hostname of the node.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"managementt_ip": {
									Description: "Returns management IP of the node.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"outofband_ip": {
									Description: "Returns out of band IP of the node.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
				},
			},
		},
		"num_dcnm_site": {
			Description: "Returns the number of DCNM site fabrics.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fabrics": {
			Description: "Returns total number of fabrics in DCNM set-up.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fabrics_in_msd": {
			Description: "Returns the number of fabrics in msd.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_ingress_replication_fabrics": {
			Description: "Returns the number of fabrics that have ingress replication type.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_local_users": {
			Description: "Returns the number of local users other than admin user.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_msd": {
			Description: "Returns the number of MSD fabrics.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_svi_vrf_count": {
			Description: "Returns the number of svi interfaces configured for VRF vlans.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_trm_enabled_count": {
			Description: "Returns the number of links where TRM is enabled.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_upg_users": {
			Description: "Number of users who have upgrade privileges excluding the admin.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"nxos_image_count": {
			Description: "Number of NXOS images uploaded to DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"outofband_ip": {
			Description: "Out of band IP of controller on DCNM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"record_type": {
			Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"record_version": {
			Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"registered_device": {
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"serial": {
			Description: "Serial number of device being inventoried. The serial number is unique per device.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"site_name": {
			Description: "Name of fabric domain of the controller.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"underlay_peering_active_links_count": {
			Description: "Returns the number of underlay peering active links.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"upg_job_count": {
			Description: "Number of upgrade jobs configured on DCNM.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"upg_status": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"job_id": {
						Description: "Returns the id of the job.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"upg_status": {
						Description: "Returns the status of the job.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"nr_version": {
			Description: "Returns the value of the version field.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
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
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"marked_for_deletion": {
						Description: "The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	return schemaMap
}

func dataSourceNiatelemetryNiaInventoryDcnm() *schema.Resource {
	var subSchema = getNiatelemetryNiaInventoryDcnmSchema()
	var model = getNiatelemetryNiaInventoryDcnmSchema()
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNiaInventoryDcnmRead,
		Schema:      model}
}

func dataSourceNiatelemetryNiaInventoryDcnmRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNiaInventoryDcnm{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOkExists("controller_health"); ok {
		x := int64(v.(int))
		o.SetControllerHealth(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOkExists("dev"); ok {
		x := (v.(bool))
		o.SetDev(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOkExists("epld_image_count"); ok {
		x := int64(v.(int))
		o.SetEpldImageCount(x)
	}

	if v, ok := d.GetOk("golden_image_details"); ok {
		x := make([]models.NiatelemetryImageDetail, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryImageDetail{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("niatelemetry.ImageDetail")
			if v, ok := l["image_name"]; ok {
				{
					x := (v.(string))
					o.SetImageName(x)
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := (v.(string))
					o.SetVersion(x)
				}
			}
			x = append(x, *o)
		}
		o.SetGoldenImageDetails(x)
	}

	if v, ok := d.GetOkExists("ha_enabled"); ok {
		x := (v.(bool))
		o.SetHaEnabled(x)
	}

	if v, ok := d.GetOk("ha_replication_status"); ok {
		x := (v.(string))
		o.SetHaReplicationStatus(x)
	}

	if v, ok := d.GetOk("install"); ok {
		x := (v.(string))
		o.SetInstall(x)
	}

	if v, ok := d.GetOk("installation_type"); ok {
		x := (v.(string))
		o.SetInstallationType(x)
	}

	if v, ok := d.GetOk("installation_type_description"); ok {
		x := (v.(string))
		o.SetInstallationTypeDescription(x)
	}

	if v, ok := d.GetOkExists("is_isn_configured"); ok {
		x := (v.(bool))
		o.SetIsIsnConfigured(x)
	}

	if v, ok := d.GetOkExists("is_media_controller"); ok {
		x := (v.(bool))
		o.SetIsMediaController(x)
	}

	if v, ok := d.GetOkExists("is_smart_license_enabled"); ok {
		x := (v.(bool))
		o.SetIsSmartLicenseEnabled(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("mode"); ok {
		x := (v.(string))
		o.SetMode(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("ndfc_fabric_name"); ok {
		x := (v.(string))
		o.SetNdfcFabricName(x)
	}

	if v, ok := d.GetOk("ndfc_oper_state"); ok {
		x := (v.(string))
		o.SetNdfcOperState(x)
	}

	if v, ok := d.GetOk("network_info"); ok {
		p := make([]models.NiatelemetryNetworkInfo, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.NiatelemetryNetworkInfo{}
			if v, ok := l["active_node"]; ok {
				{
					p := make([]models.NiatelemetryNode, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewNiatelemetryNodeWithDefaults()
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("niatelemetry.Node")
						if v, ok := l["hostname"]; ok {
							{
								x := (v.(string))
								o.SetHostname(x)
							}
						}
						if v, ok := l["managementt_ip"]; ok {
							{
								x := (v.(string))
								o.SetManagementtIp(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["outofband_ip"]; ok {
							{
								x := (v.(string))
								o.SetOutofbandIp(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetActiveNode(x)
					}
				}
			}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("niatelemetry.NetworkInfo")
			if v, ok := l["hostname"]; ok {
				{
					x := (v.(string))
					o.SetHostname(x)
				}
			}
			if v, ok := l["managementt_ip"]; ok {
				{
					x := (v.(string))
					o.SetManagementtIp(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["outofband_ip"]; ok {
				{
					x := (v.(string))
					o.SetOutofbandIp(x)
				}
			}
			if v, ok := l["standby_node"]; ok {
				{
					p := make([]models.NiatelemetryNode, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewNiatelemetryNodeWithDefaults()
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("niatelemetry.Node")
						if v, ok := l["hostname"]; ok {
							{
								x := (v.(string))
								o.SetHostname(x)
							}
						}
						if v, ok := l["managementt_ip"]; ok {
							{
								x := (v.(string))
								o.SetManagementtIp(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["outofband_ip"]; ok {
							{
								x := (v.(string))
								o.SetOutofbandIp(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetStandbyNode(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetNetworkInfo(x)
		}
	}

	if v, ok := d.GetOkExists("num_dcnm_site"); ok {
		x := int64(v.(int))
		o.SetNumDcnmSite(x)
	}

	if v, ok := d.GetOkExists("num_fabrics"); ok {
		x := int64(v.(int))
		o.SetNumFabrics(x)
	}

	if v, ok := d.GetOkExists("num_fabrics_in_msd"); ok {
		x := int64(v.(int))
		o.SetNumFabricsInMsd(x)
	}

	if v, ok := d.GetOkExists("num_ingress_replication_fabrics"); ok {
		x := int64(v.(int))
		o.SetNumIngressReplicationFabrics(x)
	}

	if v, ok := d.GetOkExists("num_local_users"); ok {
		x := int64(v.(int))
		o.SetNumLocalUsers(x)
	}

	if v, ok := d.GetOkExists("num_msd"); ok {
		x := int64(v.(int))
		o.SetNumMsd(x)
	}

	if v, ok := d.GetOkExists("num_svi_vrf_count"); ok {
		x := int64(v.(int))
		o.SetNumSviVrfCount(x)
	}

	if v, ok := d.GetOkExists("num_trm_enabled_count"); ok {
		x := int64(v.(int))
		o.SetNumTrmEnabledCount(x)
	}

	if v, ok := d.GetOkExists("num_upg_users"); ok {
		x := int64(v.(int))
		o.SetNumUpgUsers(x)
	}

	if v, ok := d.GetOkExists("nxos_image_count"); ok {
		x := int64(v.(int))
		o.SetNxosImageCount(x)
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	if v, ok := d.GetOk("outofband_ip"); ok {
		x := (v.(string))
		o.SetOutofbandIp(x)
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}

	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}

	if v, ok := d.GetOk("registered_device"); ok {
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOkExists("underlay_peering_active_links_count"); ok {
		x := int64(v.(int))
		o.SetUnderlayPeeringActiveLinksCount(x)
	}

	if v, ok := d.GetOkExists("upg_job_count"); ok {
		x := int64(v.(int))
		o.SetUpgJobCount(x)
	}

	if v, ok := d.GetOk("upg_status"); ok {
		x := make([]models.NiatelemetryJobDetail, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.NiatelemetryJobDetail{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("niatelemetry.JobDetail")
			if v, ok := l["job_id"]; ok {
				{
					x := int64(v.(int))
					o.SetJobId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["upg_status"]; ok {
				{
					x := (v.(string))
					o.SetUpgStatus(x)
				}
			}
			x = append(x, *o)
		}
		o.SetUpgStatus(x)
	}

	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNiaInventoryDcnm object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryDcnmList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventoryDcnm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of NiatelemetryNiaInventoryDcnm: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for NiatelemetryNiaInventoryDcnm data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var niatelemetryNiaInventoryDcnmResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNiaInventoryDcnmList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching NiatelemetryNiaInventoryDcnm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching NiatelemetryNiaInventoryDcnm: %s", responseErr.Error())
		}
		results := resMo.NiatelemetryNiaInventoryDcnmList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())
				temp["controller_health"] = (s.GetControllerHealth())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["dev"] = (s.GetDev())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["epld_image_count"] = (s.GetEpldImageCount())

				temp["golden_image_details"] = flattenListNiatelemetryImageDetail(s.GetGoldenImageDetails(), d)
				temp["ha_enabled"] = (s.GetHaEnabled())
				temp["ha_replication_status"] = (s.GetHaReplicationStatus())
				temp["install"] = (s.GetInstall())
				temp["installation_type"] = (s.GetInstallationType())
				temp["installation_type_description"] = (s.GetInstallationTypeDescription())
				temp["is_isn_configured"] = (s.GetIsIsnConfigured())
				temp["is_media_controller"] = (s.GetIsMediaController())
				temp["is_smart_license_enabled"] = (s.GetIsSmartLicenseEnabled())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["mode"] = (s.GetMode())
				temp["moid"] = (s.GetMoid())
				temp["ndfc_fabric_name"] = (s.GetNdfcFabricName())
				temp["ndfc_oper_state"] = (s.GetNdfcOperState())

				temp["network_info"] = flattenMapNiatelemetryNetworkInfo(s.GetNetworkInfo(), d)
				temp["num_dcnm_site"] = (s.GetNumDcnmSite())
				temp["num_fabrics"] = (s.GetNumFabrics())
				temp["num_fabrics_in_msd"] = (s.GetNumFabricsInMsd())
				temp["num_ingress_replication_fabrics"] = (s.GetNumIngressReplicationFabrics())
				temp["num_local_users"] = (s.GetNumLocalUsers())
				temp["num_msd"] = (s.GetNumMsd())
				temp["num_svi_vrf_count"] = (s.GetNumSviVrfCount())
				temp["num_trm_enabled_count"] = (s.GetNumTrmEnabledCount())
				temp["num_upg_users"] = (s.GetNumUpgUsers())
				temp["nxos_image_count"] = (s.GetNxosImageCount())
				temp["object_type"] = (s.GetObjectType())
				temp["outofband_ip"] = (s.GetOutofbandIp())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["record_type"] = (s.GetRecordType())
				temp["record_version"] = (s.GetRecordVersion())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["serial"] = (s.GetSerial())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["site_name"] = (s.GetSiteName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["underlay_peering_active_links_count"] = (s.GetUnderlayPeeringActiveLinksCount())
				temp["upg_job_count"] = (s.GetUpgJobCount())

				temp["upg_status"] = flattenListNiatelemetryJobDetail(s.GetUpgStatus(), d)
				temp["nr_version"] = (s.GetVersion())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				niatelemetryNiaInventoryDcnmResults = append(niatelemetryNiaInventoryDcnmResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryNiaInventoryDcnmResults))
	if err := d.Set("results", niatelemetryNiaInventoryDcnmResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryNiaInventoryDcnmResults[0]["moid"].(string))
	return de
}
