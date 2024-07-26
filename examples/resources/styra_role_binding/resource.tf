resource "styra_role_binding" "my_rolebinding" {
  id = "6f629e17-7630-4a51-b91e-14e0b36dbabc"
  resource_filter = {
    id   = "cb1df49f-f91d-4576-bc1d-2a13ef306a5e"
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
      id   = "6b163d86-8ce6-498f-aef2-603ebc9c11f9"
      kind = "...my_kind..."
    },
  ]
}