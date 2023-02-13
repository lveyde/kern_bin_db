#include <stdio.h>

void test_func() {
   printf("test func.\n");
}

inline static void inline_func2() {
    printf("Inline function call 2...\n");
    test_func();
    return;
}

inline static void inline_func() {
    printf("Inline function call...\n");
    inline_func2();
    return;
}

void trampoline() {
    inline_func();
    test_func();
}

int main() {
    printf("Hello World!\n");
    trampoline();
    return 0;
}
