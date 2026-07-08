public class Solution {
    public int Trap(int[] height) {
        int[] maxL = new int[height.Length];
        int[] maxR = new int[height.Length];
        int[] minLR = new int[height.Length];
        
        int maxSeen = 0;
        for(int i = 0; i < height.Length; i++) {
            maxL[i] = maxSeen;
            if(height[i] > maxSeen) maxSeen = height[i];
        }

        maxSeen = 0;
        for(int i = height.Length - 1; i >= 0; i--) {
            maxR[i] = maxSeen;
            if(height[i] > maxSeen) maxSeen = height[i];
        }

        for(int i = 0; i < height.Length; i++) {
            minLR[i] = Math.Min(maxL[i], maxR[i]);
        }

        int res = 0;
        for(int i = 0; i < height.Length; i++) {
            int trappable = minLR[i] - height[i];
            if(trappable > 0) res += trappable;
        }

        return res;
    }
}