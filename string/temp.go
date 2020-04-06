package string

//确定一个字符串的所有字符是否全都不同。假使不允许使用额外的数据结构
//public static boolean isUniqueChars(String str) {
//8             if (str.length() > 128) {
//9                 return false;
//10             }
//11             int checker = 0;
//12             for (int i = 0; i < str.length(); i++) {
//13                 int val = str.charAt(i) - 'a';
//                     利用位运算 & 1和1 才为1
//14                 if ((checker & (1 << val)) > 0) return false;
//                     |只要有1就是1,标记下所有的1
//15                 checker |= (1 << val);
//16             }
//17             return true;
//18         }
//-------------------------------


//在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)
//func revreseString(str string) string {
//	var result string
//	strLen := len(str)
//	for i := 0; i < strLen; i++ {
//		result = result + fmt.Sprintf("%c", str[strLen-i-1])
//	}
//	return result

