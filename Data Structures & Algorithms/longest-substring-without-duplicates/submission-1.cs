public class Solution {
    public int LengthOfLongestSubstring(string s) {
        int res = 0;
        var sb = new StringBuilder();
        for(int i = 0; i < s.Length; i++)
        {
            sb.Append(s[i]);
            for(int j = i+1; j < s.Length; j++)
            {
                if(sb.ToString().Contains(s[j])) break;
                sb.Append(s[j]);
            }
            if(sb.Length > res) res = sb.Length;
            sb.Clear();
        }
        return res;
    }
}
