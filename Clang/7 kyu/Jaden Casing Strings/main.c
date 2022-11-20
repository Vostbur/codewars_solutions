#include <ctype.h>

char *to_jaden_case (char *jaden_case, const char *string)
{
    char prev = ' ';
    char *r = jaden_case;

    while(*string)
    {
        if (isspace(prev))
            *jaden_case = toupper(*string);
        else
            *jaden_case = *string;
        prev = *string;
        string++;
        jaden_case++;
    }
    *jaden_case = '\0';
    return r;
}

int main()
{
    return 0;
}