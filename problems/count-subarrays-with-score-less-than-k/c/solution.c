long long countSubarrays(int* nums, int numsSize, long long k) {
    long long res = 0;
    long long sum = 0;
    for (int left = 0, right = 0; right < numsSize; right++) {
        sum += nums[right];
        for (;sum * (right-left+1) >= k && left < right;) {
            sum -= nums[left];
            left++;
        }
        if (sum * (right-left+1) < k) {
            res += right-left+1;
        }
    }

    return res;
}