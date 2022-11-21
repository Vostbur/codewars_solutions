#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <string.h>

char *binary_add(unsigned a, unsigned b, char *binary) {
  long long x = (long long)a + (long long)b;
  char *c, *byte = (char *)&x;
  bool flag = false;
  *binary = '\0';

  if (x == 0)
    strcat(binary, "0");
  else {
    for (int i = sizeof(x) - 1; i >= 0; i--) {
      for (int j = 7; j >= 0; j--) {
        c = ((byte[i] >> j) & 1) ? "1" : "0";
        if (c[0] == '1')
          flag = true;
        if (flag)
          strcat(binary, c);
      }
    }
  }

  return binary;
}

int main(void) {
  char b[33];
  printf("%s\n", binary_add(UINT_MAX, UINT_MAX, b));
  return 0;
}
