package main

import "goweb/framework"

func UserLoginController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, UserLoginController")
	return nil
}
