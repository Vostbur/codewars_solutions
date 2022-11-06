import unittest

import Kata


class TestKata(unittest.TestCase):
    def test_exp_sum(self):
        self.assertEqual(Kata.exp_sum(-1), 0)
        self.assertEqual(Kata.exp_sum(0), 1)
        self.assertEqual(Kata.exp_sum(1), 1)
        self.assertEqual(Kata.exp_sum(2), 2)
        self.assertEqual(Kata.exp_sum(3), 3)
        self.assertEqual(Kata.exp_sum(4), 5)
        self.assertEqual(Kata.exp_sum(5), 7)
        self.assertEqual(Kata.exp_sum(20), 627)
        self.assertEqual(Kata.exp_sum(30), 5604)
        self.assertEqual(Kata.exp_sum(40), 37338)
        self.assertEqual(Kata.exp_sum(43), 63261)
        self.assertEqual(Kata.exp_sum(60), 966467)
        self.assertEqual(Kata.exp_sum(70), 4087968)
        self.assertEqual(Kata.exp_sum(90), 56634173)


if __name__ == "__main__":
    unittest.main()
