package main

import (
	"goweb/framework"
	"strconv"
)

func SubjectAddController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectDelController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	var a string
	if idurl, ok := c.ParamInt("id", 200); ok {
		a = strconv.Itoa(idurl)
	}
	c.SetStatus(200).Json(a)
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectDelController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	var a string
	if idurl, ok := c.ParamInt("id", 200); ok {
		a = strconv.Itoa(idurl)
	}
	c.SetStatus(200).Json(a)
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectDelController")
	return nil
}
