data "intersight_search_search_item" "odata_top_skip_value1"{
	odata {
	 top = 143
	 skip = 200
	}
}

data "intersight_aaa_audit_record" "odata_top_skip_small_value1"{
	odata {
	 top = 80
	 skip = 40
	}
}

data "intersight_workflow_task_info" "odata_top_skip_small_value2"{
	odata {
	 top = 30
	 skip = 60
	}
}

data "intersight_search_search_item" "odata_top_and_skip_small_value3"{
	odata {
	 top = 20
	 skip = 10
	}
}

data "intersight_workflow_task_info" "odata_top_skip_value2"{
	odata {
	 top = 230
	 skip = 150
	}
}

data "intersight_aaa_audit_record" "odata_top_and_skip_value3"{
	odata {
	 top = 70
	 skip = 65
	}
}

data "intersight_workflow_task_info" "odata_top_and_skip_value4"{
	odata {
	 top = 174
	 skip = 133
	}
}

data "intersight_search_search_item" "odata_top_and_skip_value5"{
	odata {
	 top = 440
	 skip = 250
	}
}

data "intersight_aaa_audit_record" "odata_top_and_skip_value6"{
	odata {
	 top = 748
	 skip = 660
	}
}

data "intersight_aaa_audit_record" "odata_top_and_skip_large_values1"{
	odata {
	 top = 1000
	 skip = 500
	}
}

data "intersight_aaa_audit_record" "odata_top_value1"{
	odata {
	 top = 12
	}
}

data "intersight_appliance_app_status" "odata_skip_value1"{
	odata {
	 skip = 60
	}
}

data "intersight_appliance_app_status" "odata_top_value2"{
	odata {
	 top = 80
	}
}

data "intersight_search_search_item" "odata_top_value3"{
	odata {
	 top = 890
	}
}

data "intersight_workflow_task_metadata" "odata_skip_large_value1"{
	odata {
	 skip = 200
	}
}

data "intersight_workflow_task_definition" "odata_skip_large_value2"{
	odata {
		skip = 1200
	}
}

data "intersight_appliance_app_status" "odata_top_zero_value"{
	odata {
	 top = 0
	}
}

data "intersight_appliance_app_status" "odata_skip_zero_value"{
	odata {
	 skip = 0
	}
}

data "intersight_workflow_workflow_info" "odata_top_negative_value"{
	odata {
	 top = -20
	}
}

data "intersight_aaa_audit_record" "odata_top_and_skip_value8"{
	odata {
	 skip = 500
	 top = 700
	}
}

data "intersight_aaa_audit_record" "odata_top_and_skip_value9"{
	odata {
	 top = 1300
	 skip = 500
	}
}

data "intersight_workflow_task_definition" "odata_count_true_value"{
	odata {
	 nr_count = true
	}
}

data "intersight_workflow_task_metadata" "odata_count_false_value"{
	odata {
	 nr_count = false
	}
}

data "intersight_workflow_task_metadata" "odata_count_filter"{
	odata {
	 nr_count = true
	 filter = "Label eq 'Restart IMC'"
	}
}

data "intersight_iam_role" "odata_count_orderby"{
	odata {
	 nr_count = false
	 orderby = "CreateTime"
	}
}

data "intersight_hyperflex_software_distribution_component" "odata_orderby_descending"{
	odata {
	 orderby = "CreateTime desc"
	}
}

data "intersight_adapter_ext_eth_interface" "odata_orderby_ascending_select"{
	odata {
	 orderby = "CreateTime"
	 nr_select = "CreateTime,MacAddress"
	}
}

data "intersight_adapter_ext_eth_interface" "odata_orderby_ascending"{
	odata {
	 orderby = "Moid"
	}
}

data "intersight_appliance_app_status" "odata_orderby_top_skip"{
	odata {
	 orderby = "ModTime desc"
	 top = 10
	 skip = 50
	}
}

data "intersight_compute_rack_unit" "odata_filter_equal"{
	odata {
	 filter = "Name eq 'C125-WZP22360YWC'"
	}
}

data "intersight_adapter_host_eth_interface" "odata_filter_not_equal"{
	odata {
	 filter = "Moid ne '6167ee1576752d32341e2df3'"
	}
}

data "intersight_adapter_host_eth_interface" "odata_filter_orderby_select"{
	odata {
	filter = "Moid ne '6167ee1576752d32341e2df3'"
	orderby = "Name"
	nr_select = "Name,ModTime"
	}
}

data "intersight_compute_rack_unit" "odata_filter_greater_than"{
	odata {
	filter = "AvailableMemory gt 98304"
	}
}

data "intersight_compute_rack_unit" "odata_filter_greater_than_or_equal"{
	odata {
	filter = "AvailableMemory ge 98304"
	}
}

data "intersight_compute_rack_unit" "odata_filter_lesser_than"{
	odata {
	filter = "AvailableMemory lt 98304"
	}
}

data "intersight_search_search_item" "odata_filter_less_than_or_equal"{
	odata {
	filter = "MemorySize le 50"
	}
}

data "intersight_compute_rack_unit" "odata_filter_and_additional_property"{
	odata {
		filter = "AvailableMemory gt 65000"
	}
	additional_properties = jsonencode({ "Model" = "UCSC-C240-M5L"})
}

data "intersight_compute_rack_unit" "odata_filter_not_and_top_and_select"{
	odata {
		nr_select = "Model,Serial" 
		top=10 
		filter = "not(Model eq 'UCSC-C125' or Model eq 'HX220C-M5S')"
	}
}

data "intersight_compute_rack_unit" "odata_filter_in"{
	odata {
		filter = "Model in ('UCSC-C125', 'UCSC-C240-M5L')"
	}
}

data "intersight_compute_rack_unit" "odata_filter_contains"{
	odata {
		filter = "contains(Model,'C240')"
	}
}

data "intersight_server_profile" "odata_filter_startswith"{
	odata {
		filter = "startswith(Name,'cloned')"
	}
}

data "intersight_server_profile" "odata_filter_endswith"{
	odata {
		filter = "endswith(Name,'test')"
	}
}

data "intersight_appliance_app_status" "odata_top_filter_empty_array"{
	odata {
		filter = "not(Tags/any())"
		top = 5
	}
}

data "intersight_workflow_workflow_info" "odata_top_filter_not_empty_array"{
	odata {
		filter = "Message/any()"
	}
}

data "intersight_compute_rack_unit" "odata_top_skip_select"{
	odata {
		top = 20 
		skip = 1 
		nr_select = "Model,Serial"
	}
}

data "intersight_adapter_host_fc_interface" "odata_top_inlinecount_allpages"{
	odata {
		top = 2 
		inlinecount = "allpages"
	}
}

data "intersight_adapter_host_fc_interface" "odata_top_inlinecount_none"{
	odata {
		top = 3
		inlinecount = "none"
	}
}
