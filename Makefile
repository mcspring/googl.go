include $(GOROOT)/src/Make.inc

TARG=googl
GOFMT=gofmt -s -space=true -tabindent=false -tabwidth=4

GOFILES=\
	googl.go\

include $(GOROOT)/src/Make.pkg

format:
	$(GOFMT) -w googl.go
	$(GOFMT) -w googl_test.go
