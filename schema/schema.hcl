schema "public" {
}
table "webhooks" {
  schema = schema.public
  column "id" {
    type = varchar(255)
  }

  column "name" {
    type = varchar(255)
    null = false
  }

  column "url" {
    unique = true
    type   = varchar(255)
    null   = false
  }

  column "status" {
    type = varchar(255)
    null = false
  }

  column "created_at" {
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }

}

table "endpoints" {
  schema = schema.public
  column "id" {
    type = varchar(255)
  }

  column "webhook_id" {
    type = varchar(255)
    null = false
  }

  column "status" {
    type = varchar(255)
    null = false
  }

  column "url" {
    type = varchar(255)
    null = false
  }

  column "created_at" {
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
  foreign_key "webhook_fk" {
    columns     = [column.webhook_id]
    ref_columns = [table.webhooks.column.id]
  }
}

table "response" {
  schema = schema.public
  column "id" {
    type = varchar(255)
  }

  column "endpoint_id" {
    type = varchar(255)
    null = false
  }

  column "webhook_id" {
    type = varchar(255)
    null = false
  }

  column "body" {
    type = text
  }

  column "result" {
    type = text
  }

  column "created_at" {
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "endpoint_fk" {
    columns     = [column.endpoint_id]
    ref_columns = [table.endpoints.column.id]
  }
}