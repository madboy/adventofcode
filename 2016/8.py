#!/usr/bin/env python3
import sys
from collections import deque
import numpy as np

COLUMNS = 50
ROWS = 6

def turn_on_pixels(columns, rows, pixels):
    for y in range(rows):
        for x in range(columns):
            pixels[y][x] = "#"

def rotate_column(column, by, pixels):
    items = deque(pixels.transpose()[column])
    items.rotate(by)
    pixels.transpose()[column] = items

def rotate_row(row, by, pixels):
    items = deque(pixels[row])
    items.rotate(by)
    pixels[row] = items

def print_screen(pixels):
    for row in pixels:
        leds = ""
        for col in row:
            leds += col
        print(leds)

def count_on_pixels(pixels):
    count = 0
    for row in pixels:
        for col in row:
            if col == "#":
                count += 1
    return count

PIXELS = [["."]*COLUMNS for i in range(ROWS)]
PIXELS = np.array(PIXELS)

for line in sys.stdin:
    line.strip()
    parts = line.split(" ") # rect 2x1 || rotate row y=0 by 5
    if line.startswith("rect"):
        coords = parts[1].split("x")
        turn_on_pixels(int(coords[0]), int(coords[1]), PIXELS)
        continue
    # ["rotate", "row", "y=0", "by", "5"]
    rc = int(parts[2].split("=")[1]) # y=11
    rotate_by = int(parts[-1])
    if line.startswith("rotate row"):
        rotate_row(rc, rotate_by, PIXELS)
    elif line.startswith("rotate column"): # rotate column
        rotate_column(rc, rotate_by, PIXELS)

print_screen(PIXELS)
print(count_on_pixels(PIXELS))
