package string;

import java.sql.SQLOutput;

/**
 * Created by ixfosa on 2021/1/27 12:46
 */
// Java 实例 - 字符串搜索
// 以下实例使用了 String 类的 indexOf() 方法在字符串中查找子字符串出现的位置，
// 如果存在返回字符串出现的位置（第一位为0），如果不存在返回 -1：
public class Demo6 {
    public static void main(String[] args) {
        String str = "ixfosa";
        searchString(str, "fo"); // 找到字符串: 2
    }
    public static void searchString(String str, String subStr) {
        int idx = str.indexOf(subStr);
        if ( idx == -1) {
            System.out.println("没有找到字符串: " + subStr);
        }
        System.out.println("找到字符串: " + idx);
    }
}
