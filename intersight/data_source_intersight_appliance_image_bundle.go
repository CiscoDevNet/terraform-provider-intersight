package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplianceImageBundle() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceImageBundleRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"ansible_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"auto_upgrade": {
				Description: "Indicates that the software upgrade was automatically initiated by the Intersight Appliance.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dc_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"debug_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"description": {
				Description: "Short description of the software upgrade bundle.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"endpoint_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"fingerprint": {
				Description: "Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"has_error": {
				Description: "Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"infra_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"init_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the software upgrade bundle.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"notes": {
				Description: "Detailed description of the software upgrade bundle.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Description: "Software upgrade manifest's upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time.\n* `Normal` - Normal upgrade priority is used for all the software upgrades except for the critical security updates. The upgrade service of Intersight Appliance uses the Software Upgrade Policy settings to start the upgrade process.\n* `Critical` - Critical upgrade priority is used for critical updates such as security patches. The upgrade service of the Intersight Appliance starts the upgrade as specified by the upgrade properties in the software manifest file. The upgrade service will not use the settings specified in the Software Upgrade Policy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"release_time": {
				Description: "Software upgrade manifest's release date and time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"status_message": {
				Description: "Status message set during the manifest processing.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"system_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
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
			"ui_packages": {
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
						"file_path": {
							Description: "Optional file path of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_sha": {
							Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"file_size": {
							Description: "Image file size in bytes.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"file_time": {
							Description: "Image file's last modified date and time.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"filename": {
							Description: "Filename of the image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the software image package.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"package_type": {
							Description: "Image package type (e.g. service, system etc.).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "Image package version string.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"upgrade_end_time": {
				Description: "End date of the software upgrade process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"upgrade_grace_period": {
				Description: "Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"upgrade_impact_duration": {
				Description: "Duration (in minutes) for which services will be disrupted.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"upgrade_impact_enum": {
				Description: "UpgradeImpactEnum is used to indicate the kind of impact the upgrade has on currently running services on the appliance.\n* `None` - The upgrade has no effect on the system.\n* `Disruptive` - The services will not be functional during the upgrade.\n* `Disruptive-reboot` - The upgrade needs a reboot.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"upgrade_start_time": {
				Description: "Start date of the software upgrade process.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nr_version": {
				Description: "Software upgrade manifest's version.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceApplianceImageBundleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceImageBundle{}
	if v, ok := d.GetOk("auto_upgrade"); ok {
		x := (v.(bool))
		o.SetAutoUpgrade(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("fingerprint"); ok {
		x := (v.(string))
		o.SetFingerprint(x)
	}
	if v, ok := d.GetOk("has_error"); ok {
		x := (v.(bool))
		o.SetHasError(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("notes"); ok {
		x := (v.(string))
		o.SetNotes(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("priority"); ok {
		x := (v.(string))
		o.SetPriority(x)
	}
	if v, ok := d.GetOk("release_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetReleaseTime(x)
	}
	if v, ok := d.GetOk("status_message"); ok {
		x := (v.(string))
		o.SetStatusMessage(x)
	}
	if v, ok := d.GetOk("upgrade_end_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetUpgradeEndTime(x)
	}
	if v, ok := d.GetOk("upgrade_grace_period"); ok {
		x := int64(v.(int))
		o.SetUpgradeGracePeriod(x)
	}
	if v, ok := d.GetOk("upgrade_impact_duration"); ok {
		x := int64(v.(int))
		o.SetUpgradeImpactDuration(x)
	}
	if v, ok := d.GetOk("upgrade_impact_enum"); ok {
		x := (v.(string))
		o.SetUpgradeImpactEnum(x)
	}
	if v, ok := d.GetOk("upgrade_start_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetUpgradeStartTime(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceImageBundle object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceImageBundleList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching ApplianceImageBundle: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for ApplianceImageBundle list: %s", err.Error())
	}
	var s = &models.ApplianceImageBundleList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to ApplianceImageBundle list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for ApplianceImageBundle data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for ApplianceImageBundle data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.ApplianceImageBundle{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("ansible_packages", flattenListOnpremImagePackage(s.GetAnsiblePackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AnsiblePackages: %s", err.Error())
			}
			if err := d.Set("auto_upgrade", (s.GetAutoUpgrade())); err != nil {
				return diag.Errorf("error occurred while setting property AutoUpgrade: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("dc_packages", flattenListOnpremImagePackage(s.GetDcPackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property DcPackages: %s", err.Error())
			}

			if err := d.Set("debug_packages", flattenListOnpremImagePackage(s.GetDebugPackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property DebugPackages: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}

			if err := d.Set("endpoint_packages", flattenListOnpremImagePackage(s.GetEndpointPackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property EndpointPackages: %s", err.Error())
			}
			if err := d.Set("fingerprint", (s.GetFingerprint())); err != nil {
				return diag.Errorf("error occurred while setting property Fingerprint: %s", err.Error())
			}
			if err := d.Set("has_error", (s.GetHasError())); err != nil {
				return diag.Errorf("error occurred while setting property HasError: %s", err.Error())
			}

			if err := d.Set("infra_packages", flattenListOnpremImagePackage(s.GetInfraPackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InfraPackages: %s", err.Error())
			}

			if err := d.Set("init_packages", flattenListOnpremImagePackage(s.GetInitPackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InitPackages: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("notes", (s.GetNotes())); err != nil {
				return diag.Errorf("error occurred while setting property Notes: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("priority", (s.GetPriority())); err != nil {
				return diag.Errorf("error occurred while setting property Priority: %s", err.Error())
			}

			if err := d.Set("release_time", (s.GetReleaseTime()).String()); err != nil {
				return diag.Errorf("error occurred while setting property ReleaseTime: %s", err.Error())
			}

			if err := d.Set("service_packages", flattenListOnpremImagePackage(s.GetServicePackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ServicePackages: %s", err.Error())
			}
			if err := d.Set("status_message", (s.GetStatusMessage())); err != nil {
				return diag.Errorf("error occurred while setting property StatusMessage: %s", err.Error())
			}

			if err := d.Set("system_packages", flattenListOnpremImagePackage(s.GetSystemPackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SystemPackages: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("ui_packages", flattenListOnpremImagePackage(s.GetUiPackages(), d)); err != nil {
				return diag.Errorf("error occurred while setting property UiPackages: %s", err.Error())
			}

			if err := d.Set("upgrade_end_time", (s.GetUpgradeEndTime()).String()); err != nil {
				return diag.Errorf("error occurred while setting property UpgradeEndTime: %s", err.Error())
			}
			if err := d.Set("upgrade_grace_period", (s.GetUpgradeGracePeriod())); err != nil {
				return diag.Errorf("error occurred while setting property UpgradeGracePeriod: %s", err.Error())
			}
			if err := d.Set("upgrade_impact_duration", (s.GetUpgradeImpactDuration())); err != nil {
				return diag.Errorf("error occurred while setting property UpgradeImpactDuration: %s", err.Error())
			}
			if err := d.Set("upgrade_impact_enum", (s.GetUpgradeImpactEnum())); err != nil {
				return diag.Errorf("error occurred while setting property UpgradeImpactEnum: %s", err.Error())
			}

			if err := d.Set("upgrade_start_time", (s.GetUpgradeStartTime()).String()); err != nil {
				return diag.Errorf("error occurred while setting property UpgradeStartTime: %s", err.Error())
			}
			if err := d.Set("nr_version", (s.GetVersion())); err != nil {
				return diag.Errorf("error occurred while setting property Version: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
