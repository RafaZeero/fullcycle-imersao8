```C
#include <stdio.h>
#include <stdlib.h>

struct node {
int item;
struct node* left;
struct node* right;
};

// Inorder traversal
void inorderTraversal(struct node\* root) {
if (root == NULL) return;
inorderTraversal(root->left);
printf("%d ->", root->item);
inorderTraversal(root->right);
}
```

         9
      7    11
    5  8  10 12

5,7, 8, 9, 10, 11, 12
