public class Solution {
    public int MaxArea(int[] heights) {
        int s = 0;
        int e = heights.Length - 1;
        int max = -1;

        while (s < e) {
            int a = Math.Min(heights[s], heights[e]) * (e - s);
            if (a > max)
                max = a;

            if (heights[s] > heights[e])
                e--;
            else
                s++;
        }

        return max;
    }
}
