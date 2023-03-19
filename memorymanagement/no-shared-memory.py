#!/usr/bin/python3

import os
import sys

data = 1000

pid = os.fork()
if pid < 0:
    print("fail to fork", file=os.stderr)
elif pid == 0:
    data *= 2
    print("child process have data: {}".format(data))
    sys.exit(0)

os.wait()
print("parent process have data: {}".format(data))