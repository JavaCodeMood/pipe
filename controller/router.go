// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package controller

import (
	"github.com/b3log/solo.go/controller/console"
	"github.com/b3log/solo.go/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes() *gin.Engine {
	ret := gin.New()
	// TODO: D, ret.Use(favicon.New("./favicon.ico"))
	ret.Use(gin.Recovery())

	store := sessions.NewCookieStore([]byte(util.Conf.SessionSecret))
	cookiePath := "/"
	if "" != util.Conf.Context {
		cookiePath = util.Conf.Context
	}
	store.Options(sessions.Options{
		Path:     cookiePath,
		MaxAge:   util.Conf.SessionMaxAge,
		Secure:   true,
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("solo.go", store))

	ret.Any("/hp/*apis", util.HacPaiAPI())

	ret.POST("/init", initCtl)

	ret.POST("/login", loginCtl)
	ret.POST("/logout", logoutCtl)

	statusGroup := ret.Group("/status")
	statusGroup.GET("", GetPlatformStatusCtl)
	statusGroup.GET("/ping", pingCtl)

	consoleGroup := ret.Group("/console")
	consoleGroup.POST("/articles", console.AddArticleCtl)
	consoleGroup.GET("/articles", console.GetArticlesCtl)
	consoleGroup.GET("/articles/:id", console.GetArticleCtl)
	consoleGroup.DELETE("/articles/:id", console.RemoveArticleCtl)
	consoleGroup.PUT("/articles/:id", console.UpdateArticleCtl)
	consoleGroup.GET("/tags", console.GetTagsCtl)

	return ret
}
