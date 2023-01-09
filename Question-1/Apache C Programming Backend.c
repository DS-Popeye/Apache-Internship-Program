#include <stdio.h>
#include <stdlib.h>

int fibo(int a)
{
    if (a == 1)
        return 0;

    if (a == 2)
        return 1;

    int fib[a];
    fib[0] = 0;
    fib[1] = 1;
    for (int i = 2; i < a; i++)
        fib[i] = fib[i - 1] + fib[i - 2];

    return fib[a - 1];
}

typedef enum TypeTag
{
    ADD,
    SUB,
    MUL,
    FIBO,
} TypeTag;
typedef struct Node
{
    TypeTag type;
    int val;
} Node;

Node *constructNode(TypeTag _type, int _val)
{
    struct Node *tmp = (Node *)malloc(sizeof(struct Node));
    tmp->val = _val;
    tmp->type = _type;
    return tmp;
}
Node *add(int a, int b)
{
    int sum = a + b;
    return constructNode(ADD, a + b);
}
Node *addN(Node *a, Node *b)
{
    if (b == NULL)
        return constructNode(FIBO, fibo(a->val));
    return constructNode(ADD, a->val + b->val);
}
Node *sub(int a, int b)
{
    return constructNode(SUB, a - b);
}
Node *subN(Node *a, Node *b)
{
    if (b == NULL)
        return constructNode(FIBO, fibo(a->val));
    return constructNode(SUB, a->val - b->val);
}
Node *mul(int a, int b)
{
    return constructNode(MUL, a * b);
}
Node *mulN(Node *a, Node *b)
{
    if (b == NULL)
        return constructNode(FIBO, fibo(a->val));
    return constructNode(MUL, a->val * b->val);
}

Node *(*makeFunc(TypeTag tag))(int, int)
{
    if (tag == ADD)
    {
        return &add;
    }
    else if (tag == SUB)
    {
        return &sub;
    }
    else if (tag == MUL)
    {
        return &mul;
    }

    return &add;
}
Node *(*makeFun(TypeTag tag))(Node *, Node *)
{
    if (tag == ADD)
    {
        return &addN;
    }
    else if (tag == SUB)
    {
        return &subN;
    }
    else if (tag == MUL)
    {
        return &mulN;
    }

    return &addN;
}

Node *ABS(Node *tmp)
{
    tmp->val = (tmp->val < 0) ? -tmp->val : tmp->val;
    return tmp;
}

#define abs(tmp) ABS(tmp)

void calc(Node *tmp)
{

    printf("- ");
    switch (tmp->type)
    {
    case ADD:
        printf("add");
        break;
    case SUB:
        printf("sub");
        break;
    case MUL:
        printf("mul");
        break;
    case FIBO:
        printf("fibo");
        break;
    }
    printf(" : %d\n", tmp->val);
}

int main()
{

    Node *add = (*makeFunc(ADD))(10, 6);
    calc(add);
    Node *mul = (*makeFunc(MUL))(5, 4);
    calc(mul);
    Node *sub = (*makeFun(SUB))(add, mul);
    calc(sub);
    Node *fibo = (*makeFun(SUB))(abs(sub), NULL); // Get n-th from fibonacci numbers.
    calc(fibo);

    return 0;
}
