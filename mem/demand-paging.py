#!/usr/bin/python3

import mmap
import time
import datetime

ALLOC_SIZE = 100 * 1024 * 1024
ACCESS_UNIT = 10 * 1024 * 1024
PAGE_SIZE = 4096

def show_message(msg):
    print("{}: {}".format(datetime.datetime.now().strftime("%H:%M:%S"), msg))

show_message("before alloc memory. if you click enter, this program will alloc memory")
input()

memregion = mmap.mmap(-1, ALLOC_SIZE, flags=mmap.MAP_PRIVATE)
show_message("this program get new memory segment")
input()

for i in range(0, ALLOC_SIZE, PAGE_SIZE):
    memregion[i] = 0
    if i%ACCESS_UNIT == 0 and i != 0:
        show_message("{} MiB access happened".format(i//(1024*1024)))
        time.sleep(1)
show_message("all memory access is done.if you click enter, this program will finish")
input()