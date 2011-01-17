include $(GOROOT)/src/Make.inc

TARG=goo
GOFMT=gofmt -s -space=true -tabindent=false -tabwidth=4

GOFILES=\
	goo.go\

include $(GOROOT)/src/Make.pkg

format:
	$(GOFMT) -w goo.go
	$(GOFMT) -w goo_test.go
