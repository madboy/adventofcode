#!/usr/bin/env python3
import sys

COLUMNS = 50
ROWS = 6

def prepare_screen(pixels):
    for y in range(ROWS):
        for x in range(COLUMNS):
            pixels[(x,y)] = "."

def turn_on_pixels(columns, rows, pixels):
    count = 0
    for y in range(rows):
        for x in range(columns):
            pixels[(x,y)] = "#"
            count += 1
    return count

def rotate_row(row, by, pixels):
    changes = []
    for x in range(COLUMNS):
        coord = (x, row)
        if pixels[coord] == "#":
            pixels[coord] = "."
            changes.append(((x+by) % COLUMNS, row))
    for change in changes:
        pixels[change] = "#"

def rotate_column(column, by, pixels):
    changes = []
    for y in range(ROWS):
        coord = (column, y)
        if pixels[coord] == "#":
            pixels[coord] = "."
            changes.append((column, (y+by) % ROWS))
    for change in changes:
        pixels[change] = "#"

def print_screen(pixels):
    for y in range(ROWS):
        row = ""
        for x in range(COLUMNS):
            row += pixels[(x,y)]
        print(row)

def count_on_pixels(pixels):
    count = 0
    for y in range(ROWS):
        for x in range(COLUMNS):
            if pixels[(x,y)] == "#":
                count += 1
    return count

pixels = {}
prepare_screen(pixels)
scount = 0
for line in sys.stdin:
    # deal with three different actions
    # rect (turn on all the specified pixels, starting top left 0,0)
    # rotate row
    # rotate column
    parts = line.split(" ")
    if line.startswith("rect"):
        coords = parts[1].split("x")
        scount += turn_on_pixels(int(coords[0]), int(coords[1]), pixels)
    elif line.startswith("rotate row"):
        row = int(parts[2][-1])
        rotate_by = int(parts[-1])
        rotate_row(row, rotate_by, pixels)
    else: # rotate column
        column = int(parts[2][-1])
        rotate_by = int(parts[-1])
        rotate_column(column, rotate_by, pixels)

print_screen(pixels)
print(count_on_pixels(pixels), scount)
