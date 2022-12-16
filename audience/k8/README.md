We need the following env variables

k set env deploy/kong-kong KONG_PLUGINSERVER_NAMES-
k set env deploy/kong-kong 

KONG_PLUGINS=bundled,go-jwt-aud
KONG_PLUGINSERVER_NAMES=go-jwt-aud
KONG_PLUGINSERVER_GO_JWT_AUD_START_CMD=/opt/kong/plugins/go-jwt-aud
KONG_PLUGINSERVER_GO_JWT_AUD_QUERY_CMD=/opt/kong/plugins/go-jwt-aud -dump


k set env deploy/kong-kong KONG_PLUGINS="bundled,go-jwt-aud" KONG_PLUGINSERVER_NAMES="go-jwt-aud" KONG_PLUGINSERVER_GO_JWT_AUD_START_CMD=/opt/kong/plugins/go-jwt-aud KONG_PLUGINSERVER_GO_JWT_AUD_QUERY_CMD="/opt/kong/plugins/go-jwt-aud -dump"

k set env deploy/kong-kong KONG_PLUGINS- KONG_PLUGINSERVER_NAMES- KONG_PLUGINSERVER_GO_JWT_AUD_QUERY_CMD-