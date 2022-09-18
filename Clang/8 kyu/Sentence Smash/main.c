#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *smash(const char *const words[], size_t nb_words)
{
    char joint[2];
    strcpy(joint, " ");
    size_t joint_length = strlen(joint);
    size_t total = (nb_words - 1) * joint_length + 1;

    for (size_t i = 0; i < nb_words; i++)
    {
        total += strlen(words[i]);
    }

    char *target, *result;
    target = result = calloc(sizeof(char), total);

    *target = '\0';
    for (size_t i = 0; i < nb_words; i++)
    {
        strcat(target, words[i]);
        target += strlen(words[i]);
        if (i < nb_words - 1)
        {
            strcat(target, joint);
            target += joint_length;
        }
    };
    return result;
}

int main()
{
    const char *const s[] = {"first", "second", "last"};
    char *cat = smash(s, 3);
    puts(cat);
    free(cat);
    return 0;
}
