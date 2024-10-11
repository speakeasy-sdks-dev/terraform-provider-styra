resource "styra_system" "my_system" {
  bundle_download = {
    delta_bundles = true
  }
  bundle_registry = {
    disable_bundle_compatibility_check = false
    distribution_s3 = {
      access_keys    = "...my_access_keys..."
      bucket         = "...my_bucket..."
      context_path   = "...my_context_path..."
      discovery_path = "...my_discovery_path..."
      endpoint       = "...my_endpoint..."
      opa_credentials = {
        environment_credentials = {
          # ...
        }
        metadata_credentials = {
          aws_region = "...my_aws_region..."
          iam_role   = "...my_iam_role..."
        }
        web_identity_credentials = {
          aws_region   = "...my_aws_region..."
          session_name = "...my_session_name..."
        }
      }
      policy_path = "...my_policy_path..."
      region      = "...my_region..."
      role_arn    = "...my_role_arn..."
    }
    entrypoints = [
      "..."
    ]
    manual_deployment = false
    # ...        max_bundles = 10
    max_deployed_bundles = 4
    optimization_level   = 2
  }
  context_bundle_data_only = true
  context_bundle_roots = [
    "..."
  ]
  decision_mappings = {
    allowed = {
      expected = "{ \"see\": \"documentation\" }"
      negated  = false
      path     = "...my_path..."
    }
    columns = [
      {
        key  = "...my_key..."
        path = "...my_path..."
        type = "...my_type..."
      }
    ]
    reason = {
      path = "...my_path..."
    }
  }
  deployment_parameters = {
    deny_on_opa_fail = true
    discovery = {
      # ...
    }
    extra = {
      # ...
    }
    http_proxy            = "...my_http_proxy..."
    https_proxy           = "...my_https_proxy..."
    kubernetes_version    = "...my_kubernetes_version..."
    mutating_webhook_name = "...my_mutating_webhook_name..."
    namespace             = "...my_namespace..."
    no_proxy              = "...my_no_proxy..."
    timeout_seconds       = 4
    trusted_ca_certs = [
      "..."
    ]
    trusted_container_registry = "...my_trusted_container_registry..."
  }
  description   = "...my_description..."
  error_setting = "...my_error_setting..."
  external_bundles = {
    bundles = {
      persist = true
      polling = {
        long_polling_timeout_seconds = 1
        max_delay_seconds            = 3
        min_delay_seconds            = 10
      }
      resource = "...my_resource..."
      service  = "...my_service..."
      signing = {
        exclude_files = [
          "..."
        ]
        keyid = "...my_keyid..."
        public_keys = {
          algorithm   = "...my_algorithm..."
          key         = "...my_key..."
          private_key = "...my_private_key..."
          scope       = "...my_scope..."
        }
        scope = "...my_scope..."
      }
      size_limit_bytes = 1
    }
    services = [
      {
        allow_insecure_tls = false
        credentials = {
          azure_managed_identity = {
            api_version = "...my_api_version..."
            client_id   = "...my_client_id..."
            endpoint    = "...my_endpoint..."
            mi_res_id   = "...my_mi_res_id..."
            object_id   = "...my_object_id..."
            resource    = "...my_resource..."
          }
          bearer = {
            scheme     = "...my_scheme..."
            token      = "...my_token..."
            token_path = "...my_token_path..."
          }
          client_tls = {
            cert                   = "...my_cert..."
            private_key            = "...my_private_key..."
            private_key_passphrase = "...my_private_key_passphrase..."
          }
          gcp_metadata = {
            access_token_path = "...my_access_token_path..."
            audience          = "...my_audience..."
            endpoint          = "...my_endpoint..."
            id_token_path     = "...my_id_token_path..."
            scopes = [
              "..."
            ]
          }
          oauth2 = {
            additional_claims = {
              # ...
            }
            additional_headers = {
              "see" : "documentation",
            }
            additional_parameters = {
              "see" : "documentation",
            }
            client_id         = "...my_client_id..."
            client_secret     = "...my_client_secret..."
            grant_type        = "...my_grant_type..."
            include_jti_claim = true
            scopes = [
              "..."
            ]
            signing_key = "...my_signing_key..."
            thumbprint  = "...my_thumbprint..."
            token_url   = "...my_token_url..."
          }
          plugin = "...my_plugin..."
          s3_signing = {
            environment_credentials = "{ \"see\": \"documentation\" }"
            metadata_credentials = {
              aws_region = "...my_aws_region..."
              iam_role   = "...my_iam_role..."
            }
            profile_credentials = {
              aws_region = "...my_aws_region..."
              path       = "...my_path..."
              profile    = "...my_profile..."
            }
            service = "...my_service..."
            web_identity_credentials = {
              aws_region   = "...my_aws_region..."
              session_name = "...my_session_name..."
            }
          }
        }
        headers = {
          "see" : "documentation",
        }
        keys = {
          algorithm   = "...my_algorithm..."
          key         = "...my_key..."
          private_key = "...my_private_key..."
          scope       = "...my_scope..."
        }
        name                            = "...my_name..."
        response_header_timeout_seconds = 0
        tls = {
          ca_cert            = "...my_ca_cert..."
          system_ca_required = true
        }
        type = "...my_type..."
        url  = "...my_url..."
      }
    ]
  }
  external_id      = "...my_external_id..."
  filter_stacks    = false
  kafka_topic      = "...my_kafka_topic..."
  mock_opa_enabled = true
  name             = "...my_name..."
  read_only        = false
  recursive        = "...my_recursive..."
  source_control = {
    origin = {
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
  }
  type = "...my_type..."
  type_parameters = {
    # ...
  }
}