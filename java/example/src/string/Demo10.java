package string;

/**
 * Created by ixfosa on 2021/1/27 15:00
 */

// 测试两个字符串区域是否相等
// 以下实例使用了 regionMatches() 方法测试两个字符串区域是否相等：
public class Demo10 {
    public static void main(String[] args) {
        String str1 = "hello Java";
        String str2 = "java say hello";

        // str1.regionMatches(6, str2, 0, 4)
        // 表示将 str1 字符串从第 6 个字符"M"开始和 str2 字符串的第1个字符 "j" 开始逐个比较，
        // 共比较 4 对字符，由于字符串区分大小写，所以结果为false。
        //如果设置第一个参数为 true ，则表示忽略大小写区别，所以返回 true。

        boolean match1 = str1.regionMatches(6, str2, 0, 4);
        boolean match2 = str1.regionMatches(true, 6, str2, 0, 4);
        System.out.println(match1); //false
        System.out.println(match2); // true
    }
}
