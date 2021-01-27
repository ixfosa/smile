package string;

import com.sun.xml.internal.ws.server.ServerRtException;

/**
 * Created by ixfosa on 2021/1/27 14:51
 */
// Java 实例 - 字符串小写转大写
// 以下实例使用了 String toUpperCase() 方法将字符串从小写转为大写：
public class Demo9 {
    public static void main(String[] args) {
        String str = "ixfsoa";
        String strUpper = str.toUpperCase();
        String strLower = strUpper.toLowerCase();

        System.out.println("小写：" + strLower);
        System.out.println("大写：" + strUpper);
    }
}
