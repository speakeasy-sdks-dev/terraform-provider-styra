resource "styra_library" "my_library" {
  description = "...my_description..."
  id          = "fa7ed3c3-aa52-4d40-b844-4511eaa69789"
  read_only   = false
  source_control = {
    library_origin = {
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
}