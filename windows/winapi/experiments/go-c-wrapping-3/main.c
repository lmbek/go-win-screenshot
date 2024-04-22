#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int Main() {
    int number, guess, attempts = 0;

    // Seed for random number generation
    srand(time(0));

    // Generate a random number between 1 and 10
    number = rand() % 10 + 1;

    printf("Welcome to the Number Guessing Game!\n");

    do {
        printf("Guess the number (between 1 and 10): ");
        scanf("%d", &guess);

        attempts++;

        if (guess > number) {
            printf("Too high! Try again.\n");
        } else if (guess < number) {
            printf("Too low! Try again.\n");
        } else {
            printf("Congratulations! You guessed the number in %d attempts.\n", attempts);
        }
    } while (guess != number);

    return 0;
}
