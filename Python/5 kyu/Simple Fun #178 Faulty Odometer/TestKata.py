import unittest

import Kata


class TestKata(unittest.TestCase):
    def test_faulty_odometer(self):
        self.assertEqual(Kata.faulty_odometer(13), 12)
        self.assertEqual(Kata.faulty_odometer(15), 13)
        self.assertEqual(Kata.faulty_odometer(55), 40)
        self.assertEqual(Kata.faulty_odometer(2005), 1462)
        self.assertEqual(Kata.faulty_odometer(1500), 1053)
        self.assertEqual(Kata.faulty_odometer(999999), 531440)
        self.assertEqual(Kata.faulty_odometer(165826622), 69517865)


if __name__ == "__main__":
    unittest.main()
