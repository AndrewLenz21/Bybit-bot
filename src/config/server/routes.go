package server

import test "bybitbot/src/controllers/test"

//IN THIS FILE WE ARE GOING TO CALL THE CONTROLLERS
func StartRoutes() {
	test.TestController(server)
}
