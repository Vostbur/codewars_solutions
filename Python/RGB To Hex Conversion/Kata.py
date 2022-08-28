def valid(x):
    return 0 if x < 0 else 255 if x > 255 else x


def rgb(r, g, b):
    return ("{:02X}" * 3).format(valid(r), valid(g), valid(b))
