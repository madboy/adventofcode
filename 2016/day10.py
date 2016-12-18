#!/usr/bin/env python3
import sys
from collections import defaultdict

class Bot(object):
    def __init__(self, name, low, high):
        self.name = name
        self.low = low
        self.high = high
        self.compares = (-1, -1)
        self.values = []

    def __str__(self):
        return "[name: %s, low: %s, high: %s, compares=%s]" % \
        (self.name, self.low, self.high, self.compares)

    def __repr__(self):
        return self.__str__()

    def add(self, chip, bots, bins):
        self.values.append(chip)
        if len(self.values) == 2:
            self.values.sort()
            self.compares = (self.values[0], self.values[1])
            low_parts = self.low.split(" ")
            if low_parts[0] == "output":
                bins[low_parts[1]] = self.values[0]
            else:
                bots[low_parts[1]].add(self.values[0], bots, bins)

            high_parts = self.high.split(" ")
            if high_parts[0] == "output":
                bins[high_parts[1]] = self.values[1]
            else:
                bots[high_parts[1]].add(self.values[1], bots, bins)

def run():
    start_settings = []
    instructions = []
    for line in sys.stdin:
        line = line.strip()
        if line.startswith("value"):
            start_settings.append(line)
        elif line.startswith("bot"):
            instructions.append(line)

    bots = defaultdict(Bot)
    bins = defaultdict(int)

    for instruction in instructions:
        parts = instruction.split(" ")
        name = parts[1]
        low = parts[5] + " " + parts[6]
        high = parts[10] + " " + parts[11]
        bots[name] = Bot(name, low, high)

    for setting in start_settings:
        parts = setting.split(" ")
        name = parts[-1]
        chip = int(parts[1])
        bots[name].add(chip, bots, bins)

    for k, bot in bots.items():
        if bot.compares == (17, 61):
            print(bot)

    print("Value: %d" % (bins['0']*bins['1']*bins['2']))

if __name__ == '__main__':
    run()
