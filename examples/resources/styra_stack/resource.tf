resource "styra_stack" "my_stack" {
  description = "...my_description..."
  name        = "...my_name..."
  read_only   = true
  source_control = {
    stack_origin = {
      commit      = "...my_commit..."
      credentials = "...my_credentials..."
      path        = "...my_path..."
      reference   = "...my_reference..."
      ssh_credentials = {
        passphrase  = "...my_passphrase..."
        private_key = "...my_private_key..."
      }
      url = "...my_url..."
    }
    use_workspace_settings = true
  }
  type = "...my_type..."
  type_parameters = {
    # ...
  }
}