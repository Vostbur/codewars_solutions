#include <stdbool.h>
#include <stdio.h>
#include <string.h>
#include <math.h>

bool narcissistic(int num)
{
    char str[20];
    sprintf(str, "%d", num);
    int r = 0, l = strlen(str);
    for (int i = 0; i < l; i++)
        r += pow((str[i] - '0'), l);
    return r == num;
}

int main()
{
    printf("%s", narcissistic(7) ? "true" : "false");
    printf("%s", narcissistic(371) ? "true" : "false");
    printf("%s", narcissistic(122) ? "true" : "false");
    printf("%s", narcissistic(4887) ? "true" : "false");
}