return {
  no_consumer = true,
  fields = {
    redis_host = { type = "string", default = "127.0.0.1" },
    redis_port = { type = "number", default = 6379 },
    redis_password = { type = "string", default = "" }
  }
}