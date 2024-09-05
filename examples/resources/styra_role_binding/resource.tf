resource "styra_role_binding" "my_rolebinding" {
  id = "...my_id..."
  resource_filter = {
    id   = "...my_id..."
    kind = "...my_kind..."
  }
  role_id = "...my_role_id..."
  subjects = [
    {
      claim_config = {
        identity_provider = "...my_identity_provider..."
        key               = "...my_key..."
        value             = "...my_value..."
      }
      id   = "...my_id..."
      kind = "...my_kind..."
    }
  ]
}