public class p1 {
    public static void main(String[] args) throws Exception {
        int n = 0;
        for (int i = 0; i < 1000; i++) {
            if (i % 3 == 0  || i % 5 == 0) {
                n += i;
            }
        }

        System.out.println(n);
    }
}