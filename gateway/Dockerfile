FROM kong:3.1.0

USER root

COPY plugins /plugins

RUN luarocks install /plugins/kong-plugin-authentication-1.0-1.all.rock