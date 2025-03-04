env "local" {
  dev = "docker://postgres/16"
 
  migration {
    dir = "file://migrations"
    format = "atlas"

  }

  schema {
    src = "internal/idl/db/schema"
  }
}