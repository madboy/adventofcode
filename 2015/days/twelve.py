#!/usr/bin/env python
import fileinput
import json

def count(numbers, ignore):
    if type(numbers) == dict:
        if ignore and "red" in numbers.values():
            return 0
        else:
            return count(numbers.values(), ignore)
    elif type(numbers) == list:
        inner_sum = 0
        for e in numbers:
            inner_sum += count(e, ignore)
        return inner_sum
    elif type(numbers) == int:
        return numbers
    return 0

def bookkeeping(ignore):
    total = 0
    for i in j:
        total += count(i, ignore)
    return total

s = ""
for line in fileinput.input():
    s += line.strip()
j = json.loads(s)

print("first attempt:", bookkeeping(False))
print("second attempt:", bookkeeping(True))
