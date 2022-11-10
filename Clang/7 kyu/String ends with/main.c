#include <stdbool.h>
#include <stddef.h>
#include <string.h>

bool solution(const char *string, const char *ending)
{
    size_t str_len = strlen(string);
    size_t end_len = strlen(ending);

    if (str_len < end_len)
        return false;

    return false == strncmp(string + str_len - end_len, ending, end_len);
}

int main()
{
    solution("abc", "bc");
}