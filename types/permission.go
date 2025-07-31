package types

type RouteScope string

const (
	RouteScopePublic        RouteScope = "public"
	RouteScopeAuthenticated RouteScope = "authenticated"
	RouteScopeAuthorized    RouteScope = "authorized"
)

type Permission struct {
	Name             string           `json:"name"`
	Domain           string           `json:"domain"`
	Resource         string           `json:"resource"`
	Action           string           `json:"action"`
	Description      string           `json:"description"`
	ImpliedByActions []ActionResource `json:"implied_actions,omitempty"`
}

type ActionResource struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type PermissionPath struct {
	Domain   string     `json:"domain"`
	Path     string     `json:"path"`
	Resource string     `json:"resource"`
	Action   string     `json:"action"`
	Type     RouteScope `json:"type"`
}

type PermissionSeeding struct {
	Permissions []Permission     `json:"permissions"`
	Paths       []PermissionPath `json:"paths"`
}
