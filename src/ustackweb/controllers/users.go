package controllers

import (
  "ustackweb/utils"
)

type User struct {
  Id int
  Name string
}

type UsersController struct {
  BaseController
}

func (this *UsersController) Prepare() {
  this.PrepareXsrf()
  this.RequireAuth()
  this.PrepareLayout()
  this.Layout = "layouts/default.tpl.html"
}

func (this *UsersController) Get() {
  this.TplNames = "users/index.tpl.html"
  users := []User{User{Id:1,Name:"foo"},
                  User{Id:2,Name:"admin"},
                  User{Id:3,Name:"abc"},
                  User{Id:4,Name:"def"},
                  User{Id:5,Name:"hij"},
                  User{Id:6,Name:"glk"},
                  User{Id:7,Name:"uvw"},
                  User{Id:8,Name:"xyz"}}

  paginator := utils.NewPaginator(this.Ctx.Request, 3, len(users))
  this.Data["paginator"] = paginator
  this.Data["users"] = users[paginator.Offset():min(paginator.Offset()+paginator.PerPageNums, len(users))]
}

func min(x, y int) int {
  if x > y {
    return y
  }
  return x
}
