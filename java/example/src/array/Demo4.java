package array;

import java.sql.Connection;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

/**
 * Created by ixfosa on 2021/1/27 20:44
 */

// 数组反转
//  Collections.reverse(ArrayList) 将数组进行反转：
public class Demo4 {
    public static void main(String[] args) {
        Object[] arr = {1, 2, 3, 4, 5, 6};
        System.out.println(Arrays.toString(reserve(arr))); // [6, 5, 4, 3, 2, 1]

    }

    public static Object[] reserve(Object[] arr) {
        List<Object> list = Arrays.asList(arr);
        Collections.reverse(list);
        return list.toArray();
    }
}
