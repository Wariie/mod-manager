module github.com/Wariie/mod-manager

go 1.15

replace guilhem-mateo.fr/modbase/modules => ../modbase/modules

require (
	github.com/Wariie/go-woxy/com v0.0.0-20200909143036-4e7d0ed96e98
	github.com/Wariie/go-woxy/modbase v0.0.0-20200909182143-8406c03934c3
	github.com/Wariie/go-woxy/tools v0.0.0-20200909143036-4e7d0ed96e98 // indirect
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
)
