public class Solution {
    public int CharacterReplacement(string s, int k) {
        int r = 0;
        int l = 0;
        int res = 0;
        int mostOccurringLetter = 0;
        Dictionary<char, int> d = new();

        while(r < s.Length) {
            if(!d.TryAdd(s[r], 1)) {
                d[s[r]]++;
            }
            
            if(d[s[r]] > mostOccurringLetter) mostOccurringLetter = d[s[r]];

            // check window validity
            var changesToMakeRepeating = (r - l + 1) - mostOccurringLetter;

            if(changesToMakeRepeating <= k) {
                res = (r - l + 1) > res ? (r - l + 1) : res;
                // increase r
                r++;
            } else {
                // shrink window until valid again
                while(changesToMakeRepeating > k) {
                    d[s[l]]--;
                    l++;
                    changesToMakeRepeating = (r - l + 1) - mostOccurringLetter;
                }
                r++;
            }
        }

        return res;
    }
}