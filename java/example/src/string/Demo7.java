package string;

/**
 * Created by ixfosa on 2021/1/27 13:20
 */
// Java 实例 - 字符串分割

// Java 实例 Java 实例
// 以下实例使用了 split(string) 方法通过指定分隔符将字符串分割为数组：
public class Demo7 {
    public static void main(String[] args) {
        String str = "ixfosa-long-zhong";
        String[] temp = str.split("-");
        for (int i = 0; i < temp.length; i++) {
            System.out.println(temp[i]);
        }
        System.out.println("--------------fotEach--------------");
        for (String s : temp) {
            System.out.println(s);
        }
    }
}
