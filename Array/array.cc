#include <stdio.h>

int main() {
    int A[5];
    
    A[0] = 1;
    A[1] = 2;
    A[2] = 3;
    A[3] = 4;
    A[4] = 5;

    printf("One-dimensional Array\n");
    for (int i = 0; i < 5; i++) {
        printf("%d\n", A[i]);
    }

    int B[5][5];
    B[0][0] = 1;
    B[0][1] = 2;
    B[0][2] = 3;
    B[0][3] = 4;
    B[0][4] = 5;

    B[1][0] = 6;
    B[1][1] = 7;
    B[1][2] = 8;
    B[1][3] = 9;
    B[1][4] = 10;

    printf("Two-dimensional array\n");
    for (int i = 0; i < 2; i++) {
        for (int j = 0; j < 5; j++) { // Corrected the loop variable from i to j
            printf("%d ", B[i][j]);
        }
        printf("\n");
    }

    return 0;
}
