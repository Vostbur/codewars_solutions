# def faulty_odometer(num):
#     digs = '012356789'
#     result = 0
#     for i, n in enumerate(str(num)[::-1]):
#         result += digs.index(n) * 9 ** i
#     return result

def faulty_odometer(n):
    if n == 0:
        return 0
    else:
        d = n % 10
        if d > 4:
            d -= 1
        return d + 9 * faulty_odometer(n//10)
