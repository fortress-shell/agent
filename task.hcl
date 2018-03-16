job "task" {
  datacenters = ["dc1"]
  type = "batch"

  parameterized {
    payload       = "required"
    meta_required = ["config_path", "build_id", "repo", "branch", "commit", "username", "ssh_key"]
  }

  group "tasks" {
    restart {
      attempts = 0
      mode = "fail"
    }
    task "build" {
     driver = "raw_exec"
      env {
        "CONFIG_PATH"      = "${NOMAD_TASK_DIR}/payload.yml"
      }
     config {
        command = "fortress"
     }
     dispatch_payload {
        file = "payload.yml"
      }
    }
  }
}
