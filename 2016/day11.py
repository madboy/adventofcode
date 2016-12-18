#!/usr/bin/env python3
import sys


class Floor(object):
    def __init__(self):
        self.microchips = []
        self.generators = []

    def __repr__(self):
        return "chips: %s, generators: %s" % (str(self.microchips), str(self.generators))

    def set_contents(self, c):
        parts = c.split(" a ")
        for p in parts:
            if "microchip" in p:
                self.microchips.append(p.split(" ")[0])
            elif "generator" in p:
                self.generators.append((p.split(" ")[0]))
        pass


class Elevator(object):
    def __init__(self, floor):
        self.floor = floor
        self.contents = []

    def __repr__(self):
        return "[%s]: %s" % (self.floor, str(self.contents))

    @staticmethod
    def allowed(item1, item2):
        if chip(item1) and chip(item2):
            return True
        return pair(item1, item2)


def chip(item):
    return "compatible" in item


def pair(item1, item2):
    return item1.startswith(item2[:5])


def run():
    floors = {}
    for line in sys.stdin:
        line = line.strip()
        floor, contents = line.split("contains")
        nbr = floor.split(" ")[1]
        floors[nbr] = Floor()
        floors[nbr].set_contents(contents)
    elevator = Elevator("first")
    print(elevator)

    print(floors)

if __name__ == '__main__':
    run()
