package array;

import javax.swing.plaf.nimbus.NimbusLookAndFeel;
import java.util.Arrays;
import java.util.Collections;

/**
 * Created by ixfosa on 2021/1/27 21:09
 */

// 数组获取最大和最小值
// 通过 Collections 类的 Collections.max() 和 Collections.min() 方法来查找数组中的最大和最小值：
public class Demo6 {
    public static void main(String[] args) {
        Integer[] array = { 2, 5, -2, 6, -3, 8, 0, -7, -9, 4 };
        Integer max = Collections.max(Arrays.asList(array));
        Integer min = Collections.min(Arrays.asList(array));

        System.out.println("max: " + max); // 8
        System.out.println("min: " + min); // -9
    }
}
