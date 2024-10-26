#include <iostream>
#include <vector>

using namespace std;

int main() {
    int num = 0;
    printf("input array size...\n");
    scanf("%d", &num);

    vector<int> arr(num);
    // 前缀和数组长度+1，可以简化边界处理，sums[0]表示没有值
    vector<int> sums(num + 1, 0);
    for (int i = 0; i < num; i++) {
        printf("input array member[%d]:", i);
        scanf("%d", &arr[i]);

        sums[i + 1] = sums[i] + arr[i];
    }

    int left, right;
    printf("input range left,right...\n");
    scanf("%d,%d", &left, &right);
    if (left < 0 || left > num-1 || right < 0 || right > num - 1 || left > right) {
        printf("left:%d or right:%d invalid!\n", left, right);
        return -1;
    }
    printf("sum:%d\n", sums[right + 1] - sums[left]);

    return 0;
}
