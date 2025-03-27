import java.util.Scanner;
import java.util.*;
import java.util.stream.Stream;
import java.util.HashMap;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;
 
public class Main {
    public static int count = 0;
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
        String input_str = in.nextLine();
        String[] tmp2 = input_str.split(" ");
        int[] nums = new int[tmp2.length];
        for (int j = 0; j < tmp2.length; j++) {  
            nums[j] = Integer.parseInt(tmp2[j]);
        }
        int m = nums[0];
        int n = nums[1];
    
        HashMap<Integer, ArrayList<Integer>> relations = new HashMap<>();
        HashMap<Integer, Integer> relation_keys = new HashMap<>();
        for (int i = 0; i < m; i++) {
            String tmp1 = in.nextLine();
            String[] tmp3 = tmp1.split(" ");
            if (!relations.containsKey(Integer.parseInt(tmp3[0]))) {
                relations.put((Integer.parseInt(tmp3[0])), new ArrayList<>());
            }
            if (!relation_keys.containsKey(Integer.parseInt(tmp3[0]))) {
                relation_keys.put((Integer.parseInt(tmp3[0])),Integer.parseInt(tmp3[1]));
            }
 
            if (tmp3[2].length() > 2) {
                String[] tmp4 = tmp3[2].substring(1, tmp3[2].length() - 1).split(",");
                Integer[] nums1 = new Integer[tmp4.length];
                for (int j = 0; j < tmp4.length; j++) {  
                    nums1[j] = Integer.parseInt(tmp4[j]);
                }
                relations.get(Integer.parseInt(tmp3[0])).addAll(Arrays.stream(nums1).collect(Collectors.toList()));
            }
        }
        LinkedList<Integer> arr1 = new LinkedList<>();
        arr1.add(n);
    
        while (arr1.size() > 0) {
            Integer num1 = arr1.pop();
            if (relation_keys.containsKey(num1)) {
                count += relation_keys.get(num1);
                arr1.addAll(relations.get(num1));
            }
        }

        System.out.println(count);
        return;
	}
}