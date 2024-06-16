local redis = require "resty.redis"
local cjson = require "cjson"

local APIKeyAuthenticator = {
  VERSION  = "1.0.0",
  PRIORITY = 1000,
}

local red -- Redis connection object

function APIKeyAuthenticator:new()
  APIKeyAuthenticator.super.new(self, "api-key-authenticator")
end

function APIKeyAuthenticator:init_worker()
  APIKeyAuthenticator.super.init_worker(self)

  -- Create a single Redis connection per worker process
  red = redis:new()
  red:set_timeout(1000)  -- 1 second

  local ok, err = red:connect(kong.configuration.redis_host, kong.configuration.redis_port)
  if not ok then
    kong.log.err("Failed to connect to Redis: ", err)
  end

  if kong.configuration.redis_password and kong.configuration.redis_password ~= "" then
    local res, err = red:auth(kong.configuration.redis_password)
    if not res then
      kong.log.err("Failed to authenticate with Redis: ", err)
    end
  end
end

function APIKeyAuthenticator:access(conf)
  APIKeyAuthenticator.super.access(self)

  -- Get the API key from the request headers
  local api_key = kong.request.get_header("api-key")
  if not api_key then
    return kong.response.exit(401, { message = "API key is missing" })
  end

  -- Fetch the user information from Redis using the API key
  local res, err = red:get(api_key)
  if not res or res == ngx.null then
    return kong.response.exit(403, { message = "Invalid API key" })
  elseif err then
    kong.log.err("Failed to get API key from Redis: ", err)
    return kong.response.exit(500, { message = "Internal Server Error" })
  end

  -- Decode the JSON response
  local user_info, err = cjson.decode(res)
  if not user_info then
    kong.log.err("Failed to decode user info: ", err)
    return kong.response.exit(500, { message = "Internal Server Error" })
  end

  -- Set the tenantId and userId in the headers for downstream services
  kong.service.request.set_header("X-Tenant-ID", user_info.tenantId)
  kong.service.request.set_header("X-User-ID", user_info.userId)
end

return APIKeyAuthenticator
