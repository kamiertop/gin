package gin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setRouteName(t *testing.T) {
	setRouteName(http.MethodGet, "/base/user/:id", "query user by id")
	assert.Equal(t, "query user by id", routeMap[[2]string{http.MethodGet, "/base/user/:id"}])
}

func TestGetRouteName(t *testing.T) {
	engine := New()
	engine.GETEX("user/:id", "get user by id", func(c *Context) {})
	assert.Equal(t, "get user by id", GetRouteName(http.MethodGet, "/user/:id"))

	engine.Group("user").DELETEEX(":id", "delete user by id", func(c *Context) {})
	assert.Equal(t, "delete user by id", GetRouteName(http.MethodDelete, "/user/:id"))

}
