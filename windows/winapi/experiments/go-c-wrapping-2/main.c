#include <stdio.h>
#include <unistd.h>

int Count() {
    for (int i = 1; i <= 10; i++) {
        printf("%d\n", i);

        sleep(1); // Pause for 1 second
    }
    return 0;
}
