def longest_palindrome(s):
    s = s.lower()
    # Manachar’s Algorithm
    # https://www.hackerearth.com/practice/algorithms/string-algorithm/manachars-algorithm/tutorial/

    # ==== BEGIN ====
    new_string = '$#'
    for c in s:
        new_string += c + '#'
    new_string = new_string+'@'

    palindrome = [0]*len(new_string)
    c = r = 0

    for i in range(1, len(new_string)-1):
        if r > i:
            mirror = c - (i-c)
            palindrome[i] = min(r-i, palindrome[mirror])
        while new_string[i+1+palindrome[i]] == new_string[i-1-palindrome[i]]:
            palindrome[i] += 1
        if i+palindrome[i] > r:
            c = i
            r = i + palindrome[i]
    # ==== END ====

    # find the max length, remove special characters, and return
    max_center = max_length = 0
    for i in range(len(new_string)):
        if palindrome[i] > max_length:
            max_length = palindrome[i]
            max_center = i
    output = new_string[max_center-max_length: max_center+max_length]
    output = ''.join(output.split('#'))
    return output
