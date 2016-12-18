#!/usr/bin/env python3
import sys

def read_marker(marker):
    parts = marker.split("x")
    nbr_of_chars = int(parts[0])
    repetition = int(parts[1])
    return nbr_of_chars, repetition

def _read_marker(s, l):
    #read until we hit first )
    marker = ""
    c = l[s]
    while c != ")":
        marker += c
        s += 1
        c = l[s]
    noc, rep = read_marker(marker)
    s += 1 # remove the )
    return s, noc, rep

def decompress(start, end, line):
    dlen = 0
    while start < end:
        current = line[start]
        if current == "(":
            start, noc, rep = _read_marker(start+1, line)
            dlen += rep * decompress(start, noc+start, line)
            start = noc + start
        else:
            dlen += 1
            start += 1
    return dlen

def run():
    for line in sys.stdin:
        line = line.strip()
        dlen = decompress(0, len(line), line)
        print("%s... -> ... [%d]" % (line[:12], dlen))

if __name__ == '__main__':
    run()
