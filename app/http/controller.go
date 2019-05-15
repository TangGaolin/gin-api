package http

func ping(c *Context) {
	if err := svc.Ping(); err != nil {
		c.error(500, err.Error())
	}
	c.success(nil)
}

func login(c *Context) {

	userName := c.Query("user_name")
	password := c.Query("password")
	result, err := svc.Login(userName, password)
	if err != nil {
		c.error(403, err.Error())
		return
	}
	c.success(result)
	return
}

func authInfo(c *Context) {
	userName := c.Query("user_name")
	result, err := svc.AuthInfo(userName)
	if err != nil {
		c.error(404, err.Error())
		return
	}
	c.success(result)
	return
}
