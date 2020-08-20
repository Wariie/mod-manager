module github.com/Wariie/mod-manager

go 1.14

replace guilhem-mateo.fr/modbase/modules => ../modbase/modules

require (
	github.com/Wariie/go-woxy/com v0.0.0-20200820220311-e31b3dd08379
	github.com/Wariie/go-woxy/modbase v0.0.0-20200820222252-42021c704ae5
	github.com/Wariie/go-woxy/tools v0.0.0-20200820221546-cf6d3e445bfe // indirect
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/sys v0.0.0-20200820212457-1fb795427249 // indirect
)
