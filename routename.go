package gin

import "maps"

// routeMap storage: {[2]{http.MethodGet,"/base/user/:id"}: "query user by id"}
var routeMap = make(map[[2]string]string)

// setRouteName storage [method, path]
// example: setRouteName(http.MethodGet,"/base/user/:id","query user by id").
func setRouteName(httpMethod, fullPath, name string) {
	routeMap[[2]string{httpMethod, fullPath}] = name
}

// GetRouteName return route name.
func GetRouteName(httpMethod, fullPath string) string {
	return routeMap[[2]string{httpMethod, fullPath}]
}

func GetRouteMap() map[[2]string]string {
	return maps.Clone(routeMap)
}
