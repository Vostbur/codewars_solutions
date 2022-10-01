def my_first_interpreter(code):
    mem, output = 0, ""

    for c in code:
        match c:
            case "+":
                mem += 1
            case ".":
                output += chr(mem % 256)

    return output
