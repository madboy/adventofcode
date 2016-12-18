import unittest
import day11


class TestDay11(unittest.TestCase):
    def test_allowed_incompatible(self):
        a = day11.Elevator.allowed("hydrogen", "thorium")
        self.assertEqual(a, False)

    def test_allowed_compatible(self):
        a = day11.Elevator.allowed("hydrogen-compatible",
                                   "hydrogen")
        self.assertEqual(a, True)

    def test_allowed_chips(self):
        a = day11.Elevator.allowed("hydrogen-compatible",
                                   "thorium-compatible")
        self.assertEqual(a, True)

if __name__ == '__main__':
    unittest.main()
