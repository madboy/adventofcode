#!/usr/bin/env python3
import unittest
import day9

class TestDay9(unittest.TestCase):
    def test_no_decompression2(self):
        line = "ADVENT"
        decompressed_length = day9.decompress(0, len(line), line)
        self.assertEqual(decompressed_length, 6)

    def test_single_marker2(self):
        line = "A(1x5)BC"
        decompressed_length = day9.decompress(0, len(line), line)
        self.assertEqual(decompressed_length, 7)

    def test_single_marker_more_chars2(self):
        line = "(3x3)XYZ"
        decompressed_length = day9.decompress(0, len(line), line)
        self.assertEqual(decompressed_length, 9)

    def test_double_marker(self):
        line = "X(8x2)(3x3)ABCY"
        decompressed_length = day9.decompress(0, len(line), line)
        self.assertEqual(decompressed_length, 20)

    def test_many_internal_markers(self):
        line = "(27x12)(20x12)(13x14)(7x10)(1x12)A"
        decompressed_length = day9.decompress(0, len(line), line)
        self.assertEqual(decompressed_length, 241920)

    def test_read_marker(self):
        nc, r = day9.read_marker("3x2")
        self.assertEqual(nc, 3)
        self.assertEqual(r, 2)

    def test_read_marker_double_digits(self):
        nc, r = day9.read_marker("12x10")
        self.assertEqual(nc, 12)
        self.assertEqual(r, 10)

if __name__ == '__main__':
    unittest.main()
