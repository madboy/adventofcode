#!/usr/bin/env python
import unittest
import day9

class TestDay9(unittest.TestCase):
    def test_no_decompression(self):
        decompressed_length = day9.decompress("ADVENT")
        self.assertEqual(decompressed_length, 6)

    def test_single_marker(self):
        decompressed_length = day9.decompress("A(1x5)BC")
        self.assertEqual(decompressed_length, 7)

    def test_single_marker_more_chars(self):
        decompressed_length = day9.decompress("(3x3)XYZ")
        self.assertEqual(decompressed_length, 9)

if __name__ == '__main__':
    unittest.main()
