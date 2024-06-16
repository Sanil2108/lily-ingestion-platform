-- kong-plugin-authentication-1.0-1.rockspec

package = "kong-plugin-authentication"
version = "1.0-1"
source = {
   url = "file://."
}

description = {
   summary = "A custom plugin for Kong",
   detailed = [[
      This plugin does X, Y, and Z for Kong.
   ]],
   homepage = "https://yourpluginhomepage.com",
   license = "MIT"
}

dependencies = {
   "lua >= 5.1"
}

build = {
   type = "builtin",
   modules = {
      ["kong.plugins.kong-plugin-authentication.handler"] = "handler.lua",
      ["kong.plugins.kong-plugin-authentication.schema"] = "schema.lua"
   }
}
