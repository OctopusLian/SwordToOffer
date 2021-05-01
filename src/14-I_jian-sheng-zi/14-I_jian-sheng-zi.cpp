class Solution {
public:
    int cuttingRope(int n) {
        return n <= 3? n - 1 : pow(3, n / 3) * 4 / (4 - n % 3);
    }
};