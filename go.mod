module github.com/Wariie/mod-manager

go 1.14

replace guilhem-mateo.fr/modbase/modules => ../modbase/modules

require (
	github.com/Wariie/go-woxy/com v0.0.0-20200820153627-1c4025862a85
	github.com/Wariie/go-woxy/modbase v0.0.0-20200820154249-86ec98197822
	github.com/Wariie/go-woxy/tools v0.0.0-20200820153550-bcfb8f16f012 // indirect
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/sys v0.0.0-20200819171115-d785dc25833f // indirect
)
