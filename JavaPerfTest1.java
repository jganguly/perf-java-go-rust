import java.lang.management.ManagementFactory;
import java.lang.management.MemoryMXBean;
import java.lang.management.MemoryUsage;
import java.lang.management.OperatingSystemMXBean;


public class JavaPerfTest1 {

    public static void main(String[] args) {
        elapsedTime();

        // Get the OperatingSystemMXBean
        // OperatingSystemMXBean osBean = ManagementFactory.getOperatingSystemMXBean();
        // // Print CPU Usage
        // System.out.println("CPU Usage: " + osBean.getSystemLoadAverage() + "%");

        // // Get the MemoryMXBean
        // MemoryMXBean memoryBean = ManagementFactory.getMemoryMXBean();
        // // Print Heap Memory Usage
        // MemoryUsage heapMemoryUsage = memoryBean.getHeapMemoryUsage();
        // System.out.println("Heap Memory Usage: " + formatMemoryUsage(heapMemoryUsage));

        // Print Non-Heap Memory Usage
        // MemoryUsage nonHeapMemoryUsage = memoryBean.getNonHeapMemoryUsage();
        // System.out.println("Non-Heap Memory Usage: " + formatMemoryUsage(nonHeapMemoryUsage));
    }


    private static String formatMemoryUsage(MemoryUsage memoryUsage) {
        long max = memoryUsage.getMax() / (1024 * 1024);
        long used = memoryUsage.getUsed() / (1024 * 1024);
        long committed = memoryUsage.getCommitted() / (1024 * 1024);

        return "Used: " + used + "MB, Committed: " + committed + "MB, Max: " + max + "MB";
    }

    private static void runYourCode() {
        // Allocate a large array to consume memory
        int n = 1000*1000*1000;
        int[] largeArray = new int[n];

        for (int i = 0; i < n; i++) {
            largeArray[i] = calculateProduct();        
        }
    }


    private static int calculateProduct() {
        int product = 1;
        for (int i = 1; i <= 100*1; i++) {
            product *= i;
        }
        return product;
    }

    private static void elapsedTime() {
        long startTime = System.currentTimeMillis();
        runYourCode();
        long endTime = System.currentTimeMillis();
        long elapsedTime = endTime - startTime;

        System.out.println("Elapsed Time: " + elapsedTime + " milliseconds");
    }
}
