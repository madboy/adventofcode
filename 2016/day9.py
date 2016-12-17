#!/usr/bin/env python3
import sys
from enum import Enum

class State(Enum):
    normal = 1
    marker = 2

def read_marker(marker):
    parts = marker.split("x")
    nbr_of_chars = int(parts[0])
    repetition = int(parts[1])
    return nbr_of_chars, repetition

def decompress(line):
    marker = ""
    state = State.normal
    nbr_of_chars = -1
    to_repeat = ""
    repetition = 0
    output = ""
    for i, c in enumerate(line):
        if nbr_of_chars == 1:
            to_repeat += c
            output += to_repeat*repetition
            marker = ""
            state = State.normal
            nbr_of_chars = -1
            to_repeat = ""
        elif nbr_of_chars > 0:
            to_repeat += c
            nbr_of_chars -= 1
        elif state == State.marker:
            if c == ")":
                state = State.normal
                nbr_of_chars, repetition = read_marker(marker)
            else:
                marker += c
        elif c == "(":
            state = State.marker
        else:
            output += c
    return len(output)

def run():
    for line in sys.stdin:
        line = line.strip()
        dline = decompress(line)
        print("%s... -> %s... [%d]" % (line[:12], dline[:12], len(dline)))

if __name__ == '__main__':
    run()
