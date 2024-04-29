package stringer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToSnakeCase(t *testing.T) {
	require.Equal(t, "id", ToSnakeCase("ID"))
	require.Equal(t, "iso3", ToCamelCase("iso3"))
	require.Equal(t, "ticket", ToSnakeCase("Ticket"))
	require.Equal(t, "ticket_id", ToSnakeCase("TicketID"))
	require.Equal(t, "ticket_id_name", ToSnakeCase("TicketIDName"))
	require.Equal(t, "ticket_id", ToSnakeCase("TicketId"))
	require.Equal(t, "api_endpoint", ToSnakeCase("APIEndpoint"))
	require.Equal(t, "api_endpoint", ToSnakeCase("Api   Endpoint"))
	require.Equal(t, "api_endpoint", ToSnakeCase("API   endpoint"))
	require.Equal(t, "api_endpoint_id", ToSnakeCase("apiEndpointID"))
	require.Equal(t, "api_endpoint_id", ToSnakeCase("api_Endpoint_id"))
}

func TestToCamelCase(t *testing.T) {
	require.Equal(t, "id", ToCamelCase("ID"))
	require.Equal(t, "iso3", ToCamelCase("ISO3"))
	require.Equal(t, "ticket", ToCamelCase("Ticket"))
	require.Equal(t, "ticketId", ToCamelCase("TicketID"))
	require.Equal(t, "ticketIdName", ToCamelCase("TicketIDName"))
	require.Equal(t, "ticketId", ToCamelCase("TicketId"))
	require.Equal(t, "apiEndpoint", ToCamelCase("APIEndpoint"))
	require.Equal(t, "apiEndpoint", ToCamelCase("Api    Endpoint"))
	require.Equal(t, "apiEndpoint", ToCamelCase("Api    endpoint"))
	require.Equal(t, "apiEndpointId", ToCamelCase("apiEndpointID"))
	require.Equal(t, "apiEndpointId", ToCamelCase("api_endpoint_id"))
	require.Equal(t, "apiEndpointId", ToCamelCase("api_Endpoint_id"))
}
