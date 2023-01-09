
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int mod = 1e9 + 7;

int arr[200000];
int Dynamic_Programming_Approach(int n)
{
    arr[0] = 0;
    arr[1] = 1;
    arr[2] = 2;
    for (int i = 3; i <= n; i++)
        arr[i] = arr[i - 3] + arr[i - 2];
    return arr[n];
}

// Declare T[][] as global matrix
int T[2000][2000];

// Result matrix
int result[2000][2000];

// Function to multiply two matrices
void mul_2(int K)
{
    // Create an auxiliary matrix to
    // store elements of the
    // multiplication matrix
    int temp[K + 1][K + 1];
    memset(temp, 0, sizeof temp);

    // Iterate over range [0, K]
    for (int i = 1; i <= K; i++)
    {

        for (int j = 1; j <= K; j++)
        {

            for (int k = 1; k <= K; k++)
            {

                // Update temp[i][j]
                temp[i][j] = (temp[i][j] + (T[i][k] * T[k][j]) % mod) % mod;
            }
        }
    }

    // Update the final matrix
    for (int i = 1; i <= K; i++)
    {
        for (int j = 1; j <= K; j++)
        {
            T[i][j] = temp[i][j];
        }
    }
}

// Function to multiply two matrices
void mul_1(int K)
{
    // Create an auxiliary matrix to
    // store elements of the
    // multiplication matrix
    int temp[K + 1][K + 1];
    memset(temp, 0, sizeof temp);

    // Iterate over range [0, K]
    for (int i = 1; i <= K; i++)
    {

        for (int j = 1; j <= K; j++)
        {

            for (int k = 1; k <= K; k++)
            {

                // Update temp[i][j]
                temp[i][j] = (temp[i][j] + (result[i][k] * T[k][j]) % mod) % mod;
            }
        }
    }

    // Update the final matrix
    for (int i = 1; i <= K; i++)
    {
        for (int j = 1; j <= K; j++)
        {
            result[i][j] = temp[i][j];
        }
    }
}

// Function to calculate matrix^n
// using binary exponentaion
void matrix_pow(int K, int n)
{
    // Initialize result matrix
    // and unity matrix
    for (int i = 1; i <= K; i++)
    {
        for (int j = 1; j <= K; j++)
        {
            if (i == j)
                result[i][j] = 1;
        }
    }

    while (n > 0)
    {
        if (n % 2 == 1)
            mul_1(K);
        mul_2(K);
        n /= 2;
    }
}

// Function to calculate nth term
// of general recurrence
int NthTerm(int F[], int C[], int K,
            int n)
{
    memset(T, 0, sizeof T);
    memset(result, 0, sizeof result);

    // Fill T[][] with appropriate value
    for (int i = 1; i <= K; i++)
        T[i][K] = C[K - i];

    for (int i = 1; i <= K; i++)
        T[i + 1][i] = 1;

    // Function Call to calculate T^n
    matrix_pow(K, n);

    int answer = 0;

    // Calculate nth term as first
    // element of F*(T^n)
    for (int i = 1; i <= K; i++)
    {
        answer += (F[i - 1] * result[i][1]) % mod;
        answer %= mod;
    }
    return answer;
}

// Driver Code
int Matrix_Exponentiation_Approach(int N)
{
    // Given Initial terms
    int F[] = {0, 1, 2};

    // Given coefficients
    int C[] = {0, 1, 1};

    // Given K
    int K = 3;

    int ans = NthTerm(F, C, K, N);

    return ans;
}

int Brute_Force_Approach(int n)
{
    // rescursion
    if (n == 0)
        return 0;
    else if (n == 1)
        return 1;
    else if (n == 2)
        return 2;
    else
        return Brute_Force_Approach(n - 3) + Brute_Force_Approach(n - 2);
}

int main()
{
    int ans;
    ans = Matrix_Exponentiation_Approach(10);

    printf("%d\n", ans);

    return 0;
}
