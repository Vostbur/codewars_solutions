#include <stdbool.h>
#include <string.h>

bool valid_braces(const char *braces)
{
    char tracer[strlen(braces)], last_char;
    int pointer = 0;

    for (int i = 0; i < strlen(braces); i++)
    {
        if (braces[i] == '(' || braces[i] == '{' || braces[i] == '[')
        {
            tracer[pointer++] = braces[i];
        }
        else
        {
            if (pointer == 0)
            {
                return false;
            }
            last_char = tracer[pointer - 1];
            if ((braces[i] == ')' && last_char == '(') || (braces[i] == '}' && last_char == '{') || (braces[i] == ']' && last_char == '['))
            {
                pointer--;
            }
            else
            {
                break;
            }
        }
    }
    return pointer == 0;
}

int main()
{
    valid_braces("(((");
    return 0;
}