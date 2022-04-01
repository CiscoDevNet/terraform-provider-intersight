---
name: Issue report
about: Create an issue report to help us improve
title: "[ISSUE] Description"
assignees: ''

---

#### Bug Report Checklist

- [ ] Have you provided a full/minimal configuration to reproduce the issue?
- [ ] Have you [tested with the latest master] to confirm the issue still exists?
- [ ] Have you provided the terraform console logs with environment variable set to TF_LOG=trace?

<!--
Please follow the issue template below for issue reports.
-->

##### Description

<!-- Clear and concise descrption of what is the question, suggestion or issue and why this is a problem for you. -->

##### Terraform-provider-intersight version

<!-- which version of terraform-provider-intersight are you using? -->

##### Configuration file

<!-- if it is a bug, .tf file to reproduce it
If you post the code inline, please wrap it with
``` hcl
(your code here)
```
  -->

##### Actual output (Attach screenshots if applicable)

<!-- 
    Post the actual output from the execution.
    Run the terraform commands after setting the follow environment variable for detail logging.
    export TF_LOG=trace
-->

##### Related issues/PRs

<!-- has a similar issue/PR been reported/opened before? Please do a search in https://github.com/CiscoDevNet/terraform-provider-intersight/issues -->

##### Suggest a fix

<!-- You can point to what might be causing the problem (line of code or commit), or simply make a suggestion -->
