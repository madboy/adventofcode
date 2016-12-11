#!/usr/bin/env python3
import sys
import re

def check_abba(word):
    """Does word have an ABBA seuquence"""
    f = word[0]
    for s in word[1:]:
        # sequence must be differnt characters
        if f == s:
            continue
        if f + s + s + f in word:
            return True
        f = s
    return False


def supports_ipv7(hypernets, supernets):
    """For ipv7 there cannot be any ABBA in the hypernets,
    but there should be one in the supernets"""
    for hyper in hypernets:
        if check_abba(hyper):
            return False
    for supernet in supernets:
        if check_abba(supernet):
            return True
    return False

def check_aba(hyper, supernets):
    f = hyper[0]
    m = hyper[1]
    for l in hyper[2:]:
        if f == l:
            for supernet in supernets:
                if m+f+m in supernet:
                    return True
        f = m
        m = l
    return False

def supports_ssl(hypernets, supernets):
    """For ssl we only care about aba and bab patterns"""
    for hyper in hypernets:
        if check_aba(hyper, supernets):
            return True
    return False

IPV7 = 0
SSL = 0
for line in sys.stdin:
    l = line.strip()
    hypernet = re.compile("\[(\w+)\]")
    supernet = re.compile("[a-z]+")
    hypernets = hypernet.findall(line)
    # remove hyper addresses from the line
    for hyper in hypernets:
        line = line.replace(hyper, "")
    if supports_ipv7(hypernets, supernet.findall(line)):
        IPV7 += 1
    if supports_ssl(hypernets, supernet.findall(line)):
        SSL += 1

print("%d addresses supports IPV7" % IPV7)
print("%d addresses supports SSL" % SSL)
