package core

import "ginEssential/initialize"

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")


	Router.Run(":3000")
}
