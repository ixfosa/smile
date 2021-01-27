package string;

import java.util.Locale;

/**
 * Created by ixfosa on 2021/1/27 15:21
 */

// Java 实例 - 字符串格式化
//以下实例演示了通过 format() 方法来格式化字符串，还可以指定地区来格式化：

public class Demo12 {
    public static void main(String[] args){
        double e = Math.E;
        System.out.format("%f%n", e); // 2.718282
        System.out.format(Locale.CHINA  , "%-10.4f%n%n", e);  //指定本地为中国（CHINA） 2.7183
    }
}
