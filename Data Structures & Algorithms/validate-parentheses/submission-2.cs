public class Solution {
    public bool IsValid(string s) {
		if (s.Length % 2 != 0) return false;
		Stack<char> stack = new();
		for(int i = 0; i < s.Length; i++)
		{
			switch(s[i])
			{
				case '{':
				case '(':
				case '[':
				stack.Push(s[i]);
				break;
				case '}':
				if(stack.Count == 0) return false;
				if (stack.Pop() != '{') return false;
				break;
				case ')':
				if(stack.Count == 0) return false;
				if (stack.Pop() != '(') return false;
				break;
				case ']':
				if(stack.Count == 0) return false;
				if (stack.Pop() != '[') return false;
				break;
				default:
				return false;
			}
		}
		return stack.Count == 0;
    }
}
