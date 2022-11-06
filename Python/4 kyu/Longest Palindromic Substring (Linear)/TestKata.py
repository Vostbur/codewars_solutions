import unittest

import Kata


class TestKata(unittest.TestCase):
    def test_exp_sum(self):
        self.assertEqual(Kata.longest_palindrome('babad'), 'bab')
        self.assertEqual(Kata.longest_palindrome('madam'), 'madam')
        self.assertEqual(Kata.longest_palindrome('dde'), 'dd')
        self.assertEqual(Kata.longest_palindrome('ababbab'), 'babbab')
        self.assertEqual(Kata.longest_palindrome('abababa'), 'abababa')
        self.assertEqual(Kata.longest_palindrome('banana'), 'anana')
        self.assertEqual(Kata.longest_palindrome('abba'), 'abba')
        self.assertEqual(Kata.longest_palindrome('cbbd'), 'bb')
        self.assertEqual(Kata.longest_palindrome('zz'), 'zz')
        self.assertEqual(Kata.longest_palindrome('dddd'), 'dddd')
        self.assertEqual(Kata.longest_palindrome(''), '')
        self.assertEqual(Kata.longest_palindrome(
            'abcdefghijklmnopqrstuvwxyz'), 'a')
        self.assertEqual(Kata.longest_palindrome(
            'ttaaftffftfaafatf'), 'aaftffftfaa')
        self.assertEqual(Kata.longest_palindrome('bbaaacc'), 'aaa')
        self.assertEqual(Kata.longest_palindrome('m'), 'm')


if __name__ == "__main__":
    unittest.main()
