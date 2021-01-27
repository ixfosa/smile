package string;

//  字符串替换
public class Demo4 {
    public static void main(String[] args) {
        String str = "hello ixfosa";

        System.out.println(str.replace("h", "H")); // Hello ixfosa
        System.out.println(str.replaceFirst("o", "oo")); //helloo ixfosa
        System.out.println(str.replaceAll("o", "ooo")); // hellooo ixfooosa
    }
}
