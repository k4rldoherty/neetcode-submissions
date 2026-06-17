public class Solution {
    public List<List<int>> ThreeSum(int[] nums) {
        List<List<int>> res = new();
        for (int i = 0; i < nums.Length; i++) {
            // Two Sum where the value
            // is minus the current index
            var target = -nums[i];
            Dictionary<int, int> map = new();

            // -1, 0, 1, 2, -1, -4
            // -1
            // target = 1
            // d = { 0, 1 }
            // j = 1
 
            for (int j = i + 1; j < nums.Length; j++) {
                if (map.TryGetValue(nums[j], out var val)) {
                    List<int> triplet = new List<int>() { nums[i], nums[j], val};
                    triplet.Sort(); 
                    if(res.Any(x => x.SequenceEqual(triplet))) continue;
                    res.Add(triplet);
                } else {
                    map.TryAdd(target - nums[j], nums[j]);
                }
            }
        }
        return res;
    }
}
