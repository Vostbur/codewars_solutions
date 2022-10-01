import unittest

import Kata


class TestKata(unittest.TestCase):
    def test_rgb(self):
        self.assertEqual(Kata.rgb(0, 0, 0), "000000")
        self.assertEqual(Kata.rgb(1, 2, 3), "010203")
        self.assertEqual(Kata.rgb(255, 255, 255), "FFFFFF")
        self.assertEqual(Kata.rgb(254, 253, 252), "FEFDFC")
        self.assertEqual(Kata.rgb(-20, 275, 125), "00FF7D")


if __name__ == "__main__":
    unittest.main()
