package array;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

/**
 * Created by ixfosa on 2021/1/27 21:13
 */

// 数组合并
// 通过 List 类的 Arrays.toString () 方法和 List 类的 list.Addall(array1.asList(array2) 方法将两个数组合并为一个数组：
public class Demo7 {
    public static void main(String[] args) {
        Integer[] arr1 = {1, 2, 3};
        Integer[] arr2 = {1, 2, 3};
        Integer[] arr3 = {1, 2, 3};


        System.out.println(Arrays.toString(mergeArr(arr1, arr2, arr3)));
        System.out.println(Arrays.toString(mergeArr2(arr1, arr2, arr3)));
    }

    public static Object[] mergeArr(Object[] ... args) {
        int totalLength = 0;
        for (Object[] arg : args) {
            totalLength += arg.length;
        }

        Object[] target = new Object[totalLength];
        int len = 0;
        for (int i = 0; i < args.length; i++) {
            if (i == 0) {
                System.arraycopy(args[0], 0, target, 0, args[0].length);
            } else {
                len += args[i-1].length;
                System.arraycopy(args[0], 0, target, len, args[1].length);
            }
        }
        return target;
    }


    public static Object[] mergeArr2(Object[] ... args) {

        if (args.length == 0) {
            return null;
        }

        ArrayList<Object> list = new ArrayList<>();

        for (Object[] arg : args) {
            list.addAll(Arrays.asList(arg));
        }

        return list.toArray();
    }
}
