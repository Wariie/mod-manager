module github.com/Wariie/mod-manager

go 1.15

replace guilhem-mateo.fr/modbase/modules => ../modbase/modules

require (
	github.com/Wariie/go-woxy/com v0.0.0-20200909115455-2200e097047c
	github.com/Wariie/go-woxy/modbase v0.0.0-20200909122107-588ccb30e3b7
	github.com/Wariie/go-woxy/tools v0.0.0-20200909115455-2200e097047c // indirect
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
)
