#include <stdio.h>

int main(void)
{
    // for (/* A */; /* B */ ; /* C */)
    // {
        // тело цикла
    // }
    // A - какое-то начальное выражение если надо
    // B - условие которое будем проверять каждый раз перед итерацией
    // C - выражение которое выполним в конце итерации

    // /* A */
    // while (/* B */)
    // {
    //      тело цикла
    //      С
    // }


    int number = 0, result = 0;
    char c = '+';

    printf("Ну введи что-то...\nВвод: ");
    scanf("%c", &c);
    fflush(stdin);
    while (c != '0')
    {
        printf("Ты ввёл %c, это не ноль\n", c);
        printf("Ну введи что-то...\nВвод: ");
        scanf("%c", &c);
        fflush(stdin);
    }

    printf("Ну наконец-то ты ввёл ноль...\n");
    return 0;
}
