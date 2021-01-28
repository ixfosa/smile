package array;

import com.sun.xml.internal.ws.streaming.TidyXMLStreamReader;

import java.util.Arrays;
import java.util.Collections;

/**
 * Created by ixfosa on 2021/1/28 16:01
 */

// 删除数组元素
// Java 的数组是固定长度的，无法直接删除，我们可以通过创建一个新数组，把原始数组中要保留的元素放到新数组中即可：
public class Demo11 {
    public static void main(String[] args) {
        int[] oldArr = new int[] {3, 4, 5, 6, 7}; // 原始数组
        System.out.println(Arrays.toString(delArrElem(oldArr, 0)));;

    }

    public static int[] delArrElem(int[] oldArr, int idx) {

        if (oldArr.length <= idx || idx < 0) {
            throw new RuntimeException("元素越界... ");
        }

        int[] newArr = new int[oldArr.length - 1];

        if (idx == 0) {
            System.arraycopy(oldArr, 1, newArr, 0, newArr.length);
        } else {
            System.arraycopy(oldArr, 0, newArr, 0, idx);
            System.arraycopy(oldArr, idx + 1, newArr, idx, newArr.length - idx);
        }
        return newArr;
    }

}
