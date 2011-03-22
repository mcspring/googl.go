#!/usr/bin/env python

import sys
import webbrowser

if 1 >= len(sys.argv):
	sys.exit(1)

req_url = sys.argv[1]
try:
	webbrowser.open_new(req_url)

	sys.exit(0)
except Exception, err:
	sys.exit(1)

