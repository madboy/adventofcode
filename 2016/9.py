#!/usr/bin/env python3
import sys

def read_marker(marker):
    parts = marker.split("x")
    nbr_of_chars = int(parts[0])
    repetition = int(parts[1])
    return nbr_of_chars, repetition

def decompress(line):
    marker = ""
    marked = False
    nbr_of_chars = -1
    to_repeat = ""
    repetition = 0
    output = ""
    for i,c in enumerate(line):
        if nbr_of_chars == 1:
            to_repeat += c
            output += to_repeat*repetition
            marker = ""
            marked = False
            nbr_of_chars = -1
            to_repeat = ""
        elif nbr_of_chars > 0:
            to_repeat += c
            nbr_of_chars -= 1
        elif marked:
            if c == ")":
                marked = False
                nbr_of_chars, repetition = read_marker(marker)
            else:
                marker += c
        elif c == "(":
            marked = True
        else:
            output += c
    return output

for line in sys.stdin:
    line = line.strip()
    dline = decompress(line)
    print("%s... -> %s... [%d]" % (line[:12], dline[:12], len(dline)))
