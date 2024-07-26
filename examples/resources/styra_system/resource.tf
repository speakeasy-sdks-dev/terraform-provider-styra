resource "styra_system" "my_system" {
  context_bundle_data_only = true
  description              = "...my_description..."
  error_setting            = "...my_error_setting..."
  external_id              = "...my_external_id..."
  filter_stacks            = false
  kafka_topic              = "...my_kafka_topic..."
  mock_opa_enabled         = true
  name                     = "Ms. Tim Kuvalis"
  read_only                = false
  recursive                = "...my_recursive..."
  type                     = "...my_type..."
}