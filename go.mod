module github.com/Wariie/mod-manager

go 1.14

replace guilhem-mateo.fr/modbase/modules => ../modbase/modules

require (
	github.com/Wariie/go-woxy/com v0.0.0-20200820222252-42021c704ae5
	github.com/Wariie/go-woxy/modbase v0.0.0-20200820225319-9838e983e883
	github.com/Wariie/go-woxy/tools v0.0.0-20200820223903-1d84ed355fa4 // indirect
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/sys v0.0.0-20200820212457-1fb795427249 // indirect
)
